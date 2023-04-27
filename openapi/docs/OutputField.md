# OutputField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | Pointer to **string** | The name of a field, parsed as a SQL qualified name. | [optional] 
**OnError** | Pointer to **string** | Error in Mapping execution: &#39;skip&#39; or &#39;fail&#39;. | [optional] 
**Value** | Pointer to [**SqlExpression**](SqlExpression.md) |  | [optional] 

## Methods

### NewOutputField

`func NewOutputField() *OutputField`

NewOutputField instantiates a new OutputField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputFieldWithDefaults

`func NewOutputFieldWithDefaults() *OutputField`

NewOutputFieldWithDefaults instantiates a new OutputField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *OutputField) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *OutputField) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *OutputField) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *OutputField) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetOnError

`func (o *OutputField) GetOnError() string`

GetOnError returns the OnError field if non-nil, zero value otherwise.

### GetOnErrorOk

`func (o *OutputField) GetOnErrorOk() (*string, bool)`

GetOnErrorOk returns a tuple with the OnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnError

`func (o *OutputField) SetOnError(v string)`

SetOnError sets OnError field to given value.

### HasOnError

`func (o *OutputField) HasOnError() bool`

HasOnError returns a boolean if a field has been set.

### GetValue

`func (o *OutputField) GetValue() SqlExpression`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OutputField) GetValueOk() (*SqlExpression, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OutputField) SetValue(v SqlExpression)`

SetValue sets Value field to given value.

### HasValue

`func (o *OutputField) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


