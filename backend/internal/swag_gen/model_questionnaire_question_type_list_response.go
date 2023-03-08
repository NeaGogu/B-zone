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

// QuestionnaireQuestionTypeListResponse struct for QuestionnaireQuestionTypeListResponse
type QuestionnaireQuestionTypeListResponse struct {
	// 
	Items *[]QuestionnaireQuestionTypeModel `json:"items,omitempty"`
	// Count of total items with filters in place
	CountFiltered *int32 `json:"count_filtered,omitempty"`
	// Count of total items without filters in place
	CountUnfiltered *int32 `json:"count_unfiltered,omitempty"`
	// Count of items with limit in place
	CountLimited *int32 `json:"count_limited,omitempty"`
}

// NewQuestionnaireQuestionTypeListResponse instantiates a new QuestionnaireQuestionTypeListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuestionnaireQuestionTypeListResponse() *QuestionnaireQuestionTypeListResponse {
	this := QuestionnaireQuestionTypeListResponse{}
	return &this
}

// NewQuestionnaireQuestionTypeListResponseWithDefaults instantiates a new QuestionnaireQuestionTypeListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuestionnaireQuestionTypeListResponseWithDefaults() *QuestionnaireQuestionTypeListResponse {
	this := QuestionnaireQuestionTypeListResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *QuestionnaireQuestionTypeListResponse) GetItems() []QuestionnaireQuestionTypeModel {
	if o == nil || o.Items == nil {
		var ret []QuestionnaireQuestionTypeModel
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireQuestionTypeListResponse) GetItemsOk() (*[]QuestionnaireQuestionTypeModel, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *QuestionnaireQuestionTypeListResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []QuestionnaireQuestionTypeModel and assigns it to the Items field.
func (o *QuestionnaireQuestionTypeListResponse) SetItems(v []QuestionnaireQuestionTypeModel) {
	o.Items = &v
}

// GetCountFiltered returns the CountFiltered field value if set, zero value otherwise.
func (o *QuestionnaireQuestionTypeListResponse) GetCountFiltered() int32 {
	if o == nil || o.CountFiltered == nil {
		var ret int32
		return ret
	}
	return *o.CountFiltered
}

// GetCountFilteredOk returns a tuple with the CountFiltered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireQuestionTypeListResponse) GetCountFilteredOk() (*int32, bool) {
	if o == nil || o.CountFiltered == nil {
		return nil, false
	}
	return o.CountFiltered, true
}

// HasCountFiltered returns a boolean if a field has been set.
func (o *QuestionnaireQuestionTypeListResponse) HasCountFiltered() bool {
	if o != nil && o.CountFiltered != nil {
		return true
	}

	return false
}

// SetCountFiltered gets a reference to the given int32 and assigns it to the CountFiltered field.
func (o *QuestionnaireQuestionTypeListResponse) SetCountFiltered(v int32) {
	o.CountFiltered = &v
}

// GetCountUnfiltered returns the CountUnfiltered field value if set, zero value otherwise.
func (o *QuestionnaireQuestionTypeListResponse) GetCountUnfiltered() int32 {
	if o == nil || o.CountUnfiltered == nil {
		var ret int32
		return ret
	}
	return *o.CountUnfiltered
}

// GetCountUnfilteredOk returns a tuple with the CountUnfiltered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireQuestionTypeListResponse) GetCountUnfilteredOk() (*int32, bool) {
	if o == nil || o.CountUnfiltered == nil {
		return nil, false
	}
	return o.CountUnfiltered, true
}

// HasCountUnfiltered returns a boolean if a field has been set.
func (o *QuestionnaireQuestionTypeListResponse) HasCountUnfiltered() bool {
	if o != nil && o.CountUnfiltered != nil {
		return true
	}

	return false
}

// SetCountUnfiltered gets a reference to the given int32 and assigns it to the CountUnfiltered field.
func (o *QuestionnaireQuestionTypeListResponse) SetCountUnfiltered(v int32) {
	o.CountUnfiltered = &v
}

// GetCountLimited returns the CountLimited field value if set, zero value otherwise.
func (o *QuestionnaireQuestionTypeListResponse) GetCountLimited() int32 {
	if o == nil || o.CountLimited == nil {
		var ret int32
		return ret
	}
	return *o.CountLimited
}

// GetCountLimitedOk returns a tuple with the CountLimited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireQuestionTypeListResponse) GetCountLimitedOk() (*int32, bool) {
	if o == nil || o.CountLimited == nil {
		return nil, false
	}
	return o.CountLimited, true
}

// HasCountLimited returns a boolean if a field has been set.
func (o *QuestionnaireQuestionTypeListResponse) HasCountLimited() bool {
	if o != nil && o.CountLimited != nil {
		return true
	}

	return false
}

// SetCountLimited gets a reference to the given int32 and assigns it to the CountLimited field.
func (o *QuestionnaireQuestionTypeListResponse) SetCountLimited(v int32) {
	o.CountLimited = &v
}

func (o QuestionnaireQuestionTypeListResponse) MarshalJSON() ([]byte, error) {
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

type NullableQuestionnaireQuestionTypeListResponse struct {
	value *QuestionnaireQuestionTypeListResponse
	isSet bool
}

func (v NullableQuestionnaireQuestionTypeListResponse) Get() *QuestionnaireQuestionTypeListResponse {
	return v.value
}

func (v *NullableQuestionnaireQuestionTypeListResponse) Set(val *QuestionnaireQuestionTypeListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQuestionnaireQuestionTypeListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQuestionnaireQuestionTypeListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuestionnaireQuestionTypeListResponse(val *QuestionnaireQuestionTypeListResponse) *NullableQuestionnaireQuestionTypeListResponse {
	return &NullableQuestionnaireQuestionTypeListResponse{value: val, isSet: true}
}

func (v NullableQuestionnaireQuestionTypeListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuestionnaireQuestionTypeListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


