package rockset_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/option"
)

func TestListViews(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()
	rc, err := rockset.NewClient()
	require.NoError(t, err)

	_, err = rc.ListViews(ctx, option.WithViewWorkspace(persistentWorkspace))
	require.NoError(t, err)
}

func TestViewCRUD(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()
	rc, err := rockset.NewClient()
	require.NoError(t, err)

	ws := "acc"
	name := randomName("view")
	query := "select * from commons._events where _events.kind = 'COLLECTION'"
	_, err = rc.CreateView(ctx, ws, name, query)
	require.NoError(t, err)

	_, err = rc.UpdateView(ctx, ws, name, query, option.WithViewDescription(description()))
	require.NoError(t, err)

	err = rc.DeleteView(ctx, ws, name)
	require.NoError(t, err)
}
