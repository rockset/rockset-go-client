# QueryPaginationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to **[]map[string]interface{}** | List of documents returned by the query. | [optional] 
**ResultsTotalDocCount** | Pointer to **int64** | Total documents returned by the query. | [optional] 
**Pagination** | Pointer to [**PaginationInfo**](PaginationInfo.md) |  | [optional] 

## Methods

### NewQueryPaginationResponse

`func NewQueryPaginationResponse() *QueryPaginationResponse`

NewQueryPaginationResponse instantiates a new QueryPaginationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryPaginationResponseWithDefaults

`func NewQueryPaginationResponseWithDefaults() *QueryPaginationResponse`

NewQueryPaginationResponseWithDefaults instantiates a new QueryPaginationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *QueryPaginationResponse) GetResults() []map[string]interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *QueryPaginationResponse) GetResultsOk() (*[]map[string]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *QueryPaginationResponse) SetResults(v []map[string]interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *QueryPaginationResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetResultsTotalDocCount

`func (o *QueryPaginationResponse) GetResultsTotalDocCount() int64`

GetResultsTotalDocCount returns the ResultsTotalDocCount field if non-nil, zero value otherwise.

### GetResultsTotalDocCountOk

`func (o *QueryPaginationResponse) GetResultsTotalDocCountOk() (*int64, bool)`

GetResultsTotalDocCountOk returns a tuple with the ResultsTotalDocCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsTotalDocCount

`func (o *QueryPaginationResponse) SetResultsTotalDocCount(v int64)`

SetResultsTotalDocCount sets ResultsTotalDocCount field to given value.

### HasResultsTotalDocCount

`func (o *QueryPaginationResponse) HasResultsTotalDocCount() bool`

HasResultsTotalDocCount returns a boolean if a field has been set.

### GetPagination

`func (o *QueryPaginationResponse) GetPagination() PaginationInfo`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *QueryPaginationResponse) GetPaginationOk() (*PaginationInfo, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *QueryPaginationResponse) SetPagination(v PaginationInfo)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *QueryPaginationResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


