package rockset_test

import (
	"github.com/rockset/rockset-go-client/option"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"

	"github.com/rockset/rockset-go-client"
)

func TestRockClient_ListCollections_all(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()
	log := zerolog.Ctx(ctx)

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	collections, err := rc.ListCollections(ctx)
	require.NoError(t, err)

	log.Debug().Int("count", len(collections)).Msg("collections")
}

func TestRockClient_ListCollections_ws(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()
	log := zerolog.Ctx(ctx)

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	collections, err := rc.ListCollections(ctx, option.WithWorkspace("commons"))
	require.NoError(t, err)

	log.Debug().Int("count", len(collections)).Msg("collections")
}
