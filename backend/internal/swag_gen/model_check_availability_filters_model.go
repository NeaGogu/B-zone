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

// CheckAvailabilityFiltersModel struct for CheckAvailabilityFiltersModel
type CheckAvailabilityFiltersModel struct {
	// Route id
	RouteId *[]int32 `json:"route_id,omitempty"`
	// 
	DateFrom *string `json:"date_from,omitempty"`
	// 
	DateTo *string `json:"date_to,omitempty"`
	// Availability check will continue to analyze days until there is no availability in the system anymore or the number of days with available time slots has reached the max_nr_of_days_with_availability
	MaxNrOfDaysWithAvailability *int32 `json:"max_nr_of_days_with_availability,omitempty"`
}

// NewCheckAvailabilityFiltersModel instantiates a new CheckAvailabilityFiltersModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckAvailabilityFiltersModel() *CheckAvailabilityFiltersModel {
	this := CheckAvailabilityFiltersModel{}
	return &this
}

// NewCheckAvailabilityFiltersModelWithDefaults instantiates a new CheckAvailabilityFiltersModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckAvailabilityFiltersModelWithDefaults() *CheckAvailabilityFiltersModel {
	this := CheckAvailabilityFiltersModel{}
	return &this
}

// GetRouteId returns the RouteId field value if set, zero value otherwise.
func (o *CheckAvailabilityFiltersModel) GetRouteId() []int32 {
	if o == nil || o.RouteId == nil {
		var ret []int32
		return ret
	}
	return *o.RouteId
}

// GetRouteIdOk returns a tuple with the RouteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckAvailabilityFiltersModel) GetRouteIdOk() (*[]int32, bool) {
	if o == nil || o.RouteId == nil {
		return nil, false
	}
	return o.RouteId, true
}

// HasRouteId returns a boolean if a field has been set.
func (o *CheckAvailabilityFiltersModel) HasRouteId() bool {
	if o != nil && o.RouteId != nil {
		return true
	}

	return false
}

// SetRouteId gets a reference to the given []int32 and assigns it to the RouteId field.
func (o *CheckAvailabilityFiltersModel) SetRouteId(v []int32) {
	o.RouteId = &v
}

// GetDateFrom returns the DateFrom field value if set, zero value otherwise.
func (o *CheckAvailabilityFiltersModel) GetDateFrom() string {
	if o == nil || o.DateFrom == nil {
		var ret string
		return ret
	}
	return *o.DateFrom
}

// GetDateFromOk returns a tuple with the DateFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckAvailabilityFiltersModel) GetDateFromOk() (*string, bool) {
	if o == nil || o.DateFrom == nil {
		return nil, false
	}
	return o.DateFrom, true
}

// HasDateFrom returns a boolean if a field has been set.
func (o *CheckAvailabilityFiltersModel) HasDateFrom() bool {
	if o != nil && o.DateFrom != nil {
		return true
	}

	return false
}

// SetDateFrom gets a reference to the given string and assigns it to the DateFrom field.
func (o *CheckAvailabilityFiltersModel) SetDateFrom(v string) {
	o.DateFrom = &v
}

// GetDateTo returns the DateTo field value if set, zero value otherwise.
func (o *CheckAvailabilityFiltersModel) GetDateTo() string {
	if o == nil || o.DateTo == nil {
		var ret string
		return ret
	}
	return *o.DateTo
}

// GetDateToOk returns a tuple with the DateTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckAvailabilityFiltersModel) GetDateToOk() (*string, bool) {
	if o == nil || o.DateTo == nil {
		return nil, false
	}
	return o.DateTo, true
}

// HasDateTo returns a boolean if a field has been set.
func (o *CheckAvailabilityFiltersModel) HasDateTo() bool {
	if o != nil && o.DateTo != nil {
		return true
	}

	return false
}

// SetDateTo gets a reference to the given string and assigns it to the DateTo field.
func (o *CheckAvailabilityFiltersModel) SetDateTo(v string) {
	o.DateTo = &v
}

// GetMaxNrOfDaysWithAvailability returns the MaxNrOfDaysWithAvailability field value if set, zero value otherwise.
func (o *CheckAvailabilityFiltersModel) GetMaxNrOfDaysWithAvailability() int32 {
	if o == nil || o.MaxNrOfDaysWithAvailability == nil {
		var ret int32
		return ret
	}
	return *o.MaxNrOfDaysWithAvailability
}

// GetMaxNrOfDaysWithAvailabilityOk returns a tuple with the MaxNrOfDaysWithAvailability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckAvailabilityFiltersModel) GetMaxNrOfDaysWithAvailabilityOk() (*int32, bool) {
	if o == nil || o.MaxNrOfDaysWithAvailability == nil {
		return nil, false
	}
	return o.MaxNrOfDaysWithAvailability, true
}

// HasMaxNrOfDaysWithAvailability returns a boolean if a field has been set.
func (o *CheckAvailabilityFiltersModel) HasMaxNrOfDaysWithAvailability() bool {
	if o != nil && o.MaxNrOfDaysWithAvailability != nil {
		return true
	}

	return false
}

// SetMaxNrOfDaysWithAvailability gets a reference to the given int32 and assigns it to the MaxNrOfDaysWithAvailability field.
func (o *CheckAvailabilityFiltersModel) SetMaxNrOfDaysWithAvailability(v int32) {
	o.MaxNrOfDaysWithAvailability = &v
}

func (o CheckAvailabilityFiltersModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RouteId != nil {
		toSerialize["route_id"] = o.RouteId
	}
	if o.DateFrom != nil {
		toSerialize["date_from"] = o.DateFrom
	}
	if o.DateTo != nil {
		toSerialize["date_to"] = o.DateTo
	}
	if o.MaxNrOfDaysWithAvailability != nil {
		toSerialize["max_nr_of_days_with_availability"] = o.MaxNrOfDaysWithAvailability
	}
	return json.Marshal(toSerialize)
}

type NullableCheckAvailabilityFiltersModel struct {
	value *CheckAvailabilityFiltersModel
	isSet bool
}

func (v NullableCheckAvailabilityFiltersModel) Get() *CheckAvailabilityFiltersModel {
	return v.value
}

func (v *NullableCheckAvailabilityFiltersModel) Set(val *CheckAvailabilityFiltersModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckAvailabilityFiltersModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckAvailabilityFiltersModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckAvailabilityFiltersModel(val *CheckAvailabilityFiltersModel) *NullableCheckAvailabilityFiltersModel {
	return &NullableCheckAvailabilityFiltersModel{value: val, isSet: true}
}

func (v NullableCheckAvailabilityFiltersModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckAvailabilityFiltersModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

