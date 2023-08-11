# ExecuteQueryLambdaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Async** | Pointer to **bool** | If true, the query will run asynchronously for up to 30 minutes. The query request will immediately return with a query id that can be used to retrieve the query status and results. If false or not specified, the query will return with results once completed or timeout after 2 minutes. (To return results directly for shorter queries while still allowing a timeout of up to 30 minutes, set &#x60;async_options.client_timeout_ms&#x60;.)  | [optional] 
**AsyncOptions** | Pointer to [**AsyncQueryOptions**](AsyncQueryOptions.md) |  | [optional] 
**DebugThresholdMs** | Pointer to **int64** | If query execution takes longer than this value, debug information will be logged. If the query text includes the DEBUG hint and this parameter is also provided, only this value will be used and the DEBUG hint will be ignored. | [optional] 
**DefaultRowLimit** | Pointer to **int32** | Row limit to use if no limit specified in the SQL query text. | [optional] 
**InitialPaginateResponseDocCount** | Pointer to **int32** | [DEPRECATED] Use &#x60;max_initial_results&#x60; instead. Number of documents to return in addition to paginating for this query call. Only relevant if &#x60;paginate&#x60; flag is also set. | [optional] 
**MaxInitialResults** | Pointer to **int64** | This limits the maximum number of results in the initial response. A pagination cursor is returned if the number of results exceeds &#x60;max_initial_results&#x60;. If &#x60;max_initial_results&#x60; is not set, all results will be returned in the initial response up to 4 million. If &#x60;max_initial_results&#x60; is set, the value must be between 0 and 100,000. If the query is async and &#x60;client_timeout_ms&#x60; is exceeded, &#x60;max_initial_results&#x60; does not apply since none of the results will be returned with the initial response. | [optional] 
**Paginate** | Pointer to **bool** | Flag to paginate and store the results of this query for later / sequential retrieval. | [optional] 
**Parameters** | Pointer to [**[]QueryParameter**](QueryParameter.md) | List of named parameters. | [optional] 
**TimeoutMs** | Pointer to **int64** | If a query exceeds the specified timeout, the query will automatically stop and return an error. The query timeout defaults to a maximum of 2 minutes. If &#x60;async&#x60; is true, the query timeout defaults to a maximum of 30 minutes. | [optional] 
**VirtualInstanceId** | Pointer to **string** | Virtual instance on which to run the query. | [optional] 

## Methods

### NewExecuteQueryLambdaRequest

`func NewExecuteQueryLambdaRequest() *ExecuteQueryLambdaRequest`

NewExecuteQueryLambdaRequest instantiates a new ExecuteQueryLambdaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteQueryLambdaRequestWithDefaults

`func NewExecuteQueryLambdaRequestWithDefaults() *ExecuteQueryLambdaRequest`

NewExecuteQueryLambdaRequestWithDefaults instantiates a new ExecuteQueryLambdaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsync

`func (o *ExecuteQueryLambdaRequest) GetAsync() bool`

GetAsync returns the Async field if non-nil, zero value otherwise.

### GetAsyncOk

`func (o *ExecuteQueryLambdaRequest) GetAsyncOk() (*bool, bool)`

GetAsyncOk returns a tuple with the Async field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsync

`func (o *ExecuteQueryLambdaRequest) SetAsync(v bool)`

SetAsync sets Async field to given value.

### HasAsync

`func (o *ExecuteQueryLambdaRequest) HasAsync() bool`

HasAsync returns a boolean if a field has been set.

### GetAsyncOptions

`func (o *ExecuteQueryLambdaRequest) GetAsyncOptions() AsyncQueryOptions`

GetAsyncOptions returns the AsyncOptions field if non-nil, zero value otherwise.

### GetAsyncOptionsOk

`func (o *ExecuteQueryLambdaRequest) GetAsyncOptionsOk() (*AsyncQueryOptions, bool)`

GetAsyncOptionsOk returns a tuple with the AsyncOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsyncOptions

`func (o *ExecuteQueryLambdaRequest) SetAsyncOptions(v AsyncQueryOptions)`

SetAsyncOptions sets AsyncOptions field to given value.

### HasAsyncOptions

`func (o *ExecuteQueryLambdaRequest) HasAsyncOptions() bool`

HasAsyncOptions returns a boolean if a field has been set.

### GetDebugThresholdMs

`func (o *ExecuteQueryLambdaRequest) GetDebugThresholdMs() int64`

GetDebugThresholdMs returns the DebugThresholdMs field if non-nil, zero value otherwise.

### GetDebugThresholdMsOk

`func (o *ExecuteQueryLambdaRequest) GetDebugThresholdMsOk() (*int64, bool)`

GetDebugThresholdMsOk returns a tuple with the DebugThresholdMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugThresholdMs

`func (o *ExecuteQueryLambdaRequest) SetDebugThresholdMs(v int64)`

SetDebugThresholdMs sets DebugThresholdMs field to given value.

### HasDebugThresholdMs

`func (o *ExecuteQueryLambdaRequest) HasDebugThresholdMs() bool`

HasDebugThresholdMs returns a boolean if a field has been set.

