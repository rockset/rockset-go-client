package rockset_test

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/rockset/rockset-go-client"
)

func (s *AzureIntegrationsSuite) TestAzureBlob() {
	ctx := testCtx()

	_, err := s.rc.CreateAzureBlobStorageIntegration(ctx, s.integrationName, s.connectionString)
	s.Require().NoError(err)

	integration, err := s.rc.GetIntegration(ctx, s.integrationName)
	s.Require().NoError(err)
	s.NotNil(integration)

	err = s.rc.DeleteIntegration(ctx, s.integrationName)
	s.Require().NoError(err)
}

type AzureIntegrationsSuite struct {
	suite.Suite
	rc               *rockset.RockClient
	integrationName  string
	connectionString string
}

func (s *AzureIntegrationsSuite) TearDownSuite() {
	ctx := testCtx()

	// cleanup the integration if it is still around
	_, err := s.rc.GetIntegration(ctx, s.integrationName)
	if err == nil {
		s.T().Logf("deleting lingering azure integration: %s", s.integrationName)
		err = s.rc.DeleteIntegration(ctx, s.integrationName)
		s.Require().NoError(err)
	}
}

func TestAzureIntegrations(t *testing.T) {
	connectionString := skipUnlessEnvSet(t, "AZURE_CONNECTION_STRING")

	rc, err := rockset.NewClient()
	require.NoError(t, err)

	s := AzureIntegrationsSuite{
		rc:               rc,
		integrationName:  randomName(t, "azure"),
		connectionString: connectionString,
	}
	suite.Run(t, &s)
}
