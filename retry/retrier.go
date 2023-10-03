package retry

import (
	"context"
	"errors"
)

// Func is the function Retrier will call as long as it returns an error which is retryable.
type Func func() (err error)

// CheckFn is the function Retrier will call until the Func returns false or an error.
type CheckFn func() (retry bool, err error)

// Retrier is the interface used by the RockClient convenience methods to retry an operation
// which returned an errors.Error which is Retryable().
type Retrier interface {
	// Retry will retry retryFn if it returns an error which is retryable
	Retry(ctx context.Context, retryFn Func) error
	// RetryWithCheck will retry checkFn until it returns false or an error
	RetryWithCheck(ctx context.Context, checkFunc CheckFn) error
}

// RetryableError is an error which can be retried if Retryable() returns true.
type RetryableError interface {
	error
	Retryable() bool
}

// DefaultRetryableErrorCheck returns true if err is an error that is retryable, i.e. implements RetryableError.
// This function is used to determine which errors to retry for the convenience methods on the rockset.RockClient.
func DefaultRetryableErrorCheck(err error) bool {
	var re RetryableError
	if errors.As(err, &re) {
		return re.Retryable()
	}

	return false
}
