package rockset

import (
	"bytes"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"

	"github.com/rockset/rockset-go-client/openapi"
)

func TestLogResult(t *testing.T) {
	buf := &bytes.Buffer{}
	log := zerolog.New(buf).Level(zerolog.TraceLevel).With().Logger()

	logResult(log.Info(), []openapi.DocumentStatus{
		{Status: openapi.PtrString("foo")},
		{Status: openapi.PtrString("foo")},
		{Status: openapi.PtrString("bar")},
	}).Msg("test")
	assert.Equal(t, `{"level":"info","foo":2,"bar":1,"message":"test"}
`, buf.String())
}
