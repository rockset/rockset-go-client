# KafkaV3SecurityConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **string** |  | [optional] 
**Secret** | Pointer to **string** |  | [optional] 

## Methods

### NewKafkaV3SecurityConfig

`func NewKafkaV3SecurityConfig() *KafkaV3SecurityConfig`

NewKafkaV3SecurityConfig instantiates a new KafkaV3SecurityConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaV3SecurityConfigWithDefaults

`func NewKafkaV3SecurityConfigWithDefaults() *KafkaV3SecurityConfig`

NewKafkaV3SecurityConfigWithDefaults instantiates a new KafkaV3SecurityConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *KafkaV3SecurityConfig) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *KafkaV3SecurityConfig) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *KafkaV3SecurityConfig) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *KafkaV3SecurityConfig) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetSecret

`func (o *KafkaV3SecurityConfig) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *KafkaV3SecurityConfig) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *KafkaV3SecurityConfig) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *KafkaV3SecurityConfig) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


