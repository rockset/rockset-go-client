package rockset_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/rockset/rockset-go-client"
)

func (s *IntegrationsSuite) TestAzureBlob() {
	// Need to update skip, this doesn't rely on AWS variables
	skipUnlessIntegrationTest(s.T())

	name := "azuretest"
	connectionString := "DefaultEndpointsProtocol=https;AccountName=rockset;AccountKey=<your-account-key>"

	ctx := testCtx()

	rc, err := rockset.NewClient()
	s.Require().NoError(err)

	// get the integration
	getReq := rc.IntegrationsApi.GetIntegration(ctx, name)
	_, _, err = getReq.Execute()

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
	listReq := rc.IntegrationsApi.ListIntegrations(ctx)
	listResp, _, err := listReq.Execute()
	s.Require().NoError(err)

	var found bool
	for _, i := range listResp.GetData() {
		if i.Name == name {
			found = true
		}
	}
	if !found {
		s.Errorf(errors.New("azure integration not found"), "%s", name)
	}
}

type IntegrationsSuite struct {
	suite.Suite
}

func TestAzureIntegrations(t *testing.T) {
	s := new(IntegrationsSuite)
	suite.Run(t, s)
}
