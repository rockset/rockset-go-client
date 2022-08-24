# GetQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**QueryInfo**](QueryInfo.md) |  | [optional] 
**LastOffset** | Pointer to **string** | If this was a write query, this is the log offset the query was written to. | [optional] 

## Methods

### NewGetQueryResponse

`func NewGetQueryResponse() *GetQueryResponse`

NewGetQueryResponse instantiates a new GetQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetQueryResponseWithDefaults

`func NewGetQueryResponseWithDefaults() *GetQueryResponse`

NewGetQueryResponseWithDefaults instantiates a new GetQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetQueryResponse) GetData() QueryInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetQueryResponse) GetDataOk() (*QueryInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetQueryResponse) SetData(v QueryInfo)`

SetData sets Data field to given value.

### HasData

`func (o *GetQueryResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLastOffset

`func (o *GetQueryResponse) GetLastOffset() string`

GetLastOffset returns the LastOffset field if non-nil, zero value otherwise.

### GetLastOffsetOk

`func (o *GetQueryResponse) GetLastOffsetOk() (*string, bool)`

GetLastOffsetOk returns a tuple with the LastOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOffset

`func (o *GetQueryResponse) SetLastOffset(v string)`

SetLastOffset sets LastOffset field to given value.

### HasLastOffset

`func (o *GetQueryResponse) HasLastOffset() bool`

HasLastOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


