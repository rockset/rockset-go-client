# StatusAzureServiceBusSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastProcessed** | Pointer to **string** | Most recent ISO-8601 date when a message from this session was processed. | [optional] 
**SequenceNumber** | Pointer to **int64** | The last processed sequence number within this session. | [optional] 

## Methods

### NewStatusAzureServiceBusSession

`func NewStatusAzureServiceBusSession() *StatusAzureServiceBusSession`

NewStatusAzureServiceBusSession instantiates a new StatusAzureServiceBusSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusAzureServiceBusSessionWithDefaults

`func NewStatusAzureServiceBusSessionWithDefaults() *StatusAzureServiceBusSession`

NewStatusAzureServiceBusSessionWithDefaults instantiates a new StatusAzureServiceBusSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastProcessed

`func (o *StatusAzureServiceBusSession) GetLastProcessed() string`

GetLastProcessed returns the LastProcessed field if non-nil, zero value otherwise.

### GetLastProcessedOk

`func (o *StatusAzureServiceBusSession) GetLastProcessedOk() (*string, bool)`

GetLastProcessedOk returns a tuple with the LastProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProcessed

`func (o *StatusAzureServiceBusSession) SetLastProcessed(v string)`

SetLastProcessed sets LastProcessed field to given value.

### HasLastProcessed

`func (o *StatusAzureServiceBusSession) HasLastProcessed() bool`

HasLastProcessed returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *StatusAzureServiceBusSession) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *StatusAzureServiceBusSession) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *StatusAzureServiceBusSession) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *StatusAzureServiceBusSession) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


