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

// UpdateRecurrenceRelations struct for UpdateRecurrenceRelations
type UpdateRecurrenceRelations struct {
	// Route id's
	RouteIds *[]int32 `json:"route_ids,omitempty"`
}

// NewUpdateRecurrenceRelations instantiates a new UpdateRecurrenceRelations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRecurrenceRelations() *UpdateRecurrenceRelations {
	this := UpdateRecurrenceRelations{}
	return &this
}

// NewUpdateRecurrenceRelationsWithDefaults instantiates a new UpdateRecurrenceRelations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRecurrenceRelationsWithDefaults() *UpdateRecurrenceRelations {
	this := UpdateRecurrenceRelations{}
	return &this
}

// GetRouteIds returns the RouteIds field value if set, zero value otherwise.
func (o *UpdateRecurrenceRelations) GetRouteIds() []int32 {
	if o == nil || o.RouteIds == nil {
		var ret []int32
		return ret
	}
	return *o.RouteIds
}

// GetRouteIdsOk returns a tuple with the RouteIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRecurrenceRelations) GetRouteIdsOk() (*[]int32, bool) {
	if o == nil || o.RouteIds == nil {
		return nil, false
	}
	return o.RouteIds, true
}

// HasRouteIds returns a boolean if a field has been set.
func (o *UpdateRecurrenceRelations) HasRouteIds() bool {
	if o != nil && o.RouteIds != nil {
		return true
	}

	return false
}

// SetRouteIds gets a reference to the given []int32 and assigns it to the RouteIds field.
func (o *UpdateRecurrenceRelations) SetRouteIds(v []int32) {
	o.RouteIds = &v
}

func (o UpdateRecurrenceRelations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RouteIds != nil {
		toSerialize["route_ids"] = o.RouteIds
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateRecurrenceRelations struct {
	value *UpdateRecurrenceRelations
	isSet bool
}

func (v NullableUpdateRecurrenceRelations) Get() *UpdateRecurrenceRelations {
	return v.value
}

func (v *NullableUpdateRecurrenceRelations) Set(val *UpdateRecurrenceRelations) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRecurrenceRelations) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRecurrenceRelations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRecurrenceRelations(val *UpdateRecurrenceRelations) *NullableUpdateRecurrenceRelations {
	return &NullableUpdateRecurrenceRelations{value: val, isSet: true}
}

func (v NullableUpdateRecurrenceRelations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRecurrenceRelations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


