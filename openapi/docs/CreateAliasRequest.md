# CreateAliasRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Alias name | 
**Description** | Pointer to **string** | optional description | [optional] 
**Collections** | **[]string** | list of fully qualified collection names referenced by alias | 

## Methods

### NewCreateAliasRequest

`func NewCreateAliasRequest(name string, collections []string, ) *CreateAliasRequest`

NewCreateAliasRequest instantiates a new CreateAliasRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAliasRequestWithDefaults

`func NewCreateAliasRequestWithDefaults() *CreateAliasRequest`

NewCreateAliasRequestWithDefaults instantiates a new CreateAliasRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAliasRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAliasRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAliasRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateAliasRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAliasRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAliasRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAliasRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCollections

`func (o *CreateAliasRequest) GetCollections() []string`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *CreateAliasRequest) GetCollectionsOk() (*[]string, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *CreateAliasRequest) SetCollections(v []string)`

SetCollections sets Collections field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


