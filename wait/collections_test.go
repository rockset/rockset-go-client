package wait_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/wait"
)

func TestWait_untilCollectionReady(t *testing.T) {
	ctx := context.TODO()

	rs := fakeRocksetClient()
	rs.GetCollectionReturnsOnCall(0, openapi.Collection{Status: openapi.PtrString(rockset.CollectionStatusInitialized)}, nil)
	rs.GetCollectionReturnsOnCall(1, openapi.Collection{Status: openapi.PtrString(rockset.CollectionStatusCreated)}, nil)
	rs.GetCollectionReturnsOnCall(2, openapi.Collection{Status: openapi.PtrString(rockset.CollectionStatusReady)}, nil)

	w := wait.New(&rs)

	err := w.UntilCollectionReady(ctx, "workspace", "collection")
	assert.NoError(t, err)
	assert.Equal(t, 3, rs.GetCollectionCallCount())
}

func TestWait_untilCollectionGone(t *testing.T) {
	ctx := context.TODO()

	rs := fakeRocksetClient()
	rs.GetCollectionReturnsOnCall(0, openapi.Collection{Status: openapi.PtrString(rockset.CollectionStatusReady)}, nil)
	rs.GetCollectionReturnsOnCall(1, openapi.Collection{}, NotFoundErr)
	w := wait.New(&rs)

	err := w.UntilCollectionGone(ctx, "workspace", "collection")
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
	w := wait.New(&rs)

	err := w.UntilCollectionHasNewDocuments(ctx, "workspace", "collection", 2)
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
	w := wait.New(&rs)

	err := w.UntilCollectionHasDocuments(ctx, "workspace", "collection", 12)
	assert.NoError(t, err)
	assert.Equal(t, 3, rs.GetCollectionCallCount())
}
