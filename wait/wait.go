package wait

import (
	"context"
	"errors"

	"github.com/rs/zerolog"

	rockerr "github.com/rockset/rockset-go-client/errors"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/retry"
)

type Waiter struct {
	rc Rocksetter
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o fake . Rocksetter
type Rocksetter interface {
	retry.Retrier
	GetAlias(ctx context.Context, workspace string, name string) (openapi.Alias, error)
	GetCollection(ctx context.Context, workspace string, name string) (openapi.Collection, error)
}

func New(rs Rocksetter) *Waiter {
	return &Waiter{rs}
}

// resourceHasState implements RetryFn to wait until the resource has the desired state
func resourceHasState[T comparable](ctx context.Context, states []T,
	fn func(ctx context.Context) (T, error)) retry.CheckFn {
	return func() (bool, error) {
		zl := zerolog.Ctx(ctx)
		state, err := fn(ctx)
		if err != nil {
			return false, err
		}

		for _, s := range states {
			if state == s {
				return false, nil
			}
		}

		zl.Trace().Any("current", state).Msg("waiting for resource state")

		return true, nil
	}
}

// resourceIsAvailable implements RetryFn to wait until the resource is present
func resourceIsAvailable(ctx context.Context, fn func(ctx context.Context) error) retry.CheckFn {
	return func() (bool, error) {
		err := fn(ctx)

		if err == nil {
			// the resource is present, done
			return false, nil
		}

		var re rockerr.Error
		if errors.As(err, &re) {
			if re.IsNotFoundError() {
				// the resource is not present, retry
				return true, nil
			}
		}

		// we got an error, stop
		return false, err
	}
}

// resourceIsGone implements RetryFn to wait until the resource is gone
func resourceIsGone(ctx context.Context, fn func(ctx context.Context) error) retry.CheckFn {
	return func() (bool, error) {
		err := fn(ctx)

		if err == nil {
			// the resource is still present, retry the operation
			return true, nil
		}

		var re rockerr.Error
		if errors.As(err, &re) {
			if re.IsNotFoundError() {
				// the resource is no longer present
				return false, nil
			}
		}

		return false, err
	}
}
