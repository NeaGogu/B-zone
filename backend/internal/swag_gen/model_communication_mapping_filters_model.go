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

// CommunicationMappingFiltersModel struct for CommunicationMappingFiltersModel
type CommunicationMappingFiltersModel struct {
	// Show updated since
	UpdatedAtSince *time.Time `json:"updated_at_since,omitempty"`
}

// NewCommunicationMappingFiltersModel instantiates a new CommunicationMappingFiltersModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommunicationMappingFiltersModel() *CommunicationMappingFiltersModel {
	this := CommunicationMappingFiltersModel{}
	return &this
}

// NewCommunicationMappingFiltersModelWithDefaults instantiates a new CommunicationMappingFiltersModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommunicationMappingFiltersModelWithDefaults() *CommunicationMappingFiltersModel {
	this := CommunicationMappingFiltersModel{}
	return &this
}

// GetUpdatedAtSince returns the UpdatedAtSince field value if set, zero value otherwise.
func (o *CommunicationMappingFiltersModel) GetUpdatedAtSince() time.Time {
	if o == nil || o.UpdatedAtSince == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAtSince
}

// GetUpdatedAtSinceOk returns a tuple with the UpdatedAtSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationMappingFiltersModel) GetUpdatedAtSinceOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAtSince == nil {
		return nil, false
	}
	return o.UpdatedAtSince, true
}

// HasUpdatedAtSince returns a boolean if a field has been set.
func (o *CommunicationMappingFiltersModel) HasUpdatedAtSince() bool {
	if o != nil && o.UpdatedAtSince != nil {
		return true
	}

	return false
}

// SetUpdatedAtSince gets a reference to the given time.Time and assigns it to the UpdatedAtSince field.
func (o *CommunicationMappingFiltersModel) SetUpdatedAtSince(v time.Time) {
	o.UpdatedAtSince = &v
}

func (o CommunicationMappingFiltersModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UpdatedAtSince != nil {
		toSerialize["updated_at_since"] = o.UpdatedAtSince
	}
	return json.Marshal(toSerialize)
}

type NullableCommunicationMappingFiltersModel struct {
	value *CommunicationMappingFiltersModel
	isSet bool
}

func (v NullableCommunicationMappingFiltersModel) Get() *CommunicationMappingFiltersModel {
	return v.value
}

func (v *NullableCommunicationMappingFiltersModel) Set(val *CommunicationMappingFiltersModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCommunicationMappingFiltersModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCommunicationMappingFiltersModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommunicationMappingFiltersModel(val *CommunicationMappingFiltersModel) *NullableCommunicationMappingFiltersModel {
	return &NullableCommunicationMappingFiltersModel{value: val, isSet: true}
}

func (v NullableCommunicationMappingFiltersModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommunicationMappingFiltersModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

