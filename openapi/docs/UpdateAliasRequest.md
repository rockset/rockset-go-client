# UpdateAliasRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | **[]string** | List of fully qualified collection names referenced by alias. | 
**Description** | Pointer to **string** | Optional description. | [optional] 

## Methods

### NewUpdateAliasRequest

`func NewUpdateAliasRequest(collections []string, ) *UpdateAliasRequest`

NewUpdateAliasRequest instantiates a new UpdateAliasRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAliasRequestWithDefaults

`func NewUpdateAliasRequestWithDefaults() *UpdateAliasRequest`

NewUpdateAliasRequestWithDefaults instantiates a new UpdateAliasRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *UpdateAliasRequest) GetCollections() []string`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *UpdateAliasRequest) GetCollectionsOk() (*[]string, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *UpdateAliasRequest) SetCollections(v []string)`

SetCollections sets Collections field to given value.


### GetDescription

`func (o *UpdateAliasRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateAliasRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateAliasRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateAliasRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


