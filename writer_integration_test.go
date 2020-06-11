// +build integration
// go test -tags integration -v -run TestWriterIntegration -cpuprofile=cpu.out
// go tool pprof -http=:8080 cpu.out

package rockset_test

import (
	"github.com/rockset/rockset-go-client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type testObject struct {
	Uint64 uint64
}

func TestWriterIntegration(t *testing.T) {
	rc, err := rockset.NewClient(rockset.FromEnv())
	require.Nil(t, err)

	c := rockset.WriterConfig{
		BufferSize:    100,
		BatchSize:     25,
		FlushInterval: time.Millisecond * 200,
	}
	w := rockset.NewWriter(c)

	go w.Run(rc.Documents)
	go w.Worker(rc.Documents)
	go w.Worker(rc.Documents)

	const writeCount uint64 = 10_000

	for i := uint64(0); i < writeCount; i++ {
		w.C() <- rockset.WriteRequest{
			Workspace:  "commons",
			Collection: "writetest",
			Data:       testObject{i},
		}
	}

	w.Stop(true)
	w.Wait()

	assert.Equal(t, uint64(0), w.Stats().ErrorCount)
	assert.Equal(t, writeCount, w.Stats().DocumentCount)
}
