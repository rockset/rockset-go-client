# AddDocumentsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]DocumentStatus**](DocumentStatus.md) | Information about the added documents. | [optional] 
**LastOffset** | Pointer to **string** | A string representing the collection offset after completing the write. | [optional] 

## Methods

### NewAddDocumentsResponse

`func NewAddDocumentsResponse() *AddDocumentsResponse`

NewAddDocumentsResponse instantiates a new AddDocumentsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDocumentsResponseWithDefaults

`func NewAddDocumentsResponseWithDefaults() *AddDocumentsResponse`

NewAddDocumentsResponseWithDefaults instantiates a new AddDocumentsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AddDocumentsResponse) GetData() []DocumentStatus`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AddDocumentsResponse) GetDataOk() (*[]DocumentStatus, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AddDocumentsResponse) SetData(v []DocumentStatus)`

SetData sets Data field to given value.

### HasData

`func (o *AddDocumentsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLastOffset

`func (o *AddDocumentsResponse) GetLastOffset() string`

GetLastOffset returns the LastOffset field if non-nil, zero value otherwise.

### GetLastOffsetOk

`func (o *AddDocumentsResponse) GetLastOffsetOk() (*string, bool)`

GetLastOffsetOk returns a tuple with the LastOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOffset

`func (o *AddDocumentsResponse) SetLastOffset(v string)`

SetLastOffset sets LastOffset field to given value.

### HasLastOffset

`func (o *AddDocumentsResponse) HasLastOffset() bool`

HasLastOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


