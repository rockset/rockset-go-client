# Alias

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **[]string** | List of fully qualified collection names referenced by alias. | [optional] 
**CreatedAt** | Pointer to **string** | ISO-8601 date. | [optional] 
**CreatedByApikeyName** | Pointer to **string** | Name of the API key that was used to create this object if one was used. | [optional] 
**CreatorEmail** | Pointer to **string** | Email of the creator. | [optional] 
**Description** | Pointer to **string** | Alias description. | [optional] 
**ModifiedAt** | Pointer to **string** | ISO-8601 date. | [optional] 
**Name** | Pointer to **string** | Name of the alias. | [optional] 
**State** | Pointer to **string** | State of the alias. | [optional] 
**Workspace** | Pointer to **string** | Name of the workspace. | [optional] 

## Methods

### NewAlias

`func NewAlias() *Alias`

NewAlias instantiates a new Alias object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAliasWithDefaults

`func NewAliasWithDefaults() *Alias`

NewAliasWithDefaults instantiates a new Alias object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *Alias) GetCollections() []string`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *Alias) GetCollectionsOk() (*[]string, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *Alias) SetCollections(v []string)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *Alias) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Alias) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Alias) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Alias) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Alias) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedByApikeyName

`func (o *Alias) GetCreatedByApikeyName() string`

GetCreatedByApikeyName returns the CreatedByApikeyName field if non-nil, zero value otherwise.

### GetCreatedByApikeyNameOk

`func (o *Alias) GetCreatedByApikeyNameOk() (*string, bool)`

GetCreatedByApikeyNameOk returns a tuple with the CreatedByApikeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByApikeyName

`func (o *Alias) SetCreatedByApikeyName(v string)`

SetCreatedByApikeyName sets CreatedByApikeyName field to given value.

### HasCreatedByApikeyName

`func (o *Alias) HasCreatedByApikeyName() bool`

HasCreatedByApikeyName returns a boolean if a field has been set.

### GetCreatorEmail

`func (o *Alias) GetCreatorEmail() string`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *Alias) GetCreatorEmailOk() (*string, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *Alias) SetCreatorEmail(v string)`

SetCreatorEmail sets CreatorEmail field to given value.

### HasCreatorEmail

`func (o *Alias) HasCreatorEmail() bool`

HasCreatorEmail returns a boolean if a field has been set.

### GetDescription

`func (o *Alias) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Alias) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Alias) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Alias) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetModifiedAt

`func (o *Alias) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Alias) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Alias) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *Alias) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *Alias) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Alias) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Alias) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Alias) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *Alias) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Alias) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Alias) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Alias) HasState() bool`

HasState returns a boolean if a field has been set.

### GetWorkspace

`func (o *Alias) GetWorkspace() string`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *Alias) GetWorkspaceOk() (*string, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *Alias) SetWorkspace(v string)`

SetWorkspace sets Workspace field to given value.

### HasWorkspace

`func (o *Alias) HasWorkspace() bool`

HasWorkspace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


