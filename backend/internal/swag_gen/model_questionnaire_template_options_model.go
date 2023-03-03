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

// QuestionnaireTemplateOptionsModel struct for QuestionnaireTemplateOptionsModel
type QuestionnaireTemplateOptionsModel struct {
	// 
	IncludeRecordInfo *bool `json:"include_record_info,omitempty"`
	// 
	IncludeQuestions *bool `json:"include_questions,omitempty"`
	// 
	IncludeOptions *bool `json:"include_options,omitempty"`
	// 
	IncludeTexts *bool `json:"include_texts,omitempty"`
	// 
	IncludeQuestionTypeName *bool `json:"include_question_type_name,omitempty"`
	// 
	IncludeZones *bool `json:"include_zones,omitempty"`
	// 
	IncludeBrands *bool `json:"include_brands,omitempty"`
	// 
	IncludeTagIds *bool `json:"include_tag_ids,omitempty"`
	// 
	IncludeTagNames *bool `json:"include_tag_names,omitempty"`
	// 
	IncludeTags *bool `json:"include_tags,omitempty"`
	// 
	IncludeZoneIds *bool `json:"include_zone_ids,omitempty"`
	// 
	IncludeZoneNames *bool `json:"include_zone_names,omitempty"`
	// 
	IncludeBrandIds *bool `json:"include_brand_ids,omitempty"`
	// 
	IncludeBrandNames *bool `json:"include_brand_names,omitempty"`
}

// NewQuestionnaireTemplateOptionsModel instantiates a new QuestionnaireTemplateOptionsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuestionnaireTemplateOptionsModel() *QuestionnaireTemplateOptionsModel {
	this := QuestionnaireTemplateOptionsModel{}
	return &this
}

// NewQuestionnaireTemplateOptionsModelWithDefaults instantiates a new QuestionnaireTemplateOptionsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuestionnaireTemplateOptionsModelWithDefaults() *QuestionnaireTemplateOptionsModel {
	this := QuestionnaireTemplateOptionsModel{}
	return &this
}

// GetIncludeRecordInfo returns the IncludeRecordInfo field value if set, zero value otherwise.
func (o *QuestionnaireTemplateOptionsModel) GetIncludeRecordInfo() bool {
	if o == nil || o.IncludeRecordInfo == nil {
		var ret bool
		return ret
	}
	return *o.IncludeRecordInfo
}

// GetIncludeRecordInfoOk returns a tuple with the IncludeRecordInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireTemplateOptionsModel) GetIncludeRecordInfoOk() (*bool, bool) {
	if o == nil || o.IncludeRecordInfo == nil {
		return nil, false
	}
	return o.IncludeRecordInfo, true
}

// HasIncludeRecordInfo returns a boolean if a field has been set.
func (o *QuestionnaireTemplateOptionsModel) HasIncludeRecordInfo() bool {
	if o != nil && o.IncludeRecordInfo != nil {
		return true
	}

	return false
}

// SetIncludeRecordInfo gets a reference to the given bool and assigns it to the IncludeRecordInfo field.
func (o *QuestionnaireTemplateOptionsModel) SetIncludeRecordInfo(v bool) {
	o.IncludeRecordInfo = &v
}

// GetIncludeQuestions returns the IncludeQuestions field value if set, zero value otherwise.
func (o *QuestionnaireTemplateOptionsModel) GetIncludeQuestions() bool {
	if o == nil || o.IncludeQuestions == nil {
		var ret bool
		return ret
	}
	return *o.IncludeQuestions
}

// GetIncludeQuestionsOk returns a tuple with the IncludeQuestions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireTemplateOptionsModel) GetIncludeQuestionsOk() (*bool, bool) {
	if o == nil || o.IncludeQuestions == nil {
		return nil, false
	}
	return o.IncludeQuestions, true
}

// HasIncludeQuestions returns a boolean if a field has been set.
func (o *QuestionnaireTemplateOptionsModel) HasIncludeQuestions() bool {
	if o != nil && o.IncludeQuestions != nil {
		return true
	}

	return false
}

