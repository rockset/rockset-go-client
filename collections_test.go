package rockset_test

import (
	"github.com/stretchr/testify/suite"
	"testing"

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

	collection, err := s.rc.GetCollection(ctx, persistentWorkspace, persistentCollection)
	s.NoError(err)
	s.Assert().Equal(persistentCollection, collection.GetName())
}

func (s *CollectionTestSuite) TestListAllCollections() {
	ctx := testCtx()

	collections, err := s.rc.ListCollections(ctx)
	s.NoError(err)

	s.T().Logf("collections: %d", len(collections))
}

func (s *CollectionTestSuite) TestListCollectionsInWorkspace() {
	ctx := testCtx()

	collections, err := s.rc.ListCollections(ctx, option.WithWorkspace(persistentWorkspace))
	s.NoError(err)

	s.T().Logf("collections in %s: %d", persistentWorkspace, len(collections))
}
