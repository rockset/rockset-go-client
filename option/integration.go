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
	openapi.GcsIntegration
	Description *string
}
type GCSIntegrationOption func(request *GCSIntegration)

func WithGCSIntegrationDescription(desc string) GCSIntegrationOption {
	return func(o *GCSIntegration) {
		o.Description = &desc
	}
}

type GCSCredentialsFn func(*openapi.GcsIntegration)

func GCSServiceAccount(json string) GCSCredentialsFn {
	return func(o *openapi.GcsIntegration) {
		o.GcpServiceAccount = &openapi.GcpServiceAccount{ServiceAccountKeyFileJson: json}
	}
}
