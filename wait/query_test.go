package wait_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
	"github.com/rockset/rockset-go-client/wait"
)

func TestWait_untilQueryDone(t *testing.T) {
	ctx := context.TODO()

	rs := fakeRocksetClient()
	rs.GetQueryInfoReturnsOnCall(0, openapi.QueryInfo{Status: openapi.PtrString(option.QueryRunning.String())}, nil)
	rs.GetQueryInfoReturnsOnCall(1, openapi.QueryInfo{Status: openapi.PtrString(option.QueryCompleted.String())}, nil)

	err := wait.New(&rs).UntilQueryDone(ctx, "id")
	assert.NoError(t, err)
	assert.Equal(t, 2, rs.GetQueryInfoCallCount())
}
