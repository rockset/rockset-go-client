# PatchDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique id of the document which is being patched | 
**Patch** | [**[]PatchOperation**](PatchOperation.md) | List of patch operations | 

## Methods

### NewPatchDocument

`func NewPatchDocument(id string, patch []PatchOperation, ) *PatchDocument`

NewPatchDocument instantiates a new PatchDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchDocumentWithDefaults

`func NewPatchDocumentWithDefaults() *PatchDocument`

NewPatchDocumentWithDefaults instantiates a new PatchDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchDocument) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchDocument) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchDocument) SetId(v string)`

SetId sets Id field to given value.


### GetPatch

`func (o *PatchDocument) GetPatch() []PatchOperation`

GetPatch returns the Patch field if non-nil, zero value otherwise.

### GetPatchOk

`func (o *PatchDocument) GetPatchOk() (*[]PatchOperation, bool)`

GetPatchOk returns a tuple with the Patch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatch

`func (o *PatchDocument) SetPatch(v []PatchOperation)`

SetPatch sets Patch field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


