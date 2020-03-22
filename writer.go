package rockset

import (
	api "github.com/rockset/rockset-go-client/lib/go"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type WriterConfig struct {
	BufferSize    uint64
	BatchSize     uint64
	FlushInterval time.Duration
	Logger        *log.Logger
}

type Writer struct {
	wg      sync.WaitGroup
	ch      chan WriteRequest
	stop    chan bool
	buffers map[string]map[string][]interface{}
	counter uint64
	config  WriterConfig
	stats   WriteStats
	log     *log.Logger
}

type WriteRequest struct {
	Workspace  string
	Collection string
	Data       interface{}
}

type WriteStats struct {
	DocumentCount uint64
	ErrorCount    uint64
}

func NewWriter(conf WriterConfig) *Writer {
	logger := conf.Logger
	if logger == nil {
		logger = log.New(os.Stderr, "", log.LstdFlags)
	}
	return &Writer{
		wg:   sync.WaitGroup{},
		ch:   make(chan WriteRequest, conf.BufferSize),
		stop: make(chan bool, 1),
		buffers: make(map[string]map[string][]interface{}),
		config: conf,
		stats:  WriteStats{},
		log:    logger,
	}
}

func (w *Writer) Write(wr WriteRequest) {
	w.ch <- wr
}

func (w *Writer) C() chan WriteRequest {
	return w.ch
}

type DocAdder interface {
	Add(string, string, api.AddDocumentsRequest) (api.AddDocumentsResponse, *http.Response, error)
}

func (w *Writer) Run(da DocAdder) {
	w.wg.Add(1)
	defer w.wg.Done()

	ticker := time.NewTicker(w.config.FlushInterval)
	defer ticker.Stop()

	for {
		select {
		case flush := <-w.stop:
			w.log.Printf("stopping writer loop...")
			close(w.ch)
			if flush {
				w.log.Printf("flushing %d items", len(w.ch))
				for r := range w.ch {
					w.buffer(r)
				}
				w.flush(da)
			}
			return
		case <-ticker.C:
			// TODO: if we're writing at full speed, the ticker based flush slows it down
			// it should only write documents when they have been sitting in the buffer too long
			w.flush(da)
		case r := <-w.C():
			w.buffer(r)
			if w.counter >= w.config.BufferSize {
				w.flush(da)
			}
		}
	}
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

func (w *Writer) flush(da DocAdder) {
	for ws, colls := range w.buffers {
		for coll, data := range colls {
			req := api.AddDocumentsRequest{Data: data}
			// TODO: a possible optimization would be using a go routine for the request
			start := time.Now()
			_, _, err := da.Add(ws, coll, req)
			end := time.Now()
			count := len(data)
			if err == nil {
				w.log.Printf("wrote %d documents in %s", count, end.Sub(start))
				w.stats.DocumentCount += uint64(len(data))
				continue
			}
			w.stats.ErrorCount += uint64(len(data))
			// TODO: should errors be sent on an error channel instead?
			if e, ok := AsRocksetError(err); ok {
				w.log.Printf("error adding %d documents: %s", count, e.Message)
			} else {
				w.log.Printf("error adding %d documents: %v", count, err)
			}
		}
	}
	// TODO: benchmark to see if it is faster to delete keys as they are flushed instead of creating a new map
	w.buffers = make(map[string]map[string][]interface{})
	w.counter = 0
}

func (w *Writer) Stop(flush bool) {
	w.stop <- flush
}

func (w *Writer) Wait() {
	w.wg.Wait()
}

func (w *Writer) Stats() WriteStats {
	return w.stats
}
