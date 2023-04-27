# SourceAzureServiceBus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**StatusAzureServiceBus**](StatusAzureServiceBus.md) |  | [optional] 
**Subscription** | Pointer to **string** | The subscription to read from the topic. | [optional] 
**Topic** | Pointer to **string** | Name of the topic which rockset should ingest from. | [optional] 

## Methods

### NewSourceAzureServiceBus

`func NewSourceAzureServiceBus() *SourceAzureServiceBus`

NewSourceAzureServiceBus instantiates a new SourceAzureServiceBus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceAzureServiceBusWithDefaults

`func NewSourceAzureServiceBusWithDefaults() *SourceAzureServiceBus`

NewSourceAzureServiceBusWithDefaults instantiates a new SourceAzureServiceBus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *SourceAzureServiceBus) GetStatus() StatusAzureServiceBus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SourceAzureServiceBus) GetStatusOk() (*StatusAzureServiceBus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SourceAzureServiceBus) SetStatus(v StatusAzureServiceBus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SourceAzureServiceBus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscription

`func (o *SourceAzureServiceBus) GetSubscription() string`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *SourceAzureServiceBus) GetSubscriptionOk() (*string, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *SourceAzureServiceBus) SetSubscription(v string)`

SetSubscription sets Subscription field to given value.

### HasSubscription

`func (o *SourceAzureServiceBus) HasSubscription() bool`

HasSubscription returns a boolean if a field has been set.

### GetTopic

`func (o *SourceAzureServiceBus) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *SourceAzureServiceBus) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *SourceAzureServiceBus) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *SourceAzureServiceBus) HasTopic() bool`

HasTopic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


