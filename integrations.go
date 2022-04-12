package rockset

import (
	"context"

	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

func (rc *RockClient) GetIntegration(ctx context.Context, name string) (openapi.Integration, error) {
	var err error
	req := rc.IntegrationsApi.GetIntegration(ctx, name)

	var resp *openapi.GetIntegrationResponse
	err = rc.Retry(ctx, func() error {
		resp, _, err = req.Execute()
		return err
	})
	if err != nil {
		return openapi.Integration{}, err
	}

	return resp.GetData(), nil
}

func (rc *RockClient) ListIntegrations(ctx context.Context) ([]openapi.Integration, error) {
	var err error
	req := rc.IntegrationsApi.ListIntegrations(ctx)

	var resp *openapi.ListIntegrationsResponse
	err = rc.Retry(ctx, func() error {
		resp, _, err = req.Execute()
		return err
	})
	if err != nil {
		return nil, err
	}

	return resp.GetData(), nil
}

func (rc *RockClient) DeleteIntegration(ctx context.Context, name string) error {
	var err error
	req := rc.IntegrationsApi.DeleteIntegration(ctx, name)

	err = rc.Retry(ctx, func() error {
		_, _, err = req.Execute()
		return err
	})

	return err
}

func (rc *RockClient) CreateAzureBlobStorageIntegration(ctx context.Context, name string,
	connection string) (openapi.Integration, error) {
	var err error
	var resp *openapi.CreateIntegrationResponse
	q := rc.IntegrationsApi.CreateIntegration(ctx)
	req := openapi.NewCreateIntegrationRequest(name)

	req.AzureBlobStorage = &openapi.AzureBlobStorageIntegration{
		ConnectionString: connection,
	}

	err = rc.Retry(ctx, func() error {
		resp, _, err = q.Body(*req).Execute()
		return err
	})

	if err != nil {
		return openapi.Integration{}, err
	}

	return resp.GetData(), nil
}

func (rc *RockClient) CreateS3Integration(ctx context.Context, name string, creds option.AWSCredentialsFn,
	options ...option.S3IntegrationOption) (openapi.Integration, error) {
	var err error
	var resp *openapi.CreateIntegrationResponse

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
		return openapi.Integration{}, err
	}

	return resp.GetData(), nil
}

func (rc *RockClient) CreateKinesisIntegration(ctx context.Context, name string, creds option.AWSCredentialsFn,
	options ...option.KinesisIntegrationOption) (openapi.Integration, error) {
	var err error
	var resp *openapi.CreateIntegrationResponse
	q := rc.IntegrationsApi.CreateIntegration(ctx)
	req := openapi.NewCreateIntegrationRequest(name)

	c := option.AWSCredentials{}
	creds(&c)

	opts := option.KinesisIntegration{}
	for _, o := range options {
		o(&opts)
	}

	req.Kinesis = &openapi.KinesisIntegration{}
	if opts.Description != nil {
		req.Description = opts.Description
	}
	if c.AwsRole != nil {
		req.Kinesis.AwsRole = c.AwsRole
	}
	if c.AwsAccessKey != nil {
		req.Kinesis.AwsAccessKey = c.AwsAccessKey
	}

	err = rc.Retry(ctx, func() error {
		resp, _, err = q.Body(*req).Execute()
		return err
	})

	if err != nil {
		return openapi.Integration{}, err
	}

	return resp.GetData(), nil
}

// CreateDynamoDBIntegration creates a new AWS DynamoDB integration. It requires AWS credentials using either
// option.AWSKeys() or option.AWSRole(), and an S3 bucket which is used to export the DynamoDB tables.
func (rc *RockClient) CreateDynamoDBIntegration(ctx context.Context, name string, creds option.AWSCredentialsFn,
	s3BucketName string, options ...option.DynamoDBIntegrationOption) (openapi.Integration, error) {
	var err error
	var resp *openapi.CreateIntegrationResponse

	q := rc.IntegrationsApi.CreateIntegration(ctx)
	req := openapi.NewCreateIntegrationRequest(name)

	c := option.AWSCredentials{}
	creds(&c)

	opts := option.DynamoDBIntegration{}
	for _, o := range options {
		o(&opts)
	}

	req.Dynamodb = &openapi.DynamodbIntegration{
		S3ExportBucketName: &s3BucketName,
	}
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
		return openapi.Integration{}, err
	}

	return resp.GetData(), nil
}

