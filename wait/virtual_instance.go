package wait

import (
	"context"

	"github.com/rockset/rockset-go-client/option"
)

func (w *Waiter) UntilVirtualInstanceActive(ctx context.Context, id string) error {
	return w.rc.RetryWithCheck(ctx, ResourceHasState(ctx, []string{option.VirtualInstanceActive},
		func(ctx context.Context) (string, error) {
			vi, err := w.rc.GetVirtualInstance(ctx, id)
			return vi.GetState(), err
		}))
}

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
