# Workspace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | ISO-8601 date of when workspace was created. | [optional] 
**CreatedBy** | Pointer to **string** | Email of user who created the workspace. | [optional] 
**Name** | Pointer to **string** | Descriptive label and unique identifier. | [optional] 
**Description** | Pointer to **string** | Longer explanation for the workspace. | [optional] 
**CollectionCount** | Pointer to **int64** | Number of collections that are immediate children of workspace. | [optional] 

## Methods

### NewWorkspace

`func NewWorkspace() *Workspace`

NewWorkspace instantiates a new Workspace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceWithDefaults

`func NewWorkspaceWithDefaults() *Workspace`

NewWorkspaceWithDefaults instantiates a new Workspace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Workspace) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Workspace) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Workspace) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Workspace) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Workspace) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Workspace) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Workspace) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Workspace) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetName

`func (o *Workspace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Workspace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Workspace) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Workspace) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Workspace) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Workspace) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Workspace) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Workspace) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCollectionCount

`func (o *Workspace) GetCollectionCount() int64`

GetCollectionCount returns the CollectionCount field if non-nil, zero value otherwise.

### GetCollectionCountOk

`func (o *Workspace) GetCollectionCountOk() (*int64, bool)`

GetCollectionCountOk returns a tuple with the CollectionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionCount

`func (o *Workspace) SetCollectionCount(v int64)`

SetCollectionCount sets CollectionCount field to given value.

### HasCollectionCount

`func (o *Workspace) HasCollectionCount() bool`

HasCollectionCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


