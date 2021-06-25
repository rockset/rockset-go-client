# ExecuteQueryLambdaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameters** | Pointer to [**[]QueryParameter**](QueryParameter.md) | list of named parameters | [optional] 
**DefaultRowLimit** | Pointer to **int32** | Row limit to use if no limit specified in the SQL query text | [optional] 
**GenerateWarnings** | Pointer to **bool** | Whether to generate warnings | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


