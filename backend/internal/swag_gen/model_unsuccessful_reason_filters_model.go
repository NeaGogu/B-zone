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

// UnsuccessfulReasonFiltersModel struct for UnsuccessfulReasonFiltersModel
type UnsuccessfulReasonFiltersModel struct {
	// UnsuccessfulReason id's
	Id *[]int32 `json:"id,omitempty"`
}

// NewUnsuccessfulReasonFiltersModel instantiates a new UnsuccessfulReasonFiltersModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnsuccessfulReasonFiltersModel() *UnsuccessfulReasonFiltersModel {
	this := UnsuccessfulReasonFiltersModel{}
	return &this
}

// NewUnsuccessfulReasonFiltersModelWithDefaults instantiates a new UnsuccessfulReasonFiltersModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnsuccessfulReasonFiltersModelWithDefaults() *UnsuccessfulReasonFiltersModel {
	this := UnsuccessfulReasonFiltersModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UnsuccessfulReasonFiltersModel) GetId() []int32 {
	if o == nil || o.Id == nil {
		var ret []int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnsuccessfulReasonFiltersModel) GetIdOk() (*[]int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UnsuccessfulReasonFiltersModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given []int32 and assigns it to the Id field.
func (o *UnsuccessfulReasonFiltersModel) SetId(v []int32) {
	o.Id = &v
}

func (o UnsuccessfulReasonFiltersModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableUnsuccessfulReasonFiltersModel struct {
	value *UnsuccessfulReasonFiltersModel
	isSet bool
}

func (v NullableUnsuccessfulReasonFiltersModel) Get() *UnsuccessfulReasonFiltersModel {
	return v.value
}

func (v *NullableUnsuccessfulReasonFiltersModel) Set(val *UnsuccessfulReasonFiltersModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUnsuccessfulReasonFiltersModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUnsuccessfulReasonFiltersModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnsuccessfulReasonFiltersModel(val *UnsuccessfulReasonFiltersModel) *NullableUnsuccessfulReasonFiltersModel {
	return &NullableUnsuccessfulReasonFiltersModel{value: val, isSet: true}
}

func (v NullableUnsuccessfulReasonFiltersModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnsuccessfulReasonFiltersModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


