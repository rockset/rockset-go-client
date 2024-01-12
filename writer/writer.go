package writer

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/fatih/structs"
	"github.com/rs/zerolog"

	"github.com/rockset/rockset-go-client/openapi"
)

const (
	// MaxDocumentCount is the max number of documents that can be sent in one batch
	MaxDocumentCount     = 1_000
	DefaultDocumentCount = MaxDocumentCount
	DefaultFlushInterval = time.Second
)

// Config is a struct containing the configurable parameters for a Writer
type Config struct {
	// BatchDocumentCount is the max document count to send in one request. Must be less than or equal to MaxDocumentCount
	BatchDocumentCount uint64

	// TODO implement document size restriction.

	// Workers number of workers to use, defaults to 1.
	Workers uint64
	// FlushInterval is the longest interval to wait before flushing buffered documents.
	FlushInterval time.Duration
	// ConversionFn is a function that converts the data in a Request into a map[string]any.
	// If not set, struct.Map() is used.
	ConversionFn func(any) (map[string]any, error)
}

// Validate validates a Config
func (c Config) Validate() error {
	if c.BatchDocumentCount < 1 {
		return fmt.Errorf("document count must be greater than 0")
	}
	if c.BatchDocumentCount > MaxDocumentCount {
		return fmt.Errorf("document count must be less than or equal to %d", MaxDocumentCount)
	}

	if c.FlushInterval == time.Duration(0) {
		return fmt.Errorf("batch size must be greater than 0")
	}

	return nil
}

// Writer is a helper package that writes documents to Rockset collections
type Writer struct {
	adder  DocumentAdder
	config Config
	stats  Stats

	buffers      map[string]map[string][]any
	bufferedDocs uint64

	m      sync.Mutex
	ticker *time.Ticker

	wg      sync.WaitGroup
	workers int

	addDocRequests chan addDocRequest
	writeRequests  chan Request
	stop           chan struct{}
}

// Request contains a single document to be written to Rockset
type Request struct {
	Workspace  string
	Collection string
	Data       any
}

// addDocRequest contains a list of documents to be written to Rockset
type addDocRequest struct {
	workspace  string
	collection string
	data       []any
}

// Stats holds counters for the documents written to Rockset
type Stats struct {
	DocumentCount uint64
	ErrorCount    uint64
}

// DocumentAdder is the interface used to write documents to Rockset, and is implemented by rockset.RockClient.
type DocumentAdder interface {
	AddDocuments(ctx context.Context, workspace, collection string, docs []interface{}) ([]openapi.DocumentStatus, error)
}

// New creates a new Writer
func New(conf Config, client DocumentAdder) (*Writer, error) {
	// set defaults if the zero value is used
	if conf.Workers == 0 {
		conf.Workers = 1
	}
	if conf.BatchDocumentCount == 0 {
		conf.BatchDocumentCount = DefaultDocumentCount
	}
	if conf.FlushInterval == 0 {
		conf.FlushInterval = DefaultFlushInterval
	}
	if conf.ConversionFn == nil {
		conf.ConversionFn = StructConversion
	}

	if err := conf.Validate(); err != nil {
		return nil, err
	}
	return &Writer{
		adder:          client,
		wg:             sync.WaitGroup{},
		writeRequests:  make(chan Request, conf.BatchDocumentCount),
		addDocRequests: make(chan addDocRequest, conf.Workers),
		stop:           make(chan struct{}, 1),
		buffers:        make(map[string]map[string][]any),
		config:         conf,
		stats:          Stats{},
	}, nil
}

// C returns the Request channel, which is used to send documents to be added to Rockset.
func (w *Writer) C() chan<- Request {
	return w.writeRequests
}

// Stop cleanly stops the Writer and flushes any buffered item, and closes the Request channel.
func (w *Writer) Stop() {
	w.stop <- struct{}{}
}

