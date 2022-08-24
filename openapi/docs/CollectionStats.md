# CollectionStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocCount** | Pointer to **int64** | Number of documents in the collection. | [optional] 
**PurgedDocCount** | Pointer to **int64** | Number of documents purged from the collection. | [optional] 
**FillProgress** | Pointer to **float64** | Number between 0 and 1 that indicates progress of collection creation. | [optional] 
**LastQueriedMs** | Pointer to **int64** | Milliseconds since Unix epoch Jan 1, 1970. | [optional] 
**LastUpdatedMs** | Pointer to **int64** | Milliseconds since Unix epoch Jan 1, 1970. | [optional] 
**TotalSize** | Pointer to **int64** | Total collection size in bytes. | [optional] 
**TotalIndexSize** | Pointer to **int64** | Total collection index size in bytes. | [optional] 
**RowIndexSize** | Pointer to **int64** | Total collection row index size in bytes. | [optional] 
**ColumnIndexSize** | Pointer to **int64** | Total collection column index size in bytes. | [optional] 
**InvertedIndexSize** | Pointer to **int64** | Total collection inverted index size in bytes. | [optional] 
**RangeIndexSize** | Pointer to **int64** | Total collection range index size in bytes. | [optional] 
**PurgedDocSize** | Pointer to **int64** | Total size of bytes purged in bytes. | [optional] 
**BytesInserted** | Pointer to **int64** | Total number of bytes inserted into the collection. | [optional] 
**BytesOverwritten** | Pointer to **int64** | Total number of bytes overwritten in writing into the collection. | [optional] 
**BulkBytesInserted** | Pointer to **int64** | Total number of bytes inserted into the collection during bulk. | [optional] 
**BulkBytesOverwritten** | Pointer to **int64** | Total number of bytes overwritten in writing into the collection during bulk. | [optional] 

## Methods

### NewCollectionStats

`func NewCollectionStats() *CollectionStats`

NewCollectionStats instantiates a new CollectionStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionStatsWithDefaults

`func NewCollectionStatsWithDefaults() *CollectionStats`

NewCollectionStatsWithDefaults instantiates a new CollectionStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocCount

`func (o *CollectionStats) GetDocCount() int64`

GetDocCount returns the DocCount field if non-nil, zero value otherwise.

### GetDocCountOk

`func (o *CollectionStats) GetDocCountOk() (*int64, bool)`

GetDocCountOk returns a tuple with the DocCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocCount

`func (o *CollectionStats) SetDocCount(v int64)`

SetDocCount sets DocCount field to given value.

### HasDocCount

`func (o *CollectionStats) HasDocCount() bool`

HasDocCount returns a boolean if a field has been set.

### GetPurgedDocCount

`func (o *CollectionStats) GetPurgedDocCount() int64`

GetPurgedDocCount returns the PurgedDocCount field if non-nil, zero value otherwise.

### GetPurgedDocCountOk

`func (o *CollectionStats) GetPurgedDocCountOk() (*int64, bool)`

GetPurgedDocCountOk returns a tuple with the PurgedDocCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgedDocCount

`func (o *CollectionStats) SetPurgedDocCount(v int64)`

SetPurgedDocCount sets PurgedDocCount field to given value.

### HasPurgedDocCount

`func (o *CollectionStats) HasPurgedDocCount() bool`

HasPurgedDocCount returns a boolean if a field has been set.

### GetFillProgress

`func (o *CollectionStats) GetFillProgress() float64`

GetFillProgress returns the FillProgress field if non-nil, zero value otherwise.

### GetFillProgressOk

`func (o *CollectionStats) GetFillProgressOk() (*float64, bool)`

GetFillProgressOk returns a tuple with the FillProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillProgress

`func (o *CollectionStats) SetFillProgress(v float64)`

SetFillProgress sets FillProgress field to given value.

### HasFillProgress

`func (o *CollectionStats) HasFillProgress() bool`

HasFillProgress returns a boolean if a field has been set.

### GetLastQueriedMs

`func (o *CollectionStats) GetLastQueriedMs() int64`

GetLastQueriedMs returns the LastQueriedMs field if non-nil, zero value otherwise.

### GetLastQueriedMsOk

`func (o *CollectionStats) GetLastQueriedMsOk() (*int64, bool)`

GetLastQueriedMsOk returns a tuple with the LastQueriedMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueriedMs

