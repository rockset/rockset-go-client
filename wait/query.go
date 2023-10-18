package wait

import (
	"context"

	"github.com/rockset/rockset-go-client/option"
)

// UntilQueryDone waits until queryID has either completed.
// Returns ErrBadWaitState if the query failed or was cancelled.
func (w *Waiter) UntilQueryDone(ctx context.Context, queryID string) error {
	return w.rc.RetryWithCheck(ctx,
		ResourceHasState(ctx, []option.QueryState{option.QueryCompleted},
			[]option.QueryState{option.QueryError, option.QueryCancelled},
			func(ctx context.Context) (option.QueryState, error) {
				q, err := w.rc.GetQueryInfo(ctx, queryID)
				return option.QueryState(q.GetStatus()), err
			}))
}
