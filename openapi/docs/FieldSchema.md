# FieldSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | Pointer to **string** | The name of a field, parsed as a SQL qualified name | [optional] 
**FieldOptions** | Pointer to [**FieldOptions**](FieldOptions.md) |  | [optional] 

## Methods

### NewFieldSchema

`func NewFieldSchema() *FieldSchema`

NewFieldSchema instantiates a new FieldSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldSchemaWithDefaults

`func NewFieldSchemaWithDefaults() *FieldSchema`

NewFieldSchemaWithDefaults instantiates a new FieldSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *FieldSchema) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *FieldSchema) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *FieldSchema) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *FieldSchema) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetFieldOptions

`func (o *FieldSchema) GetFieldOptions() FieldOptions`

GetFieldOptions returns the FieldOptions field if non-nil, zero value otherwise.

### GetFieldOptionsOk

`func (o *FieldSchema) GetFieldOptionsOk() (*FieldOptions, bool)`

GetFieldOptionsOk returns a tuple with the FieldOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldOptions

`func (o *FieldSchema) SetFieldOptions(v FieldOptions)`

SetFieldOptions sets FieldOptions field to given value.

### HasFieldOptions

`func (o *FieldSchema) HasFieldOptions() bool`

HasFieldOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


