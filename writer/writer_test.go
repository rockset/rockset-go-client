package writer_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/internal/test"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
	"github.com/rockset/rockset-go-client/writer"
)

type testObject struct {
	Foo uint64
	Bar string
}

// TODO replace with counterfeiter
type fakeAdder struct {
	m     sync.Mutex
	count int
}

func (f *fakeAdder) AddDocuments(_ context.Context, _, collection string,
	docs []interface{}) ([]openapi.DocumentStatus, error) {
	f.m.Lock()
	f.count = +len(docs)
	f.m.Unlock()
	time.Sleep(10 * time.Millisecond)

	var added = "ADDED"
	var result = make([]openapi.DocumentStatus, len(docs))
	for i := 0; i < len(docs); i++ {
		result[i] = openapi.DocumentStatus{
			Collection: &collection,
			Status:     &added,
		}
	}

	return result, nil
}

type WriterSuite struct {
	suite.Suite
}

func TestWriterSuite(t *testing.T) {
	s := WriterSuite{}
	suite.Run(t, &s)
}

func (s *WriterSuite) TestWriter() {
	ctx := test.Context()
	flushInterval := time.Millisecond * 50
	c := writer.Config{
		BatchDocumentCount: 30,
		FlushInterval:      flushInterval,
		ConversionFn:       writer.JSONConversion,
	}
	fa := &fakeAdder{}
	w, err := writer.New(c, fa)
	s.Require().NoError(err)

	go w.Run(ctx)

	// write enough documents that the writer will have to flush using the ticker
	const writeCount uint64 = 100
	for i := uint64(0); i < writeCount; i++ {
		w.C() <- writer.Request{
			Workspace:  "workspace",
			Collection: "collection",
			Data:       testObject{Foo: i},
		}
	}

	// wait for the ticker to flush the documents
	time.Sleep(flushInterval + time.Millisecond*20)
	s.Equal(writeCount, w.Stats().DocumentCount)

	w.Stop()
	w.Wait()

	s.Equal(uint64(0), w.Stats().ErrorCount)
	s.Equal(writeCount, w.Stats().DocumentCount)
}

func (s *WriterSuite) TestCancellation() {
	ctx, cancel := context.WithCancel(test.Context())
	c := writer.Config{
		BatchDocumentCount: 30,
		FlushInterval:      time.Millisecond * 20,
	}
	fa := &fakeAdder{}
	w, err := writer.New(c, fa)
	s.Require().NoError(err)

	go w.Run(ctx)

	w.C() <- writer.Request{
		Workspace:  "workspace",
		Collection: "collection",
		Data:       testObject{},
	}
	time.Sleep(30 * time.Millisecond)
	w.C() <- writer.Request{
		Workspace:  "workspace",
		Collection: "collection",
		Data:       testObject{},
	}

	cancel()
	w.Wait()

	s.Equal(uint64(1), w.Stats().ErrorCount)
	s.Equal(uint64(1), w.Stats().DocumentCount)
}

func (s *WriterSuite) TestShutdown() {
	ctx := test.Context()
	c := writer.Config{
		BatchDocumentCount: 30,
		FlushInterval:      time.Millisecond * 500,
	}
	fa := &fakeAdder{}
	w, err := writer.New(c, fa)
	s.Require().NoError(err)

	go w.Run(ctx)

	const writeCount uint64 = 100
	for i := uint64(0); i < writeCount; i++ {
		w.C() <- writer.Request{
			Workspace:  "workspace",
			Collection: "collection",
			Data:       testObject{Foo: i},
		}
	}

	w.Stop()
	w.Wait()

	s.Equal(uint64(0), w.Stats().ErrorCount)
	s.Equal(writeCount, w.Stats().DocumentCount)
}

func BenchmarkWriter10(b *testing.B) {
	c := writer.Config{
		FlushInterval: time.Millisecond * 500,
	}

	benchmarkWriter(b, c)
}

func BenchmarkWriter50(b *testing.B) {
	c := writer.Config{
		FlushInterval: time.Millisecond * 500,
	}

	benchmarkWriter(b, c)
}

func BenchmarkWriter100(b *testing.B) {
	c := writer.Config{
		FlushInterval: time.Millisecond * 500,
	}

	benchmarkWriter(b, c)
}

func benchmarkWriter(b *testing.B, c writer.Config) {
	ctx := test.Context()
	fa := &fakeAdder{}
	w, err := writer.New(c, fa)
	require.NoError(b, err)

	go w.Run(ctx)

	doc := writer.Request{
		Workspace:  "workspace",
		Collection: "collection",
		Data:       testObject{},
	}

	for i := 0; i < b.N; i++ {
		w.C() <- doc
	}

	w.Stop()
	w.Wait()
}

// go test -v -run TestWriterIntegration -cpuprofile=cpu.out
// go tool pprof -http=:8080 cpu.out

func TestWriterIntegration(t *testing.T) {
	test.SkipUnlessIntegrationTest(t)
	ctx := test.Context()

	rc, err := rockset.NewClient()
	require.Nil(t, err)

	c := writer.Config{
		FlushInterval: time.Millisecond * 200,
		Workers:       4,
	}
	w, err := writer.New(c, rc)
	require.NoError(t, err)

	go w.Run(ctx)

	const writeCount uint64 = 10_000

	name := test.RandomName("write")
	t.Logf("writing %d documents to %s", writeCount, name)
	t0 := time.Now()

	for i := uint64(0); i < writeCount; i++ {
		w.C() <- writer.Request{
			Workspace:  "commons",
			Collection: "writetest",
			Data:       testObject{i, name},
		}
	}

	w.Stop()
	w.Wait()

	t.Logf("wrote %d documents in %s", writeCount, time.Since(t0))

	assert.Equal(t, uint64(0), w.Stats().ErrorCount)
	assert.Equal(t, writeCount, w.Stats().DocumentCount)

	// try to get the documents written back from Rockset
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		resp, err := rc.Query(ctx, "SELECT COUNT(*) AS cnt FROM commons.writetest WHERE Bar = :name",
			option.WithParameter("name", "string", name))
		require.NoError(t, err)

		require.Len(t, resp.Results, 1)
		r := resp.Results[0]
		cnt := r["cnt"]
		count := cnt.(float64)
		if uint64(count) == writeCount {
			return
		}
	}

	t.Errorf("failed to get %d records", writeCount)
}
