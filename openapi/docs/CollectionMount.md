# CollectionMount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionPath** | Pointer to **string** | Collection path. | [optional] 
**CreatedAt** | Pointer to **string** | ISO-8601 date. | [optional] 
**Id** | Pointer to **string** | Mount ID. | [optional] 
**LastRefreshTimeMillis** | Pointer to **int64** | Unix timestamp of most recent refresh. Not applicable for live mounts. | [optional] 
**Rrn** | Pointer to **string** | Mount RRN. | [optional] 
**SnapshotExpirationTimeMillis** | Pointer to **int64** | Time in millis at which the snapshot expires. | [optional] 
**State** | Pointer to **string** | Mount type. | [optional] 
**Stats** | Pointer to [**CollectionMountStats**](CollectionMountStats.md) |  | [optional] 
**Type** | Pointer to **string** | Mount type. | [optional] 
**VirtualInstanceId** | Pointer to **string** | Virtual instance ID. | [optional] 
**VirtualInstanceRrn** | Pointer to **string** | Virtual Instance RRN. | [optional] 

## Methods

### NewCollectionMount

`func NewCollectionMount() *CollectionMount`

NewCollectionMount instantiates a new CollectionMount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionMountWithDefaults

`func NewCollectionMountWithDefaults() *CollectionMount`

NewCollectionMountWithDefaults instantiates a new CollectionMount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionPath

`func (o *CollectionMount) GetCollectionPath() string`

GetCollectionPath returns the CollectionPath field if non-nil, zero value otherwise.

### GetCollectionPathOk

`func (o *CollectionMount) GetCollectionPathOk() (*string, bool)`

GetCollectionPathOk returns a tuple with the CollectionPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionPath

`func (o *CollectionMount) SetCollectionPath(v string)`

SetCollectionPath sets CollectionPath field to given value.

### HasCollectionPath

`func (o *CollectionMount) HasCollectionPath() bool`

HasCollectionPath returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CollectionMount) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CollectionMount) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CollectionMount) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CollectionMount) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *CollectionMount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CollectionMount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CollectionMount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CollectionMount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastRefreshTimeMillis

`func (o *CollectionMount) GetLastRefreshTimeMillis() int64`

GetLastRefreshTimeMillis returns the LastRefreshTimeMillis field if non-nil, zero value otherwise.

### GetLastRefreshTimeMillisOk

`func (o *CollectionMount) GetLastRefreshTimeMillisOk() (*int64, bool)`

GetLastRefreshTimeMillisOk returns a tuple with the LastRefreshTimeMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRefreshTimeMillis

`func (o *CollectionMount) SetLastRefreshTimeMillis(v int64)`

SetLastRefreshTimeMillis sets LastRefreshTimeMillis field to given value.

### HasLastRefreshTimeMillis

`func (o *CollectionMount) HasLastRefreshTimeMillis() bool`

HasLastRefreshTimeMillis returns a boolean if a field has been set.

### GetRrn

`func (o *CollectionMount) GetRrn() string`

GetRrn returns the Rrn field if non-nil, zero value otherwise.

### GetRrnOk

`func (o *CollectionMount) GetRrnOk() (*string, bool)`

GetRrnOk returns a tuple with the Rrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrn

`func (o *CollectionMount) SetRrn(v string)`

SetRrn sets Rrn field to given value.

### HasRrn

`func (o *CollectionMount) HasRrn() bool`

HasRrn returns a boolean if a field has been set.

### GetSnapshotExpirationTimeMillis

`func (o *CollectionMount) GetSnapshotExpirationTimeMillis() int64`

GetSnapshotExpirationTimeMillis returns the SnapshotExpirationTimeMillis field if non-nil, zero value otherwise.

### GetSnapshotExpirationTimeMillisOk

`func (o *CollectionMount) GetSnapshotExpirationTimeMillisOk() (*int64, bool)`

GetSnapshotExpirationTimeMillisOk returns a tuple with the SnapshotExpirationTimeMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotExpirationTimeMillis

`func (o *CollectionMount) SetSnapshotExpirationTimeMillis(v int64)`

SetSnapshotExpirationTimeMillis sets SnapshotExpirationTimeMillis field to given value.

### HasSnapshotExpirationTimeMillis

`func (o *CollectionMount) HasSnapshotExpirationTimeMillis() bool`

HasSnapshotExpirationTimeMillis returns a boolean if a field has been set.

### GetState

`func (o *CollectionMount) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CollectionMount) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CollectionMount) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CollectionMount) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStats

`func (o *CollectionMount) GetStats() CollectionMountStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *CollectionMount) GetStatsOk() (*CollectionMountStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *CollectionMount) SetStats(v CollectionMountStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *CollectionMount) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetType

`func (o *CollectionMount) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CollectionMount) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CollectionMount) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CollectionMount) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVirtualInstanceId

`func (o *CollectionMount) GetVirtualInstanceId() string`

GetVirtualInstanceId returns the VirtualInstanceId field if non-nil, zero value otherwise.

### GetVirtualInstanceIdOk

`func (o *CollectionMount) GetVirtualInstanceIdOk() (*string, bool)`

GetVirtualInstanceIdOk returns a tuple with the VirtualInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualInstanceId

`func (o *CollectionMount) SetVirtualInstanceId(v string)`

SetVirtualInstanceId sets VirtualInstanceId field to given value.

### HasVirtualInstanceId

`func (o *CollectionMount) HasVirtualInstanceId() bool`

HasVirtualInstanceId returns a boolean if a field has been set.

### GetVirtualInstanceRrn

`func (o *CollectionMount) GetVirtualInstanceRrn() string`

GetVirtualInstanceRrn returns the VirtualInstanceRrn field if non-nil, zero value otherwise.

### GetVirtualInstanceRrnOk

`func (o *CollectionMount) GetVirtualInstanceRrnOk() (*string, bool)`

GetVirtualInstanceRrnOk returns a tuple with the VirtualInstanceRrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualInstanceRrn

`func (o *CollectionMount) SetVirtualInstanceRrn(v string)`

SetVirtualInstanceRrn sets VirtualInstanceRrn field to given value.

### HasVirtualInstanceRrn

`func (o *CollectionMount) HasVirtualInstanceRrn() bool`

HasVirtualInstanceRrn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


