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

// CreateUserRequest struct for CreateUserRequest
type CreateUserRequest struct {
	// User email, must be unique.
	Email string `json:"email"`
	// List of roles for a given user.
	Roles []string `json:"roles"`
}

// NewCreateUserRequest instantiates a new CreateUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUserRequest(email string, roles []string) *CreateUserRequest {
	this := CreateUserRequest{}
	this.Email = email
	this.Roles = roles
	return &this
}

// NewCreateUserRequestWithDefaults instantiates a new CreateUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUserRequestWithDefaults() *CreateUserRequest {
	this := CreateUserRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *CreateUserRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CreateUserRequest) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CreateUserRequest) SetEmail(v string) {
	o.Email = v
}

// GetRoles returns the Roles field value
func (o *CreateUserRequest) GetRoles() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *CreateUserRequest) GetRolesOk() ([]string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *CreateUserRequest) SetRoles(v []string) {
	o.Roles = v
}

func (o CreateUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["roles"] = o.Roles
	}
	return json.Marshal(toSerialize)
}

type NullableCreateUserRequest struct {
	value *CreateUserRequest
	isSet bool
}

func (v NullableCreateUserRequest) Get() *CreateUserRequest {
	return v.value
}

func (v *NullableCreateUserRequest) Set(val *CreateUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUserRequest(val *CreateUserRequest) *NullableCreateUserRequest {
	return &NullableCreateUserRequest{value: val, isSet: true}
}

func (v NullableCreateUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


