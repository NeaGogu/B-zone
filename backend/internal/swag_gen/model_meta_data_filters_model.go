/*
 * Bumbal Client Api
 *
 * Bumbal API documentation
 *
 * API version: 2.0
 * Contact: info@bumbal.eu
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// MetaDataFiltersModel struct for MetaDataFiltersModel
type MetaDataFiltersModel struct {
	// MetaData id's
	Id *[]int32 `json:"id,omitempty"`
	// MetaData names
	Name *[]string `json:"name,omitempty"`
	// MetaData object names
	ObjectTypeName *[]string `json:"object_type_name,omitempty"`
	// MetaData object id's
	ObjectId *[]int32 `json:"object_id,omitempty"`
	// Show updated since
	UpdatedAtSince *time.Time `json:"updated_at_since,omitempty"`
	// Show updated till
	UpdatedAtTill *time.Time `json:"updated_at_till,omitempty"`
	// Show create since
	CreatedAtSince *time.Time `json:"created_at_since,omitempty"`
	// Show created till
	CreatedAtTill *time.Time `json:"created_at_till,omitempty"`
}

// NewMetaDataFiltersModel instantiates a new MetaDataFiltersModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetaDataFiltersModel() *MetaDataFiltersModel {
	this := MetaDataFiltersModel{}
	return &this
}

// NewMetaDataFiltersModelWithDefaults instantiates a new MetaDataFiltersModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaDataFiltersModelWithDefaults() *MetaDataFiltersModel {
	this := MetaDataFiltersModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MetaDataFiltersModel) GetId() []int32 {
	if o == nil || o.Id == nil {
		var ret []int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaDataFiltersModel) GetIdOk() (*[]int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MetaDataFiltersModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given []int32 and assigns it to the Id field.
func (o *MetaDataFiltersModel) SetId(v []int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MetaDataFiltersModel) GetName() []string {
	if o == nil || o.Name == nil {
		var ret []string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaDataFiltersModel) GetNameOk() (*[]string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MetaDataFiltersModel) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given []string and assigns it to the Name field.
func (o *MetaDataFiltersModel) SetName(v []string) {
	o.Name = &v
}

// GetObjectTypeName returns the ObjectTypeName field value if set, zero value otherwise.
func (o *MetaDataFiltersModel) GetObjectTypeName() []string {
	if o == nil || o.ObjectTypeName == nil {
		var ret []string
		return ret
	}
	return *o.ObjectTypeName
}

// GetObjectTypeNameOk returns a tuple with the ObjectTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaDataFiltersModel) GetObjectTypeNameOk() (*[]string, bool) {
	if o == nil || o.ObjectTypeName == nil {
		return nil, false
	}
	return o.ObjectTypeName, true
}

// HasObjectTypeName returns a boolean if a field has been set.
func (o *MetaDataFiltersModel) HasObjectTypeName() bool {
	if o != nil && o.ObjectTypeName != nil {
		return true
	}

	return false
}

// SetObjectTypeName gets a reference to the given []string and assigns it to the ObjectTypeName field.
func (o *MetaDataFiltersModel) SetObjectTypeName(v []string) {
	o.ObjectTypeName = &v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *MetaDataFiltersModel) GetObjectId() []int32 {
	if o == nil || o.ObjectId == nil {
		var ret []int32
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaDataFiltersModel) GetObjectIdOk() (*[]int32, bool) {
	if o == nil || o.ObjectId == nil {
		return nil, false
	}
	return o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *MetaDataFiltersModel) HasObjectId() bool {
	if o != nil && o.ObjectId != nil {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given []int32 and assigns it to the ObjectId field.
func (o *MetaDataFiltersModel) SetObjectId(v []int32) {
	o.ObjectId = &v
}

// GetUpdatedAtSince returns the UpdatedAtSince field value if set, zero value otherwise.
func (o *MetaDataFiltersModel) GetUpdatedAtSince() time.Time {
	if o == nil || o.UpdatedAtSince == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAtSince
}

// GetUpdatedAtSinceOk returns a tuple with the UpdatedAtSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaDataFiltersModel) GetUpdatedAtSinceOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAtSince == nil {
		return nil, false
	}
	return o.UpdatedAtSince, true
}

// HasUpdatedAtSince returns a boolean if a field has been set.
func (o *MetaDataFiltersModel) HasUpdatedAtSince() bool {
	if o != nil && o.UpdatedAtSince != nil {
		return true
	}

	return false
}

// SetUpdatedAtSince gets a reference to the given time.Time and assigns it to the UpdatedAtSince field.
func (o *MetaDataFiltersModel) SetUpdatedAtSince(v time.Time) {
	o.UpdatedAtSince = &v
}

// GetUpdatedAtTill returns the UpdatedAtTill field value if set, zero value otherwise.
func (o *MetaDataFiltersModel) GetUpdatedAtTill() time.Time {
	if o == nil || o.UpdatedAtTill == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAtTill
}

// GetUpdatedAtTillOk returns a tuple with the UpdatedAtTill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaDataFiltersModel) GetUpdatedAtTillOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAtTill == nil {
		return nil, false
	}
	return o.UpdatedAtTill, true
}

// HasUpdatedAtTill returns a boolean if a field has been set.
func (o *MetaDataFiltersModel) HasUpdatedAtTill() bool {
	if o != nil && o.UpdatedAtTill != nil {
		return true
	}

	return false
}

// SetUpdatedAtTill gets a reference to the given time.Time and assigns it to the UpdatedAtTill field.
func (o *MetaDataFiltersModel) SetUpdatedAtTill(v time.Time) {
	o.UpdatedAtTill = &v
}

// GetCreatedAtSince returns the CreatedAtSince field value if set, zero value otherwise.
func (o *MetaDataFiltersModel) GetCreatedAtSince() time.Time {
	if o == nil || o.CreatedAtSince == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAtSince
}

// GetCreatedAtSinceOk returns a tuple with the CreatedAtSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaDataFiltersModel) GetCreatedAtSinceOk() (*time.Time, bool) {
	if o == nil || o.CreatedAtSince == nil {
		return nil, false
	}
	return o.CreatedAtSince, true
}

// HasCreatedAtSince returns a boolean if a field has been set.
func (o *MetaDataFiltersModel) HasCreatedAtSince() bool {
	if o != nil && o.CreatedAtSince != nil {
		return true
	}

	return false
}

// SetCreatedAtSince gets a reference to the given time.Time and assigns it to the CreatedAtSince field.
func (o *MetaDataFiltersModel) SetCreatedAtSince(v time.Time) {
	o.CreatedAtSince = &v
}

// GetCreatedAtTill returns the CreatedAtTill field value if set, zero value otherwise.
func (o *MetaDataFiltersModel) GetCreatedAtTill() time.Time {
	if o == nil || o.CreatedAtTill == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAtTill
}

// GetCreatedAtTillOk returns a tuple with the CreatedAtTill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetaDataFiltersModel) GetCreatedAtTillOk() (*time.Time, bool) {
	if o == nil || o.CreatedAtTill == nil {
		return nil, false
	}
	return o.CreatedAtTill, true
}

// HasCreatedAtTill returns a boolean if a field has been set.
func (o *MetaDataFiltersModel) HasCreatedAtTill() bool {
	if o != nil && o.CreatedAtTill != nil {
		return true
	}

	return false
}

// SetCreatedAtTill gets a reference to the given time.Time and assigns it to the CreatedAtTill field.
func (o *MetaDataFiltersModel) SetCreatedAtTill(v time.Time) {
	o.CreatedAtTill = &v
}

func (o MetaDataFiltersModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ObjectTypeName != nil {
		toSerialize["object_type_name"] = o.ObjectTypeName
	}
	if o.ObjectId != nil {
		toSerialize["object_id"] = o.ObjectId
	}
	if o.UpdatedAtSince != nil {
		toSerialize["updated_at_since"] = o.UpdatedAtSince
	}
	if o.UpdatedAtTill != nil {
		toSerialize["updated_at_till"] = o.UpdatedAtTill
	}
	if o.CreatedAtSince != nil {
		toSerialize["created_at_since"] = o.CreatedAtSince
	}
	if o.CreatedAtTill != nil {
		toSerialize["created_at_till"] = o.CreatedAtTill
	}
	return json.Marshal(toSerialize)
}

type NullableMetaDataFiltersModel struct {
	value *MetaDataFiltersModel
	isSet bool
}

func (v NullableMetaDataFiltersModel) Get() *MetaDataFiltersModel {
	return v.value
}

func (v *NullableMetaDataFiltersModel) Set(val *MetaDataFiltersModel) {
	v.value = val
	v.isSet = true
}

func (v NullableMetaDataFiltersModel) IsSet() bool {
	return v.isSet
}

func (v *NullableMetaDataFiltersModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetaDataFiltersModel(val *MetaDataFiltersModel) *NullableMetaDataFiltersModel {
	return &NullableMetaDataFiltersModel{value: val, isSet: true}
}

func (v NullableMetaDataFiltersModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetaDataFiltersModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

