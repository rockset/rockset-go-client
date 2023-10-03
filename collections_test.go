package rockset_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/dataset"
	"github.com/rockset/rockset-go-client/internal/test"
	"github.com/rockset/rockset-go-client/option"
)

type CollectionTestSuite struct {
	suite.Suite
	rc          *rockset.RockClient
	ws          string
	collections []string
}

func TestCollectionIntegrationSuite(t *testing.T) {
	rc, randomName := vcrTestClient(t, t.Name())

	suite.Run(t, &CollectionTestSuite{rc: rc, ws: randomName("collection")})
}

func (s *CollectionTestSuite) SetupSuite() {
	ctx := test.Context()
	_, err := s.rc.CreateWorkspace(ctx, s.ws)
	s.Require().NoError(err)

	err = s.rc.WaitUntilWorkspaceAvailable(ctx, s.ws)
	s.Require().NoError(err)
}

func (s *CollectionTestSuite) TearDownSuite() {
	ctx := test.Context()
	var deleted []string

	for _, c := range s.collections {
		err := s.rc.DeleteCollection(ctx, s.ws, c)
		if s.NoError(err) {
			deleted = append(deleted, c)
		}
	}

	for _, c := range deleted {
		err := s.rc.WaitUntilCollectionGone(ctx, s.ws, c)
		s.NoError(err)
	}

	err := s.rc.DeleteWorkspace(ctx, s.ws)
	s.NoError(err)
}

func (s *CollectionTestSuite) BeforeTest(suiteName, testName string) {
	s.rc, _ = vcrTestClient(s.T(), fmt.Sprintf("%s/%s", suiteName, testName))
}

func (s *CollectionTestSuite) TestGetCollection() {
	ctx := test.Context()

	collection, err := s.rc.GetCollection(ctx, persistentWorkspace, persistentCollection)
	s.NoError(err)
	s.Equal(persistentCollection, collection.GetName())
}

func (s *CollectionTestSuite) TestListAllCollections() {
	ctx := test.Context()

	collections, err := s.rc.ListCollections(ctx)
	s.NoError(err)

	s.T().Logf("collections: %d", len(collections))
}

func (s *CollectionTestSuite) TestListCollectionsInWorkspace() {
	ctx := test.Context()

	collections, err := s.rc.ListCollections(ctx, option.WithWorkspace(persistentWorkspace))
	s.NoError(err)

	s.T().Logf("collections in %s: %d", persistentWorkspace, len(collections))
}

func (s *CollectionTestSuite) TestCreateSampleCitiesCollection() {
	ctx := test.Context()
	name := randomName("cities")

	_, err := s.rc.CreateCollection(ctx, s.ws, name,
		option.WithSampleDataset(dataset.Cities))
	s.Require().NoError(err)
	s.collections = append(s.collections, name)

	err = s.rc.WaitUntilCollectionHasDocuments(ctx, s.ws, name, int64(145_658))
	s.NoError(err)
}

func (s *CollectionTestSuite) TestCreateSampleMoviesCollection() {
	ctx := test.Context()
	name := randomName("movies")

	_, err := s.rc.CreateCollection(ctx, s.ws, name,
		option.WithStorageCompressionType(option.StorageCompressionLZ4),
		option.WithSampleDatasetPattern("movies/*"))
	s.Require().NoError(err)
	s.collections = append(s.collections, name)

	err = s.rc.WaitUntilCollectionHasDocuments(ctx, s.ws, name, int64(2_830))
	s.NoError(err)
}

func (s *CollectionTestSuite) TestUpdateCollection() {
	ctx := test.Context()
	name := randomName("update")

	_, err := s.rc.CreateCollection(ctx, s.ws, name,
		option.WithStorageCompressionType(option.StorageCompressionZSTD),
		option.WithCollectionDescription("initial"),
		option.WithIngestTransformation("SELECT * FROM _input"))
	s.Require().NoError(err)
	s.collections = append(s.collections, name)

	err = s.rc.WaitUntilCollectionReady(ctx, s.ws, name)
	s.Require().NoError(err)

	_, err = s.rc.UpdateCollection(ctx, s.ws, name,
		option.WithStorageCompressionType(option.StorageCompressionLZ4))
	s.Error(err)

	_, err = s.rc.UpdateCollection(ctx, s.ws, name,
		option.WithCollectionDescription("updated"),
		option.WithIngestTransformation("SELECT * FROM _input WHERE _input.foo != 'bar'"))
	s.NoError(err)
}
