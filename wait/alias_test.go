package wait_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/wait"
)

func TestWait_untilAliasReady(t *testing.T) {
	ctx := context.TODO()

	rs := fakeRocksetClient()
	rs.GetAliasReturnsOnCall(0, openapi.Alias{}, ErrNotFound)
	rs.GetAliasReturnsOnCall(1, openapi.Alias{}, nil)

	err := wait.New(&rs).UntilAliasAvailable(ctx, "workspace", "alias")
	assert.NoError(t, err)
	assert.Equal(t, 2, rs.GetAliasCallCount())
}

func TestWait_untilAliasGone(t *testing.T) {
	ctx := context.TODO()

	rs := fakeRocksetClient()
	rs.GetAliasReturnsOnCall(0, openapi.Alias{}, nil)
	rs.GetAliasReturnsOnCall(1, openapi.Alias{}, ErrNotFound)

	err := wait.New(&rs).UntilAliasGone(ctx, "workspace", "alias")
	assert.NoError(t, err)
	assert.Equal(t, 2, rs.GetAliasCallCount())
}
