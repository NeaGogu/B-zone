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

// SettingsRetrieveListArguments struct for SettingsRetrieveListArguments
type SettingsRetrieveListArguments struct {
	Options *SettingsOptionsModel `json:"options,omitempty"`
	Filters *SettingsFiltersModel `json:"filters,omitempty"`
	// 
	Limit *int64 `json:"limit,omitempty"`
	// 
	Offset *int64 `json:"offset,omitempty"`
	// 
	SearchText *string `json:"search_text,omitempty"`
}

// NewSettingsRetrieveListArguments instantiates a new SettingsRetrieveListArguments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingsRetrieveListArguments() *SettingsRetrieveListArguments {
	this := SettingsRetrieveListArguments{}
	return &this
}

// NewSettingsRetrieveListArgumentsWithDefaults instantiates a new SettingsRetrieveListArguments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingsRetrieveListArgumentsWithDefaults() *SettingsRetrieveListArguments {
	this := SettingsRetrieveListArguments{}
	return &this
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *SettingsRetrieveListArguments) GetOptions() SettingsOptionsModel {
	if o == nil || o.Options == nil {
		var ret SettingsOptionsModel
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsRetrieveListArguments) GetOptionsOk() (*SettingsOptionsModel, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *SettingsRetrieveListArguments) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given SettingsOptionsModel and assigns it to the Options field.
func (o *SettingsRetrieveListArguments) SetOptions(v SettingsOptionsModel) {
	o.Options = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *SettingsRetrieveListArguments) GetFilters() SettingsFiltersModel {
	if o == nil || o.Filters == nil {
		var ret SettingsFiltersModel
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsRetrieveListArguments) GetFiltersOk() (*SettingsFiltersModel, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *SettingsRetrieveListArguments) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given SettingsFiltersModel and assigns it to the Filters field.
func (o *SettingsRetrieveListArguments) SetFilters(v SettingsFiltersModel) {
	o.Filters = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *SettingsRetrieveListArguments) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsRetrieveListArguments) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *SettingsRetrieveListArguments) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *SettingsRetrieveListArguments) SetLimit(v int64) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *SettingsRetrieveListArguments) GetOffset() int64 {
	if o == nil || o.Offset == nil {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsRetrieveListArguments) GetOffsetOk() (*int64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *SettingsRetrieveListArguments) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *SettingsRetrieveListArguments) SetOffset(v int64) {
	o.Offset = &v
}

// GetSearchText returns the SearchText field value if set, zero value otherwise.
func (o *SettingsRetrieveListArguments) GetSearchText() string {
	if o == nil || o.SearchText == nil {
		var ret string
		return ret
	}
	return *o.SearchText
}

// GetSearchTextOk returns a tuple with the SearchText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SettingsRetrieveListArguments) GetSearchTextOk() (*string, bool) {
	if o == nil || o.SearchText == nil {
		return nil, false
	}
	return o.SearchText, true
}

// HasSearchText returns a boolean if a field has been set.
func (o *SettingsRetrieveListArguments) HasSearchText() bool {
	if o != nil && o.SearchText != nil {
		return true
	}

	return false
}

// SetSearchText gets a reference to the given string and assigns it to the SearchText field.
func (o *SettingsRetrieveListArguments) SetSearchText(v string) {
	o.SearchText = &v
}

func (o SettingsRetrieveListArguments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.SearchText != nil {
		toSerialize["search_text"] = o.SearchText
	}
	return json.Marshal(toSerialize)
}

type NullableSettingsRetrieveListArguments struct {
	value *SettingsRetrieveListArguments
	isSet bool
}

func (v NullableSettingsRetrieveListArguments) Get() *SettingsRetrieveListArguments {
	return v.value
}

func (v *NullableSettingsRetrieveListArguments) Set(val *SettingsRetrieveListArguments) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingsRetrieveListArguments) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingsRetrieveListArguments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingsRetrieveListArguments(val *SettingsRetrieveListArguments) *NullableSettingsRetrieveListArguments {
	return &NullableSettingsRetrieveListArguments{value: val, isSet: true}
}

func (v NullableSettingsRetrieveListArguments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingsRetrieveListArguments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

