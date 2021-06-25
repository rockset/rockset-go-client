# SourceFileUpload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | **string** | name of the file | 
**FileSize** | **int64** | size of the file in bytes | 
**FileUploadTime** | **string** | time of file upload | 

## Methods

### NewSourceFileUpload

`func NewSourceFileUpload(fileName string, fileSize int64, fileUploadTime string, ) *SourceFileUpload`

NewSourceFileUpload instantiates a new SourceFileUpload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceFileUploadWithDefaults

`func NewSourceFileUploadWithDefaults() *SourceFileUpload`

NewSourceFileUploadWithDefaults instantiates a new SourceFileUpload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *SourceFileUpload) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *SourceFileUpload) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *SourceFileUpload) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetFileSize

`func (o *SourceFileUpload) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *SourceFileUpload) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *SourceFileUpload) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.


### GetFileUploadTime

`func (o *SourceFileUpload) GetFileUploadTime() string`

GetFileUploadTime returns the FileUploadTime field if non-nil, zero value otherwise.

### GetFileUploadTimeOk

`func (o *SourceFileUpload) GetFileUploadTimeOk() (*string, bool)`

GetFileUploadTimeOk returns a tuple with the FileUploadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUploadTime

`func (o *SourceFileUpload) SetFileUploadTime(v string)`

SetFileUploadTime sets FileUploadTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


