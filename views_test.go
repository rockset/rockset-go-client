package rockset_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/rockset/rockset-go-client/internal/test"
	"github.com/rockset/rockset-go-client/option"
)

func TestListViews(t *testing.T) {
	rc, _ := vcrTestClient(t, t.Name())
	ctx := test.Context()

	_, err := rc.ListViews(ctx, option.WithViewWorkspace(test.PersistentWorkspace))
	require.NoError(t, err)
}

func TestViewCRUD(t *testing.T) {
	rc, _ := vcrTestClient(t, t.Name())
	ctx := test.Context()

	ws := "acc"
	name := test.RandomName("view")
	query := "select * from commons._events where _events.kind = 'COLLECTION'"
	_, err := rc.CreateView(ctx, ws, name, query)
	require.NoError(t, err)

	_, err = rc.UpdateView(ctx, ws, name, query, option.WithViewDescription(test.Description()))
	require.NoError(t, err)

	err = rc.DeleteView(ctx, ws, name)
	require.NoError(t, err)
}
