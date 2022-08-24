# DynamodbIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsAccessKey** | Pointer to [**AwsAccessKey**](AwsAccessKey.md) |  | [optional] 
**AwsRole** | Pointer to [**AwsRole**](AwsRole.md) |  | [optional] 
**S3ExportBucketName** | Pointer to **string** | S3 bucket used for export during collection initial dump. | [optional] 

## Methods

### NewDynamodbIntegration

`func NewDynamodbIntegration() *DynamodbIntegration`

NewDynamodbIntegration instantiates a new DynamodbIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamodbIntegrationWithDefaults

`func NewDynamodbIntegrationWithDefaults() *DynamodbIntegration`

NewDynamodbIntegrationWithDefaults instantiates a new DynamodbIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsAccessKey

`func (o *DynamodbIntegration) GetAwsAccessKey() AwsAccessKey`

GetAwsAccessKey returns the AwsAccessKey field if non-nil, zero value otherwise.

### GetAwsAccessKeyOk

`func (o *DynamodbIntegration) GetAwsAccessKeyOk() (*AwsAccessKey, bool)`

GetAwsAccessKeyOk returns a tuple with the AwsAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKey

`func (o *DynamodbIntegration) SetAwsAccessKey(v AwsAccessKey)`

SetAwsAccessKey sets AwsAccessKey field to given value.

### HasAwsAccessKey

`func (o *DynamodbIntegration) HasAwsAccessKey() bool`

HasAwsAccessKey returns a boolean if a field has been set.

### GetAwsRole

`func (o *DynamodbIntegration) GetAwsRole() AwsRole`

GetAwsRole returns the AwsRole field if non-nil, zero value otherwise.

### GetAwsRoleOk

`func (o *DynamodbIntegration) GetAwsRoleOk() (*AwsRole, bool)`

GetAwsRoleOk returns a tuple with the AwsRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRole

`func (o *DynamodbIntegration) SetAwsRole(v AwsRole)`

SetAwsRole sets AwsRole field to given value.

### HasAwsRole

`func (o *DynamodbIntegration) HasAwsRole() bool`

HasAwsRole returns a boolean if a field has been set.

### GetS3ExportBucketName

`func (o *DynamodbIntegration) GetS3ExportBucketName() string`

GetS3ExportBucketName returns the S3ExportBucketName field if non-nil, zero value otherwise.

### GetS3ExportBucketNameOk

`func (o *DynamodbIntegration) GetS3ExportBucketNameOk() (*string, bool)`

GetS3ExportBucketNameOk returns a tuple with the S3ExportBucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3ExportBucketName

`func (o *DynamodbIntegration) SetS3ExportBucketName(v string)`

SetS3ExportBucketName sets S3ExportBucketName field to given value.

### HasS3ExportBucketName

`func (o *DynamodbIntegration) HasS3ExportBucketName() bool`

HasS3ExportBucketName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


