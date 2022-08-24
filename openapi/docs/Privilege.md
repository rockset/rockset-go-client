# Privilege

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The action allowed by this privilege. | [optional] 
**ResourceName** | Pointer to **string** | The resources on which the action is allowed. Defaults to &#39;*All*&#39; if not specified. | [optional] 
**Cluster** | Pointer to **string** | Cluster ID (&#x60;usw2a1&#x60; for us-west-2, &#x60;use1a1&#x60; for us-east-1, &#x60;euc1a1&#x60; for eu-central-1) for which the action is allowed. Defaults to &#39;*All*&#39; if not specified. | [optional] 

## Methods

### NewPrivilege

`func NewPrivilege() *Privilege`

NewPrivilege instantiates a new Privilege object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivilegeWithDefaults

`func NewPrivilegeWithDefaults() *Privilege`

NewPrivilegeWithDefaults instantiates a new Privilege object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *Privilege) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Privilege) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Privilege) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *Privilege) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetResourceName

`func (o *Privilege) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *Privilege) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *Privilege) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *Privilege) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetCluster

`func (o *Privilege) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *Privilege) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *Privilege) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *Privilege) HasCluster() bool`

HasCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


