# BulkStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataDownloadedBytes** | Pointer to **int64** | Size in bytes of documents downloaded from source during an ongoing or completed bulk ingest. This includes documents that are dropped and reingested. | [optional] 
**DataIndexedBytes** | Pointer to **int64** | Size in bytes of documents indexed. This is the total size of documents after transformations and dropping before indexes are built. | [optional] 
**DataIndexedThroughputBytes** | Pointer to **float64** | Throughput of documents indexed in the last minute measured in bytes/s. This is based off the data_indexed_bytes size. Throughput during the download stage is shown on a per-source granularity in the sources field of the Collection response. | [optional] 
**DocumentsDownloaded** | Pointer to **int64** | Number of documents downloaded from source during an ongoing or completed bulk ingest. This includes documents that are dropped and reingested. | [optional] 
**DownloadComputeMs** | Pointer to **int64** | Bulk ingest compute units in milliseconds used for downloading documents. | [optional] 
**DownloadingStageDoneAt** | Pointer to **string** | ISO-8601 date of when the downloading stage was completed. | [optional] 
**FinalizingStageDoneAt** | Pointer to **string** | ISO-8601 date of when the finalizing stage was completed. | [optional] 
**IndexComputeMs** | Pointer to **int64** | Bulk ingest compute units in milliseconds used for indexing documents. | [optional] 
**IndexingStageDoneAt** | Pointer to **string** | ISO-8601 date of when the indexing stage was completed. | [optional] 
**InitializingStageDoneAt** | Pointer to **string** | ISO-8601 date of when the initializing stage was completed. | [optional] 
**PreIndexSizeBytes** | Pointer to **int64** | Size in bytes of documents before being indexed. This is the total size of documents after decompression, transformations, and dropping. This is equal to data_indexed_bytes after the indexing stage is done unless there are retries during indexing the data. | [optional] 
**ProvisioningStageDoneAt** | Pointer to **string** | ISO-8601 date of when the provisioning stage was completed. | [optional] 
**StartedAt** | Pointer to **string** | ISO-8601 date of when the bulk ingest was started. | [optional] 
**TotalIndexSizeBytes** | Pointer to **int64** | Total size of indexes after the completed bulk ingest. This is the same as collection size. | [optional] 
**TransformationComputeMs** | Pointer to **int64** | Bulk ingest compute units in milliseconds used for ingest transformation. | [optional] 

## Methods

### NewBulkStats

`func NewBulkStats() *BulkStats`

NewBulkStats instantiates a new BulkStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkStatsWithDefaults

`func NewBulkStatsWithDefaults() *BulkStats`

NewBulkStatsWithDefaults instantiates a new BulkStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataDownloadedBytes

`func (o *BulkStats) GetDataDownloadedBytes() int64`

GetDataDownloadedBytes returns the DataDownloadedBytes field if non-nil, zero value otherwise.

### GetDataDownloadedBytesOk

`func (o *BulkStats) GetDataDownloadedBytesOk() (*int64, bool)`

GetDataDownloadedBytesOk returns a tuple with the DataDownloadedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDownloadedBytes

`func (o *BulkStats) SetDataDownloadedBytes(v int64)`

SetDataDownloadedBytes sets DataDownloadedBytes field to given value.

### HasDataDownloadedBytes

`func (o *BulkStats) HasDataDownloadedBytes() bool`

HasDataDownloadedBytes returns a boolean if a field has been set.

### GetDataIndexedBytes

`func (o *BulkStats) GetDataIndexedBytes() int64`

GetDataIndexedBytes returns the DataIndexedBytes field if non-nil, zero value otherwise.

### GetDataIndexedBytesOk

`func (o *BulkStats) GetDataIndexedBytesOk() (*int64, bool)`

GetDataIndexedBytesOk returns a tuple with the DataIndexedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataIndexedBytes

`func (o *BulkStats) SetDataIndexedBytes(v int64)`

SetDataIndexedBytes sets DataIndexedBytes field to given value.

### HasDataIndexedBytes

`func (o *BulkStats) HasDataIndexedBytes() bool`

HasDataIndexedBytes returns a boolean if a field has been set.

### GetDataIndexedThroughputBytes

`func (o *BulkStats) GetDataIndexedThroughputBytes() float64`

GetDataIndexedThroughputBytes returns the DataIndexedThroughputBytes field if non-nil, zero value otherwise.

### GetDataIndexedThroughputBytesOk

`func (o *BulkStats) GetDataIndexedThroughputBytesOk() (*float64, bool)`

