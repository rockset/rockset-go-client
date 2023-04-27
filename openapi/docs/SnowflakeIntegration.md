# SnowflakeIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsAccessKey** | Pointer to [**AwsAccessKey**](AwsAccessKey.md) |  | [optional] 
**AwsRole** | Pointer to [**AwsRole**](AwsRole.md) |  | [optional] 
**DefaultWarehouse** | **string** | default snowflake data warehouse name for query execution. Warehouse name can be overridden in the collection. | 
**Password** | **string** | Snowflake database password. | 
**S3ExportPath** | **string** | S3 path used for running &#39;COPY INTO&#39; command on snowflake table. | 
**SnowflakeUrl** | **string** | Snowflake browser url. | 
**UserRole** | Pointer to **string** | Snowflake user role. If unspecified, will use the default user role. | [optional] 
**Username** | **string** | Snowflake database username. | 

## Methods

### NewSnowflakeIntegration

`func NewSnowflakeIntegration(defaultWarehouse string, password string, s3ExportPath string, snowflakeUrl string, username string, ) *SnowflakeIntegration`

NewSnowflakeIntegration instantiates a new SnowflakeIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnowflakeIntegrationWithDefaults

`func NewSnowflakeIntegrationWithDefaults() *SnowflakeIntegration`

NewSnowflakeIntegrationWithDefaults instantiates a new SnowflakeIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsAccessKey

`func (o *SnowflakeIntegration) GetAwsAccessKey() AwsAccessKey`

GetAwsAccessKey returns the AwsAccessKey field if non-nil, zero value otherwise.

### GetAwsAccessKeyOk

`func (o *SnowflakeIntegration) GetAwsAccessKeyOk() (*AwsAccessKey, bool)`

GetAwsAccessKeyOk returns a tuple with the AwsAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsAccessKey

`func (o *SnowflakeIntegration) SetAwsAccessKey(v AwsAccessKey)`

SetAwsAccessKey sets AwsAccessKey field to given value.

### HasAwsAccessKey

`func (o *SnowflakeIntegration) HasAwsAccessKey() bool`

HasAwsAccessKey returns a boolean if a field has been set.

### GetAwsRole

`func (o *SnowflakeIntegration) GetAwsRole() AwsRole`

GetAwsRole returns the AwsRole field if non-nil, zero value otherwise.

### GetAwsRoleOk

`func (o *SnowflakeIntegration) GetAwsRoleOk() (*AwsRole, bool)`

GetAwsRoleOk returns a tuple with the AwsRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRole

`func (o *SnowflakeIntegration) SetAwsRole(v AwsRole)`

SetAwsRole sets AwsRole field to given value.

### HasAwsRole

`func (o *SnowflakeIntegration) HasAwsRole() bool`

HasAwsRole returns a boolean if a field has been set.

### GetDefaultWarehouse

`func (o *SnowflakeIntegration) GetDefaultWarehouse() string`

GetDefaultWarehouse returns the DefaultWarehouse field if non-nil, zero value otherwise.

### GetDefaultWarehouseOk

`func (o *SnowflakeIntegration) GetDefaultWarehouseOk() (*string, bool)`

GetDefaultWarehouseOk returns a tuple with the DefaultWarehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultWarehouse

`func (o *SnowflakeIntegration) SetDefaultWarehouse(v string)`

SetDefaultWarehouse sets DefaultWarehouse field to given value.


### GetPassword

`func (o *SnowflakeIntegration) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SnowflakeIntegration) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SnowflakeIntegration) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetS3ExportPath

`func (o *SnowflakeIntegration) GetS3ExportPath() string`

GetS3ExportPath returns the S3ExportPath field if non-nil, zero value otherwise.

### GetS3ExportPathOk

`func (o *SnowflakeIntegration) GetS3ExportPathOk() (*string, bool)`

GetS3ExportPathOk returns a tuple with the S3ExportPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3ExportPath

`func (o *SnowflakeIntegration) SetS3ExportPath(v string)`

SetS3ExportPath sets S3ExportPath field to given value.


### GetSnowflakeUrl

`func (o *SnowflakeIntegration) GetSnowflakeUrl() string`

GetSnowflakeUrl returns the SnowflakeUrl field if non-nil, zero value otherwise.

### GetSnowflakeUrlOk

`func (o *SnowflakeIntegration) GetSnowflakeUrlOk() (*string, bool)`

GetSnowflakeUrlOk returns a tuple with the SnowflakeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnowflakeUrl

`func (o *SnowflakeIntegration) SetSnowflakeUrl(v string)`

SetSnowflakeUrl sets SnowflakeUrl field to given value.


### GetUserRole

`func (o *SnowflakeIntegration) GetUserRole() string`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *SnowflakeIntegration) GetUserRoleOk() (*string, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRole

`func (o *SnowflakeIntegration) SetUserRole(v string)`

SetUserRole sets UserRole field to given value.

### HasUserRole

`func (o *SnowflakeIntegration) HasUserRole() bool`

HasUserRole returns a boolean if a field has been set.

### GetUsername

`func (o *SnowflakeIntegration) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SnowflakeIntegration) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SnowflakeIntegration) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


