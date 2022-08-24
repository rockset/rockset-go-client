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

// MongoDbIntegration struct for MongoDbIntegration
type MongoDbIntegration struct {
	// MongoDB connection URI string.
	ConnectionUri string `json:"connection_uri"`
}

// NewMongoDbIntegration instantiates a new MongoDbIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMongoDbIntegration(connectionUri string) *MongoDbIntegration {
	this := MongoDbIntegration{}
	this.ConnectionUri = connectionUri
	return &this
}

// NewMongoDbIntegrationWithDefaults instantiates a new MongoDbIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMongoDbIntegrationWithDefaults() *MongoDbIntegration {
	this := MongoDbIntegration{}
	return &this
}

// GetConnectionUri returns the ConnectionUri field value
func (o *MongoDbIntegration) GetConnectionUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionUri
}

// GetConnectionUriOk returns a tuple with the ConnectionUri field value
// and a boolean to check if the value has been set.
func (o *MongoDbIntegration) GetConnectionUriOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConnectionUri, true
}

// SetConnectionUri sets field value
func (o *MongoDbIntegration) SetConnectionUri(v string) {
	o.ConnectionUri = v
}

func (o MongoDbIntegration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["connection_uri"] = o.ConnectionUri
	}
	return json.Marshal(toSerialize)
}

type NullableMongoDbIntegration struct {
	value *MongoDbIntegration
	isSet bool
}

func (v NullableMongoDbIntegration) Get() *MongoDbIntegration {
	return v.value
}

func (v *NullableMongoDbIntegration) Set(val *MongoDbIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableMongoDbIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableMongoDbIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMongoDbIntegration(val *MongoDbIntegration) *NullableMongoDbIntegration {
	return &NullableMongoDbIntegration{value: val, isSet: true}
}

func (v NullableMongoDbIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMongoDbIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