### GetDefaultRowLimit

`func (o *ExecuteQueryLambdaRequest) GetDefaultRowLimit() int32`

GetDefaultRowLimit returns the DefaultRowLimit field if non-nil, zero value otherwise.

### GetDefaultRowLimitOk

`func (o *ExecuteQueryLambdaRequest) GetDefaultRowLimitOk() (*int32, bool)`

GetDefaultRowLimitOk returns a tuple with the DefaultRowLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRowLimit

`func (o *ExecuteQueryLambdaRequest) SetDefaultRowLimit(v int32)`

SetDefaultRowLimit sets DefaultRowLimit field to given value.

### HasDefaultRowLimit

`func (o *ExecuteQueryLambdaRequest) HasDefaultRowLimit() bool`

HasDefaultRowLimit returns a boolean if a field has been set.

### GetInitialPaginateResponseDocCount

`func (o *ExecuteQueryLambdaRequest) GetInitialPaginateResponseDocCount() int32`

GetInitialPaginateResponseDocCount returns the InitialPaginateResponseDocCount field if non-nil, zero value otherwise.

### GetInitialPaginateResponseDocCountOk

`func (o *ExecuteQueryLambdaRequest) GetInitialPaginateResponseDocCountOk() (*int32, bool)`

GetInitialPaginateResponseDocCountOk returns a tuple with the InitialPaginateResponseDocCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialPaginateResponseDocCount

`func (o *ExecuteQueryLambdaRequest) SetInitialPaginateResponseDocCount(v int32)`

SetInitialPaginateResponseDocCount sets InitialPaginateResponseDocCount field to given value.

### HasInitialPaginateResponseDocCount

`func (o *ExecuteQueryLambdaRequest) HasInitialPaginateResponseDocCount() bool`

HasInitialPaginateResponseDocCount returns a boolean if a field has been set.

### GetMaxInitialResults

`func (o *ExecuteQueryLambdaRequest) GetMaxInitialResults() int64`

GetMaxInitialResults returns the MaxInitialResults field if non-nil, zero value otherwise.

### GetMaxInitialResultsOk

`func (o *ExecuteQueryLambdaRequest) GetMaxInitialResultsOk() (*int64, bool)`

GetMaxInitialResultsOk returns a tuple with the MaxInitialResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInitialResults

`func (o *ExecuteQueryLambdaRequest) SetMaxInitialResults(v int64)`

SetMaxInitialResults sets MaxInitialResults field to given value.

### HasMaxInitialResults

`func (o *ExecuteQueryLambdaRequest) HasMaxInitialResults() bool`

HasMaxInitialResults returns a boolean if a field has been set.

### GetPaginate

`func (o *ExecuteQueryLambdaRequest) GetPaginate() bool`

GetPaginate returns the Paginate field if non-nil, zero value otherwise.

### GetPaginateOk

`func (o *ExecuteQueryLambdaRequest) GetPaginateOk() (*bool, bool)`

GetPaginateOk returns a tuple with the Paginate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaginate

`func (o *ExecuteQueryLambdaRequest) SetPaginate(v bool)`

SetPaginate sets Paginate field to given value.

### HasPaginate

`func (o *ExecuteQueryLambdaRequest) HasPaginate() bool`

HasPaginate returns a boolean if a field has been set.

### GetParameters

`func (o *ExecuteQueryLambdaRequest) GetParameters() []QueryParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ExecuteQueryLambdaRequest) GetParametersOk() (*[]QueryParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ExecuteQueryLambdaRequest) SetParameters(v []QueryParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ExecuteQueryLambdaRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetTimeoutMs

`func (o *ExecuteQueryLambdaRequest) GetTimeoutMs() int64`

GetTimeoutMs returns the TimeoutMs field if non-nil, zero value otherwise.

### GetTimeoutMsOk

`func (o *ExecuteQueryLambdaRequest) GetTimeoutMsOk() (*int64, bool)`

GetTimeoutMsOk returns a tuple with the TimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMs

`func (o *ExecuteQueryLambdaRequest) SetTimeoutMs(v int64)`

SetTimeoutMs sets TimeoutMs field to given value.

### HasTimeoutMs

`func (o *ExecuteQueryLambdaRequest) HasTimeoutMs() bool`

HasTimeoutMs returns a boolean if a field has been set.

### GetVirtualInstanceId

`func (o *ExecuteQueryLambdaRequest) GetVirtualInstanceId() string`

GetVirtualInstanceId returns the VirtualInstanceId field if non-nil, zero value otherwise.

### GetVirtualInstanceIdOk

`func (o *ExecuteQueryLambdaRequest) GetVirtualInstanceIdOk() (*string, bool)`

GetVirtualInstanceIdOk returns a tuple with the VirtualInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualInstanceId

`func (o *ExecuteQueryLambdaRequest) SetVirtualInstanceId(v string)`

SetVirtualInstanceId sets VirtualInstanceId field to given value.

### HasVirtualInstanceId

`func (o *ExecuteQueryLambdaRequest) HasVirtualInstanceId() bool`

HasVirtualInstanceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


