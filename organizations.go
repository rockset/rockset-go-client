package rockset

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/rockset/rockset-go-client/openapi"
)

func (rc *RockClient) GetOrganization(ctx context.Context) (openapi.Organization, error) {
	log := zerolog.Ctx(ctx)
	var err error
	var resp openapi.OrganizationResponse
	getReq := rc.OrganizationsApi.GetOrganization(ctx)

	err = rc.Retry(ctx, func() (bool, error) {
		resp, _, err = getReq.Execute()
		if err != nil {
			re := NewError(err)
			if re.Retryable() {
				log.Debug().Str("error", re.GetMessage()).Msg("got retryable error")
				return true, nil
			}
			return false, re
		}

		return false, nil
	})

	if err != nil {
		return openapi.Organization{}, err
	}

	return *resp.Data, nil
}
