package wait

import (
	"context"
	"time"

	"github.com/rockset/rockset-go-client/option"
)

// UntilQueryLambdaVersionGone waits until a query lambda is deleted, i.e. GetQueryLambda() returns "not found".
func (w *Waiter) UntilQueryLambdaVersionGone(ctx context.Context, workspace, name, version string) error {
	return w.rc.RetryWithCheck(ctx, ResourceIsGone(ctx, func(ctx context.Context) error {
		_, err := w.rc.GetQueryLambdaVersion(ctx, workspace, name, version)
		return err
	}))
}

// UntilQueryLambdaVersionActive waits until the query lambda version is active. Returns an error if the query lambda
// contains invalid SQL.
func (w *Waiter) UntilQueryLambdaVersionActive(ctx context.Context, workspace, name, version string) error {
	return w.rc.RetryWithCheck(ctx,
		ResourceHasState(ctx,
			[]option.QueryLambdaState{option.QueryLambdaActive}, []option.QueryLambdaState{option.QueryLambdaInvalidSQL},
			func(ctx context.Context) (option.QueryLambdaState, error) {
				ql, err := w.rc.GetQueryLambdaVersion(ctx, workspace, name, version)
				return option.QueryLambdaState(ql.GetState()), err
			}))
}

// UntilQueryLambdaTagPropagated waits until the query lambda tag has been propagated throughout the system.
func (w *Waiter) UntilQueryLambdaTagPropagated(ctx context.Context, workspace, name, tag string) error {
	t := time.NewTimer(w.QueryLambdaTagPropagation) // eww
	defer t.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.C:
		return nil
	}
}
