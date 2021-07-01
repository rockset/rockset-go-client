package rockset_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/openapi"
)

type fakeAdder struct {
	count int
}

func (f *fakeAdder) AddDocuments(ctx context.Context, workspace, collection string,
	docs []interface{}) ([]openapi.DocumentStatus, error) {
	f.count++
	time.Sleep(100 * time.Millisecond)
	return []openapi.DocumentStatus{}, nil
}

func TestWriter(t *testing.T) {
	ctx := testCtx()
	c := rockset.WriterConfig{
		BufferSize:    30,
		BatchSize:     7,
		FlushInterval: time.Millisecond * 50,
	}
	fa := &fakeAdder{}
	w := rockset.NewWriter(c, fa)

	go w.Run(ctx)
	go w.Worker(ctx)

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

	assert.Equal(t, uint64(0), w.Stats().ErrorCount)
	assert.Equal(t, writeCount, w.Stats().DocumentCount)
}

func TestWriterTimeout(t *testing.T) {
	ctx, cancel := context.WithCancel(testCtx())
	c := rockset.WriterConfig{
		BufferSize:    30,
		BatchSize:     7,
		FlushInterval: time.Millisecond * 10,
	}
	fa := &fakeAdder{}
	w := rockset.NewWriter(c, fa)

	go w.Run(ctx)
	go w.Worker(ctx)

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

	assert.Equal(t, uint64(0), w.Stats().ErrorCount)
	assert.Equal(t, uint64(1), w.Stats().DocumentCount)
}

func BenchmarkWriter10(b *testing.B) {
	c := rockset.WriterConfig{
		BufferSize:    1000,
		BatchSize:     10,
		FlushInterval: time.Millisecond * 500,
	}

	benchmarkWriter(b, c)
}

func BenchmarkWriter50(b *testing.B) {
	c := rockset.WriterConfig{
		BufferSize:    1000,
		BatchSize:     50,
		FlushInterval: time.Millisecond * 500,
	}

	benchmarkWriter(b, c)
}

func BenchmarkWriter100(b *testing.B) {
	c := rockset.WriterConfig{
		BufferSize:    1000,
		BatchSize:     100,
		FlushInterval: time.Millisecond * 500,
	}

	benchmarkWriter(b, c)
}

func benchmarkWriter(b *testing.B, c rockset.WriterConfig) {
	ctx := testCtx()
	fa := &fakeAdder{}
	w := rockset.NewWriter(c, fa)

	go w.Run(ctx)
	go w.Worker(ctx)

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

type testObject struct {
	Foo uint64
	Bar string
}

func TestWriterIntegration(t *testing.T) {
	skipUnlessIntegrationTest(t)
	ctx := testCtx()

	rc, err := rockset.NewClient()
	require.Nil(t, err)

	c := rockset.WriterConfig{
		BufferSize:    1000,
		BatchSize:     250,
		FlushInterval: time.Millisecond * 200,
	}
	w := rockset.NewWriter(c, rc)

	go w.Run(ctx)
	go w.Worker(ctx)
	go w.Worker(ctx)

	const writeCount uint64 = 10_000

	for i := uint64(0); i < writeCount; i++ {
		w.C() <- rockset.WriteRequest{
			Workspace:  "commons",
			Collection: "writetest",
			Data:       testObject{i, "test"},
		}
	}

	w.Stop()
	w.Wait()

	assert.Equal(t, uint64(0), w.Stats().ErrorCount)
	assert.Equal(t, writeCount, w.Stats().DocumentCount)
}