`func (o *CollectionStats) SetLastQueriedMs(v int64)`

SetLastQueriedMs sets LastQueriedMs field to given value.

### HasLastQueriedMs

`func (o *CollectionStats) HasLastQueriedMs() bool`

HasLastQueriedMs returns a boolean if a field has been set.

### GetLastUpdatedMs

`func (o *CollectionStats) GetLastUpdatedMs() int64`

GetLastUpdatedMs returns the LastUpdatedMs field if non-nil, zero value otherwise.

### GetLastUpdatedMsOk

`func (o *CollectionStats) GetLastUpdatedMsOk() (*int64, bool)`

GetLastUpdatedMsOk returns a tuple with the LastUpdatedMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedMs

`func (o *CollectionStats) SetLastUpdatedMs(v int64)`

SetLastUpdatedMs sets LastUpdatedMs field to given value.

### HasLastUpdatedMs

`func (o *CollectionStats) HasLastUpdatedMs() bool`

HasLastUpdatedMs returns a boolean if a field has been set.

### GetTotalSize

`func (o *CollectionStats) GetTotalSize() int64`

GetTotalSize returns the TotalSize field if non-nil, zero value otherwise.

### GetTotalSizeOk

`func (o *CollectionStats) GetTotalSizeOk() (*int64, bool)`

GetTotalSizeOk returns a tuple with the TotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSize

`func (o *CollectionStats) SetTotalSize(v int64)`

SetTotalSize sets TotalSize field to given value.

### HasTotalSize

`func (o *CollectionStats) HasTotalSize() bool`

HasTotalSize returns a boolean if a field has been set.

### GetTotalIndexSize

`func (o *CollectionStats) GetTotalIndexSize() int64`

GetTotalIndexSize returns the TotalIndexSize field if non-nil, zero value otherwise.

### GetTotalIndexSizeOk

`func (o *CollectionStats) GetTotalIndexSizeOk() (*int64, bool)`

GetTotalIndexSizeOk returns a tuple with the TotalIndexSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIndexSize

`func (o *CollectionStats) SetTotalIndexSize(v int64)`

SetTotalIndexSize sets TotalIndexSize field to given value.

### HasTotalIndexSize

`func (o *CollectionStats) HasTotalIndexSize() bool`

HasTotalIndexSize returns a boolean if a field has been set.

### GetRowIndexSize

`func (o *CollectionStats) GetRowIndexSize() int64`

GetRowIndexSize returns the RowIndexSize field if non-nil, zero value otherwise.

### GetRowIndexSizeOk

`func (o *CollectionStats) GetRowIndexSizeOk() (*int64, bool)`

GetRowIndexSizeOk returns a tuple with the RowIndexSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowIndexSize

`func (o *CollectionStats) SetRowIndexSize(v int64)`

SetRowIndexSize sets RowIndexSize field to given value.

### HasRowIndexSize

`func (o *CollectionStats) HasRowIndexSize() bool`

HasRowIndexSize returns a boolean if a field has been set.

### GetColumnIndexSize

`func (o *CollectionStats) GetColumnIndexSize() int64`

GetColumnIndexSize returns the ColumnIndexSize field if non-nil, zero value otherwise.

### GetColumnIndexSizeOk

`func (o *CollectionStats) GetColumnIndexSizeOk() (*int64, bool)`

GetColumnIndexSizeOk returns a tuple with the ColumnIndexSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnIndexSize

`func (o *CollectionStats) SetColumnIndexSize(v int64)`

SetColumnIndexSize sets ColumnIndexSize field to given value.

### HasColumnIndexSize

`func (o *CollectionStats) HasColumnIndexSize() bool`

HasColumnIndexSize returns a boolean if a field has been set.

### GetInvertedIndexSize

`func (o *CollectionStats) GetInvertedIndexSize() int64`

GetInvertedIndexSize returns the InvertedIndexSize field if non-nil, zero value otherwise.

### GetInvertedIndexSizeOk

`func (o *CollectionStats) GetInvertedIndexSizeOk() (*int64, bool)`

GetInvertedIndexSizeOk returns a tuple with the InvertedIndexSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvertedIndexSize

`func (o *CollectionStats) SetInvertedIndexSize(v int64)`

SetInvertedIndexSize sets InvertedIndexSize field to given value.

### HasInvertedIndexSize

`func (o *CollectionStats) HasInvertedIndexSize() bool`

