# OrgMembership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organization** | [**Organization**](Organization.md) |  | 
**Roles** | **[]string** |  | 
**InviteState** | **string** |  | 
**InvitedBy** | **string** |  | 
**Email** | Pointer to **string** | email of the user | [optional] 
**ExpiresAt** | **string** | ISO-8601 date | [readonly] 
**CreatedAt** | **string** | ISO-8601 date | 

## Methods

### NewOrgMembership

`func NewOrgMembership(organization Organization, roles []string, inviteState string, invitedBy string, expiresAt string, createdAt string, ) *OrgMembership`

NewOrgMembership instantiates a new OrgMembership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgMembershipWithDefaults

`func NewOrgMembershipWithDefaults() *OrgMembership`

NewOrgMembershipWithDefaults instantiates a new OrgMembership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganization

`func (o *OrgMembership) GetOrganization() Organization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *OrgMembership) GetOrganizationOk() (*Organization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *OrgMembership) SetOrganization(v Organization)`

SetOrganization sets Organization field to given value.


### GetRoles

`func (o *OrgMembership) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *OrgMembership) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *OrgMembership) SetRoles(v []string)`

SetRoles sets Roles field to given value.


### GetInviteState

`func (o *OrgMembership) GetInviteState() string`

GetInviteState returns the InviteState field if non-nil, zero value otherwise.

### GetInviteStateOk

`func (o *OrgMembership) GetInviteStateOk() (*string, bool)`

GetInviteStateOk returns a tuple with the InviteState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteState

`func (o *OrgMembership) SetInviteState(v string)`

SetInviteState sets InviteState field to given value.


### GetInvitedBy

`func (o *OrgMembership) GetInvitedBy() string`

GetInvitedBy returns the InvitedBy field if non-nil, zero value otherwise.

### GetInvitedByOk

`func (o *OrgMembership) GetInvitedByOk() (*string, bool)`

GetInvitedByOk returns a tuple with the InvitedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedBy

`func (o *OrgMembership) SetInvitedBy(v string)`

SetInvitedBy sets InvitedBy field to given value.


### GetEmail

`func (o *OrgMembership) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrgMembership) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrgMembership) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrgMembership) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExpiresAt

`func (o *OrgMembership) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *OrgMembership) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *OrgMembership) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### GetCreatedAt

`func (o *OrgMembership) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrgMembership) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrgMembership) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


