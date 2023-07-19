# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusters** | Pointer to [**[]Cluster**](Cluster.md) | List of clusters associated with this org. | [optional] 
**CreatedAt** | Pointer to **string** | ISO-8601 date. | [optional] 
**DisplayName** | Pointer to **string** | Name of the organization. | [optional] 
**ExternalId** | Pointer to **string** | Organization&#39;s unique external ID within Rockset. | [optional] 
**Id** | Pointer to **string** | Unique identifier for the organization. | [optional] 
**RocksetUser** | Pointer to **string** | Rockset&#39;s global AWS user. | [optional] 
**SsoConnection** | Pointer to **string** | Connection name of SSO connection. | [optional] 
**SsoOnly** | Pointer to **bool** | Whether or not SSO is the only permitted form of auth. | [optional] 

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

### GetClusters

`func (o *Organization) GetClusters() []Cluster`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *Organization) GetClustersOk() (*[]Cluster, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *Organization) SetClusters(v []Cluster)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *Organization) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

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

### GetExternalId

`func (o *Organization) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Organization) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Organization) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Organization) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

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

### GetRocksetUser

`func (o *Organization) GetRocksetUser() string`

GetRocksetUser returns the RocksetUser field if non-nil, zero value otherwise.

### GetRocksetUserOk

`func (o *Organization) GetRocksetUserOk() (*string, bool)`

GetRocksetUserOk returns a tuple with the RocksetUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRocksetUser

`func (o *Organization) SetRocksetUser(v string)`

SetRocksetUser sets RocksetUser field to given value.

### HasRocksetUser

`func (o *Organization) HasRocksetUser() bool`

HasRocksetUser returns a boolean if a field has been set.

### GetSsoConnection

`func (o *Organization) GetSsoConnection() string`

GetSsoConnection returns the SsoConnection field if non-nil, zero value otherwise.

### GetSsoConnectionOk

`func (o *Organization) GetSsoConnectionOk() (*string, bool)`

GetSsoConnectionOk returns a tuple with the SsoConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoConnection

`func (o *Organization) SetSsoConnection(v string)`

SetSsoConnection sets SsoConnection field to given value.

### HasSsoConnection

`func (o *Organization) HasSsoConnection() bool`

HasSsoConnection returns a boolean if a field has been set.

### GetSsoOnly

`func (o *Organization) GetSsoOnly() bool`

GetSsoOnly returns the SsoOnly field if non-nil, zero value otherwise.

### GetSsoOnlyOk

`func (o *Organization) GetSsoOnlyOk() (*bool, bool)`

GetSsoOnlyOk returns a tuple with the SsoOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoOnly

`func (o *Organization) SetSsoOnly(v bool)`

SetSsoOnly sets SsoOnly field to given value.

### HasSsoOnly

`func (o *Organization) HasSsoOnly() bool`

HasSsoOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


