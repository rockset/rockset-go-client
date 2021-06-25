# CreateIpAllowlistRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | IP Allowlist policy name | 
**Description** | Pointer to **string** | optional description | [optional] 
**IpAddress** | **string** | individual IP address or range of IP addresses in CIDR notation | 

## Methods

### NewCreateIpAllowlistRequest

`func NewCreateIpAllowlistRequest(name string, ipAddress string, ) *CreateIpAllowlistRequest`

NewCreateIpAllowlistRequest instantiates a new CreateIpAllowlistRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIpAllowlistRequestWithDefaults

`func NewCreateIpAllowlistRequestWithDefaults() *CreateIpAllowlistRequest`

NewCreateIpAllowlistRequestWithDefaults instantiates a new CreateIpAllowlistRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateIpAllowlistRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateIpAllowlistRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateIpAllowlistRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateIpAllowlistRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateIpAllowlistRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateIpAllowlistRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateIpAllowlistRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpAddress

`func (o *CreateIpAllowlistRequest) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *CreateIpAllowlistRequest) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *CreateIpAllowlistRequest) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


