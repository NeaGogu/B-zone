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

// ApiResponse21 struct for ApiResponse21
type ApiResponse21 struct {
	// Message describing the code
	Message *string `json:"message,omitempty"`
	// Ready
	Type *string `json:"type,omitempty"`
	// 
	Code *float32 `json:"code,omitempty"`
	AdditionalData *map[string]interface{} `json:"additional_data,omitempty"`
}

// NewApiResponse21 instantiates a new ApiResponse21 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiResponse21() *ApiResponse21 {
	this := ApiResponse21{}
	return &this
}

// NewApiResponse21WithDefaults instantiates a new ApiResponse21 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiResponse21WithDefaults() *ApiResponse21 {
	this := ApiResponse21{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ApiResponse21) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponse21) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ApiResponse21) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ApiResponse21) SetMessage(v string) {
	o.Message = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApiResponse21) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponse21) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApiResponse21) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApiResponse21) SetType(v string) {
	o.Type = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ApiResponse21) GetCode() float32 {
	if o == nil || o.Code == nil {
		var ret float32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponse21) GetCodeOk() (*float32, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ApiResponse21) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given float32 and assigns it to the Code field.
func (o *ApiResponse21) SetCode(v float32) {
	o.Code = &v
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *ApiResponse21) GetAdditionalData() map[string]interface{} {
	if o == nil || o.AdditionalData == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponse21) GetAdditionalDataOk() (*map[string]interface{}, bool) {
	if o == nil || o.AdditionalData == nil {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *ApiResponse21) HasAdditionalData() bool {
	if o != nil && o.AdditionalData != nil {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]interface{} and assigns it to the AdditionalData field.
func (o *ApiResponse21) SetAdditionalData(v map[string]interface{}) {
	o.AdditionalData = &v
}

func (o ApiResponse21) MarshalJSON() ([]byte, error) {
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

type NullableApiResponse21 struct {
	value *ApiResponse21
	isSet bool
}

func (v NullableApiResponse21) Get() *ApiResponse21 {
	return v.value
}

func (v *NullableApiResponse21) Set(val *ApiResponse21) {
	v.value = val
	v.isSet = true
}

func (v NullableApiResponse21) IsSet() bool {
	return v.isSet
}

func (v *NullableApiResponse21) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiResponse21(val *ApiResponse21) *NullableApiResponse21 {
	return &NullableApiResponse21{value: val, isSet: true}
}

func (v NullableApiResponse21) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiResponse21) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


