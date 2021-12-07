# SourceKafka

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KafkaTopicName** | **string** | The Kafka topic to be tailed | 
**Status** | Pointer to [**StatusKafka**](StatusKafka.md) |  | [optional] 
**UseV3** | Pointer to **bool** |  | [optional] 
**OffsetResetPolicy** | Pointer to **string** |  | [optional] 

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

### GetUseV3

`func (o *SourceKafka) GetUseV3() bool`

GetUseV3 returns the UseV3 field if non-nil, zero value otherwise.

### GetUseV3Ok

`func (o *SourceKafka) GetUseV3Ok() (*bool, bool)`

GetUseV3Ok returns a tuple with the UseV3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseV3

`func (o *SourceKafka) SetUseV3(v bool)`

SetUseV3 sets UseV3 field to given value.

### HasUseV3

`func (o *SourceKafka) HasUseV3() bool`

HasUseV3 returns a boolean if a field has been set.

### GetOffsetResetPolicy

`func (o *SourceKafka) GetOffsetResetPolicy() string`

GetOffsetResetPolicy returns the OffsetResetPolicy field if non-nil, zero value otherwise.

### GetOffsetResetPolicyOk

`func (o *SourceKafka) GetOffsetResetPolicyOk() (*string, bool)`

GetOffsetResetPolicyOk returns a tuple with the OffsetResetPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetResetPolicy

`func (o *SourceKafka) SetOffsetResetPolicy(v string)`

SetOffsetResetPolicy sets OffsetResetPolicy field to given value.

### HasOffsetResetPolicy

`func (o *SourceKafka) HasOffsetResetPolicy() bool`

HasOffsetResetPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