// SetIncludeQuestions gets a reference to the given bool and assigns it to the IncludeQuestions field.
func (o *QuestionnaireTemplateOptionsModel) SetIncludeQuestions(v bool) {
	o.IncludeQuestions = &v
}

// GetIncludeOptions returns the IncludeOptions field value if set, zero value otherwise.
func (o *QuestionnaireTemplateOptionsModel) GetIncludeOptions() bool {
	if o == nil || o.IncludeOptions == nil {
		var ret bool
		return ret
	}
	return *o.IncludeOptions
}

// GetIncludeOptionsOk returns a tuple with the IncludeOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireTemplateOptionsModel) GetIncludeOptionsOk() (*bool, bool) {
	if o == nil || o.IncludeOptions == nil {
		return nil, false
	}
	return o.IncludeOptions, true
}

// HasIncludeOptions returns a boolean if a field has been set.
func (o *QuestionnaireTemplateOptionsModel) HasIncludeOptions() bool {
	if o != nil && o.IncludeOptions != nil {
		return true
	}

	return false
}

// SetIncludeOptions gets a reference to the given bool and assigns it to the IncludeOptions field.
func (o *QuestionnaireTemplateOptionsModel) SetIncludeOptions(v bool) {
	o.IncludeOptions = &v
}

// GetIncludeTexts returns the IncludeTexts field value if set, zero value otherwise.
func (o *QuestionnaireTemplateOptionsModel) GetIncludeTexts() bool {
	if o == nil || o.IncludeTexts == nil {
		var ret bool
		return ret
	}
	return *o.IncludeTexts
}

// GetIncludeTextsOk returns a tuple with the IncludeTexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireTemplateOptionsModel) GetIncludeTextsOk() (*bool, bool) {
	if o == nil || o.IncludeTexts == nil {
		return nil, false
	}
	return o.IncludeTexts, true
}

// HasIncludeTexts returns a boolean if a field has been set.
func (o *QuestionnaireTemplateOptionsModel) HasIncludeTexts() bool {
	if o != nil && o.IncludeTexts != nil {
		return true
	}

	return false
}

// SetIncludeTexts gets a reference to the given bool and assigns it to the IncludeTexts field.
func (o *QuestionnaireTemplateOptionsModel) SetIncludeTexts(v bool) {
	o.IncludeTexts = &v
}

// GetIncludeQuestionTypeName returns the IncludeQuestionTypeName field value if set, zero value otherwise.
func (o *QuestionnaireTemplateOptionsModel) GetIncludeQuestionTypeName() bool {
	if o == nil || o.IncludeQuestionTypeName == nil {
		var ret bool
		return ret
	}
	return *o.IncludeQuestionTypeName
}

// GetIncludeQuestionTypeNameOk returns a tuple with the IncludeQuestionTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireTemplateOptionsModel) GetIncludeQuestionTypeNameOk() (*bool, bool) {
	if o == nil || o.IncludeQuestionTypeName == nil {
		return nil, false
	}
	return o.IncludeQuestionTypeName, true
}

// HasIncludeQuestionTypeName returns a boolean if a field has been set.
func (o *QuestionnaireTemplateOptionsModel) HasIncludeQuestionTypeName() bool {
	if o != nil && o.IncludeQuestionTypeName != nil {
		return true
	}

	return false
}

// SetIncludeQuestionTypeName gets a reference to the given bool and assigns it to the IncludeQuestionTypeName field.
func (o *QuestionnaireTemplateOptionsModel) SetIncludeQuestionTypeName(v bool) {
	o.IncludeQuestionTypeName = &v
}

// GetIncludeZones returns the IncludeZones field value if set, zero value otherwise.
func (o *QuestionnaireTemplateOptionsModel) GetIncludeZones() bool {
	if o == nil || o.IncludeZones == nil {
		var ret bool
		return ret
	}
	return *o.IncludeZones
}

