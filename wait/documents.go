package wait

import "context"

// UntilQueryable waits until the collection includes the documents at or after the specified offset(s).
func (w *Waiter) UntilQueryable(ctx context.Context, workspace, name string, offsets []string) error {
	return w.rc.RetryWithCheck(ctx, func() (retry bool, err error) {
		commit, err := w.rc.GetCollectionCommit(ctx, workspace, name, offsets)

		return !commit.GetPassed(), err
	})
}
