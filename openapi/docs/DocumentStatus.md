# DocumentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | Pointer to **string** | collection name | [optional] 
**Error** | Pointer to [**ErrorModel**](ErrorModel.md) |  | [optional] 
**Id** | Pointer to **string** | unique document ID | [optional] 
**Status** | Pointer to **string** | status of the document | [optional] 
**PatchId** | Pointer to **string** | unique id used to represent each patch request | [optional] 

## Methods

### NewDocumentStatus

`func NewDocumentStatus() *DocumentStatus`

NewDocumentStatus instantiates a new DocumentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentStatusWithDefaults

`func NewDocumentStatusWithDefaults() *DocumentStatus`

NewDocumentStatusWithDefaults instantiates a new DocumentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *DocumentStatus) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *DocumentStatus) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *DocumentStatus) SetCollection(v string)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *DocumentStatus) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetError

`func (o *DocumentStatus) GetError() ErrorModel`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DocumentStatus) GetErrorOk() (*ErrorModel, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DocumentStatus) SetError(v ErrorModel)`

SetError sets Error field to given value.

### HasError

`func (o *DocumentStatus) HasError() bool`

HasError returns a boolean if a field has been set.

### GetId

`func (o *DocumentStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DocumentStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *DocumentStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DocumentStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DocumentStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DocumentStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPatchId

`func (o *DocumentStatus) GetPatchId() string`

GetPatchId returns the PatchId field if non-nil, zero value otherwise.

### GetPatchIdOk

`func (o *DocumentStatus) GetPatchIdOk() (*string, bool)`

GetPatchIdOk returns a tuple with the PatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchId

`func (o *DocumentStatus) SetPatchId(v string)`

SetPatchId sets PatchId field to given value.

### HasPatchId

`func (o *DocumentStatus) HasPatchId() bool`

HasPatchId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


