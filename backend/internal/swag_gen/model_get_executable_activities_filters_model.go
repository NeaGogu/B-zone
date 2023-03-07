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

// GetExecutableActivitiesFiltersModel struct for GetExecutableActivitiesFiltersModel
type GetExecutableActivitiesFiltersModel struct {
	// Unique Identifier for route
	RouteId *int32 `json:"route_id,omitempty"`
}

// NewGetExecutableActivitiesFiltersModel instantiates a new GetExecutableActivitiesFiltersModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetExecutableActivitiesFiltersModel() *GetExecutableActivitiesFiltersModel {
	this := GetExecutableActivitiesFiltersModel{}
	return &this
}

// NewGetExecutableActivitiesFiltersModelWithDefaults instantiates a new GetExecutableActivitiesFiltersModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetExecutableActivitiesFiltersModelWithDefaults() *GetExecutableActivitiesFiltersModel {
	this := GetExecutableActivitiesFiltersModel{}
	return &this
}

// GetRouteId returns the RouteId field value if set, zero value otherwise.
func (o *GetExecutableActivitiesFiltersModel) GetRouteId() int32 {
	if o == nil || o.RouteId == nil {
		var ret int32
		return ret
	}
	return *o.RouteId
}

// GetRouteIdOk returns a tuple with the RouteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExecutableActivitiesFiltersModel) GetRouteIdOk() (*int32, bool) {
	if o == nil || o.RouteId == nil {
		return nil, false
	}
	return o.RouteId, true
}

// HasRouteId returns a boolean if a field has been set.
func (o *GetExecutableActivitiesFiltersModel) HasRouteId() bool {
	if o != nil && o.RouteId != nil {
		return true
	}

	return false
}

// SetRouteId gets a reference to the given int32 and assigns it to the RouteId field.
func (o *GetExecutableActivitiesFiltersModel) SetRouteId(v int32) {
	o.RouteId = &v
}

func (o GetExecutableActivitiesFiltersModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RouteId != nil {
		toSerialize["route_id"] = o.RouteId
	}
	return json.Marshal(toSerialize)
}

type NullableGetExecutableActivitiesFiltersModel struct {
	value *GetExecutableActivitiesFiltersModel
	isSet bool
}

func (v NullableGetExecutableActivitiesFiltersModel) Get() *GetExecutableActivitiesFiltersModel {
	return v.value
}

func (v *NullableGetExecutableActivitiesFiltersModel) Set(val *GetExecutableActivitiesFiltersModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGetExecutableActivitiesFiltersModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGetExecutableActivitiesFiltersModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetExecutableActivitiesFiltersModel(val *GetExecutableActivitiesFiltersModel) *NullableGetExecutableActivitiesFiltersModel {
	return &NullableGetExecutableActivitiesFiltersModel{value: val, isSet: true}
}

func (v NullableGetExecutableActivitiesFiltersModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetExecutableActivitiesFiltersModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

