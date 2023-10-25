package retry

import (
	"context"
	"math/rand"
	"time"

	"github.com/rs/zerolog"
)

// Exponential is used to perform API call retries with exponential backoff.
type Exponential struct {
	// MaxBackoff is the max time the exponential backoff can wait
	MaxBackoff time.Duration
	// WaitInterval is the initial interval wait between consecutive calls
	WaitInterval time.Duration
	// JitterFraction is the fraction of wait interval to use as jitter [0,1.0]
	JitterFraction float64
	// RetryableErrorCheck is the function that determines if an error should be retried.
	// If nil, it uses the RetryableError().
	RetryableErrorCheck func(error) bool
}

// NewExponential creates a new Retrier which uses exponential backoff.
func NewExponential() Exponential {
	return Exponential{}
}

const DefaultJitterFraction = .05

// Retry retries retryFn until it returns an error. Uses exponential backoff.
// If the retryFn returns an error, it is wrapped in an Error, which implements RetryableError
// so the RetryableErrorCheck can determine if it should retry the operation.
func (r Exponential) Retry(ctx context.Context, retryFn Func) error {
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
	jitterFraction := DefaultJitterFraction
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
			//nolint:gosec
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
func (r Exponential) RetryWithCheck(ctx context.Context, checkFn CheckFn) error {
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
	jitterFraction := DefaultJitterFraction
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
			//nolint:gosec
			var jitter = time.Duration(jitterFraction*rand.Float64()) * waitInterval
			waitInterval *= 2
			waitInterval += jitter
			if waitInterval > maxBackoff {
				waitInterval = maxBackoff
			}
		}
	}
}
