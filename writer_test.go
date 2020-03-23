package rockset_test

import (
	"github.com/rockset/rockset-go-client"
	api "github.com/rockset/rockset-go-client/lib/go"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"testing"
	"time"
)

type fakeAdder struct {
	count int
}

func (f *fakeAdder) Add(ws, coll string, request api.AddDocumentsRequest) (api.AddDocumentsResponse, *http.Response, error) {
	f.count++
	time.Sleep(10 * time.Millisecond)
	return api.AddDocumentsResponse{}, nil, nil
}

func TestWriter(t *testing.T) {
	c := rockset.WriterConfig{
		BufferSize:    30,
		BatchSize:     7,
		FlushInterval: time.Millisecond * 50,
	}
	w := rockset.NewWriter(c)

	go w.Run(&fakeAdder{})

	const writeCount uint64 = 100

	for i := uint64(0); i < writeCount; i++ {
		w.C() <- rockset.WriteRequest{
			Workspace:  "workspace",
			Collection: "collection",
			Data:       strconv.FormatUint(i, 10),
		}
	}

	w.Stop(true)
	w.Wait()

	assert.Equal(t, uint64(0), w.Stats().ErrorCount)
	assert.Equal(t, writeCount, w.Stats().DocumentCount)
}

func TestWriterTimeout(t *testing.T) {
	c := rockset.WriterConfig{
		BufferSize:    30,
		BatchSize:     7,
		FlushInterval: time.Millisecond * 10,
	}
	w := rockset.NewWriter(c)

	go w.Run(&fakeAdder{})

	w.C() <- rockset.WriteRequest{
		Workspace:  "workspace",
		Collection: "collection",
		Data:       "",
	}
	time.Sleep(20 * time.Millisecond)

	w.Stop(true)
	w.Wait()

	assert.Equal(t, uint64(0), w.Stats().ErrorCount)
	assert.Equal(t, uint64(1), w.Stats().DocumentCount)

}

var discardLogger = log.New(ioutil.Discard, "", log.LstdFlags)

func BenchmarkWriter10(b *testing.B) {
	c := rockset.WriterConfig{
		BufferSize:    1000,
		BatchSize:     10,
		FlushInterval: time.Millisecond * 500,
		Logger:        discardLogger,
	}

	benchmarkWriter(b, c)
}

func BenchmarkWriter50(b *testing.B) {
	c := rockset.WriterConfig{
		BufferSize:    1000,
		BatchSize:     50,
		FlushInterval: time.Millisecond * 500,
		Logger:        discardLogger,
	}

	benchmarkWriter(b, c)
}

func BenchmarkWriter100(b *testing.B) {
	c := rockset.WriterConfig{
		BufferSize:    1000,
		BatchSize:     100,
		FlushInterval: time.Millisecond * 500,
		Logger:        discardLogger,
	}

	benchmarkWriter(b, c)
}

func benchmarkWriter(b *testing.B, c rockset.WriterConfig) {
	w := rockset.NewWriter(c)

	go w.Run(&fakeAdder{})

	doc := rockset.WriteRequest{
		Workspace:  "workspace",
		Collection: "collection",
		Data:       "foo",
	}

	for i := 0; i < b.N; i++ {
		w.C() <- doc
	}

	w.Stop(true)
	w.Wait()
}
