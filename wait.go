package rockset

import (
	"context"
	"errors"

	"github.com/rs/zerolog"
)

const (
	collectionStatusREADY = "READY"
)

// WaitUntilAliasAvailable waits until the alias is available.
func (rc *RockClient) WaitUntilAliasAvailable(ctx context.Context, workspace, alias string) error {
	return rc.RetryWithCheck(ctx, func() (bool, error) {
		_, _, err := rc.AliasesApi.GetAlias(ctx, workspace, alias).Execute()

		if err == nil {
			return false, nil
		}

		re := NewError(err)
		if re.IsNotFoundError() {
			return true, nil
		}
		if re.Retryable() {
			return true, nil
		}

		return false, err
	})
}

// WaitUntilCollectionReady waits until the collection is ready.
func (rc *RockClient) WaitUntilCollectionReady(ctx context.Context, workspace, name string) error {
	return rc.RetryWithCheck(ctx, rc.collectionHasState(ctx, workspace, name, collectionStatusREADY))
}

// WaitUntilCollectionGone waits until a collection marked for deletion is gone, i.e. GetCollection()
// returns "not found".
func (rc *RockClient) WaitUntilCollectionGone(ctx context.Context, workspace, name string) error {
	return rc.RetryWithCheck(ctx, rc.collectionIsGone(ctx, workspace, name))
}

// WaitUntilViewGone waits until a view marked for deletion is gone, i.e. GetView()
// returns "not found".
func (rc *RockClient) WaitUntilViewGone(ctx context.Context, workspace, name string) error {
	return rc.RetryWithCheck(ctx, rc.viewIsGone(ctx, workspace, name))
}

// WaitUntilCollectionDocuments waits until the collection has at least count new documents
func (rc *RockClient) WaitUntilCollectionDocuments(ctx context.Context, workspace, name string, count int64) error {
	waiter := docWaiter{rc: rc}
	return rc.RetryWithCheck(ctx, waiter.collectionHasNewDocs(ctx, workspace, name, count))
}

// WaitUntilWorkspaceAvailable waits until the workspace is available.
func (rc *RockClient) WaitUntilWorkspaceAvailable(ctx context.Context, workspace string) error {
	return rc.RetryWithCheck(ctx, rc.workspaceIsAvailable(ctx, workspace))
}

// TODO(pme) refactor viewIsGone() and collectionIsGone() to be DRY

func (rc *RockClient) workspaceIsAvailable(ctx context.Context, workspace string) RetryCheck {
	return func() (bool, error) {
		_, err := rc.GetWorkspace(ctx, workspace)

		if err == nil {
			// the collection still exist, and we are done
			return false, nil
		}

		var re Error
		if errors.As(err, &re) {
			if re.IsNotFoundError() {
				// the view is no longer present
				return true, nil
			}
			if re.Retryable() {
				return true, nil
			}
		}

		return false, err
	}
}

// viewIsGone implements RetryFn to wait until the view is deleted
func (rc *RockClient) viewIsGone(ctx context.Context, workspace, name string) RetryCheck {
	return func() (bool, error) {
		_, err := rc.GetView(ctx, workspace, name)

		if err == nil {
			// the collection still exist, retry
			return true, nil
		}

		var re Error
		if errors.As(err, &re) {
			if re.IsNotFoundError() {
				// the view is no longer present
				return false, nil
			}
			if re.Retryable() {
				return true, nil
			}
		}

		return false, err
	}
}

func (rc *RockClient) collectionIsGone(ctx context.Context, workspace, name string) RetryCheck {
	return func() (bool, error) {
		_, err := rc.GetCollection(ctx, workspace, name)

		if err == nil {
			// the collection still exist, retry
			return true, nil
		}

		var re Error
		if errors.As(err, &re) {
			if re.IsNotFoundError() {
				// the collection is gone
				return false, nil
			}
			if re.Retryable() {
				return true, nil
			}
		}

		return false, err
	}
}

func (rc *RockClient) collectionHasState(ctx context.Context, workspace, name, state string) RetryCheck {
	return func() (bool, error) {
		log := zerolog.Ctx(ctx)

		c, err := rc.GetCollection(ctx, workspace, name)
		if err != nil {
			re := NewError(err)
			if re.Retryable() {
				return true, nil
			}

			return false, err
		}

		log.Debug().Str("status", c.GetStatus()).Str("workspace", workspace).
			Str("collection", name).Msg("collectionHasState()")
		if c.GetStatus() == state {
			return false, nil
		}

		return true, nil
	}
}

type docWaiter struct {
	rc        *RockClient
	prevCount int64
}

func (d *docWaiter) collectionHasNewDocs(ctx context.Context, workspace, name string, count int64) RetryCheck {
	d.prevCount = -1
	return func() (bool, error) {
		log := zerolog.Ctx(ctx)

		c, err := d.rc.GetCollection(ctx, workspace, name)
		if err != nil {
			re := NewError(err)
			if re.Retryable() {
				return true, nil
			}

			return false, err
		}

		current := c.Stats.GetDocCount()
		log.Debug().Str("workspace", workspace).Int64("current", current).
			Int64("previous", d.prevCount).Str("collection", name).
			Int64("count", count).Msg("collectionHasNewDocs()")

		if d.prevCount == -1 {
			d.prevCount = current
		}

		if current-d.prevCount >= count {
			return false, nil
		}

		return true, nil
	}
}
