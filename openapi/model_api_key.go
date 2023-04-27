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

// ApiKey API keys are used to authenticate requests to Rockset's API. An API key is tied to the user who creates it.
type ApiKey struct {
	// Date that API key was created (ISO-8601 format).
	CreatedAt *string `json:"created_at,omitempty"`
	// Email of API key owner.
	CreatedBy *string `json:"created_by,omitempty"`
	// The expiration date of this API key.
	ExpiryTime *string `json:"expiry_time,omitempty"`
	// API key string of 64 alphanumeric characters.
	Key string `json:"key"`
	// Date that API key was most recently used (ISO-8601 format).
	LastAccessTime *string `json:"last_access_time,omitempty"`
	// Name of the API key.
	Name string `json:"name"`
	// Role specifying access control. If not specified, API key will have access to all of the associated user's roles.
	Role *string `json:"role,omitempty"`
	// Current state of this key.
	State *string `json:"state,omitempty"`
}

// NewApiKey instantiates a new ApiKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiKey(key string, name string) *ApiKey {
	this := ApiKey{}
	this.Key = key
	this.Name = name
	return &this
}

// NewApiKeyWithDefaults instantiates a new ApiKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiKeyWithDefaults() *ApiKey {
	this := ApiKey{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ApiKey) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKey) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ApiKey) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ApiKey) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ApiKey) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKey) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ApiKey) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *ApiKey) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetExpiryTime returns the ExpiryTime field value if set, zero value otherwise.
func (o *ApiKey) GetExpiryTime() string {
	if o == nil || o.ExpiryTime == nil {
		var ret string
		return ret
	}
	return *o.ExpiryTime
}

// GetExpiryTimeOk returns a tuple with the ExpiryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKey) GetExpiryTimeOk() (*string, bool) {
	if o == nil || o.ExpiryTime == nil {
		return nil, false
	}
	return o.ExpiryTime, true
}

// HasExpiryTime returns a boolean if a field has been set.
func (o *ApiKey) HasExpiryTime() bool {
	if o != nil && o.ExpiryTime != nil {
		return true
	}

	return false
}

// SetExpiryTime gets a reference to the given string and assigns it to the ExpiryTime field.
func (o *ApiKey) SetExpiryTime(v string) {
	o.ExpiryTime = &v
}

// GetKey returns the Key field value
func (o *ApiKey) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ApiKey) GetKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ApiKey) SetKey(v string) {
	o.Key = v
}

// GetLastAccessTime returns the LastAccessTime field value if set, zero value otherwise.
func (o *ApiKey) GetLastAccessTime() string {
	if o == nil || o.LastAccessTime == nil {
		var ret string
		return ret
	}
	return *o.LastAccessTime
}

// GetLastAccessTimeOk returns a tuple with the LastAccessTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKey) GetLastAccessTimeOk() (*string, bool) {
	if o == nil || o.LastAccessTime == nil {
		return nil, false
	}
	return o.LastAccessTime, true
}

// HasLastAccessTime returns a boolean if a field has been set.
func (o *ApiKey) HasLastAccessTime() bool {
	if o != nil && o.LastAccessTime != nil {
		return true
	}

	return false
}

// SetLastAccessTime gets a reference to the given string and assigns it to the LastAccessTime field.
func (o *ApiKey) SetLastAccessTime(v string) {
	o.LastAccessTime = &v
}

// GetName returns the Name field value
func (o *ApiKey) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiKey) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiKey) SetName(v string) {
	o.Name = v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *ApiKey) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKey) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *ApiKey) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *ApiKey) SetRole(v string) {
	o.Role = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ApiKey) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKey) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ApiKey) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ApiKey) SetState(v string) {
	o.State = &v
}

func (o ApiKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.ExpiryTime != nil {
		toSerialize["expiry_time"] = o.ExpiryTime
	}
	if true {
		toSerialize["key"] = o.Key
	}
	if o.LastAccessTime != nil {
		toSerialize["last_access_time"] = o.LastAccessTime
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableApiKey struct {
	value *ApiKey
	isSet bool
}

func (v NullableApiKey) Get() *ApiKey {
	return v.value
}

func (v *NullableApiKey) Set(val *ApiKey) {
	v.value = val
	v.isSet = true
}

func (v NullableApiKey) IsSet() bool {
	return v.isSet
}

func (v *NullableApiKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiKey(val *ApiKey) *NullableApiKey {
	return &NullableApiKey{value: val, isSet: true}
}

func (v NullableApiKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


