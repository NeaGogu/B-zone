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

// EquipmentDeleteNotFoundResponse struct for EquipmentDeleteNotFoundResponse
type EquipmentDeleteNotFoundResponse struct {
	// 
	Code *int32 `json:"code,omitempty"`
	// Message describing the code
	Message *string `json:"message,omitempty"`
}

// NewEquipmentDeleteNotFoundResponse instantiates a new EquipmentDeleteNotFoundResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentDeleteNotFoundResponse() *EquipmentDeleteNotFoundResponse {
	this := EquipmentDeleteNotFoundResponse{}
	return &this
}

// NewEquipmentDeleteNotFoundResponseWithDefaults instantiates a new EquipmentDeleteNotFoundResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentDeleteNotFoundResponseWithDefaults() *EquipmentDeleteNotFoundResponse {
	this := EquipmentDeleteNotFoundResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *EquipmentDeleteNotFoundResponse) GetCode() int32 {
	if o == nil || o.Code == nil {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentDeleteNotFoundResponse) GetCodeOk() (*int32, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *EquipmentDeleteNotFoundResponse) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *EquipmentDeleteNotFoundResponse) SetCode(v int32) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *EquipmentDeleteNotFoundResponse) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentDeleteNotFoundResponse) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *EquipmentDeleteNotFoundResponse) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *EquipmentDeleteNotFoundResponse) SetMessage(v string) {
	o.Message = &v
}

func (o EquipmentDeleteNotFoundResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableEquipmentDeleteNotFoundResponse struct {
	value *EquipmentDeleteNotFoundResponse
	isSet bool
}

func (v NullableEquipmentDeleteNotFoundResponse) Get() *EquipmentDeleteNotFoundResponse {
	return v.value
}

func (v *NullableEquipmentDeleteNotFoundResponse) Set(val *EquipmentDeleteNotFoundResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentDeleteNotFoundResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentDeleteNotFoundResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentDeleteNotFoundResponse(val *EquipmentDeleteNotFoundResponse) *NullableEquipmentDeleteNotFoundResponse {
	return &NullableEquipmentDeleteNotFoundResponse{value: val, isSet: true}
}

func (v NullableEquipmentDeleteNotFoundResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentDeleteNotFoundResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

