package wait_test

import (
	"context"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/wait"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWait_untilAliasReady(t *testing.T) {
	ctx := context.TODO()

	rs := fakeRocksetClient()
	rs.GetAliasReturnsOnCall(0, openapi.Alias{}, NotFoundErr)
	rs.GetAliasReturnsOnCall(1, openapi.Alias{}, nil)

	w := wait.New(&rs)

	err := w.UntilAliasAvailable(ctx, "workspace", "alias")
	assert.NoError(t, err)
	assert.Equal(t, 2, rs.GetAliasCallCount())
}

func TestWait_untilAliasGone(t *testing.T) {
	ctx := context.TODO()

	rs := fakeRocksetClient()
	rs.GetAliasReturnsOnCall(0, openapi.Alias{}, nil)
	rs.GetAliasReturnsOnCall(1, openapi.Alias{}, NotFoundErr)

	w := wait.New(&rs)

	err := w.UntilAliasGone(ctx, "workspace", "alias")
	assert.NoError(t, err)
	assert.Equal(t, 2, rs.GetAliasCallCount())
}
