package rockset_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/rockset/rockset-go-client"
)

type AzureIntegrationsSuite struct {
	suite.Suite
	rc               *rockset.RockClient
	integrationName  string
	connectionString string
}

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
	t.Parallel()

	s := AzureIntegrationsSuite{
		rc:               testClient(t),
		integrationName:  randomName("azure"),
		connectionString: connectionString,
	}
	suite.Run(t, &s)
}
