package test

import (
	"context"

	"github.com/rockset/rockset-go-client/retry"
)

// Retrier is used for testing
type Retrier struct{}

func (t Retrier) Retry(ctx context.Context, retryFn retry.Func) error {
	for {
		err := retryFn()
		if err == nil {
			return nil
		}
		if !retry.DefaultRetryableErrorCheck(err) {
			return err
		}
	}
}

func (t Retrier) RetryWithCheck(ctx context.Context, checkFunc retry.CheckFn) error {
	for {
		r, err := checkFunc()
		if err == nil {
			return nil
		}
		if !r {
			return nil
		}
		if !retry.DefaultRetryableErrorCheck(err) {
			return err
		}
	}
}
