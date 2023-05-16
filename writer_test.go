package rockset_test

import (
	"context"
	"github.com/rockset/rockset-go-client/option"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/openapi"
)

type testObject struct {
	Foo uint64
	Bar string
}

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
	ctx := testCtx()
	c := rockset.WriterConfig{
		BatchDocumentCount: 30,
		FlushInterval:      time.Millisecond * 50,
	}
	fa := &fakeAdder{}
	w, err := rockset.NewWriter(c, fa)
	s.Require().NoError(err)

	go w.Run(ctx)

	const writeCount uint64 = 100
	for i := uint64(0); i < writeCount; i++ {
		w.C() <- rockset.WriteRequest{
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

func (s *WriterSuite) TestCancellation() {
	ctx, cancel := context.WithCancel(testCtx())
	c := rockset.WriterConfig{
		BatchDocumentCount: 30,
		FlushInterval:      time.Millisecond * 10,
	}
	fa := &fakeAdder{}
	w, err := rockset.NewWriter(c, fa)
	s.Require().NoError(err)

	go w.Run(ctx)

	w.C() <- rockset.WriteRequest{
		Workspace:  "workspace",
		Collection: "collection",
		Data:       testObject{},
	}
	time.Sleep(20 * time.Millisecond)
	w.C() <- rockset.WriteRequest{
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
	ctx := testCtx()
	c := rockset.WriterConfig{
		BatchDocumentCount: 30,
		FlushInterval:      time.Millisecond * 500,
	}
	fa := &fakeAdder{}
	w, err := rockset.NewWriter(c, fa)
	s.Require().NoError(err)

	go w.Run(ctx)

	const writeCount uint64 = 100
	for i := uint64(0); i < writeCount; i++ {
		w.C() <- rockset.WriteRequest{
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
	c := rockset.WriterConfig{
		FlushInterval: time.Millisecond * 500,
	}

	benchmarkWriter(b, c)
}

func BenchmarkWriter50(b *testing.B) {
	c := rockset.WriterConfig{
		FlushInterval: time.Millisecond * 500,
	}

	benchmarkWriter(b, c)
}

func BenchmarkWriter100(b *testing.B) {
	c := rockset.WriterConfig{
		FlushInterval: time.Millisecond * 500,
	}

	benchmarkWriter(b, c)
}

func benchmarkWriter(b *testing.B, c rockset.WriterConfig) {
	ctx := testCtx()
	fa := &fakeAdder{}
	w, err := rockset.NewWriter(c, fa)
	require.NoError(b, err)

	go w.Run(ctx)

	doc := rockset.WriteRequest{
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
	skipUnlessIntegrationTest(t)
	ctx := testCtx()

	rc, err := rockset.NewClient()
	require.Nil(t, err)

	c := rockset.WriterConfig{
		FlushInterval: time.Millisecond * 200,
		Workers:       4,
	}
	w, err := rockset.NewWriter(c, rc)
	require.NoError(t, err)

	go w.Run(ctx)

	const writeCount uint64 = 10_000

	name := randomName("write")
	t.Logf("writing %d documents to %s", writeCount, name)
	t0 := time.Now()

	for i := uint64(0); i < writeCount; i++ {
		w.C() <- rockset.WriteRequest{
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
	// sleep 1s to allow the data to be queryable
	time.Sleep(time.Second)

	resp, err := rc.Query(ctx, "SELECT COUNT(*) AS cnt FROM commons.writetest WHERE Bar = :name",
		option.WithParameter("name", "string", name))
	require.NoError(t, err)

	require.Len(t, resp.Results, 1)
	r := resp.Results[0]
	cnt := r["cnt"]
	count := cnt.(float64)
	assert.Equal(t, writeCount, uint64(count))
}
