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

// CheckAvailabilityArguments struct for CheckAvailabilityArguments
type CheckAvailabilityArguments struct {
	Data *CheckAvailabilityDataModel `json:"data,omitempty"`
	Options *CheckAvailabilityOptionsModel `json:"options,omitempty"`
	Filters *CheckAvailabilityFiltersModel `json:"filters,omitempty"`
}

// NewCheckAvailabilityArguments instantiates a new CheckAvailabilityArguments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckAvailabilityArguments() *CheckAvailabilityArguments {
	this := CheckAvailabilityArguments{}
	return &this
}

// NewCheckAvailabilityArgumentsWithDefaults instantiates a new CheckAvailabilityArguments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckAvailabilityArgumentsWithDefaults() *CheckAvailabilityArguments {
	this := CheckAvailabilityArguments{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CheckAvailabilityArguments) GetData() CheckAvailabilityDataModel {
	if o == nil || o.Data == nil {
		var ret CheckAvailabilityDataModel
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckAvailabilityArguments) GetDataOk() (*CheckAvailabilityDataModel, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CheckAvailabilityArguments) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given CheckAvailabilityDataModel and assigns it to the Data field.
func (o *CheckAvailabilityArguments) SetData(v CheckAvailabilityDataModel) {
	o.Data = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *CheckAvailabilityArguments) GetOptions() CheckAvailabilityOptionsModel {
	if o == nil || o.Options == nil {
		var ret CheckAvailabilityOptionsModel
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckAvailabilityArguments) GetOptionsOk() (*CheckAvailabilityOptionsModel, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *CheckAvailabilityArguments) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given CheckAvailabilityOptionsModel and assigns it to the Options field.
func (o *CheckAvailabilityArguments) SetOptions(v CheckAvailabilityOptionsModel) {
	o.Options = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *CheckAvailabilityArguments) GetFilters() CheckAvailabilityFiltersModel {
	if o == nil || o.Filters == nil {
		var ret CheckAvailabilityFiltersModel
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckAvailabilityArguments) GetFiltersOk() (*CheckAvailabilityFiltersModel, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *CheckAvailabilityArguments) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given CheckAvailabilityFiltersModel and assigns it to the Filters field.
func (o *CheckAvailabilityArguments) SetFilters(v CheckAvailabilityFiltersModel) {
	o.Filters = &v
}

func (o CheckAvailabilityArguments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	return json.Marshal(toSerialize)
}

type NullableCheckAvailabilityArguments struct {
	value *CheckAvailabilityArguments
	isSet bool
}

func (v NullableCheckAvailabilityArguments) Get() *CheckAvailabilityArguments {
	return v.value
}

func (v *NullableCheckAvailabilityArguments) Set(val *CheckAvailabilityArguments) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckAvailabilityArguments) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckAvailabilityArguments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckAvailabilityArguments(val *CheckAvailabilityArguments) *NullableCheckAvailabilityArguments {
	return &NullableCheckAvailabilityArguments{value: val, isSet: true}
}

func (v NullableCheckAvailabilityArguments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckAvailabilityArguments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


