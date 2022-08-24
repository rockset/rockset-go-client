# CancelQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**QueryInfo**](QueryInfo.md) |  | [optional] 

## Methods

### NewCancelQueryResponse

`func NewCancelQueryResponse() *CancelQueryResponse`

NewCancelQueryResponse instantiates a new CancelQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelQueryResponseWithDefaults

`func NewCancelQueryResponseWithDefaults() *CancelQueryResponse`

NewCancelQueryResponseWithDefaults instantiates a new CancelQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CancelQueryResponse) GetData() QueryInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CancelQueryResponse) GetDataOk() (*QueryInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CancelQueryResponse) SetData(v QueryInfo)`

SetData sets Data field to given value.

### HasData

`func (o *CancelQueryResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


