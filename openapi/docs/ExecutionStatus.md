# ExecutionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Message describing the execution status. | [optional] 
**State** | Pointer to **string** | The scheduled query lambda execution status | [optional] 

## Methods

### NewExecutionStatus

`func NewExecutionStatus() *ExecutionStatus`

NewExecutionStatus instantiates a new ExecutionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionStatusWithDefaults

`func NewExecutionStatusWithDefaults() *ExecutionStatus`

NewExecutionStatusWithDefaults instantiates a new ExecutionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ExecutionStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ExecutionStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ExecutionStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ExecutionStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetState

`func (o *ExecutionStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ExecutionStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ExecutionStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ExecutionStatus) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


