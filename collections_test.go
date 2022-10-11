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

func (s *CollectionTestSuite) TestListAllCollections(t *testing.T) {
	ctx := testCtx()
	log := zerolog.Ctx(ctx)

	collections, err := s.rc.ListCollections(ctx)
	require.NoError(t, err)

	log.Debug().Int("count", len(collections)).Msg("collections")
}

func (s *CollectionTestSuite) TestListCollectionsInWorkspace(t *testing.T) {
	ctx := testCtx()
	log := zerolog.Ctx(ctx)

	collections, err := s.rc.ListCollections(ctx, option.WithWorkspace("commons"))
	require.NoError(t, err)

	log.Debug().Int("count", len(collections)).Msg("collections")
}