// GetIncludeZonesOk returns a tuple with the IncludeZones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireTemplateOptionsModel) GetIncludeZonesOk() (*bool, bool) {
	if o == nil || o.IncludeZones == nil {
		return nil, false
	}
	return o.IncludeZones, true
}

// HasIncludeZones returns a boolean if a field has been set.
func (o *QuestionnaireTemplateOptionsModel) HasIncludeZones() bool {
	if o != nil && o.IncludeZones != nil {
		return true
	}

	return false
}

// SetIncludeZones gets a reference to the given bool and assigns it to the IncludeZones field.
func (o *QuestionnaireTemplateOptionsModel) SetIncludeZones(v bool) {
	o.IncludeZones = &v
}

// GetIncludeBrands returns the IncludeBrands field value if set, zero value otherwise.
func (o *QuestionnaireTemplateOptionsModel) GetIncludeBrands() bool {
	if o == nil || o.IncludeBrands == nil {
		var ret bool
		return ret
	}
	return *o.IncludeBrands
}

// GetIncludeBrandsOk returns a tuple with the IncludeBrands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireTemplateOptionsModel) GetIncludeBrandsOk() (*bool, bool) {
	if o == nil || o.IncludeBrands == nil {
		return nil, false
	}
	return o.IncludeBrands, true
}

// HasIncludeBrands returns a boolean if a field has been set.
func (o *QuestionnaireTemplateOptionsModel) HasIncludeBrands() bool {
	if o != nil && o.IncludeBrands != nil {
		return true
	}

	return false
}

// SetIncludeBrands gets a reference to the given bool and assigns it to the IncludeBrands field.
func (o *QuestionnaireTemplateOptionsModel) SetIncludeBrands(v bool) {
	o.IncludeBrands = &v
}

// GetIncludeTagIds returns the IncludeTagIds field value if set, zero value otherwise.
func (o *QuestionnaireTemplateOptionsModel) GetIncludeTagIds() bool {
	if o == nil || o.IncludeTagIds == nil {
		var ret bool
		return ret
	}
	return *o.IncludeTagIds
}

// GetIncludeTagIdsOk returns a tuple with the IncludeTagIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireTemplateOptionsModel) GetIncludeTagIdsOk() (*bool, bool) {
	if o == nil || o.IncludeTagIds == nil {
		return nil, false
	}
	return o.IncludeTagIds, true
}

// HasIncludeTagIds returns a boolean if a field has been set.
func (o *QuestionnaireTemplateOptionsModel) HasIncludeTagIds() bool {
	if o != nil && o.IncludeTagIds != nil {
		return true
	}

	return false
}

// SetIncludeTagIds gets a reference to the given bool and assigns it to the IncludeTagIds field.
func (o *QuestionnaireTemplateOptionsModel) SetIncludeTagIds(v bool) {
	o.IncludeTagIds = &v
}

// GetIncludeTagNames returns the IncludeTagNames field value if set, zero value otherwise.
func (o *QuestionnaireTemplateOptionsModel) GetIncludeTagNames() bool {
	if o == nil || o.IncludeTagNames == nil {
		var ret bool
		return ret
	}
	return *o.IncludeTagNames
}

// GetIncludeTagNamesOk returns a tuple with the IncludeTagNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireTemplateOptionsModel) GetIncludeTagNamesOk() (*bool, bool) {
	if o == nil || o.IncludeTagNames == nil {
		return nil, false
	}
	return o.IncludeTagNames, true
}

// HasIncludeTagNames returns a boolean if a field has been set.
func (o *QuestionnaireTemplateOptionsModel) HasIncludeTagNames() bool {
	if o != nil && o.IncludeTagNames != nil {
		return true
	}

	return false
}

// SetIncludeTagNames gets a reference to the given bool and assigns it to the IncludeTagNames field.
func (o *QuestionnaireTemplateOptionsModel) SetIncludeTagNames(v bool) {
	o.IncludeTagNames = &v
}

