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

// checks if the DeleteDocumentsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteDocumentsResponse{}

// DeleteDocumentsResponse struct for DeleteDocumentsResponse
type DeleteDocumentsResponse struct {
	// Information about deleted documents.
	Data []DocumentStatus `json:"data,omitempty"`
	// A string representing the collection offset after completing the deletes.
	LastOffset *string `json:"last_offset,omitempty"`
}

// NewDeleteDocumentsResponse instantiates a new DeleteDocumentsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteDocumentsResponse() *DeleteDocumentsResponse {
	this := DeleteDocumentsResponse{}
	return &this
}

// NewDeleteDocumentsResponseWithDefaults instantiates a new DeleteDocumentsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteDocumentsResponseWithDefaults() *DeleteDocumentsResponse {
	this := DeleteDocumentsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DeleteDocumentsResponse) GetData() []DocumentStatus {
	if o == nil || IsNil(o.Data) {
		var ret []DocumentStatus
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteDocumentsResponse) GetDataOk() ([]DocumentStatus, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DeleteDocumentsResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []DocumentStatus and assigns it to the Data field.
func (o *DeleteDocumentsResponse) SetData(v []DocumentStatus) {
	o.Data = v
}

// GetLastOffset returns the LastOffset field value if set, zero value otherwise.
func (o *DeleteDocumentsResponse) GetLastOffset() string {
	if o == nil || IsNil(o.LastOffset) {
		var ret string
		return ret
	}
	return *o.LastOffset
}

// GetLastOffsetOk returns a tuple with the LastOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteDocumentsResponse) GetLastOffsetOk() (*string, bool) {
	if o == nil || IsNil(o.LastOffset) {
		return nil, false
	}
	return o.LastOffset, true
}

// HasLastOffset returns a boolean if a field has been set.
func (o *DeleteDocumentsResponse) HasLastOffset() bool {
	if o != nil && !IsNil(o.LastOffset) {
		return true
	}

	return false
}

// SetLastOffset gets a reference to the given string and assigns it to the LastOffset field.
func (o *DeleteDocumentsResponse) SetLastOffset(v string) {
	o.LastOffset = &v
}

func (o DeleteDocumentsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteDocumentsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.LastOffset) {
		toSerialize["last_offset"] = o.LastOffset
	}
	return toSerialize, nil
}

type NullableDeleteDocumentsResponse struct {
	value *DeleteDocumentsResponse
	isSet bool
}

func (v NullableDeleteDocumentsResponse) Get() *DeleteDocumentsResponse {
	return v.value
}

func (v *NullableDeleteDocumentsResponse) Set(val *DeleteDocumentsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteDocumentsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteDocumentsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteDocumentsResponse(val *DeleteDocumentsResponse) *NullableDeleteDocumentsResponse {
	return &NullableDeleteDocumentsResponse{value: val, isSet: true}
}

func (v NullableDeleteDocumentsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteDocumentsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


