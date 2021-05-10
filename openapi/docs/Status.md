# Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | Status of the Source&#39;s ingestion, one of: INITIALIZING, WATCHING, PROCESSING, COMPLETED, ERROR | [optional] 
**Since** | Pointer to **string** | ISO-8601 date when state was triggered | [optional] 
**Message** | Pointer to **string** | state message | [optional] 
**LastProcessedAt** | Pointer to **string** | ISO-8601 date when source was last processed | [optional] 
**LastProcessedItem** | Pointer to **string** | last source item processed by ingester | [optional] 
**TotalProcessedItems** | Pointer to **int64** | Total items processed of source | [optional] 
**LastErrorAt** | Pointer to **string** | ISO-8601 date when last error occurred | [optional] 
**LastErrorItem** | Pointer to **string** | last source item that errored | [optional] 
**LastErrorReason** | Pointer to **string** | reason for the last error | [optional] 
**TotalErrorItems** | Pointer to **int64** | Total items that errored | [optional] 

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

### GetSince

`func (o *Status) GetSince() string`

GetSince returns the Since field if non-nil, zero value otherwise.

### GetSinceOk

`func (o *Status) GetSinceOk() (*string, bool)`

GetSinceOk returns a tuple with the Since field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSince

`func (o *Status) SetSince(v string)`

SetSince sets Since field to given value.

### HasSince

`func (o *Status) HasSince() bool`

HasSince returns a boolean if a field has been set.

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

### GetLastErrorAt

`func (o *Status) GetLastErrorAt() string`

GetLastErrorAt returns the LastErrorAt field if non-nil, zero value otherwise.

### GetLastErrorAtOk

`func (o *Status) GetLastErrorAtOk() (*string, bool)`

GetLastErrorAtOk returns a tuple with the LastErrorAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorAt

`func (o *Status) SetLastErrorAt(v string)`

SetLastErrorAt sets LastErrorAt field to given value.

### HasLastErrorAt

`func (o *Status) HasLastErrorAt() bool`

HasLastErrorAt returns a boolean if a field has been set.

### GetLastErrorItem

`func (o *Status) GetLastErrorItem() string`

GetLastErrorItem returns the LastErrorItem field if non-nil, zero value otherwise.

### GetLastErrorItemOk

`func (o *Status) GetLastErrorItemOk() (*string, bool)`

GetLastErrorItemOk returns a tuple with the LastErrorItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorItem

`func (o *Status) SetLastErrorItem(v string)`

SetLastErrorItem sets LastErrorItem field to given value.

### HasLastErrorItem

`func (o *Status) HasLastErrorItem() bool`

HasLastErrorItem returns a boolean if a field has been set.

### GetLastErrorReason

`func (o *Status) GetLastErrorReason() string`

GetLastErrorReason returns the LastErrorReason field if non-nil, zero value otherwise.

### GetLastErrorReasonOk

`func (o *Status) GetLastErrorReasonOk() (*string, bool)`

GetLastErrorReasonOk returns a tuple with the LastErrorReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorReason

`func (o *Status) SetLastErrorReason(v string)`

SetLastErrorReason sets LastErrorReason field to given value.

### HasLastErrorReason

`func (o *Status) HasLastErrorReason() bool`

HasLastErrorReason returns a boolean if a field has been set.

### GetTotalErrorItems

`func (o *Status) GetTotalErrorItems() int64`

GetTotalErrorItems returns the TotalErrorItems field if non-nil, zero value otherwise.

### GetTotalErrorItemsOk

`func (o *Status) GetTotalErrorItemsOk() (*int64, bool)`

GetTotalErrorItemsOk returns a tuple with the TotalErrorItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalErrorItems

`func (o *Status) SetTotalErrorItems(v int64)`

SetTotalErrorItems sets TotalErrorItems field to given value.

### HasTotalErrorItems

`func (o *Status) HasTotalErrorItems() bool`

HasTotalErrorItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


