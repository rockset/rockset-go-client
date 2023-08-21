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

// checks if the QueryResponseStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryResponseStats{}

// QueryResponseStats struct for QueryResponseStats
type QueryResponseStats struct {
	// Query time in milliseconds.
	ElapsedTimeMs *int64 `json:"elapsed_time_ms,omitempty"`
	// Time query was throttled by admission control.
	ThrottledTimeMicros *int64 `json:"throttled_time_micros,omitempty"`
}

// NewQueryResponseStats instantiates a new QueryResponseStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryResponseStats() *QueryResponseStats {
	this := QueryResponseStats{}
	return &this
}

// NewQueryResponseStatsWithDefaults instantiates a new QueryResponseStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryResponseStatsWithDefaults() *QueryResponseStats {
	this := QueryResponseStats{}
	return &this
}

// GetElapsedTimeMs returns the ElapsedTimeMs field value if set, zero value otherwise.
func (o *QueryResponseStats) GetElapsedTimeMs() int64 {
	if o == nil || IsNil(o.ElapsedTimeMs) {
		var ret int64
		return ret
	}
	return *o.ElapsedTimeMs
}

// GetElapsedTimeMsOk returns a tuple with the ElapsedTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryResponseStats) GetElapsedTimeMsOk() (*int64, bool) {
	if o == nil || IsNil(o.ElapsedTimeMs) {
		return nil, false
	}
	return o.ElapsedTimeMs, true
}

// HasElapsedTimeMs returns a boolean if a field has been set.
func (o *QueryResponseStats) HasElapsedTimeMs() bool {
	if o != nil && !IsNil(o.ElapsedTimeMs) {
		return true
	}

	return false
}

// SetElapsedTimeMs gets a reference to the given int64 and assigns it to the ElapsedTimeMs field.
func (o *QueryResponseStats) SetElapsedTimeMs(v int64) {
	o.ElapsedTimeMs = &v
}

// GetThrottledTimeMicros returns the ThrottledTimeMicros field value if set, zero value otherwise.
func (o *QueryResponseStats) GetThrottledTimeMicros() int64 {
	if o == nil || IsNil(o.ThrottledTimeMicros) {
		var ret int64
		return ret
	}
	return *o.ThrottledTimeMicros
}

// GetThrottledTimeMicrosOk returns a tuple with the ThrottledTimeMicros field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryResponseStats) GetThrottledTimeMicrosOk() (*int64, bool) {
	if o == nil || IsNil(o.ThrottledTimeMicros) {
		return nil, false
	}
	return o.ThrottledTimeMicros, true
}

// HasThrottledTimeMicros returns a boolean if a field has been set.
func (o *QueryResponseStats) HasThrottledTimeMicros() bool {
	if o != nil && !IsNil(o.ThrottledTimeMicros) {
		return true
	}

	return false
}

// SetThrottledTimeMicros gets a reference to the given int64 and assigns it to the ThrottledTimeMicros field.
func (o *QueryResponseStats) SetThrottledTimeMicros(v int64) {
	o.ThrottledTimeMicros = &v
}

func (o QueryResponseStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryResponseStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ElapsedTimeMs) {
		toSerialize["elapsed_time_ms"] = o.ElapsedTimeMs
	}
	if !IsNil(o.ThrottledTimeMicros) {
		toSerialize["throttled_time_micros"] = o.ThrottledTimeMicros
	}
	return toSerialize, nil
}

type NullableQueryResponseStats struct {
	value *QueryResponseStats
	isSet bool
}

func (v NullableQueryResponseStats) Get() *QueryResponseStats {
	return v.value
}

func (v *NullableQueryResponseStats) Set(val *QueryResponseStats) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryResponseStats) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryResponseStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryResponseStats(val *QueryResponseStats) *NullableQueryResponseStats {
	return &NullableQueryResponseStats{value: val, isSet: true}
}

func (v NullableQueryResponseStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryResponseStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


