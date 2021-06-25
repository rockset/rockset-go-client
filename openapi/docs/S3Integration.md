# S3Integration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsAccessKey** | Pointer to [**AwsAccessKey**](AwsAccessKey.md) |  | [optional] 
**AwsRole** | Pointer to [**AwsRole**](AwsRole.md) |  | [optional] 

## Methods

### NewS3Integration

`func NewS3Integration() *S3Integration`

NewS3Integration instantiates a new S3Integration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3IntegrationWithDefaults

`func NewS3IntegrationWithDefaults() *S3Integration`

NewS3IntegrationWithDefaults instantiates a new S3Integration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsAccessKey

`func (o *S3Integration) GetAwsAccessKey() AwsAccessKey`

GetAwsAccessKey returns the AwsAccessKey field if non-nil, zero value otherwise.

### GetAwsAccessKeyOk

`func (o *S3Integration) GetAwsAccessKeyOk() (*AwsAccessKey, bool)`

GetAwsAccessKeyOk returns a tuple with the AwsAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKey

`func (o *S3Integration) SetAwsAccessKey(v AwsAccessKey)`

SetAwsAccessKey sets AwsAccessKey field to given value.

### HasAwsAccessKey

`func (o *S3Integration) HasAwsAccessKey() bool`

HasAwsAccessKey returns a boolean if a field has been set.

### GetAwsRole

`func (o *S3Integration) GetAwsRole() AwsRole`

GetAwsRole returns the AwsRole field if non-nil, zero value otherwise.

### GetAwsRoleOk

`func (o *S3Integration) GetAwsRoleOk() (*AwsRole, bool)`

GetAwsRoleOk returns a tuple with the AwsRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRole

`func (o *S3Integration) SetAwsRole(v AwsRole)`

SetAwsRole sets AwsRole field to given value.

### HasAwsRole

`func (o *S3Integration) HasAwsRole() bool`

HasAwsRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


