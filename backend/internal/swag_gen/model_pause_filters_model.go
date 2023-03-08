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

// PauseFiltersModel struct for PauseFiltersModel
type PauseFiltersModel struct {
	// Unique Identifier(s)
	Id *[]int32 `json:"id,omitempty"`
	// Pause name
	Name *[]string `json:"name,omitempty"`
	// Show updated since
	UpdatedAtSince *time.Time `json:"updated_at_since,omitempty"`
	// Show updated till
	UpdatedAtTill *time.Time `json:"updated_at_till,omitempty"`
}

// NewPauseFiltersModel instantiates a new PauseFiltersModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPauseFiltersModel() *PauseFiltersModel {
	this := PauseFiltersModel{}
	return &this
}

// NewPauseFiltersModelWithDefaults instantiates a new PauseFiltersModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPauseFiltersModelWithDefaults() *PauseFiltersModel {
	this := PauseFiltersModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PauseFiltersModel) GetId() []int32 {
	if o == nil || o.Id == nil {
		var ret []int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PauseFiltersModel) GetIdOk() (*[]int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PauseFiltersModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given []int32 and assigns it to the Id field.
func (o *PauseFiltersModel) SetId(v []int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PauseFiltersModel) GetName() []string {
	if o == nil || o.Name == nil {
		var ret []string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PauseFiltersModel) GetNameOk() (*[]string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PauseFiltersModel) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given []string and assigns it to the Name field.
func (o *PauseFiltersModel) SetName(v []string) {
	o.Name = &v
}

// GetUpdatedAtSince returns the UpdatedAtSince field value if set, zero value otherwise.
func (o *PauseFiltersModel) GetUpdatedAtSince() time.Time {
	if o == nil || o.UpdatedAtSince == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAtSince
}

// GetUpdatedAtSinceOk returns a tuple with the UpdatedAtSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PauseFiltersModel) GetUpdatedAtSinceOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAtSince == nil {
		return nil, false
	}
	return o.UpdatedAtSince, true
}

// HasUpdatedAtSince returns a boolean if a field has been set.
func (o *PauseFiltersModel) HasUpdatedAtSince() bool {
	if o != nil && o.UpdatedAtSince != nil {
		return true
	}

	return false
}

// SetUpdatedAtSince gets a reference to the given time.Time and assigns it to the UpdatedAtSince field.
func (o *PauseFiltersModel) SetUpdatedAtSince(v time.Time) {
	o.UpdatedAtSince = &v
}

// GetUpdatedAtTill returns the UpdatedAtTill field value if set, zero value otherwise.
func (o *PauseFiltersModel) GetUpdatedAtTill() time.Time {
	if o == nil || o.UpdatedAtTill == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAtTill
}

// GetUpdatedAtTillOk returns a tuple with the UpdatedAtTill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PauseFiltersModel) GetUpdatedAtTillOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAtTill == nil {
		return nil, false
	}
	return o.UpdatedAtTill, true
}

// HasUpdatedAtTill returns a boolean if a field has been set.
func (o *PauseFiltersModel) HasUpdatedAtTill() bool {
	if o != nil && o.UpdatedAtTill != nil {
		return true
	}

	return false
}

// SetUpdatedAtTill gets a reference to the given time.Time and assigns it to the UpdatedAtTill field.
func (o *PauseFiltersModel) SetUpdatedAtTill(v time.Time) {
	o.UpdatedAtTill = &v
}

func (o PauseFiltersModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.UpdatedAtSince != nil {
		toSerialize["updated_at_since"] = o.UpdatedAtSince
	}
	if o.UpdatedAtTill != nil {
		toSerialize["updated_at_till"] = o.UpdatedAtTill
	}
	return json.Marshal(toSerialize)
}

type NullablePauseFiltersModel struct {
	value *PauseFiltersModel
	isSet bool
}

func (v NullablePauseFiltersModel) Get() *PauseFiltersModel {
	return v.value
}

func (v *NullablePauseFiltersModel) Set(val *PauseFiltersModel) {
	v.value = val
	v.isSet = true
}

func (v NullablePauseFiltersModel) IsSet() bool {
	return v.isSet
}

func (v *NullablePauseFiltersModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePauseFiltersModel(val *PauseFiltersModel) *NullablePauseFiltersModel {
	return &NullablePauseFiltersModel{value: val, isSet: true}
}

func (v NullablePauseFiltersModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePauseFiltersModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


