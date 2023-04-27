# KafkaIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsRole** | Pointer to [**AwsRole**](AwsRole.md) |  | [optional] 
**BootstrapServers** | Pointer to **string** | The Kafka bootstrap server url(s). Required only for V3 integration. | [optional] 
**ConnectionString** | Pointer to **string** | Kafka connection string. | [optional] 
**KafkaDataFormat** | Pointer to **string** | The format of the Kafka topics being tailed. | [optional] 
**KafkaTopicNames** | Pointer to **[]string** | Kafka topics to tail. | [optional] 
**SchemaRegistryConfig** | Pointer to [**SchemaRegistryConfig**](SchemaRegistryConfig.md) |  | [optional] 
**SecurityConfig** | Pointer to [**KafkaV3SecurityConfig**](KafkaV3SecurityConfig.md) |  | [optional] 
**SourceStatusByTopic** | Pointer to [**map[string]StatusKafka**](StatusKafka.md) | The status of the Kafka source by topic. | [optional] [readonly] 
**UseV3** | Pointer to **bool** |  | [optional] 

## Methods

### NewKafkaIntegration

`func NewKafkaIntegration() *KafkaIntegration`

NewKafkaIntegration instantiates a new KafkaIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaIntegrationWithDefaults

`func NewKafkaIntegrationWithDefaults() *KafkaIntegration`

NewKafkaIntegrationWithDefaults instantiates a new KafkaIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsRole

`func (o *KafkaIntegration) GetAwsRole() AwsRole`

GetAwsRole returns the AwsRole field if non-nil, zero value otherwise.

### GetAwsRoleOk

`func (o *KafkaIntegration) GetAwsRoleOk() (*AwsRole, bool)`

GetAwsRoleOk returns a tuple with the AwsRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRole

`func (o *KafkaIntegration) SetAwsRole(v AwsRole)`

SetAwsRole sets AwsRole field to given value.

### HasAwsRole

`func (o *KafkaIntegration) HasAwsRole() bool`

HasAwsRole returns a boolean if a field has been set.

### GetBootstrapServers

`func (o *KafkaIntegration) GetBootstrapServers() string`

GetBootstrapServers returns the BootstrapServers field if non-nil, zero value otherwise.

### GetBootstrapServersOk

`func (o *KafkaIntegration) GetBootstrapServersOk() (*string, bool)`

GetBootstrapServersOk returns a tuple with the BootstrapServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapServers

`func (o *KafkaIntegration) SetBootstrapServers(v string)`

SetBootstrapServers sets BootstrapServers field to given value.

### HasBootstrapServers

`func (o *KafkaIntegration) HasBootstrapServers() bool`

HasBootstrapServers returns a boolean if a field has been set.

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

### HasKafkaDataFormat

`func (o *KafkaIntegration) HasKafkaDataFormat() bool`

HasKafkaDataFormat returns a boolean if a field has been set.

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

### HasKafkaTopicNames

`func (o *KafkaIntegration) HasKafkaTopicNames() bool`

HasKafkaTopicNames returns a boolean if a field has been set.

### GetSchemaRegistryConfig

`func (o *KafkaIntegration) GetSchemaRegistryConfig() SchemaRegistryConfig`

GetSchemaRegistryConfig returns the SchemaRegistryConfig field if non-nil, zero value otherwise.

### GetSchemaRegistryConfigOk

`func (o *KafkaIntegration) GetSchemaRegistryConfigOk() (*SchemaRegistryConfig, bool)`

GetSchemaRegistryConfigOk returns a tuple with the SchemaRegistryConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaRegistryConfig

`func (o *KafkaIntegration) SetSchemaRegistryConfig(v SchemaRegistryConfig)`

SetSchemaRegistryConfig sets SchemaRegistryConfig field to given value.

### HasSchemaRegistryConfig

`func (o *KafkaIntegration) HasSchemaRegistryConfig() bool`

HasSchemaRegistryConfig returns a boolean if a field has been set.

### GetSecurityConfig

`func (o *KafkaIntegration) GetSecurityConfig() KafkaV3SecurityConfig`

GetSecurityConfig returns the SecurityConfig field if non-nil, zero value otherwise.

### GetSecurityConfigOk

`func (o *KafkaIntegration) GetSecurityConfigOk() (*KafkaV3SecurityConfig, bool)`

GetSecurityConfigOk returns a tuple with the SecurityConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityConfig

`func (o *KafkaIntegration) SetSecurityConfig(v KafkaV3SecurityConfig)`

SetSecurityConfig sets SecurityConfig field to given value.

### HasSecurityConfig

`func (o *KafkaIntegration) HasSecurityConfig() bool`

HasSecurityConfig returns a boolean if a field has been set.

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

### GetUseV3

`func (o *KafkaIntegration) GetUseV3() bool`

GetUseV3 returns the UseV3 field if non-nil, zero value otherwise.

### GetUseV3Ok

`func (o *KafkaIntegration) GetUseV3Ok() (*bool, bool)`

GetUseV3Ok returns a tuple with the UseV3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseV3

`func (o *KafkaIntegration) SetUseV3(v bool)`

SetUseV3 sets UseV3 field to given value.

### HasUseV3

`func (o *KafkaIntegration) HasUseV3() bool`

HasUseV3 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


