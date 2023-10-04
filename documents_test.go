package rockset_test

import (
	"testing"

	"github.com/fatih/structs"
	"github.com/stretchr/testify/suite"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/internal/test"
)

type DocumentIntegrationSuite struct {
	suite.Suite
	rc         *rockset.RockClient
	ws         string
	collection string
	id         string
}

func TestDocumentSuite(t *testing.T) {
	rc, randomName := vcrTestClient(t, t.Name())

	s := DocumentIntegrationSuite{
		rc:         rc,
		ws:         randomName("docs"),
		collection: randomName("docs"),
	}
	suite.Run(t, &s)
}

func (s *DocumentIntegrationSuite) SetupSuite() {
	ctx := test.Context()

	_, err := s.rc.CreateWorkspace(ctx, s.ws)
	s.Require().NoError(err)

	// create a temporary write-api collection
	_, err = s.rc.CreateCollection(ctx, s.ws, s.collection)
	s.Require().NoError(err)
	err = s.rc.Wait.UntilCollectionReady(ctx, s.ws, s.collection)
	s.Require().NoError(err)
}

func (s *DocumentIntegrationSuite) TearDownSuite() {
	ctx := test.Context()

	err := s.rc.DeleteCollection(ctx, s.ws, s.collection)
	s.Require().NoError(err)

	err = s.rc.Wait.UntilCollectionGone(ctx, s.ws, s.collection)
	s.Require().NoError(err)

	err = s.rc.DeleteWorkspace(ctx, s.ws)
	s.Require().NoError(err)
}

func (s *DocumentIntegrationSuite) TestAddDocument() {
	ctx := test.Context()
	rc, _ := vcrTestClient(s.T(), s.T().Name())

	type doc struct {
		Foo string `json:"foo"`
		Bar bool   `json:"bar"`
	}

	var docs = []interface{}{
		structs.Map(doc{Foo: "foo"}),
	}

	res, err := rc.AddDocuments(ctx, s.ws, s.collection, docs)
	s.Require().NoError(err)
	s.Require().Len(res, 1)
	s.id = res[0].GetId()
}

func (s *DocumentIntegrationSuite) TestPatchDocument() {
	ctx := test.Context()
	rc, _ := vcrTestClient(s.T(), s.T().Name())
	s.T().Logf("t: %s", s.T().Name())

	patches := []rockset.PatchDocument{
		{
			ID: s.id,
			Patches: []rockset.PatchOperation{
				{
					Op:    "replace",
					Path:  "/Bar",
					Value: true,
				},
			},
		},
	}
	res, err := rc.PatchDocuments(ctx, s.ws, s.collection, patches)
	s.Require().NoError(err)
	s.Require().Len(res, 1)
	s.Assert().Equal("PATCHED", res[0].GetStatus())
}

// TestRemoveDocument should really be named TestDeleteDocument as the call is DeleteDocuments()
// but we want the operations in alphabetical order, add/patch/remove (delete).
func (s *DocumentIntegrationSuite) TestRemoveDocument() {
	ctx := test.Context()
	rc, _ := vcrTestClient(s.T(), s.T().Name())

	res, err := rc.DeleteDocuments(ctx, s.ws, s.collection, []string{s.id})
	s.Require().NoError(err)
	s.Require().Len(res, 1)
	s.Assert().Equal("DELETED", res[0].GetStatus())
}
