package rockset

import (
	"context"
	"net/http"

	"github.com/rs/zerolog"

	rockerr "github.com/rockset/rockset-go-client/errors"
	"github.com/rockset/rockset-go-client/openapi"
)

// GetOrganization gets the current organization.
func (rc *RockClient) GetOrganization(ctx context.Context) (openapi.Organization, error) {
	var err error
	var httpResp *http.Response
	var resp *openapi.OrganizationResponse
	log := zerolog.Ctx(ctx)
	getReq := rc.OrganizationsApi.GetOrganization(ctx)

	err = rc.Retry(ctx, func() error {
		resp, httpResp, err = getReq.Execute()

		return rockerr.NewWithStatusCode(err, httpResp)
	})

	if err != nil {
		return openapi.Organization{}, err
	}

	data := resp.GetData()
	log.Debug().Str("name", data.GetDisplayName()).Msg("organization")

	return data, nil
}
