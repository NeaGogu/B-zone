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

// QuestionnaireModel struct for QuestionnaireModel
type QuestionnaireModel struct {
	// Unique Identifier
	Id *int64 `json:"id,omitempty"`
	// if active=0: Questionnaire has been removed and is no longer visible in any bumbal interface
	Active *bool `json:"active,omitempty"`
	// ID of the parent questionnaire template, just for reference
	QuestionnaireTemplateId *int64 `json:"questionnaire_template_id,omitempty"`
	// Questionnaire name
	Name *string `json:"name,omitempty"`
	// ID of the questionnaire type, just for reference
	QuestionnaireTypeId *int64 `json:"questionnaire_type_id,omitempty"`
	// name of the questionnaire type
	QuestionnaireTypeName *string `json:"questionnaire_type_name,omitempty"`
	// ISO lang code
	LangCode *string `json:"lang_code,omitempty"`
	// 
	Answers *[]QuestionnaireAnswerModel `json:"answers,omitempty"`
	// Object type IDs available for this questionnaire
	ObjectType *map[string]interface{} `json:"object_type,omitempty"`
	// Object type name for the bound object to this questionnaire
	ObjectTypeName *string `json:"object_type_name,omitempty"`
	// Object ID
	ObjectId *int32 `json:"object_id,omitempty"`
}

// NewQuestionnaireModel instantiates a new QuestionnaireModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuestionnaireModel() *QuestionnaireModel {
	this := QuestionnaireModel{}
	return &this
}

// NewQuestionnaireModelWithDefaults instantiates a new QuestionnaireModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuestionnaireModelWithDefaults() *QuestionnaireModel {
	this := QuestionnaireModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *QuestionnaireModel) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireModel) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *QuestionnaireModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *QuestionnaireModel) SetId(v int64) {
	o.Id = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *QuestionnaireModel) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireModel) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *QuestionnaireModel) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *QuestionnaireModel) SetActive(v bool) {
	o.Active = &v
}

// GetQuestionnaireTemplateId returns the QuestionnaireTemplateId field value if set, zero value otherwise.
func (o *QuestionnaireModel) GetQuestionnaireTemplateId() int64 {
	if o == nil || o.QuestionnaireTemplateId == nil {
		var ret int64
		return ret
	}
	return *o.QuestionnaireTemplateId
}

// GetQuestionnaireTemplateIdOk returns a tuple with the QuestionnaireTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireModel) GetQuestionnaireTemplateIdOk() (*int64, bool) {
	if o == nil || o.QuestionnaireTemplateId == nil {
		return nil, false
	}
	return o.QuestionnaireTemplateId, true
}

// HasQuestionnaireTemplateId returns a boolean if a field has been set.
func (o *QuestionnaireModel) HasQuestionnaireTemplateId() bool {
	if o != nil && o.QuestionnaireTemplateId != nil {
		return true
	}

	return false
}

// SetQuestionnaireTemplateId gets a reference to the given int64 and assigns it to the QuestionnaireTemplateId field.
func (o *QuestionnaireModel) SetQuestionnaireTemplateId(v int64) {
	o.QuestionnaireTemplateId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *QuestionnaireModel) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireModel) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *QuestionnaireModel) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *QuestionnaireModel) SetName(v string) {
	o.Name = &v
}

// GetQuestionnaireTypeId returns the QuestionnaireTypeId field value if set, zero value otherwise.
func (o *QuestionnaireModel) GetQuestionnaireTypeId() int64 {
	if o == nil || o.QuestionnaireTypeId == nil {
		var ret int64
		return ret
	}
	return *o.QuestionnaireTypeId
}

// GetQuestionnaireTypeIdOk returns a tuple with the QuestionnaireTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireModel) GetQuestionnaireTypeIdOk() (*int64, bool) {
	if o == nil || o.QuestionnaireTypeId == nil {
		return nil, false
	}
	return o.QuestionnaireTypeId, true
}

// HasQuestionnaireTypeId returns a boolean if a field has been set.
func (o *QuestionnaireModel) HasQuestionnaireTypeId() bool {
	if o != nil && o.QuestionnaireTypeId != nil {
		return true
	}

	return false
}

// SetQuestionnaireTypeId gets a reference to the given int64 and assigns it to the QuestionnaireTypeId field.
func (o *QuestionnaireModel) SetQuestionnaireTypeId(v int64) {
	o.QuestionnaireTypeId = &v
}

// GetQuestionnaireTypeName returns the QuestionnaireTypeName field value if set, zero value otherwise.
func (o *QuestionnaireModel) GetQuestionnaireTypeName() string {
	if o == nil || o.QuestionnaireTypeName == nil {
		var ret string
		return ret
	}
	return *o.QuestionnaireTypeName
}

// GetQuestionnaireTypeNameOk returns a tuple with the QuestionnaireTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireModel) GetQuestionnaireTypeNameOk() (*string, bool) {
	if o == nil || o.QuestionnaireTypeName == nil {
		return nil, false
	}
	return o.QuestionnaireTypeName, true
}

// HasQuestionnaireTypeName returns a boolean if a field has been set.
func (o *QuestionnaireModel) HasQuestionnaireTypeName() bool {
	if o != nil && o.QuestionnaireTypeName != nil {
		return true
	}

	return false
}

// SetQuestionnaireTypeName gets a reference to the given string and assigns it to the QuestionnaireTypeName field.
func (o *QuestionnaireModel) SetQuestionnaireTypeName(v string) {
	o.QuestionnaireTypeName = &v
}

