package rockset

import (
	"context"
	"time"

	"github.com/rockset/rockset-go-client/openapi"
	"github.com/rockset/rockset-go-client/option"
)

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

func (rc *RockClient) DeleteCollection(ctx context.Context, workspace, name string) error {
	deleteReq := rc.CollectionsApi.DeleteCollection(ctx, workspace, name)

	err := rc.Retry(ctx, func() error {
		_, _, err := deleteReq.Execute()

		return err
	})

	return err
}

// CreateCollection is a convenience method to create a collection, which uses exponential backoff in case
// the API call is ratelimted.
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
