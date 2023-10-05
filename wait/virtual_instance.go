package wait

import (
	"context"

	"github.com/rockset/rockset-go-client/option"
)

// UntilVirtualInstanceActive waits until the Virtual Instance is active.
func (w *Waiter) UntilVirtualInstanceActive(ctx context.Context, id string) error {
	return w.rc.RetryWithCheck(ctx, ResourceHasState(ctx, []string{option.VirtualInstanceActive},
		func(ctx context.Context) (string, error) {
			vi, err := w.rc.GetVirtualInstance(ctx, id)
			return vi.GetState(), err
		}))
}

// UntilVirtualInstanceGone waits until the Virtual Instance is gone.
func (w *Waiter) UntilVirtualInstanceGone(ctx context.Context, id string) error {
	return w.rc.RetryWithCheck(ctx, ResourceIsGone(ctx, func(ctx context.Context) error {
		_, err := w.rc.GetVirtualInstance(ctx, id)
		return err
	}))
}

// UntilVirtualInstanceSuspended waits until the Virtual Instance is suspended.
func (w *Waiter) UntilVirtualInstanceSuspended(ctx context.Context, id string) error {
	return w.rc.RetryWithCheck(ctx, ResourceHasState(ctx, []string{option.VirtualInstanceSuspended},
		func(ctx context.Context) (string, error) {
			vi, err := w.rc.GetVirtualInstance(ctx, id)
			return vi.GetState(), err
		}))
}

// UntilMountActive waits until the collection mount is active, and queries can be issued to it on the
// virtual instance.
func (w *Waiter) UntilMountActive(ctx context.Context, vID, workspace, collection string) error {
	return w.rc.RetryWithCheck(ctx, ResourceHasState(ctx, []string{option.MountActive},
		func(ctx context.Context) (string, error) {
			cm, err := w.rc.GetCollectionMount(ctx, vID, workspace+"."+collection)
			return cm.GetState(), err
		}))
}

// UntilMountGone waits until the collection mount is gone.
func (w *Waiter) UntilMountGone(ctx context.Context, vID, workspace, collection string) error {
	return w.rc.RetryWithCheck(ctx, ResourceIsGone(ctx, func(ctx context.Context) error {
		_, err := w.rc.GetCollectionMount(ctx, vID, workspace+"."+collection)
		return err
	}))
}
