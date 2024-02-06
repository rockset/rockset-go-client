package wait_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/wait"
)

func TestWaiter_UntilQueryable(t *testing.T) {
	ctx := context.TODO()

	rs := fakeRocksetClient()
	rs.GetCollectionCommitReturnsOnCall(0, openapi.GetCollectionCommitData{
		Passed: openapi.PtrBool(false),
	}, nil)
	rs.GetCollectionCommitReturnsOnCall(1, openapi.GetCollectionCommitData{
		Passed: openapi.PtrBool(true),
	}, nil)

	err := wait.New(&rs).UntilQueryable(ctx, "ws", "ql", []string{"foo", "bar"})
	assert.NoError(t, err)
	assert.Equal(t, 2, rs.GetCollectionCommitCallCount())
}
