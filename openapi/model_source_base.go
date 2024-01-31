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

// checks if the SourceBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceBase{}

// SourceBase struct for SourceBase
type SourceBase struct {
	AzureBlobStorage *SourceAzBlobStorageBase `json:"azure_blob_storage,omitempty"`
	Dynamodb *SourceDynamoDbBase `json:"dynamodb,omitempty"`
	Gcs *SourceGcsBase `json:"gcs,omitempty"`
	S3 *SourceS3Base `json:"s3,omitempty"`
}

// NewSourceBase instantiates a new SourceBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceBase() *SourceBase {
	this := SourceBase{}
	return &this
}

// NewSourceBaseWithDefaults instantiates a new SourceBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceBaseWithDefaults() *SourceBase {
	this := SourceBase{}
	return &this
}

// GetAzureBlobStorage returns the AzureBlobStorage field value if set, zero value otherwise.
func (o *SourceBase) GetAzureBlobStorage() SourceAzBlobStorageBase {
	if o == nil || IsNil(o.AzureBlobStorage) {
		var ret SourceAzBlobStorageBase
		return ret
	}
	return *o.AzureBlobStorage
}

// GetAzureBlobStorageOk returns a tuple with the AzureBlobStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceBase) GetAzureBlobStorageOk() (*SourceAzBlobStorageBase, bool) {
	if o == nil || IsNil(o.AzureBlobStorage) {
		return nil, false
	}
	return o.AzureBlobStorage, true
}

// HasAzureBlobStorage returns a boolean if a field has been set.
func (o *SourceBase) HasAzureBlobStorage() bool {
	if o != nil && !IsNil(o.AzureBlobStorage) {
		return true
	}

	return false
}

// SetAzureBlobStorage gets a reference to the given SourceAzBlobStorageBase and assigns it to the AzureBlobStorage field.
func (o *SourceBase) SetAzureBlobStorage(v SourceAzBlobStorageBase) {
	o.AzureBlobStorage = &v
}

// GetDynamodb returns the Dynamodb field value if set, zero value otherwise.
func (o *SourceBase) GetDynamodb() SourceDynamoDbBase {
	if o == nil || IsNil(o.Dynamodb) {
		var ret SourceDynamoDbBase
		return ret
	}
	return *o.Dynamodb
}

// GetDynamodbOk returns a tuple with the Dynamodb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceBase) GetDynamodbOk() (*SourceDynamoDbBase, bool) {
	if o == nil || IsNil(o.Dynamodb) {
		return nil, false
	}
	return o.Dynamodb, true
}

// HasDynamodb returns a boolean if a field has been set.
func (o *SourceBase) HasDynamodb() bool {
	if o != nil && !IsNil(o.Dynamodb) {
		return true
	}

	return false
}

// SetDynamodb gets a reference to the given SourceDynamoDbBase and assigns it to the Dynamodb field.
func (o *SourceBase) SetDynamodb(v SourceDynamoDbBase) {
	o.Dynamodb = &v
}

// GetGcs returns the Gcs field value if set, zero value otherwise.
func (o *SourceBase) GetGcs() SourceGcsBase {
	if o == nil || IsNil(o.Gcs) {
		var ret SourceGcsBase
		return ret
	}
	return *o.Gcs
}

// GetGcsOk returns a tuple with the Gcs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceBase) GetGcsOk() (*SourceGcsBase, bool) {
	if o == nil || IsNil(o.Gcs) {
		return nil, false
	}
	return o.Gcs, true
}

// HasGcs returns a boolean if a field has been set.
func (o *SourceBase) HasGcs() bool {
	if o != nil && !IsNil(o.Gcs) {
		return true
	}

	return false
}

// SetGcs gets a reference to the given SourceGcsBase and assigns it to the Gcs field.
func (o *SourceBase) SetGcs(v SourceGcsBase) {
	o.Gcs = &v
}

// GetS3 returns the S3 field value if set, zero value otherwise.
func (o *SourceBase) GetS3() SourceS3Base {
	if o == nil || IsNil(o.S3) {
		var ret SourceS3Base
		return ret
	}
	return *o.S3
}

// GetS3Ok returns a tuple with the S3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceBase) GetS3Ok() (*SourceS3Base, bool) {
	if o == nil || IsNil(o.S3) {
		return nil, false
	}
	return o.S3, true
}

// HasS3 returns a boolean if a field has been set.
func (o *SourceBase) HasS3() bool {
	if o != nil && !IsNil(o.S3) {
		return true
	}

	return false
}

// SetS3 gets a reference to the given SourceS3Base and assigns it to the S3 field.
func (o *SourceBase) SetS3(v SourceS3Base) {
	o.S3 = &v
}

func (o SourceBase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AzureBlobStorage) {
		toSerialize["azure_blob_storage"] = o.AzureBlobStorage
	}
	if !IsNil(o.Dynamodb) {
		toSerialize["dynamodb"] = o.Dynamodb
	}
	if !IsNil(o.Gcs) {
		toSerialize["gcs"] = o.Gcs
	}
	if !IsNil(o.S3) {
		toSerialize["s3"] = o.S3
	}
	return toSerialize, nil
}

type NullableSourceBase struct {
	value *SourceBase
	isSet bool
}

func (v NullableSourceBase) Get() *SourceBase {
	return v.value
}

func (v *NullableSourceBase) Set(val *SourceBase) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceBase) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceBase(val *SourceBase) *NullableSourceBase {
	return &NullableSourceBase{value: val, isSet: true}
}

func (v NullableSourceBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


