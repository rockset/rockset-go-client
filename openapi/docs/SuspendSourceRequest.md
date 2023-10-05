# SuspendSourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResumeAfterDuration** | Pointer to **string** | duration to suspend source; 1h is the default | [optional] 

## Methods

### NewSuspendSourceRequest

`func NewSuspendSourceRequest() *SuspendSourceRequest`

NewSuspendSourceRequest instantiates a new SuspendSourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuspendSourceRequestWithDefaults

`func NewSuspendSourceRequestWithDefaults() *SuspendSourceRequest`

NewSuspendSourceRequestWithDefaults instantiates a new SuspendSourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResumeAfterDuration

`func (o *SuspendSourceRequest) GetResumeAfterDuration() string`

GetResumeAfterDuration returns the ResumeAfterDuration field if non-nil, zero value otherwise.

### GetResumeAfterDurationOk

`func (o *SuspendSourceRequest) GetResumeAfterDurationOk() (*string, bool)`

GetResumeAfterDurationOk returns a tuple with the ResumeAfterDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeAfterDuration

`func (o *SuspendSourceRequest) SetResumeAfterDuration(v string)`

SetResumeAfterDuration sets ResumeAfterDuration field to given value.

### HasResumeAfterDuration

`func (o *SuspendSourceRequest) HasResumeAfterDuration() bool`

HasResumeAfterDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


