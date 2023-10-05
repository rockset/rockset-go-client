# SourceS3Settings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**S3ScanFrequency** | Pointer to **string** | Rockset scans an S3 bucket based on a defined time interval. The scan frequency determines the length of time between a new scan and the previous scan. If the previous scan finds new objects or updates to existing objects, Rockset immediately scans the bucket again after processing changes from the previous scan. | [optional] 

## Methods

### NewSourceS3Settings

`func NewSourceS3Settings() *SourceS3Settings`

NewSourceS3Settings instantiates a new SourceS3Settings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceS3SettingsWithDefaults

`func NewSourceS3SettingsWithDefaults() *SourceS3Settings`

NewSourceS3SettingsWithDefaults instantiates a new SourceS3Settings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetS3ScanFrequency

`func (o *SourceS3Settings) GetS3ScanFrequency() string`

GetS3ScanFrequency returns the S3ScanFrequency field if non-nil, zero value otherwise.

### GetS3ScanFrequencyOk

`func (o *SourceS3Settings) GetS3ScanFrequencyOk() (*string, bool)`

GetS3ScanFrequencyOk returns a tuple with the S3ScanFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3ScanFrequency

`func (o *SourceS3Settings) SetS3ScanFrequency(v string)`

SetS3ScanFrequency sets S3ScanFrequency field to given value.

### HasS3ScanFrequency

`func (o *SourceS3Settings) HasS3ScanFrequency() bool`

HasS3ScanFrequency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


