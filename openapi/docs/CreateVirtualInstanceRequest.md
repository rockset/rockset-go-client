# CreateVirtualInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Requested virtual instance type. | [optional] 
**Name** | **string** | Unique identifier for virtual instance, can contain alphanumeric or dash characters. | 
**Description** | Pointer to **string** | Description of requested virtual instance. | [optional] 
**AutoSuspendSeconds** | Pointer to **int32** | Number of seconds without queries after which the VI is suspended | [optional] 

## Methods

### NewCreateVirtualInstanceRequest

`func NewCreateVirtualInstanceRequest(name string, ) *CreateVirtualInstanceRequest`

NewCreateVirtualInstanceRequest instantiates a new CreateVirtualInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVirtualInstanceRequestWithDefaults

`func NewCreateVirtualInstanceRequestWithDefaults() *CreateVirtualInstanceRequest`

NewCreateVirtualInstanceRequestWithDefaults instantiates a new CreateVirtualInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateVirtualInstanceRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateVirtualInstanceRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateVirtualInstanceRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateVirtualInstanceRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *CreateVirtualInstanceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateVirtualInstanceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateVirtualInstanceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateVirtualInstanceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateVirtualInstanceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateVirtualInstanceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateVirtualInstanceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAutoSuspendSeconds

`func (o *CreateVirtualInstanceRequest) GetAutoSuspendSeconds() int32`

GetAutoSuspendSeconds returns the AutoSuspendSeconds field if non-nil, zero value otherwise.

### GetAutoSuspendSecondsOk

`func (o *CreateVirtualInstanceRequest) GetAutoSuspendSecondsOk() (*int32, bool)`

GetAutoSuspendSecondsOk returns a tuple with the AutoSuspendSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSuspendSeconds

`func (o *CreateVirtualInstanceRequest) SetAutoSuspendSeconds(v int32)`

SetAutoSuspendSeconds sets AutoSuspendSeconds field to given value.

### HasAutoSuspendSeconds

`func (o *CreateVirtualInstanceRequest) HasAutoSuspendSeconds() bool`

HasAutoSuspendSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


