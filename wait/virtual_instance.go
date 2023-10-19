package wait

import (
	"context"
	"errors"
	"regexp"

	rockerr "github.com/rockset/rockset-go-client/errors"
	"github.com/rockset/rockset-go-client/option"
	"github.com/rockset/rockset-go-client/retry"
)

// UntilVirtualInstanceActive waits until the Virtual Instance is active.
func (w *Waiter) UntilVirtualInstanceActive(ctx context.Context, id string) error {
	return w.rc.RetryWithCheck(ctx, ResourceHasState(ctx, []option.VirtualInstanceState{option.VirtualInstanceActive}, nil,
		func(ctx context.Context) (option.VirtualInstanceState, error) {
			vi, err := w.rc.GetVirtualInstance(ctx, id)
			return option.VirtualInstanceState(vi.GetState()), err
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
	return w.rc.RetryWithCheck(ctx, ResourceHasState(ctx,
		[]option.VirtualInstanceState{option.VirtualInstanceSuspended}, nil,
		func(ctx context.Context) (option.VirtualInstanceState, error) {
			vi, err := w.rc.GetVirtualInstance(ctx, id)
			return option.VirtualInstanceState(vi.GetState()), err
		}))
}

// UntilMountActive waits until the collection mount is active, and queries can be issued to it on the
// virtual instance.
func (w *Waiter) UntilMountActive(ctx context.Context, vID, workspace, collection string) error {
	return w.rc.RetryWithCheck(ctx, ResourceHasState(ctx, []option.MountState{option.MountActive}, nil,
		func(ctx context.Context) (option.MountState, error) {
			cm, err := w.rc.GetCollectionMount(ctx, vID, workspace+"."+collection)
			return option.MountState(cm.GetState()), err
		}))
}

// UntilMountGone waits until the collection mount is gone.
func (w *Waiter) UntilMountGone(ctx context.Context, vID, workspace, collection string) error {
	// TODO this can't use ResourceIsGone as the API returns 400 instead of the expected 404 when the
	//  mount doesn't exist
	return w.rc.RetryWithCheck(ctx, mountIsGone(ctx, func(ctx context.Context) error {
		_, err := w.rc.GetCollectionMount(ctx, vID, workspace+"."+collection)
		return err
	}))
}

var mountRe = regexp.MustCompile("Collection .* is not mounted on virtual instance .*")

// TODO(pme) once ORC-3594 is fixed, this test can be removed
func mountIsGone(ctx context.Context, fn func(ctx context.Context) error) retry.CheckFn {
	return func() (bool, error) {
		err := fn(ctx)

		if err == nil {
			// the resource is still present, retry the operation
			return true, nil
		}

		var re rockerr.Error
		if errors.As(err, &re) {
			if re.IsNotFoundError() {
				// the resource is no longer present
				return false, nil
			}
			if mountRe.MatchString(re.GetMessage()) {
				return false, nil
			}
		}

		return false, err
	}
}
