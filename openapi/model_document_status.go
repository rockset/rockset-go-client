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

// DocumentStatus struct for DocumentStatus
type DocumentStatus struct {
	// Collection name.
	Collection *string `json:"_collection,omitempty"`
	// Unique document ID.
	Id *string `json:"_id,omitempty"`
	Error *ErrorModel `json:"error,omitempty"`
	// Unique id used to represent each patch request.
	PatchId *string `json:"patch_id,omitempty"`
	// Status of the document.
	Status *string `json:"status,omitempty"`
}

// NewDocumentStatus instantiates a new DocumentStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentStatus() *DocumentStatus {
	this := DocumentStatus{}
	return &this
}

// NewDocumentStatusWithDefaults instantiates a new DocumentStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentStatusWithDefaults() *DocumentStatus {
	this := DocumentStatus{}
	return &this
}

// GetCollection returns the Collection field value if set, zero value otherwise.
func (o *DocumentStatus) GetCollection() string {
	if o == nil || o.Collection == nil {
		var ret string
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentStatus) GetCollectionOk() (*string, bool) {
	if o == nil || o.Collection == nil {
		return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *DocumentStatus) HasCollection() bool {
	if o != nil && o.Collection != nil {
		return true
	}

	return false
}

// SetCollection gets a reference to the given string and assigns it to the Collection field.
func (o *DocumentStatus) SetCollection(v string) {
	o.Collection = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DocumentStatus) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentStatus) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DocumentStatus) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DocumentStatus) SetId(v string) {
	o.Id = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *DocumentStatus) GetError() ErrorModel {
	if o == nil || o.Error == nil {
		var ret ErrorModel
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentStatus) GetErrorOk() (*ErrorModel, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *DocumentStatus) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given ErrorModel and assigns it to the Error field.
func (o *DocumentStatus) SetError(v ErrorModel) {
	o.Error = &v
}

// GetPatchId returns the PatchId field value if set, zero value otherwise.
func (o *DocumentStatus) GetPatchId() string {
	if o == nil || o.PatchId == nil {
		var ret string
		return ret
	}
	return *o.PatchId
}

// GetPatchIdOk returns a tuple with the PatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentStatus) GetPatchIdOk() (*string, bool) {
	if o == nil || o.PatchId == nil {
		return nil, false
	}
	return o.PatchId, true
}

// HasPatchId returns a boolean if a field has been set.
func (o *DocumentStatus) HasPatchId() bool {
	if o != nil && o.PatchId != nil {
		return true
	}

	return false
}

// SetPatchId gets a reference to the given string and assigns it to the PatchId field.
func (o *DocumentStatus) SetPatchId(v string) {
	o.PatchId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DocumentStatus) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentStatus) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DocumentStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DocumentStatus) SetStatus(v string) {
	o.Status = &v
}

func (o DocumentStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Collection != nil {
		toSerialize["_collection"] = o.Collection
	}
	if o.Id != nil {
		toSerialize["_id"] = o.Id
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.PatchId != nil {
		toSerialize["patch_id"] = o.PatchId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableDocumentStatus struct {
	value *DocumentStatus
	isSet bool
}

func (v NullableDocumentStatus) Get() *DocumentStatus {
	return v.value
}

func (v *NullableDocumentStatus) Set(val *DocumentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentStatus(val *DocumentStatus) *NullableDocumentStatus {
	return &NullableDocumentStatus{value: val, isSet: true}
}

func (v NullableDocumentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


