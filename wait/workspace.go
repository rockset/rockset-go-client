package wait

import "context"

// UntilWorkspaceAvailable waits until the workspace is available.
func (w *Waiter) UntilWorkspaceAvailable(ctx context.Context, workspace string) error {
	return w.rc.RetryWithCheck(ctx, ResourceIsAvailable(ctx, func(ctx context.Context) error {
		_, err := w.rc.GetWorkspace(ctx, workspace)
		return err
	}))
}

// UntilWorkspaceGone waits until the workspace is gone, i.e. GetWorkspace() returns "not found".
func (w *Waiter) UntilWorkspaceGone(ctx context.Context, workspace string) error {
	return w.rc.RetryWithCheck(ctx, ResourceIsGone(ctx, func(ctx context.Context) error {
		_, err := w.rc.GetWorkspace(ctx, workspace)
		return err
	}))
}
