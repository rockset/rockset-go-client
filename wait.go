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
	return rc.RetryWithCheck(ctx, resourceIsAvailable(ctx, func(ctx context.Context) error {
		_, err := rc.GetAlias(ctx, workspace, alias)
		return err
	}))
}

// WaitUntilAliasGone waits until the alias is gone.
func (rc *RockClient) WaitUntilAliasGone(ctx context.Context, workspace, alias string) error {
	return rc.RetryWithCheck(ctx, resourceIsGone(ctx, func(ctx context.Context) error {
		_, err := rc.GetAlias(ctx, workspace, alias)
		return err
	}))
}

// WaitUntilQueryCompleted waits until queryID has either completed, errored, or been cancelled.
func (rc *RockClient) WaitUntilQueryCompleted(ctx context.Context, queryID string) error {
	// TODO should this only wait for COMPLETED and return an error for ERROR and CANCELLED?
	return rc.RetryWithCheck(ctx, resourceHasState(ctx, []QueryState{QueryCompleted, QueryError, QueryCancelled},
		func(ctx context.Context) (QueryState, error) {
			q, err := rc.GetQueryInfo(ctx, queryID)
			return QueryState(q.GetStatus()), err
		}))
}

// WaitUntilCollectionReady waits until the collection is ready.
func (rc *RockClient) WaitUntilCollectionReady(ctx context.Context, workspace, name string) error {
	return rc.RetryWithCheck(ctx, resourceHasState(ctx, []string{collectionStatusREADY},
		func(ctx context.Context) (string, error) {
			c, err := rc.GetCollection(ctx, workspace, name)
			return c.GetStatus(), err
		}))
}

func (rc *RockClient) WaitUntilVirtualInstanceActive(ctx context.Context, id string) error {
	return rc.RetryWithCheck(ctx, resourceHasState(ctx, []string{VirtualInstanceActive},
		func(ctx context.Context) (string, error) {
			vi, err := rc.GetVirtualInstance(ctx, id)
			return vi.GetState(), err
		}))
}

func (rc *RockClient) WaitUntilVirtualInstanceSuspended(ctx context.Context, id string) error {
	return rc.RetryWithCheck(ctx, resourceHasState(ctx, []string{VirtualInstanceSuspended},
		func(ctx context.Context) (string, error) {
			vi, err := rc.GetVirtualInstance(ctx, id)
			return vi.GetState(), err
		}))
}

// WaitUntilCollectionGone waits until a collection marked for deletion is gone, i.e. GetCollection()
// returns "not found".
func (rc *RockClient) WaitUntilCollectionGone(ctx context.Context, workspace, name string) error {
	return rc.RetryWithCheck(ctx, resourceIsGone(ctx, func(ctx context.Context) error {
		_, err := rc.GetCollection(ctx, workspace, name)
		return err
	}))
}

// WaitUntilViewGone waits until a view marked for deletion is gone, i.e. GetView()
// returns "not found".
func (rc *RockClient) WaitUntilViewGone(ctx context.Context, workspace, name string) error {
	return rc.RetryWithCheck(ctx, resourceIsGone(ctx, func(ctx context.Context) error {
		_, err := rc.GetView(ctx, workspace, name)
		return err
	}))
}

// WaitUntilWorkspaceAvailable waits until the workspace is available.
func (rc *RockClient) WaitUntilWorkspaceAvailable(ctx context.Context, workspace string) error {
	return rc.RetryWithCheck(ctx, resourceIsAvailable(ctx, func(ctx context.Context) error {
		_, err := rc.GetWorkspace(ctx, workspace)
		return err
	}))
}

// WaitUntilWorkspaceGone waits until the workspace is gone, i.e. GetWorkspace() returns "not found".
func (rc *RockClient) WaitUntilWorkspaceGone(ctx context.Context, workspace string) error {
	return rc.RetryWithCheck(ctx, resourceIsGone(ctx, func(ctx context.Context) error {
		_, err := rc.GetWorkspace(ctx, workspace)
		return err
	}))
}

// WaitUntilCollectionHasNewDocuments waits until the collection has at least count new documents
// (measured from when the method is called).
func (rc *RockClient) WaitUntilCollectionHasNewDocuments(ctx context.Context, workspace, name string,
	count int64) error {
	waiter := docWaiter{rc: rc}
	return rc.RetryWithCheck(ctx, waiter.collectionHasNewDocs(ctx, workspace, name, count))
}

// WaitUntilCollectionHasDocuments waits until the collection has at least count documents
func (rc *RockClient) WaitUntilCollectionHasDocuments(ctx context.Context, workspace, name string, count int64) error {
	waiter := docWaiter{rc: rc}
	return rc.RetryWithCheck(ctx, waiter.collectionHasDocs(ctx, workspace, name, count))
}

// resourceIsAvailable implements RetryFn to wait until the resource is present
func resourceIsAvailable(ctx context.Context, fn func(ctx context.Context) error) RetryCheck {
	return func() (bool, error) {
		err := fn(ctx)

		if err == nil {
			// the resource is present, done
			return false, nil
		}

		var re Error
		if errors.As(err, &re) {
			if re.IsNotFoundError() {
				// the resource is not present, retry
				return true, nil
			}
		}

		// we got an error, stop
		return false, err
	}
}

// resourceIsGone implements RetryFn to wait until the resource is gone
func resourceIsGone(ctx context.Context, fn func(ctx context.Context) error) RetryCheck {
	return func() (bool, error) {
		err := fn(ctx)

		if err == nil {
			// the resource is still present, retry the operation
			return true, nil
		}

		var re Error
		if errors.As(err, &re) {
			if re.IsNotFoundError() {
				// the resource is no longer present
				return false, nil
			}
		}

		return false, err
	}
}

// resourceHasState implements RetryFn to wait until the resource is has the desired state
func resourceHasState[T comparable](ctx context.Context, states []T, fn func(ctx context.Context) (T, error)) RetryCheck {
	return func() (bool, error) {
		zl := zerolog.Ctx(ctx)
		state, err := fn(ctx)
		if err != nil {
			return false, err
		}

		for _, s := range states {
			if state == s {
				return false, nil
			}
		}

		zl.Trace().Any("current", state).Msg("waiting for resource state")

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

func (d *docWaiter) collectionHasDocs(ctx context.Context, workspace, name string, count int64) RetryCheck {
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

		if current >= count {
			return false, nil
		}

		return true, nil
	}
}
