# SourceS3Base

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settings** | Pointer to [**SourceS3Settings**](SourceS3Settings.md) |  | [optional] 

## Methods

### NewSourceS3Base

`func NewSourceS3Base() *SourceS3Base`

NewSourceS3Base instantiates a new SourceS3Base object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceS3BaseWithDefaults

`func NewSourceS3BaseWithDefaults() *SourceS3Base`

NewSourceS3BaseWithDefaults instantiates a new SourceS3Base object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettings

`func (o *SourceS3Base) GetSettings() SourceS3Settings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *SourceS3Base) GetSettingsOk() (*SourceS3Settings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *SourceS3Base) SetSettings(v SourceS3Settings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *SourceS3Base) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