// GetIncludeTags returns the IncludeTags field value if set, zero value otherwise.
func (o *QuestionnaireTemplateOptionsModel) GetIncludeTags() bool {
	if o == nil || o.IncludeTags == nil {
		var ret bool
		return ret
	}
	return *o.IncludeTags
}

// GetIncludeTagsOk returns a tuple with the IncludeTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireTemplateOptionsModel) GetIncludeTagsOk() (*bool, bool) {
	if o == nil || o.IncludeTags == nil {
		return nil, false
	}
	return o.IncludeTags, true
}

// HasIncludeTags returns a boolean if a field has been set.
func (o *QuestionnaireTemplateOptionsModel) HasIncludeTags() bool {
	if o != nil && o.IncludeTags != nil {
		return true
	}

	return false
}

// SetIncludeTags gets a reference to the given bool and assigns it to the IncludeTags field.
func (o *QuestionnaireTemplateOptionsModel) SetIncludeTags(v bool) {
	o.IncludeTags = &v
}

// GetIncludeZoneIds returns the IncludeZoneIds field value if set, zero value otherwise.
func (o *QuestionnaireTemplateOptionsModel) GetIncludeZoneIds() bool {
	if o == nil || o.IncludeZoneIds == nil {
		var ret bool
		return ret
	}
	return *o.IncludeZoneIds
}

// GetIncludeZoneIdsOk returns a tuple with the IncludeZoneIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireTemplateOptionsModel) GetIncludeZoneIdsOk() (*bool, bool) {
	if o == nil || o.IncludeZoneIds == nil {
		return nil, false
	}
	return o.IncludeZoneIds, true
}

// HasIncludeZoneIds returns a boolean if a field has been set.
func (o *QuestionnaireTemplateOptionsModel) HasIncludeZoneIds() bool {
	if o != nil && o.IncludeZoneIds != nil {
		return true
	}

	return false
}

// SetIncludeZoneIds gets a reference to the given bool and assigns it to the IncludeZoneIds field.
func (o *QuestionnaireTemplateOptionsModel) SetIncludeZoneIds(v bool) {
	o.IncludeZoneIds = &v
}

// GetIncludeZoneNames returns the IncludeZoneNames field value if set, zero value otherwise.
func (o *QuestionnaireTemplateOptionsModel) GetIncludeZoneNames() bool {
	if o == nil || o.IncludeZoneNames == nil {
		var ret bool
		return ret
	}
	return *o.IncludeZoneNames
}

// GetIncludeZoneNamesOk returns a tuple with the IncludeZoneNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireTemplateOptionsModel) GetIncludeZoneNamesOk() (*bool, bool) {
	if o == nil || o.IncludeZoneNames == nil {
		return nil, false
	}
	return o.IncludeZoneNames, true
}

// HasIncludeZoneNames returns a boolean if a field has been set.
func (o *QuestionnaireTemplateOptionsModel) HasIncludeZoneNames() bool {
	if o != nil && o.IncludeZoneNames != nil {
		return true
	}

	return false
}

// SetIncludeZoneNames gets a reference to the given bool and assigns it to the IncludeZoneNames field.
func (o *QuestionnaireTemplateOptionsModel) SetIncludeZoneNames(v bool) {
	o.IncludeZoneNames = &v
}

// GetIncludeBrandIds returns the IncludeBrandIds field value if set, zero value otherwise.
func (o *QuestionnaireTemplateOptionsModel) GetIncludeBrandIds() bool {
	if o == nil || o.IncludeBrandIds == nil {
		var ret bool
		return ret
	}
	return *o.IncludeBrandIds
}

// GetIncludeBrandIdsOk returns a tuple with the IncludeBrandIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireTemplateOptionsModel) GetIncludeBrandIdsOk() (*bool, bool) {
	if o == nil || o.IncludeBrandIds == nil {
		return nil, false
	}
	return o.IncludeBrandIds, true
}

// HasIncludeBrandIds returns a boolean if a field has been set.
func (o *QuestionnaireTemplateOptionsModel) HasIncludeBrandIds() bool {
	if o != nil && o.IncludeBrandIds != nil {
		return true
	}

	return false
}