GetDataIndexedThroughputBytesOk returns a tuple with the DataIndexedThroughputBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataIndexedThroughputBytes

`func (o *BulkStats) SetDataIndexedThroughputBytes(v float64)`

SetDataIndexedThroughputBytes sets DataIndexedThroughputBytes field to given value.

### HasDataIndexedThroughputBytes

`func (o *BulkStats) HasDataIndexedThroughputBytes() bool`

HasDataIndexedThroughputBytes returns a boolean if a field has been set.

### GetDocumentsDownloaded

`func (o *BulkStats) GetDocumentsDownloaded() int64`

GetDocumentsDownloaded returns the DocumentsDownloaded field if non-nil, zero value otherwise.

### GetDocumentsDownloadedOk

`func (o *BulkStats) GetDocumentsDownloadedOk() (*int64, bool)`

GetDocumentsDownloadedOk returns a tuple with the DocumentsDownloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentsDownloaded

`func (o *BulkStats) SetDocumentsDownloaded(v int64)`

SetDocumentsDownloaded sets DocumentsDownloaded field to given value.

### HasDocumentsDownloaded

`func (o *BulkStats) HasDocumentsDownloaded() bool`

HasDocumentsDownloaded returns a boolean if a field has been set.

### GetDownloadComputeMs

`func (o *BulkStats) GetDownloadComputeMs() int64`

GetDownloadComputeMs returns the DownloadComputeMs field if non-nil, zero value otherwise.

### GetDownloadComputeMsOk

`func (o *BulkStats) GetDownloadComputeMsOk() (*int64, bool)`

GetDownloadComputeMsOk returns a tuple with the DownloadComputeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadComputeMs

`func (o *BulkStats) SetDownloadComputeMs(v int64)`

SetDownloadComputeMs sets DownloadComputeMs field to given value.

### HasDownloadComputeMs

`func (o *BulkStats) HasDownloadComputeMs() bool`

HasDownloadComputeMs returns a boolean if a field has been set.

### GetDownloadingStageDoneAt

`func (o *BulkStats) GetDownloadingStageDoneAt() string`

GetDownloadingStageDoneAt returns the DownloadingStageDoneAt field if non-nil, zero value otherwise.

### GetDownloadingStageDoneAtOk

`func (o *BulkStats) GetDownloadingStageDoneAtOk() (*string, bool)`

GetDownloadingStageDoneAtOk returns a tuple with the DownloadingStageDoneAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadingStageDoneAt

`func (o *BulkStats) SetDownloadingStageDoneAt(v string)`

SetDownloadingStageDoneAt sets DownloadingStageDoneAt field to given value.

### HasDownloadingStageDoneAt

`func (o *BulkStats) HasDownloadingStageDoneAt() bool`

HasDownloadingStageDoneAt returns a boolean if a field has been set.

### GetFinalizingStageDoneAt

`func (o *BulkStats) GetFinalizingStageDoneAt() string`

GetFinalizingStageDoneAt returns the FinalizingStageDoneAt field if non-nil, zero value otherwise.

### GetFinalizingStageDoneAtOk

`func (o *BulkStats) GetFinalizingStageDoneAtOk() (*string, bool)`

GetFinalizingStageDoneAtOk returns a tuple with the FinalizingStageDoneAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalizingStageDoneAt

`func (o *BulkStats) SetFinalizingStageDoneAt(v string)`

SetFinalizingStageDoneAt sets FinalizingStageDoneAt field to given value.

### HasFinalizingStageDoneAt

`func (o *BulkStats) HasFinalizingStageDoneAt() bool`

HasFinalizingStageDoneAt returns a boolean if a field has been set.

### GetIndexComputeMs

`func (o *BulkStats) GetIndexComputeMs() int64`

GetIndexComputeMs returns the IndexComputeMs field if non-nil, zero value otherwise.

### GetIndexComputeMsOk

`func (o *BulkStats) GetIndexComputeMsOk() (*int64, bool)`

GetIndexComputeMsOk returns a tuple with the IndexComputeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexComputeMs

`func (o *BulkStats) SetIndexComputeMs(v int64)`

SetIndexComputeMs sets IndexComputeMs field to given value.

### HasIndexComputeMs

`func (o *BulkStats) HasIndexComputeMs() bool`

HasIndexComputeMs returns a boolean if a field has been set.

### GetIndexingStageDoneAt

`func (o *BulkStats) GetIndexingStageDoneAt() string`

GetIndexingStageDoneAt returns the IndexingStageDoneAt field if non-nil, zero value otherwise.