// GetLangCode returns the LangCode field value if set, zero value otherwise.
func (o *QuestionnaireModel) GetLangCode() string {
	if o == nil || o.LangCode == nil {
		var ret string
		return ret
	}
	return *o.LangCode
}

// GetLangCodeOk returns a tuple with the LangCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireModel) GetLangCodeOk() (*string, bool) {
	if o == nil || o.LangCode == nil {
		return nil, false
	}
	return o.LangCode, true
}

// HasLangCode returns a boolean if a field has been set.
func (o *QuestionnaireModel) HasLangCode() bool {
	if o != nil && o.LangCode != nil {
		return true
	}

	return false
}

// SetLangCode gets a reference to the given string and assigns it to the LangCode field.
func (o *QuestionnaireModel) SetLangCode(v string) {
	o.LangCode = &v
}

// GetAnswers returns the Answers field value if set, zero value otherwise.
func (o *QuestionnaireModel) GetAnswers() []QuestionnaireAnswerModel {
	if o == nil || o.Answers == nil {
		var ret []QuestionnaireAnswerModel
		return ret
	}
	return *o.Answers
}

// GetAnswersOk returns a tuple with the Answers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireModel) GetAnswersOk() (*[]QuestionnaireAnswerModel, bool) {
	if o == nil || o.Answers == nil {
		return nil, false
	}
	return o.Answers, true
}

// HasAnswers returns a boolean if a field has been set.
func (o *QuestionnaireModel) HasAnswers() bool {
	if o != nil && o.Answers != nil {
		return true
	}

	return false
}

// SetAnswers gets a reference to the given []QuestionnaireAnswerModel and assigns it to the Answers field.
func (o *QuestionnaireModel) SetAnswers(v []QuestionnaireAnswerModel) {
	o.Answers = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *QuestionnaireModel) GetObjectType() map[string]interface{} {
	if o == nil || o.ObjectType == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireModel) GetObjectTypeOk() (*map[string]interface{}, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *QuestionnaireModel) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given map[string]interface{} and assigns it to the ObjectType field.
func (o *QuestionnaireModel) SetObjectType(v map[string]interface{}) {
	o.ObjectType = &v
}

// GetObjectTypeName returns the ObjectTypeName field value if set, zero value otherwise.
func (o *QuestionnaireModel) GetObjectTypeName() string {
	if o == nil || o.ObjectTypeName == nil {
		var ret string
		return ret
	}
	return *o.ObjectTypeName
}

// GetObjectTypeNameOk returns a tuple with the ObjectTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireModel) GetObjectTypeNameOk() (*string, bool) {
	if o == nil || o.ObjectTypeName == nil {
		return nil, false
	}
	return o.ObjectTypeName, true
}

// HasObjectTypeName returns a boolean if a field has been set.
func (o *QuestionnaireModel) HasObjectTypeName() bool {
	if o != nil && o.ObjectTypeName != nil {
		return true
	}

	return false
}

// SetObjectTypeName gets a reference to the given string and assigns it to the ObjectTypeName field.
func (o *QuestionnaireModel) SetObjectTypeName(v string) {
	o.ObjectTypeName = &v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *QuestionnaireModel) GetObjectId() int32 {
	if o == nil || o.ObjectId == nil {
		var ret int32
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireModel) GetObjectIdOk() (*int32, bool) {
	if o == nil || o.ObjectId == nil {
		return nil, false
	}
	return o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *QuestionnaireModel) HasObjectId() bool {
	if o != nil && o.ObjectId != nil {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given int32 and assigns it to the ObjectId field.
func (o *QuestionnaireModel) SetObjectId(v int32) {
	o.ObjectId = &v
}

func (o QuestionnaireModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.QuestionnaireTemplateId != nil {
		toSerialize["questionnaire_template_id"] = o.QuestionnaireTemplateId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.QuestionnaireTypeId != nil {
		toSerialize["questionnaire_type_id"] = o.QuestionnaireTypeId
	}
	if o.QuestionnaireTypeName != nil {
		toSerialize["questionnaire_type_name"] = o.QuestionnaireTypeName
	}
	if o.LangCode != nil {
		toSerialize["lang_code"] = o.LangCode
	}
	if o.Answers != nil {
		toSerialize["answers"] = o.Answers
	}
	if o.ObjectType != nil {
		toSerialize["object_type"] = o.ObjectType
	}
	if o.ObjectTypeName != nil {
		toSerialize["object_type_name"] = o.ObjectTypeName
	}
	if o.ObjectId != nil {
		toSerialize["object_id"] = o.ObjectId
	}
	return json.Marshal(toSerialize)
}

type NullableQuestionnaireModel struct {
	value *QuestionnaireModel
	isSet bool
}

func (v NullableQuestionnaireModel) Get() *QuestionnaireModel {
	return v.value
}

func (v *NullableQuestionnaireModel) Set(val *QuestionnaireModel) {
	v.value = val
	v.isSet = true
}

func (v NullableQuestionnaireModel) IsSet() bool {
	return v.isSet
}

func (v *NullableQuestionnaireModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuestionnaireModel(val *QuestionnaireModel) *NullableQuestionnaireModel {
	return &NullableQuestionnaireModel{value: val, isSet: true}
}

func (v NullableQuestionnaireModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuestionnaireModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


