# EventTimeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | name of the field containing event time | 
**Format** | Pointer to **string** | format of time field, can be one of: milliseconds_since_epoch, seconds_since_epoch | [optional] 
**TimeZone** | Pointer to **string** | default time zone, in standard IANA format | [optional] 

## Methods

### NewEventTimeInfo

`func NewEventTimeInfo(field string, ) *EventTimeInfo`

NewEventTimeInfo instantiates a new EventTimeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventTimeInfoWithDefaults

`func NewEventTimeInfoWithDefaults() *EventTimeInfo`

NewEventTimeInfoWithDefaults instantiates a new EventTimeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *EventTimeInfo) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *EventTimeInfo) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *EventTimeInfo) SetField(v string)`

SetField sets Field field to given value.


### GetFormat

`func (o *EventTimeInfo) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *EventTimeInfo) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *EventTimeInfo) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *EventTimeInfo) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetTimeZone

`func (o *EventTimeInfo) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *EventTimeInfo) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *EventTimeInfo) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *EventTimeInfo) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


