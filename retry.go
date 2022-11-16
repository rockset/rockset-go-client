package rockset

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/rs/zerolog"
)

// RetryFunc is the function Retrier will call as long as it returns an error which is retryable.
type RetryFunc func() (err error)

// RetryCheck is the function Retrier will call until the RetryFunc returns false or an error.
type RetryCheck func() (retry bool, err error)

// Retrier is the interface used by the RockClient convenience methods to retry an operation
// which returned a rockset.Error which is Retryable().
type Retrier interface {
	// Retry will retry retryFn if it returns an error which is retryable
	Retry(ctx context.Context, retryFn RetryFunc) error
	// RetryWithCheck will retry checkFn until it returns false or an error
	RetryWithCheck(ctx context.Context, checkFunc RetryCheck) error
}

// RetryableError is an error which can be retried
type RetryableError interface {
	error
	Retryable() bool
}

// DefaultRetryableErrorCheck returns true if err is an error that is retryable, i.e. implements RetryableError.
// This function is used to determine which errors to retry for the convenience methods on the RockClient.
func DefaultRetryableErrorCheck(err error) bool {
	var re RetryableError
	if errors.As(err, &re) {
		return re.Retryable()
	}

	return false
}

// ExponentialRetry is used to perform API cal retries with exponential backoff
type ExponentialRetry struct {
	// MaxBackoff is the max time the exponential backoff can wait
	MaxBackoff time.Duration
	// WaitInterval is the initial interval wait between consecutive calls
	WaitInterval time.Duration
	// JitterFraction is the fraction of wait interval to use as jitter [0,1.0]
	JitterFraction float64
	// RetryableErrorCheck is the function that determines if an error should be retried. If nil, it uses the RetryableError().
	RetryableErrorCheck func(error) bool
}

// Retry retries retryFn until it returns false, or an error. Uses exponential backoff.
// If the retryFn returns an error, it is wrapped in an Error, which implements RetryableError
// so the RetryableErrorCheck can determine if it should retry the operation.
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
	jitterFraction := .05
	if r.JitterFraction != 0 {
		jitterFraction = r.JitterFraction
	}
	checkFn := DefaultRetryableErrorCheck
	if r.RetryableErrorCheck != nil {
		checkFn = r.RetryableErrorCheck
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

		// if it already isn't a RetryableError, wrap the error so the caller can determine if it is a retryable error
		var re RetryableError
		if !errors.As(err, &re) {
			err = NewError(err)
		}

		if !checkFn(err) {
			return err
		}
		log.Debug().Err(err).Msg("retrying due to retryable error")

		t := time.NewTimer(waitInterval)
		select {
		case <-ctx.Done():
			log.Debug().Msg("wait cancelled")
			t.Stop()

			return ctx.Err()
		case t := <-t.C:
			log.Trace().Str("t", t.String()).Msg("wait time")
			var jitter = time.Duration(jitterFraction*rand.Float64()) * waitInterval
			waitInterval *= 2
			waitInterval += jitter
			if waitInterval > maxBackoff {
				waitInterval = maxBackoff
			}
		}
	}
}

// RetryWithCheck will retry checkFn until it returns false or an error. If checkFn returns false, RetryWithCheck will
// return nil, otherwise it'll return the error.
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
	jitterFraction := .05
	if r.JitterFraction != 0 {
		jitterFraction = r.JitterFraction
	}

	defer func() {
		log.Debug().Str("d", time.Since(t0).String()).Msg("retry duration")
	}()

	for {
		t1 := time.Now()
		retry, err := checkFn()
		log.Debug().Str("d", time.Since(t1).String()).Msg("call duration")

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
			var jitter = time.Duration(jitterFraction*rand.Float64()) * waitInterval
			waitInterval *= 2
			waitInterval += jitter
			if waitInterval > maxBackoff {
				waitInterval = maxBackoff
			}
		}
	}
}
