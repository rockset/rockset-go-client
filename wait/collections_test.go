package wait_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
	"github.com/rockset/rockset-go-client/wait"
)

func TestWait_untilCollectionReady(t *testing.T) {
	ctx := context.TODO()

	rs := fakeRocksetClient()
	rs.GetCollectionReturnsOnCall(0, openapi.Collection{Status: openapi.PtrString(option.CollectionStatusInitialized)}, nil)
	rs.GetCollectionReturnsOnCall(1, openapi.Collection{Status: openapi.PtrString(option.CollectionStatusCreated)}, nil)
	rs.GetCollectionReturnsOnCall(2, openapi.Collection{Status: openapi.PtrString(option.CollectionStatusReady)}, nil)

	err := wait.New(&rs).UntilCollectionReady(ctx, "workspace", "collection")
	assert.NoError(t, err)
	assert.Equal(t, 3, rs.GetCollectionCallCount())
}

func TestWait_untilCollectionGone(t *testing.T) {
	ctx := context.TODO()

	rs := fakeRocksetClient()
	rs.GetCollectionReturnsOnCall(0, openapi.Collection{Status: openapi.PtrString(option.CollectionStatusReady)}, nil)
	rs.GetCollectionReturnsOnCall(1, openapi.Collection{}, NotFoundErr)

	err := wait.New(&rs).UntilCollectionGone(ctx, "workspace", "collection")
	assert.NoError(t, err)
	assert.Equal(t, 2, rs.GetCollectionCallCount())
}

func TestWait_untilCollectionHasNewDocuments(t *testing.T) {
	ctx := context.TODO()

	rs := fakeRocksetClient()
	for i, c := range []int64{0, 1, 2} {
		var count = c
		rs.GetCollectionReturnsOnCall(i, openapi.Collection{Stats: &openapi.CollectionStats{DocCount: &count}}, nil)
	}

	err := wait.New(&rs).UntilCollectionHasNewDocuments(ctx, "workspace", "collection", 2)
	assert.NoError(t, err)
	assert.Equal(t, 3, rs.GetCollectionCallCount())
}

func TestWait_untilCollectionHasDocuments(t *testing.T) {
	ctx := context.TODO()

	rs := fakeRocksetClient()
	for i, c := range []int64{10, 10, 12} {
		var count = c
		rs.GetCollectionReturnsOnCall(i, openapi.Collection{Stats: &openapi.CollectionStats{DocCount: &count}}, nil)
	}

	err := wait.New(&rs).UntilCollectionHasDocuments(ctx, "workspace", "collection", 12)
	assert.NoError(t, err)
	assert.Equal(t, 3, rs.GetCollectionCallCount())
}
