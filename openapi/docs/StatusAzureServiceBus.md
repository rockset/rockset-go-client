# StatusAzureServiceBus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstProcessedAt** | Pointer to **time.Time** | Service Bus first message processed time in ISO-8601 format | [optional] 
**RecordsProcessed** | Pointer to **int64** | Number of records processed | [optional] 
**LastProcessedAt** | Pointer to **time.Time** | ISO-8601 date when the last message was processed | [optional] 

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

`func (o *StatusAzureServiceBus) GetFirstProcessedAt() time.Time`

GetFirstProcessedAt returns the FirstProcessedAt field if non-nil, zero value otherwise.

### GetFirstProcessedAtOk

`func (o *StatusAzureServiceBus) GetFirstProcessedAtOk() (*time.Time, bool)`

GetFirstProcessedAtOk returns a tuple with the FirstProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstProcessedAt

`func (o *StatusAzureServiceBus) SetFirstProcessedAt(v time.Time)`

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

### GetLastProcessedAt

`func (o *StatusAzureServiceBus) GetLastProcessedAt() time.Time`

GetLastProcessedAt returns the LastProcessedAt field if non-nil, zero value otherwise.

### GetLastProcessedAtOk

`func (o *StatusAzureServiceBus) GetLastProcessedAtOk() (*time.Time, bool)`

GetLastProcessedAtOk returns a tuple with the LastProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProcessedAt

`func (o *StatusAzureServiceBus) SetLastProcessedAt(v time.Time)`

SetLastProcessedAt sets LastProcessedAt field to given value.

### HasLastProcessedAt

`func (o *StatusAzureServiceBus) HasLastProcessedAt() bool`

HasLastProcessedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


