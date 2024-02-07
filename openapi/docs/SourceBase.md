# SourceBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureBlobStorage** | Pointer to [**SourceAzBlobStorageBase**](SourceAzBlobStorageBase.md) |  | [optional] 
**Dynamodb** | Pointer to [**SourceDynamoDbBase**](SourceDynamoDbBase.md) |  | [optional] 
**Gcs** | Pointer to [**SourceGcsBase**](SourceGcsBase.md) |  | [optional] 
**S3** | Pointer to [**SourceS3Base**](SourceS3Base.md) |  | [optional] 

## Methods

### NewSourceBase

`func NewSourceBase() *SourceBase`

NewSourceBase instantiates a new SourceBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceBaseWithDefaults

`func NewSourceBaseWithDefaults() *SourceBase`

NewSourceBaseWithDefaults instantiates a new SourceBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureBlobStorage

`func (o *SourceBase) GetAzureBlobStorage() SourceAzBlobStorageBase`

GetAzureBlobStorage returns the AzureBlobStorage field if non-nil, zero value otherwise.

### GetAzureBlobStorageOk

`func (o *SourceBase) GetAzureBlobStorageOk() (*SourceAzBlobStorageBase, bool)`

GetAzureBlobStorageOk returns a tuple with the AzureBlobStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureBlobStorage

`func (o *SourceBase) SetAzureBlobStorage(v SourceAzBlobStorageBase)`

SetAzureBlobStorage sets AzureBlobStorage field to given value.

### HasAzureBlobStorage

`func (o *SourceBase) HasAzureBlobStorage() bool`

HasAzureBlobStorage returns a boolean if a field has been set.

### GetDynamodb

`func (o *SourceBase) GetDynamodb() SourceDynamoDbBase`

GetDynamodb returns the Dynamodb field if non-nil, zero value otherwise.

### GetDynamodbOk

`func (o *SourceBase) GetDynamodbOk() (*SourceDynamoDbBase, bool)`

GetDynamodbOk returns a tuple with the Dynamodb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamodb

`func (o *SourceBase) SetDynamodb(v SourceDynamoDbBase)`

SetDynamodb sets Dynamodb field to given value.

### HasDynamodb

`func (o *SourceBase) HasDynamodb() bool`

HasDynamodb returns a boolean if a field has been set.

### GetGcs

`func (o *SourceBase) GetGcs() SourceGcsBase`

GetGcs returns the Gcs field if non-nil, zero value otherwise.

### GetGcsOk

`func (o *SourceBase) GetGcsOk() (*SourceGcsBase, bool)`

GetGcsOk returns a tuple with the Gcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcs

`func (o *SourceBase) SetGcs(v SourceGcsBase)`

SetGcs sets Gcs field to given value.

### HasGcs

`func (o *SourceBase) HasGcs() bool`

HasGcs returns a boolean if a field has been set.

### GetS3

`func (o *SourceBase) GetS3() SourceS3Base`

GetS3 returns the S3 field if non-nil, zero value otherwise.

### GetS3Ok

`func (o *SourceBase) GetS3Ok() (*SourceS3Base, bool)`

GetS3Ok returns a tuple with the S3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3

`func (o *SourceBase) SetS3(v SourceS3Base)`

SetS3 sets S3 field to given value.

### HasS3

`func (o *SourceBase) HasS3() bool`

HasS3 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


