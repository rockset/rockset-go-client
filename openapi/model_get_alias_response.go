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

// GetAliasResponse struct for GetAliasResponse
type GetAliasResponse struct {
	Data *Alias `json:"data,omitempty"`
}

// NewGetAliasResponse instantiates a new GetAliasResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAliasResponse() *GetAliasResponse {
	this := GetAliasResponse{}
	return &this
}

// NewGetAliasResponseWithDefaults instantiates a new GetAliasResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAliasResponseWithDefaults() *GetAliasResponse {
	this := GetAliasResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetAliasResponse) GetData() Alias {
	if o == nil || o.Data == nil {
		var ret Alias
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAliasResponse) GetDataOk() (*Alias, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetAliasResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given Alias and assigns it to the Data field.
func (o *GetAliasResponse) SetData(v Alias) {
	o.Data = &v
}

func (o GetAliasResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetAliasResponse struct {
	value *GetAliasResponse
	isSet bool
}

func (v NullableGetAliasResponse) Get() *GetAliasResponse {
	return v.value
}

func (v *NullableGetAliasResponse) Set(val *GetAliasResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAliasResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAliasResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAliasResponse(val *GetAliasResponse) *NullableGetAliasResponse {
	return &NullableGetAliasResponse{value: val, isSet: true}
}

func (v NullableGetAliasResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAliasResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


