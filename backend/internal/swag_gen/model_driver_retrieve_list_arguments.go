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

// DriverRetrieveListArguments struct for DriverRetrieveListArguments
type DriverRetrieveListArguments struct {
	Options *DriverOptionsModel `json:"options,omitempty"`
	Filters *DriverFiltersModel `json:"filters,omitempty"`
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

// NewDriverRetrieveListArguments instantiates a new DriverRetrieveListArguments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDriverRetrieveListArguments() *DriverRetrieveListArguments {
	this := DriverRetrieveListArguments{}
	return &this
}

// NewDriverRetrieveListArgumentsWithDefaults instantiates a new DriverRetrieveListArguments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDriverRetrieveListArgumentsWithDefaults() *DriverRetrieveListArguments {
	this := DriverRetrieveListArguments{}
	return &this
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *DriverRetrieveListArguments) GetOptions() DriverOptionsModel {
	if o == nil || o.Options == nil {
		var ret DriverOptionsModel
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriverRetrieveListArguments) GetOptionsOk() (*DriverOptionsModel, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *DriverRetrieveListArguments) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given DriverOptionsModel and assigns it to the Options field.
func (o *DriverRetrieveListArguments) SetOptions(v DriverOptionsModel) {
	o.Options = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *DriverRetrieveListArguments) GetFilters() DriverFiltersModel {
	if o == nil || o.Filters == nil {
		var ret DriverFiltersModel
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriverRetrieveListArguments) GetFiltersOk() (*DriverFiltersModel, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *DriverRetrieveListArguments) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given DriverFiltersModel and assigns it to the Filters field.
func (o *DriverRetrieveListArguments) SetFilters(v DriverFiltersModel) {
	o.Filters = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *DriverRetrieveListArguments) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriverRetrieveListArguments) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *DriverRetrieveListArguments) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *DriverRetrieveListArguments) SetLimit(v int64) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *DriverRetrieveListArguments) GetOffset() int64 {
	if o == nil || o.Offset == nil {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriverRetrieveListArguments) GetOffsetOk() (*int64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *DriverRetrieveListArguments) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *DriverRetrieveListArguments) SetOffset(v int64) {
	o.Offset = &v
}

// GetSortingColumn returns the SortingColumn field value if set, zero value otherwise.
func (o *DriverRetrieveListArguments) GetSortingColumn() string {
	if o == nil || o.SortingColumn == nil {
		var ret string
		return ret
	}
	return *o.SortingColumn
}

// GetSortingColumnOk returns a tuple with the SortingColumn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriverRetrieveListArguments) GetSortingColumnOk() (*string, bool) {
	if o == nil || o.SortingColumn == nil {
		return nil, false
	}
	return o.SortingColumn, true
}

// HasSortingColumn returns a boolean if a field has been set.
func (o *DriverRetrieveListArguments) HasSortingColumn() bool {
	if o != nil && o.SortingColumn != nil {
		return true
	}

	return false
}

// SetSortingColumn gets a reference to the given string and assigns it to the SortingColumn field.
func (o *DriverRetrieveListArguments) SetSortingColumn(v string) {
	o.SortingColumn = &v
}

// GetSortingDirection returns the SortingDirection field value if set, zero value otherwise.
func (o *DriverRetrieveListArguments) GetSortingDirection() string {
	if o == nil || o.SortingDirection == nil {
		var ret string
		return ret
	}
	return *o.SortingDirection
}

// GetSortingDirectionOk returns a tuple with the SortingDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriverRetrieveListArguments) GetSortingDirectionOk() (*string, bool) {
	if o == nil || o.SortingDirection == nil {
		return nil, false
	}
	return o.SortingDirection, true
}

// HasSortingDirection returns a boolean if a field has been set.
func (o *DriverRetrieveListArguments) HasSortingDirection() bool {
	if o != nil && o.SortingDirection != nil {
		return true
	}

	return false
}

// SetSortingDirection gets a reference to the given string and assigns it to the SortingDirection field.
func (o *DriverRetrieveListArguments) SetSortingDirection(v string) {
	o.SortingDirection = &v
}

// GetSearchText returns the SearchText field value if set, zero value otherwise.
func (o *DriverRetrieveListArguments) GetSearchText() string {
	if o == nil || o.SearchText == nil {
		var ret string
		return ret
	}
	return *o.SearchText
}

// GetSearchTextOk returns a tuple with the SearchText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriverRetrieveListArguments) GetSearchTextOk() (*string, bool) {
	if o == nil || o.SearchText == nil {
		return nil, false
	}
	return o.SearchText, true
}

// HasSearchText returns a boolean if a field has been set.
func (o *DriverRetrieveListArguments) HasSearchText() bool {
	if o != nil && o.SearchText != nil {
		return true
	}

	return false
}

// SetSearchText gets a reference to the given string and assigns it to the SearchText field.
func (o *DriverRetrieveListArguments) SetSearchText(v string) {
	o.SearchText = &v
}

// GetAsList returns the AsList field value if set, zero value otherwise.
func (o *DriverRetrieveListArguments) GetAsList() bool {
	if o == nil || o.AsList == nil {
		var ret bool
		return ret
	}
	return *o.AsList
}

// GetAsListOk returns a tuple with the AsList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DriverRetrieveListArguments) GetAsListOk() (*bool, bool) {
	if o == nil || o.AsList == nil {
		return nil, false
	}
	return o.AsList, true
}

// HasAsList returns a boolean if a field has been set.
func (o *DriverRetrieveListArguments) HasAsList() bool {
	if o != nil && o.AsList != nil {
		return true
	}

	return false
}

// SetAsList gets a reference to the given bool and assigns it to the AsList field.
func (o *DriverRetrieveListArguments) SetAsList(v bool) {
	o.AsList = &v
}

func (o DriverRetrieveListArguments) MarshalJSON() ([]byte, error) {
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

type NullableDriverRetrieveListArguments struct {
	value *DriverRetrieveListArguments
	isSet bool
}

func (v NullableDriverRetrieveListArguments) Get() *DriverRetrieveListArguments {
	return v.value
}

func (v *NullableDriverRetrieveListArguments) Set(val *DriverRetrieveListArguments) {
	v.value = val
	v.isSet = true
}

func (v NullableDriverRetrieveListArguments) IsSet() bool {
	return v.isSet
}

func (v *NullableDriverRetrieveListArguments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDriverRetrieveListArguments(val *DriverRetrieveListArguments) *NullableDriverRetrieveListArguments {
	return &NullableDriverRetrieveListArguments{value: val, isSet: true}
}

func (v NullableDriverRetrieveListArguments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDriverRetrieveListArguments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


