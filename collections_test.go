package rockset_test

import (
	"github.com/stretchr/testify/suite"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/option"
)

type CollectionTestSuite struct {
	suite.Suite
	rc *rockset.RockClient
}

func TestCollectionTestSuite(t *testing.T) {
	skipUnlessIntegrationTest(t)
	suite.Run(t, &CollectionTestSuite{})
}

func (s *CollectionTestSuite) SetupSuite() {
	rc, err := rockset.NewClient()
	s.NoError(err)
	s.rc = rc
}

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
