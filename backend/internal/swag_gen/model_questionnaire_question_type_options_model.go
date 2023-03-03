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

// QuestionnaireQuestionTypeOptionsModel struct for QuestionnaireQuestionTypeOptionsModel
type QuestionnaireQuestionTypeOptionsModel struct {
	// 
	IncludeRecordInfo *bool `json:"include_record_info,omitempty"`
}

// NewQuestionnaireQuestionTypeOptionsModel instantiates a new QuestionnaireQuestionTypeOptionsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuestionnaireQuestionTypeOptionsModel() *QuestionnaireQuestionTypeOptionsModel {
	this := QuestionnaireQuestionTypeOptionsModel{}
	return &this
}

// NewQuestionnaireQuestionTypeOptionsModelWithDefaults instantiates a new QuestionnaireQuestionTypeOptionsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuestionnaireQuestionTypeOptionsModelWithDefaults() *QuestionnaireQuestionTypeOptionsModel {
	this := QuestionnaireQuestionTypeOptionsModel{}
	return &this
}

// GetIncludeRecordInfo returns the IncludeRecordInfo field value if set, zero value otherwise.
func (o *QuestionnaireQuestionTypeOptionsModel) GetIncludeRecordInfo() bool {
	if o == nil || o.IncludeRecordInfo == nil {
		var ret bool
		return ret
	}
	return *o.IncludeRecordInfo
}

// GetIncludeRecordInfoOk returns a tuple with the IncludeRecordInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireQuestionTypeOptionsModel) GetIncludeRecordInfoOk() (*bool, bool) {
	if o == nil || o.IncludeRecordInfo == nil {
		return nil, false
	}
	return o.IncludeRecordInfo, true
}

// HasIncludeRecordInfo returns a boolean if a field has been set.
func (o *QuestionnaireQuestionTypeOptionsModel) HasIncludeRecordInfo() bool {
	if o != nil && o.IncludeRecordInfo != nil {
		return true
	}

	return false
}

// SetIncludeRecordInfo gets a reference to the given bool and assigns it to the IncludeRecordInfo field.
func (o *QuestionnaireQuestionTypeOptionsModel) SetIncludeRecordInfo(v bool) {
	o.IncludeRecordInfo = &v
}

func (o QuestionnaireQuestionTypeOptionsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IncludeRecordInfo != nil {
		toSerialize["include_record_info"] = o.IncludeRecordInfo
	}
	return json.Marshal(toSerialize)
}

type NullableQuestionnaireQuestionTypeOptionsModel struct {
	value *QuestionnaireQuestionTypeOptionsModel
	isSet bool
}

func (v NullableQuestionnaireQuestionTypeOptionsModel) Get() *QuestionnaireQuestionTypeOptionsModel {
	return v.value
}

func (v *NullableQuestionnaireQuestionTypeOptionsModel) Set(val *QuestionnaireQuestionTypeOptionsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableQuestionnaireQuestionTypeOptionsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableQuestionnaireQuestionTypeOptionsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuestionnaireQuestionTypeOptionsModel(val *QuestionnaireQuestionTypeOptionsModel) *NullableQuestionnaireQuestionTypeOptionsModel {
	return &NullableQuestionnaireQuestionTypeOptionsModel{value: val, isSet: true}
}

func (v NullableQuestionnaireQuestionTypeOptionsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuestionnaireQuestionTypeOptionsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


