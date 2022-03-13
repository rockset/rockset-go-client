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

// View struct for View
type View struct {
	Path *string `json:"path,omitempty"`
	// name of the view
	Name *string `json:"name,omitempty"`
	// view description
	Description *string `json:"description,omitempty"`
	// name of the workspace
	Workspace *string `json:"workspace,omitempty"`
	// email of the creator
	CreatorEmail *string `json:"creator_email,omitempty"`
	// email of the owner
	OwnerEmail *string `json:"owner_email,omitempty"`
	// SQL query of the view
	QuerySql *string `json:"query_sql,omitempty"`
	// list of entities referenced by view. An entity can be a view, alias or collection
	Entities []string `json:"entities,omitempty"`
	// state of the view
	State *string `json:"state,omitempty"`
	// ISO-8601 date
	CreatedAt *string `json:"created_at,omitempty"`
	// ISO-8601 date
	ModifiedAt *string `json:"modified_at,omitempty"`
}

// NewView instantiates a new View object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewView() *View {
	this := View{}
	return &this
}

// NewViewWithDefaults instantiates a new View object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewWithDefaults() *View {
	this := View{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *View) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *View) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *View) SetPath(v string) {
	o.Path = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *View) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *View) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *View) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *View) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *View) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *View) SetDescription(v string) {
	o.Description = &v
}

// GetWorkspace returns the Workspace field value if set, zero value otherwise.
func (o *View) GetWorkspace() string {
	if o == nil || o.Workspace == nil {
		var ret string
		return ret
	}
	return *o.Workspace
}

// GetWorkspaceOk returns a tuple with the Workspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetWorkspaceOk() (*string, bool) {
	if o == nil || o.Workspace == nil {
		return nil, false
	}
	return o.Workspace, true
}

// HasWorkspace returns a boolean if a field has been set.
func (o *View) HasWorkspace() bool {
	if o != nil && o.Workspace != nil {
		return true
	}

	return false
}

// SetWorkspace gets a reference to the given string and assigns it to the Workspace field.
func (o *View) SetWorkspace(v string) {
	o.Workspace = &v
}

// GetCreatorEmail returns the CreatorEmail field value if set, zero value otherwise.
func (o *View) GetCreatorEmail() string {
	if o == nil || o.CreatorEmail == nil {
		var ret string
		return ret
	}
	return *o.CreatorEmail
}

// GetCreatorEmailOk returns a tuple with the CreatorEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetCreatorEmailOk() (*string, bool) {
	if o == nil || o.CreatorEmail == nil {
		return nil, false
	}
	return o.CreatorEmail, true
}

// HasCreatorEmail returns a boolean if a field has been set.
func (o *View) HasCreatorEmail() bool {
	if o != nil && o.CreatorEmail != nil {
		return true
	}

	return false
}

// SetCreatorEmail gets a reference to the given string and assigns it to the CreatorEmail field.
func (o *View) SetCreatorEmail(v string) {
	o.CreatorEmail = &v
}

// GetOwnerEmail returns the OwnerEmail field value if set, zero value otherwise.
func (o *View) GetOwnerEmail() string {
	if o == nil || o.OwnerEmail == nil {
		var ret string
		return ret
	}
	return *o.OwnerEmail
}

// GetOwnerEmailOk returns a tuple with the OwnerEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetOwnerEmailOk() (*string, bool) {
	if o == nil || o.OwnerEmail == nil {
		return nil, false
	}
	return o.OwnerEmail, true
}

// HasOwnerEmail returns a boolean if a field has been set.
func (o *View) HasOwnerEmail() bool {
	if o != nil && o.OwnerEmail != nil {
		return true
	}

	return false
}

// SetOwnerEmail gets a reference to the given string and assigns it to the OwnerEmail field.
func (o *View) SetOwnerEmail(v string) {
	o.OwnerEmail = &v
}

// GetQuerySql returns the QuerySql field value if set, zero value otherwise.
func (o *View) GetQuerySql() string {
	if o == nil || o.QuerySql == nil {
		var ret string
		return ret
	}
	return *o.QuerySql
}

// GetQuerySqlOk returns a tuple with the QuerySql field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetQuerySqlOk() (*string, bool) {
	if o == nil || o.QuerySql == nil {
		return nil, false
	}
	return o.QuerySql, true
}

// HasQuerySql returns a boolean if a field has been set.
func (o *View) HasQuerySql() bool {
	if o != nil && o.QuerySql != nil {
		return true
	}

	return false
}

// SetQuerySql gets a reference to the given string and assigns it to the QuerySql field.
func (o *View) SetQuerySql(v string) {
	o.QuerySql = &v
}

// GetEntities returns the Entities field value if set, zero value otherwise.
func (o *View) GetEntities() []string {
	if o == nil || o.Entities == nil {
		var ret []string
		return ret
	}
	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetEntitiesOk() ([]string, bool) {
	if o == nil || o.Entities == nil {
		return nil, false
	}
	return o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *View) HasEntities() bool {
	if o != nil && o.Entities != nil {
		return true
	}

	return false
}

// SetEntities gets a reference to the given []string and assigns it to the Entities field.
func (o *View) SetEntities(v []string) {
	o.Entities = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *View) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *View) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *View) SetState(v string) {
	o.State = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *View) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *View) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *View) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *View) GetModifiedAt() string {
	if o == nil || o.ModifiedAt == nil {
		var ret string
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetModifiedAtOk() (*string, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *View) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given string and assigns it to the ModifiedAt field.
func (o *View) SetModifiedAt(v string) {
	o.ModifiedAt = &v
}

func (o View) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Workspace != nil {
		toSerialize["workspace"] = o.Workspace
	}
	if o.CreatorEmail != nil {
		toSerialize["creator_email"] = o.CreatorEmail
	}
	if o.OwnerEmail != nil {
		toSerialize["owner_email"] = o.OwnerEmail
	}
	if o.QuerySql != nil {
		toSerialize["query_sql"] = o.QuerySql
	}
	if o.Entities != nil {
		toSerialize["entities"] = o.Entities
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.ModifiedAt != nil {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	return json.Marshal(toSerialize)
}

type NullableView struct {
	value *View
	isSet bool
}

func (v NullableView) Get() *View {
	return v.value
}

func (v *NullableView) Set(val *View) {
	v.value = val
	v.isSet = true
}

func (v NullableView) IsSet() bool {
	return v.isSet
}

func (v *NullableView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableView(val *View) *NullableView {
	return &NullableView{value: val, isSet: true}
}

func (v NullableView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


