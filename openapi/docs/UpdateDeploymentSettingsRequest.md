# UpdateDeploymentSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultQueryVi** | Pointer to **string** | RRN of the Virtual Instance that all queries will be routed to by default | [optional] 

## Methods

### NewUpdateDeploymentSettingsRequest

`func NewUpdateDeploymentSettingsRequest() *UpdateDeploymentSettingsRequest`

NewUpdateDeploymentSettingsRequest instantiates a new UpdateDeploymentSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeploymentSettingsRequestWithDefaults

`func NewUpdateDeploymentSettingsRequestWithDefaults() *UpdateDeploymentSettingsRequest`

NewUpdateDeploymentSettingsRequestWithDefaults instantiates a new UpdateDeploymentSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultQueryVi

`func (o *UpdateDeploymentSettingsRequest) GetDefaultQueryVi() string`

GetDefaultQueryVi returns the DefaultQueryVi field if non-nil, zero value otherwise.

### GetDefaultQueryViOk

`func (o *UpdateDeploymentSettingsRequest) GetDefaultQueryViOk() (*string, bool)`

GetDefaultQueryViOk returns a tuple with the DefaultQueryVi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultQueryVi

`func (o *UpdateDeploymentSettingsRequest) SetDefaultQueryVi(v string)`

SetDefaultQueryVi sets DefaultQueryVi field to given value.

### HasDefaultQueryVi

`func (o *UpdateDeploymentSettingsRequest) HasDefaultQueryVi() bool`

HasDefaultQueryVi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


