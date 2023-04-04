package rockset_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/dataset"
	"github.com/rockset/rockset-go-client/option"
)

type CollectionTestSuite struct {
	suite.Suite
	rc          *rockset.RockClient
	ws          string
	collections []string
}

func TestCollectionIntegrationSuite(t *testing.T) {
	skipUnlessIntegrationTest(t)
	t.Parallel()

	suite.Run(t, &CollectionTestSuite{rc: testClient(t), ws: randomName("collection")})
}

func (s *CollectionTestSuite) SetupSuite() {
	ctx := testCtx()
	_, err := s.rc.CreateWorkspace(ctx, s.ws)
	s.NoError(err)
}

func (s *CollectionTestSuite) TearDownSuite() {
	ctx := testCtx()

	for _, c := range s.collections {
		err := s.rc.WaitUntilCollectionGone(ctx, s.ws, c)
		s.Assert().NoError(err)
	}
	err := s.rc.DeleteWorkspace(ctx, s.ws)
	s.NoError(err)
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

func (s *CollectionTestSuite) TestCreateSampleCitiesCollection() {
	ctx := testCtx()
	name := randomName("cities")
	s.collections = append(s.collections, name)

	_, err := s.rc.CreateCollection(ctx, s.ws, name,
		option.WithSampleDataset(dataset.Cities))
	s.Require().NoError(err)

	err = s.rc.WaitUntilCollectionHasDocuments(ctx, s.ws, name, int64(145_658))
	s.Assert().NoError(err)

	err = s.rc.DeleteCollection(ctx, s.ws, name)
	s.Assert().NoError(err)

}

func (s *CollectionTestSuite) TestCreateSampleMoviesCollection() {
	ctx := testCtx()
	name := randomName("movies")
	s.collections = append(s.collections, name)

	_, err := s.rc.CreateCollection(ctx, s.ws, name,
		option.WithSampleDatasetPattern("movies/*"))
	s.Require().NoError(err)

	err = s.rc.WaitUntilCollectionHasDocuments(ctx, s.ws, name, int64(2_830))
	s.Assert().NoError(err)

	err = s.rc.DeleteCollection(ctx, s.ws, name)
	s.Assert().NoError(err)
}
