# QueryLambdaVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workspace** | Pointer to **string** | Workspace of this Query Lambda. | [optional] 
**CreatedBy** | Pointer to **string** | User that created this Query Lambda. | [optional] 
**CreatedAt** | Pointer to **string** | ISO-8601 date of when Query Lambda was created. | [optional] 
**Name** | Pointer to **string** | Query Lambda name. | [optional] 
**Version** | Pointer to **string** | Query Lambda version. | [optional] 
**Description** | Pointer to **string** | Optional description. | [optional] 
**Sql** | Pointer to [**QueryLambdaSql**](QueryLambdaSql.md) |  | [optional] 
**Collections** | Pointer to **[]string** | Collections queried by underlying SQL query. | [optional] 
**State** | Pointer to **string** | Status of this Query Lambda. | [optional] 
**Stats** | Pointer to [**QueryLambdaStats**](QueryLambdaStats.md) |  | [optional] 
**PublicAccessId** | Pointer to **string** | Public access ID associated with this QL version | [optional] 

## Methods

### NewQueryLambdaVersion

`func NewQueryLambdaVersion() *QueryLambdaVersion`

NewQueryLambdaVersion instantiates a new QueryLambdaVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryLambdaVersionWithDefaults

`func NewQueryLambdaVersionWithDefaults() *QueryLambdaVersion`

NewQueryLambdaVersionWithDefaults instantiates a new QueryLambdaVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspace

`func (o *QueryLambdaVersion) GetWorkspace() string`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *QueryLambdaVersion) GetWorkspaceOk() (*string, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *QueryLambdaVersion) SetWorkspace(v string)`

SetWorkspace sets Workspace field to given value.

### HasWorkspace

`func (o *QueryLambdaVersion) HasWorkspace() bool`

HasWorkspace returns a boolean if a field has been set.

### GetCreatedBy

`func (o *QueryLambdaVersion) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *QueryLambdaVersion) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *QueryLambdaVersion) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *QueryLambdaVersion) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *QueryLambdaVersion) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QueryLambdaVersion) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QueryLambdaVersion) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *QueryLambdaVersion) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetName

`func (o *QueryLambdaVersion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QueryLambdaVersion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QueryLambdaVersion) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QueryLambdaVersion) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *QueryLambdaVersion) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *QueryLambdaVersion) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *QueryLambdaVersion) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *QueryLambdaVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDescription

`func (o *QueryLambdaVersion) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QueryLambdaVersion) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QueryLambdaVersion) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QueryLambdaVersion) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSql

`func (o *QueryLambdaVersion) GetSql() QueryLambdaSql`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *QueryLambdaVersion) GetSqlOk() (*QueryLambdaSql, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *QueryLambdaVersion) SetSql(v QueryLambdaSql)`

SetSql sets Sql field to given value.

### HasSql

`func (o *QueryLambdaVersion) HasSql() bool`

HasSql returns a boolean if a field has been set.

### GetCollections

`func (o *QueryLambdaVersion) GetCollections() []string`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *QueryLambdaVersion) GetCollectionsOk() (*[]string, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *QueryLambdaVersion) SetCollections(v []string)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *QueryLambdaVersion) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetState

`func (o *QueryLambdaVersion) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *QueryLambdaVersion) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *QueryLambdaVersion) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *QueryLambdaVersion) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStats

`func (o *QueryLambdaVersion) GetStats() QueryLambdaStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *QueryLambdaVersion) GetStatsOk() (*QueryLambdaStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *QueryLambdaVersion) SetStats(v QueryLambdaStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *QueryLambdaVersion) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetPublicAccessId

`func (o *QueryLambdaVersion) GetPublicAccessId() string`

GetPublicAccessId returns the PublicAccessId field if non-nil, zero value otherwise.

### GetPublicAccessIdOk

`func (o *QueryLambdaVersion) GetPublicAccessIdOk() (*string, bool)`

GetPublicAccessIdOk returns a tuple with the PublicAccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAccessId

`func (o *QueryLambdaVersion) SetPublicAccessId(v string)`

SetPublicAccessId sets PublicAccessId field to given value.

### HasPublicAccessId

`func (o *QueryLambdaVersion) HasPublicAccessId() bool`

HasPublicAccessId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


