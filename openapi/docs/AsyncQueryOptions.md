# AsyncQueryOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientTimeoutMs** | Pointer to **int64** | The maximum amount of time that the client is willing to wait for the query to complete. If the query is not complete by this timeout, a response will be returned with a &#x60;query_id&#x60; that can be used to check the status of the query and retrieve results once the query has completed. | [optional] 
**MaxInitialResults** | Pointer to **int64** | The maximum number of results you will receive as a client. If the query exceeds this limit, the remaining results can be requested using a returned pagination cursor. In addition, there is a maximum response size of 100MiB so fewer than &#x60;max_results&#x60; may be returned. | [optional] 
**TimeoutMs** | Pointer to **int64** | The maximum amount of time that the system will attempt to complete query execution before aborting the query and returning an error. | [optional] 

## Methods

### NewAsyncQueryOptions

`func NewAsyncQueryOptions() *AsyncQueryOptions`

NewAsyncQueryOptions instantiates a new AsyncQueryOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncQueryOptionsWithDefaults

`func NewAsyncQueryOptionsWithDefaults() *AsyncQueryOptions`

NewAsyncQueryOptionsWithDefaults instantiates a new AsyncQueryOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientTimeoutMs

`func (o *AsyncQueryOptions) GetClientTimeoutMs() int64`

GetClientTimeoutMs returns the ClientTimeoutMs field if non-nil, zero value otherwise.

### GetClientTimeoutMsOk

`func (o *AsyncQueryOptions) GetClientTimeoutMsOk() (*int64, bool)`

GetClientTimeoutMsOk returns a tuple with the ClientTimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTimeoutMs

`func (o *AsyncQueryOptions) SetClientTimeoutMs(v int64)`

SetClientTimeoutMs sets ClientTimeoutMs field to given value.

### HasClientTimeoutMs

`func (o *AsyncQueryOptions) HasClientTimeoutMs() bool`

HasClientTimeoutMs returns a boolean if a field has been set.

### GetMaxInitialResults

`func (o *AsyncQueryOptions) GetMaxInitialResults() int64`

GetMaxInitialResults returns the MaxInitialResults field if non-nil, zero value otherwise.

### GetMaxInitialResultsOk

`func (o *AsyncQueryOptions) GetMaxInitialResultsOk() (*int64, bool)`

GetMaxInitialResultsOk returns a tuple with the MaxInitialResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxInitialResults

`func (o *AsyncQueryOptions) SetMaxInitialResults(v int64)`

SetMaxInitialResults sets MaxInitialResults field to given value.

### HasMaxInitialResults

`func (o *AsyncQueryOptions) HasMaxInitialResults() bool`

HasMaxInitialResults returns a boolean if a field has been set.

### GetTimeoutMs

`func (o *AsyncQueryOptions) GetTimeoutMs() int64`

GetTimeoutMs returns the TimeoutMs field if non-nil, zero value otherwise.

### GetTimeoutMsOk

`func (o *AsyncQueryOptions) GetTimeoutMsOk() (*int64, bool)`

GetTimeoutMsOk returns a tuple with the TimeoutMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutMs

`func (o *AsyncQueryOptions) SetTimeoutMs(v int64)`

SetTimeoutMs sets TimeoutMs field to given value.

### HasTimeoutMs

`func (o *AsyncQueryOptions) HasTimeoutMs() bool`

HasTimeoutMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


