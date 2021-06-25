package option

import "github.com/rockset/rockset-go-client/openapi"

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

type RedshiftIntegration struct {
	openapi.RedshiftIntegration
	Description *string
}
type RedshiftIntegrationOption func(request *RedshiftIntegration)

func WithRedshiftIntegrationDescription(desc string) RedshiftIntegrationOption {
	return func(o *RedshiftIntegration) {
		o.Description = &desc
	}
}

func WithRedshiftIntegrationConfig(s3BucketPath, username, password, host string,
	port int32) RedshiftIntegrationOption {
	return func(o *RedshiftIntegration) {
		o.RedshiftIntegration.Username = username
		o.RedshiftIntegration.Password = password
		o.RedshiftIntegration.Host = host
		o.RedshiftIntegration.Port = port
		o.RedshiftIntegration.S3BucketPath = s3BucketPath
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

type KafkaIntegration struct {
	Description *string
}
type KafkaIntegrationOption func(request *KafkaIntegration)

func WithKafkaIntegrationDescription(desc string) KafkaIntegrationOption {
	return func(o *KafkaIntegration) {
		o.Description = &desc
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
