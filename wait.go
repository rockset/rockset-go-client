package rockset

import (
	"context"
	"errors"
	"fmt"
	"github.com/rockset/rockset-go-client/retry"

	"github.com/rs/zerolog"

	rockerr "github.com/rockset/rockset-go-client/errors"
	"github.com/rockset/rockset-go-client/option"
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

// WaitUntilMountActive waits until the collection mount is active, and queries can be issued to it on the
// virtual instance.
func (rc *RockClient) WaitUntilMountActive(ctx context.Context, vID, workspace, collection string) error {
	return rc.RetryWithCheck(ctx, resourceHasState(ctx, []string{MountActive},
		func(ctx context.Context) (string, error) {
			cm, err := rc.GetCollectionMount(ctx, vID, workspace+"."+collection)
			return cm.GetState(), err
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

// resourceIsAvailable implements RetryFn to wait until the resource is present
func resourceIsAvailable(ctx context.Context, fn func(ctx context.Context) error) retry.CheckFn {
	return func() (bool, error) {
		err := fn(ctx)

		if err == nil {
			// the resource is present, done
			return false, nil
		}

		var re rockerr.Error
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
func resourceIsGone(ctx context.Context, fn func(ctx context.Context) error) retry.CheckFn {
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
		}

		return false, err
	}
}

// resourceHasState implements RetryFn to wait until the resource is has the desired state
func resourceHasState[T comparable](ctx context.Context, states []T, fn func(ctx context.Context) (T, error)) retry.CheckFn {
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
