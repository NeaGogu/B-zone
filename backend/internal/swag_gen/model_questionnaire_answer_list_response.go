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

// QuestionnaireAnswerListResponse struct for QuestionnaireAnswerListResponse
type QuestionnaireAnswerListResponse struct {
	// 
	Items *[]QuestionnaireAnswerModel `json:"items,omitempty"`
	// Count of total items with filters in place
	CountFiltered *int32 `json:"count_filtered,omitempty"`
	// Count of total items without filters in place
	CountUnfiltered *int32 `json:"count_unfiltered,omitempty"`
	// Count of items with limit in place
	CountLimited *int32 `json:"count_limited,omitempty"`
}

// NewQuestionnaireAnswerListResponse instantiates a new QuestionnaireAnswerListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuestionnaireAnswerListResponse() *QuestionnaireAnswerListResponse {
	this := QuestionnaireAnswerListResponse{}
	return &this
}

// NewQuestionnaireAnswerListResponseWithDefaults instantiates a new QuestionnaireAnswerListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuestionnaireAnswerListResponseWithDefaults() *QuestionnaireAnswerListResponse {
	this := QuestionnaireAnswerListResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *QuestionnaireAnswerListResponse) GetItems() []QuestionnaireAnswerModel {
	if o == nil || o.Items == nil {
		var ret []QuestionnaireAnswerModel
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireAnswerListResponse) GetItemsOk() (*[]QuestionnaireAnswerModel, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *QuestionnaireAnswerListResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []QuestionnaireAnswerModel and assigns it to the Items field.
func (o *QuestionnaireAnswerListResponse) SetItems(v []QuestionnaireAnswerModel) {
	o.Items = &v
}

// GetCountFiltered returns the CountFiltered field value if set, zero value otherwise.
func (o *QuestionnaireAnswerListResponse) GetCountFiltered() int32 {
	if o == nil || o.CountFiltered == nil {
		var ret int32
		return ret
	}
	return *o.CountFiltered
}

// GetCountFilteredOk returns a tuple with the CountFiltered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireAnswerListResponse) GetCountFilteredOk() (*int32, bool) {
	if o == nil || o.CountFiltered == nil {
		return nil, false
	}
	return o.CountFiltered, true
}

// HasCountFiltered returns a boolean if a field has been set.
func (o *QuestionnaireAnswerListResponse) HasCountFiltered() bool {
	if o != nil && o.CountFiltered != nil {
		return true
	}

	return false
}

// SetCountFiltered gets a reference to the given int32 and assigns it to the CountFiltered field.
func (o *QuestionnaireAnswerListResponse) SetCountFiltered(v int32) {
	o.CountFiltered = &v
}

// GetCountUnfiltered returns the CountUnfiltered field value if set, zero value otherwise.
func (o *QuestionnaireAnswerListResponse) GetCountUnfiltered() int32 {
	if o == nil || o.CountUnfiltered == nil {
		var ret int32
		return ret
	}
	return *o.CountUnfiltered
}

// GetCountUnfilteredOk returns a tuple with the CountUnfiltered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireAnswerListResponse) GetCountUnfilteredOk() (*int32, bool) {
	if o == nil || o.CountUnfiltered == nil {
		return nil, false
	}
	return o.CountUnfiltered, true
}

// HasCountUnfiltered returns a boolean if a field has been set.
func (o *QuestionnaireAnswerListResponse) HasCountUnfiltered() bool {
	if o != nil && o.CountUnfiltered != nil {
		return true
	}

	return false
}

// SetCountUnfiltered gets a reference to the given int32 and assigns it to the CountUnfiltered field.
func (o *QuestionnaireAnswerListResponse) SetCountUnfiltered(v int32) {
	o.CountUnfiltered = &v
}

// GetCountLimited returns the CountLimited field value if set, zero value otherwise.
func (o *QuestionnaireAnswerListResponse) GetCountLimited() int32 {
	if o == nil || o.CountLimited == nil {
		var ret int32
		return ret
	}
	return *o.CountLimited
}

// GetCountLimitedOk returns a tuple with the CountLimited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireAnswerListResponse) GetCountLimitedOk() (*int32, bool) {
	if o == nil || o.CountLimited == nil {
		return nil, false
	}
	return o.CountLimited, true
}

// HasCountLimited returns a boolean if a field has been set.
func (o *QuestionnaireAnswerListResponse) HasCountLimited() bool {
	if o != nil && o.CountLimited != nil {
		return true
	}

	return false
}

// SetCountLimited gets a reference to the given int32 and assigns it to the CountLimited field.
func (o *QuestionnaireAnswerListResponse) SetCountLimited(v int32) {
	o.CountLimited = &v
}

func (o QuestionnaireAnswerListResponse) MarshalJSON() ([]byte, error) {
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

type NullableQuestionnaireAnswerListResponse struct {
	value *QuestionnaireAnswerListResponse
	isSet bool
}

func (v NullableQuestionnaireAnswerListResponse) Get() *QuestionnaireAnswerListResponse {
	return v.value
}

func (v *NullableQuestionnaireAnswerListResponse) Set(val *QuestionnaireAnswerListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQuestionnaireAnswerListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQuestionnaireAnswerListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuestionnaireAnswerListResponse(val *QuestionnaireAnswerListResponse) *NullableQuestionnaireAnswerListResponse {
	return &NullableQuestionnaireAnswerListResponse{value: val, isSet: true}
}

func (v NullableQuestionnaireAnswerListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuestionnaireAnswerListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


