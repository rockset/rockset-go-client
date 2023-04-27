# QueryLambda

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **[]string** | Collections/aliases queried by underlying SQL query. | [optional] 
**LastUpdated** | Pointer to **string** | ISO-8601 date of when Query Lambda was last updated. | [optional] 
**LastUpdatedBy** | Pointer to **string** | User that created this Query Lambda. | [optional] 
**LatestVersion** | Pointer to [**QueryLambdaVersion**](QueryLambdaVersion.md) |  | [optional] 
**Name** | Pointer to **string** | Query Lambda name. | [optional] 
**VersionCount** | Pointer to **int32** | Number of Query Lambda versions. | [optional] 
**Workspace** | Pointer to **string** | Workspace of this Query Lambda. | [optional] 

## Methods

### NewQueryLambda

`func NewQueryLambda() *QueryLambda`

NewQueryLambda instantiates a new QueryLambda object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryLambdaWithDefaults

`func NewQueryLambdaWithDefaults() *QueryLambda`

NewQueryLambdaWithDefaults instantiates a new QueryLambda object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *QueryLambda) GetCollections() []string`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *QueryLambda) GetCollectionsOk() (*[]string, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *QueryLambda) SetCollections(v []string)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *QueryLambda) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetLastUpdated

`func (o *QueryLambda) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *QueryLambda) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *QueryLambda) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *QueryLambda) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *QueryLambda) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *QueryLambda) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *QueryLambda) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *QueryLambda) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetLatestVersion

`func (o *QueryLambda) GetLatestVersion() QueryLambdaVersion`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *QueryLambda) GetLatestVersionOk() (*QueryLambdaVersion, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *QueryLambda) SetLatestVersion(v QueryLambdaVersion)`

SetLatestVersion sets LatestVersion field to given value.

### HasLatestVersion

`func (o *QueryLambda) HasLatestVersion() bool`

HasLatestVersion returns a boolean if a field has been set.

### GetName

`func (o *QueryLambda) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QueryLambda) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QueryLambda) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QueryLambda) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersionCount

`func (o *QueryLambda) GetVersionCount() int32`

GetVersionCount returns the VersionCount field if non-nil, zero value otherwise.

### GetVersionCountOk

`func (o *QueryLambda) GetVersionCountOk() (*int32, bool)`

GetVersionCountOk returns a tuple with the VersionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionCount

`func (o *QueryLambda) SetVersionCount(v int32)`

SetVersionCount sets VersionCount field to given value.

### HasVersionCount

`func (o *QueryLambda) HasVersionCount() bool`

HasVersionCount returns a boolean if a field has been set.

### GetWorkspace

`func (o *QueryLambda) GetWorkspace() string`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *QueryLambda) GetWorkspaceOk() (*string, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *QueryLambda) SetWorkspace(v string)`

SetWorkspace sets Workspace field to given value.

### HasWorkspace

`func (o *QueryLambda) HasWorkspace() bool`

HasWorkspace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


