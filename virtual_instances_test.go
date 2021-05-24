package rockset_test

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/rockset/rockset-go-client"
)

func TestRockClient_ListVirtualInstances(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()
	log := zerolog.Ctx(ctx)

	rc, err := rockset.NewClient(rockset.FromEnv())
	require.NoError(t, err)

	vis, err := rc.ListVirtualInstances(ctx)
	require.NoError(t, err)

	assert.NotEmpty(t, vis)

	for _, vi := range vis {
		log.Debug().Str("id", vi.GetId()).Str("state", vi.GetState()).
			Str("type", vi.GetCurrentType()).Str("size", vi.GetCurrentSize()).
			Msg("virtual instance")
	}
}
