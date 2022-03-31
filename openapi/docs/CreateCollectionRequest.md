# CreateCollectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | unique identifier for collection, can contain alphanumeric or dash characters | 
**Description** | Pointer to **string** | text describing the collection | [optional] 
**Sources** | Pointer to [**[]Source**](Source.md) | list of sources from which to ingest data | [optional] 
**RetentionSecs** | Pointer to **int64** | number of seconds after which data is purged, based on event time | [optional] 
**InsertOnly** | Pointer to **bool** | If true disallows updates and deletes, but makes indexing more efficient | [optional] 
**EventTimeInfo** | Pointer to [**EventTimeInfo**](EventTimeInfo.md) |  | [optional] 
**FieldMappings** | Pointer to [**[]FieldMappingV2**](FieldMappingV2.md) | list of mappings | [optional] 
**FieldMappingQuery** | Pointer to [**FieldMappingQuery**](FieldMappingQuery.md) |  | [optional] 
**ClusteringKey** | Pointer to [**[]FieldPartition**](FieldPartition.md) | list of clustering fields | [optional] 

## Methods

### NewCreateCollectionRequest

`func NewCreateCollectionRequest(name string, ) *CreateCollectionRequest`

NewCreateCollectionRequest instantiates a new CreateCollectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCollectionRequestWithDefaults

`func NewCreateCollectionRequestWithDefaults() *CreateCollectionRequest`

NewCreateCollectionRequestWithDefaults instantiates a new CreateCollectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateCollectionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCollectionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCollectionRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateCollectionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCollectionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCollectionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateCollectionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSources

`func (o *CreateCollectionRequest) GetSources() []Source`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *CreateCollectionRequest) GetSourcesOk() (*[]Source, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *CreateCollectionRequest) SetSources(v []Source)`

SetSources sets Sources field to given value.

### HasSources

`func (o *CreateCollectionRequest) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetRetentionSecs

`func (o *CreateCollectionRequest) GetRetentionSecs() int64`

GetRetentionSecs returns the RetentionSecs field if non-nil, zero value otherwise.

### GetRetentionSecsOk

`func (o *CreateCollectionRequest) GetRetentionSecsOk() (*int64, bool)`

GetRetentionSecsOk returns a tuple with the RetentionSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionSecs

`func (o *CreateCollectionRequest) SetRetentionSecs(v int64)`

SetRetentionSecs sets RetentionSecs field to given value.

### HasRetentionSecs

`func (o *CreateCollectionRequest) HasRetentionSecs() bool`

HasRetentionSecs returns a boolean if a field has been set.

### GetInsertOnly

`func (o *CreateCollectionRequest) GetInsertOnly() bool`

GetInsertOnly returns the InsertOnly field if non-nil, zero value otherwise.

### GetInsertOnlyOk

`func (o *CreateCollectionRequest) GetInsertOnlyOk() (*bool, bool)`

GetInsertOnlyOk returns a tuple with the InsertOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertOnly

`func (o *CreateCollectionRequest) SetInsertOnly(v bool)`

SetInsertOnly sets InsertOnly field to given value.

### HasInsertOnly

`func (o *CreateCollectionRequest) HasInsertOnly() bool`

HasInsertOnly returns a boolean if a field has been set.

### GetEventTimeInfo

`func (o *CreateCollectionRequest) GetEventTimeInfo() EventTimeInfo`

GetEventTimeInfo returns the EventTimeInfo field if non-nil, zero value otherwise.

### GetEventTimeInfoOk

`func (o *CreateCollectionRequest) GetEventTimeInfoOk() (*EventTimeInfo, bool)`

GetEventTimeInfoOk returns a tuple with the EventTimeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimeInfo

`func (o *CreateCollectionRequest) SetEventTimeInfo(v EventTimeInfo)`

SetEventTimeInfo sets EventTimeInfo field to given value.

### HasEventTimeInfo

`func (o *CreateCollectionRequest) HasEventTimeInfo() bool`

HasEventTimeInfo returns a boolean if a field has been set.

### GetFieldMappings

`func (o *CreateCollectionRequest) GetFieldMappings() []FieldMappingV2`

GetFieldMappings returns the FieldMappings field if non-nil, zero value otherwise.

### GetFieldMappingsOk

`func (o *CreateCollectionRequest) GetFieldMappingsOk() (*[]FieldMappingV2, bool)`

GetFieldMappingsOk returns a tuple with the FieldMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMappings

`func (o *CreateCollectionRequest) SetFieldMappings(v []FieldMappingV2)`

SetFieldMappings sets FieldMappings field to given value.

### HasFieldMappings

`func (o *CreateCollectionRequest) HasFieldMappings() bool`

HasFieldMappings returns a boolean if a field has been set.

### GetFieldMappingQuery

`func (o *CreateCollectionRequest) GetFieldMappingQuery() FieldMappingQuery`

GetFieldMappingQuery returns the FieldMappingQuery field if non-nil, zero value otherwise.

### GetFieldMappingQueryOk

`func (o *CreateCollectionRequest) GetFieldMappingQueryOk() (*FieldMappingQuery, bool)`

GetFieldMappingQueryOk returns a tuple with the FieldMappingQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldMappingQuery

`func (o *CreateCollectionRequest) SetFieldMappingQuery(v FieldMappingQuery)`

SetFieldMappingQuery sets FieldMappingQuery field to given value.

### HasFieldMappingQuery

`func (o *CreateCollectionRequest) HasFieldMappingQuery() bool`

HasFieldMappingQuery returns a boolean if a field has been set.

### GetClusteringKey

`func (o *CreateCollectionRequest) GetClusteringKey() []FieldPartition`

GetClusteringKey returns the ClusteringKey field if non-nil, zero value otherwise.

### GetClusteringKeyOk

`func (o *CreateCollectionRequest) GetClusteringKeyOk() (*[]FieldPartition, bool)`

GetClusteringKeyOk returns a tuple with the ClusteringKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusteringKey

`func (o *CreateCollectionRequest) SetClusteringKey(v []FieldPartition)`

SetClusteringKey sets ClusteringKey field to given value.

### HasClusteringKey

`func (o *CreateCollectionRequest) HasClusteringKey() bool`

HasClusteringKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


