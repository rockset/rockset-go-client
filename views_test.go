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

	_, err = rc.ListViews(ctx, option.WithViewWorkspace("commons"))
	require.NoError(t, err)
}

func TestViewCRUD(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()
	rc, err := rockset.NewClient()
	require.NoError(t, err)

	ws := "commons"
	randomName := "not-so-random-name"
	query := "select * from commons._events where _events.kind = 'COLLECTION'"
	_, err = rc.CreateView(ctx, ws, randomName, query)
	require.NoError(t, err)

	_, err = rc.UpdateView(ctx, ws, randomName, query, option.WithViewDescription("description"))
	require.NoError(t, err)

	err = rc.DeleteView(ctx, ws, randomName)
	require.NoError(t, err)
}
