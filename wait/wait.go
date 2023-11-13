package wait

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/rs/zerolog"

	rockerr "github.com/rockset/rockset-go-client/errors"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/retry"
)

type Waiter struct {
	rc                        ResourceGetter
	QueryLambdaTagPropagation time.Duration
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o fake . ResourceGetter
type ResourceGetter interface {
	retry.Retrier
	GetAlias(ctx context.Context, workspace string, name string) (openapi.Alias, error)
	GetCollection(ctx context.Context, workspace string, name string) (openapi.Collection, error)
	GetCollectionMount(ctx context.Context, id, collectionPath string) (openapi.CollectionMount, error)
	GetIntegration(ctx context.Context, name string) (openapi.Integration, error)
	GetQueryInfo(ctx context.Context, queryID string) (openapi.QueryInfo, error)
	GetQueryLambdaVersion(ctx context.Context, workspace, name, version string) (openapi.QueryLambdaVersion, error)
	GetView(ctx context.Context, workspace, name string) (openapi.View, error)
	GetVirtualInstance(ctx context.Context, id string) (openapi.VirtualInstance, error)
	GetWorkspace(ctx context.Context, name string) (openapi.Workspace, error)
}

const queryLambdaTagPropagation = 2 * time.Minute

func New(rs ResourceGetter) *Waiter {
	return &Waiter{rs, queryLambdaTagPropagation}
}

// ResourceHasState implements RetryFn to wait until the resource has the desired state, and if a bad state is
// encountered it will return ErrBadWaitState
func ResourceHasState[T fmt.Stringer](ctx context.Context, validStates, badStates []T,
	// TODO should T be Stringer instead? Then all
	fn func(ctx context.Context) (T, error)) retry.CheckFn {
	return func() (bool, error) {
		zl := zerolog.Ctx(ctx)
		state, err := fn(ctx)
		if err != nil {
			return false, err
		}

		for _, s := range validStates {
			if state.String() == s.String() {
				return false, nil
			}
		}
		for _, s := range badStates {
			if state.String() == s.String() {
				return false, rockerr.NewBadWaitStateError(state.String()) // fmt.Errorf("%w: %v", ErrBadWaitState, state)
			}
		}

		zl.Trace().Any("current", state).Msg("waiting for resource state")

		return true, nil
	}
}

// ResourceIsAvailable implements RetryFn to wait until the resource is present
func ResourceIsAvailable(ctx context.Context, fn func(ctx context.Context) error) retry.CheckFn {
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

// ResourceIsGone implements RetryFn to wait until the resource is gone
func ResourceIsGone(ctx context.Context, fn func(ctx context.Context) error) retry.CheckFn {
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
