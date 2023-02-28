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

// SystemProviderFiltersModel struct for SystemProviderFiltersModel
type SystemProviderFiltersModel struct {
	// System Provider id's
	Id *[]int32 `json:"id,omitempty"`
	// System Provider names
	Name *[]string `json:"name,omitempty"`
}

// NewSystemProviderFiltersModel instantiates a new SystemProviderFiltersModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemProviderFiltersModel() *SystemProviderFiltersModel {
	this := SystemProviderFiltersModel{}
	return &this
}

// NewSystemProviderFiltersModelWithDefaults instantiates a new SystemProviderFiltersModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemProviderFiltersModelWithDefaults() *SystemProviderFiltersModel {
	this := SystemProviderFiltersModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SystemProviderFiltersModel) GetId() []int32 {
	if o == nil || o.Id == nil {
		var ret []int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemProviderFiltersModel) GetIdOk() (*[]int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SystemProviderFiltersModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given []int32 and assigns it to the Id field.
func (o *SystemProviderFiltersModel) SetId(v []int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SystemProviderFiltersModel) GetName() []string {
	if o == nil || o.Name == nil {
		var ret []string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemProviderFiltersModel) GetNameOk() (*[]string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SystemProviderFiltersModel) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given []string and assigns it to the Name field.
func (o *SystemProviderFiltersModel) SetName(v []string) {
	o.Name = &v
}

func (o SystemProviderFiltersModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableSystemProviderFiltersModel struct {
	value *SystemProviderFiltersModel
	isSet bool
}

func (v NullableSystemProviderFiltersModel) Get() *SystemProviderFiltersModel {
	return v.value
}

func (v *NullableSystemProviderFiltersModel) Set(val *SystemProviderFiltersModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemProviderFiltersModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemProviderFiltersModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemProviderFiltersModel(val *SystemProviderFiltersModel) *NullableSystemProviderFiltersModel {
	return &NullableSystemProviderFiltersModel{value: val, isSet: true}
}

func (v NullableSystemProviderFiltersModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemProviderFiltersModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


