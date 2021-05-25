package rockset

import (
	"context"
	"time"

	"github.com/rs/zerolog"

	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

func (rc *RockClient) GetCollection(ctx context.Context, workspace, name string) (openapi.Collection, error) {
	var err error
	var resp openapi.GetCollectionResponse
	getReq := rc.CollectionsApi.GetCollection(ctx, workspace, name)

	err = rc.Retry(ctx, func() error {
		resp, _, err = getReq.Execute()

		return err
	})

	if err != nil {
		return openapi.Collection{}, err
	}

	return *resp.Data, nil
}

func (rc *RockClient) DeleteCollection(ctx context.Context, workspace, name string) error {
	deleteReq := rc.CollectionsApi.DeleteCollection(ctx, workspace, name)

	err := rc.Retry(ctx, func() error {
		_, _, err := deleteReq.Execute()

		return err
	})

	return err
}

// CreateCollection is a convenience method to create a collection, which uses exponential backoff in case
// the API call is ratelimted. It will overwite the request.Name field with the argument name.
func (rc *RockClient) CreateCollection(ctx context.Context, workspace, name string,
	request *openapi.CreateCollectionRequest) (openapi.Collection, error) {
	log := zerolog.Ctx(ctx)

	createReq := rc.CollectionsApi.CreateCollection(ctx, workspace)
	if request.Name != name {
		log.Warn().Str("request.Name", request.Name).Str("name", name).
			Msg("name differs from request, using argument")
	}

	request.Name = name

	var err error
	var createResp openapi.CreateCollectionResponse
	err = rc.Retry(ctx, func() error {
		createResp, _, err = createReq.Body(*request).Execute()
		return err
	})
	if err != nil {
		return openapi.Collection{}, err
	}

	return *createResp.Data, nil
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

	var err error
	var createResp openapi.CreateCollectionResponse
	err = rc.Retry(ctx, func() error {
		createResp, _, err = createReq.Body(*createParams).Execute()
		return err
	})
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

	// TODO retry
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

	// TODO retry
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
				IncrementalField: nil, // TODO this is an optional string field
			},
			FormatParams: &f,
		},
	}

	for _, o := range options {
		o(createParams)
	}

	// TODO retry
	createResp, _, err := createReq.Body(*createParams).Execute()
	if err != nil {
		return openapi.Collection{}, err
	}

	return *createResp.Data, nil
}

func (rc *RockClient) CreateDynamoDBCollection(ctx context.Context,
	workspace, name, description, integration, region, tableName string, maxRCU int64,
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
				Rcu:       &maxRCU, // TODO optional long int field "Max RCU usage for scan"
			},
			FormatParams: &f,
		},
	}

	for _, o := range options {
		o(createParams)
	}

	// TODO retry
	createResp, _, err := createReq.Body(*createParams).Execute()
	if err != nil {
		return openapi.Collection{}, err
	}

	return *createResp.Data, nil
}

func (rc *RockClient) CreateFileUploadCollection(ctx context.Context,
	workspace, name, description, fileName string, fileSize int64,
	fileUploadTime time.Time,
	format Format, options ...option.CollectionOption) (openapi.Collection, error) {
	createReq := rc.CollectionsApi.CreateCollection(ctx, workspace)
	createParams := openapi.NewCreateCollectionRequest(name)
	createParams.Description = &description

	f := openapi.FormatParams{}
	format(&f)

	createParams.Sources = &[]openapi.Source{
		{
			FileUpload: &openapi.SourceFileUpload{
				// TODO how do you send the actual file contents?!?!
				FileName:       fileName,
				FileSize:       fileSize,
				FileUploadTime: fileUploadTime.Format(time.RFC3339),
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

	// TODO retry
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

	// TODO retry
	createResp, _, err := createReq.Body(*createParams).Execute()
	if err != nil {
		return openapi.Collection{}, err
	}

	return *createResp.Data, nil
}
