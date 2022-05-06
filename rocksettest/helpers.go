package rocksettest

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
)

// helper function to skip if envVar is not set
// returns the value of os.Getenv(envVar)
func SkipUnlessEnv(t *testing.T, envVar string) string {
	value := os.Getenv(envVar)
	if value == "" {
		t.Skipf("skipping as %s is not set", envVar)
	}

	return value
}

// helper function to create a context with a zerolog logger
func TestCtx() context.Context {
	ctx := context.Background()
	console := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	log := zerolog.New(console).Level(zerolog.TraceLevel).With().Timestamp().Logger()

	return log.WithContext(ctx)
}
