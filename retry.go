package rockset

import (
	"context"
	"errors"
	"time"

	"github.com/rs/zerolog"
)

// RetryFunc is the function Retrier will call as long as it returns an error which is retryable.
type RetryFunc func() (err error)

// RetryCheck is the function Retrier will call until the RetryFunc returns false or and error.
type RetryCheck func() (retry bool, err error)

// Retrier is the interface used by the RockClient convenience methods to retry an operation
// which returned a rockset.Error which is Retryable().
type Retrier interface {
	// Retry will retry retryFn if it returns an error which is retryable
	Retry(ctx context.Context, retryFn RetryFunc) error
	// RetryWithCheck will retry checkFn until it returns false or an error
	RetryWithCheck(ctx context.Context, checkFunc RetryCheck) error
}

// RetryableError returns true if err is a Rockset error that is retryable
func RetryableError(err error) bool {
	var re Error
	if errors.As(err, &re) {
		if re.Retryable() {
			return true
		}
	}

	return false
}

// ExponentialRetry is used to perform API cal retries with exponential backoff
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
		log.Debug().Str("d", time.Since(t0).String()).Msg("total duration")
	}()

	for {
		t1 := time.Now()
		err := retryFn()
		log.Trace().Str("d", time.Since(t1).String()).Msg("api call curation")

		// no error, so no need to retry
		if err == nil {
			return nil
		}

		// wrap the error in a rockset.Error so we can check if it is retryable
		re := NewError(err)
		if !re.Retryable() {
			return re
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

// RetryWithCheck will retry checkFn until it returns false or an error
func (r ExponentialRetry) RetryWithCheck(ctx context.Context, checkFn RetryCheck) error {
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
		retry, err := checkFn()
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
