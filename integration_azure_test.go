package rockset_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/rockset/rockset-go-client"
)

func (s *IntegrationsSuite) TestAzureBlob() {
	connectionString := skipUnlessEnvSet(s.T(), "ROCKSET_AZURE_CONNECTION_STRING")

	name := "azuretest"
	ctx := testCtx()

	rc, err := rockset.NewClient()
	s.Require().NoError(err)

	// get the integration
	_, err = rc.GetIntegration(ctx, name)

	if err != nil {
		// check if it is missing
		var re rockset.Error
		if errors.As(err, &re) {
			if !re.IsNotFoundError() {
				s.Require().NoError(err)
			}
		}
	} else {
		// the integration exists, delete it
		deleteReq := rc.IntegrationsApi.DeleteIntegration(ctx, name)
		_, _, err = deleteReq.Execute()
		s.Require().NoError(err)
	}

	// create a new integration
	_, err = rc.CreateAzureBlobStorageIntegration(ctx, name, connectionString)
	s.Require().NoError(err)

	// list integrations and look for the newly created integration
	integration, err := rc.GetIntegration(ctx, name)
	s.Require().NoError(err)
	s.NotNil(integration)
}

type IntegrationsSuite struct {
	suite.Suite
}

func TestAzureIntegrations(t *testing.T) {
	s := new(IntegrationsSuite)
	suite.Run(t, s)
}
