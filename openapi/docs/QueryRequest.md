# QueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Async** | Pointer to **bool** | If true, the query will run asynchronously for up to 30 minutes. The query request will immediately return with a query id that can be used to retrieve the query status and results. If false or not specified, the query will return with results once completed or timeout after 2 minutes. (To return results directly for shorter queries while still allowing a timeout of up to 30 minutes, set &#x60;async_options.client_timeout_ms&#x60;.)  | [optional] 
**AsyncOptions** | Pointer to [**AsyncQueryOptions**](AsyncQueryOptions.md) |  | [optional] 
**DebugThresholdMs** | Pointer to **int64** | If query execution takes longer than this value, debug information will be logged. If the query text includes the DEBUG hint and this parameter is also provided, only this value will be used and the DEBUG hint will be ignored. | [optional] 
**MaxInitialResults** | Pointer to **int64** | This limits the maximum number of results in the initial response. A pagination cursor is returned if the number of results exceeds &#x60;max_initial_results&#x60;. If &#x60;max_initial_results&#x60; is not set, all results will be returned in the initial response up to 4 million. If &#x60;max_initial_results&#x60; is set, the value must be between 0 and 100,000. If the query is async and &#x60;client_timeout_ms&#x60; is exceeded, &#x60;max_initial_results&#x60; does not apply since none of the results will be returned with the initial response. | [optional] 
**Sql** | [**QueryRequestSql**](QueryRequestSql.md) |  | 
**TimeoutMs** | Pointer to **int64** | If a query exceeds the specified timeout, the query will automatically stop and return an error. The query timeout defaults to a maximum of 2 minutes. If &#x60;async&#x60; is true, the query timeout defaults to a maximum of 30 minutes. | [optional] 

## Methods

### NewQueryRequest

`func NewQueryRequest(sql QueryRequestSql, ) *QueryRequest`

NewQueryRequest instantiates a new QueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryRequestWithDefaults

`func NewQueryRequestWithDefaults() *QueryRequest`

NewQueryRequestWithDefaults instantiates a new QueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsync

`func (o *QueryRequest) GetAsync() bool`

GetAsync returns the Async field if non-nil, zero value otherwise.

### GetAsyncOk

`func (o *QueryRequest) GetAsyncOk() (*bool, bool)`

GetAsyncOk returns a tuple with the Async field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsync

`func (o *QueryRequest) SetAsync(v bool)`

SetAsync sets Async field to given value.

### HasAsync

`func (o *QueryRequest) HasAsync() bool`

HasAsync returns a boolean if a field has been set.

### GetAsyncOptions

`func (o *QueryRequest) GetAsyncOptions() AsyncQueryOptions`

GetAsyncOptions returns the AsyncOptions field if non-nil, zero value otherwise.

### GetAsyncOptionsOk

`func (o *QueryRequest) GetAsyncOptionsOk() (*AsyncQueryOptions, bool)`

GetAsyncOptionsOk returns a tuple with the AsyncOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsyncOptions

`func (o *QueryRequest) SetAsyncOptions(v AsyncQueryOptions)`

SetAsyncOptions sets AsyncOptions field to given value.

### HasAsyncOptions

`func (o *QueryRequest) HasAsyncOptions() bool`

HasAsyncOptions returns a boolean if a field has been set.

### GetDebugThresholdMs

`func (o *QueryRequest) GetDebugThresholdMs() int64`

GetDebugThresholdMs returns the DebugThresholdMs field if non-nil, zero value otherwise.

### GetDebugThresholdMsOk

`func (o *QueryRequest) GetDebugThresholdMsOk() (*int64, bool)`

GetDebugThresholdMsOk returns a tuple with the DebugThresholdMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugThresholdMs

`func (o *QueryRequest) SetDebugThresholdMs(v int64)`

SetDebugThresholdMs sets DebugThresholdMs field to given value.

### HasDebugThresholdMs

`func (o *QueryRequest) HasDebugThresholdMs() bool`

HasDebugThresholdMs returns a boolean if a field has been set.

### GetMaxInitialResults

`func (o *QueryRequest) GetMaxInitialResults() int64`

GetMaxInitialResults returns the MaxInitialResults field if non-nil, zero value otherwise.

### GetMaxInitialResultsOk

`func (o *QueryRequest) GetMaxInitialResultsOk() (*int64, bool)`

GetMaxInitialResultsOk returns a tuple with the MaxInitialResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInitialResults

`func (o *QueryRequest) SetMaxInitialResults(v int64)`

SetMaxInitialResults sets MaxInitialResults field to given value.

### HasMaxInitialResults

`func (o *QueryRequest) HasMaxInitialResults() bool`

HasMaxInitialResults returns a boolean if a field has been set.

### GetSql

`func (o *QueryRequest) GetSql() QueryRequestSql`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *QueryRequest) GetSqlOk() (*QueryRequestSql, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *QueryRequest) SetSql(v QueryRequestSql)`

SetSql sets Sql field to given value.


### GetTimeoutMs

`func (o *QueryRequest) GetTimeoutMs() int64`

GetTimeoutMs returns the TimeoutMs field if non-nil, zero value otherwise.

### GetTimeoutMsOk

`func (o *QueryRequest) GetTimeoutMsOk() (*int64, bool)`

GetTimeoutMsOk returns a tuple with the TimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMs

`func (o *QueryRequest) SetTimeoutMs(v int64)`

SetTimeoutMs sets TimeoutMs field to given value.

### HasTimeoutMs

`func (o *QueryRequest) HasTimeoutMs() bool`

HasTimeoutMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


