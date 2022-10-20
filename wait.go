package rockset

import (
	"context"
	"errors"
	"fmt"
	"github.com/rockset/rockset-go-client/option"
	"github.com/rs/zerolog"
)

const (
	collectionStatusREADY = "READY"
)

// WaitUntilKafkaIntegrationActive waits until all topics in a Kafka integration are in ACTIVE state.
func (rc *RockClient) WaitUntilKafkaIntegrationActive(ctx context.Context, integration string) error {
	return rc.RetryWithCheck(ctx, func() (bool, error) {
		zl := zerolog.Ctx(ctx)

		i, err := rc.GetIntegration(ctx, integration)
		if err != nil {
			return false, err
		}

		if i.Kafka == nil {
			return false, fmt.Errorf("not a kafka integration: %s", integration)
		}

		if i.Kafka.SourceStatusByTopic == nil {
			// for v3 integrations this will be nil
			zl.Trace().Bool("v3", i.Kafka.GetUseV3()).Msg("no topics found")
			return false, nil
		}

		var allActive = true
		for topic, status := range *i.Kafka.SourceStatusByTopic {
			zl.Trace().Str("state", status.GetState()).Str("topic", topic).Send()
			if status.GetState() != string(option.KafkaIntegrationActive) {
				allActive = false
			}
		}

		return !allActive, nil
	})
}

// WaitUntilAliasAvailable waits until the alias is available.
func (rc *RockClient) WaitUntilAliasAvailable(ctx context.Context, workspace, alias string) error {
	return rc.RetryWithCheck(ctx, func() (bool, error) {
		_, err := rc.GetAlias(ctx, workspace, alias)

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

// WaitUntilQueryCompleted waits until queryID has either completed, errored, or been cancelled.
func (rc *RockClient) WaitUntilQueryCompleted(ctx context.Context, queryID string) error {
	// TODO should this only wait for COMPLETED and return an error for ERROR and CANCELLED?
	return rc.RetryWithCheck(ctx, rc.queryHasStatus(ctx, queryID, []QueryState{QueryCompleted, QueryError, QueryCancelled}))
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

// WaitUntilCollectionHasNewDocuments waits until the collection has at least count new documents
// (measured from when the method is called).
func (rc *RockClient) WaitUntilCollectionHasNewDocuments(ctx context.Context, workspace, name string, count int64) error {
	waiter := docWaiter{rc: rc}
	return rc.RetryWithCheck(ctx, waiter.collectionHasNewDocs(ctx, workspace, name, count))
}

// WaitUntilCollectionHasDocuments waits until the collection has at least count documents
func (rc *RockClient) WaitUntilCollectionHasDocuments(ctx context.Context, workspace, name string, count int64) error {
	return rc.RetryWithCheck(ctx, rc.collectionHasDocs(ctx, workspace, name, count))
}

// WaitUntilWorkspaceAvailable waits until the workspace is available.
func (rc *RockClient) WaitUntilWorkspaceAvailable(ctx context.Context, workspace string) error {
	return rc.RetryWithCheck(ctx, rc.workspaceIsAvailable(ctx, workspace))
}

// TODO(pme) refactor viewIsGone() and collectionIsGone() to be DRY

func (rc *RockClient) queryHasStatus(ctx context.Context, queryID string, statuses []QueryState) RetryCheck {
	return func() (bool, error) {
		res, err := rc.GetQueryInfo(ctx, queryID)
		if err != nil {
			return false, err
		}

		for _, s := range statuses {
			if string(s) == res.GetStatus() {
				return false, nil
			}
		}

		return true, nil
	}
}

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
		zl := zerolog.Ctx(ctx)
		_, err := rc.GetView(ctx, workspace, name)

		if err == nil {
			zl.Debug().Str("workspace", workspace).Str("view", name).Msg("view is still present")
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
		zl := zerolog.Ctx(ctx)
		_, err := rc.GetCollection(ctx, workspace, name)

		if err == nil {
			zl.Debug().Str("workspace", workspace).Str("collection", name).Msg("collection is still present")
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
		zl := zerolog.Ctx(ctx)
		c, err := rc.GetCollection(ctx, workspace, name)
		if err != nil {
			re := NewError(err)
			if re.Retryable() {
				return true, nil
			}

			return false, err
		}

		zl.Debug().Str("status", c.GetStatus()).Str("workspace", workspace).
			Str("desired", c.GetStatus()).Str("collection", name).Msg("collectionHasState()")
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
		zl := zerolog.Ctx(ctx)
		c, err := d.rc.GetCollection(ctx, workspace, name)
		if err != nil {
			re := NewError(err)
			if re.Retryable() {
				return true, nil
			}

			return false, err
		}

		current := c.Stats.GetDocCount()
		zl.Debug().Str("workspace", workspace).Int64("current", current).
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

func (rc *RockClient) collectionHasDocs(ctx context.Context, workspace, name string, count int64) RetryCheck {
	return func() (bool, error) {
		zl := zerolog.Ctx(ctx)
		c, err := rc.GetCollection(ctx, workspace, name)
		if err != nil {
			re := NewError(err)
			if re.Retryable() {
				return true, nil
			}

			return false, err
		}

		current := c.Stats.GetDocCount()
		zl.Debug().Str("workspace", workspace).Int64("current", current).Str("collection", name).
			Int64("count", count).Msg("collectionHasDocs()")

		if current >= count {
			return false, nil
		}

		return true, nil
	}
}
