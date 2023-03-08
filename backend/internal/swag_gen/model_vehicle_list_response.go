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

// VehicleListResponse struct for VehicleListResponse
type VehicleListResponse struct {
	// 
	Items *[]VehicleModel `json:"items,omitempty"`
	// Count of total items with filters in place
	CountFiltered *int32 `json:"count_filtered,omitempty"`
	// Count of total items without filters in place
	CountUnfiltered *int32 `json:"count_unfiltered,omitempty"`
	// Count of items with limit in place
	CountLimited *int32 `json:"count_limited,omitempty"`
}

// NewVehicleListResponse instantiates a new VehicleListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVehicleListResponse() *VehicleListResponse {
	this := VehicleListResponse{}
	return &this
}

// NewVehicleListResponseWithDefaults instantiates a new VehicleListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVehicleListResponseWithDefaults() *VehicleListResponse {
	this := VehicleListResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *VehicleListResponse) GetItems() []VehicleModel {
	if o == nil || o.Items == nil {
		var ret []VehicleModel
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleListResponse) GetItemsOk() (*[]VehicleModel, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *VehicleListResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []VehicleModel and assigns it to the Items field.
func (o *VehicleListResponse) SetItems(v []VehicleModel) {
	o.Items = &v
}

// GetCountFiltered returns the CountFiltered field value if set, zero value otherwise.
func (o *VehicleListResponse) GetCountFiltered() int32 {
	if o == nil || o.CountFiltered == nil {
		var ret int32
		return ret
	}
	return *o.CountFiltered
}

// GetCountFilteredOk returns a tuple with the CountFiltered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleListResponse) GetCountFilteredOk() (*int32, bool) {
	if o == nil || o.CountFiltered == nil {
		return nil, false
	}
	return o.CountFiltered, true
}

// HasCountFiltered returns a boolean if a field has been set.
func (o *VehicleListResponse) HasCountFiltered() bool {
	if o != nil && o.CountFiltered != nil {
		return true
	}

	return false
}

// SetCountFiltered gets a reference to the given int32 and assigns it to the CountFiltered field.
func (o *VehicleListResponse) SetCountFiltered(v int32) {
	o.CountFiltered = &v
}

// GetCountUnfiltered returns the CountUnfiltered field value if set, zero value otherwise.
func (o *VehicleListResponse) GetCountUnfiltered() int32 {
	if o == nil || o.CountUnfiltered == nil {
		var ret int32
		return ret
	}
	return *o.CountUnfiltered
}

// GetCountUnfilteredOk returns a tuple with the CountUnfiltered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleListResponse) GetCountUnfilteredOk() (*int32, bool) {
	if o == nil || o.CountUnfiltered == nil {
		return nil, false
	}
	return o.CountUnfiltered, true
}

// HasCountUnfiltered returns a boolean if a field has been set.
func (o *VehicleListResponse) HasCountUnfiltered() bool {
	if o != nil && o.CountUnfiltered != nil {
		return true
	}

	return false
}

// SetCountUnfiltered gets a reference to the given int32 and assigns it to the CountUnfiltered field.
func (o *VehicleListResponse) SetCountUnfiltered(v int32) {
	o.CountUnfiltered = &v
}

// GetCountLimited returns the CountLimited field value if set, zero value otherwise.
func (o *VehicleListResponse) GetCountLimited() int32 {
	if o == nil || o.CountLimited == nil {
		var ret int32
		return ret
	}
	return *o.CountLimited
}

// GetCountLimitedOk returns a tuple with the CountLimited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VehicleListResponse) GetCountLimitedOk() (*int32, bool) {
	if o == nil || o.CountLimited == nil {
		return nil, false
	}
	return o.CountLimited, true
}

// HasCountLimited returns a boolean if a field has been set.
func (o *VehicleListResponse) HasCountLimited() bool {
	if o != nil && o.CountLimited != nil {
		return true
	}

	return false
}

// SetCountLimited gets a reference to the given int32 and assigns it to the CountLimited field.
func (o *VehicleListResponse) SetCountLimited(v int32) {
	o.CountLimited = &v
}

func (o VehicleListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.CountFiltered != nil {
		toSerialize["count_filtered"] = o.CountFiltered
	}
	if o.CountUnfiltered != nil {
		toSerialize["count_unfiltered"] = o.CountUnfiltered
	}
	if o.CountLimited != nil {
		toSerialize["count_limited"] = o.CountLimited
	}
	return json.Marshal(toSerialize)
}

type NullableVehicleListResponse struct {
	value *VehicleListResponse
	isSet bool
}

func (v NullableVehicleListResponse) Get() *VehicleListResponse {
	return v.value
}

func (v *NullableVehicleListResponse) Set(val *VehicleListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVehicleListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVehicleListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVehicleListResponse(val *VehicleListResponse) *NullableVehicleListResponse {
	return &NullableVehicleListResponse{value: val, isSet: true}
}

func (v NullableVehicleListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVehicleListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


