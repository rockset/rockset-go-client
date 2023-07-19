# SourceSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceCollectionPath** | Pointer to **string** | Path of source collection to restore the snapshot from. | [optional] 
**SourceSnapshotId** | Pointer to **string** | Snapshot id of the snapshot that the new collection will be created from. | [optional] 

## Methods

### NewSourceSnapshot

`func NewSourceSnapshot() *SourceSnapshot`

NewSourceSnapshot instantiates a new SourceSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceSnapshotWithDefaults

`func NewSourceSnapshotWithDefaults() *SourceSnapshot`

NewSourceSnapshotWithDefaults instantiates a new SourceSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceCollectionPath

`func (o *SourceSnapshot) GetSourceCollectionPath() string`

GetSourceCollectionPath returns the SourceCollectionPath field if non-nil, zero value otherwise.

### GetSourceCollectionPathOk

`func (o *SourceSnapshot) GetSourceCollectionPathOk() (*string, bool)`

GetSourceCollectionPathOk returns a tuple with the SourceCollectionPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCollectionPath

`func (o *SourceSnapshot) SetSourceCollectionPath(v string)`

SetSourceCollectionPath sets SourceCollectionPath field to given value.

### HasSourceCollectionPath

`func (o *SourceSnapshot) HasSourceCollectionPath() bool`

HasSourceCollectionPath returns a boolean if a field has been set.

### GetSourceSnapshotId

`func (o *SourceSnapshot) GetSourceSnapshotId() string`

GetSourceSnapshotId returns the SourceSnapshotId field if non-nil, zero value otherwise.

### GetSourceSnapshotIdOk

`func (o *SourceSnapshot) GetSourceSnapshotIdOk() (*string, bool)`

GetSourceSnapshotIdOk returns a tuple with the SourceSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSnapshotId

`func (o *SourceSnapshot) SetSourceSnapshotId(v string)`

SetSourceSnapshotId sets SourceSnapshotId field to given value.

### HasSourceSnapshotId

`func (o *SourceSnapshot) HasSourceSnapshotId() bool`

HasSourceSnapshotId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


