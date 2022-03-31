# SourceAzureBlobStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Container** | Pointer to **string** | name of Azure blob Storage container you want to ingest from | [optional] 
**Prefix** | Pointer to **string** | Prefix that selects blobs to ingest. | [optional] 
**Pattern** | Pointer to **string** | Glob-style pattern that selects keys to ingest. Only either prefix or pattern can be specified. | [optional] 
**BlobCountDownloaded** | Pointer to **int64** |  | [optional] [readonly] 
**BlobCountTotal** | Pointer to **int64** |  | [optional] [readonly] 
**BlobBytesTotal** | Pointer to **int64** |  | [optional] [readonly] 

## Methods

### NewSourceAzureBlobStorage

`func NewSourceAzureBlobStorage() *SourceAzureBlobStorage`

NewSourceAzureBlobStorage instantiates a new SourceAzureBlobStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceAzureBlobStorageWithDefaults

`func NewSourceAzureBlobStorageWithDefaults() *SourceAzureBlobStorage`

NewSourceAzureBlobStorageWithDefaults instantiates a new SourceAzureBlobStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainer

`func (o *SourceAzureBlobStorage) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *SourceAzureBlobStorage) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *SourceAzureBlobStorage) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *SourceAzureBlobStorage) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetPrefix

`func (o *SourceAzureBlobStorage) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *SourceAzureBlobStorage) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *SourceAzureBlobStorage) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *SourceAzureBlobStorage) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetPattern

`func (o *SourceAzureBlobStorage) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *SourceAzureBlobStorage) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *SourceAzureBlobStorage) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *SourceAzureBlobStorage) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetBlobCountDownloaded

`func (o *SourceAzureBlobStorage) GetBlobCountDownloaded() int64`

GetBlobCountDownloaded returns the BlobCountDownloaded field if non-nil, zero value otherwise.

### GetBlobCountDownloadedOk

`func (o *SourceAzureBlobStorage) GetBlobCountDownloadedOk() (*int64, bool)`

GetBlobCountDownloadedOk returns a tuple with the BlobCountDownloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobCountDownloaded

`func (o *SourceAzureBlobStorage) SetBlobCountDownloaded(v int64)`

SetBlobCountDownloaded sets BlobCountDownloaded field to given value.

### HasBlobCountDownloaded

`func (o *SourceAzureBlobStorage) HasBlobCountDownloaded() bool`

HasBlobCountDownloaded returns a boolean if a field has been set.

### GetBlobCountTotal

`func (o *SourceAzureBlobStorage) GetBlobCountTotal() int64`

GetBlobCountTotal returns the BlobCountTotal field if non-nil, zero value otherwise.

### GetBlobCountTotalOk

`func (o *SourceAzureBlobStorage) GetBlobCountTotalOk() (*int64, bool)`

GetBlobCountTotalOk returns a tuple with the BlobCountTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobCountTotal

`func (o *SourceAzureBlobStorage) SetBlobCountTotal(v int64)`

SetBlobCountTotal sets BlobCountTotal field to given value.

### HasBlobCountTotal

`func (o *SourceAzureBlobStorage) HasBlobCountTotal() bool`

HasBlobCountTotal returns a boolean if a field has been set.

### GetBlobBytesTotal

`func (o *SourceAzureBlobStorage) GetBlobBytesTotal() int64`

GetBlobBytesTotal returns the BlobBytesTotal field if non-nil, zero value otherwise.

### GetBlobBytesTotalOk

`func (o *SourceAzureBlobStorage) GetBlobBytesTotalOk() (*int64, bool)`

GetBlobBytesTotalOk returns a tuple with the BlobBytesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobBytesTotal

`func (o *SourceAzureBlobStorage) SetBlobBytesTotal(v int64)`

SetBlobBytesTotal sets BlobBytesTotal field to given value.

### HasBlobBytesTotal

`func (o *SourceAzureBlobStorage) HasBlobBytesTotal() bool`

HasBlobBytesTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


