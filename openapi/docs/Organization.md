# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletionScheduledAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | unique identifier for the organization | [optional] 
**CreatedAt** | Pointer to **string** | ISO-8601 date | [optional] 
**DisplayName** | Pointer to **string** | name of the organization | [optional] 

## Methods

### NewOrganization

`func NewOrganization() *Organization`

NewOrganization instantiates a new Organization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWithDefaults

`func NewOrganizationWithDefaults() *Organization`

NewOrganizationWithDefaults instantiates a new Organization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletionScheduledAt

`func (o *Organization) GetDeletionScheduledAt() string`

GetDeletionScheduledAt returns the DeletionScheduledAt field if non-nil, zero value otherwise.

### GetDeletionScheduledAtOk

`func (o *Organization) GetDeletionScheduledAtOk() (*string, bool)`

GetDeletionScheduledAtOk returns a tuple with the DeletionScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionScheduledAt

`func (o *Organization) SetDeletionScheduledAt(v string)`

SetDeletionScheduledAt sets DeletionScheduledAt field to given value.

### HasDeletionScheduledAt

`func (o *Organization) HasDeletionScheduledAt() bool`

HasDeletionScheduledAt returns a boolean if a field has been set.

### GetId

`func (o *Organization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Organization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Organization) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Organization) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Organization) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Organization) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Organization) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Organization) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisplayName

`func (o *Organization) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Organization) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Organization) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Organization) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


