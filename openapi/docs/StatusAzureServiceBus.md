# StatusAzureServiceBus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstProcessedAt** | Pointer to **string** | Service Bus first message processed time in ISO-8601 format. | [optional] 
**RecordsProcessed** | Pointer to **int64** | Number of records processed. | [optional] 
**Sessions** | Pointer to [**map[string]StatusAzureServiceBusSession**](StatusAzureServiceBusSession.md) | Sessions processed. | [optional] 

## Methods

### NewStatusAzureServiceBus

`func NewStatusAzureServiceBus() *StatusAzureServiceBus`

NewStatusAzureServiceBus instantiates a new StatusAzureServiceBus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusAzureServiceBusWithDefaults

`func NewStatusAzureServiceBusWithDefaults() *StatusAzureServiceBus`

NewStatusAzureServiceBusWithDefaults instantiates a new StatusAzureServiceBus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstProcessedAt

`func (o *StatusAzureServiceBus) GetFirstProcessedAt() string`

GetFirstProcessedAt returns the FirstProcessedAt field if non-nil, zero value otherwise.

### GetFirstProcessedAtOk

`func (o *StatusAzureServiceBus) GetFirstProcessedAtOk() (*string, bool)`

GetFirstProcessedAtOk returns a tuple with the FirstProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstProcessedAt

`func (o *StatusAzureServiceBus) SetFirstProcessedAt(v string)`

SetFirstProcessedAt sets FirstProcessedAt field to given value.

### HasFirstProcessedAt

`func (o *StatusAzureServiceBus) HasFirstProcessedAt() bool`

HasFirstProcessedAt returns a boolean if a field has been set.

### GetRecordsProcessed

`func (o *StatusAzureServiceBus) GetRecordsProcessed() int64`

GetRecordsProcessed returns the RecordsProcessed field if non-nil, zero value otherwise.

### GetRecordsProcessedOk

`func (o *StatusAzureServiceBus) GetRecordsProcessedOk() (*int64, bool)`

GetRecordsProcessedOk returns a tuple with the RecordsProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsProcessed

`func (o *StatusAzureServiceBus) SetRecordsProcessed(v int64)`

SetRecordsProcessed sets RecordsProcessed field to given value.

### HasRecordsProcessed

`func (o *StatusAzureServiceBus) HasRecordsProcessed() bool`

HasRecordsProcessed returns a boolean if a field has been set.

### GetSessions

`func (o *StatusAzureServiceBus) GetSessions() map[string]StatusAzureServiceBusSession`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *StatusAzureServiceBus) GetSessionsOk() (*map[string]StatusAzureServiceBusSession, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *StatusAzureServiceBus) SetSessions(v map[string]StatusAzureServiceBusSession)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *StatusAzureServiceBus) HasSessions() bool`

HasSessions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


