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

// ConfigModel struct for ConfigModel
type ConfigModel struct {
	// 
	Options *string `json:"options,omitempty"`
	// 
	KeyRing *string `json:"keyRing,omitempty"`
}

// NewConfigModel instantiates a new ConfigModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigModel() *ConfigModel {
	this := ConfigModel{}
	return &this
}

// NewConfigModelWithDefaults instantiates a new ConfigModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigModelWithDefaults() *ConfigModel {
	this := ConfigModel{}
	return &this
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ConfigModel) GetOptions() string {
	if o == nil || o.Options == nil {
		var ret string
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigModel) GetOptionsOk() (*string, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ConfigModel) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given string and assigns it to the Options field.
func (o *ConfigModel) SetOptions(v string) {
	o.Options = &v
}

// GetKeyRing returns the KeyRing field value if set, zero value otherwise.
func (o *ConfigModel) GetKeyRing() string {
	if o == nil || o.KeyRing == nil {
		var ret string
		return ret
	}
	return *o.KeyRing
}

// GetKeyRingOk returns a tuple with the KeyRing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigModel) GetKeyRingOk() (*string, bool) {
	if o == nil || o.KeyRing == nil {
		return nil, false
	}
	return o.KeyRing, true
}

// HasKeyRing returns a boolean if a field has been set.
func (o *ConfigModel) HasKeyRing() bool {
	if o != nil && o.KeyRing != nil {
		return true
	}

	return false
}

// SetKeyRing gets a reference to the given string and assigns it to the KeyRing field.
func (o *ConfigModel) SetKeyRing(v string) {
	o.KeyRing = &v
}

func (o ConfigModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.KeyRing != nil {
		toSerialize["keyRing"] = o.KeyRing
	}
	return json.Marshal(toSerialize)
}

type NullableConfigModel struct {
	value *ConfigModel
	isSet bool
}

func (v NullableConfigModel) Get() *ConfigModel {
	return v.value
}

func (v *NullableConfigModel) Set(val *ConfigModel) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigModel) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigModel(val *ConfigModel) *NullableConfigModel {
	return &NullableConfigModel{value: val, isSet: true}
}

func (v NullableConfigModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

