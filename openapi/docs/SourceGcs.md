# SourceGcs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** | name of GCS bucket you want to ingest from | [optional] 
**Prefix** | Pointer to **string** | Prefix that selects keys to ingest. | [optional] 

## Methods

### NewSourceGcs

`func NewSourceGcs() *SourceGcs`

NewSourceGcs instantiates a new SourceGcs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceGcsWithDefaults

`func NewSourceGcsWithDefaults() *SourceGcs`

NewSourceGcsWithDefaults instantiates a new SourceGcs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *SourceGcs) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *SourceGcs) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *SourceGcs) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *SourceGcs) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetPrefix

`func (o *SourceGcs) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *SourceGcs) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *SourceGcs) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *SourceGcs) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


