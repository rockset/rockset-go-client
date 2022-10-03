package option

import (
	"github.com/rockset/rockset-go-client/openapi"
)

type AWSCredentials struct {
	*openapi.AwsAccessKey
	*openapi.AwsRole
}
type AWSCredentialsFn func(o *AWSCredentials)

func AWSKeys(accessKey, secretKey string) AWSCredentialsFn {
	return func(o *AWSCredentials) {
		o.AwsAccessKey = &openapi.AwsAccessKey{
			AwsAccessKeyId:     accessKey,
			AwsSecretAccessKey: secretKey,
		}
	}
}

func AWSRole(roleARN string) AWSCredentialsFn {
	return func(o *AWSCredentials) {
		o.AwsRole = &openapi.AwsRole{
			AwsRoleArn: roleARN,
		}
	}
}

type S3Integration struct {
	openapi.S3Integration
	Description *string
}
type S3IntegrationOption func(request *S3Integration)

func WithS3IntegrationDescription(desc string) S3IntegrationOption {
	return func(o *S3Integration) {
		o.Description = &desc
	}
}

type KinesisIntegration struct {
	openapi.KinesisIntegration
	Description *string
}
type KinesisIntegrationOption func(request *KinesisIntegration)

func WithKinesisIntegrationDescription(desc string) KinesisIntegrationOption {
	return func(o *KinesisIntegration) {
		o.Description = &desc
	}
}

type DynamoDBIntegration struct {
	openapi.DynamodbIntegration
	Description *string
}
type DynamoDBIntegrationOption func(request *DynamoDBIntegration)

func WithDynamoDBIntegrationDescription(desc string) DynamoDBIntegrationOption {
	return func(o *DynamoDBIntegration) {
		o.Description = &desc
	}
}

type GCSIntegration struct {
	Description *string
}
type GCSIntegrationOption func(request *GCSIntegration)

func WithGCSIntegrationDescription(desc string) GCSIntegrationOption {
	return func(o *GCSIntegration) {
		o.Description = &desc
	}
}

type SegmentIntegration struct {
	Description *string
}
type SegmentIntegrationOption func(request *SegmentIntegration)

func WithSegmentIntegrationDescription(desc string) SegmentIntegrationOption {
	return func(o *SegmentIntegration) {
		o.Description = &desc
	}
}

// KafkaFormat is the definition of the Kafka format
type KafkaFormat string

// String returns the string representation of the Kafka format
func (f KafkaFormat) String() string {
	return string(f)
}

const (
	// KafkaFormatJSON is the JSON format for Kafka.
	KafkaFormatJSON KafkaFormat = "JSON"
	// KafkaFormatAVRO is the AVRO format for Kafka. If used, the option.WithKafkaSchemaRegistryConfig()
	// must be used when creating the integration.
	KafkaFormatAVRO KafkaFormat = "AVRO"
)

// KafkaIntegrationTopicState is the definition of topic states for a Kafka integration.
type KafkaIntegrationTopicState string

const (
	KafkaIntegrationActive    KafkaIntegrationTopicState = "ACTIVE"
	KafkaIntegrationNoDocsYet KafkaIntegrationTopicState = "NO_DOCS_YET"
	KafkaIntegrationDormant   KafkaIntegrationTopicState = "DORMANT"
)

type KafkaIntegration struct {
	Description *string
	Config      openapi.KafkaIntegration
}

type KafkaIntegrationOption func(request *KafkaIntegration)

func WithKafkaIntegrationDescription(desc string) KafkaIntegrationOption {
	return func(o *KafkaIntegration) {
		o.Description = &desc
	}
}

// WithKafkaIntegrationTopic adds a topic name to the integration. Can be specified multiple times.
func WithKafkaIntegrationTopic(topic string) KafkaIntegrationOption {
	return func(o *KafkaIntegration) {
		o.Config.KafkaTopicNames = append(o.Config.KafkaTopicNames, topic)
	}
}

func WithKafkaConnectionString(s string) KafkaIntegrationOption {
	return func(o *KafkaIntegration) {
		o.Config.ConnectionString = &s
	}
}

func WithKafkaIntegrationConfig(config openapi.KafkaIntegration) KafkaIntegrationOption {
	return func(o *KafkaIntegration) {
		o.Config = config
	}
}

func WithKafkaDataFormat(format KafkaFormat) KafkaIntegrationOption {
	return func(o *KafkaIntegration) {
		o.Config.KafkaDataFormat = openapi.PtrString(string(format))
	}
}

func WithKafkaV3() KafkaIntegrationOption {
	return func(o *KafkaIntegration) {
		o.Config.UseV3 = openapi.PtrBool(true)
	}
}

func WithKafkaBootstrapServers(servers string) KafkaIntegrationOption {
	return func(o *KafkaIntegration) {
		o.Config.BootstrapServers = &servers
	}
}

func WithKafkaSecurityConfig(apiKey, secret string) KafkaIntegrationOption {
	return func(o *KafkaIntegration) {
		o.Config.SecurityConfig = &openapi.KafkaV3SecurityConfig{
			ApiKey: &apiKey,
			Secret: &secret,
		}
	}
}

// WithKafkaSchemaRegistryConfig is required when the Kafka format is rockset.KafkaFormatJSON.
// A Kafka integration without schema registry configured can only be used to ingest from topics
// serving JSON messages.
func WithKafkaSchemaRegistryConfig(url, apiKey, secret string) KafkaIntegrationOption {
	return func(o *KafkaIntegration) {
		o.Config.SchemaRegistryConfig = &openapi.SchemaRegistryConfig{
			Url:    &url,
			Key:    &apiKey,
			Secret: &secret,
		}
	}
}

func WithKafkaStatusByTopic(topic string, status openapi.StatusKafka) KafkaIntegrationOption {
	return func(o *KafkaIntegration) {
		if o.Config.SourceStatusByTopic == nil {
			m := make(map[string]openapi.StatusKafka)
			o.Config.SourceStatusByTopic = &m
		}
		(*o.Config.SourceStatusByTopic)[topic] = status
	}
}

type MongoDBIntegration struct {
	Description *string
}
type MongoDBIntegrationOption func(request *MongoDBIntegration)

func WithMongoDBIntegrationDescription(desc string) MongoDBIntegrationOption {
	return func(o *MongoDBIntegration) {
		o.Description = &desc
	}
}
