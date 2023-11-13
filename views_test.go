package rockset_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/rockset/rockset-go-client/internal/test"
	"github.com/rockset/rockset-go-client/option"
)

func TestListViews(t *testing.T) {
	rc, _ := vcrTestClient(t, t.Name())
	ctx := test.Context()

	_, err := rc.ListViews(ctx)
	require.NoError(t, err)

	_, err = rc.ListViews(ctx, option.WithViewWorkspace(test.PersistentWorkspace))
	require.NoError(t, err)
}

func TestViewCRUD(t *testing.T) {
	rc, randomName := vcrTestClient(t, t.Name())
	ctx := test.Context()

	ws := "acc"
	name := randomName("view")
	query := "select * from commons._events where _events.kind = 'COLLECTION'"
	view, err := rc.CreateView(ctx, ws, name, query, option.WithViewDescription("test"))
	require.NoError(t, err)
	assert.Equal(t, name, view.GetName())

	_, err = rc.UpdateView(ctx, ws, name, query, option.WithViewDescription(test.Description()))
	require.NoError(t, err)

	view, err = rc.GetView(ctx, ws, name)
	require.NoError(t, err)
	assert.Equal(t, name, view.GetName())

	err = rc.DeleteView(ctx, ws, name)
	require.NoError(t, err)
}
