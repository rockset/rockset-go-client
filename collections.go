package rockset

import (
	"context"
	"time"

	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

// GetCollection gets information about a collection.
func (rc *RockClient) GetCollection(ctx context.Context, workspace, name string) (openapi.Collection, error) {
	var err error
	var resp *openapi.GetCollectionResponse
	getReq := rc.CollectionsApi.GetCollection(ctx, workspace, name)

	err = rc.Retry(ctx, func() error {
		resp, _, err = getReq.Execute()

		return err
	})

	if err != nil {
		return openapi.Collection{}, err
	}

	return resp.GetData(), nil
}

// ListCollections lists all collections, or in a specific workspace is option.WithWorkspace() is used.
func (rc *RockClient) ListCollections(ctx context.Context,
	options ...option.ListCollectionOption) ([]openapi.Collection, error) {
	var err error

	opts := option.ListCollectionOptions{}

	for _, o := range options {
		o(&opts)
	}

	var resp *openapi.ListCollectionsResponse

	if opts.Workspace == nil {
		listReq := rc.CollectionsApi.ListCollections(ctx)
		err = rc.Retry(ctx, func() error {
			resp, _, err = listReq.Execute()

			return err
		})
	} else {
		listWsReq := rc.CollectionsApi.WorkspaceCollections(ctx, *opts.Workspace)

		err = rc.Retry(ctx, func() error {
			resp, _, err = listWsReq.Execute()

			return err
		})
	}

	if err != nil {
		return nil, err
	}

	return resp.GetData(), nil
}

// DeleteCollection deletes  collection.
func (rc *RockClient) DeleteCollection(ctx context.Context, workspace, name string) error {
	deleteReq := rc.CollectionsApi.DeleteCollection(ctx, workspace, name)

	err := rc.Retry(ctx, func() error {
		_, _, err := deleteReq.Execute()

		return err
	})

	return err
}

// CreateCollection is used to create a collection from one or more sources:
//  - DynamoDB (see CreateDynamoDBIntegration())
//  - GCS (see CreateGCSIntegration())
//  - Kafka (see CreateKafkaIntegration())
//  - Kinesis (see CreateKinesisIntegration())
//  - MongoDB (see CreateMongoDBIntegration())
//  - S3 (see CreateS3Integration())
//
// It uses exponential backoff in case the API call is rate-limted.
//
// To create a collection from multiple sources, use:
// 	 c, err := rc.CreateCollection(ctx, "commons", "example",
//	   option.WithCollectionDescription("created by go example code"),
//	   option.WithS3Source("s3-integration-name", "rockset-go-tests",
//	     option.WithCSVFormat(
//	       []string{"city", "country", "population", "visited"},
//	       []option.ColumnType{
//	         option.ColumnTypeString, option.ColumnTypeString, option.ColumnTypeInteger, option.ColumnTypeBool,
//	       },
//	       option.WithEncoding("UTF-8"),
//	       option.WithEscapeChar("\\"),
//	       option.WithQuoteChar(`"`),
//	       option.WithSeparator(","),
//	    ),
//	    option.WithS3Prefix("cities.csv"),
//	  ),
//    option.WithKafkaSource("kafka-integration-name", "topic", option.KafkaStartingOffsetEarliest, option.WithJSONFormat(),
//      option.WithKafkaSourceV3(),
//    ),
//    option.WithCollectionRetention(time.Hour),
//    option.WithInsertOnly(),
//	  option.WithFieldMappingQuery("SELECT * FROM _input"),
//  )
func (rc *RockClient) CreateCollection(ctx context.Context, workspace, name string,
	options ...option.CollectionOption) (openapi.Collection, error) {
	var err error
	var resp *openapi.CreateCollectionResponse

	request := openapi.CreateCollectionRequest{}
	request.Name = name

	for _, o := range options {
		o(&request)
	}

	createReq := rc.CollectionsApi.CreateCollection(ctx, workspace)

	err = rc.Retry(ctx, func() error {
		resp, _, err = createReq.Body(request).Execute()
		return err
	})
	if err != nil {
		return openapi.Collection{}, err
	}

	return resp.GetData(), nil
}

// CreateS3Collection creates an S3 collection from an existing S3 integration.
// Not specifying a format will default to JSON.
// Deprecated: use CreateCollection() with option.WithS3Source() instead.
func (rc *RockClient) CreateS3Collection(ctx context.Context,
	workspace, name, description, integration, bucket, pattern string,
	format option.Format, options ...option.CollectionOption) (openapi.Collection, error) {
	var s3opts []option.S3SourceOption
	if pattern != "" {
		s3opts = append(s3opts, option.WithS3Pattern(pattern))
	}

	opts := []option.CollectionOption{
		option.WithS3Source(integration, bucket, format, s3opts...),
	}

	if description != "" {
		opts = append(opts, option.WithCollectionDescription(description))
	}
	opts = append(opts, options...)

	return rc.CreateCollection(ctx, workspace, name, opts...)
}

func (rc *RockClient) CreateKinesisCollection(ctx context.Context,
	workspace, name, description, integration, region, stream string,
	format option.Format, options ...option.CollectionOption) (openapi.Collection, error) {
	var err error
	var resp *openapi.CreateCollectionResponse

	createReq := rc.CollectionsApi.CreateCollection(ctx, workspace)
	createParams := openapi.NewCreateCollectionRequest(name)
	createParams.Description = &description

	f := openapi.FormatParams{}
	format(&f)

	createParams.Sources = []openapi.Source{
		{
			IntegrationName: &integration,
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

	err = rc.Retry(ctx, func() error {
		resp, _, err = createReq.Body(*createParams).Execute()
		return err
	})

	if err != nil {
		return openapi.Collection{}, err
	}

	return resp.GetData(), nil
}

func (rc *RockClient) CreateGCSCollection(ctx context.Context,
	workspace, name, description, integration, bucket, prefix string,
	format option.Format, options ...option.CollectionOption) (openapi.Collection, error) {
	var err error
	var resp *openapi.CreateCollectionResponse

	createReq := rc.CollectionsApi.CreateCollection(ctx, workspace)
	createParams := openapi.NewCreateCollectionRequest(name)
	createParams.Description = &description

	f := openapi.FormatParams{}
	format(&f)

	createParams.Sources = []openapi.Source{
		{
			IntegrationName: &integration,
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

	err = rc.Retry(ctx, func() error {
		resp, _, err = createReq.Body(*createParams).Execute()
		return err
	})

	if err != nil {
		return openapi.Collection{}, err
	}

	return resp.GetData(), nil
}

func (rc *RockClient) CreateDynamoDBCollection(ctx context.Context,
	workspace, name, description, integration, region, tableName string, maxRCU int64,
	format option.Format, options ...option.CollectionOption) (openapi.Collection, error) {
	var err error
	var resp *openapi.CreateCollectionResponse

	createReq := rc.CollectionsApi.CreateCollection(ctx, workspace)
	createParams := openapi.NewCreateCollectionRequest(name)
	createParams.Description = &description

	f := openapi.FormatParams{}
	format(&f)

	createParams.Sources = []openapi.Source{
		{
			IntegrationName: &integration,
			Dynamodb: &openapi.SourceDynamoDb{
				AwsRegion: &region,
				TableName: tableName,
			},
			FormatParams: &f,
		},
	}

	for _, o := range options {
		o(createParams)
	}

	err = rc.Retry(ctx, func() error {
		resp, _, err = createReq.Body(*createParams).Execute()
		return err
	})

	if err != nil {
		return openapi.Collection{}, err
	}

	return resp.GetData(), nil
}

func (rc *RockClient) CreateFileUploadCollection(ctx context.Context,
	workspace, name, description, fileName string, fileSize int64,
	fileUploadTime time.Time,
	format option.Format, options ...option.CollectionOption) (openapi.Collection, error) {
	var err error
	var resp *openapi.CreateCollectionResponse

	createReq := rc.CollectionsApi.CreateCollection(ctx, workspace)
	createParams := openapi.NewCreateCollectionRequest(name)
	createParams.Description = &description

	f := openapi.FormatParams{}
	format(&f)

	createParams.Sources = []openapi.Source{
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
	err = rc.Retry(ctx, func() error {
		resp, _, err = createReq.Body(*createParams).Execute()
		return err
	})
	if err != nil {
		return openapi.Collection{}, err
	}

	return resp.GetData(), nil
}

// CreateKafkaCollection creates a single collection from a Kafka integration. Requires using
// option.WithKafkaSource() to configure the Kafka source options.
//
//	rc, err := rockset.NewClient()
//	if err != nil { ... }
//
//	c, err := rc.CreateKafkaCollection(ctx, "workspace", "collection",
//	    option.WithCollectionRetention(time.Hour),
//	    option.WithKafkaSource("integration-name", "topic", option.KafkaStartingOffsetEarliest,
//	        option.WithJSONFormat(),
//	    ))
//
//	if err != nil { ... }
//	if err = rc.WaitUntilCollectionReady(ctx, "workspace", "collection"); err != nil {
//	    ...
//	}
func (rc *RockClient) CreateKafkaCollection(ctx context.Context, workspace, name string,
	options ...option.CollectionOption) (openapi.Collection, error) {
	var err error
	var resp *openapi.CreateCollectionResponse

	createReq := rc.CollectionsApi.CreateCollection(ctx, workspace)
	createParams := openapi.NewCreateCollectionRequest(name)

	createParams.Sources = []openapi.Source{}

	for _, o := range options {
		o(createParams)
	}

	err = rc.Retry(ctx, func() error {
		resp, _, err = createReq.Body(*createParams).Execute()
		return err
	})

	if err != nil {
		return openapi.Collection{}, err
	}

	return resp.GetData(), nil
}

func (rc *RockClient) CreateMongoDBCollection(ctx context.Context,
	workspace, name, description, integration, database, collection string,
	format option.Format, options ...option.CollectionOption) (openapi.Collection, error) {
	var err error
	var resp *openapi.CreateCollectionResponse

	createReq := rc.CollectionsApi.CreateCollection(ctx, workspace)
	createParams := openapi.NewCreateCollectionRequest(name)
	createParams.Description = &description

	f := openapi.FormatParams{}
	format(&f)

	createParams.Sources = []openapi.Source{
		{
			IntegrationName: &integration,
			Mongodb: &openapi.SourceMongoDb{
				DatabaseName:   database,
				CollectionName: collection,
			},
			FormatParams: &f,
		},
	}

	for _, o := range options {
		o(createParams)
	}

	err = rc.Retry(ctx, func() error {
		resp, _, err = createReq.Body(*createParams).Execute()
		return err
	})
	if err != nil {
		return openapi.Collection{}, err
	}

	return resp.GetData(), nil
}
