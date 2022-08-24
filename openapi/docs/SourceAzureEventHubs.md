# SourceAzureEventHubs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HubId** | Pointer to **string** | Name of the hub which rockset should ingest from. | [optional] 
**OffsetResetPolicy** | Pointer to **string** | The offset reset policy. | [optional] 
**Status** | Pointer to [**StatusAzureEventHubs**](StatusAzureEventHubs.md) |  | [optional] 

## Methods

### NewSourceAzureEventHubs

`func NewSourceAzureEventHubs() *SourceAzureEventHubs`

NewSourceAzureEventHubs instantiates a new SourceAzureEventHubs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceAzureEventHubsWithDefaults

`func NewSourceAzureEventHubsWithDefaults() *SourceAzureEventHubs`

NewSourceAzureEventHubsWithDefaults instantiates a new SourceAzureEventHubs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHubId

`func (o *SourceAzureEventHubs) GetHubId() string`

GetHubId returns the HubId field if non-nil, zero value otherwise.

### GetHubIdOk

`func (o *SourceAzureEventHubs) GetHubIdOk() (*string, bool)`

GetHubIdOk returns a tuple with the HubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubId

`func (o *SourceAzureEventHubs) SetHubId(v string)`

SetHubId sets HubId field to given value.

### HasHubId

`func (o *SourceAzureEventHubs) HasHubId() bool`

HasHubId returns a boolean if a field has been set.

### GetOffsetResetPolicy

`func (o *SourceAzureEventHubs) GetOffsetResetPolicy() string`

GetOffsetResetPolicy returns the OffsetResetPolicy field if non-nil, zero value otherwise.

### GetOffsetResetPolicyOk

`func (o *SourceAzureEventHubs) GetOffsetResetPolicyOk() (*string, bool)`

GetOffsetResetPolicyOk returns a tuple with the OffsetResetPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetResetPolicy

`func (o *SourceAzureEventHubs) SetOffsetResetPolicy(v string)`

SetOffsetResetPolicy sets OffsetResetPolicy field to given value.

### HasOffsetResetPolicy

`func (o *SourceAzureEventHubs) HasOffsetResetPolicy() bool`

HasOffsetResetPolicy returns a boolean if a field has been set.

### GetStatus

`func (o *SourceAzureEventHubs) GetStatus() StatusAzureEventHubs`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SourceAzureEventHubs) GetStatusOk() (*StatusAzureEventHubs, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SourceAzureEventHubs) SetStatus(v StatusAzureEventHubs)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SourceAzureEventHubs) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


