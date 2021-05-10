# QueryLambdaStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastExecuted** | Pointer to **string** | ISO-8601 date | [optional] 
**LastExecutedBy** | Pointer to **string** | user who last executed Query Lambda | [optional] 
**LastExecutionError** | Pointer to **string** | ISO-8601 date of last execution failure | [optional] 
**LastExecutionErrorMessage** | Pointer to **string** | error message associated with last failed execution | [optional] 

## Methods

### NewQueryLambdaStats

`func NewQueryLambdaStats() *QueryLambdaStats`

NewQueryLambdaStats instantiates a new QueryLambdaStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryLambdaStatsWithDefaults

`func NewQueryLambdaStatsWithDefaults() *QueryLambdaStats`

NewQueryLambdaStatsWithDefaults instantiates a new QueryLambdaStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastExecuted

`func (o *QueryLambdaStats) GetLastExecuted() string`

GetLastExecuted returns the LastExecuted field if non-nil, zero value otherwise.

### GetLastExecutedOk

`func (o *QueryLambdaStats) GetLastExecutedOk() (*string, bool)`

GetLastExecutedOk returns a tuple with the LastExecuted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecuted

`func (o *QueryLambdaStats) SetLastExecuted(v string)`

SetLastExecuted sets LastExecuted field to given value.

### HasLastExecuted

`func (o *QueryLambdaStats) HasLastExecuted() bool`

HasLastExecuted returns a boolean if a field has been set.

### GetLastExecutedBy

`func (o *QueryLambdaStats) GetLastExecutedBy() string`

GetLastExecutedBy returns the LastExecutedBy field if non-nil, zero value otherwise.

### GetLastExecutedByOk

`func (o *QueryLambdaStats) GetLastExecutedByOk() (*string, bool)`

GetLastExecutedByOk returns a tuple with the LastExecutedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutedBy

`func (o *QueryLambdaStats) SetLastExecutedBy(v string)`

SetLastExecutedBy sets LastExecutedBy field to given value.

### HasLastExecutedBy

`func (o *QueryLambdaStats) HasLastExecutedBy() bool`

HasLastExecutedBy returns a boolean if a field has been set.

### GetLastExecutionError

`func (o *QueryLambdaStats) GetLastExecutionError() string`

GetLastExecutionError returns the LastExecutionError field if non-nil, zero value otherwise.

### GetLastExecutionErrorOk

`func (o *QueryLambdaStats) GetLastExecutionErrorOk() (*string, bool)`

GetLastExecutionErrorOk returns a tuple with the LastExecutionError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutionError

`func (o *QueryLambdaStats) SetLastExecutionError(v string)`

SetLastExecutionError sets LastExecutionError field to given value.

### HasLastExecutionError

`func (o *QueryLambdaStats) HasLastExecutionError() bool`

HasLastExecutionError returns a boolean if a field has been set.

### GetLastExecutionErrorMessage

`func (o *QueryLambdaStats) GetLastExecutionErrorMessage() string`

GetLastExecutionErrorMessage returns the LastExecutionErrorMessage field if non-nil, zero value otherwise.

### GetLastExecutionErrorMessageOk

`func (o *QueryLambdaStats) GetLastExecutionErrorMessageOk() (*string, bool)`

GetLastExecutionErrorMessageOk returns a tuple with the LastExecutionErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutionErrorMessage

`func (o *QueryLambdaStats) SetLastExecutionErrorMessage(v string)`

SetLastExecutionErrorMessage sets LastExecutionErrorMessage field to given value.

### HasLastExecutionErrorMessage

`func (o *QueryLambdaStats) HasLastExecutionErrorMessage() bool`

HasLastExecutionErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


