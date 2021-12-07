package rockset

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/rockset/rockset-go-client/openapi"
)

// GetOrganization gets the current organization.
func (rc *RockClient) GetOrganization(ctx context.Context) (openapi.Organization, error) {
	var err error
	var resp openapi.OrganizationResponse
	log := zerolog.Ctx(ctx)
	getReq := rc.OrganizationsApi.GetOrganization(ctx)

	err = rc.Retry(ctx, func() error {
		resp, _, err = getReq.Execute()
		return err
	})

	if err != nil {
		return openapi.Organization{}, err
	}

	data := resp.GetData()
	log.Debug().Str("name", data.GetDisplayName()).Msg("organization")

	return data, nil
}
