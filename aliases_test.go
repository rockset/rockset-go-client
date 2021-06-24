package rockset_test

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/option"
)

func TestRockClient_GetAlias(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	alias, err := rc.GetAlias(ctx, "commons", "getalias")
	require.NoError(t, err)
	assert.Equal(t, "pme@rockset.com", alias.GetCreatorEmail())
}

func TestRockClient_ListAliases(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()
	log := zerolog.Ctx(ctx)

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	aliases, err := rc.ListAliases(ctx)
	require.NoError(t, err)
	for _, a := range aliases {
		log.Printf("workspace: %s", a.GetName())
	}
}

func TestRockClient_ListAliasesForWorkspace(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()
	log := zerolog.Ctx(ctx)

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	aliases, err := rc.ListAliases(ctx, option.WithAliasWorkspace("common"))
	require.NoError(t, err)
	for _, a := range aliases {
		log.Printf("workspace: %s", a.GetName())
	}
}

func TestRockClient_Aliases(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	ws := "commons"
	alias := "randomstring"

	_, err = rc.CreateAlias(ctx, ws, alias, []string{"commons.writetest"})
	require.NoError(t, err)

	err = rc.WaitUntilAliasAvailable(ctx, ws, alias)
	require.NoError(t, err)

	// update

	err = rc.WaitUntilAliasAvailable(ctx, ws, alias)
	require.NoError(t, err)

	err = rc.DeleteAlias(ctx, ws, alias)
	require.NoError(t, err)
}
