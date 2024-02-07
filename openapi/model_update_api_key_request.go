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

// checks if the UpdateApiKeyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateApiKeyRequest{}

// UpdateApiKeyRequest struct for UpdateApiKeyRequest
type UpdateApiKeyRequest struct {
	// If set to true, the expiration time for this key will be cleared.
	ClearExpiryTime *bool `json:"clear_expiry_time,omitempty"`
	// If provided, the API key will automatically expire at this time (ISO-8601 format). Requires premium.
	ExpiryTime *string `json:"expiry_time,omitempty"`
	// State that the api key should be set to.
	State *string `json:"state,omitempty"`
}

// NewUpdateApiKeyRequest instantiates a new UpdateApiKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateApiKeyRequest() *UpdateApiKeyRequest {
	this := UpdateApiKeyRequest{}
	return &this
}

// NewUpdateApiKeyRequestWithDefaults instantiates a new UpdateApiKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateApiKeyRequestWithDefaults() *UpdateApiKeyRequest {
	this := UpdateApiKeyRequest{}
	return &this
}

// GetClearExpiryTime returns the ClearExpiryTime field value if set, zero value otherwise.
func (o *UpdateApiKeyRequest) GetClearExpiryTime() bool {
	if o == nil || IsNil(o.ClearExpiryTime) {
		var ret bool
		return ret
	}
	return *o.ClearExpiryTime
}

// GetClearExpiryTimeOk returns a tuple with the ClearExpiryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateApiKeyRequest) GetClearExpiryTimeOk() (*bool, bool) {
	if o == nil || IsNil(o.ClearExpiryTime) {
		return nil, false
	}
	return o.ClearExpiryTime, true
}

// HasClearExpiryTime returns a boolean if a field has been set.
func (o *UpdateApiKeyRequest) HasClearExpiryTime() bool {
	if o != nil && !IsNil(o.ClearExpiryTime) {
		return true
	}

	return false
}

// SetClearExpiryTime gets a reference to the given bool and assigns it to the ClearExpiryTime field.
func (o *UpdateApiKeyRequest) SetClearExpiryTime(v bool) {
	o.ClearExpiryTime = &v
}

// GetExpiryTime returns the ExpiryTime field value if set, zero value otherwise.
func (o *UpdateApiKeyRequest) GetExpiryTime() string {
	if o == nil || IsNil(o.ExpiryTime) {
		var ret string
		return ret
	}
	return *o.ExpiryTime
}

// GetExpiryTimeOk returns a tuple with the ExpiryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateApiKeyRequest) GetExpiryTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiryTime) {
		return nil, false
	}
	return o.ExpiryTime, true
}

// HasExpiryTime returns a boolean if a field has been set.
func (o *UpdateApiKeyRequest) HasExpiryTime() bool {
	if o != nil && !IsNil(o.ExpiryTime) {
		return true
	}

	return false
}

// SetExpiryTime gets a reference to the given string and assigns it to the ExpiryTime field.
func (o *UpdateApiKeyRequest) SetExpiryTime(v string) {
	o.ExpiryTime = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *UpdateApiKeyRequest) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateApiKeyRequest) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *UpdateApiKeyRequest) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *UpdateApiKeyRequest) SetState(v string) {
	o.State = &v
}

func (o UpdateApiKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateApiKeyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClearExpiryTime) {
		toSerialize["clear_expiry_time"] = o.ClearExpiryTime
	}
	if !IsNil(o.ExpiryTime) {
		toSerialize["expiry_time"] = o.ExpiryTime
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableUpdateApiKeyRequest struct {
	value *UpdateApiKeyRequest
	isSet bool
}

func (v NullableUpdateApiKeyRequest) Get() *UpdateApiKeyRequest {
	return v.value
}

func (v *NullableUpdateApiKeyRequest) Set(val *UpdateApiKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateApiKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateApiKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateApiKeyRequest(val *UpdateApiKeyRequest) *NullableUpdateApiKeyRequest {
	return &NullableUpdateApiKeyRequest{value: val, isSet: true}
}

func (v NullableUpdateApiKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateApiKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
