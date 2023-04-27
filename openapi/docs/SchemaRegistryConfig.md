# SchemaRegistryConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The secure API key for schema registry. | [optional] 
**Secret** | Pointer to **string** | The secure API password registry. | [optional] 
**Url** | Pointer to **string** | Schema registry URL. | [optional] 

## Methods

### NewSchemaRegistryConfig

`func NewSchemaRegistryConfig() *SchemaRegistryConfig`

NewSchemaRegistryConfig instantiates a new SchemaRegistryConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaRegistryConfigWithDefaults

`func NewSchemaRegistryConfigWithDefaults() *SchemaRegistryConfig`

NewSchemaRegistryConfigWithDefaults instantiates a new SchemaRegistryConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *SchemaRegistryConfig) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SchemaRegistryConfig) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SchemaRegistryConfig) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SchemaRegistryConfig) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSecret

`func (o *SchemaRegistryConfig) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *SchemaRegistryConfig) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *SchemaRegistryConfig) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *SchemaRegistryConfig) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetUrl

`func (o *SchemaRegistryConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SchemaRegistryConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SchemaRegistryConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SchemaRegistryConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


