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

// ApiResponse11 struct for ApiResponse11
type ApiResponse11 struct {
	// Message describing the code
	Message *string `json:"message,omitempty"`
	// Ready
	Type *string `json:"type,omitempty"`
	// 
	Code *float32 `json:"code,omitempty"`
	AdditionalData *map[string]interface{} `json:"additional_data,omitempty"`
}

// NewApiResponse11 instantiates a new ApiResponse11 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiResponse11() *ApiResponse11 {
	this := ApiResponse11{}
	return &this
}

// NewApiResponse11WithDefaults instantiates a new ApiResponse11 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiResponse11WithDefaults() *ApiResponse11 {
	this := ApiResponse11{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ApiResponse11) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponse11) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ApiResponse11) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ApiResponse11) SetMessage(v string) {
	o.Message = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApiResponse11) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponse11) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApiResponse11) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApiResponse11) SetType(v string) {
	o.Type = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ApiResponse11) GetCode() float32 {
	if o == nil || o.Code == nil {
		var ret float32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponse11) GetCodeOk() (*float32, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ApiResponse11) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given float32 and assigns it to the Code field.
func (o *ApiResponse11) SetCode(v float32) {
	o.Code = &v
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *ApiResponse11) GetAdditionalData() map[string]interface{} {
	if o == nil || o.AdditionalData == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponse11) GetAdditionalDataOk() (*map[string]interface{}, bool) {
	if o == nil || o.AdditionalData == nil {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *ApiResponse11) HasAdditionalData() bool {
	if o != nil && o.AdditionalData != nil {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]interface{} and assigns it to the AdditionalData field.
func (o *ApiResponse11) SetAdditionalData(v map[string]interface{}) {
	o.AdditionalData = &v
}

func (o ApiResponse11) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.AdditionalData != nil {
		toSerialize["additional_data"] = o.AdditionalData
	}
	return json.Marshal(toSerialize)
}

type NullableApiResponse11 struct {
	value *ApiResponse11
	isSet bool
}

func (v NullableApiResponse11) Get() *ApiResponse11 {
	return v.value
}

func (v *NullableApiResponse11) Set(val *ApiResponse11) {
	v.value = val
	v.isSet = true
}

func (v NullableApiResponse11) IsSet() bool {
	return v.isSet
}

func (v *NullableApiResponse11) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiResponse11(val *ApiResponse11) *NullableApiResponse11 {
	return &NullableApiResponse11{value: val, isSet: true}
}

func (v NullableApiResponse11) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiResponse11) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

