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

// DynamodbIntegration struct for DynamodbIntegration
type DynamodbIntegration struct {
	AwsAccessKey *AwsAccessKey `json:"aws_access_key,omitempty"`
	AwsRole *AwsRole `json:"aws_role,omitempty"`
	S3ExportBucketName *string `json:"s3_export_bucket_name,omitempty"`
}

// NewDynamodbIntegration instantiates a new DynamodbIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamodbIntegration() *DynamodbIntegration {
	this := DynamodbIntegration{}
	return &this
}

// NewDynamodbIntegrationWithDefaults instantiates a new DynamodbIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamodbIntegrationWithDefaults() *DynamodbIntegration {
	this := DynamodbIntegration{}
	return &this
}

// GetAwsAccessKey returns the AwsAccessKey field value if set, zero value otherwise.
func (o *DynamodbIntegration) GetAwsAccessKey() AwsAccessKey {
	if o == nil || o.AwsAccessKey == nil {
		var ret AwsAccessKey
		return ret
	}
	return *o.AwsAccessKey
}

// GetAwsAccessKeyOk returns a tuple with the AwsAccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamodbIntegration) GetAwsAccessKeyOk() (*AwsAccessKey, bool) {
	if o == nil || o.AwsAccessKey == nil {
		return nil, false
	}
	return o.AwsAccessKey, true
}

// HasAwsAccessKey returns a boolean if a field has been set.
func (o *DynamodbIntegration) HasAwsAccessKey() bool {
	if o != nil && o.AwsAccessKey != nil {
		return true
	}

	return false
}

// SetAwsAccessKey gets a reference to the given AwsAccessKey and assigns it to the AwsAccessKey field.
func (o *DynamodbIntegration) SetAwsAccessKey(v AwsAccessKey) {
	o.AwsAccessKey = &v
}

// GetAwsRole returns the AwsRole field value if set, zero value otherwise.
func (o *DynamodbIntegration) GetAwsRole() AwsRole {
	if o == nil || o.AwsRole == nil {
		var ret AwsRole
		return ret
	}
	return *o.AwsRole
}

// GetAwsRoleOk returns a tuple with the AwsRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamodbIntegration) GetAwsRoleOk() (*AwsRole, bool) {
	if o == nil || o.AwsRole == nil {
		return nil, false
	}
	return o.AwsRole, true
}

// HasAwsRole returns a boolean if a field has been set.
func (o *DynamodbIntegration) HasAwsRole() bool {
	if o != nil && o.AwsRole != nil {
		return true
	}

	return false
}

// SetAwsRole gets a reference to the given AwsRole and assigns it to the AwsRole field.
func (o *DynamodbIntegration) SetAwsRole(v AwsRole) {
	o.AwsRole = &v
}

// GetS3ExportBucketName returns the S3ExportBucketName field value if set, zero value otherwise.
func (o *DynamodbIntegration) GetS3ExportBucketName() string {
	if o == nil || o.S3ExportBucketName == nil {
		var ret string
		return ret
	}
	return *o.S3ExportBucketName
}

// GetS3ExportBucketNameOk returns a tuple with the S3ExportBucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamodbIntegration) GetS3ExportBucketNameOk() (*string, bool) {
	if o == nil || o.S3ExportBucketName == nil {
		return nil, false
	}
	return o.S3ExportBucketName, true
}

// HasS3ExportBucketName returns a boolean if a field has been set.
func (o *DynamodbIntegration) HasS3ExportBucketName() bool {
	if o != nil && o.S3ExportBucketName != nil {
		return true
	}

	return false
}

// SetS3ExportBucketName gets a reference to the given string and assigns it to the S3ExportBucketName field.
func (o *DynamodbIntegration) SetS3ExportBucketName(v string) {
	o.S3ExportBucketName = &v
}

func (o DynamodbIntegration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AwsAccessKey != nil {
		toSerialize["aws_access_key"] = o.AwsAccessKey
	}
	if o.AwsRole != nil {
		toSerialize["aws_role"] = o.AwsRole
	}
	if o.S3ExportBucketName != nil {
		toSerialize["s3_export_bucket_name"] = o.S3ExportBucketName
	}
	return json.Marshal(toSerialize)
}

type NullableDynamodbIntegration struct {
	value *DynamodbIntegration
	isSet bool
}

func (v NullableDynamodbIntegration) Get() *DynamodbIntegration {
	return v.value
}

func (v *NullableDynamodbIntegration) Set(val *DynamodbIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamodbIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamodbIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamodbIntegration(val *DynamodbIntegration) *NullableDynamodbIntegration {
	return &NullableDynamodbIntegration{value: val, isSet: true}
}

func (v NullableDynamodbIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamodbIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


