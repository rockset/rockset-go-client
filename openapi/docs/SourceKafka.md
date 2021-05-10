# SourceKafka

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KafkaTopicName** | **string** | The Kafka topic to be tailed | 
**Status** | Pointer to [**StatusKafka**](StatusKafka.md) |  | [optional] 

## Methods

### NewSourceKafka

`func NewSourceKafka(kafkaTopicName string, ) *SourceKafka`

NewSourceKafka instantiates a new SourceKafka object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceKafkaWithDefaults

`func NewSourceKafkaWithDefaults() *SourceKafka`

NewSourceKafkaWithDefaults instantiates a new SourceKafka object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKafkaTopicName

`func (o *SourceKafka) GetKafkaTopicName() string`

GetKafkaTopicName returns the KafkaTopicName field if non-nil, zero value otherwise.

### GetKafkaTopicNameOk

`func (o *SourceKafka) GetKafkaTopicNameOk() (*string, bool)`

GetKafkaTopicNameOk returns a tuple with the KafkaTopicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaTopicName

`func (o *SourceKafka) SetKafkaTopicName(v string)`

SetKafkaTopicName sets KafkaTopicName field to given value.


### GetStatus

`func (o *SourceKafka) GetStatus() StatusKafka`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SourceKafka) GetStatusOk() (*StatusKafka, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SourceKafka) SetStatus(v StatusKafka)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SourceKafka) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


