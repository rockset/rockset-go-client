# ExecuteQueryLambdaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameters** | Pointer to [**[]QueryParameter**](QueryParameter.md) | List of named parameters. | [optional] 
**DefaultRowLimit** | Pointer to **int32** | Row limit to use if no limit specified in the SQL query text. | [optional] 
**GenerateWarnings** | Pointer to **bool** | Whether to generate warnings. | [optional] 
**Paginate** | Pointer to **bool** | Flag to paginate and store the results of this query for later / sequential retrieval. | [optional] 
**InitialPaginateResponseDocCount** | Pointer to **int32** | Number of documents to return in addition to paginating for this query call. Only relevant if &#x60;paginate&#x60; flag is also set. | [optional] 
**AsyncOptions** | Pointer to [**AsyncQueryOptions**](AsyncQueryOptions.md) |  | [optional] 
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

### GetGenerateWarnings

`func (o *ExecuteQueryLambdaRequest) GetGenerateWarnings() bool`

GetGenerateWarnings returns the GenerateWarnings field if non-nil, zero value otherwise.

### GetGenerateWarningsOk

`func (o *ExecuteQueryLambdaRequest) GetGenerateWarningsOk() (*bool, bool)`

GetGenerateWarningsOk returns a tuple with the GenerateWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateWarnings

`func (o *ExecuteQueryLambdaRequest) SetGenerateWarnings(v bool)`

SetGenerateWarnings sets GenerateWarnings field to given value.

### HasGenerateWarnings

`func (o *ExecuteQueryLambdaRequest) HasGenerateWarnings() bool`

HasGenerateWarnings returns a boolean if a field has been set.

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


