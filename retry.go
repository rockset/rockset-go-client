package rockset

import (
	"context"
	"time"

	"github.com/rs/zerolog"
)

// RetryFunc is the function Retrier will continually call until it either returns true or an error.
type RetryFunc func() (retry bool, err error)

// Retrier is the interface used by the RockClient convenience methods to retry a noperation
// which returned a rockset.Error which is Retryable()
type Retrier interface {
	Retry(ctx context.Context, retryFn RetryFunc) error
}

// ExponentialRetry is used to perform retries with exponential backoff
type ExponentialRetry struct {
	// MaxBackoff is the max time the exponential backoff can wait
	MaxBackoff time.Duration
	// WaitInterval is the initial interval wait between consecutive calls
	WaitInterval time.Duration
}

// Retry retries retryFn until it returns false, or an error. Uses exponential backoff.
func (r ExponentialRetry) Retry(ctx context.Context, retryFn RetryFunc) error {
	t0 := time.Now()
	log := zerolog.Ctx(ctx)

	var maxBackoff = 8 * time.Second
	if r.MaxBackoff != 0 {
		maxBackoff = r.MaxBackoff
	}
	waitInterval := time.Second
	if r.WaitInterval != 0 {
		waitInterval = r.WaitInterval
	}

	defer func() {
		log.Debug().Str("d", time.Since(t0).String()).Msg("retry duration")
	}()

	for {
		t1 := time.Now()
		retry, err := retryFn()
		log.Debug().Str("d", time.Since(t1).String()).Msg("call curation")
		if err != nil {
			return err
		}
		if !retry {
			return nil
		}

		t := time.NewTimer(waitInterval)
		select {
		case <-ctx.Done():
			log.Debug().Msg("wait cancelled")
			t.Stop()

			return ctx.Err()
		case t := <-t.C:
			log.Trace().Str("t", t.String()).Msg("wait time")
			if waitInterval < maxBackoff {
				waitInterval *= 2
			}
		}
	}
}

