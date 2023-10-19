package wait

import (
	"context"

	"github.com/rs/zerolog"

	rockerr "github.com/rockset/rockset-go-client/errors"
	"github.com/rockset/rockset-go-client/option"
)

// UntilCollectionReady waits until the collection is ready.
func (w *Waiter) UntilCollectionReady(ctx context.Context, workspace, name string) error {
	return w.rc.RetryWithCheck(ctx, ResourceHasState(ctx, []option.CollectionStatus{option.CollectionStatusReady}, nil,
		func(ctx context.Context) (option.CollectionStatus, error) {
			c, err := w.rc.GetCollection(ctx, workspace, name)
			return option.CollectionStatus(c.GetStatus()), err
		}))
}

// UntilCollectionGone waits until a collection marked for deletion is gone, i.e. GetCollection()
// returns "not found".
func (w *Waiter) UntilCollectionGone(ctx context.Context, workspace, name string) error {
	return w.rc.RetryWithCheck(ctx, ResourceIsGone(ctx, func(ctx context.Context) error {
		_, err := w.rc.GetCollection(ctx, workspace, name)
		return err
	}))
}

// UntilCollectionHasNewDocuments waits until the collection has at least count new documents
// (measured from when the method is called).
func (w *Waiter) UntilCollectionHasNewDocuments(ctx context.Context, workspace, name string, count int64) error {
	var prevCount int64 = -1
	return w.rc.RetryWithCheck(ctx, func() (retry bool, err error) {
		zl := zerolog.Ctx(ctx)
		c, err := w.rc.GetCollection(ctx, workspace, name)
		if err != nil {
			re := rockerr.New(err)
			if re.Retryable() {
				return true, nil
			}

			return false, err
		}

		current := c.Stats.GetDocCount()
		zl.Debug().Str("workspace", workspace).Int64("current", current).
			Int64("previous", prevCount).Str("collection", name).
			Int64("count", count).Msg("collectionHasNewDocs()")

		if prevCount == -1 {
			prevCount = current
		}

		if current-prevCount >= count {
			return false, nil
		}

		return true, nil
	})
}

// UntilCollectionHasDocuments waits until the collection has at least count documents
func (w *Waiter) UntilCollectionHasDocuments(ctx context.Context, workspace, name string, count int64) error {
	return w.rc.RetryWithCheck(ctx, func() (retry bool, err error) {
		zl := zerolog.Ctx(ctx)
		c, err := w.rc.GetCollection(ctx, workspace, name)
		if err != nil {
			re := rockerr.New(err)
			if re.Retryable() {
				return true, nil
			}

			return false, err
		}

		current := c.Stats.GetDocCount()
		zl.Debug().Str("workspace", workspace).Int64("current", current).
			Str("collection", name).
			Int64("count", count).Msg("collectionHasNewDocs()")

		if current >= count {
			return false, nil
		}

		return true, nil
	})
}
