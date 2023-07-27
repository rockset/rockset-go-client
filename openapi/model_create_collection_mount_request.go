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

// CreateCollectionMountRequest struct for CreateCollectionMountRequest
type CreateCollectionMountRequest struct {
	// Collections to mount.
	CollectionPaths []string `json:"collection_paths,omitempty"`
}

// NewCreateCollectionMountRequest instantiates a new CreateCollectionMountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCollectionMountRequest() *CreateCollectionMountRequest {
	this := CreateCollectionMountRequest{}
	return &this
}

// NewCreateCollectionMountRequestWithDefaults instantiates a new CreateCollectionMountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCollectionMountRequestWithDefaults() *CreateCollectionMountRequest {
	this := CreateCollectionMountRequest{}
	return &this
}

// GetCollectionPaths returns the CollectionPaths field value if set, zero value otherwise.
func (o *CreateCollectionMountRequest) GetCollectionPaths() []string {
	if o == nil || o.CollectionPaths == nil {
		var ret []string
		return ret
	}
	return o.CollectionPaths
}

// GetCollectionPathsOk returns a tuple with the CollectionPaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCollectionMountRequest) GetCollectionPathsOk() ([]string, bool) {
	if o == nil || o.CollectionPaths == nil {
		return nil, false
	}
	return o.CollectionPaths, true
}

// HasCollectionPaths returns a boolean if a field has been set.
func (o *CreateCollectionMountRequest) HasCollectionPaths() bool {
	if o != nil && o.CollectionPaths != nil {
		return true
	}

	return false
}

// SetCollectionPaths gets a reference to the given []string and assigns it to the CollectionPaths field.
func (o *CreateCollectionMountRequest) SetCollectionPaths(v []string) {
	o.CollectionPaths = v
}

func (o CreateCollectionMountRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CollectionPaths != nil {
		toSerialize["collection_paths"] = o.CollectionPaths
	}
	return json.Marshal(toSerialize)
}

type NullableCreateCollectionMountRequest struct {
	value *CreateCollectionMountRequest
	isSet bool
}

func (v NullableCreateCollectionMountRequest) Get() *CreateCollectionMountRequest {
	return v.value
}

func (v *NullableCreateCollectionMountRequest) Set(val *CreateCollectionMountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCollectionMountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCollectionMountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCollectionMountRequest(val *CreateCollectionMountRequest) *NullableCreateCollectionMountRequest {
	return &NullableCreateCollectionMountRequest{value: val, isSet: true}
}

func (v NullableCreateCollectionMountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCollectionMountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


