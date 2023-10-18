package option

import (
	"time"

	"github.com/rockset/rockset-go-client/dataset"
	"github.com/rockset/rockset-go-client/openapi"
)

type CollectionStatus string

func (c CollectionStatus) String() string { return string(c) }

const (
	CollectionStatusCreated     CollectionStatus = "CREATED"
	CollectionStatusInitialized CollectionStatus = "INITIALIZED"
	CollectionStatusReady       CollectionStatus = "READY"
)

type ListCollectionOptions struct {
	Workspace *string
}

type ListCollectionOption func(o *ListCollectionOptions)

func WithWorkspace(name string) func(o *ListCollectionOptions) {
	return func(o *ListCollectionOptions) {
		o.Workspace = &name
	}
}

// WithCollectionRequest is used to pass a openapi.CreateCollectionRequest to a CreateCollection() call.
func WithCollectionRequest(request openapi.CreateCollectionRequest) CollectionOption {
	return func(o *openapi.CreateCollectionRequest) {
		// TODO this is not a good way of copying the request as if we add a new field
		//   to the request, it won't get copied here
		o.Name = request.Name
		o.Description = request.Description
		o.Sources = request.Sources
		o.RetentionSecs = request.RetentionSecs
		o.EventTimeInfo = request.EventTimeInfo
		o.FieldMappingQuery = request.FieldMappingQuery
		o.ClusteringKey = request.ClusteringKey
	}
}

type S3SourceOption func(o *openapi.SourceS3)

func WithS3Region(region string) S3SourceOption {
	return func(o *openapi.SourceS3) {
		o.Region = &region
	}
}

// WithS3Prefix is used to pick the S3 prefix to ingest objects from.
// Can't be used together with WithS3Pattern()
func WithS3Prefix(prefix string) S3SourceOption {
	return func(o *openapi.SourceS3) {
		o.Prefix = &prefix
	}
}

// WithS3Pattern is used to pick the S3 pattern to ingest objects from.
// Can't be used together with WithS3Prefix()
func WithS3Pattern(pattern string) S3SourceOption {
	return func(o *openapi.SourceS3) {
		o.Pattern = &pattern
	}
}

func WithS3Source(IntegrationName, bucket string, fmt Format, options ...S3SourceOption) CollectionOption {
	src := openapi.SourceS3{
		Bucket: bucket,
	}
	for _, o := range options {
		o(&src)
	}

	fp := openapi.FormatParams{}
	fmt(&fp)

	return func(o *openapi.CreateCollectionRequest) {
		o.Sources = append(o.Sources, openapi.Source{
			IntegrationName: &IntegrationName,
			S3:              &src,
			FormatParams:    &fp,
		})
	}
}

type KafkaSourceOption func(o *openapi.SourceKafka)

type KafkaStartingOffset string

const (
	KafkaStartingOffsetEarliest KafkaStartingOffset = "EARLIEST"
	KafkaStartingOffsetLatest   KafkaStartingOffset = "LATEST"
)

func WithKafkaSourceV3() KafkaSourceOption {
	return func(o *openapi.SourceKafka) {
		o.UseV3 = openapi.PtrBool(true)
	}
}

func WithKafkaConsumerGroupID(id string) KafkaSourceOption {
	return func(o *openapi.SourceKafka) {
		o.ConsumerGroupId = &id
	}
}

type CollectionOption func(o *openapi.CreateCollectionRequest)

// WithCollectionDescription sets the description for the collection.
func WithCollectionDescription(d string) CollectionOption {
	return func(o *openapi.CreateCollectionRequest) {
		o.Description = &d
	}
}

func WithKafkaSource(IntegrationName, topic string, startingOffset KafkaStartingOffset, fmt Format,
	options ...KafkaSourceOption) CollectionOption {
	src := openapi.SourceKafka{
		KafkaTopicName:    &topic,
		OffsetResetPolicy: openapi.PtrString(string(startingOffset)),
	}
	for _, o := range options {
		o(&src)
	}

	fp := openapi.FormatParams{}
	fmt(&fp)

	return func(o *openapi.CreateCollectionRequest) {
		o.Sources = append(o.Sources, openapi.Source{
			IntegrationName: &IntegrationName,
			Kafka:           &src,
			FormatParams:    &fp,
		})
	}
}

// WithCollectionRetention sets the retention in seconds for documents.
func WithCollectionRetention(d time.Duration) CollectionOption {
	return func(o *openapi.CreateCollectionRequest) {
		s := int64(d.Seconds())
		o.RetentionSecs = &s
	}
}

// WithCollectionRetentionSeconds sets the retention in seconds for documents.
func WithCollectionRetentionSeconds(s int64) CollectionOption {
	return func(o *openapi.CreateCollectionRequest) {
		o.RetentionSecs = &s
	}
}

// WithDynamoDBMaxRCU sets the max RCU for a DynamoDB collection.
func WithDynamoDBMaxRCU(maxRCU int64) CollectionOption {
	return func(o *openapi.CreateCollectionRequest) {
		if o.Sources == nil || len(o.Sources) != 1 {
			return
		}
		if (o.Sources)[0].Dynamodb == nil {
			return
		}
		(o.Sources)[0].Dynamodb.Rcu = &maxRCU
	}
}

