package rockset_test

import (
	"github.com/stretchr/testify/suite"
	"testing"

	"github.com/fatih/structs"
	"github.com/stretchr/testify/require"

	"github.com/rockset/rockset-go-client"
)

type DocumentIntegrationSuite struct {
	suite.Suite
	rc         *rockset.RockClient
	collection string
	id         string
}

func TestDocumentSuite(t *testing.T) {
	skipUnlessIntegrationTest(t)

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	s := DocumentIntegrationSuite{
		rc:         rc,
		collection: randomName(t, "docs"),
	}
	suite.Run(t, &s)
}

func (s *DocumentIntegrationSuite) SetupSuite() {
	ctx := testCtx()

	// create a temporary write-api collection
	_, err := s.rc.CreateCollection(ctx, persistentWorkspace, s.collection)
	s.Require().NoError(err)
	err = s.rc.WaitUntilCollectionReady(ctx, persistentWorkspace, s.collection)
	s.Require().NoError(err)
}

func (s *DocumentIntegrationSuite) TearDownSuite() {
	ctx := testCtx()

	err := s.rc.DeleteCollection(ctx, persistentWorkspace, s.collection)
	s.Require().NoError(err)
}

func (s *DocumentIntegrationSuite) TestAddDocument() {
	ctx := testCtx()

	type doc struct {
		Foo string `json:"foo"`
		Bar bool   `json:"bar"`
	}

	var docs = []interface{}{
		structs.Map(doc{Foo: "foo"}),
	}

	res, err := s.rc.AddDocuments(ctx, persistentWorkspace, "patch", docs)
	s.Require().NoError(err)
	s.Require().Len(res, 1)
	s.id = res[0].GetId()
}

func (s *DocumentIntegrationSuite) TestPatchDocument() {
	ctx := testCtx()

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
	res, err := s.rc.PatchDocuments(ctx, persistentWorkspace, "patch", patches)
	s.Require().NoError(err)
	s.Require().Len(res, 1)
	s.Assert().Equal("PATCHED", res[0].GetStatus())
}

// TestRemoveDocument should really be named TestDeleteDocument as the call is DeleteDocuments()
// but we want the operations in alphabetical order, add/patch/remove (delete).
func (s *DocumentIntegrationSuite) TestRemoveDocument() {
	ctx := testCtx()

	res, err := s.rc.DeleteDocuments(ctx, "tests", "patch", []string{s.id})
	s.Require().NoError(err)
	s.Require().Len(res, 1)
	s.Assert().Equal("DELETED", res[0].GetStatus())
}