HasInvertedIndexSize returns a boolean if a field has been set.

### GetRangeIndexSize

`func (o *CollectionStats) GetRangeIndexSize() int64`

GetRangeIndexSize returns the RangeIndexSize field if non-nil, zero value otherwise.

### GetRangeIndexSizeOk

`func (o *CollectionStats) GetRangeIndexSizeOk() (*int64, bool)`

GetRangeIndexSizeOk returns a tuple with the RangeIndexSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeIndexSize

`func (o *CollectionStats) SetRangeIndexSize(v int64)`

SetRangeIndexSize sets RangeIndexSize field to given value.

### HasRangeIndexSize

`func (o *CollectionStats) HasRangeIndexSize() bool`

HasRangeIndexSize returns a boolean if a field has been set.

### GetPurgedDocSize

`func (o *CollectionStats) GetPurgedDocSize() int64`

GetPurgedDocSize returns the PurgedDocSize field if non-nil, zero value otherwise.

### GetPurgedDocSizeOk

`func (o *CollectionStats) GetPurgedDocSizeOk() (*int64, bool)`

GetPurgedDocSizeOk returns a tuple with the PurgedDocSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgedDocSize

`func (o *CollectionStats) SetPurgedDocSize(v int64)`

SetPurgedDocSize sets PurgedDocSize field to given value.

### HasPurgedDocSize

`func (o *CollectionStats) HasPurgedDocSize() bool`

HasPurgedDocSize returns a boolean if a field has been set.

### GetBytesInserted

`func (o *CollectionStats) GetBytesInserted() int64`

GetBytesInserted returns the BytesInserted field if non-nil, zero value otherwise.

### GetBytesInsertedOk

`func (o *CollectionStats) GetBytesInsertedOk() (*int64, bool)`

GetBytesInsertedOk returns a tuple with the BytesInserted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesInserted

`func (o *CollectionStats) SetBytesInserted(v int64)`

SetBytesInserted sets BytesInserted field to given value.

### HasBytesInserted

`func (o *CollectionStats) HasBytesInserted() bool`

HasBytesInserted returns a boolean if a field has been set.

### GetBytesOverwritten

`func (o *CollectionStats) GetBytesOverwritten() int64`

GetBytesOverwritten returns the BytesOverwritten field if non-nil, zero value otherwise.

### GetBytesOverwrittenOk

`func (o *CollectionStats) GetBytesOverwrittenOk() (*int64, bool)`

GetBytesOverwrittenOk returns a tuple with the BytesOverwritten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesOverwritten

`func (o *CollectionStats) SetBytesOverwritten(v int64)`

SetBytesOverwritten sets BytesOverwritten field to given value.

### HasBytesOverwritten

`func (o *CollectionStats) HasBytesOverwritten() bool`

HasBytesOverwritten returns a boolean if a field has been set.

### GetBulkBytesInserted

`func (o *CollectionStats) GetBulkBytesInserted() int64`

GetBulkBytesInserted returns the BulkBytesInserted field if non-nil, zero value otherwise.

### GetBulkBytesInsertedOk

`func (o *CollectionStats) GetBulkBytesInsertedOk() (*int64, bool)`

GetBulkBytesInsertedOk returns a tuple with the BulkBytesInserted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkBytesInserted

`func (o *CollectionStats) SetBulkBytesInserted(v int64)`

SetBulkBytesInserted sets BulkBytesInserted field to given value.

### HasBulkBytesInserted

`func (o *CollectionStats) HasBulkBytesInserted() bool`

HasBulkBytesInserted returns a boolean if a field has been set.

### GetBulkBytesOverwritten

`func (o *CollectionStats) GetBulkBytesOverwritten() int64`

GetBulkBytesOverwritten returns the BulkBytesOverwritten field if non-nil, zero value otherwise.

### GetBulkBytesOverwrittenOk

`func (o *CollectionStats) GetBulkBytesOverwrittenOk() (*int64, bool)`

GetBulkBytesOverwrittenOk returns a tuple with the BulkBytesOverwritten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkBytesOverwritten

`func (o *CollectionStats) SetBulkBytesOverwritten(v int64)`

SetBulkBytesOverwritten sets BulkBytesOverwritten field to given value.

### HasBulkBytesOverwritten

`func (o *CollectionStats) HasBulkBytesOverwritten() bool`

HasBulkBytesOverwritten returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


