# KafkaIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KafkaTopicNames** | **[]string** | Kafka topics to tail | 
**SourceStatusByTopic** | Pointer to [**map[string]StatusKafka**](StatusKafka.md) | The status of the Kafka source by topic | [optional] [readonly] 
**KafkaDataFormat** | **string** | The format of the Kafka topics being tailed | 
**ConnectionString** | Pointer to **string** | kafka connection string | [optional] [readonly] 

## Methods

### NewKafkaIntegration

`func NewKafkaIntegration(kafkaTopicNames []string, kafkaDataFormat string, ) *KafkaIntegration`

NewKafkaIntegration instantiates a new KafkaIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaIntegrationWithDefaults

`func NewKafkaIntegrationWithDefaults() *KafkaIntegration`

NewKafkaIntegrationWithDefaults instantiates a new KafkaIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKafkaTopicNames

`func (o *KafkaIntegration) GetKafkaTopicNames() []string`

GetKafkaTopicNames returns the KafkaTopicNames field if non-nil, zero value otherwise.

### GetKafkaTopicNamesOk

`func (o *KafkaIntegration) GetKafkaTopicNamesOk() (*[]string, bool)`

GetKafkaTopicNamesOk returns a tuple with the KafkaTopicNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaTopicNames

`func (o *KafkaIntegration) SetKafkaTopicNames(v []string)`

SetKafkaTopicNames sets KafkaTopicNames field to given value.


### GetSourceStatusByTopic

`func (o *KafkaIntegration) GetSourceStatusByTopic() map[string]StatusKafka`

GetSourceStatusByTopic returns the SourceStatusByTopic field if non-nil, zero value otherwise.

### GetSourceStatusByTopicOk

`func (o *KafkaIntegration) GetSourceStatusByTopicOk() (*map[string]StatusKafka, bool)`

GetSourceStatusByTopicOk returns a tuple with the SourceStatusByTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceStatusByTopic

`func (o *KafkaIntegration) SetSourceStatusByTopic(v map[string]StatusKafka)`

SetSourceStatusByTopic sets SourceStatusByTopic field to given value.

### HasSourceStatusByTopic

`func (o *KafkaIntegration) HasSourceStatusByTopic() bool`

HasSourceStatusByTopic returns a boolean if a field has been set.

### GetKafkaDataFormat

`func (o *KafkaIntegration) GetKafkaDataFormat() string`

GetKafkaDataFormat returns the KafkaDataFormat field if non-nil, zero value otherwise.

### GetKafkaDataFormatOk

`func (o *KafkaIntegration) GetKafkaDataFormatOk() (*string, bool)`

GetKafkaDataFormatOk returns a tuple with the KafkaDataFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaDataFormat

`func (o *KafkaIntegration) SetKafkaDataFormat(v string)`

SetKafkaDataFormat sets KafkaDataFormat field to given value.


### GetConnectionString

`func (o *KafkaIntegration) GetConnectionString() string`

GetConnectionString returns the ConnectionString field if non-nil, zero value otherwise.

### GetConnectionStringOk

`func (o *KafkaIntegration) GetConnectionStringOk() (*string, bool)`

GetConnectionStringOk returns a tuple with the ConnectionString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionString

`func (o *KafkaIntegration) SetConnectionString(v string)`

SetConnectionString sets ConnectionString field to given value.

### HasConnectionString

`func (o *KafkaIntegration) HasConnectionString() bool`

HasConnectionString returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


