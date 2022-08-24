# SourceGcs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** | Name of GCS bucket you want to ingest from. | [optional] 
**Prefix** | Pointer to **string** | Prefix that selects keys to ingest. | [optional] 
**Pattern** | Pointer to **string** | Glob-style pattern that selects keys to ingest. Only either prefix or pattern can be specified. | [optional] 
**ObjectCountDownloaded** | Pointer to **int64** |  | [optional] [readonly] 
**ObjectCountTotal** | Pointer to **int64** |  | [optional] [readonly] 
**ObjectBytesTotal** | Pointer to **int64** |  | [optional] [readonly] 
**ObjectBytesDownloaded** | Pointer to **int64** |  | [optional] [readonly] 

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

### GetPattern

`func (o *SourceGcs) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *SourceGcs) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *SourceGcs) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *SourceGcs) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetObjectCountDownloaded

`func (o *SourceGcs) GetObjectCountDownloaded() int64`

GetObjectCountDownloaded returns the ObjectCountDownloaded field if non-nil, zero value otherwise.

### GetObjectCountDownloadedOk

`func (o *SourceGcs) GetObjectCountDownloadedOk() (*int64, bool)`

GetObjectCountDownloadedOk returns a tuple with the ObjectCountDownloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectCountDownloaded

`func (o *SourceGcs) SetObjectCountDownloaded(v int64)`

SetObjectCountDownloaded sets ObjectCountDownloaded field to given value.

### HasObjectCountDownloaded

`func (o *SourceGcs) HasObjectCountDownloaded() bool`

HasObjectCountDownloaded returns a boolean if a field has been set.

### GetObjectCountTotal

`func (o *SourceGcs) GetObjectCountTotal() int64`

GetObjectCountTotal returns the ObjectCountTotal field if non-nil, zero value otherwise.

### GetObjectCountTotalOk

`func (o *SourceGcs) GetObjectCountTotalOk() (*int64, bool)`

GetObjectCountTotalOk returns a tuple with the ObjectCountTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectCountTotal

`func (o *SourceGcs) SetObjectCountTotal(v int64)`

SetObjectCountTotal sets ObjectCountTotal field to given value.

### HasObjectCountTotal

`func (o *SourceGcs) HasObjectCountTotal() bool`

HasObjectCountTotal returns a boolean if a field has been set.

### GetObjectBytesTotal

`func (o *SourceGcs) GetObjectBytesTotal() int64`

GetObjectBytesTotal returns the ObjectBytesTotal field if non-nil, zero value otherwise.

### GetObjectBytesTotalOk

`func (o *SourceGcs) GetObjectBytesTotalOk() (*int64, bool)`

GetObjectBytesTotalOk returns a tuple with the ObjectBytesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectBytesTotal

`func (o *SourceGcs) SetObjectBytesTotal(v int64)`

SetObjectBytesTotal sets ObjectBytesTotal field to given value.

### HasObjectBytesTotal

`func (o *SourceGcs) HasObjectBytesTotal() bool`

HasObjectBytesTotal returns a boolean if a field has been set.

### GetObjectBytesDownloaded

`func (o *SourceGcs) GetObjectBytesDownloaded() int64`

GetObjectBytesDownloaded returns the ObjectBytesDownloaded field if non-nil, zero value otherwise.

### GetObjectBytesDownloadedOk

`func (o *SourceGcs) GetObjectBytesDownloadedOk() (*int64, bool)`

GetObjectBytesDownloadedOk returns a tuple with the ObjectBytesDownloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectBytesDownloaded

`func (o *SourceGcs) SetObjectBytesDownloaded(v int64)`

SetObjectBytesDownloaded sets ObjectBytesDownloaded field to given value.

### HasObjectBytesDownloaded

`func (o *SourceGcs) HasObjectBytesDownloaded() bool`

HasObjectBytesDownloaded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


