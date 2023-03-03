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

// ApiResponse403 struct for ApiResponse403
type ApiResponse403 struct {
	// Error code
	Code *int32 `json:"code,omitempty"`
	// Message describing the error
	Message *string `json:"message,omitempty"`
}

// NewApiResponse403 instantiates a new ApiResponse403 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiResponse403() *ApiResponse403 {
	this := ApiResponse403{}
	return &this
}

// NewApiResponse403WithDefaults instantiates a new ApiResponse403 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiResponse403WithDefaults() *ApiResponse403 {
	this := ApiResponse403{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ApiResponse403) GetCode() int32 {
	if o == nil || o.Code == nil {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponse403) GetCodeOk() (*int32, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ApiResponse403) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *ApiResponse403) SetCode(v int32) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ApiResponse403) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResponse403) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ApiResponse403) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ApiResponse403) SetMessage(v string) {
	o.Message = &v
}

func (o ApiResponse403) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableApiResponse403 struct {
	value *ApiResponse403
	isSet bool
}

func (v NullableApiResponse403) Get() *ApiResponse403 {
	return v.value
}

func (v *NullableApiResponse403) Set(val *ApiResponse403) {
	v.value = val
	v.isSet = true
}

func (v NullableApiResponse403) IsSet() bool {
	return v.isSet
}

func (v *NullableApiResponse403) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiResponse403(val *ApiResponse403) *NullableApiResponse403 {
	return &NullableApiResponse403{value: val, isSet: true}
}

func (v NullableApiResponse403) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiResponse403) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


