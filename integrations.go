package rockset

import (
	"context"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

func (rc *RockClient) DeleteIntegration(ctx context.Context, name string) error {
	req := rc.IntegrationsApi.DeleteIntegration(ctx, name)
	_, _, err := req.Execute()
	return err
}

func (rc *RockClient) CreateS3Integration(ctx context.Context, name string, creds option.AWSCredentials,
	options ...option.S3IntegrationOption) (openapi.CreateIntegrationResponse, error) {
	q := rc.IntegrationsApi.CreateIntegration(ctx)
	req := openapi.NewCreateIntegrationRequest(name)

	opts := option.S3Integration{}
	creds(&opts)
	for _, o := range options {
		o(&opts)
	}

	req.S3 = &openapi.S3Integration{
		AwsAccessKey: nil,
		AwsRole:      nil,
	}
	if opts.Description != nil {
		req.Description = opts.Description
	}
	if opts.AwsRole != nil {
		req.S3.AwsRole = opts.AwsRole
	}
	if opts.AwsAccessKey != nil {
		req.S3.AwsAccessKey = opts.AwsAccessKey
	}

	resp, _, err := q.Body(*req).Execute()
	if err != nil {
		return openapi.CreateIntegrationResponse{}, err
	}

	return resp, nil
}
