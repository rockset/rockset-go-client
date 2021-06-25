# FieldOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IndexMode** | Pointer to **string** |  index or no_index | [optional] 
**RangeIndexMode** | Pointer to **string** |  v1_index or no_index | [optional] 
**TypeIndexMode** | Pointer to **string** |  index or no_index | [optional] 
**ColumnIndexMode** | Pointer to **string** |  store or no_store | [optional] 

## Methods

### NewFieldOptions

`func NewFieldOptions() *FieldOptions`

NewFieldOptions instantiates a new FieldOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldOptionsWithDefaults

`func NewFieldOptionsWithDefaults() *FieldOptions`

NewFieldOptionsWithDefaults instantiates a new FieldOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexMode

`func (o *FieldOptions) GetIndexMode() string`

GetIndexMode returns the IndexMode field if non-nil, zero value otherwise.

### GetIndexModeOk

`func (o *FieldOptions) GetIndexModeOk() (*string, bool)`

GetIndexModeOk returns a tuple with the IndexMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexMode

`func (o *FieldOptions) SetIndexMode(v string)`

SetIndexMode sets IndexMode field to given value.

### HasIndexMode

`func (o *FieldOptions) HasIndexMode() bool`

HasIndexMode returns a boolean if a field has been set.

### GetRangeIndexMode

`func (o *FieldOptions) GetRangeIndexMode() string`

GetRangeIndexMode returns the RangeIndexMode field if non-nil, zero value otherwise.

### GetRangeIndexModeOk

`func (o *FieldOptions) GetRangeIndexModeOk() (*string, bool)`

GetRangeIndexModeOk returns a tuple with the RangeIndexMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeIndexMode

`func (o *FieldOptions) SetRangeIndexMode(v string)`

SetRangeIndexMode sets RangeIndexMode field to given value.

### HasRangeIndexMode

`func (o *FieldOptions) HasRangeIndexMode() bool`

HasRangeIndexMode returns a boolean if a field has been set.

### GetTypeIndexMode

`func (o *FieldOptions) GetTypeIndexMode() string`

GetTypeIndexMode returns the TypeIndexMode field if non-nil, zero value otherwise.

### GetTypeIndexModeOk

`func (o *FieldOptions) GetTypeIndexModeOk() (*string, bool)`

GetTypeIndexModeOk returns a tuple with the TypeIndexMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeIndexMode

`func (o *FieldOptions) SetTypeIndexMode(v string)`

SetTypeIndexMode sets TypeIndexMode field to given value.

### HasTypeIndexMode

`func (o *FieldOptions) HasTypeIndexMode() bool`

HasTypeIndexMode returns a boolean if a field has been set.

### GetColumnIndexMode

`func (o *FieldOptions) GetColumnIndexMode() string`

GetColumnIndexMode returns the ColumnIndexMode field if non-nil, zero value otherwise.

### GetColumnIndexModeOk

`func (o *FieldOptions) GetColumnIndexModeOk() (*string, bool)`

GetColumnIndexModeOk returns a tuple with the ColumnIndexMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnIndexMode

`func (o *FieldOptions) SetColumnIndexMode(v string)`

SetColumnIndexMode sets ColumnIndexMode field to given value.

### HasColumnIndexMode

`func (o *FieldOptions) HasColumnIndexMode() bool`

HasColumnIndexMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


