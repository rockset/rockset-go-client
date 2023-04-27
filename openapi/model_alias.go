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

// Alias struct for Alias
type Alias struct {
	// List of fully qualified collection names referenced by alias.
	Collections []string `json:"collections,omitempty"`
	// ISO-8601 date.
	CreatedAt *string `json:"created_at,omitempty"`
	// Email of the creator.
	CreatorEmail *string `json:"creator_email,omitempty"`
	// Alias description.
	Description *string `json:"description,omitempty"`
	// ISO-8601 date.
	ModifiedAt *string `json:"modified_at,omitempty"`
	// Name of the alias.
	Name *string `json:"name,omitempty"`
	// State of the alias.
	State *string `json:"state,omitempty"`
	// Name of the workspace.
	Workspace *string `json:"workspace,omitempty"`
}

// NewAlias instantiates a new Alias object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlias() *Alias {
	this := Alias{}
	return &this
}

// NewAliasWithDefaults instantiates a new Alias object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAliasWithDefaults() *Alias {
	this := Alias{}
	return &this
}

// GetCollections returns the Collections field value if set, zero value otherwise.
func (o *Alias) GetCollections() []string {
	if o == nil || o.Collections == nil {
		var ret []string
		return ret
	}
	return o.Collections
}

// GetCollectionsOk returns a tuple with the Collections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alias) GetCollectionsOk() ([]string, bool) {
	if o == nil || o.Collections == nil {
		return nil, false
	}
	return o.Collections, true
}

// HasCollections returns a boolean if a field has been set.
func (o *Alias) HasCollections() bool {
	if o != nil && o.Collections != nil {
		return true
	}

	return false
}

// SetCollections gets a reference to the given []string and assigns it to the Collections field.
func (o *Alias) SetCollections(v []string) {
	o.Collections = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Alias) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alias) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Alias) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Alias) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCreatorEmail returns the CreatorEmail field value if set, zero value otherwise.
func (o *Alias) GetCreatorEmail() string {
	if o == nil || o.CreatorEmail == nil {
		var ret string
		return ret
	}
	return *o.CreatorEmail
}

// GetCreatorEmailOk returns a tuple with the CreatorEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alias) GetCreatorEmailOk() (*string, bool) {
	if o == nil || o.CreatorEmail == nil {
		return nil, false
	}
	return o.CreatorEmail, true
}

// HasCreatorEmail returns a boolean if a field has been set.
func (o *Alias) HasCreatorEmail() bool {
	if o != nil && o.CreatorEmail != nil {
		return true
	}

	return false
}

// SetCreatorEmail gets a reference to the given string and assigns it to the CreatorEmail field.
func (o *Alias) SetCreatorEmail(v string) {
	o.CreatorEmail = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Alias) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alias) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Alias) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Alias) SetDescription(v string) {
	o.Description = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *Alias) GetModifiedAt() string {
	if o == nil || o.ModifiedAt == nil {
		var ret string
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alias) GetModifiedAtOk() (*string, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *Alias) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given string and assigns it to the ModifiedAt field.
func (o *Alias) SetModifiedAt(v string) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Alias) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alias) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Alias) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Alias) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Alias) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alias) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Alias) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Alias) SetState(v string) {
	o.State = &v
}

// GetWorkspace returns the Workspace field value if set, zero value otherwise.
func (o *Alias) GetWorkspace() string {
	if o == nil || o.Workspace == nil {
		var ret string
		return ret
	}
	return *o.Workspace
}

// GetWorkspaceOk returns a tuple with the Workspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alias) GetWorkspaceOk() (*string, bool) {
	if o == nil || o.Workspace == nil {
		return nil, false
	}
	return o.Workspace, true
}

// HasWorkspace returns a boolean if a field has been set.
func (o *Alias) HasWorkspace() bool {
	if o != nil && o.Workspace != nil {
		return true
	}

	return false
}

// SetWorkspace gets a reference to the given string and assigns it to the Workspace field.
func (o *Alias) SetWorkspace(v string) {
	o.Workspace = &v
}

func (o Alias) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Collections != nil {
		toSerialize["collections"] = o.Collections
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.CreatorEmail != nil {
		toSerialize["creator_email"] = o.CreatorEmail
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ModifiedAt != nil {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Workspace != nil {
		toSerialize["workspace"] = o.Workspace
	}
	return json.Marshal(toSerialize)
}

type NullableAlias struct {
	value *Alias
	isSet bool
}

func (v NullableAlias) Get() *Alias {
	return v.value
}

func (v *NullableAlias) Set(val *Alias) {
	v.value = val
	v.isSet = true
}

func (v NullableAlias) IsSet() bool {
	return v.isSet
}

func (v *NullableAlias) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlias(val *Alias) *NullableAlias {
	return &NullableAlias{value: val, isSet: true}
}

func (v NullableAlias) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlias) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


