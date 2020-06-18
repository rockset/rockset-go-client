package rockset

import (
	api "github.com/rockset/rockset-go-client/lib/go"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

// WriterConfig is a struct containing the configurable parameters for the Writer
type WriterConfig struct {
	BufferSize    uint64
	BatchSize     uint64
	Workers       uint64
	FlushInterval time.Duration
	Logger        *log.Logger
}

// Writer is
type Writer struct {
	config         WriterConfig
	stats          WriteStats
	addDocRequests chan addDocRequest
	buffers        map[string]map[string][]interface{}
	counter        uint64
	log            *log.Logger
	m              sync.Mutex
	stop           chan bool
	timeout        *time.Timer
	timedout       bool
	wg             sync.WaitGroup
	workers        int
	writeRequests  chan WriteRequest
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
func NewWriter(conf WriterConfig) *Writer {
	logger := conf.Logger
	if logger == nil {
		logger = log.New(os.Stderr, "", log.LstdFlags|log.Lmicroseconds)
	}
	return &Writer{
		wg:             sync.WaitGroup{},
		writeRequests:  make(chan WriteRequest, conf.BufferSize),
		addDocRequests: make(chan addDocRequest, conf.Workers),
		stop:           make(chan bool, 1),
		buffers:        make(map[string]map[string][]interface{}),
		config:         conf,
		stats:          WriteStats{},
		log:            logger,
	}
}

// C returns the WriteRequest channel
func (w *Writer) C() chan WriteRequest {
	return w.writeRequests
}

// Adder is the interface used to write documents to Rockset.
type Adder interface {
	Add(string, string, api.AddDocumentsRequest) (api.AddDocumentsResponse, *http.Response, error)
}

// Run starts the reader loop that gets write requests from the channel and batches them
// so the workers can add them to the collection. At least one Worker is needs to be started.
//
// rs, err := rockset.NewClient()
// w := NewWriter(WriterConfig{})
// w.Run(rs.Documents)
func (w *Writer) Run(da Adder) {
	w.wg.Add(1)
	defer w.wg.Done()

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
		case flush := <-w.stop:
			w.log.Printf("stopping writer loop...")
			close(w.writeRequests)
			if flush {
				w.log.Printf("flushing %d items", len(w.writeRequests))
				// this can add more then BatchSize to the payload
				for r := range w.writeRequests {
					w.buffer(r)
				}
				w.flush(da)
			}
			close(w.addDocRequests)
			return
		case <-w.timeout.C:
			w.log.Printf("flushing %d documents due to timeout", w.counter)
			w.flush(da)
		case r := <-w.C():
			w.buffer(r)
			// only start the timer once we have data in the buffer
			if w.timedout {
				w.timedout = false
				w.timeout = time.NewTimer(w.config.FlushInterval)
			}
			if w.counter >= w.config.BatchSize {
				w.flush(da)
			}
		}
	}
}

// Worker runs a worker that writes batches of documents to the Rockset API.
// It needs to be started as a go routine or it will block.
func (w *Writer) Worker(da Adder) {
	w.wg.Add(1)
	defer w.wg.Done()

	w.m.Lock()
	index := w.workers
	w.workers++
	w.m.Unlock()

	w.log.Printf("worker #%d started", index)

	for dr := range w.addDocRequests {
		req := api.AddDocumentsRequest{Data: dr.data}
		start := time.Now()
		_, _, err := da.Add(dr.workspace, dr.collection, req)
		end := time.Now()
		count := len(dr.data)
		if err == nil {
			w.log.Printf("worker #%d wrote %d documents in %s", index, count, end.Sub(start))
			w.m.Lock()
			w.stats.DocumentCount += uint64(len(dr.data))
			w.m.Unlock()
			continue
		}

		// TODO: handle 429 from the API

		w.m.Lock()
		w.stats.ErrorCount += uint64(len(dr.data))
		w.m.Unlock()
		// TODO: should errors be sent on an error channel instead?
		if e, ok := AsRocksetError(err); ok {
			w.log.Printf("worker #%d error adding %d documents: %s", index, count, e.Message)
		} else {
			w.log.Printf("worker #%d error adding %d documents: %v", index, count, err)
		}

	}
	w.log.Printf("worker #%d done", index)
}

// Stop signals the reader loop to stop, and then flushes any remaining data
func (w *Writer) Stop(flush bool) {
	w.stop <- flush
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

func (w *Writer) buffer(r WriteRequest) {
	ws, found := w.buffers[r.Workspace]
	if !found {
		ws = make(map[string][]interface{})
		w.buffers[r.Workspace] = ws
	}
	ws[r.Collection] = append(w.buffers[r.Workspace][r.Collection], r.Data)
	w.counter++
}

func (w *Writer) flush(da Adder) {
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

type addDocRequest struct {
	workspace  string
	collection string
	data       []interface{}
}
