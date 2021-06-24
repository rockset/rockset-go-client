# SourceS3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** | AWS credential with ListObjects and GetObject access | [optional] [readonly] 
**SecretAccess** | Pointer to **string** | AWS credential with ListObjects and GetObject access | [optional] [readonly] 
**Prefix** | Pointer to **string** | Prefix that selects keys to ingest. | [optional] 
**Pattern** | Pointer to **string** | Pattern that selects keys to ingest. | [optional] 
**Region** | Pointer to **string** | AWS region containing source bucket | [optional] 
**Bucket** | **string** | address of S3 bucket containing data | 
**Prefixes** | **[]string** | list of prefixes to paths from which data should be ingested | [readonly] 
**Format** | Pointer to **string** | do not use | [optional] 
**Mappings** | Pointer to [**[]FieldMask**](FieldMask.md) | custom transformation on data field | [optional] 
**ObjectCountDownloaded** | Pointer to **int64** |  | [optional] 
**ObjectCountTotal** | Pointer to **int64** |  | [optional] 
**ObjectBytesTotal** | Pointer to **int64** |  | [optional] 

## Methods

### NewSourceS3

`func NewSourceS3(bucket string, prefixes []string, ) *SourceS3`

NewSourceS3 instantiates a new SourceS3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceS3WithDefaults

`func NewSourceS3WithDefaults() *SourceS3`

NewSourceS3WithDefaults instantiates a new SourceS3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *SourceS3) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *SourceS3) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *SourceS3) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *SourceS3) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetSecretAccess

`func (o *SourceS3) GetSecretAccess() string`

GetSecretAccess returns the SecretAccess field if non-nil, zero value otherwise.

### GetSecretAccessOk

`func (o *SourceS3) GetSecretAccessOk() (*string, bool)`

GetSecretAccessOk returns a tuple with the SecretAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccess

`func (o *SourceS3) SetSecretAccess(v string)`

SetSecretAccess sets SecretAccess field to given value.

### HasSecretAccess

`func (o *SourceS3) HasSecretAccess() bool`

HasSecretAccess returns a boolean if a field has been set.

### GetPrefix

`func (o *SourceS3) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *SourceS3) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *SourceS3) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *SourceS3) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetPattern

`func (o *SourceS3) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *SourceS3) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *SourceS3) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *SourceS3) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetRegion

`func (o *SourceS3) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *SourceS3) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *SourceS3) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *SourceS3) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetBucket

`func (o *SourceS3) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *SourceS3) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *SourceS3) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetPrefixes

`func (o *SourceS3) GetPrefixes() []string`

GetPrefixes returns the Prefixes field if non-nil, zero value otherwise.

### GetPrefixesOk

`func (o *SourceS3) GetPrefixesOk() (*[]string, bool)`

GetPrefixesOk returns a tuple with the Prefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixes

`func (o *SourceS3) SetPrefixes(v []string)`

SetPrefixes sets Prefixes field to given value.


### GetFormat

`func (o *SourceS3) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *SourceS3) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *SourceS3) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *SourceS3) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetMappings

`func (o *SourceS3) GetMappings() []FieldMask`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *SourceS3) GetMappingsOk() (*[]FieldMask, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *SourceS3) SetMappings(v []FieldMask)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *SourceS3) HasMappings() bool`

HasMappings returns a boolean if a field has been set.

### GetObjectCountDownloaded

`func (o *SourceS3) GetObjectCountDownloaded() int64`

GetObjectCountDownloaded returns the ObjectCountDownloaded field if non-nil, zero value otherwise.

### GetObjectCountDownloadedOk

`func (o *SourceS3) GetObjectCountDownloadedOk() (*int64, bool)`

GetObjectCountDownloadedOk returns a tuple with the ObjectCountDownloaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectCountDownloaded

`func (o *SourceS3) SetObjectCountDownloaded(v int64)`

SetObjectCountDownloaded sets ObjectCountDownloaded field to given value.

### HasObjectCountDownloaded

`func (o *SourceS3) HasObjectCountDownloaded() bool`

HasObjectCountDownloaded returns a boolean if a field has been set.

### GetObjectCountTotal

`func (o *SourceS3) GetObjectCountTotal() int64`

GetObjectCountTotal returns the ObjectCountTotal field if non-nil, zero value otherwise.

### GetObjectCountTotalOk

`func (o *SourceS3) GetObjectCountTotalOk() (*int64, bool)`

GetObjectCountTotalOk returns a tuple with the ObjectCountTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectCountTotal

`func (o *SourceS3) SetObjectCountTotal(v int64)`

SetObjectCountTotal sets ObjectCountTotal field to given value.

### HasObjectCountTotal

`func (o *SourceS3) HasObjectCountTotal() bool`

HasObjectCountTotal returns a boolean if a field has been set.

### GetObjectBytesTotal

`func (o *SourceS3) GetObjectBytesTotal() int64`

GetObjectBytesTotal returns the ObjectBytesTotal field if non-nil, zero value otherwise.

### GetObjectBytesTotalOk

`func (o *SourceS3) GetObjectBytesTotalOk() (*int64, bool)`

GetObjectBytesTotalOk returns a tuple with the ObjectBytesTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectBytesTotal

`func (o *SourceS3) SetObjectBytesTotal(v int64)`

SetObjectBytesTotal sets ObjectBytesTotal field to given value.

### HasObjectBytesTotal

`func (o *SourceS3) HasObjectBytesTotal() bool`

HasObjectBytesTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


