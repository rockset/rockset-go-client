package rockset

import (
	"context"
	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

func (rc *RockClient) GetCollection(ctx context.Context, workspace, name string) (openapi.Collection, error) {
	getReq := rc.CollectionsApi.GetCollection(ctx, workspace, name)
	resp, _, err := getReq.Execute()
	if err != nil {
		return openapi.Collection{}, err
	}

	return *resp.Data, nil
}

func (rc *RockClient) DeleteCollection(ctx context.Context, workspace, name string) error {
	deleteReq := rc.CollectionsApi.DeleteCollection(ctx, workspace, name)
	_, _, err := deleteReq.Execute()
	return err
}

// CreateS3Collection creates an S3 collection from an existing S3 integration.
// Not specifying a format will default to JSON.
func (rc *RockClient) CreateS3Collection(ctx context.Context,
	workspace, name, description, integration, bucket, pattern string,
	format Format, options ...option.CollectionOption) (openapi.Collection, error) {
	createReq := rc.CollectionsApi.CreateCollection(ctx, workspace)
	createParams := openapi.NewCreateCollectionRequest(name)
	createParams.Description = &description

	f := openapi.FormatParams{}
	format(&f)

	createParams.Sources = &[]openapi.Source{
		{
			IntegrationName: integration,
			S3: &openapi.SourceS3{
				Pattern:  &pattern,
				Bucket:   bucket,
				Prefix:   nil,
				Prefixes: nil,
			},
			FormatParams: &f,
		},
	}

	for _, o := range options {
		o(createParams)
	}

	createResp, _, err := createReq.Body(*createParams).Execute()
	if err != nil {
		return openapi.Collection{}, err
	}

	return *createResp.Data, nil
}

func (rc *RockClient) CreateKinesisCollection(ctx context.Context,
	workspace, name, description, integration, region, stream string,
	format Format, options ...option.CollectionOption) (openapi.Collection, error) {
	createReq := rc.CollectionsApi.CreateCollection(ctx, workspace)
	createParams := openapi.NewCreateCollectionRequest(name)
	createParams.Description = &description

	f := openapi.FormatParams{}
	format(&f)

	createParams.Sources = &[]openapi.Source{
		{
			IntegrationName: integration,
			Kinesis: &openapi.SourceKinesis{
				StreamName: stream,
				AwsRegion:  &region,
			},
			FormatParams: &f,
		},
	}

	for _, o := range options {
		o(createParams)
	}

	createResp, _, err := createReq.Body(*createParams).Execute()
	if err != nil {
		return openapi.Collection{}, err
	}

	return *createResp.Data, nil
}

func (rc *RockClient) CreateGCSCollection(ctx context.Context,
	workspace, name, description, integration, bucket, prefix string,
	format Format, options ...option.CollectionOption) (openapi.Collection, error) {
	createReq := rc.CollectionsApi.CreateCollection(ctx, workspace)
	createParams := openapi.NewCreateCollectionRequest(name)
	createParams.Description = &description

	f := openapi.FormatParams{}
	format(&f)

	createParams.Sources = &[]openapi.Source{
		{
			IntegrationName: integration,
			Gcs: &openapi.SourceGcs{
				Bucket: &bucket,
				Prefix: &prefix,
			},
			FormatParams: &f,
		},
	}

	for _, o := range options {
		o(createParams)
	}

	createResp, _, err := createReq.Body(*createParams).Execute()
	if err != nil {
		return openapi.Collection{}, err
	}

	return *createResp.Data, nil
}

func (rc *RockClient) CreateRedshiftCollection(ctx context.Context,
	workspace, name, description, integration, database, schema, tableName string,
	format Format, options ...option.CollectionOption) (openapi.Collection, error) {
	createReq := rc.CollectionsApi.CreateCollection(ctx, workspace)
	createParams := openapi.NewCreateCollectionRequest(name)
	createParams.Description = &description

	f := openapi.FormatParams{}
	format(&f)

	createParams.Sources = &[]openapi.Source{
		{
			IntegrationName: integration,
			Redshift: &openapi.SourceRedshift{
				Database:         database,
				Schema:           schema,
				TableName:        tableName,
				IncrementalField: nil, // TODO
			},
			FormatParams: &f,
		},
	}

	for _, o := range options {
		o(createParams)
	}

	createResp, _, err := createReq.Body(*createParams).Execute()
	if err != nil {
		return openapi.Collection{}, err
	}

	return *createResp.Data, nil
}

func (rc *RockClient) CreateDynamoDBCollection(ctx context.Context,
	workspace, name, description, integration, region, tableName string,
	format Format, options ...option.CollectionOption) (openapi.Collection, error) {
	createReq := rc.CollectionsApi.CreateCollection(ctx, workspace)
	createParams := openapi.NewCreateCollectionRequest(name)
	createParams.Description = &description

	f := openapi.FormatParams{}
	format(&f)

	createParams.Sources = &[]openapi.Source{
		{
			IntegrationName: integration,
			Dynamodb: &openapi.SourceDynamoDb{
				AwsRegion: &region,
				TableName: tableName,
				Status:    nil, // TODO
				Rcu:       nil, // TODO
			},
			FormatParams: &f,
		},
	}

	for _, o := range options {
		o(createParams)
	}

	createResp, _, err := createReq.Body(*createParams).Execute()
	if err != nil {
		return openapi.Collection{}, err
	}

	return *createResp.Data, nil
}

func (rc *RockClient) CreateFileUploadCollection(ctx context.Context,
	workspace, name, description, integration, fileName string,
	format Format, options ...option.CollectionOption) (openapi.Collection, error) {
	createReq := rc.CollectionsApi.CreateCollection(ctx, workspace)
	createParams := openapi.NewCreateCollectionRequest(name)
	createParams.Description = &description

	f := openapi.FormatParams{}
	format(&f)

	createParams.Sources = &[]openapi.Source{
		{
			IntegrationName: integration,
			FileUpload: &openapi.SourceFileUpload{
				FileName:       fileName,
				FileSize:       0,  // TODO
				FileUploadTime: "", // TODO
			},
			FormatParams: &f,
		},
	}

	for _, o := range options {
		o(createParams)
	}

	createResp, _, err := createReq.Body(*createParams).Execute()
	if err != nil {
		return openapi.Collection{}, err
	}

	return *createResp.Data, nil
}

func (rc *RockClient) CreateKafkaCollection(ctx context.Context,
	workspace, name, description, integration, topic string,
	format Format, options ...option.CollectionOption) (openapi.Collection, error) {
	createReq := rc.CollectionsApi.CreateCollection(ctx, workspace)
	createParams := openapi.NewCreateCollectionRequest(name)
	createParams.Description = &description

	f := openapi.FormatParams{}
	format(&f)

	createParams.Sources = &[]openapi.Source{
		{
			IntegrationName: integration,
			Kafka: &openapi.SourceKafka{
				KafkaTopicName: topic,
				Status:         nil, // TODO
			},
			FormatParams: &f,
		},
	}

	for _, o := range options {
		o(createParams)
	}

	createResp, _, err := createReq.Body(*createParams).Execute()
	if err != nil {
		return openapi.Collection{}, err
	}

	return *createResp.Data, nil
}

func (rc *RockClient) CreateMongoDBCollection(ctx context.Context,
	workspace, name, description, integration, database, collection string,
	format Format, options ...option.CollectionOption) (openapi.Collection, error) {
	createReq := rc.CollectionsApi.CreateCollection(ctx, workspace)
	createParams := openapi.NewCreateCollectionRequest(name)
	createParams.Description = &description

	f := openapi.FormatParams{}
	format(&f)

	createParams.Sources = &[]openapi.Source{
		{
			IntegrationName: integration,
			Mongodb: &openapi.SourceMongoDb{
				DatabaseName:   database,
				CollectionName: collection,
				Status:         nil, // TODO
			},
			FormatParams: &f,
		},
	}

	for _, o := range options {
		o(createParams)
	}

	createResp, _, err := createReq.Body(*createParams).Execute()
	if err != nil {
		return openapi.Collection{}, err
	}

	return *createResp.Data, nil
}