func (rc *RockClient) CreateGCSIntegration(ctx context.Context, name, serviceAccountKeyFileJSON string,
	options ...option.GCSIntegrationOption) (openapi.Integration, error) {
	var err error
	var resp *openapi.CreateIntegrationResponse

	q := rc.IntegrationsApi.CreateIntegration(ctx)
	req := openapi.NewCreateIntegrationRequest(name)

	req.Gcs = &openapi.GcsIntegration{
		GcpServiceAccount: &openapi.GcpServiceAccount{
			ServiceAccountKeyFileJson: serviceAccountKeyFileJSON,
		},
	}

	opts := option.GCSIntegration{}
	for _, o := range options {
		o(&opts)
	}
	if opts.Description != nil {
		req.Description = opts.Description
	}

	err = rc.Retry(ctx, func() error {
		resp, _, err = q.Body(*req).Execute()
		return err
	})

	if err != nil {
		return openapi.Integration{}, err
	}

	return resp.GetData(), nil
}

func (rc *RockClient) CreateSegmentIntegration(ctx context.Context, name, connectionString string,
	options ...option.SegmentIntegrationOption) (openapi.Integration, error) {
	var err error
	var resp *openapi.CreateIntegrationResponse

	q := rc.IntegrationsApi.CreateIntegration(ctx)
	req := openapi.NewCreateIntegrationRequest(name)

	opts := option.SegmentIntegration{}
	for _, o := range options {
		o(&opts)
	}

	req.Segment.ConnectionString = &connectionString
	if opts.Description != nil {
		req.Description = opts.Description
	}

	err = rc.Retry(ctx, func() error {
		resp, _, err = q.Body(*req).Execute()
		return err
	})

	if err != nil {
		return openapi.Integration{}, err
	}

	return resp.GetData(), nil
}

func (rc *RockClient) CreateKafkaIntegration(ctx context.Context, name string, topics []string, format KafkaFormat,
	options ...option.KafkaIntegrationOption) (openapi.Integration, error) {
	var err error
	var resp *openapi.CreateIntegrationResponse

	q := rc.IntegrationsApi.CreateIntegration(ctx)
	req := openapi.NewCreateIntegrationRequest(name)

	opts := option.KafkaIntegration{}
	for _, o := range options {
		o(&opts)
	}

	f := format.String()
	req.Kafka = &openapi.KafkaIntegration{
		KafkaTopicNames: topics,
		KafkaDataFormat: &f,
	}
	if opts.Description != nil {
		req.Description = opts.Description
	}

	err = rc.Retry(ctx, func() error {
		resp, _, err = q.Body(*req).Execute()
		return err
	})

	if err != nil {
		return openapi.Integration{}, err
	}

	return resp.GetData(), nil
}

func (rc *RockClient) CreateMongoDBIntegration(ctx context.Context, name, connectionURI string,
	options ...option.MongoDBIntegrationOption) (openapi.Integration, error) {
	var err error
	var resp *openapi.CreateIntegrationResponse

	q := rc.IntegrationsApi.CreateIntegration(ctx)
	req := openapi.NewCreateIntegrationRequest(name)

	opts := option.MongoDBIntegration{}
	for _, o := range options {
		o(&opts)
	}

	req.Mongodb = &openapi.MongoDbIntegration{
		ConnectionUri: connectionURI,
	}
	if opts.Description != nil {
		req.Description = opts.Description
	}

	err = rc.Retry(ctx, func() error {
		resp, _, err = q.Body(*req).Execute()
		return err
	})

	if err != nil {
		return openapi.Integration{}, err
	}

	return resp.GetData(), nil
}
