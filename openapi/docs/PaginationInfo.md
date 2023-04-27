# PaginationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPageDocCount** | Pointer to **int64** | Number of documents returned in this result set. | [optional] 
**NextCursor** | Pointer to **string** | Cursor to use to get the list of documents. | [optional] 
**NextCursorOffset** | Pointer to **int64** | The doc offset that next_cursor starts at. | [optional] 
**NextPageLink** | Pointer to **string** | Direct link to the next page of results. | [optional] [readonly] 
**StartCursor** | Pointer to **string** | Cursor used to retrieve the first set of documents. | [optional] 

## Methods

### NewPaginationInfo

`func NewPaginationInfo() *PaginationInfo`

NewPaginationInfo instantiates a new PaginationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationInfoWithDefaults

`func NewPaginationInfoWithDefaults() *PaginationInfo`

NewPaginationInfoWithDefaults instantiates a new PaginationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPageDocCount

`func (o *PaginationInfo) GetCurrentPageDocCount() int64`

GetCurrentPageDocCount returns the CurrentPageDocCount field if non-nil, zero value otherwise.

### GetCurrentPageDocCountOk

`func (o *PaginationInfo) GetCurrentPageDocCountOk() (*int64, bool)`

GetCurrentPageDocCountOk returns a tuple with the CurrentPageDocCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPageDocCount

`func (o *PaginationInfo) SetCurrentPageDocCount(v int64)`

SetCurrentPageDocCount sets CurrentPageDocCount field to given value.

### HasCurrentPageDocCount

`func (o *PaginationInfo) HasCurrentPageDocCount() bool`

HasCurrentPageDocCount returns a boolean if a field has been set.

### GetNextCursor

`func (o *PaginationInfo) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *PaginationInfo) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *PaginationInfo) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *PaginationInfo) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetNextCursorOffset

`func (o *PaginationInfo) GetNextCursorOffset() int64`

GetNextCursorOffset returns the NextCursorOffset field if non-nil, zero value otherwise.

### GetNextCursorOffsetOk

`func (o *PaginationInfo) GetNextCursorOffsetOk() (*int64, bool)`

GetNextCursorOffsetOk returns a tuple with the NextCursorOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursorOffset

`func (o *PaginationInfo) SetNextCursorOffset(v int64)`

SetNextCursorOffset sets NextCursorOffset field to given value.

### HasNextCursorOffset

`func (o *PaginationInfo) HasNextCursorOffset() bool`

HasNextCursorOffset returns a boolean if a field has been set.

### GetNextPageLink

`func (o *PaginationInfo) GetNextPageLink() string`

GetNextPageLink returns the NextPageLink field if non-nil, zero value otherwise.

### GetNextPageLinkOk

`func (o *PaginationInfo) GetNextPageLinkOk() (*string, bool)`

GetNextPageLinkOk returns a tuple with the NextPageLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageLink

`func (o *PaginationInfo) SetNextPageLink(v string)`

SetNextPageLink sets NextPageLink field to given value.

### HasNextPageLink

`func (o *PaginationInfo) HasNextPageLink() bool`

HasNextPageLink returns a boolean if a field has been set.

### GetStartCursor

`func (o *PaginationInfo) GetStartCursor() string`

GetStartCursor returns the StartCursor field if non-nil, zero value otherwise.

### GetStartCursorOk

`func (o *PaginationInfo) GetStartCursorOk() (*string, bool)`

GetStartCursorOk returns a tuple with the StartCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartCursor

`func (o *PaginationInfo) SetStartCursor(v string)`

SetStartCursor sets StartCursor field to given value.

### HasStartCursor

`func (o *PaginationInfo) HasStartCursor() bool`

HasStartCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


