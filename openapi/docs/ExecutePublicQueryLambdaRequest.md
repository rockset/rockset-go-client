# ExecutePublicQueryLambdaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultRowLimit** | Pointer to **int32** | Row limit to use if no limit specified in the SQL query text. | [optional] 
**GenerateWarnings** | Pointer to **bool** | Whether to generate warnings. | [optional] 
**Parameters** | Pointer to [**[]QueryParameter**](QueryParameter.md) | List of named parameters. | [optional] 

## Methods

### NewExecutePublicQueryLambdaRequest

`func NewExecutePublicQueryLambdaRequest() *ExecutePublicQueryLambdaRequest`

NewExecutePublicQueryLambdaRequest instantiates a new ExecutePublicQueryLambdaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutePublicQueryLambdaRequestWithDefaults

`func NewExecutePublicQueryLambdaRequestWithDefaults() *ExecutePublicQueryLambdaRequest`

NewExecutePublicQueryLambdaRequestWithDefaults instantiates a new ExecutePublicQueryLambdaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultRowLimit

`func (o *ExecutePublicQueryLambdaRequest) GetDefaultRowLimit() int32`

GetDefaultRowLimit returns the DefaultRowLimit field if non-nil, zero value otherwise.

### GetDefaultRowLimitOk

`func (o *ExecutePublicQueryLambdaRequest) GetDefaultRowLimitOk() (*int32, bool)`

GetDefaultRowLimitOk returns a tuple with the DefaultRowLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRowLimit

`func (o *ExecutePublicQueryLambdaRequest) SetDefaultRowLimit(v int32)`

SetDefaultRowLimit sets DefaultRowLimit field to given value.

### HasDefaultRowLimit

`func (o *ExecutePublicQueryLambdaRequest) HasDefaultRowLimit() bool`

HasDefaultRowLimit returns a boolean if a field has been set.

### GetGenerateWarnings

`func (o *ExecutePublicQueryLambdaRequest) GetGenerateWarnings() bool`

GetGenerateWarnings returns the GenerateWarnings field if non-nil, zero value otherwise.

### GetGenerateWarningsOk

`func (o *ExecutePublicQueryLambdaRequest) GetGenerateWarningsOk() (*bool, bool)`

GetGenerateWarningsOk returns a tuple with the GenerateWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateWarnings

`func (o *ExecutePublicQueryLambdaRequest) SetGenerateWarnings(v bool)`

SetGenerateWarnings sets GenerateWarnings field to given value.

### HasGenerateWarnings

`func (o *ExecutePublicQueryLambdaRequest) HasGenerateWarnings() bool`

HasGenerateWarnings returns a boolean if a field has been set.

### GetParameters

`func (o *ExecutePublicQueryLambdaRequest) GetParameters() []QueryParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ExecutePublicQueryLambdaRequest) GetParametersOk() (*[]QueryParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ExecutePublicQueryLambdaRequest) SetParameters(v []QueryParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ExecutePublicQueryLambdaRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


