package wait

import (
	"context"

	"github.com/rockset/rockset-go-client/option"
)

// UntilQueryDone waits until queryID has either completed, errored, or been cancelled.
func (w *Waiter) UntilQueryDone(ctx context.Context, queryID string) error {
	// TODO should this only wait for COMPLETED and return an error for ERROR and CANCELLED?
	return w.rc.RetryWithCheck(ctx, ResourceHasState(ctx, []option.QueryState{option.QueryCompleted, option.QueryError, option.QueryCancelled},
		func(ctx context.Context) (option.QueryState, error) {
			q, err := w.rc.GetQueryInfo(ctx, queryID)
			return option.QueryState(q.GetStatus()), err
		}))
}
