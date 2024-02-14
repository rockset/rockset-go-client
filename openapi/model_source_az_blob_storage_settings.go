/*
REST API

Rockset's REST API allows for creating and managing all resources in Rockset. Each supported endpoint is documented below.  All requests must be authorized with a Rockset API key, which can be created in the [Rockset console](https://console.rockset.com). The API key must be provided as `ApiKey <api_key>` in the `Authorization` request header. For example: ``` Authorization: ApiKey aB35kDjg93J5nsf4GjwMeErAVd832F7ad4vhsW1S02kfZiab42sTsfW5Sxt25asT ```  All endpoints are only accessible via https.  Build something awesome!

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SourceAzBlobStorageSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceAzBlobStorageSettings{}

// SourceAzBlobStorageSettings struct for SourceAzBlobStorageSettings
type SourceAzBlobStorageSettings struct {
	// Rockset scans an Azure blob Storage container based on a defined time interval. The scan frequency determines the length of time between a new scan and the previous scan. If the previous scan finds new objects or updates to existing objects, Rockset immediately scans the bucket again after processing changes from the previous scan. Duration value is of type ISO 8601 (e.g. PT5H, PT4M, PT3S). It doesn't account for DST, leap seconds and leap years. Minimum value: PT1S. Maximum value: PT1H.
	AzblobScanFrequency *string `json:"azblob_scan_frequency,omitempty"`
}

// NewSourceAzBlobStorageSettings instantiates a new SourceAzBlobStorageSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceAzBlobStorageSettings() *SourceAzBlobStorageSettings {
	this := SourceAzBlobStorageSettings{}
	return &this
}

// NewSourceAzBlobStorageSettingsWithDefaults instantiates a new SourceAzBlobStorageSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceAzBlobStorageSettingsWithDefaults() *SourceAzBlobStorageSettings {
	this := SourceAzBlobStorageSettings{}
	return &this
}

// GetAzblobScanFrequency returns the AzblobScanFrequency field value if set, zero value otherwise.
func (o *SourceAzBlobStorageSettings) GetAzblobScanFrequency() string {
	if o == nil || IsNil(o.AzblobScanFrequency) {
		var ret string
		return ret
	}
	return *o.AzblobScanFrequency
}

// GetAzblobScanFrequencyOk returns a tuple with the AzblobScanFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceAzBlobStorageSettings) GetAzblobScanFrequencyOk() (*string, bool) {
	if o == nil || IsNil(o.AzblobScanFrequency) {
		return nil, false
	}
	return o.AzblobScanFrequency, true
}

// HasAzblobScanFrequency returns a boolean if a field has been set.
func (o *SourceAzBlobStorageSettings) HasAzblobScanFrequency() bool {
	if o != nil && !IsNil(o.AzblobScanFrequency) {
		return true
	}

	return false
}

// SetAzblobScanFrequency gets a reference to the given string and assigns it to the AzblobScanFrequency field.
func (o *SourceAzBlobStorageSettings) SetAzblobScanFrequency(v string) {
	o.AzblobScanFrequency = &v
}

func (o SourceAzBlobStorageSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceAzBlobStorageSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AzblobScanFrequency) {
		toSerialize["azblob_scan_frequency"] = o.AzblobScanFrequency
	}
	return toSerialize, nil
}

type NullableSourceAzBlobStorageSettings struct {
	value *SourceAzBlobStorageSettings
	isSet bool
}

func (v NullableSourceAzBlobStorageSettings) Get() *SourceAzBlobStorageSettings {
	return v.value
}

func (v *NullableSourceAzBlobStorageSettings) Set(val *SourceAzBlobStorageSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceAzBlobStorageSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceAzBlobStorageSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceAzBlobStorageSettings(val *SourceAzBlobStorageSettings) *NullableSourceAzBlobStorageSettings {
	return &NullableSourceAzBlobStorageSettings{value: val, isSet: true}
}

func (v NullableSourceAzBlobStorageSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceAzBlobStorageSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


