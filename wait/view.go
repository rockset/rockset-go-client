package wait

import "context"

// UntilViewGone waits until a view marked for deletion is gone, i.e. GetView() returns "not found".
func (w *Waiter) UntilViewGone(ctx context.Context, workspace, name string) error {
	return w.rc.RetryWithCheck(ctx, ResourceIsGone(ctx, func(ctx context.Context) error {
		_, err := w.rc.GetView(ctx, workspace, name)
		return err
	}))
}
