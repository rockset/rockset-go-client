# Offsets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commit** | Pointer to **string** | An encoded value representing the most recent offsets that have been committed. If the fence offset is less than or equal to the commit, then passed is true. | [optional] 

## Methods

### NewOffsets

`func NewOffsets() *Offsets`

NewOffsets instantiates a new Offsets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOffsetsWithDefaults

`func NewOffsetsWithDefaults() *Offsets`

NewOffsetsWithDefaults instantiates a new Offsets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommit

`func (o *Offsets) GetCommit() string`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *Offsets) GetCommitOk() (*string, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *Offsets) SetCommit(v string)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *Offsets) HasCommit() bool`

HasCommit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


