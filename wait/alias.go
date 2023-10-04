package wait

import "context"

// UntilAliasAvailable waits until the alias is available.
func (w *Waiter) UntilAliasAvailable(ctx context.Context, workspace, alias string) error {
	return w.rc.RetryWithCheck(ctx, ResourceIsAvailable(ctx, func(ctx context.Context) error {
		_, err := w.rc.GetAlias(ctx, workspace, alias)
		return err
	}))
}

// UntilAliasGone waits until the alias is gone.
func (w *Waiter) UntilAliasGone(ctx context.Context, workspace, alias string) error {
	return w.rc.RetryWithCheck(ctx, ResourceIsGone(ctx, func(ctx context.Context) error {
		_, err := w.rc.GetAlias(ctx, workspace, alias)
		return err
	}))
}
