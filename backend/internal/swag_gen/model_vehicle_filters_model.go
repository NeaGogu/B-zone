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

// VehicleFiltersModel struct for VehicleFiltersModel
type VehicleFiltersModel struct {
	// Vehicle id's
	Id *[]int32 `json:"id,omitempty"`
	// Vehicle type id's
	VehicleTypeId *[]int32 `json:"vehicle_type_id,omitempty"`
	// Vehicle Link ids
	Links *[]map[string]interface{} `json:"links,omitempty"`
	// Show updated since
	UpdatedAtSince *time.Time `json:"updated_at_since,omitempty"`
	// Show updated till
	UpdatedAtTill *time.Time `json:"updated_at_till,omitempty"`
}

// NewVehicleFiltersModel instantiates a new VehicleFiltersModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVehicleFiltersModel() *VehicleFiltersModel {
	this := VehicleFiltersModel{}
	return &this
}

// NewVehicleFiltersModelWithDefaults instantiates a new VehicleFiltersModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVehicleFiltersModelWithDefaults() *VehicleFiltersModel {
	this := VehicleFiltersModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VehicleFiltersModel) GetId() []int32 {
	if o == nil || o.Id == nil {
		var ret []int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleFiltersModel) GetIdOk() (*[]int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VehicleFiltersModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given []int32 and assigns it to the Id field.
func (o *VehicleFiltersModel) SetId(v []int32) {
	o.Id = &v
}

// GetVehicleTypeId returns the VehicleTypeId field value if set, zero value otherwise.
func (o *VehicleFiltersModel) GetVehicleTypeId() []int32 {
	if o == nil || o.VehicleTypeId == nil {
		var ret []int32
		return ret
	}
	return *o.VehicleTypeId
}

// GetVehicleTypeIdOk returns a tuple with the VehicleTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleFiltersModel) GetVehicleTypeIdOk() (*[]int32, bool) {
	if o == nil || o.VehicleTypeId == nil {
		return nil, false
	}
	return o.VehicleTypeId, true
}

// HasVehicleTypeId returns a boolean if a field has been set.
func (o *VehicleFiltersModel) HasVehicleTypeId() bool {
	if o != nil && o.VehicleTypeId != nil {
		return true
	}

	return false
}

// SetVehicleTypeId gets a reference to the given []int32 and assigns it to the VehicleTypeId field.
func (o *VehicleFiltersModel) SetVehicleTypeId(v []int32) {
	o.VehicleTypeId = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *VehicleFiltersModel) GetLinks() []map[string]interface{} {
	if o == nil || o.Links == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleFiltersModel) GetLinksOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *VehicleFiltersModel) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []map[string]interface{} and assigns it to the Links field.
func (o *VehicleFiltersModel) SetLinks(v []map[string]interface{}) {
	o.Links = &v
}

// GetUpdatedAtSince returns the UpdatedAtSince field value if set, zero value otherwise.
func (o *VehicleFiltersModel) GetUpdatedAtSince() time.Time {
	if o == nil || o.UpdatedAtSince == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAtSince
}

// GetUpdatedAtSinceOk returns a tuple with the UpdatedAtSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleFiltersModel) GetUpdatedAtSinceOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAtSince == nil {
		return nil, false
	}
	return o.UpdatedAtSince, true
}

// HasUpdatedAtSince returns a boolean if a field has been set.
func (o *VehicleFiltersModel) HasUpdatedAtSince() bool {
	if o != nil && o.UpdatedAtSince != nil {
		return true
	}

	return false
}

// SetUpdatedAtSince gets a reference to the given time.Time and assigns it to the UpdatedAtSince field.
func (o *VehicleFiltersModel) SetUpdatedAtSince(v time.Time) {
	o.UpdatedAtSince = &v
}

// GetUpdatedAtTill returns the UpdatedAtTill field value if set, zero value otherwise.
func (o *VehicleFiltersModel) GetUpdatedAtTill() time.Time {
	if o == nil || o.UpdatedAtTill == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAtTill
}

// GetUpdatedAtTillOk returns a tuple with the UpdatedAtTill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleFiltersModel) GetUpdatedAtTillOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAtTill == nil {
		return nil, false
	}
	return o.UpdatedAtTill, true
}

// HasUpdatedAtTill returns a boolean if a field has been set.
func (o *VehicleFiltersModel) HasUpdatedAtTill() bool {
	if o != nil && o.UpdatedAtTill != nil {
		return true
	}

	return false
}

// SetUpdatedAtTill gets a reference to the given time.Time and assigns it to the UpdatedAtTill field.
func (o *VehicleFiltersModel) SetUpdatedAtTill(v time.Time) {
	o.UpdatedAtTill = &v
}

func (o VehicleFiltersModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.VehicleTypeId != nil {
		toSerialize["vehicle_type_id"] = o.VehicleTypeId
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.UpdatedAtSince != nil {
		toSerialize["updated_at_since"] = o.UpdatedAtSince
	}
	if o.UpdatedAtTill != nil {
		toSerialize["updated_at_till"] = o.UpdatedAtTill
	}
	return json.Marshal(toSerialize)
}

type NullableVehicleFiltersModel struct {
	value *VehicleFiltersModel
	isSet bool
}

func (v NullableVehicleFiltersModel) Get() *VehicleFiltersModel {
	return v.value
}

func (v *NullableVehicleFiltersModel) Set(val *VehicleFiltersModel) {
	v.value = val
	v.isSet = true
}

func (v NullableVehicleFiltersModel) IsSet() bool {
	return v.isSet
}

func (v *NullableVehicleFiltersModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVehicleFiltersModel(val *VehicleFiltersModel) *NullableVehicleFiltersModel {
	return &NullableVehicleFiltersModel{value: val, isSet: true}
}

func (v NullableVehicleFiltersModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVehicleFiltersModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