### GetIndexingStageDoneAtOk

`func (o *BulkStats) GetIndexingStageDoneAtOk() (*string, bool)`

GetIndexingStageDoneAtOk returns a tuple with the IndexingStageDoneAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingStageDoneAt

`func (o *BulkStats) SetIndexingStageDoneAt(v string)`

SetIndexingStageDoneAt sets IndexingStageDoneAt field to given value.

### HasIndexingStageDoneAt

`func (o *BulkStats) HasIndexingStageDoneAt() bool`

HasIndexingStageDoneAt returns a boolean if a field has been set.

### GetInitializingStageDoneAt

`func (o *BulkStats) GetInitializingStageDoneAt() string`

GetInitializingStageDoneAt returns the InitializingStageDoneAt field if non-nil, zero value otherwise.

### GetInitializingStageDoneAtOk

`func (o *BulkStats) GetInitializingStageDoneAtOk() (*string, bool)`

GetInitializingStageDoneAtOk returns a tuple with the InitializingStageDoneAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializingStageDoneAt

`func (o *BulkStats) SetInitializingStageDoneAt(v string)`

SetInitializingStageDoneAt sets InitializingStageDoneAt field to given value.

### HasInitializingStageDoneAt

`func (o *BulkStats) HasInitializingStageDoneAt() bool`

HasInitializingStageDoneAt returns a boolean if a field has been set.

### GetPreIndexSizeBytes

`func (o *BulkStats) GetPreIndexSizeBytes() int64`

GetPreIndexSizeBytes returns the PreIndexSizeBytes field if non-nil, zero value otherwise.

### GetPreIndexSizeBytesOk

`func (o *BulkStats) GetPreIndexSizeBytesOk() (*int64, bool)`

GetPreIndexSizeBytesOk returns a tuple with the PreIndexSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreIndexSizeBytes

`func (o *BulkStats) SetPreIndexSizeBytes(v int64)`

SetPreIndexSizeBytes sets PreIndexSizeBytes field to given value.

### HasPreIndexSizeBytes

`func (o *BulkStats) HasPreIndexSizeBytes() bool`

HasPreIndexSizeBytes returns a boolean if a field has been set.

### GetProvisioningStageDoneAt

`func (o *BulkStats) GetProvisioningStageDoneAt() string`

GetProvisioningStageDoneAt returns the ProvisioningStageDoneAt field if non-nil, zero value otherwise.

### GetProvisioningStageDoneAtOk

`func (o *BulkStats) GetProvisioningStageDoneAtOk() (*string, bool)`

GetProvisioningStageDoneAtOk returns a tuple with the ProvisioningStageDoneAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStageDoneAt

`func (o *BulkStats) SetProvisioningStageDoneAt(v string)`

SetProvisioningStageDoneAt sets ProvisioningStageDoneAt field to given value.

### HasProvisioningStageDoneAt

`func (o *BulkStats) HasProvisioningStageDoneAt() bool`

HasProvisioningStageDoneAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *BulkStats) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BulkStats) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BulkStats) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *BulkStats) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetTotalIndexSizeBytes

`func (o *BulkStats) GetTotalIndexSizeBytes() int64`

GetTotalIndexSizeBytes returns the TotalIndexSizeBytes field if non-nil, zero value otherwise.

### GetTotalIndexSizeBytesOk

`func (o *BulkStats) GetTotalIndexSizeBytesOk() (*int64, bool)`

GetTotalIndexSizeBytesOk returns a tuple with the TotalIndexSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIndexSizeBytes

`func (o *BulkStats) SetTotalIndexSizeBytes(v int64)`

SetTotalIndexSizeBytes sets TotalIndexSizeBytes field to given value.

### HasTotalIndexSizeBytes

`func (o *BulkStats) HasTotalIndexSizeBytes() bool`

HasTotalIndexSizeBytes returns a boolean if a field has been set.

### GetTransformationComputeMs

`func (o *BulkStats) GetTransformationComputeMs() int64`

GetTransformationComputeMs returns the TransformationComputeMs field if non-nil, zero value otherwise.

### GetTransformationComputeMsOk

`func (o *BulkStats) GetTransformationComputeMsOk() (*int64, bool)`

GetTransformationComputeMsOk returns a tuple with the TransformationComputeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformationComputeMs

`func (o *BulkStats) SetTransformationComputeMs(v int64)`

SetTransformationComputeMs sets TransformationComputeMs field to given value.

### HasTransformationComputeMs

`func (o *BulkStats) HasTransformationComputeMs() bool`

HasTransformationComputeMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


