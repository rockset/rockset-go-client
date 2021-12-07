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

// ListRolesResponse struct for ListRolesResponse
type ListRolesResponse struct {
	// List of all roles.
	Data *[]Role `json:"data,omitempty"`
}

// NewListRolesResponse instantiates a new ListRolesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListRolesResponse() *ListRolesResponse {
	this := ListRolesResponse{}
	return &this
}

// NewListRolesResponseWithDefaults instantiates a new ListRolesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListRolesResponseWithDefaults() *ListRolesResponse {
	this := ListRolesResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListRolesResponse) GetData() []Role {
	if o == nil || o.Data == nil {
		var ret []Role
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRolesResponse) GetDataOk() (*[]Role, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListRolesResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Role and assigns it to the Data field.
func (o *ListRolesResponse) SetData(v []Role) {
	o.Data = &v
}

func (o ListRolesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListRolesResponse struct {
	value *ListRolesResponse
	isSet bool
}

func (v NullableListRolesResponse) Get() *ListRolesResponse {
	return v.value
}

func (v *NullableListRolesResponse) Set(val *ListRolesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListRolesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListRolesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRolesResponse(val *ListRolesResponse) *NullableListRolesResponse {
	return &NullableListRolesResponse{value: val, isSet: true}
}

func (v NullableListRolesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRolesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


