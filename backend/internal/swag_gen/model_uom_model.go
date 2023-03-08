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
)

// UomModel struct for UomModel
type UomModel struct {
	// Unique ID
	Id *int64 `json:"id,omitempty"`
	// name of unit
	Name *string `json:"name,omitempty"`
	UomGroup *UomGroupModel `json:"uom_group,omitempty"`
}

// NewUomModel instantiates a new UomModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUomModel() *UomModel {
	this := UomModel{}
	return &this
}

// NewUomModelWithDefaults instantiates a new UomModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUomModelWithDefaults() *UomModel {
	this := UomModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UomModel) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UomModel) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UomModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *UomModel) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UomModel) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UomModel) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UomModel) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UomModel) SetName(v string) {
	o.Name = &v
}

// GetUomGroup returns the UomGroup field value if set, zero value otherwise.
func (o *UomModel) GetUomGroup() UomGroupModel {
	if o == nil || o.UomGroup == nil {
		var ret UomGroupModel
		return ret
	}
	return *o.UomGroup
}

// GetUomGroupOk returns a tuple with the UomGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UomModel) GetUomGroupOk() (*UomGroupModel, bool) {
	if o == nil || o.UomGroup == nil {
		return nil, false
	}
	return o.UomGroup, true
}

// HasUomGroup returns a boolean if a field has been set.
func (o *UomModel) HasUomGroup() bool {
	if o != nil && o.UomGroup != nil {
		return true
	}

	return false
}

// SetUomGroup gets a reference to the given UomGroupModel and assigns it to the UomGroup field.
func (o *UomModel) SetUomGroup(v UomGroupModel) {
	o.UomGroup = &v
}

func (o UomModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.UomGroup != nil {
		toSerialize["uom_group"] = o.UomGroup
	}
	return json.Marshal(toSerialize)
}

type NullableUomModel struct {
	value *UomModel
	isSet bool
}

func (v NullableUomModel) Get() *UomModel {
	return v.value
}

func (v *NullableUomModel) Set(val *UomModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUomModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUomModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUomModel(val *UomModel) *NullableUomModel {
	return &NullableUomModel{value: val, isSet: true}
}

func (v NullableUomModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUomModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