// WithFieldMappingQuery sets the field mapping query.
// Deprecated: use WithIngestTransformation instead.
func WithFieldMappingQuery(sql string) CollectionOption {
	return func(o *openapi.CreateCollectionRequest) {
		o.FieldMappingQuery = &openapi.FieldMappingQuery{
			Sql: &sql,
		}
	}
}

// WithIngestTransformation lets you run a [SQL query] over the data from your source and only persists the output
// of that query to the collection. This gives you the power of SQL to drop, rename, or combine fields,
// filter out incoming rows, and even aggregate incoming documents in real-time with rollups.
//
// [SQL query]: https://rockset.com/docs/ingest-transformation/
func WithIngestTransformation(sql string) CollectionOption {
	return func(o *openapi.CreateCollectionRequest) {
		o.FieldMappingQuery = &openapi.FieldMappingQuery{
			Sql: &sql,
		}
	}
}

func WithEventTimeInfo(field, timeZone string, format EventTimeInfoFormat) CollectionOption {
	return func(o *openapi.CreateCollectionRequest) {
		o.EventTimeInfo = &openapi.EventTimeInfo{
			Field:    field,
			Format:   openapi.PtrString(string(format)),
			TimeZone: &timeZone,
		}
	}
}

type EventTimeInfoFormat string

const (
	MillisecondsSinceEpoch EventTimeInfoFormat = "milliseconds_since_epoch"
	SecondsSinceEpoch      EventTimeInfoFormat = "seconds_since_epoch"
)

// WithCollectionClusteringKey adds a clustering key. Can be specified multiple times.
// Deprecated: use WithFieldMappingQuery() instead and with a CLUSTER BY clause.
func WithCollectionClusteringKey(fieldName, fieldType string, keys []string) CollectionOption {
	return func(o *openapi.CreateCollectionRequest) {
		if o.ClusteringKey == nil {
			o.ClusteringKey = []openapi.FieldPartition{}
		}

		o.ClusteringKey = append(o.ClusteringKey, openapi.FieldPartition{
			FieldName: &fieldName,
			Type:      &fieldType,
			Keys:      keys,
		})
	}
}

type FieldMissingAction string

func (f FieldMissingAction) String() string { return string(f) }

const (
	FieldMissingSkip FieldMissingAction = "SKIP"
	FieldMissingPass FieldMissingAction = "PASS"
)

type InputFieldFn func(field *openapi.InputField)

func InputField(fieldName string, ifMissing FieldMissingAction, drop bool, parameterName string) InputFieldFn {
	missing := ifMissing.String()
	return func(field *openapi.InputField) {
		field.FieldName = &fieldName
		field.IfMissing = &missing
		field.IsDrop = &drop
		field.Param = &parameterName
	}
}

type OnError string

func (e OnError) String() string { return string(e) }

const (
	OnErrorSkip OnError = "SKIP"
	OnErrorFail OnError = "FAIL"
)

type OutputFieldFn func(field *openapi.OutputField)

func OutputField(fieldName string, sql string, onError OnError) OutputFieldFn {
	return func(field *openapi.OutputField) {
		e := onError.String()
		field.FieldName = &fieldName
		field.Value = &openapi.SqlExpression{Sql: &sql}
		field.OnError = &e
	}
}

// WithCollectionFieldMapping adds a field mapping to the collection.
// If dropAll is true, the input and output fields are not set.
func WithCollectionFieldMapping(name string, dropAll bool, outputField OutputFieldFn,
	inputFields ...InputFieldFn) CollectionOption {
	return func(o *openapi.CreateCollectionRequest) {
		mapping := openapi.FieldMappingV2{
			Name: &name,
		}

		if dropAll {
			mapping.IsDropAllFields = &dropAll
		} else {
			out := openapi.OutputField{}
			outputField(&out)
			mapping.OutputField = &out

			inputs := make([]openapi.InputField, len(inputFields))
			for i, fn := range inputFields {
				var in openapi.InputField
				fn(&in)
				inputs[i] = in
			}
			mapping.InputFields = inputs
		}
	}
}

// WithSampleDataset creates a new collection from the Rockset public datasets.
func WithSampleDataset(ds dataset.Sample) CollectionOption {
	pattern := dataset.Lookup(ds)
	return func(o *openapi.CreateCollectionRequest) {
		o.Sources = append(o.Sources, openapi.Source{
			S3: &openapi.SourceS3{
				Bucket:  dataset.RocksetPublicDatasets,
				Pattern: &pattern,
			},
		})
	}
}

// WithSampleDatasetPattern creates a new collection from the Rockset public datasets,
// using pattern to select which file(s) to include in the collection.
func WithSampleDatasetPattern(pattern string) CollectionOption {
	return func(o *openapi.CreateCollectionRequest) {
		o.Sources = append(o.Sources, openapi.Source{
			S3: &openapi.SourceS3{
				Bucket:  dataset.RocksetPublicDatasets,
				Pattern: &pattern,
			},
		})
	}
}

// StorageCompressionType defined the RocksDB storage compression types.
type StorageCompressionType string

func (s StorageCompressionType) String() string { return string(s) }

const (
	StorageCompressionLZ4  = "LZ4"
	StorageCompressionZSTD = "ZSTD"
)

// WithStorageCompressionType sets the RocksDB storage compression type.
func WithStorageCompressionType(compressionType StorageCompressionType) CollectionOption {
	return func(o *openapi.CreateCollectionRequest) {
		t := compressionType.String()
		o.StorageCompressionType = &t
	}
}
