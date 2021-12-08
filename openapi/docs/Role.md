# Role

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleName** | Pointer to **string** | Unique identifier for the role. | [optional] 
**Description** | Pointer to **string** | Description for the role. | [optional] 
**OwnerEmail** | Pointer to **string** | Email of the user who currently owns the role. | [optional] 
**CreatedBy** | Pointer to **string** | Email of the user who created the role. | [optional] 
**Privileges** | Pointer to [**[]Privilege**](Privilege.md) | List of privileges associated with the role. | [optional] 
**CreatedAt** | Pointer to **string** | ISO-8601 date of when the role was created. | [optional] 

## Methods

### NewRole

`func NewRole() *Role`

NewRole instantiates a new Role object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleWithDefaults

`func NewRoleWithDefaults() *Role`

NewRoleWithDefaults instantiates a new Role object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleName

`func (o *Role) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *Role) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *Role) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *Role) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetDescription

`func (o *Role) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Role) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Role) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Role) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwnerEmail

`func (o *Role) GetOwnerEmail() string`

GetOwnerEmail returns the OwnerEmail field if non-nil, zero value otherwise.

### GetOwnerEmailOk

`func (o *Role) GetOwnerEmailOk() (*string, bool)`

GetOwnerEmailOk returns a tuple with the OwnerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerEmail

`func (o *Role) SetOwnerEmail(v string)`

SetOwnerEmail sets OwnerEmail field to given value.

### HasOwnerEmail

`func (o *Role) HasOwnerEmail() bool`

HasOwnerEmail returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Role) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Role) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Role) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Role) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetPrivileges

`func (o *Role) GetPrivileges() []Privilege`

GetPrivileges returns the Privileges field if non-nil, zero value otherwise.

### GetPrivilegesOk

`func (o *Role) GetPrivilegesOk() (*[]Privilege, bool)`

GetPrivilegesOk returns a tuple with the Privileges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivileges

`func (o *Role) SetPrivileges(v []Privilege)`

SetPrivileges sets Privileges field to given value.

### HasPrivileges

`func (o *Role) HasPrivileges() bool`

HasPrivileges returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Role) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Role) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Role) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Role) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


