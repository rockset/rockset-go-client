# Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | unique identifier for the cluster | [optional] 
**ClusterType** | Pointer to **string** | cluster type | [optional] 
**AwsRegion** | Pointer to **string** | aws region | [optional] 
**Domain** | Pointer to **string** | domain of org using cluster | [optional] 
**TopLevelDomain** | Pointer to **string** | top level domain of org using cluster | [optional] 
**ApiserverUrl** | Pointer to **string** | api server url for cluster | [optional] 

## Methods

### NewCluster

`func NewCluster() *Cluster`

NewCluster instantiates a new Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWithDefaults

`func NewClusterWithDefaults() *Cluster`

NewClusterWithDefaults instantiates a new Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Cluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cluster) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Cluster) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClusterType

`func (o *Cluster) GetClusterType() string`

GetClusterType returns the ClusterType field if non-nil, zero value otherwise.

### GetClusterTypeOk

`func (o *Cluster) GetClusterTypeOk() (*string, bool)`

GetClusterTypeOk returns a tuple with the ClusterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterType

`func (o *Cluster) SetClusterType(v string)`

SetClusterType sets ClusterType field to given value.

### HasClusterType

`func (o *Cluster) HasClusterType() bool`

HasClusterType returns a boolean if a field has been set.

### GetAwsRegion

`func (o *Cluster) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *Cluster) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *Cluster) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *Cluster) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.

### GetDomain

`func (o *Cluster) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Cluster) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Cluster) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Cluster) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetTopLevelDomain

`func (o *Cluster) GetTopLevelDomain() string`

GetTopLevelDomain returns the TopLevelDomain field if non-nil, zero value otherwise.

### GetTopLevelDomainOk

`func (o *Cluster) GetTopLevelDomainOk() (*string, bool)`

GetTopLevelDomainOk returns a tuple with the TopLevelDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLevelDomain

`func (o *Cluster) SetTopLevelDomain(v string)`

SetTopLevelDomain sets TopLevelDomain field to given value.

### HasTopLevelDomain

`func (o *Cluster) HasTopLevelDomain() bool`

HasTopLevelDomain returns a boolean if a field has been set.

### GetApiserverUrl

`func (o *Cluster) GetApiserverUrl() string`

GetApiserverUrl returns the ApiserverUrl field if non-nil, zero value otherwise.

### GetApiserverUrlOk

`func (o *Cluster) GetApiserverUrlOk() (*string, bool)`

GetApiserverUrlOk returns a tuple with the ApiserverUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiserverUrl

`func (o *Cluster) SetApiserverUrl(v string)`

SetApiserverUrl sets ApiserverUrl field to given value.

### HasApiserverUrl

`func (o *Cluster) HasApiserverUrl() bool`

HasApiserverUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


