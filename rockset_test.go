package rockset_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"

	"github.com/rockset/rockset-go-client"
)

// helper function to skip unless ROCKSET_APIKEY is set
func skipUnlessIntegrationTest(t *testing.T) {
	if os.Getenv(rockset.APIKeyEnvironmentVariableName) == "" {
		t.Skipf("skipping as %s is not set", rockset.APIKeyEnvironmentVariableName)
	}
}

// helper function to create a context with a zerolog logger
func testCtx() context.Context {
	ctx := context.Background()
	console := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	log := zerolog.New(console).Level(zerolog.TraceLevel).With().Timestamp().Logger()

	return log.WithContext(ctx)
}
