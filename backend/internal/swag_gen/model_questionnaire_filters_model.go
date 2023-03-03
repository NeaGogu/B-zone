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

// QuestionnaireFiltersModel struct for QuestionnaireFiltersModel
type QuestionnaireFiltersModel struct {
	// Object type IDs available for this questionnaire
	ObjectType *int32 `json:"object_type,omitempty"`
	// Object id's available for this questionnaire
	ObjectId *int32 `json:"object_id,omitempty"`
	// Object type name available for this questionnaire
	ObjectTypeName *string `json:"object_type_name,omitempty"`
	// ISO lang code
	LangCode *string `json:"lang_code,omitempty"`
	// Questionnaire type name
	QuestionnaireTypeName *map[string]interface{} `json:"questionnaire_type_name,omitempty"`
}

// NewQuestionnaireFiltersModel instantiates a new QuestionnaireFiltersModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuestionnaireFiltersModel() *QuestionnaireFiltersModel {
	this := QuestionnaireFiltersModel{}
	return &this
}

// NewQuestionnaireFiltersModelWithDefaults instantiates a new QuestionnaireFiltersModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuestionnaireFiltersModelWithDefaults() *QuestionnaireFiltersModel {
	this := QuestionnaireFiltersModel{}
	return &this
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *QuestionnaireFiltersModel) GetObjectType() int32 {
	if o == nil || o.ObjectType == nil {
		var ret int32
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireFiltersModel) GetObjectTypeOk() (*int32, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *QuestionnaireFiltersModel) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given int32 and assigns it to the ObjectType field.
func (o *QuestionnaireFiltersModel) SetObjectType(v int32) {
	o.ObjectType = &v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *QuestionnaireFiltersModel) GetObjectId() int32 {
	if o == nil || o.ObjectId == nil {
		var ret int32
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireFiltersModel) GetObjectIdOk() (*int32, bool) {
	if o == nil || o.ObjectId == nil {
		return nil, false
	}
	return o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *QuestionnaireFiltersModel) HasObjectId() bool {
	if o != nil && o.ObjectId != nil {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given int32 and assigns it to the ObjectId field.
func (o *QuestionnaireFiltersModel) SetObjectId(v int32) {
	o.ObjectId = &v
}

// GetObjectTypeName returns the ObjectTypeName field value if set, zero value otherwise.
func (o *QuestionnaireFiltersModel) GetObjectTypeName() string {
	if o == nil || o.ObjectTypeName == nil {
		var ret string
		return ret
	}
	return *o.ObjectTypeName
}

// GetObjectTypeNameOk returns a tuple with the ObjectTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireFiltersModel) GetObjectTypeNameOk() (*string, bool) {
	if o == nil || o.ObjectTypeName == nil {
		return nil, false
	}
	return o.ObjectTypeName, true
}

// HasObjectTypeName returns a boolean if a field has been set.
func (o *QuestionnaireFiltersModel) HasObjectTypeName() bool {
	if o != nil && o.ObjectTypeName != nil {
		return true
	}

	return false
}

// SetObjectTypeName gets a reference to the given string and assigns it to the ObjectTypeName field.
func (o *QuestionnaireFiltersModel) SetObjectTypeName(v string) {
	o.ObjectTypeName = &v
}

// GetLangCode returns the LangCode field value if set, zero value otherwise.
func (o *QuestionnaireFiltersModel) GetLangCode() string {
	if o == nil || o.LangCode == nil {
		var ret string
		return ret
	}
	return *o.LangCode
}

// GetLangCodeOk returns a tuple with the LangCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireFiltersModel) GetLangCodeOk() (*string, bool) {
	if o == nil || o.LangCode == nil {
		return nil, false
	}
	return o.LangCode, true
}

// HasLangCode returns a boolean if a field has been set.
func (o *QuestionnaireFiltersModel) HasLangCode() bool {
	if o != nil && o.LangCode != nil {
		return true
	}

	return false
}

// SetLangCode gets a reference to the given string and assigns it to the LangCode field.
func (o *QuestionnaireFiltersModel) SetLangCode(v string) {
	o.LangCode = &v
}

// GetQuestionnaireTypeName returns the QuestionnaireTypeName field value if set, zero value otherwise.
func (o *QuestionnaireFiltersModel) GetQuestionnaireTypeName() map[string]interface{} {
	if o == nil || o.QuestionnaireTypeName == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.QuestionnaireTypeName
}

// GetQuestionnaireTypeNameOk returns a tuple with the QuestionnaireTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireFiltersModel) GetQuestionnaireTypeNameOk() (*map[string]interface{}, bool) {
	if o == nil || o.QuestionnaireTypeName == nil {
		return nil, false
	}
	return o.QuestionnaireTypeName, true
}

// HasQuestionnaireTypeName returns a boolean if a field has been set.
func (o *QuestionnaireFiltersModel) HasQuestionnaireTypeName() bool {
	if o != nil && o.QuestionnaireTypeName != nil {
		return true
	}

	return false
}

// SetQuestionnaireTypeName gets a reference to the given map[string]interface{} and assigns it to the QuestionnaireTypeName field.
func (o *QuestionnaireFiltersModel) SetQuestionnaireTypeName(v map[string]interface{}) {
	o.QuestionnaireTypeName = &v
}

func (o QuestionnaireFiltersModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjectType != nil {
		toSerialize["object_type"] = o.ObjectType
	}
	if o.ObjectId != nil {
		toSerialize["object_id"] = o.ObjectId
	}
	if o.ObjectTypeName != nil {
		toSerialize["object_type_name"] = o.ObjectTypeName
	}
	if o.LangCode != nil {
		toSerialize["lang_code"] = o.LangCode
	}
	if o.QuestionnaireTypeName != nil {
		toSerialize["questionnaire_type_name"] = o.QuestionnaireTypeName
	}
	return json.Marshal(toSerialize)
}

type NullableQuestionnaireFiltersModel struct {
	value *QuestionnaireFiltersModel
	isSet bool
}

func (v NullableQuestionnaireFiltersModel) Get() *QuestionnaireFiltersModel {
	return v.value
}

func (v *NullableQuestionnaireFiltersModel) Set(val *QuestionnaireFiltersModel) {
	v.value = val
	v.isSet = true
}

func (v NullableQuestionnaireFiltersModel) IsSet() bool {
	return v.isSet
}

func (v *NullableQuestionnaireFiltersModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuestionnaireFiltersModel(val *QuestionnaireFiltersModel) *NullableQuestionnaireFiltersModel {
	return &NullableQuestionnaireFiltersModel{value: val, isSet: true}
}

func (v NullableQuestionnaireFiltersModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuestionnaireFiltersModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


