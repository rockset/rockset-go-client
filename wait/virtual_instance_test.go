package wait_test

import (
	"context"
	rockerr "github.com/rockset/rockset-go-client/errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
	"github.com/rockset/rockset-go-client/wait"
)

func TestWait_untilVirtualInstanceActive(t *testing.T) {
	ctx := context.TODO()

	rs := fakeRocksetClient()
	rs.GetVirtualInstanceReturnsOnCall(0, openapi.VirtualInstance{State: openapi.PtrString(option.VirtualInstanceInitializing.String())}, nil)
	rs.GetVirtualInstanceReturnsOnCall(1, openapi.VirtualInstance{State: openapi.PtrString(option.VirtualInstanceActive.String())}, nil)

	err := wait.New(&rs).UntilVirtualInstanceActive(ctx, "id")
	assert.NoError(t, err)
	assert.Equal(t, 2, rs.GetVirtualInstanceCallCount())
}

func TestWait_untilVirtualInstanceSuspended(t *testing.T) {
	ctx := context.TODO()

	rs := fakeRocksetClient()
	rs.GetVirtualInstanceReturnsOnCall(0, openapi.VirtualInstance{State: openapi.PtrString(option.VirtualInstanceActive.String())}, nil)
	rs.GetVirtualInstanceReturnsOnCall(1, openapi.VirtualInstance{State: openapi.PtrString(option.VirtualInstanceSuspending.String())}, nil)
	rs.GetVirtualInstanceReturnsOnCall(2, openapi.VirtualInstance{State: openapi.PtrString(option.VirtualInstanceSuspended.String())}, nil)

	err := wait.New(&rs).UntilVirtualInstanceSuspended(ctx, "id")
	assert.NoError(t, err)
	assert.Equal(t, 3, rs.GetVirtualInstanceCallCount())
}

func TestWait_untilVirtualInstanceGone(t *testing.T) {
	ctx := context.TODO()

	rs := fakeRocksetClient()
	rs.GetVirtualInstanceReturnsOnCall(0, openapi.VirtualInstance{State: openapi.PtrString(option.VirtualInstanceActive.String())}, nil)
	rs.GetVirtualInstanceReturnsOnCall(1, openapi.VirtualInstance{}, NotFoundErr)

	err := wait.New(&rs).UntilVirtualInstanceGone(ctx, "id")
	assert.NoError(t, err)
	assert.Equal(t, 2, rs.GetVirtualInstanceCallCount())
}

func TestWait_untilMountActive(t *testing.T) {
	ctx := context.TODO()

	rs := fakeRocksetClient()
	rs.GetCollectionMountReturnsOnCall(0, openapi.CollectionMount{State: openapi.PtrString(option.MountCreating.String())}, nil)
	rs.GetCollectionMountReturnsOnCall(1, openapi.CollectionMount{State: openapi.PtrString(option.MountActive.String())}, nil)

	err := wait.New(&rs).UntilMountActive(ctx, "id", "workspace", "collection")
	assert.NoError(t, err)
	assert.Equal(t, 2, rs.GetCollectionMountCallCount())
}

func TestWait_untilMountGone404(t *testing.T) {
	ctx := context.TODO()

	rs := fakeRocksetClient()
	rs.GetCollectionMountReturnsOnCall(0, openapi.CollectionMount{State: openapi.PtrString(option.MountActive.String())}, nil)
	rs.GetCollectionMountReturnsOnCall(1, openapi.CollectionMount{State: openapi.PtrString(option.MountDeleting.String())}, nil)
	rs.GetCollectionMountReturnsOnCall(2, openapi.CollectionMount{}, NotFoundErr)

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
	rs.GetCollectionMountReturnsOnCall(0, openapi.CollectionMount{State: openapi.PtrString(option.MountActive.String())}, nil)
	rs.GetCollectionMountReturnsOnCall(1, openapi.CollectionMount{State: openapi.PtrString(option.MountDeleting.String())}, nil)
	rs.GetCollectionMountReturnsOnCall(2, openapi.CollectionMount{}, e404)

	err := wait.New(&rs).UntilMountGone(ctx, "id", "workspace", "collection")
	assert.NoError(t, err)
	assert.Equal(t, 3, rs.GetCollectionMountCallCount())
}
