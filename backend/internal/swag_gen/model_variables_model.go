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

// VariablesModel struct for VariablesModel
type VariablesModel struct {
	// 
	Content *string `json:"content,omitempty"`
}

// NewVariablesModel instantiates a new VariablesModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariablesModel() *VariablesModel {
	this := VariablesModel{}
	return &this
}

// NewVariablesModelWithDefaults instantiates a new VariablesModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariablesModelWithDefaults() *VariablesModel {
	this := VariablesModel{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *VariablesModel) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariablesModel) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *VariablesModel) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *VariablesModel) SetContent(v string) {
	o.Content = &v
}

func (o VariablesModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	return json.Marshal(toSerialize)
}

type NullableVariablesModel struct {
	value *VariablesModel
	isSet bool
}

func (v NullableVariablesModel) Get() *VariablesModel {
	return v.value
}

func (v *NullableVariablesModel) Set(val *VariablesModel) {
	v.value = val
	v.isSet = true
}

func (v NullableVariablesModel) IsSet() bool {
	return v.isSet
}

func (v *NullableVariablesModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariablesModel(val *VariablesModel) *NullableVariablesModel {
	return &NullableVariablesModel{value: val, isSet: true}
}

func (v NullableVariablesModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariablesModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

