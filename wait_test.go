package rockset

import (
	"context"
	"errors"
	"github.com/rockset/rockset-go-client/fakes"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"strings"
	"testing"
)

func fakeError(code int) error {
	s := strings.Join(strings.Fields(http.StatusText(code)), "")
	return Error{
		ErrorModel: &openapi.ErrorModel{
			Type: &s,
		},
		Cause: errors.New("fake error"),
	}
}

func TestQueryHasStatus(t *testing.T) {
	ctx := context.TODO()

	f := &fakes.FakeQueryInfoGetter{}
	for i, status := range []QueryState{QueryQueued, QueryRunning, QueryCompleted} {
		v := string(status)
		f.GetQueryInfoReturnsOnCall(i, openapi.QueryInfo{Status: &v}, nil)
	}

	rc, err := NewClient(WithAPIKey("fake"), WithAPIServer("fake"))
	require.NoError(t, err)

	fn := queryHasStatus(ctx, f, "id", []QueryState{QueryCompleted})
	err = rc.RetryWithCheck(ctx, fn)
	assert.NoError(t, err)

	assert.Equal(t, 3, f.GetQueryInfoCallCount())
}

func TestCollectionHasState(t *testing.T) {
	ctx := context.TODO()

	f := &fakes.FakeCollectionGetter{}
	for i, status := range []string{"error", "foo", "bar", "foobar"} {
		s := status
		if i == 0 {
			f.GetCollectionReturnsOnCall(i, openapi.Collection{}, fakeError(http.StatusGatewayTimeout))
			continue
		}
		f.GetCollectionReturnsOnCall(i, openapi.Collection{Status: &s}, nil)
	}

	rc, err := NewClient(WithAPIKey("fake"), WithAPIServer("fake"))
	require.NoError(t, err)

	fn := collectionHasState(ctx, f, "ws", "c", "foobar")
	err = rc.RetryWithCheck(ctx, fn)
	assert.NoError(t, err)

	assert.Equal(t, 4, f.GetCollectionCallCount())
}
