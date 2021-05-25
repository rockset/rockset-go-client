package rockset

import (
	"context"

	"github.com/rockset/rockset-go-client/openapi"
)

func (rc *RockClient) GetOrganization(ctx context.Context) (openapi.Organization, error) {
	var err error
	var resp openapi.OrganizationResponse
	getReq := rc.OrganizationsApi.GetOrganization(ctx)

	err = rc.Retry(ctx, func() error {
		resp, _, err = getReq.Execute()
		return err
	})

	if err != nil {
		return openapi.Organization{}, err
	}

	return *resp.Data, nil
}
