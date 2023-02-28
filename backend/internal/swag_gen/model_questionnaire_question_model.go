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

// QuestionnaireQuestionModel struct for QuestionnaireQuestionModel
type QuestionnaireQuestionModel struct {
	// ID of the questionnaire id
	QuestionnaireId *int32 `json:"questionnaire_id,omitempty"`
	// ID of the questionnaire question id
	QuestionnaireQuestionId *int32 `json:"questionnaire_question_id,omitempty"`
	// order of the question inside the questionnaire
	Order *int32 `json:"order,omitempty"`
	// Question
	Question *string `json:"question,omitempty"`
	// 
	AnswerOptions *[]QuestionnaireQuestionOptionModel `json:"answer_options,omitempty"`
	// Answer
	Answer *string `json:"answer,omitempty"`
	// Chosen option as answer
	ChosenOptions *[]int32 `json:"chosen_options,omitempty"`
	// Question type name
	QuestionnaireQuestionTypeName *string `json:"questionnaire_question_type_name,omitempty"`
}

// NewQuestionnaireQuestionModel instantiates a new QuestionnaireQuestionModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuestionnaireQuestionModel() *QuestionnaireQuestionModel {
	this := QuestionnaireQuestionModel{}
	return &this
}

// NewQuestionnaireQuestionModelWithDefaults instantiates a new QuestionnaireQuestionModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuestionnaireQuestionModelWithDefaults() *QuestionnaireQuestionModel {
	this := QuestionnaireQuestionModel{}
	return &this
}

// GetQuestionnaireId returns the QuestionnaireId field value if set, zero value otherwise.
func (o *QuestionnaireQuestionModel) GetQuestionnaireId() int32 {
	if o == nil || o.QuestionnaireId == nil {
		var ret int32
		return ret
	}
	return *o.QuestionnaireId
}

// GetQuestionnaireIdOk returns a tuple with the QuestionnaireId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireQuestionModel) GetQuestionnaireIdOk() (*int32, bool) {
	if o == nil || o.QuestionnaireId == nil {
		return nil, false
	}
	return o.QuestionnaireId, true
}

// HasQuestionnaireId returns a boolean if a field has been set.
func (o *QuestionnaireQuestionModel) HasQuestionnaireId() bool {
	if o != nil && o.QuestionnaireId != nil {
		return true
	}

	return false
}

// SetQuestionnaireId gets a reference to the given int32 and assigns it to the QuestionnaireId field.
func (o *QuestionnaireQuestionModel) SetQuestionnaireId(v int32) {
	o.QuestionnaireId = &v
}

// GetQuestionnaireQuestionId returns the QuestionnaireQuestionId field value if set, zero value otherwise.
func (o *QuestionnaireQuestionModel) GetQuestionnaireQuestionId() int32 {
	if o == nil || o.QuestionnaireQuestionId == nil {
		var ret int32
		return ret
	}
	return *o.QuestionnaireQuestionId
}

// GetQuestionnaireQuestionIdOk returns a tuple with the QuestionnaireQuestionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireQuestionModel) GetQuestionnaireQuestionIdOk() (*int32, bool) {
	if o == nil || o.QuestionnaireQuestionId == nil {
		return nil, false
	}
	return o.QuestionnaireQuestionId, true
}

// HasQuestionnaireQuestionId returns a boolean if a field has been set.
func (o *QuestionnaireQuestionModel) HasQuestionnaireQuestionId() bool {
	if o != nil && o.QuestionnaireQuestionId != nil {
		return true
	}

	return false
}

// SetQuestionnaireQuestionId gets a reference to the given int32 and assigns it to the QuestionnaireQuestionId field.
func (o *QuestionnaireQuestionModel) SetQuestionnaireQuestionId(v int32) {
	o.QuestionnaireQuestionId = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *QuestionnaireQuestionModel) GetOrder() int32 {
	if o == nil || o.Order == nil {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireQuestionModel) GetOrderOk() (*int32, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *QuestionnaireQuestionModel) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *QuestionnaireQuestionModel) SetOrder(v int32) {
	o.Order = &v
}

// GetQuestion returns the Question field value if set, zero value otherwise.
func (o *QuestionnaireQuestionModel) GetQuestion() string {
	if o == nil || o.Question == nil {
		var ret string
		return ret
	}
	return *o.Question
}

// GetQuestionOk returns a tuple with the Question field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireQuestionModel) GetQuestionOk() (*string, bool) {
	if o == nil || o.Question == nil {
		return nil, false
	}
	return o.Question, true
}

// HasQuestion returns a boolean if a field has been set.
func (o *QuestionnaireQuestionModel) HasQuestion() bool {
	if o != nil && o.Question != nil {
		return true
	}

	return false
}

// SetQuestion gets a reference to the given string and assigns it to the Question field.
func (o *QuestionnaireQuestionModel) SetQuestion(v string) {
	o.Question = &v
}

// GetAnswerOptions returns the AnswerOptions field value if set, zero value otherwise.
func (o *QuestionnaireQuestionModel) GetAnswerOptions() []QuestionnaireQuestionOptionModel {
	if o == nil || o.AnswerOptions == nil {
		var ret []QuestionnaireQuestionOptionModel
		return ret
	}
	return *o.AnswerOptions
}

// GetAnswerOptionsOk returns a tuple with the AnswerOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireQuestionModel) GetAnswerOptionsOk() (*[]QuestionnaireQuestionOptionModel, bool) {
	if o == nil || o.AnswerOptions == nil {
		return nil, false
	}
	return o.AnswerOptions, true
}

