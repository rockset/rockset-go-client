# Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | Status of the Source&#39;s ingestion. | [optional] 
**Message** | Pointer to **string** | State message. | [optional] 
**LastProcessedAt** | Pointer to **string** | ISO-8601 date when source was last processed. | [optional] 
**LastProcessedItem** | Pointer to **string** | Last source item processed by ingester. | [optional] 
**TotalProcessedItems** | Pointer to **int64** | Total items processed of source. | [optional] 
**DetectedSizeBytes** | Pointer to **int64** | Size in bytes detected for the source at collection initialization. This size can be 0 or null for event stream sources. | [optional] 

## Methods

### NewStatus

`func NewStatus() *Status`

NewStatus instantiates a new Status object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusWithDefaults

`func NewStatusWithDefaults() *Status`

NewStatusWithDefaults instantiates a new Status object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *Status) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Status) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Status) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Status) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMessage

`func (o *Status) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Status) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Status) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Status) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetLastProcessedAt

`func (o *Status) GetLastProcessedAt() string`

GetLastProcessedAt returns the LastProcessedAt field if non-nil, zero value otherwise.

### GetLastProcessedAtOk

`func (o *Status) GetLastProcessedAtOk() (*string, bool)`

GetLastProcessedAtOk returns a tuple with the LastProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProcessedAt

`func (o *Status) SetLastProcessedAt(v string)`

SetLastProcessedAt sets LastProcessedAt field to given value.

### HasLastProcessedAt

`func (o *Status) HasLastProcessedAt() bool`

HasLastProcessedAt returns a boolean if a field has been set.

### GetLastProcessedItem

`func (o *Status) GetLastProcessedItem() string`

GetLastProcessedItem returns the LastProcessedItem field if non-nil, zero value otherwise.

### GetLastProcessedItemOk

`func (o *Status) GetLastProcessedItemOk() (*string, bool)`

GetLastProcessedItemOk returns a tuple with the LastProcessedItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProcessedItem

`func (o *Status) SetLastProcessedItem(v string)`

SetLastProcessedItem sets LastProcessedItem field to given value.

### HasLastProcessedItem

`func (o *Status) HasLastProcessedItem() bool`

HasLastProcessedItem returns a boolean if a field has been set.

### GetTotalProcessedItems

`func (o *Status) GetTotalProcessedItems() int64`

GetTotalProcessedItems returns the TotalProcessedItems field if non-nil, zero value otherwise.

### GetTotalProcessedItemsOk

`func (o *Status) GetTotalProcessedItemsOk() (*int64, bool)`

GetTotalProcessedItemsOk returns a tuple with the TotalProcessedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProcessedItems

`func (o *Status) SetTotalProcessedItems(v int64)`

SetTotalProcessedItems sets TotalProcessedItems field to given value.

### HasTotalProcessedItems

`func (o *Status) HasTotalProcessedItems() bool`

HasTotalProcessedItems returns a boolean if a field has been set.

### GetDetectedSizeBytes

`func (o *Status) GetDetectedSizeBytes() int64`

GetDetectedSizeBytes returns the DetectedSizeBytes field if non-nil, zero value otherwise.

### GetDetectedSizeBytesOk

`func (o *Status) GetDetectedSizeBytesOk() (*int64, bool)`

GetDetectedSizeBytesOk returns a tuple with the DetectedSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedSizeBytes

`func (o *Status) SetDetectedSizeBytes(v int64)`

SetDetectedSizeBytes sets DetectedSizeBytes field to given value.

### HasDetectedSizeBytes

`func (o *Status) HasDetectedSizeBytes() bool`

HasDetectedSizeBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


