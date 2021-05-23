package rockset_test

import (
	"testing"

	"github.com/rockset/rockset-go-client"
	"github.com/rockset/rockset-go-client/option"

	"github.com/stretchr/testify/require"
)

func TestRockClient_S3Integration(t *testing.T) {
	ctx := testCtx()
	name := "s3test"

	rc, err := rockset.NewClient(rockset.FromEnv())
	require.NoError(t, err)

	// get the integration
	getReq := rc.IntegrationsApi.GetIntegration(ctx, name)
	_, _, err = getReq.Execute()
	if err != nil {
		// check if it is missing
		var re rockset.Error
		if rockset.AsError(err, &re) {
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
	for _, i := range *listResp.Data {
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
