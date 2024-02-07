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

// checks if the SourceSystem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceSystem{}

// SourceSystem struct for SourceSystem
type SourceSystem struct {
	// The type of this system source.
	Type *string `json:"type,omitempty"`
	// The workspace for which collections will have logs created. If unspecified, logs will be created for collections in all workspaces. Currently only supported for the INGEST_LOGS system source.
	Workspace *string `json:"workspace,omitempty"`
}

// NewSourceSystem instantiates a new SourceSystem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceSystem() *SourceSystem {
	this := SourceSystem{}
	return &this
}

// NewSourceSystemWithDefaults instantiates a new SourceSystem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceSystemWithDefaults() *SourceSystem {
	this := SourceSystem{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SourceSystem) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceSystem) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SourceSystem) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SourceSystem) SetType(v string) {
	o.Type = &v
}

// GetWorkspace returns the Workspace field value if set, zero value otherwise.
func (o *SourceSystem) GetWorkspace() string {
	if o == nil || IsNil(o.Workspace) {
		var ret string
		return ret
	}
	return *o.Workspace
}

// GetWorkspaceOk returns a tuple with the Workspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceSystem) GetWorkspaceOk() (*string, bool) {
	if o == nil || IsNil(o.Workspace) {
		return nil, false
	}
	return o.Workspace, true
}

// HasWorkspace returns a boolean if a field has been set.
func (o *SourceSystem) HasWorkspace() bool {
	if o != nil && !IsNil(o.Workspace) {
		return true
	}

	return false
}

// SetWorkspace gets a reference to the given string and assigns it to the Workspace field.
func (o *SourceSystem) SetWorkspace(v string) {
	o.Workspace = &v
}

func (o SourceSystem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceSystem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Workspace) {
		toSerialize["workspace"] = o.Workspace
	}
	return toSerialize, nil
}

type NullableSourceSystem struct {
	value *SourceSystem
	isSet bool
}

func (v NullableSourceSystem) Get() *SourceSystem {
	return v.value
}

func (v *NullableSourceSystem) Set(val *SourceSystem) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceSystem) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceSystem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceSystem(val *SourceSystem) *NullableSourceSystem {
	return &NullableSourceSystem{value: val, isSet: true}
}

func (v NullableSourceSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceSystem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
