package wait_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	rockerr "github.com/rockset/rockset-go-client/errors"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
	"github.com/rockset/rockset-go-client/wait"
)

func TestWait_untilVirtualInstanceActive(t *testing.T) {
	ctx := context.TODO()

	rs := fakeRocksetClient()
	rs.GetVirtualInstanceReturnsOnCall(0, openapi.VirtualInstance{
		State: stringPtr(option.VirtualInstanceInitializing)}, nil)
	rs.GetVirtualInstanceReturnsOnCall(1, openapi.VirtualInstance{
		State: stringPtr(option.VirtualInstanceActive)}, nil)

	err := wait.New(&rs).UntilVirtualInstanceActive(ctx, "id")
	assert.NoError(t, err)
	assert.Equal(t, 2, rs.GetVirtualInstanceCallCount())
}

func TestWait_untilVirtualInstanceSuspended(t *testing.T) {
	ctx := context.TODO()

	rs := fakeRocksetClient()
	rs.GetVirtualInstanceReturnsOnCall(0, openapi.VirtualInstance{
		State: stringPtr(option.VirtualInstanceActive)}, nil)
	rs.GetVirtualInstanceReturnsOnCall(1, openapi.VirtualInstance{
		State: stringPtr(option.VirtualInstanceSuspending)}, nil)
	rs.GetVirtualInstanceReturnsOnCall(2, openapi.VirtualInstance{
		State: stringPtr(option.VirtualInstanceSuspended)}, nil)

	err := wait.New(&rs).UntilVirtualInstanceSuspended(ctx, "id")
	assert.NoError(t, err)
	assert.Equal(t, 3, rs.GetVirtualInstanceCallCount())
}

func TestWait_untilVirtualInstanceGone(t *testing.T) {
	ctx := context.TODO()

	rs := fakeRocksetClient()
	rs.GetVirtualInstanceReturnsOnCall(0, openapi.VirtualInstance{
		State: stringPtr(option.VirtualInstanceActive)}, nil)
	rs.GetVirtualInstanceReturnsOnCall(1, openapi.VirtualInstance{}, ErrNotFound)

	err := wait.New(&rs).UntilVirtualInstanceGone(ctx, "id")
	assert.NoError(t, err)
	assert.Equal(t, 2, rs.GetVirtualInstanceCallCount())
}

func TestWait_untilMountActive(t *testing.T) {
	ctx := context.TODO()

	rs := fakeRocksetClient()
	rs.GetCollectionMountReturnsOnCall(0, openapi.CollectionMount{State: stringPtr(option.MountCreating)}, nil)
	rs.GetCollectionMountReturnsOnCall(1, openapi.CollectionMount{State: stringPtr(option.MountActive)}, nil)

	err := wait.New(&rs).UntilMountActive(ctx, "id", "workspace", "collection")
	assert.NoError(t, err)
	assert.Equal(t, 2, rs.GetCollectionMountCallCount())
}

func TestWait_untilMountGone404(t *testing.T) {
	ctx := context.TODO()

	rs := fakeRocksetClient()
	rs.GetCollectionMountReturnsOnCall(0, openapi.CollectionMount{State: stringPtr(option.MountActive)}, nil)
	rs.GetCollectionMountReturnsOnCall(1, openapi.CollectionMount{State: stringPtr(option.MountDeleting)}, nil)
	rs.GetCollectionMountReturnsOnCall(2, openapi.CollectionMount{}, ErrNotFound)

	err := wait.New(&rs).UntilMountGone(ctx, "id", "workspace", "collection")
	assert.NoError(t, err)
	assert.Equal(t, 3, rs.GetCollectionMountCallCount())
}

func TestWait_untilMountGone400(t *testing.T) {
	// TODO(pme) once ORC-3594 is fixed, this test can be removed
	ctx := context.TODO()

	msg := "Collection persistent.patch is not mounted on virtual instance 3b5d1aeb-5645-495e-a328-adee6a037824."
	e404 := rockerr.Error{ErrorModel: openapi.NewErrorModel(), StatusCode: http.StatusBadRequest}
	e404.ErrorModel.Message = &msg

	rs := fakeRocksetClient()
	rs.GetCollectionMountReturnsOnCall(0, openapi.CollectionMount{State: stringPtr(option.MountActive)}, nil)
	rs.GetCollectionMountReturnsOnCall(1, openapi.CollectionMount{State: stringPtr(option.MountDeleting)}, nil)
	rs.GetCollectionMountReturnsOnCall(2, openapi.CollectionMount{}, e404)

	err := wait.New(&rs).UntilMountGone(ctx, "id", "workspace", "collection")
	assert.NoError(t, err)
	assert.Equal(t, 3, rs.GetCollectionMountCallCount())
}
