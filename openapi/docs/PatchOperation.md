# PatchOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | **string** | [JSON Patch operation](https://datatracker.ietf.org/doc/html/rfc6902#page-4) to be performed in this patch. Case insensitive. | 
**Path** | **string** | [JSON Pointer](https://datatracker.ietf.org/doc/html/rfc6901) referencing a location in the target document where the operation is performed | 
**Value** | Pointer to **map[string]interface{}** | Value used in the patch operation. Required for &#x60;ADD&#x60;, &#x60;REPLACE&#x60;, &#x60;TEST&#x60;, and &#x60;INCREMENT&#x60; operations. | [optional] 
**From** | Pointer to **string** | [JSON Pointer](https://datatracker.ietf.org/doc/html/rfc6901) referencing a location in the target document. Required for &#x60;COPY&#x60; and &#x60;MOVE&#x60; operations. | [optional] 

## Methods

### NewPatchOperation

`func NewPatchOperation(op string, path string, ) *PatchOperation`

NewPatchOperation instantiates a new PatchOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchOperationWithDefaults

`func NewPatchOperationWithDefaults() *PatchOperation`

NewPatchOperationWithDefaults instantiates a new PatchOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *PatchOperation) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *PatchOperation) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *PatchOperation) SetOp(v string)`

SetOp sets Op field to given value.


### GetPath

`func (o *PatchOperation) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PatchOperation) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PatchOperation) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *PatchOperation) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PatchOperation) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PatchOperation) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *PatchOperation) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetFrom

`func (o *PatchOperation) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *PatchOperation) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *PatchOperation) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *PatchOperation) HasFrom() bool`

HasFrom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


