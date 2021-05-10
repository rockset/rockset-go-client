# FieldMask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputPath** | **[]string** |  | 
**Mask** | Pointer to [**FieldMaskMask**](FieldMaskMask.md) |  | [optional] 

## Methods

### NewFieldMask

`func NewFieldMask(inputPath []string, ) *FieldMask`

NewFieldMask instantiates a new FieldMask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldMaskWithDefaults

`func NewFieldMaskWithDefaults() *FieldMask`

NewFieldMaskWithDefaults instantiates a new FieldMask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputPath

`func (o *FieldMask) GetInputPath() []string`

GetInputPath returns the InputPath field if non-nil, zero value otherwise.

### GetInputPathOk

`func (o *FieldMask) GetInputPathOk() (*[]string, bool)`

GetInputPathOk returns a tuple with the InputPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPath

`func (o *FieldMask) SetInputPath(v []string)`

SetInputPath sets InputPath field to given value.


### GetMask

`func (o *FieldMask) GetMask() FieldMaskMask`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *FieldMask) GetMaskOk() (*FieldMaskMask, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *FieldMask) SetMask(v FieldMaskMask)`

SetMask sets Mask field to given value.

### HasMask

`func (o *FieldMask) HasMask() bool`

HasMask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


