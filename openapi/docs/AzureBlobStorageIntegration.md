# AzureBlobStorageIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionString** | **string** | credentials for Azure Blob Storage | 

## Methods

### NewAzureBlobStorageIntegration

`func NewAzureBlobStorageIntegration(connectionString string, ) *AzureBlobStorageIntegration`

NewAzureBlobStorageIntegration instantiates a new AzureBlobStorageIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureBlobStorageIntegrationWithDefaults

`func NewAzureBlobStorageIntegrationWithDefaults() *AzureBlobStorageIntegration`

NewAzureBlobStorageIntegrationWithDefaults instantiates a new AzureBlobStorageIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionString

`func (o *AzureBlobStorageIntegration) GetConnectionString() string`

GetConnectionString returns the ConnectionString field if non-nil, zero value otherwise.

### GetConnectionStringOk

`func (o *AzureBlobStorageIntegration) GetConnectionStringOk() (*string, bool)`

GetConnectionStringOk returns a tuple with the ConnectionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionString

`func (o *AzureBlobStorageIntegration) SetConnectionString(v string)`

SetConnectionString sets ConnectionString field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


