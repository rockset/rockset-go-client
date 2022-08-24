# ListQueriesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]QueryInfo**](QueryInfo.md) |  | [optional] 

## Methods

### NewListQueriesResponse

`func NewListQueriesResponse() *ListQueriesResponse`

NewListQueriesResponse instantiates a new ListQueriesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListQueriesResponseWithDefaults

`func NewListQueriesResponseWithDefaults() *ListQueriesResponse`

NewListQueriesResponseWithDefaults instantiates a new ListQueriesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListQueriesResponse) GetData() []QueryInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListQueriesResponse) GetDataOk() (*[]QueryInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListQueriesResponse) SetData(v []QueryInfo)`

SetData sets Data field to given value.

### HasData

`func (o *ListQueriesResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