// HasAnswerOptions returns a boolean if a field has been set.
func (o *QuestionnaireQuestionModel) HasAnswerOptions() bool {
	if o != nil && o.AnswerOptions != nil {
		return true
	}

	return false
}

// SetAnswerOptions gets a reference to the given []QuestionnaireQuestionOptionModel and assigns it to the AnswerOptions field.
func (o *QuestionnaireQuestionModel) SetAnswerOptions(v []QuestionnaireQuestionOptionModel) {
	o.AnswerOptions = &v
}

// GetAnswer returns the Answer field value if set, zero value otherwise.
func (o *QuestionnaireQuestionModel) GetAnswer() string {
	if o == nil || o.Answer == nil {
		var ret string
		return ret
	}
	return *o.Answer
}

// GetAnswerOk returns a tuple with the Answer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireQuestionModel) GetAnswerOk() (*string, bool) {
	if o == nil || o.Answer == nil {
		return nil, false
	}
	return o.Answer, true
}

// HasAnswer returns a boolean if a field has been set.
func (o *QuestionnaireQuestionModel) HasAnswer() bool {
	if o != nil && o.Answer != nil {
		return true
	}

	return false
}

// SetAnswer gets a reference to the given string and assigns it to the Answer field.
func (o *QuestionnaireQuestionModel) SetAnswer(v string) {
	o.Answer = &v
}

// GetChosenOptions returns the ChosenOptions field value if set, zero value otherwise.
func (o *QuestionnaireQuestionModel) GetChosenOptions() []int32 {
	if o == nil || o.ChosenOptions == nil {
		var ret []int32
		return ret
	}
	return *o.ChosenOptions
}

// GetChosenOptionsOk returns a tuple with the ChosenOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireQuestionModel) GetChosenOptionsOk() (*[]int32, bool) {
	if o == nil || o.ChosenOptions == nil {
		return nil, false
	}
	return o.ChosenOptions, true
}

// HasChosenOptions returns a boolean if a field has been set.
func (o *QuestionnaireQuestionModel) HasChosenOptions() bool {
	if o != nil && o.ChosenOptions != nil {
		return true
	}

	return false
}

// SetChosenOptions gets a reference to the given []int32 and assigns it to the ChosenOptions field.
func (o *QuestionnaireQuestionModel) SetChosenOptions(v []int32) {
	o.ChosenOptions = &v
}

// GetQuestionnaireQuestionTypeName returns the QuestionnaireQuestionTypeName field value if set, zero value otherwise.
func (o *QuestionnaireQuestionModel) GetQuestionnaireQuestionTypeName() string {
	if o == nil || o.QuestionnaireQuestionTypeName == nil {
		var ret string
		return ret
	}
	return *o.QuestionnaireQuestionTypeName
}

// GetQuestionnaireQuestionTypeNameOk returns a tuple with the QuestionnaireQuestionTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireQuestionModel) GetQuestionnaireQuestionTypeNameOk() (*string, bool) {
	if o == nil || o.QuestionnaireQuestionTypeName == nil {
		return nil, false
	}
	return o.QuestionnaireQuestionTypeName, true
}

// HasQuestionnaireQuestionTypeName returns a boolean if a field has been set.
func (o *QuestionnaireQuestionModel) HasQuestionnaireQuestionTypeName() bool {
	if o != nil && o.QuestionnaireQuestionTypeName != nil {
		return true
	}

	return false
}

// SetQuestionnaireQuestionTypeName gets a reference to the given string and assigns it to the QuestionnaireQuestionTypeName field.
func (o *QuestionnaireQuestionModel) SetQuestionnaireQuestionTypeName(v string) {
	o.QuestionnaireQuestionTypeName = &v
}

func (o QuestionnaireQuestionModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.QuestionnaireId != nil {
		toSerialize["questionnaire_id"] = o.QuestionnaireId
	}
	if o.QuestionnaireQuestionId != nil {
		toSerialize["questionnaire_question_id"] = o.QuestionnaireQuestionId
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.Question != nil {
		toSerialize["question"] = o.Question
	}
	if o.AnswerOptions != nil {
		toSerialize["answer_options"] = o.AnswerOptions
	}
	if o.Answer != nil {
		toSerialize["answer"] = o.Answer
	}
	if o.ChosenOptions != nil {
		toSerialize["chosen_options"] = o.ChosenOptions
	}
	if o.QuestionnaireQuestionTypeName != nil {
		toSerialize["questionnaire_question_type_name"] = o.QuestionnaireQuestionTypeName
	}
	return json.Marshal(toSerialize)
}

type NullableQuestionnaireQuestionModel struct {
	value *QuestionnaireQuestionModel
	isSet bool
}

func (v NullableQuestionnaireQuestionModel) Get() *QuestionnaireQuestionModel {
	return v.value
}

func (v *NullableQuestionnaireQuestionModel) Set(val *QuestionnaireQuestionModel) {
	v.value = val
	v.isSet = true
}

func (v NullableQuestionnaireQuestionModel) IsSet() bool {
	return v.isSet
}

func (v *NullableQuestionnaireQuestionModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuestionnaireQuestionModel(val *QuestionnaireQuestionModel) *NullableQuestionnaireQuestionModel {
	return &NullableQuestionnaireQuestionModel{value: val, isSet: true}
}

func (v NullableQuestionnaireQuestionModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuestionnaireQuestionModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


