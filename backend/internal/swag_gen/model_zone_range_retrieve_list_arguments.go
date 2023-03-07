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

// ZoneRangeRetrieveListArguments struct for ZoneRangeRetrieveListArguments
type ZoneRangeRetrieveListArguments struct {
	Options *map[string]interface{} `json:"options,omitempty"`
	Filters *ZoneRangeFiltersModel `json:"filters,omitempty"`
	// 
	Limit *int64 `json:"limit,omitempty"`
	// 
	Offset *int64 `json:"offset,omitempty"`
	// 
	SearchText *string `json:"search_text,omitempty"`
}

// NewZoneRangeRetrieveListArguments instantiates a new ZoneRangeRetrieveListArguments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneRangeRetrieveListArguments() *ZoneRangeRetrieveListArguments {
	this := ZoneRangeRetrieveListArguments{}
	return &this
}

// NewZoneRangeRetrieveListArgumentsWithDefaults instantiates a new ZoneRangeRetrieveListArguments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneRangeRetrieveListArgumentsWithDefaults() *ZoneRangeRetrieveListArguments {
	this := ZoneRangeRetrieveListArguments{}
	return &this
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ZoneRangeRetrieveListArguments) GetOptions() map[string]interface{} {
	if o == nil || o.Options == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneRangeRetrieveListArguments) GetOptionsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ZoneRangeRetrieveListArguments) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given map[string]interface{} and assigns it to the Options field.
func (o *ZoneRangeRetrieveListArguments) SetOptions(v map[string]interface{}) {
	o.Options = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ZoneRangeRetrieveListArguments) GetFilters() ZoneRangeFiltersModel {
	if o == nil || o.Filters == nil {
		var ret ZoneRangeFiltersModel
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneRangeRetrieveListArguments) GetFiltersOk() (*ZoneRangeFiltersModel, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ZoneRangeRetrieveListArguments) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given ZoneRangeFiltersModel and assigns it to the Filters field.
func (o *ZoneRangeRetrieveListArguments) SetFilters(v ZoneRangeFiltersModel) {
	o.Filters = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ZoneRangeRetrieveListArguments) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneRangeRetrieveListArguments) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ZoneRangeRetrieveListArguments) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *ZoneRangeRetrieveListArguments) SetLimit(v int64) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *ZoneRangeRetrieveListArguments) GetOffset() int64 {
	if o == nil || o.Offset == nil {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneRangeRetrieveListArguments) GetOffsetOk() (*int64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *ZoneRangeRetrieveListArguments) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *ZoneRangeRetrieveListArguments) SetOffset(v int64) {
	o.Offset = &v
}

// GetSearchText returns the SearchText field value if set, zero value otherwise.
func (o *ZoneRangeRetrieveListArguments) GetSearchText() string {
	if o == nil || o.SearchText == nil {
		var ret string
		return ret
	}
	return *o.SearchText
}

// GetSearchTextOk returns a tuple with the SearchText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneRangeRetrieveListArguments) GetSearchTextOk() (*string, bool) {
	if o == nil || o.SearchText == nil {
		return nil, false
	}
	return o.SearchText, true
}

// HasSearchText returns a boolean if a field has been set.
func (o *ZoneRangeRetrieveListArguments) HasSearchText() bool {
	if o != nil && o.SearchText != nil {
		return true
	}

	return false
}

// SetSearchText gets a reference to the given string and assigns it to the SearchText field.
func (o *ZoneRangeRetrieveListArguments) SetSearchText(v string) {
	o.SearchText = &v
}

func (o ZoneRangeRetrieveListArguments) MarshalJSON() ([]byte, error) {
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

type NullableZoneRangeRetrieveListArguments struct {
	value *ZoneRangeRetrieveListArguments
	isSet bool
}

func (v NullableZoneRangeRetrieveListArguments) Get() *ZoneRangeRetrieveListArguments {
	return v.value
}

func (v *NullableZoneRangeRetrieveListArguments) Set(val *ZoneRangeRetrieveListArguments) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneRangeRetrieveListArguments) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneRangeRetrieveListArguments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneRangeRetrieveListArguments(val *ZoneRangeRetrieveListArguments) *NullableZoneRangeRetrieveListArguments {
	return &NullableZoneRangeRetrieveListArguments{value: val, isSet: true}
}

func (v NullableZoneRangeRetrieveListArguments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneRangeRetrieveListArguments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

