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

// SqlExpression struct for SqlExpression
type SqlExpression struct {
	// The name of a sql function.
	Sql *string `json:"sql,omitempty"`
}

// NewSqlExpression instantiates a new SqlExpression object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSqlExpression() *SqlExpression {
	this := SqlExpression{}
	return &this
}

// NewSqlExpressionWithDefaults instantiates a new SqlExpression object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSqlExpressionWithDefaults() *SqlExpression {
	this := SqlExpression{}
	return &this
}

// GetSql returns the Sql field value if set, zero value otherwise.
func (o *SqlExpression) GetSql() string {
	if o == nil || o.Sql == nil {
		var ret string
		return ret
	}
	return *o.Sql
}

// GetSqlOk returns a tuple with the Sql field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlExpression) GetSqlOk() (*string, bool) {
	if o == nil || o.Sql == nil {
		return nil, false
	}
	return o.Sql, true
}

// HasSql returns a boolean if a field has been set.
func (o *SqlExpression) HasSql() bool {
	if o != nil && o.Sql != nil {
		return true
	}

	return false
}

// SetSql gets a reference to the given string and assigns it to the Sql field.
func (o *SqlExpression) SetSql(v string) {
	o.Sql = &v
}

func (o SqlExpression) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Sql != nil {
		toSerialize["sql"] = o.Sql
	}
	return json.Marshal(toSerialize)
}

type NullableSqlExpression struct {
	value *SqlExpression
	isSet bool
}

func (v NullableSqlExpression) Get() *SqlExpression {
	return v.value
}

func (v *NullableSqlExpression) Set(val *SqlExpression) {
	v.value = val
	v.isSet = true
}

func (v NullableSqlExpression) IsSet() bool {
	return v.isSet
}

func (v *NullableSqlExpression) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSqlExpression(val *SqlExpression) *NullableSqlExpression {
	return &NullableSqlExpression{value: val, isSet: true}
}

func (v NullableSqlExpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSqlExpression) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


