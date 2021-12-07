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

// FieldOptions struct for FieldOptions
type FieldOptions struct {
	//  index or no_index
	IndexMode *string `json:"index_mode,omitempty"`
	//  v1_index or no_index
	RangeIndexMode *string `json:"range_index_mode,omitempty"`
	//  index or no_index
	TypeIndexMode *string `json:"type_index_mode,omitempty"`
	//  store or no_store
	ColumnIndexMode *string `json:"column_index_mode,omitempty"`
}

// NewFieldOptions instantiates a new FieldOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldOptions() *FieldOptions {
	this := FieldOptions{}
	return &this
}

// NewFieldOptionsWithDefaults instantiates a new FieldOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldOptionsWithDefaults() *FieldOptions {
	this := FieldOptions{}
	return &this
}

// GetIndexMode returns the IndexMode field value if set, zero value otherwise.
func (o *FieldOptions) GetIndexMode() string {
	if o == nil || o.IndexMode == nil {
		var ret string
		return ret
	}
	return *o.IndexMode
}

// GetIndexModeOk returns a tuple with the IndexMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldOptions) GetIndexModeOk() (*string, bool) {
	if o == nil || o.IndexMode == nil {
		return nil, false
	}
	return o.IndexMode, true
}

// HasIndexMode returns a boolean if a field has been set.
func (o *FieldOptions) HasIndexMode() bool {
	if o != nil && o.IndexMode != nil {
		return true
	}

	return false
}

// SetIndexMode gets a reference to the given string and assigns it to the IndexMode field.
func (o *FieldOptions) SetIndexMode(v string) {
	o.IndexMode = &v
}

// GetRangeIndexMode returns the RangeIndexMode field value if set, zero value otherwise.
func (o *FieldOptions) GetRangeIndexMode() string {
	if o == nil || o.RangeIndexMode == nil {
		var ret string
		return ret
	}
	return *o.RangeIndexMode
}

// GetRangeIndexModeOk returns a tuple with the RangeIndexMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldOptions) GetRangeIndexModeOk() (*string, bool) {
	if o == nil || o.RangeIndexMode == nil {
		return nil, false
	}
	return o.RangeIndexMode, true
}

// HasRangeIndexMode returns a boolean if a field has been set.
func (o *FieldOptions) HasRangeIndexMode() bool {
	if o != nil && o.RangeIndexMode != nil {
		return true
	}

	return false
}

// SetRangeIndexMode gets a reference to the given string and assigns it to the RangeIndexMode field.
func (o *FieldOptions) SetRangeIndexMode(v string) {
	o.RangeIndexMode = &v
}

// GetTypeIndexMode returns the TypeIndexMode field value if set, zero value otherwise.
func (o *FieldOptions) GetTypeIndexMode() string {
	if o == nil || o.TypeIndexMode == nil {
		var ret string
		return ret
	}
	return *o.TypeIndexMode
}

// GetTypeIndexModeOk returns a tuple with the TypeIndexMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldOptions) GetTypeIndexModeOk() (*string, bool) {
	if o == nil || o.TypeIndexMode == nil {
		return nil, false
	}
	return o.TypeIndexMode, true
}

// HasTypeIndexMode returns a boolean if a field has been set.
func (o *FieldOptions) HasTypeIndexMode() bool {
	if o != nil && o.TypeIndexMode != nil {
		return true
	}

	return false
}

// SetTypeIndexMode gets a reference to the given string and assigns it to the TypeIndexMode field.
func (o *FieldOptions) SetTypeIndexMode(v string) {
	o.TypeIndexMode = &v
}

// GetColumnIndexMode returns the ColumnIndexMode field value if set, zero value otherwise.
func (o *FieldOptions) GetColumnIndexMode() string {
	if o == nil || o.ColumnIndexMode == nil {
		var ret string
		return ret
	}
	return *o.ColumnIndexMode
}

// GetColumnIndexModeOk returns a tuple with the ColumnIndexMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldOptions) GetColumnIndexModeOk() (*string, bool) {
	if o == nil || o.ColumnIndexMode == nil {
		return nil, false
	}
	return o.ColumnIndexMode, true
}

// HasColumnIndexMode returns a boolean if a field has been set.
func (o *FieldOptions) HasColumnIndexMode() bool {
	if o != nil && o.ColumnIndexMode != nil {
		return true
	}

	return false
}

// SetColumnIndexMode gets a reference to the given string and assigns it to the ColumnIndexMode field.
func (o *FieldOptions) SetColumnIndexMode(v string) {
	o.ColumnIndexMode = &v
}

func (o FieldOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IndexMode != nil {
		toSerialize["index_mode"] = o.IndexMode
	}
	if o.RangeIndexMode != nil {
		toSerialize["range_index_mode"] = o.RangeIndexMode
	}
	if o.TypeIndexMode != nil {
		toSerialize["type_index_mode"] = o.TypeIndexMode
	}
	if o.ColumnIndexMode != nil {
		toSerialize["column_index_mode"] = o.ColumnIndexMode
	}
	return json.Marshal(toSerialize)
}

type NullableFieldOptions struct {
	value *FieldOptions
	isSet bool
}

func (v NullableFieldOptions) Get() *FieldOptions {
	return v.value
}

func (v *NullableFieldOptions) Set(val *FieldOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldOptions(val *FieldOptions) *NullableFieldOptions {
	return &NullableFieldOptions{value: val, isSet: true}
}

func (v NullableFieldOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


