# FieldPartition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | Pointer to **string** | The name of a field, parsed as a SQL qualified name | [optional] 
**Type** | Pointer to **string** | The type of partitions on a field | [optional] 
**Keys** | Pointer to **[]string** | The values for partitioning of a field. Unneeded if the partition type is AUTO. | [optional] 

## Methods

### NewFieldPartition

`func NewFieldPartition() *FieldPartition`

NewFieldPartition instantiates a new FieldPartition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldPartitionWithDefaults

`func NewFieldPartitionWithDefaults() *FieldPartition`

NewFieldPartitionWithDefaults instantiates a new FieldPartition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *FieldPartition) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *FieldPartition) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *FieldPartition) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *FieldPartition) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetType

`func (o *FieldPartition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FieldPartition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FieldPartition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FieldPartition) HasType() bool`

HasType returns a boolean if a field has been set.

### GetKeys

`func (o *FieldPartition) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *FieldPartition) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *FieldPartition) SetKeys(v []string)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *FieldPartition) HasKeys() bool`

HasKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


