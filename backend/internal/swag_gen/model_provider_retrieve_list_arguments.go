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

// ProviderRetrieveListArguments struct for ProviderRetrieveListArguments
type ProviderRetrieveListArguments struct {
	Filters *ProviderFiltersModel `json:"filters,omitempty"`
	// 
	Limit *int64 `json:"limit,omitempty"`
	// 
	Offset *int64 `json:"offset,omitempty"`
	// Sorting Column
	SortingColumn *string `json:"sorting_column,omitempty"`
	// Sorting Direction
	SortingDirection *string `json:"sorting_direction,omitempty"`
	// 
	SearchText *string `json:"search_text,omitempty"`
	// 
	AsList *bool `json:"as_list,omitempty"`
}

// NewProviderRetrieveListArguments instantiates a new ProviderRetrieveListArguments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProviderRetrieveListArguments() *ProviderRetrieveListArguments {
	this := ProviderRetrieveListArguments{}
	return &this
}

// NewProviderRetrieveListArgumentsWithDefaults instantiates a new ProviderRetrieveListArguments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderRetrieveListArgumentsWithDefaults() *ProviderRetrieveListArguments {
	this := ProviderRetrieveListArguments{}
	return &this
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ProviderRetrieveListArguments) GetFilters() ProviderFiltersModel {
	if o == nil || o.Filters == nil {
		var ret ProviderFiltersModel
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderRetrieveListArguments) GetFiltersOk() (*ProviderFiltersModel, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ProviderRetrieveListArguments) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given ProviderFiltersModel and assigns it to the Filters field.
func (o *ProviderRetrieveListArguments) SetFilters(v ProviderFiltersModel) {
	o.Filters = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ProviderRetrieveListArguments) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderRetrieveListArguments) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ProviderRetrieveListArguments) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *ProviderRetrieveListArguments) SetLimit(v int64) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *ProviderRetrieveListArguments) GetOffset() int64 {
	if o == nil || o.Offset == nil {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderRetrieveListArguments) GetOffsetOk() (*int64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *ProviderRetrieveListArguments) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *ProviderRetrieveListArguments) SetOffset(v int64) {
	o.Offset = &v
}

// GetSortingColumn returns the SortingColumn field value if set, zero value otherwise.
func (o *ProviderRetrieveListArguments) GetSortingColumn() string {
	if o == nil || o.SortingColumn == nil {
		var ret string
		return ret
	}
	return *o.SortingColumn
}

// GetSortingColumnOk returns a tuple with the SortingColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderRetrieveListArguments) GetSortingColumnOk() (*string, bool) {
	if o == nil || o.SortingColumn == nil {
		return nil, false
	}
	return o.SortingColumn, true
}

// HasSortingColumn returns a boolean if a field has been set.
func (o *ProviderRetrieveListArguments) HasSortingColumn() bool {
	if o != nil && o.SortingColumn != nil {
		return true
	}

	return false
}

// SetSortingColumn gets a reference to the given string and assigns it to the SortingColumn field.
func (o *ProviderRetrieveListArguments) SetSortingColumn(v string) {
	o.SortingColumn = &v
}

// GetSortingDirection returns the SortingDirection field value if set, zero value otherwise.
func (o *ProviderRetrieveListArguments) GetSortingDirection() string {
	if o == nil || o.SortingDirection == nil {
		var ret string
		return ret
	}
	return *o.SortingDirection
}

// GetSortingDirectionOk returns a tuple with the SortingDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderRetrieveListArguments) GetSortingDirectionOk() (*string, bool) {
	if o == nil || o.SortingDirection == nil {
		return nil, false
	}
	return o.SortingDirection, true
}

// HasSortingDirection returns a boolean if a field has been set.
func (o *ProviderRetrieveListArguments) HasSortingDirection() bool {
	if o != nil && o.SortingDirection != nil {
		return true
	}

	return false
}

// SetSortingDirection gets a reference to the given string and assigns it to the SortingDirection field.
func (o *ProviderRetrieveListArguments) SetSortingDirection(v string) {
	o.SortingDirection = &v
}

// GetSearchText returns the SearchText field value if set, zero value otherwise.
func (o *ProviderRetrieveListArguments) GetSearchText() string {
	if o == nil || o.SearchText == nil {
		var ret string
		return ret
	}
	return *o.SearchText
}

// GetSearchTextOk returns a tuple with the SearchText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderRetrieveListArguments) GetSearchTextOk() (*string, bool) {
	if o == nil || o.SearchText == nil {
		return nil, false
	}
	return o.SearchText, true
}

// HasSearchText returns a boolean if a field has been set.
func (o *ProviderRetrieveListArguments) HasSearchText() bool {
	if o != nil && o.SearchText != nil {
		return true
	}

	return false
}

// SetSearchText gets a reference to the given string and assigns it to the SearchText field.
func (o *ProviderRetrieveListArguments) SetSearchText(v string) {
	o.SearchText = &v
}

// GetAsList returns the AsList field value if set, zero value otherwise.
func (o *ProviderRetrieveListArguments) GetAsList() bool {
	if o == nil || o.AsList == nil {
		var ret bool
		return ret
	}
	return *o.AsList
}

// GetAsListOk returns a tuple with the AsList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProviderRetrieveListArguments) GetAsListOk() (*bool, bool) {
	if o == nil || o.AsList == nil {
		return nil, false
	}
	return o.AsList, true
}

// HasAsList returns a boolean if a field has been set.
func (o *ProviderRetrieveListArguments) HasAsList() bool {
	if o != nil && o.AsList != nil {
		return true
	}

	return false
}

// SetAsList gets a reference to the given bool and assigns it to the AsList field.
func (o *ProviderRetrieveListArguments) SetAsList(v bool) {
	o.AsList = &v
}

func (o ProviderRetrieveListArguments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.SortingColumn != nil {
		toSerialize["sorting_column"] = o.SortingColumn
	}
	if o.SortingDirection != nil {
		toSerialize["sorting_direction"] = o.SortingDirection
	}
	if o.SearchText != nil {
		toSerialize["search_text"] = o.SearchText
	}
	if o.AsList != nil {
		toSerialize["as_list"] = o.AsList
	}
	return json.Marshal(toSerialize)
}

type NullableProviderRetrieveListArguments struct {
	value *ProviderRetrieveListArguments
	isSet bool
}

func (v NullableProviderRetrieveListArguments) Get() *ProviderRetrieveListArguments {
	return v.value
}

func (v *NullableProviderRetrieveListArguments) Set(val *ProviderRetrieveListArguments) {
	v.value = val
	v.isSet = true
}

func (v NullableProviderRetrieveListArguments) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderRetrieveListArguments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderRetrieveListArguments(val *ProviderRetrieveListArguments) *NullableProviderRetrieveListArguments {
	return &NullableProviderRetrieveListArguments{value: val, isSet: true}
}

func (v NullableProviderRetrieveListArguments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderRetrieveListArguments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


