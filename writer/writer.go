package writer

import (
	"context"
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

	buffers      map[string]map[string][]interface{}
	bufferedDocs uint64

	m       sync.Mutex
	timeout *time.Timer

	wg      sync.WaitGroup
	workers int

	addDocRequests chan addDocRequest
	writeRequests  chan Request
	stop           chan struct{}
}

// Request contains the data to be written to the Rockset collection in the request
type Request struct {
	Workspace  string
	Collection string
	Data       interface{}
}

type addDocRequest struct {
	workspace  string
	collection string
	data       []interface{}
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

	if err := conf.Validate(); err != nil {
		return nil, err
	}
	return &Writer{
		adder:          client,
		wg:             sync.WaitGroup{},
		writeRequests:  make(chan Request, conf.BatchDocumentCount),
		addDocRequests: make(chan addDocRequest, conf.Workers),
		stop:           make(chan struct{}, 1),
		buffers:        make(map[string]map[string][]interface{}),
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

	// start one worker
	var i uint64
	for ; i <= w.config.Workers; i++ {
		go w.Worker(ctx)
	}

	log := zerolog.Ctx(ctx)

	// start the timer and make sure we don't leave it behind
	w.timeout = time.NewTimer(w.config.FlushInterval)
	defer w.timeout.Stop()

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
				w.buffer(r)
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
		case <-w.timeout.C:
			// the interval timer has fired, and it is time to flush any buffered documents
			if w.bufferedDocs > 0 {
				log.Debug().Uint64("counter", w.bufferedDocs).Msg("flushing documents due to timeout")
				w.flush()
			}
			// restart the timer
			w.timeout.Reset(w.config.FlushInterval)
		case r := <-w.writeRequests:
			// read a write request off the channel and buffer it
			w.buffer(r)
			if w.bufferedDocs >= w.config.BatchDocumentCount {
				log.Debug().Uint64("counter", w.bufferedDocs).Msg("flushing documents due to batch size")
				w.flush()
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
func (w *Writer) buffer(r Request) {
	_, found := w.buffers[r.Workspace]
	if !found {
		w.buffers[r.Workspace] = make(map[string][]interface{})
	}

	w.buffers[r.Workspace][r.Collection] = append(w.buffers[r.Workspace][r.Collection], structs.Map(r.Data))
	w.bufferedDocs++
}

func (w *Writer) flush() {
	w.timeout.Stop()
	for ws, colls := range w.buffers {
		for coll, data := range colls {
			w.addDocRequests <- addDocRequest{
				workspace:  ws,
				collection: coll,
				data:       data,
			}
		}
	}
	// TODO use clear() once we update to go1.21
	w.buffers = make(map[string]map[string][]interface{})
	w.bufferedDocs = 0
}
