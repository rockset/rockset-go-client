package rockset_test

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/option"
)

func TestRockClient_ListAliases(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()
	log := zerolog.Ctx(ctx)

	rc, err := rockset.NewClient(rockset.FromEnv())
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

	rc, err := rockset.NewClient(rockset.FromEnv())
	require.NoError(t, err)

	aliases, err := rc.ListAliases(ctx, option.WithAliasWorkspace("common"))
	require.NoError(t, err)
	for _, a := range aliases {
		log.Printf("workspace: %s", a.GetName())
	}
}
