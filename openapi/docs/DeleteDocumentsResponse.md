# DeleteDocumentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]DocumentStatus**](DocumentStatus.md) | Information about deleted documents. | [optional] 
**LastOffset** | Pointer to **string** | A string representing the collection offset after completing the deletes. | [optional] 

## Methods

### NewDeleteDocumentsResponse

`func NewDeleteDocumentsResponse() *DeleteDocumentsResponse`

NewDeleteDocumentsResponse instantiates a new DeleteDocumentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteDocumentsResponseWithDefaults

`func NewDeleteDocumentsResponseWithDefaults() *DeleteDocumentsResponse`

NewDeleteDocumentsResponseWithDefaults instantiates a new DeleteDocumentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DeleteDocumentsResponse) GetData() []DocumentStatus`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeleteDocumentsResponse) GetDataOk() (*[]DocumentStatus, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeleteDocumentsResponse) SetData(v []DocumentStatus)`

SetData sets Data field to given value.

### HasData

`func (o *DeleteDocumentsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLastOffset

`func (o *DeleteDocumentsResponse) GetLastOffset() string`

GetLastOffset returns the LastOffset field if non-nil, zero value otherwise.

### GetLastOffsetOk

`func (o *DeleteDocumentsResponse) GetLastOffsetOk() (*string, bool)`

GetLastOffsetOk returns a tuple with the LastOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOffset

`func (o *DeleteDocumentsResponse) SetLastOffset(v string)`

SetLastOffset sets LastOffset field to given value.

### HasLastOffset

`func (o *DeleteDocumentsResponse) HasLastOffset() bool`

HasLastOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


