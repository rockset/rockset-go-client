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

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	suite.Run(t, &CollectionTestSuite{rc: rc})
}

func (s *CollectionTestSuite) TestGetCollection() {
	ctx := testCtx()

	const cName = "_events"
	collection, err := s.rc.GetCollection(ctx, "commons", cName)
	s.NoError(err)
	s.Assert().Equal(cName, collection.GetName())
}

func (s *CollectionTestSuite) TestListAllCollections() {
	ctx := testCtx()
	log := zerolog.Ctx(ctx)

	collections, err := s.rc.ListCollections(ctx)
	s.NoError(err)

	log.Debug().Int("count", len(collections)).Msg("collections")
}

func (s *CollectionTestSuite) TestListCollectionsInWorkspace() {
	ctx := testCtx()
	log := zerolog.Ctx(ctx)

	collections, err := s.rc.ListCollections(ctx, option.WithWorkspace("commons"))
	s.NoError(err)

	log.Debug().Int("count", len(collections)).Msg("collections")
}