// Run starts the reader loop that gets write requests from the channel and batches them
// so the workers can add them to the collection(s). It starts the number Worker specified in the Config,
// but more can be started it if the number of collections written to is large.
//
// rs, _ := rockset.NewClient()
// w := writer.New(Config{}, rs.Documents)
// go w.Run(ctx)
// ...
// w.Stop()
// w.Wait()
//
//nolint:funlen
func (w *Writer) Run(ctx context.Context) {
	// TODO: should this panic if started more than once?
	w.wg.Add(1)
	defer w.wg.Done()

	// start workers
	for i := uint64(0); i <= w.config.Workers; i++ {
		go w.Worker(ctx)
	}

	log := zerolog.Ctx(ctx)

	// start the ticker and make sure we don't leave it behind
	w.ticker = time.NewTicker(w.config.FlushInterval)
	defer w.ticker.Stop()

	for {
		select {
		case <-w.stop:
			// we're being asked to cleanly stop and flush all data
			log.Debug().Msg("stopping writer loop...")
			close(w.writeRequests)
			log.Info().Int("counter", len(w.writeRequests)).Msgf("flushing documents")
			// this can add more than BatchSize to the payload, but as this is the last
			// time it runs it should be ok
			for r := range w.writeRequests {
				w.buffer(r, log)
			}
			w.flush()
			close(w.addDocRequests)
			return
		case <-ctx.Done():
			// the context was cancelled so we need to abort
			log.Debug().Msg("cancelling writer loop...")
			close(w.writeRequests)
			// drop queued write requests
			if len(w.writeRequests) > 0 {
				log.Warn().Int("counter", len(w.writeRequests)).Msgf("dropping write requests")
				w.m.Lock()
				w.stats.ErrorCount += uint64(len(w.writeRequests))
				w.m.Unlock()
				for range w.writeRequests {
					// just drop them on the floor
				}
			}
			close(w.addDocRequests)

			// check if we will drop buffered data
			var counter uint64
			for _, colls := range w.buffers {
				for _, data := range colls {
					counter += uint64(len(data))
				}
			}
			if counter > 0 {
				w.m.Lock()
				w.stats.ErrorCount += counter
				w.m.Unlock()
				log.Warn().Uint64("counter", counter).Msg("dropping buffered documents")
			}

			return
		case <-w.ticker.C:
			// the ticker has fired, and it is time to flush any buffered documents
			if w.bufferedDocs > 0 {
				log.Debug().Uint64("counter", w.bufferedDocs).
					Str("interval", w.config.FlushInterval.String()).Msg("flushing documents due to timeout")
				w.flush()
			}
		case r := <-w.writeRequests:
			// read a write request off the channel and buffer it
			w.buffer(r, log)
			if w.bufferedDocs >= w.config.BatchDocumentCount {
				log.Debug().Uint64("counter", w.bufferedDocs).Uint64("size", w.config.BatchDocumentCount).
					Msg("flushing documents due to batch size")
				w.flush()
				// reset ticker after flushing
				w.ticker.Reset(w.config.FlushInterval)
			}
		}
	}
}

// Worker runs a worker that writes batches of documents to the Rockset API.
// It needs to be started as a go routine, or it will block.
func (w *Writer) Worker(ctx context.Context) {
	w.wg.Add(1)
	defer w.wg.Done()

	log := zerolog.Ctx(ctx)

	w.m.Lock()
	index := w.workers
	w.workers++
	w.m.Unlock()

	log.Trace().Int("index", index).Msg("worker started")

	for dr := range w.addDocRequests {
		var errors uint64
		responses, err := w.adder.AddDocuments(ctx, dr.workspace, dr.collection, dr.data)
		// TODO recover from panic in AddDocuments
		var count = uint64(len(dr.data))
		if err != nil {
			w.m.Lock()
			w.stats.ErrorCount += count
			w.m.Unlock()

			// TODO: should errors be sent on an error channel instead?
			log.Error().Msgf("worker #%d error adding %d documents: %v", index, count, err)
			continue
		}

		for _, r := range responses {
			if r.GetStatus() != "ADDED" {
				errors++
				log.Error().Str("status", r.GetStatus()).Msg("incorrect status")
			}
		}
		log.Debug().Int("index", index).Uint64("counter", count-errors).Uint64("errors", errors).
			Str("workspace", dr.workspace).Str("collection", dr.collection).Msg("worker wrote documents")
		w.m.Lock()
		w.stats.DocumentCount += count - errors
		w.stats.ErrorCount += errors
		w.m.Unlock()
	}

	log.Debug().Int("index", index).Msg("worker done")
}

// Wait waits until the reader loop and all workers have finished
func (w *Writer) Wait() {
	w.wg.Wait()
}

// Stats returns a struct with document write statistics
func (w *Writer) Stats() Stats {
	w.m.Lock()
	stats := w.stats
	defer w.m.Unlock()
	return stats
}

// Workers returns the number of workers
func (w *Writer) Workers() int {
	w.m.Lock()
	defer w.m.Unlock()
	return w.workers
}

// buffer adds Request into a per workspace and collection buffer
func (w *Writer) buffer(r Request, log *zerolog.Logger) {
	_, found := w.buffers[r.Workspace]
	if !found {
		w.buffers[r.Workspace] = make(map[string][]any)
	}

	doc, err := w.config.ConversionFn(r.Data)
	if err != nil {
		log.Error().Err(err).Msg("error preparing document")
		w.m.Lock()
		w.stats.ErrorCount++
		w.m.Unlock()
		return
	}

	w.buffers[r.Workspace][r.Collection] = append(w.buffers[r.Workspace][r.Collection], doc)
	w.bufferedDocs++
}

// flush takes the buffered documents and sends them to the workers
func (w *Writer) flush() {
	for ws, collections := range w.buffers {
		for collection, data := range collections {
			w.addDocRequests <- addDocRequest{
				workspace:  ws,
				collection: collection,
				data:       data,
			}
		}
	}

	clear(w.buffers)
	w.bufferedDocs = 0
}

// JSONConversion converts any to a map[string]any using json.Marshal and then json.Unmarshal
func JSONConversion(a any) (map[string]any, error) {
	data, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}

	var m map[string]any
	if err = json.Unmarshal(data, &m); err != nil {
		return nil, err
	}

	return m, nil
}

// StructConversion converts any to a map[string]any using structs.Map, which doesn't honor json tags
func StructConversion(a any) (map[string]any, error) {
	return structs.Map(a), nil
}
