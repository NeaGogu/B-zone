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

// AddTagToObjectArguments struct for AddTagToObjectArguments
type AddTagToObjectArguments struct {
	Data *AddTagToObjectDataModel `json:"data,omitempty"`
	Options *map[string]interface{} `json:"options,omitempty"`
	Filters *AddTagToObjectFiltersModel `json:"filters,omitempty"`
}

// NewAddTagToObjectArguments instantiates a new AddTagToObjectArguments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddTagToObjectArguments() *AddTagToObjectArguments {
	this := AddTagToObjectArguments{}
	return &this
}

// NewAddTagToObjectArgumentsWithDefaults instantiates a new AddTagToObjectArguments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddTagToObjectArgumentsWithDefaults() *AddTagToObjectArguments {
	this := AddTagToObjectArguments{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AddTagToObjectArguments) GetData() AddTagToObjectDataModel {
	if o == nil || o.Data == nil {
		var ret AddTagToObjectDataModel
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTagToObjectArguments) GetDataOk() (*AddTagToObjectDataModel, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AddTagToObjectArguments) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given AddTagToObjectDataModel and assigns it to the Data field.
func (o *AddTagToObjectArguments) SetData(v AddTagToObjectDataModel) {
	o.Data = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *AddTagToObjectArguments) GetOptions() map[string]interface{} {
	if o == nil || o.Options == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTagToObjectArguments) GetOptionsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *AddTagToObjectArguments) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given map[string]interface{} and assigns it to the Options field.
func (o *AddTagToObjectArguments) SetOptions(v map[string]interface{}) {
	o.Options = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *AddTagToObjectArguments) GetFilters() AddTagToObjectFiltersModel {
	if o == nil || o.Filters == nil {
		var ret AddTagToObjectFiltersModel
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTagToObjectArguments) GetFiltersOk() (*AddTagToObjectFiltersModel, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *AddTagToObjectArguments) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given AddTagToObjectFiltersModel and assigns it to the Filters field.
func (o *AddTagToObjectArguments) SetFilters(v AddTagToObjectFiltersModel) {
	o.Filters = &v
}

func (o AddTagToObjectArguments) MarshalJSON() ([]byte, error) {
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

type NullableAddTagToObjectArguments struct {
	value *AddTagToObjectArguments
	isSet bool
}

func (v NullableAddTagToObjectArguments) Get() *AddTagToObjectArguments {
	return v.value
}

func (v *NullableAddTagToObjectArguments) Set(val *AddTagToObjectArguments) {
	v.value = val
	v.isSet = true
}

func (v NullableAddTagToObjectArguments) IsSet() bool {
	return v.isSet
}

func (v *NullableAddTagToObjectArguments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddTagToObjectArguments(val *AddTagToObjectArguments) *NullableAddTagToObjectArguments {
	return &NullableAddTagToObjectArguments{value: val, isSet: true}
}

func (v NullableAddTagToObjectArguments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddTagToObjectArguments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


