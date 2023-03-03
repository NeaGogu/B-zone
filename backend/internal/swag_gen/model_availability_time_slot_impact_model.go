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

// AvailabilityTimeSlotImpactModel struct for AvailabilityTimeSlotImpactModel
type AvailabilityTimeSlotImpactModel struct {
	// 
	Type *string `json:"type,omitempty"`
	// 
	Value *int32 `json:"value,omitempty"`
}

// NewAvailabilityTimeSlotImpactModel instantiates a new AvailabilityTimeSlotImpactModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailabilityTimeSlotImpactModel() *AvailabilityTimeSlotImpactModel {
	this := AvailabilityTimeSlotImpactModel{}
	return &this
}

// NewAvailabilityTimeSlotImpactModelWithDefaults instantiates a new AvailabilityTimeSlotImpactModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailabilityTimeSlotImpactModelWithDefaults() *AvailabilityTimeSlotImpactModel {
	this := AvailabilityTimeSlotImpactModel{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AvailabilityTimeSlotImpactModel) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityTimeSlotImpactModel) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AvailabilityTimeSlotImpactModel) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AvailabilityTimeSlotImpactModel) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AvailabilityTimeSlotImpactModel) GetValue() int32 {
	if o == nil || o.Value == nil {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvailabilityTimeSlotImpactModel) GetValueOk() (*int32, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AvailabilityTimeSlotImpactModel) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *AvailabilityTimeSlotImpactModel) SetValue(v int32) {
	o.Value = &v
}

func (o AvailabilityTimeSlotImpactModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableAvailabilityTimeSlotImpactModel struct {
	value *AvailabilityTimeSlotImpactModel
	isSet bool
}

func (v NullableAvailabilityTimeSlotImpactModel) Get() *AvailabilityTimeSlotImpactModel {
	return v.value
}

func (v *NullableAvailabilityTimeSlotImpactModel) Set(val *AvailabilityTimeSlotImpactModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailabilityTimeSlotImpactModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailabilityTimeSlotImpactModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailabilityTimeSlotImpactModel(val *AvailabilityTimeSlotImpactModel) *NullableAvailabilityTimeSlotImpactModel {
	return &NullableAvailabilityTimeSlotImpactModel{value: val, isSet: true}
}

func (v NullableAvailabilityTimeSlotImpactModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailabilityTimeSlotImpactModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


