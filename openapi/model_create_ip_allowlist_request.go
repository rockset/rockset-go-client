/*
 * REST API
 *
 * Rockset's REST API allows for creating and managing all resources in Rockset. Each supported endpoint is documented below.  All requests must be authorized with a Rockset API key, which can be created in the [Rockset console](https://console.rockset.com). The API key must be provided as `ApiKey <api_key>` in the `Authorization` request header. For example: ``` Authorization: ApiKey aB35kDjg93J5nsf4GjwMeErAVd832F7ad4vhsW1S02kfZiab42sTsfW5Sxt25asT ```  All endpoints are only accessible via https.  Build something awesome!
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CreateIpAllowlistRequest struct for CreateIpAllowlistRequest
type CreateIpAllowlistRequest struct {
	// IP Allowlist policy name
	Name string `json:"name"`
	// optional description
	Description *string `json:"description,omitempty"`
	// individual IP address or range of IP addresses in CIDR notation
	IpAddress string `json:"ip_address"`
}

// NewCreateIpAllowlistRequest instantiates a new CreateIpAllowlistRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateIpAllowlistRequest(name string, ipAddress string) *CreateIpAllowlistRequest {
	this := CreateIpAllowlistRequest{}
	this.Name = name
	this.IpAddress = ipAddress
	return &this
}

// NewCreateIpAllowlistRequestWithDefaults instantiates a new CreateIpAllowlistRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateIpAllowlistRequestWithDefaults() *CreateIpAllowlistRequest {
	this := CreateIpAllowlistRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateIpAllowlistRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateIpAllowlistRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateIpAllowlistRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateIpAllowlistRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateIpAllowlistRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateIpAllowlistRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateIpAllowlistRequest) SetDescription(v string) {
	o.Description = &v
}

// GetIpAddress returns the IpAddress field value
func (o *CreateIpAllowlistRequest) GetIpAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value
// and a boolean to check if the value has been set.
func (o *CreateIpAllowlistRequest) GetIpAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IpAddress, true
}

// SetIpAddress sets field value
func (o *CreateIpAllowlistRequest) SetIpAddress(v string) {
	o.IpAddress = v
}

func (o CreateIpAllowlistRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["ip_address"] = o.IpAddress
	}
	return json.Marshal(toSerialize)
}

type NullableCreateIpAllowlistRequest struct {
	value *CreateIpAllowlistRequest
	isSet bool
}

func (v NullableCreateIpAllowlistRequest) Get() *CreateIpAllowlistRequest {
	return v.value
}

func (v *NullableCreateIpAllowlistRequest) Set(val *CreateIpAllowlistRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateIpAllowlistRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateIpAllowlistRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateIpAllowlistRequest(val *CreateIpAllowlistRequest) *NullableCreateIpAllowlistRequest {
	return &NullableCreateIpAllowlistRequest{value: val, isSet: true}
}

func (v NullableCreateIpAllowlistRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateIpAllowlistRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


