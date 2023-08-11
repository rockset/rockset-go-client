# GetCollectionCommit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**GetCollectionCommitData**](GetCollectionCommitData.md) |  | [optional] 
**Offsets** | Pointer to [**Offsets**](Offsets.md) |  | [optional] 

## Methods

### NewGetCollectionCommit

`func NewGetCollectionCommit() *GetCollectionCommit`

NewGetCollectionCommit instantiates a new GetCollectionCommit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCollectionCommitWithDefaults

`func NewGetCollectionCommitWithDefaults() *GetCollectionCommit`

NewGetCollectionCommitWithDefaults instantiates a new GetCollectionCommit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetCollectionCommit) GetData() GetCollectionCommitData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCollectionCommit) GetDataOk() (*GetCollectionCommitData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCollectionCommit) SetData(v GetCollectionCommitData)`

SetData sets Data field to given value.

### HasData

`func (o *GetCollectionCommit) HasData() bool`

HasData returns a boolean if a field has been set.

### GetOffsets

`func (o *GetCollectionCommit) GetOffsets() Offsets`

GetOffsets returns the Offsets field if non-nil, zero value otherwise.

### GetOffsetsOk

`func (o *GetCollectionCommit) GetOffsetsOk() (*Offsets, bool)`

GetOffsetsOk returns a tuple with the Offsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsets

`func (o *GetCollectionCommit) SetOffsets(v Offsets)`

SetOffsets sets Offsets field to given value.

### HasOffsets

`func (o *GetCollectionCommit) HasOffsets() bool`

HasOffsets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


