# StatusDynamoDbV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InitialDumpCompletionPercentage** | Pointer to **float64** |  | [optional] 
**State** | Pointer to **string** | State of current ingest for this table. | [optional] 
**StreamLastProcessedAt** | Pointer to **string** | ISO-8601 date when source was last processed. | [optional] 

## Methods

### NewStatusDynamoDbV2

`func NewStatusDynamoDbV2() *StatusDynamoDbV2`

NewStatusDynamoDbV2 instantiates a new StatusDynamoDbV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusDynamoDbV2WithDefaults

`func NewStatusDynamoDbV2WithDefaults() *StatusDynamoDbV2`

NewStatusDynamoDbV2WithDefaults instantiates a new StatusDynamoDbV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInitialDumpCompletionPercentage

`func (o *StatusDynamoDbV2) GetInitialDumpCompletionPercentage() float64`

GetInitialDumpCompletionPercentage returns the InitialDumpCompletionPercentage field if non-nil, zero value otherwise.

### GetInitialDumpCompletionPercentageOk

`func (o *StatusDynamoDbV2) GetInitialDumpCompletionPercentageOk() (*float64, bool)`

GetInitialDumpCompletionPercentageOk returns a tuple with the InitialDumpCompletionPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialDumpCompletionPercentage

`func (o *StatusDynamoDbV2) SetInitialDumpCompletionPercentage(v float64)`

SetInitialDumpCompletionPercentage sets InitialDumpCompletionPercentage field to given value.

### HasInitialDumpCompletionPercentage

`func (o *StatusDynamoDbV2) HasInitialDumpCompletionPercentage() bool`

HasInitialDumpCompletionPercentage returns a boolean if a field has been set.

### GetState

`func (o *StatusDynamoDbV2) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StatusDynamoDbV2) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StatusDynamoDbV2) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StatusDynamoDbV2) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStreamLastProcessedAt

`func (o *StatusDynamoDbV2) GetStreamLastProcessedAt() string`

GetStreamLastProcessedAt returns the StreamLastProcessedAt field if non-nil, zero value otherwise.

### GetStreamLastProcessedAtOk

`func (o *StatusDynamoDbV2) GetStreamLastProcessedAtOk() (*string, bool)`

GetStreamLastProcessedAtOk returns a tuple with the StreamLastProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamLastProcessedAt

`func (o *StatusDynamoDbV2) SetStreamLastProcessedAt(v string)`

SetStreamLastProcessedAt sets StreamLastProcessedAt field to given value.

### HasStreamLastProcessedAt

`func (o *StatusDynamoDbV2) HasStreamLastProcessedAt() bool`

HasStreamLastProcessedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


