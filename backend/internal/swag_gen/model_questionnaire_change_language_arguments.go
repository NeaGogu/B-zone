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

// QuestionnaireChangeLanguageArguments struct for QuestionnaireChangeLanguageArguments
type QuestionnaireChangeLanguageArguments struct {
	// Questionnaire ID
	QuestionnaireId *int32 `json:"questionnaire_id,omitempty"`
	// Language code
	LangCode *string `json:"lang_code,omitempty"`
}

// NewQuestionnaireChangeLanguageArguments instantiates a new QuestionnaireChangeLanguageArguments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuestionnaireChangeLanguageArguments() *QuestionnaireChangeLanguageArguments {
	this := QuestionnaireChangeLanguageArguments{}
	return &this
}

// NewQuestionnaireChangeLanguageArgumentsWithDefaults instantiates a new QuestionnaireChangeLanguageArguments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuestionnaireChangeLanguageArgumentsWithDefaults() *QuestionnaireChangeLanguageArguments {
	this := QuestionnaireChangeLanguageArguments{}
	return &this
}

// GetQuestionnaireId returns the QuestionnaireId field value if set, zero value otherwise.
func (o *QuestionnaireChangeLanguageArguments) GetQuestionnaireId() int32 {
	if o == nil || o.QuestionnaireId == nil {
		var ret int32
		return ret
	}
	return *o.QuestionnaireId
}

// GetQuestionnaireIdOk returns a tuple with the QuestionnaireId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireChangeLanguageArguments) GetQuestionnaireIdOk() (*int32, bool) {
	if o == nil || o.QuestionnaireId == nil {
		return nil, false
	}
	return o.QuestionnaireId, true
}

// HasQuestionnaireId returns a boolean if a field has been set.
func (o *QuestionnaireChangeLanguageArguments) HasQuestionnaireId() bool {
	if o != nil && o.QuestionnaireId != nil {
		return true
	}

	return false
}

// SetQuestionnaireId gets a reference to the given int32 and assigns it to the QuestionnaireId field.
func (o *QuestionnaireChangeLanguageArguments) SetQuestionnaireId(v int32) {
	o.QuestionnaireId = &v
}

// GetLangCode returns the LangCode field value if set, zero value otherwise.
func (o *QuestionnaireChangeLanguageArguments) GetLangCode() string {
	if o == nil || o.LangCode == nil {
		var ret string
		return ret
	}
	return *o.LangCode
}

// GetLangCodeOk returns a tuple with the LangCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireChangeLanguageArguments) GetLangCodeOk() (*string, bool) {
	if o == nil || o.LangCode == nil {
		return nil, false
	}
	return o.LangCode, true
}

// HasLangCode returns a boolean if a field has been set.
func (o *QuestionnaireChangeLanguageArguments) HasLangCode() bool {
	if o != nil && o.LangCode != nil {
		return true
	}

	return false
}

// SetLangCode gets a reference to the given string and assigns it to the LangCode field.
func (o *QuestionnaireChangeLanguageArguments) SetLangCode(v string) {
	o.LangCode = &v
}

func (o QuestionnaireChangeLanguageArguments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.QuestionnaireId != nil {
		toSerialize["questionnaire_id"] = o.QuestionnaireId
	}
	if o.LangCode != nil {
		toSerialize["lang_code"] = o.LangCode
	}
	return json.Marshal(toSerialize)
}

type NullableQuestionnaireChangeLanguageArguments struct {
	value *QuestionnaireChangeLanguageArguments
	isSet bool
}

func (v NullableQuestionnaireChangeLanguageArguments) Get() *QuestionnaireChangeLanguageArguments {
	return v.value
}

func (v *NullableQuestionnaireChangeLanguageArguments) Set(val *QuestionnaireChangeLanguageArguments) {
	v.value = val
	v.isSet = true
}

func (v NullableQuestionnaireChangeLanguageArguments) IsSet() bool {
	return v.isSet
}

func (v *NullableQuestionnaireChangeLanguageArguments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuestionnaireChangeLanguageArguments(val *QuestionnaireChangeLanguageArguments) *NullableQuestionnaireChangeLanguageArguments {
	return &NullableQuestionnaireChangeLanguageArguments{value: val, isSet: true}
}

func (v NullableQuestionnaireChangeLanguageArguments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuestionnaireChangeLanguageArguments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

