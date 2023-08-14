# GetCollectionCommitData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fence** | Pointer to **string** | The max offset passed in the request which is the latest offset that was compared to the commit. | [optional] 
**Passed** | Pointer to **bool** | True if all the data in the offsets are available to be queried. False if one or more of the offsets is still behind the request offsets. | [optional] 

## Methods

### NewGetCollectionCommitData

`func NewGetCollectionCommitData() *GetCollectionCommitData`

NewGetCollectionCommitData instantiates a new GetCollectionCommitData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCollectionCommitDataWithDefaults

`func NewGetCollectionCommitDataWithDefaults() *GetCollectionCommitData`

NewGetCollectionCommitDataWithDefaults instantiates a new GetCollectionCommitData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFence

`func (o *GetCollectionCommitData) GetFence() string`

GetFence returns the Fence field if non-nil, zero value otherwise.

### GetFenceOk

`func (o *GetCollectionCommitData) GetFenceOk() (*string, bool)`

GetFenceOk returns a tuple with the Fence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFence

`func (o *GetCollectionCommitData) SetFence(v string)`

SetFence sets Fence field to given value.

### HasFence

`func (o *GetCollectionCommitData) HasFence() bool`

HasFence returns a boolean if a field has been set.

### GetPassed

`func (o *GetCollectionCommitData) GetPassed() bool`

GetPassed returns the Passed field if non-nil, zero value otherwise.

### GetPassedOk

`func (o *GetCollectionCommitData) GetPassedOk() (*bool, bool)`

GetPassedOk returns a tuple with the Passed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassed

`func (o *GetCollectionCommitData) SetPassed(v bool)`

SetPassed sets Passed field to given value.

### HasPassed

`func (o *GetCollectionCommitData) HasPassed() bool`

HasPassed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