// SetIncludeBrandIds gets a reference to the given bool and assigns it to the IncludeBrandIds field.
func (o *QuestionnaireTemplateOptionsModel) SetIncludeBrandIds(v bool) {
	o.IncludeBrandIds = &v
}

// GetIncludeBrandNames returns the IncludeBrandNames field value if set, zero value otherwise.
func (o *QuestionnaireTemplateOptionsModel) GetIncludeBrandNames() bool {
	if o == nil || o.IncludeBrandNames == nil {
		var ret bool
		return ret
	}
	return *o.IncludeBrandNames
}

// GetIncludeBrandNamesOk returns a tuple with the IncludeBrandNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuestionnaireTemplateOptionsModel) GetIncludeBrandNamesOk() (*bool, bool) {
	if o == nil || o.IncludeBrandNames == nil {
		return nil, false
	}
	return o.IncludeBrandNames, true
}

// HasIncludeBrandNames returns a boolean if a field has been set.
func (o *QuestionnaireTemplateOptionsModel) HasIncludeBrandNames() bool {
	if o != nil && o.IncludeBrandNames != nil {
		return true
	}

	return false
}

// SetIncludeBrandNames gets a reference to the given bool and assigns it to the IncludeBrandNames field.
func (o *QuestionnaireTemplateOptionsModel) SetIncludeBrandNames(v bool) {
	o.IncludeBrandNames = &v
}

func (o QuestionnaireTemplateOptionsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IncludeRecordInfo != nil {
		toSerialize["include_record_info"] = o.IncludeRecordInfo
	}
	if o.IncludeQuestions != nil {
		toSerialize["include_questions"] = o.IncludeQuestions
	}
	if o.IncludeOptions != nil {
		toSerialize["include_options"] = o.IncludeOptions
	}
	if o.IncludeTexts != nil {
		toSerialize["include_texts"] = o.IncludeTexts
	}
	if o.IncludeQuestionTypeName != nil {
		toSerialize["include_question_type_name"] = o.IncludeQuestionTypeName
	}
	if o.IncludeZones != nil {
		toSerialize["include_zones"] = o.IncludeZones
	}
	if o.IncludeBrands != nil {
		toSerialize["include_brands"] = o.IncludeBrands
	}
	if o.IncludeTagIds != nil {
		toSerialize["include_tag_ids"] = o.IncludeTagIds
	}
	if o.IncludeTagNames != nil {
		toSerialize["include_tag_names"] = o.IncludeTagNames
	}
	if o.IncludeTags != nil {
		toSerialize["include_tags"] = o.IncludeTags
	}
	if o.IncludeZoneIds != nil {
		toSerialize["include_zone_ids"] = o.IncludeZoneIds
	}
	if o.IncludeZoneNames != nil {
		toSerialize["include_zone_names"] = o.IncludeZoneNames
	}
	if o.IncludeBrandIds != nil {
		toSerialize["include_brand_ids"] = o.IncludeBrandIds
	}
	if o.IncludeBrandNames != nil {
		toSerialize["include_brand_names"] = o.IncludeBrandNames
	}
	return json.Marshal(toSerialize)
}

type NullableQuestionnaireTemplateOptionsModel struct {
	value *QuestionnaireTemplateOptionsModel
	isSet bool
}

func (v NullableQuestionnaireTemplateOptionsModel) Get() *QuestionnaireTemplateOptionsModel {
	return v.value
}

func (v *NullableQuestionnaireTemplateOptionsModel) Set(val *QuestionnaireTemplateOptionsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableQuestionnaireTemplateOptionsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableQuestionnaireTemplateOptionsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuestionnaireTemplateOptionsModel(val *QuestionnaireTemplateOptionsModel) *NullableQuestionnaireTemplateOptionsModel {
	return &NullableQuestionnaireTemplateOptionsModel{value: val, isSet: true}
}

func (v NullableQuestionnaireTemplateOptionsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuestionnaireTemplateOptionsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


