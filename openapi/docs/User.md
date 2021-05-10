# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | ISO-8601 date | [optional] 
**Email** | **string** | user email | 
**FirstName** | Pointer to **string** | user first name | [optional] 
**LastName** | Pointer to **string** | user last name | [optional] 
**Roles** | Pointer to **[]string** | List of roles for a given user | [optional] 
**State** | Pointer to **string** | state of user - NEW / ACTIVE | [optional] 
**Org** | Pointer to **string** |  | [optional] 
**InviteState** | Pointer to **string** |  | [optional] 
**Orgs** | Pointer to [**[]Organization**](Organization.md) |  | [optional] 
**OrgMemberships** | Pointer to [**[]OrgMembership**](OrgMembership.md) |  | [optional] 

## Methods

### NewUser

`func NewUser(email string, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *User) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *User) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *User) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *User) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *User) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *User) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *User) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *User) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *User) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *User) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *User) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *User) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetRoles

`func (o *User) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *User) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *User) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *User) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetState

`func (o *User) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *User) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *User) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *User) HasState() bool`

HasState returns a boolean if a field has been set.

### GetOrg

`func (o *User) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *User) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *User) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *User) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetInviteState

`func (o *User) GetInviteState() string`

GetInviteState returns the InviteState field if non-nil, zero value otherwise.

### GetInviteStateOk

`func (o *User) GetInviteStateOk() (*string, bool)`

GetInviteStateOk returns a tuple with the InviteState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteState

`func (o *User) SetInviteState(v string)`

SetInviteState sets InviteState field to given value.

### HasInviteState

`func (o *User) HasInviteState() bool`

HasInviteState returns a boolean if a field has been set.

### GetOrgs

`func (o *User) GetOrgs() []Organization`

GetOrgs returns the Orgs field if non-nil, zero value otherwise.

### GetOrgsOk

`func (o *User) GetOrgsOk() (*[]Organization, bool)`

GetOrgsOk returns a tuple with the Orgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgs

`func (o *User) SetOrgs(v []Organization)`

SetOrgs sets Orgs field to given value.

### HasOrgs

`func (o *User) HasOrgs() bool`

HasOrgs returns a boolean if a field has been set.

### GetOrgMemberships

`func (o *User) GetOrgMemberships() []OrgMembership`

GetOrgMemberships returns the OrgMemberships field if non-nil, zero value otherwise.

### GetOrgMembershipsOk

`func (o *User) GetOrgMembershipsOk() (*[]OrgMembership, bool)`

GetOrgMembershipsOk returns a tuple with the OrgMemberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgMemberships

`func (o *User) SetOrgMemberships(v []OrgMembership)`

SetOrgMemberships sets OrgMemberships field to given value.

### HasOrgMemberships

`func (o *User) HasOrgMemberships() bool`

HasOrgMemberships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


