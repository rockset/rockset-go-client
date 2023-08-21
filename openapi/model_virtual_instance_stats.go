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

// checks if the VirtualInstanceStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualInstanceStats{}

// VirtualInstanceStats struct for VirtualInstanceStats
type VirtualInstanceStats struct {
	// Milliseconds since Unix epoch Jan 1, 1970.
	LastQueriedMs *int64 `json:"last_queried_ms,omitempty"`
}

// NewVirtualInstanceStats instantiates a new VirtualInstanceStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualInstanceStats() *VirtualInstanceStats {
	this := VirtualInstanceStats{}
	return &this
}

// NewVirtualInstanceStatsWithDefaults instantiates a new VirtualInstanceStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualInstanceStatsWithDefaults() *VirtualInstanceStats {
	this := VirtualInstanceStats{}
	return &this
}

// GetLastQueriedMs returns the LastQueriedMs field value if set, zero value otherwise.
func (o *VirtualInstanceStats) GetLastQueriedMs() int64 {
	if o == nil || IsNil(o.LastQueriedMs) {
		var ret int64
		return ret
	}
	return *o.LastQueriedMs
}

// GetLastQueriedMsOk returns a tuple with the LastQueriedMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualInstanceStats) GetLastQueriedMsOk() (*int64, bool) {
	if o == nil || IsNil(o.LastQueriedMs) {
		return nil, false
	}
	return o.LastQueriedMs, true
}

// HasLastQueriedMs returns a boolean if a field has been set.
func (o *VirtualInstanceStats) HasLastQueriedMs() bool {
	if o != nil && !IsNil(o.LastQueriedMs) {
		return true
	}

	return false
}

// SetLastQueriedMs gets a reference to the given int64 and assigns it to the LastQueriedMs field.
func (o *VirtualInstanceStats) SetLastQueriedMs(v int64) {
	o.LastQueriedMs = &v
}

func (o VirtualInstanceStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualInstanceStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LastQueriedMs) {
		toSerialize["last_queried_ms"] = o.LastQueriedMs
	}
	return toSerialize, nil
}

type NullableVirtualInstanceStats struct {
	value *VirtualInstanceStats
	isSet bool
}

func (v NullableVirtualInstanceStats) Get() *VirtualInstanceStats {
	return v.value
}

func (v *NullableVirtualInstanceStats) Set(val *VirtualInstanceStats) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualInstanceStats) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualInstanceStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualInstanceStats(val *VirtualInstanceStats) *NullableVirtualInstanceStats {
	return &NullableVirtualInstanceStats{value: val, isSet: true}
}

func (v NullableVirtualInstanceStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualInstanceStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


