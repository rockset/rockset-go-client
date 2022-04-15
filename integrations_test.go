package rockset_test

import (
	"context"
	"errors"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/option"
)

func (s *IntegrationsSuite) TestAzureBlob() {
	s.T().Skip("pending better abstraction for local testing")
}

func TestRockClient_S3Integration(t *testing.T) {
	skipUnlessIntegrationTest(t)

	ctx := testCtx()
	name := "s3test"

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	// get the integration
	getReq := rc.IntegrationsApi.GetIntegration(ctx, name)
	_, _, err = getReq.Execute()
	if err != nil {
		// check if it is missing
		var re rockset.Error
		if errors.As(err, &re) {
			if !re.IsNotFoundError() {
				require.NoError(t, err)
			}
		}
	} else {
		// the integration exists, delete it
		deleteReq := rc.IntegrationsApi.DeleteIntegration(ctx, name)
		_, _, err = deleteReq.Execute()
		require.NoError(t, err)
	}

	// create a new integration
	_, err = rc.CreateS3Integration(ctx, name,
		option.AWSRole("arn:aws:iam::469279130686:role/rockset-s3-integration"),
		option.WithS3IntegrationDescription("test"))
	require.NoError(t, err)

	// list integrations and look for the newly created integration
	listReq := rc.IntegrationsApi.ListIntegrations(ctx)
	listResp, _, err := listReq.Execute()
	require.NoError(t, err)

	var found bool
	for _, i := range listResp.GetData() {
		if i.Name == name {
			found = true
		}
	}
	if !found {
		t.Errorf("could not find %s", name)
	}

	// delete
	deleteReq := rc.IntegrationsApi.DeleteIntegration(ctx, name)
	_, _, err = deleteReq.Execute()
	require.NoError(t, err)
}

func TestRockClient_CreateGCSIntegration(t *testing.T) {
	skipUnlessIntegrationTest(t)

	saKeyFile := os.Getenv("GCP_SA_KEY_JSON")
	if saKeyFile == "" {
		t.Skip("environment variable GCP_SA_KEY_JSON set")
	}

	ctx := testCtx()
	rc, err := rockset.NewClient()
	require.NoError(t, err)

	iName := "testGCSIntegration"
	gcs, err := rc.GetIntegration(ctx, iName)
	if err == nil {
		err = rc.DeleteIntegration(ctx, iName)
		require.NoError(t, err)
		log.Printf("collection deleted: %s", iName)
	} else {
		var re rockset.Error
		if !errors.As(err, &re) {
			t.Fatalf("GetIntegration failed for %s: %v", iName, err)
		}
		if !re.IsNotFoundError() {
			t.Fatalf("GetIntegration failed for %s: %v", iName, re)
		}
	}

	gcs, err = rc.CreateGCSIntegration(ctx, iName, saKeyFile,
		option.WithGCSIntegrationDescription("created by rockset integration tests"))
	require.NoError(t, err)

	log.Printf("created gcs integration: %s", gcs.GetName())
}

type IntegrationsSuite struct {
	suite.Suite
}

func TestIntegrations(t *testing.T) {
	s := new(IntegrationsSuite)
	suite.Run(t, s)
}

type MockIntegrationsAPI struct {
	mock.Mock
}

func (m *MockIntegrationsAPI) CreateIntegration(ctx context.Context) {

}
