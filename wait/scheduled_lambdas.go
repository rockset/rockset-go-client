package wait

import "context"

// UntilScheduledLambdaAvailable waits until the scheduledLambda is available.
func (w *Waiter) UntilScheduledLambdaAvailable(ctx context.Context, workspace, scheduledLambda string) error {
	return w.rc.RetryWithCheck(ctx, ResourceIsAvailable(ctx, func(ctx context.Context) error {
		_, err := w.rc.GetScheduledLambda(ctx, workspace, scheduledLambda)
		return err
	}))
}

// UntilScheduledLambdaGone waits until the scheduledLambda is gone.
func (w *Waiter) UntilScheduledLambdaGone(ctx context.Context, workspace, scheduledLambda string) error {
	return w.rc.RetryWithCheck(ctx, ResourceIsGone(ctx, func(ctx context.Context) error {
		_, err := w.rc.GetScheduledLambda(ctx, workspace, scheduledLambda)
		return err
	}))
}
