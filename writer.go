package rockset

import (
	"context"
	"sync"
	"time"

	"github.com/fatih/structs"
	"github.com/rs/zerolog"

	"github.com/rockset/rockset-go-client/openapi"
)

// WriterConfig is a struct containing the configurable parameters for a Writer
type WriterConfig struct {
	BufferSize    uint64
	BatchSize     uint64
	Workers       uint64
	FlushInterval time.Duration
}

// Writer is a helper that writes structs to Rockset collections
type Writer struct {
	adder          DocumentAdder
	config         WriterConfig
	stats          WriteStats
	addDocRequests chan addDocRequest
	buffers        map[string]map[string][]interface{}
	counter        uint64
	m              sync.Mutex
	timeout        *time.Timer
	timedout       bool
	wg             sync.WaitGroup
	workers        int
	writeRequests  chan WriteRequest
	stop           chan struct{}
}

// WriteRequest contains the data to be written to a Rockset collection
type WriteRequest struct {
	Workspace  string
	Collection string
	Data       interface{}
}

// WriteStats holds counters for the documents written to Rockset
type WriteStats struct {
	DocumentCount uint64
	ErrorCount    uint64
}

// NewWriter creates a new Writer
func NewWriter(conf WriterConfig, adder DocumentAdder) *Writer {
	return &Writer{
		adder:          adder,
		wg:             sync.WaitGroup{},
		writeRequests:  make(chan WriteRequest, conf.BufferSize),
		addDocRequests: make(chan addDocRequest, conf.Workers),
		stop:           make(chan struct{}, 1),
		buffers:        make(map[string]map[string][]interface{}),
		config:         conf,
		stats:          WriteStats{},
	}
}

// C returns the WriteRequest channel.
func (w *Writer) C() chan<- WriteRequest {
	return w.writeRequests
}

// DocumentAdder is the interface used to write documents to Rockset.
type DocumentAdder interface {
	AddDocuments(ctx context.Context, workspace, collection string, docs []interface{}) ([]openapi.DocumentStatus, error)
}

// Stop cleanly stops the Writer and flushes any buffered item, and closes the WriteRequest channel.
func (w *Writer) Stop() {
	w.stop <- struct{}{}
}

// Run starts the reader loop that gets write requests from the channel and batches them
// so the workers can add them to the collection. At least one Worker is needs to be started.
//
// rs, err := rockset.NewClient()
// w := NewWriter(WriterConfig{})
// w.Run(rs.Documents)
func (w *Writer) Run(ctx context.Context) {
	// TODO: should this panic if started more than once?
	w.wg.Add(1)
	defer w.wg.Done()

	log := zerolog.Ctx(ctx)

	// this is done to have a timer available for the select
	w.timeout = time.NewTimer(0)
	<-w.timeout.C
	w.timedout = true

	defer func() {
		if !w.timedout {
			w.timeout.Stop()
		}
	}()

	for {
		select {
		case <-w.stop:
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

			// drop buffered data
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
				log.Warn().Uint64("counter", counter).Msg("dropping buffer")
			}

			return
		case <-w.timeout.C:
			log.Debug().Uint64("counter", w.counter).Msg("flushing documents due to timeout")
			w.flush()
		case r := <-w.writeRequests:
			w.buffer(r)
			// only start the timer once we have data in the buffer
			if w.timedout {
				w.timedout = false
				w.timeout = time.NewTimer(w.config.FlushInterval)
			}
			if w.counter >= w.config.BatchSize {
				log.Debug().Uint64("counter", w.counter).Msg("flushing documents due to batch size")
				w.flush()
			}
		}
	}
}

type addDocRequest struct {
	workspace  string
	collection string
	data       []interface{}
}

// Worker runs a worker that writes batches of documents to the Rockset API.
// It needs to be started as a go routine or it will block.
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
		var count uint64 = uint64(len(dr.data))
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
func (w *Writer) Stats() WriteStats {
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

// buffer adds WriteRequest into a per workspace and collection buffer
func (w *Writer) buffer(r WriteRequest) {
	wsBuffer, found := w.buffers[r.Workspace]
	if !found {
		wsBuffer = make(map[string][]interface{})
		w.buffers[r.Workspace] = wsBuffer
	}

	w.buffers[r.Workspace][r.Collection] = append(w.buffers[r.Workspace][r.Collection], structs.Map(r.Data))
	w.counter++
}

func (w *Writer) flush() {
	w.timeout.Stop()
	w.timedout = true
	for ws, colls := range w.buffers {
		for coll, data := range colls {
			w.addDocRequests <- addDocRequest{
				workspace:  ws,
				collection: coll,
				data:       data,
			}
		}
	}
	// TODO: benchmark to see if it is faster to delete keys as they are flushed instead of creating a new map
	w.buffers = make(map[string]map[string][]interface{})
	w.counter = 0
}
