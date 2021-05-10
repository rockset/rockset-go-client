# ListIpAllowlistsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]IpAllowlist**](IpAllowlist.md) | list of all IP Allowlist network policies | [optional] 

## Methods

### NewListIpAllowlistsResponse

`func NewListIpAllowlistsResponse() *ListIpAllowlistsResponse`

NewListIpAllowlistsResponse instantiates a new ListIpAllowlistsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIpAllowlistsResponseWithDefaults

`func NewListIpAllowlistsResponseWithDefaults() *ListIpAllowlistsResponse`

NewListIpAllowlistsResponseWithDefaults instantiates a new ListIpAllowlistsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListIpAllowlistsResponse) GetData() []IpAllowlist`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListIpAllowlistsResponse) GetDataOk() (*[]IpAllowlist, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListIpAllowlistsResponse) SetData(v []IpAllowlist)`

SetData sets Data field to given value.

### HasData

`func (o *ListIpAllowlistsResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


