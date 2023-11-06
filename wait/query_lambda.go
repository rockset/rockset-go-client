package wait

import (
	"context"

	"github.com/rockset/rockset-go-client/option"
)

// UntilQueryLambdaVersionGone waits until a query lambda is deleted, i.e. GetQueryLambda() returns "not found".
func (w *Waiter) UntilQueryLambdaVersionGone(ctx context.Context, workspace, name, version string) error {
	return w.rc.RetryWithCheck(ctx, ResourceIsGone(ctx, func(ctx context.Context) error {
		_, err := w.rc.GetQueryLambdaVersion(ctx, workspace, name, version)
		return err
	}))
}

// UntilQueryLambdaVersionActive waits until the Virtual Instance is active.
func (w *Waiter) UntilQueryLambdaVersionActive(ctx context.Context, workspace, name, version string) error {
	return w.rc.RetryWithCheck(ctx,
		ResourceHasState(ctx,
			[]option.QueryLambdaState{option.QueryLambdaActive}, []option.QueryLambdaState{option.QueryLambdaInvalidSQL},
			func(ctx context.Context) (option.QueryLambdaState, error) {
				ql, err := w.rc.GetQueryLambdaVersion(ctx, workspace, name, version)
				return option.QueryLambdaState(ql.GetState()), err
			}))
}
