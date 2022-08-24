# InputField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | Pointer to **string** | The name of a field, parsed as a SQL qualified name. | [optional] 
**IfMissing** | Pointer to **string** | Define the behaviour if fieldName is missing or is null. | [optional] 
**IsDrop** | Pointer to **bool** | If true, then drop fieldName from the document. | [optional] 
**Param** | Pointer to **string** | Sql parameter name. | [optional] 

## Methods

### NewInputField

`func NewInputField() *InputField`

NewInputField instantiates a new InputField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputFieldWithDefaults

`func NewInputFieldWithDefaults() *InputField`

NewInputFieldWithDefaults instantiates a new InputField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *InputField) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *InputField) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *InputField) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *InputField) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetIfMissing

`func (o *InputField) GetIfMissing() string`

GetIfMissing returns the IfMissing field if non-nil, zero value otherwise.

### GetIfMissingOk

`func (o *InputField) GetIfMissingOk() (*string, bool)`

GetIfMissingOk returns a tuple with the IfMissing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfMissing

`func (o *InputField) SetIfMissing(v string)`

SetIfMissing sets IfMissing field to given value.

### HasIfMissing

`func (o *InputField) HasIfMissing() bool`

HasIfMissing returns a boolean if a field has been set.

### GetIsDrop

`func (o *InputField) GetIsDrop() bool`

GetIsDrop returns the IsDrop field if non-nil, zero value otherwise.

### GetIsDropOk

`func (o *InputField) GetIsDropOk() (*bool, bool)`

GetIsDropOk returns a tuple with the IsDrop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDrop

`func (o *InputField) SetIsDrop(v bool)`

SetIsDrop sets IsDrop field to given value.

### HasIsDrop

`func (o *InputField) HasIsDrop() bool`

HasIsDrop returns a boolean if a field has been set.

### GetParam

`func (o *InputField) GetParam() string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *InputField) GetParamOk() (*string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *InputField) SetParam(v string)`

SetParam sets Param field to given value.

### HasParam

`func (o *InputField) HasParam() bool`

HasParam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


