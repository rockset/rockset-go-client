# IpAllowlist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | IP Allowlist policy name | 
**Description** | Pointer to **string** | optional description | [optional] 
**CreatorEmail** | Pointer to **string** | email of the creator | [optional] 
**IpAddress** | **string** | individual IP address or range of IP addresses in CIDR notation | 
**CreatedAt** | Pointer to **string** | ISO-8601 date | [optional] 

## Methods

### NewIpAllowlist

`func NewIpAllowlist(name string, ipAddress string, ) *IpAllowlist`

NewIpAllowlist instantiates a new IpAllowlist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpAllowlistWithDefaults

`func NewIpAllowlistWithDefaults() *IpAllowlist`

NewIpAllowlistWithDefaults instantiates a new IpAllowlist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IpAllowlist) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IpAllowlist) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IpAllowlist) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *IpAllowlist) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IpAllowlist) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IpAllowlist) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IpAllowlist) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatorEmail

`func (o *IpAllowlist) GetCreatorEmail() string`

GetCreatorEmail returns the CreatorEmail field if non-nil, zero value otherwise.

### GetCreatorEmailOk

`func (o *IpAllowlist) GetCreatorEmailOk() (*string, bool)`

GetCreatorEmailOk returns a tuple with the CreatorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorEmail

`func (o *IpAllowlist) SetCreatorEmail(v string)`

SetCreatorEmail sets CreatorEmail field to given value.

### HasCreatorEmail

`func (o *IpAllowlist) HasCreatorEmail() bool`

HasCreatorEmail returns a boolean if a field has been set.

### GetIpAddress

`func (o *IpAllowlist) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *IpAllowlist) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *IpAllowlist) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetCreatedAt

`func (o *IpAllowlist) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IpAllowlist) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IpAllowlist) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IpAllowlist) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


