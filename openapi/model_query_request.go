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

// QueryRequest struct for QueryRequest
type QueryRequest struct {
	AsyncOptions *AsyncQueryOptions `json:"async_options,omitempty"`
	Sql QueryRequestSql `json:"sql"`
}

// NewQueryRequest instantiates a new QueryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryRequest(sql QueryRequestSql) *QueryRequest {
	this := QueryRequest{}
	this.Sql = sql
	return &this
}

// NewQueryRequestWithDefaults instantiates a new QueryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryRequestWithDefaults() *QueryRequest {
	this := QueryRequest{}
	return &this
}

// GetAsyncOptions returns the AsyncOptions field value if set, zero value otherwise.
func (o *QueryRequest) GetAsyncOptions() AsyncQueryOptions {
	if o == nil || o.AsyncOptions == nil {
		var ret AsyncQueryOptions
		return ret
	}
	return *o.AsyncOptions
}

// GetAsyncOptionsOk returns a tuple with the AsyncOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryRequest) GetAsyncOptionsOk() (*AsyncQueryOptions, bool) {
	if o == nil || o.AsyncOptions == nil {
		return nil, false
	}
	return o.AsyncOptions, true
}

// HasAsyncOptions returns a boolean if a field has been set.
func (o *QueryRequest) HasAsyncOptions() bool {
	if o != nil && o.AsyncOptions != nil {
		return true
	}

	return false
}

// SetAsyncOptions gets a reference to the given AsyncQueryOptions and assigns it to the AsyncOptions field.
func (o *QueryRequest) SetAsyncOptions(v AsyncQueryOptions) {
	o.AsyncOptions = &v
}

// GetSql returns the Sql field value
func (o *QueryRequest) GetSql() QueryRequestSql {
	if o == nil {
		var ret QueryRequestSql
		return ret
	}

	return o.Sql
}

// GetSqlOk returns a tuple with the Sql field value
// and a boolean to check if the value has been set.
func (o *QueryRequest) GetSqlOk() (*QueryRequestSql, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Sql, true
}

// SetSql sets field value
func (o *QueryRequest) SetSql(v QueryRequestSql) {
	o.Sql = v
}

func (o QueryRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AsyncOptions != nil {
		toSerialize["async_options"] = o.AsyncOptions
	}
	if true {
		toSerialize["sql"] = o.Sql
	}
	return json.Marshal(toSerialize)
}

type NullableQueryRequest struct {
	value *QueryRequest
	isSet bool
}

func (v NullableQueryRequest) Get() *QueryRequest {
	return v.value
}

func (v *NullableQueryRequest) Set(val *QueryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryRequest(val *QueryRequest) *NullableQueryRequest {
	return &NullableQueryRequest{value: val, isSet: true}
}

func (v NullableQueryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


