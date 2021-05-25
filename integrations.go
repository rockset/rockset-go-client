package rockset

import (
	"context"

	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

func (rc *RockClient) DeleteIntegration(ctx context.Context, name string) error {
	var err error
	req := rc.IntegrationsApi.DeleteIntegration(ctx, name)

	err = rc.Retry(ctx, func() error {
		_, _, err = req.Execute()
		return err
	})

	return err
}

func (rc *RockClient) CreateS3Integration(ctx context.Context, name string, creds option.AWSCredentialsFn,
	options ...option.S3IntegrationOption) (openapi.CreateIntegrationResponse, error) {
	var err error
	var resp openapi.CreateIntegrationResponse

	q := rc.IntegrationsApi.CreateIntegration(ctx)
	req := openapi.NewCreateIntegrationRequest(name)

	c := option.AWSCredentials{}
	creds(&c)

	opts := option.S3Integration{}
	for _, o := range options {
		o(&opts)
	}

	req.S3 = &openapi.S3Integration{}

	if opts.Description != nil {
		req.Description = opts.Description
	}
	if c.AwsRole != nil {
		req.S3.AwsRole = c.AwsRole
	}
	if c.AwsAccessKey != nil {
		req.S3.AwsAccessKey = c.AwsAccessKey
	}

	err = rc.Retry(ctx, func() error {
		resp, _, err = q.Body(*req).Execute()
		return err
	})

	if err != nil {
		return openapi.CreateIntegrationResponse{}, err
	}

	return resp, nil
}

func (rc *RockClient) CreateKinesisIntegration(ctx context.Context, name string, creds option.AWSCredentialsFn,
	options ...option.KinesisIntegrationOption) (openapi.CreateIntegrationResponse, error) {
	var err error
	var resp openapi.CreateIntegrationResponse
	q := rc.IntegrationsApi.CreateIntegration(ctx)
	req := openapi.NewCreateIntegrationRequest(name)

	c := option.AWSCredentials{}
	creds(&c)

	opts := option.KinesisIntegration{}
	for _, o := range options {
		o(&opts)
	}

	req.S3 = &openapi.S3Integration{}
	if opts.Description != nil {
		req.Description = opts.Description
	}
	if c.AwsRole != nil {
		req.S3.AwsRole = c.AwsRole
	}
	if c.AwsAccessKey != nil {
		req.S3.AwsAccessKey = c.AwsAccessKey
	}

	err = rc.Retry(ctx, func() error {
		resp, _, err = q.Body(*req).Execute()
		return err
	})

	if err != nil {
		return openapi.CreateIntegrationResponse{}, err
	}

	return resp, nil
}

func (rc *RockClient) CreateDynamoDBIntegration(ctx context.Context, name string, creds option.AWSCredentialsFn,
	options ...option.DynamoDBIntegrationOption) (openapi.CreateIntegrationResponse, error) {
	var err error
	var resp openapi.CreateIntegrationResponse

	q := rc.IntegrationsApi.CreateIntegration(ctx)
	req := openapi.NewCreateIntegrationRequest(name)

	c := option.AWSCredentials{}
	creds(&c)

	opts := option.DynamoDBIntegration{}
	for _, o := range options {
		o(&opts)
	}

	req.Dynamodb = &openapi.DynamodbIntegration{}
	if opts.Description != nil {
		req.Description = opts.Description
	}
	if c.AwsRole != nil {
		req.Dynamodb.AwsRole = c.AwsRole
	}
	if c.AwsAccessKey != nil {
		req.Dynamodb.AwsAccessKey = c.AwsAccessKey
	}

	err = rc.Retry(ctx, func() error {
		resp, _, err = q.Body(*req).Execute()
		return err
	})

	if err != nil {
		return openapi.CreateIntegrationResponse{}, err
	}

	return resp, nil
}

// TODO redshift

func (rc *RockClient) CreateGCSIntegration(ctx context.Context, name string, creds option.GCSCredentialsFn,
	options ...option.GCSIntegrationOption) (openapi.CreateIntegrationResponse, error) {
	var err error
	var resp openapi.CreateIntegrationResponse

	q := rc.IntegrationsApi.CreateIntegration(ctx)
	req := openapi.NewCreateIntegrationRequest(name)

	sa := openapi.GcsIntegration{}
	creds(&sa)

	opts := option.GCSIntegration{}
	for _, o := range options {
		o(&opts)
	}

	req.Gcs = &sa
	if opts.Description != nil {
		req.Description = opts.Description
	}

	err = rc.Retry(ctx, func() error {
		resp, _, err = q.Body(*req).Execute()
		return err
	})

	if err != nil {
		return openapi.CreateIntegrationResponse{}, err
	}

	return resp, nil
}

// TODO segment
// TODO kafka
// TODO mongodb
