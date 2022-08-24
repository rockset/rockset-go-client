# FieldMappingV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A user specified string that is a name for this mapping. | [optional] 
**IsDropAllFields** | Pointer to **bool** | A boolean that determines whether to drop all fields in this document. If set, input and output fields should not be set | [optional] 
**InputFields** | Pointer to [**[]InputField**](InputField.md) | A List of InputField for this mapping. | [optional] 
**OutputField** | Pointer to [**OutputField**](OutputField.md) |  | [optional] 

## Methods

### NewFieldMappingV2

`func NewFieldMappingV2() *FieldMappingV2`

NewFieldMappingV2 instantiates a new FieldMappingV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldMappingV2WithDefaults

`func NewFieldMappingV2WithDefaults() *FieldMappingV2`

NewFieldMappingV2WithDefaults instantiates a new FieldMappingV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FieldMappingV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FieldMappingV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FieldMappingV2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FieldMappingV2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsDropAllFields

`func (o *FieldMappingV2) GetIsDropAllFields() bool`

GetIsDropAllFields returns the IsDropAllFields field if non-nil, zero value otherwise.

### GetIsDropAllFieldsOk

`func (o *FieldMappingV2) GetIsDropAllFieldsOk() (*bool, bool)`

GetIsDropAllFieldsOk returns a tuple with the IsDropAllFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDropAllFields

`func (o *FieldMappingV2) SetIsDropAllFields(v bool)`

SetIsDropAllFields sets IsDropAllFields field to given value.

### HasIsDropAllFields

`func (o *FieldMappingV2) HasIsDropAllFields() bool`

HasIsDropAllFields returns a boolean if a field has been set.

### GetInputFields

`func (o *FieldMappingV2) GetInputFields() []InputField`

GetInputFields returns the InputFields field if non-nil, zero value otherwise.

### GetInputFieldsOk

`func (o *FieldMappingV2) GetInputFieldsOk() (*[]InputField, bool)`

GetInputFieldsOk returns a tuple with the InputFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFields

`func (o *FieldMappingV2) SetInputFields(v []InputField)`

SetInputFields sets InputFields field to given value.

### HasInputFields

`func (o *FieldMappingV2) HasInputFields() bool`

HasInputFields returns a boolean if a field has been set.

### GetOutputField

`func (o *FieldMappingV2) GetOutputField() OutputField`

GetOutputField returns the OutputField field if non-nil, zero value otherwise.

### GetOutputFieldOk

`func (o *FieldMappingV2) GetOutputFieldOk() (*OutputField, bool)`

GetOutputFieldOk returns a tuple with the OutputField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputField

`func (o *FieldMappingV2) SetOutputField(v OutputField)`

SetOutputField sets OutputField field to given value.

### HasOutputField

`func (o *FieldMappingV2) HasOutputField() bool`

HasOutputField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


