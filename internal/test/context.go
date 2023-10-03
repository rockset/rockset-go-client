package test

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog"
)

// Context is helper function to create a context with a zerolog logger
func Context() context.Context {
	return ContextWithLevel(zerolog.WarnLevel)
}

func ContextWithLevel(lvl zerolog.Level) context.Context {
	ctx := context.Background()
	console := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	log := zerolog.New(console).Level(lvl).With().Timestamp().Logger()

	return log.WithContext(ctx)
}
