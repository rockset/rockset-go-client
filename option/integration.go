package option

import "github.com/rockset/rockset-go-client/openapi"

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

type AWSCredentials func(o *S3Integration)

func AWSKeys(accessKey, secretKey string) AWSCredentials {
	return func(o *S3Integration) {
		o.AwsAccessKey = &openapi.AwsAccessKey{
			AwsAccessKeyId:     accessKey,
			AwsSecretAccessKey: secretKey,
		}
	}
}

func AWSRole(roleARN string) AWSCredentials {
	return func(o *S3Integration) {
		o.AwsRole = &openapi.AwsRole{
			AwsRoleArn: roleARN,
		}
	}
}
