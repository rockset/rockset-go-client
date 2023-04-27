# View

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | ISO-8601 date. | [optional] 
**CreatorEmail** | Pointer to **string** | Email of the creator. | [optional] 
**Description** | Pointer to **string** | View description. | [optional] 
**Entities** | Pointer to **[]string** | List of entities referenced by view. An entity can be a view, alias or collection. | [optional] 
**ModifiedAt** | Pointer to **string** | ISO-8601 date. | [optional] 
**Name** | Pointer to **string** | Name of the view. | [optional] 
**OwnerEmail** | Pointer to **string** | Email of the owner, note: deprecated and will always be null. | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**QuerySql** | Pointer to **string** | SQL query of the view. | [optional] 
**State** | Pointer to **string** | State of the view. | [optional] 
**Workspace** | Pointer to **string** | Name of the workspace. | [optional] 

## Methods

### NewView

`func NewView() *View`

NewView instantiates a new View object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewWithDefaults

`func NewViewWithDefaults() *View`

NewViewWithDefaults instantiates a new View object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *View) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *View) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *View) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *View) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatorEmail

`func (o *View) GetCreatorEmail() string`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *View) GetCreatorEmailOk() (*string, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *View) SetCreatorEmail(v string)`

SetCreatorEmail sets CreatorEmail field to given value.

### HasCreatorEmail

`func (o *View) HasCreatorEmail() bool`

HasCreatorEmail returns a boolean if a field has been set.

### GetDescription

`func (o *View) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *View) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *View) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *View) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntities

`func (o *View) GetEntities() []string`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *View) GetEntitiesOk() (*[]string, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *View) SetEntities(v []string)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *View) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetModifiedAt

`func (o *View) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *View) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *View) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *View) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetName

`func (o *View) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *View) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *View) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *View) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwnerEmail

`func (o *View) GetOwnerEmail() string`

GetOwnerEmail returns the OwnerEmail field if non-nil, zero value otherwise.

### GetOwnerEmailOk

`func (o *View) GetOwnerEmailOk() (*string, bool)`

GetOwnerEmailOk returns a tuple with the OwnerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerEmail

`func (o *View) SetOwnerEmail(v string)`

SetOwnerEmail sets OwnerEmail field to given value.

### HasOwnerEmail

`func (o *View) HasOwnerEmail() bool`

HasOwnerEmail returns a boolean if a field has been set.

### GetPath

`func (o *View) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *View) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *View) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *View) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetQuerySql

`func (o *View) GetQuerySql() string`

GetQuerySql returns the QuerySql field if non-nil, zero value otherwise.

### GetQuerySqlOk

`func (o *View) GetQuerySqlOk() (*string, bool)`

GetQuerySqlOk returns a tuple with the QuerySql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerySql

`func (o *View) SetQuerySql(v string)`

SetQuerySql sets QuerySql field to given value.

### HasQuerySql

`func (o *View) HasQuerySql() bool`

HasQuerySql returns a boolean if a field has been set.

### GetState

`func (o *View) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *View) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *View) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *View) HasState() bool`

HasState returns a boolean if a field has been set.

### GetWorkspace

`func (o *View) GetWorkspace() string`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *View) GetWorkspaceOk() (*string, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *View) SetWorkspace(v string)`

SetWorkspace sets Workspace field to given value.

### HasWorkspace

`func (o *View) HasWorkspace() bool`

HasWorkspace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


