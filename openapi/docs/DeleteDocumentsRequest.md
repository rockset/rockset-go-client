# DeleteDocumentsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DeleteDocumentsRequestData**](DeleteDocumentsRequestData.md) | array of document IDs | 

## Methods

### NewDeleteDocumentsRequest

`func NewDeleteDocumentsRequest(data []DeleteDocumentsRequestData, ) *DeleteDocumentsRequest`

NewDeleteDocumentsRequest instantiates a new DeleteDocumentsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteDocumentsRequestWithDefaults

`func NewDeleteDocumentsRequestWithDefaults() *DeleteDocumentsRequest`

NewDeleteDocumentsRequestWithDefaults instantiates a new DeleteDocumentsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DeleteDocumentsRequest) GetData() []DeleteDocumentsRequestData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeleteDocumentsRequest) GetDataOk() (*[]DeleteDocumentsRequestData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeleteDocumentsRequest) SetData(v []DeleteDocumentsRequestData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


