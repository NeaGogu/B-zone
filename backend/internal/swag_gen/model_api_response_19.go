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

// ApiResponse19 struct for ApiResponse19
type ApiResponse19 struct {
	// Message describing the error
	Message *string `json:"message,omitempty"`
	// Error code
	Code *float32 `json:"code,omitempty"`
}

// NewApiResponse19 instantiates a new ApiResponse19 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiResponse19() *ApiResponse19 {
	this := ApiResponse19{}
	return &this
}

// NewApiResponse19WithDefaults instantiates a new ApiResponse19 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiResponse19WithDefaults() *ApiResponse19 {
	this := ApiResponse19{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ApiResponse19) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponse19) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ApiResponse19) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ApiResponse19) SetMessage(v string) {
	o.Message = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ApiResponse19) GetCode() float32 {
	if o == nil || o.Code == nil {
		var ret float32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponse19) GetCodeOk() (*float32, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ApiResponse19) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given float32 and assigns it to the Code field.
func (o *ApiResponse19) SetCode(v float32) {
	o.Code = &v
}

func (o ApiResponse19) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	return json.Marshal(toSerialize)
}

type NullableApiResponse19 struct {
	value *ApiResponse19
	isSet bool
}

func (v NullableApiResponse19) Get() *ApiResponse19 {
	return v.value
}

func (v *NullableApiResponse19) Set(val *ApiResponse19) {
	v.value = val
	v.isSet = true
}

func (v NullableApiResponse19) IsSet() bool {
	return v.isSet
}

func (v *NullableApiResponse19) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiResponse19(val *ApiResponse19) *NullableApiResponse19 {
	return &NullableApiResponse19{value: val, isSet: true}
}

func (v NullableApiResponse19) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiResponse19) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


