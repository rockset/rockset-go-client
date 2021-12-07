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

// FieldMask struct for FieldMask
type FieldMask struct {
	InputPath []string `json:"input_path"`
	Mask *FieldMaskMask `json:"mask,omitempty"`
}

// NewFieldMask instantiates a new FieldMask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldMask(inputPath []string) *FieldMask {
	this := FieldMask{}
	this.InputPath = inputPath
	return &this
}

// NewFieldMaskWithDefaults instantiates a new FieldMask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldMaskWithDefaults() *FieldMask {
	this := FieldMask{}
	return &this
}

// GetInputPath returns the InputPath field value
func (o *FieldMask) GetInputPath() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.InputPath
}

// GetInputPathOk returns a tuple with the InputPath field value
// and a boolean to check if the value has been set.
func (o *FieldMask) GetInputPathOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InputPath, true
}

// SetInputPath sets field value
func (o *FieldMask) SetInputPath(v []string) {
	o.InputPath = v
}

// GetMask returns the Mask field value if set, zero value otherwise.
func (o *FieldMask) GetMask() FieldMaskMask {
	if o == nil || o.Mask == nil {
		var ret FieldMaskMask
		return ret
	}
	return *o.Mask
}

// GetMaskOk returns a tuple with the Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldMask) GetMaskOk() (*FieldMaskMask, bool) {
	if o == nil || o.Mask == nil {
		return nil, false
	}
	return o.Mask, true
}

// HasMask returns a boolean if a field has been set.
func (o *FieldMask) HasMask() bool {
	if o != nil && o.Mask != nil {
		return true
	}

	return false
}

// SetMask gets a reference to the given FieldMaskMask and assigns it to the Mask field.
func (o *FieldMask) SetMask(v FieldMaskMask) {
	o.Mask = &v
}

func (o FieldMask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["input_path"] = o.InputPath
	}
	if o.Mask != nil {
		toSerialize["mask"] = o.Mask
	}
	return json.Marshal(toSerialize)
}

type NullableFieldMask struct {
	value *FieldMask
	isSet bool
}

func (v NullableFieldMask) Get() *FieldMask {
	return v.value
}

func (v *NullableFieldMask) Set(val *FieldMask) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldMask) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldMask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldMask(val *FieldMask) *NullableFieldMask {
	return &NullableFieldMask{value: val, isSet: true}
}

func (v NullableFieldMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldMask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


