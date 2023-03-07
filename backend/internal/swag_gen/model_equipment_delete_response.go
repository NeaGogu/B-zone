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

// EquipmentDeleteResponse struct for EquipmentDeleteResponse
type EquipmentDeleteResponse struct {
	// 
	Code *int32 `json:"code,omitempty"`
	// Ready
	Type *string `json:"type,omitempty"`
	// Message describing the code
	Message *string `json:"message,omitempty"`
	// 
	AdditionalData *map[string]interface{} `json:"additional_data,omitempty"`
}

// NewEquipmentDeleteResponse instantiates a new EquipmentDeleteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentDeleteResponse() *EquipmentDeleteResponse {
	this := EquipmentDeleteResponse{}
	return &this
}

// NewEquipmentDeleteResponseWithDefaults instantiates a new EquipmentDeleteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentDeleteResponseWithDefaults() *EquipmentDeleteResponse {
	this := EquipmentDeleteResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *EquipmentDeleteResponse) GetCode() int32 {
	if o == nil || o.Code == nil {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentDeleteResponse) GetCodeOk() (*int32, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *EquipmentDeleteResponse) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *EquipmentDeleteResponse) SetCode(v int32) {
	o.Code = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EquipmentDeleteResponse) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentDeleteResponse) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EquipmentDeleteResponse) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EquipmentDeleteResponse) SetType(v string) {
	o.Type = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *EquipmentDeleteResponse) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentDeleteResponse) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *EquipmentDeleteResponse) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *EquipmentDeleteResponse) SetMessage(v string) {
	o.Message = &v
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *EquipmentDeleteResponse) GetAdditionalData() map[string]interface{} {
	if o == nil || o.AdditionalData == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentDeleteResponse) GetAdditionalDataOk() (*map[string]interface{}, bool) {
	if o == nil || o.AdditionalData == nil {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *EquipmentDeleteResponse) HasAdditionalData() bool {
	if o != nil && o.AdditionalData != nil {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]interface{} and assigns it to the AdditionalData field.
func (o *EquipmentDeleteResponse) SetAdditionalData(v map[string]interface{}) {
	o.AdditionalData = &v
}

func (o EquipmentDeleteResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.AdditionalData != nil {
		toSerialize["additional_data"] = o.AdditionalData
	}
	return json.Marshal(toSerialize)
}

type NullableEquipmentDeleteResponse struct {
	value *EquipmentDeleteResponse
	isSet bool
}

func (v NullableEquipmentDeleteResponse) Get() *EquipmentDeleteResponse {
	return v.value
}

func (v *NullableEquipmentDeleteResponse) Set(val *EquipmentDeleteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentDeleteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentDeleteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentDeleteResponse(val *EquipmentDeleteResponse) *NullableEquipmentDeleteResponse {
	return &NullableEquipmentDeleteResponse{value: val, isSet: true}
}

func (v NullableEquipmentDeleteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentDeleteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

