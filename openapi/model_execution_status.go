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

// checks if the ExecutionStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExecutionStatus{}

// ExecutionStatus struct for ExecutionStatus
type ExecutionStatus struct {
	// Message describing the execution status.
	Message *string `json:"message,omitempty"`
	// The scheduled query lambda execution status
	State *string `json:"state,omitempty"`
}

// NewExecutionStatus instantiates a new ExecutionStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecutionStatus() *ExecutionStatus {
	this := ExecutionStatus{}
	return &this
}

// NewExecutionStatusWithDefaults instantiates a new ExecutionStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecutionStatusWithDefaults() *ExecutionStatus {
	this := ExecutionStatus{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ExecutionStatus) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionStatus) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ExecutionStatus) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ExecutionStatus) SetMessage(v string) {
	o.Message = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ExecutionStatus) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecutionStatus) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ExecutionStatus) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ExecutionStatus) SetState(v string) {
	o.State = &v
}

func (o ExecutionStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExecutionStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableExecutionStatus struct {
	value *ExecutionStatus
	isSet bool
}

func (v NullableExecutionStatus) Get() *ExecutionStatus {
	return v.value
}

func (v *NullableExecutionStatus) Set(val *ExecutionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableExecutionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableExecutionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecutionStatus(val *ExecutionStatus) *NullableExecutionStatus {
	return &NullableExecutionStatus{value: val, isSet: true}
}

func (v NullableExecutionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecutionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


