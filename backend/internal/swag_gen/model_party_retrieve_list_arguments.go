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

// PartyRetrieveListArguments struct for PartyRetrieveListArguments
type PartyRetrieveListArguments struct {
	Options *PartyOptionsModel `json:"options,omitempty"`
	Filters *PartyFiltersModel `json:"filters,omitempty"`
	// 
	Limit *int64 `json:"limit,omitempty"`
	// 
	Offset *int64 `json:"offset,omitempty"`
	// 
	SearchText *string `json:"search_text,omitempty"`
	// 
	AsList *bool `json:"as_list,omitempty"`
	// 
	CountOnly *bool `json:"count_only,omitempty"`
}

// NewPartyRetrieveListArguments instantiates a new PartyRetrieveListArguments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartyRetrieveListArguments() *PartyRetrieveListArguments {
	this := PartyRetrieveListArguments{}
	return &this
}

// NewPartyRetrieveListArgumentsWithDefaults instantiates a new PartyRetrieveListArguments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartyRetrieveListArgumentsWithDefaults() *PartyRetrieveListArguments {
	this := PartyRetrieveListArguments{}
	return &this
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *PartyRetrieveListArguments) GetOptions() PartyOptionsModel {
	if o == nil || o.Options == nil {
		var ret PartyOptionsModel
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyRetrieveListArguments) GetOptionsOk() (*PartyOptionsModel, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *PartyRetrieveListArguments) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given PartyOptionsModel and assigns it to the Options field.
func (o *PartyRetrieveListArguments) SetOptions(v PartyOptionsModel) {
	o.Options = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *PartyRetrieveListArguments) GetFilters() PartyFiltersModel {
	if o == nil || o.Filters == nil {
		var ret PartyFiltersModel
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyRetrieveListArguments) GetFiltersOk() (*PartyFiltersModel, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *PartyRetrieveListArguments) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given PartyFiltersModel and assigns it to the Filters field.
func (o *PartyRetrieveListArguments) SetFilters(v PartyFiltersModel) {
	o.Filters = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *PartyRetrieveListArguments) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyRetrieveListArguments) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *PartyRetrieveListArguments) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *PartyRetrieveListArguments) SetLimit(v int64) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *PartyRetrieveListArguments) GetOffset() int64 {
	if o == nil || o.Offset == nil {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyRetrieveListArguments) GetOffsetOk() (*int64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *PartyRetrieveListArguments) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *PartyRetrieveListArguments) SetOffset(v int64) {
	o.Offset = &v
}

// GetSearchText returns the SearchText field value if set, zero value otherwise.
func (o *PartyRetrieveListArguments) GetSearchText() string {
	if o == nil || o.SearchText == nil {
		var ret string
		return ret
	}
	return *o.SearchText
}

// GetSearchTextOk returns a tuple with the SearchText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyRetrieveListArguments) GetSearchTextOk() (*string, bool) {
	if o == nil || o.SearchText == nil {
		return nil, false
	}
	return o.SearchText, true
}

// HasSearchText returns a boolean if a field has been set.
func (o *PartyRetrieveListArguments) HasSearchText() bool {
	if o != nil && o.SearchText != nil {
		return true
	}

	return false
}

// SetSearchText gets a reference to the given string and assigns it to the SearchText field.
func (o *PartyRetrieveListArguments) SetSearchText(v string) {
	o.SearchText = &v
}

// GetAsList returns the AsList field value if set, zero value otherwise.
func (o *PartyRetrieveListArguments) GetAsList() bool {
	if o == nil || o.AsList == nil {
		var ret bool
		return ret
	}
	return *o.AsList
}

// GetAsListOk returns a tuple with the AsList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyRetrieveListArguments) GetAsListOk() (*bool, bool) {
	if o == nil || o.AsList == nil {
		return nil, false
	}
	return o.AsList, true
}

// HasAsList returns a boolean if a field has been set.
func (o *PartyRetrieveListArguments) HasAsList() bool {
	if o != nil && o.AsList != nil {
		return true
	}

	return false
}

// SetAsList gets a reference to the given bool and assigns it to the AsList field.
func (o *PartyRetrieveListArguments) SetAsList(v bool) {
	o.AsList = &v
}

// GetCountOnly returns the CountOnly field value if set, zero value otherwise.
func (o *PartyRetrieveListArguments) GetCountOnly() bool {
	if o == nil || o.CountOnly == nil {
		var ret bool
		return ret
	}
	return *o.CountOnly
}

// GetCountOnlyOk returns a tuple with the CountOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PartyRetrieveListArguments) GetCountOnlyOk() (*bool, bool) {
	if o == nil || o.CountOnly == nil {
		return nil, false
	}
	return o.CountOnly, true
}

// HasCountOnly returns a boolean if a field has been set.
func (o *PartyRetrieveListArguments) HasCountOnly() bool {
	if o != nil && o.CountOnly != nil {
		return true
	}

	return false
}

// SetCountOnly gets a reference to the given bool and assigns it to the CountOnly field.
func (o *PartyRetrieveListArguments) SetCountOnly(v bool) {
	o.CountOnly = &v
}

func (o PartyRetrieveListArguments) MarshalJSON() ([]byte, error) {
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
	if o.AsList != nil {
		toSerialize["as_list"] = o.AsList
	}
	if o.CountOnly != nil {
		toSerialize["count_only"] = o.CountOnly
	}
	return json.Marshal(toSerialize)
}

type NullablePartyRetrieveListArguments struct {
	value *PartyRetrieveListArguments
	isSet bool
}

func (v NullablePartyRetrieveListArguments) Get() *PartyRetrieveListArguments {
	return v.value
}

func (v *NullablePartyRetrieveListArguments) Set(val *PartyRetrieveListArguments) {
	v.value = val
	v.isSet = true
}

func (v NullablePartyRetrieveListArguments) IsSet() bool {
	return v.isSet
}

func (v *NullablePartyRetrieveListArguments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartyRetrieveListArguments(val *PartyRetrieveListArguments) *NullablePartyRetrieveListArguments {
	return &NullablePartyRetrieveListArguments{value: val, isSet: true}
}

func (v NullablePartyRetrieveListArguments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartyRetrieveListArguments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


