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

// checks if the ScheduledLambdaResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScheduledLambdaResponse{}

// ScheduledLambdaResponse struct for ScheduledLambdaResponse
type ScheduledLambdaResponse struct {
	Data *ScheduledLambda `json:"data,omitempty"`
}

// NewScheduledLambdaResponse instantiates a new ScheduledLambdaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduledLambdaResponse() *ScheduledLambdaResponse {
	this := ScheduledLambdaResponse{}
	return &this
}

// NewScheduledLambdaResponseWithDefaults instantiates a new ScheduledLambdaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduledLambdaResponseWithDefaults() *ScheduledLambdaResponse {
	this := ScheduledLambdaResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ScheduledLambdaResponse) GetData() ScheduledLambda {
	if o == nil || IsNil(o.Data) {
		var ret ScheduledLambda
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledLambdaResponse) GetDataOk() (*ScheduledLambda, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ScheduledLambdaResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ScheduledLambda and assigns it to the Data field.
func (o *ScheduledLambdaResponse) SetData(v ScheduledLambda) {
	o.Data = &v
}

func (o ScheduledLambdaResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScheduledLambdaResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableScheduledLambdaResponse struct {
	value *ScheduledLambdaResponse
	isSet bool
}

func (v NullableScheduledLambdaResponse) Get() *ScheduledLambdaResponse {
	return v.value
}

func (v *NullableScheduledLambdaResponse) Set(val *ScheduledLambdaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduledLambdaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduledLambdaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduledLambdaResponse(val *ScheduledLambdaResponse) *NullableScheduledLambdaResponse {
	return &NullableScheduledLambdaResponse{value: val, isSet: true}
}

func (v NullableScheduledLambdaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduledLambdaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
