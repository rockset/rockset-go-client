# ValidateQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **[]string** | list of collection specified in query | 
**Collections** | **[]string** | list of collection specified in query | 
**Parameters** | **[]string** | list of parameters specified in query | 

## Methods

### NewValidateQueryResponse

`func NewValidateQueryResponse(name []string, collections []string, parameters []string, ) *ValidateQueryResponse`

NewValidateQueryResponse instantiates a new ValidateQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateQueryResponseWithDefaults

`func NewValidateQueryResponseWithDefaults() *ValidateQueryResponse`

NewValidateQueryResponseWithDefaults instantiates a new ValidateQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ValidateQueryResponse) GetName() []string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ValidateQueryResponse) GetNameOk() (*[]string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ValidateQueryResponse) SetName(v []string)`

SetName sets Name field to given value.


### GetCollections

`func (o *ValidateQueryResponse) GetCollections() []string`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *ValidateQueryResponse) GetCollectionsOk() (*[]string, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *ValidateQueryResponse) SetCollections(v []string)`

SetCollections sets Collections field to given value.


### GetParameters

`func (o *ValidateQueryResponse) GetParameters() []string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ValidateQueryResponse) GetParametersOk() (*[]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ValidateQueryResponse) SetParameters(v []string)`

SetParameters sets Parameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


