package wait_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	rockerr "github.com/rockset/rockset-go-client/errors"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
	"github.com/rockset/rockset-go-client/wait"
)

func TestWait_untilQueryLambdaActive(t *testing.T) {
	ctx := context.TODO()

	rs := fakeRocksetClient()
	rs.GetQueryLambdaVersionReturnsOnCall(0, openapi.QueryLambdaVersion{State: &emptyString}, nil)
	rs.GetQueryLambdaVersionReturnsOnCall(1, openapi.QueryLambdaVersion{
		State: stringPtr(option.QueryLambdaActive)}, nil)

	err := wait.New(&rs).UntilQueryLambdaVersionActive(ctx, "ws", "ql", "v")
	assert.NoError(t, err)
	assert.Equal(t, 2, rs.GetQueryLambdaVersionCallCount())
}

func TestWait_untilQueryLambdaActive_invalid(t *testing.T) {
	ctx := context.TODO()

	rs := fakeRocksetClient()
	rs.GetQueryLambdaVersionReturnsOnCall(0, openapi.QueryLambdaVersion{State: &emptyString}, nil)
	rs.GetQueryLambdaVersionReturnsOnCall(1, openapi.QueryLambdaVersion{
		State: stringPtr(option.QueryLambdaInvalidSQL)}, nil)

	err := wait.New(&rs).UntilQueryLambdaVersionActive(ctx, "ws", "ql", "v")
	assert.ErrorIs(t, err, rockerr.ErrBadWaitState)
	assert.Equal(t, 2, rs.GetQueryLambdaVersionCallCount())
}
