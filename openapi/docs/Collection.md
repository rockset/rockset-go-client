# Collection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | ISO-8601 date. | [optional] 
**CreatedBy** | Pointer to **string** | Email of user who created the collection. | [optional] 
**Name** | Pointer to **string** | Unique identifer for collection, can contain alphanumeric or dash characters. | [optional] 
**Description** | Pointer to **string** | Text describing the collection. | [optional] 
**Workspace** | Pointer to **string** | Name of the workspace that the collection is in. | [optional] 
**Status** | Pointer to **string** | Current status of collection. | [optional] 
**Sources** | Pointer to [**[]Source**](Source.md) | List of sources from which collection ingests. | [optional] 
**Stats** | Pointer to [**CollectionStats**](CollectionStats.md) |  | [optional] 
**RetentionSecs** | Pointer to **int64** | Number of seconds after which data is purged based on event time. | [optional] 
**FieldMappings** | Pointer to [**[]FieldMappingV2**](FieldMappingV2.md) | List of mappings applied on all documents in a collection. | [optional] 
**FieldMappingQuery** | Pointer to [**FieldMappingQuery**](FieldMappingQuery.md) |  | [optional] 
**ClusteringKey** | Pointer to [**[]FieldPartition**](FieldPartition.md) | List of clustering fields for a collection. | [optional] 
**Aliases** | Pointer to [**[]Alias**](Alias.md) | List of aliases for a collection. | [optional] 
**ReadOnly** | Pointer to **bool** | Whether the collection is read-only or not. | [optional] 
**InsertOnly** | Pointer to **bool** | Whether the collection is insert only or not. | [optional] 

## Methods

### NewCollection

`func NewCollection() *Collection`

NewCollection instantiates a new Collection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionWithDefaults

`func NewCollectionWithDefaults() *Collection`

NewCollectionWithDefaults instantiates a new Collection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Collection) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Collection) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Collection) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Collection) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Collection) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Collection) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Collection) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Collection) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetName

`func (o *Collection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Collection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Collection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Collection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Collection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Collection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Collection) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Collection) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetWorkspace

`func (o *Collection) GetWorkspace() string`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *Collection) GetWorkspaceOk() (*string, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *Collection) SetWorkspace(v string)`

SetWorkspace sets Workspace field to given value.

### HasWorkspace

`func (o *Collection) HasWorkspace() bool`

HasWorkspace returns a boolean if a field has been set.

### GetStatus

`func (o *Collection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Collection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Collection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Collection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSources

`func (o *Collection) GetSources() []Source`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *Collection) GetSourcesOk() (*[]Source, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *Collection) SetSources(v []Source)`

SetSources sets Sources field to given value.

### HasSources

`func (o *Collection) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetStats

`func (o *Collection) GetStats() CollectionStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Collection) GetStatsOk() (*CollectionStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Collection) SetStats(v CollectionStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *Collection) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetRetentionSecs

`func (o *Collection) GetRetentionSecs() int64`

GetRetentionSecs returns the RetentionSecs field if non-nil, zero value otherwise.

### GetRetentionSecsOk

`func (o *Collection) GetRetentionSecsOk() (*int64, bool)`

GetRetentionSecsOk returns a tuple with the RetentionSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionSecs

`func (o *Collection) SetRetentionSecs(v int64)`

SetRetentionSecs sets RetentionSecs field to given value.

### HasRetentionSecs

`func (o *Collection) HasRetentionSecs() bool`

HasRetentionSecs returns a boolean if a field has been set.

### GetFieldMappings

`func (o *Collection) GetFieldMappings() []FieldMappingV2`

GetFieldMappings returns the FieldMappings field if non-nil, zero value otherwise.

### GetFieldMappingsOk

`func (o *Collection) GetFieldMappingsOk() (*[]FieldMappingV2, bool)`

GetFieldMappingsOk returns a tuple with the FieldMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMappings

`func (o *Collection) SetFieldMappings(v []FieldMappingV2)`

SetFieldMappings sets FieldMappings field to given value.

### HasFieldMappings

`func (o *Collection) HasFieldMappings() bool`

HasFieldMappings returns a boolean if a field has been set.

### GetFieldMappingQuery

`func (o *Collection) GetFieldMappingQuery() FieldMappingQuery`

GetFieldMappingQuery returns the FieldMappingQuery field if non-nil, zero value otherwise.

### GetFieldMappingQueryOk

`func (o *Collection) GetFieldMappingQueryOk() (*FieldMappingQuery, bool)`

GetFieldMappingQueryOk returns a tuple with the FieldMappingQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMappingQuery

`func (o *Collection) SetFieldMappingQuery(v FieldMappingQuery)`

SetFieldMappingQuery sets FieldMappingQuery field to given value.

### HasFieldMappingQuery

`func (o *Collection) HasFieldMappingQuery() bool`

HasFieldMappingQuery returns a boolean if a field has been set.

### GetClusteringKey

`func (o *Collection) GetClusteringKey() []FieldPartition`

GetClusteringKey returns the ClusteringKey field if non-nil, zero value otherwise.

### GetClusteringKeyOk

`func (o *Collection) GetClusteringKeyOk() (*[]FieldPartition, bool)`

GetClusteringKeyOk returns a tuple with the ClusteringKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusteringKey

`func (o *Collection) SetClusteringKey(v []FieldPartition)`

SetClusteringKey sets ClusteringKey field to given value.

### HasClusteringKey

`func (o *Collection) HasClusteringKey() bool`

HasClusteringKey returns a boolean if a field has been set.

### GetAliases

`func (o *Collection) GetAliases() []Alias`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *Collection) GetAliasesOk() (*[]Alias, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *Collection) SetAliases(v []Alias)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *Collection) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetReadOnly

`func (o *Collection) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *Collection) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *Collection) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *Collection) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetInsertOnly

`func (o *Collection) GetInsertOnly() bool`

GetInsertOnly returns the InsertOnly field if non-nil, zero value otherwise.

### GetInsertOnlyOk

`func (o *Collection) GetInsertOnlyOk() (*bool, bool)`

GetInsertOnlyOk returns a tuple with the InsertOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertOnly

`func (o *Collection) SetInsertOnly(v bool)`

SetInsertOnly sets InsertOnly field to given value.

### HasInsertOnly

`func (o *Collection) HasInsertOnly() bool`

HasInsertOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


