# RedshiftIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsAccessKey** | Pointer to [**AwsAccessKey**](AwsAccessKey.md) |  | [optional] 
**Username** | **string** | Username associated with Redshift cluster | 
**Password** | **string** | Password associated with Redshift cluster | 
**Host** | **string** | Redshift Cluster host | 
**Port** | **int32** | Redshift Cluster port | 
**S3BucketPath** | **string** | unload S3 bucket path | 

## Methods

### NewRedshiftIntegration

`func NewRedshiftIntegration(username string, password string, host string, port int32, s3BucketPath string, ) *RedshiftIntegration`

NewRedshiftIntegration instantiates a new RedshiftIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedshiftIntegrationWithDefaults

`func NewRedshiftIntegrationWithDefaults() *RedshiftIntegration`

NewRedshiftIntegrationWithDefaults instantiates a new RedshiftIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsAccessKey

`func (o *RedshiftIntegration) GetAwsAccessKey() AwsAccessKey`

GetAwsAccessKey returns the AwsAccessKey field if non-nil, zero value otherwise.

### GetAwsAccessKeyOk

`func (o *RedshiftIntegration) GetAwsAccessKeyOk() (*AwsAccessKey, bool)`

GetAwsAccessKeyOk returns a tuple with the AwsAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKey

`func (o *RedshiftIntegration) SetAwsAccessKey(v AwsAccessKey)`

SetAwsAccessKey sets AwsAccessKey field to given value.

### HasAwsAccessKey

`func (o *RedshiftIntegration) HasAwsAccessKey() bool`

HasAwsAccessKey returns a boolean if a field has been set.

### GetUsername

`func (o *RedshiftIntegration) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RedshiftIntegration) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RedshiftIntegration) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *RedshiftIntegration) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RedshiftIntegration) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RedshiftIntegration) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetHost

`func (o *RedshiftIntegration) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RedshiftIntegration) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RedshiftIntegration) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *RedshiftIntegration) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RedshiftIntegration) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RedshiftIntegration) SetPort(v int32)`

SetPort sets Port field to given value.


### GetS3BucketPath

`func (o *RedshiftIntegration) GetS3BucketPath() string`

GetS3BucketPath returns the S3BucketPath field if non-nil, zero value otherwise.

### GetS3BucketPathOk

`func (o *RedshiftIntegration) GetS3BucketPathOk() (*string, bool)`

GetS3BucketPathOk returns a tuple with the S3BucketPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3BucketPath

`func (o *RedshiftIntegration) SetS3BucketPath(v string)`

SetS3BucketPath sets S3BucketPath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


