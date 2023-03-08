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

// EquipmentCreateResponse struct for EquipmentCreateResponse
type EquipmentCreateResponse struct {
	// 
	Code *int32 `json:"code,omitempty"`
	// Ready
	Type *string `json:"type,omitempty"`
	// Message describing the code
	Message *string `json:"message,omitempty"`
	// 
	AdditionalData *map[string]interface{} `json:"additional_data,omitempty"`
}

// NewEquipmentCreateResponse instantiates a new EquipmentCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentCreateResponse() *EquipmentCreateResponse {
	this := EquipmentCreateResponse{}
	return &this
}

// NewEquipmentCreateResponseWithDefaults instantiates a new EquipmentCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentCreateResponseWithDefaults() *EquipmentCreateResponse {
	this := EquipmentCreateResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *EquipmentCreateResponse) GetCode() int32 {
	if o == nil || o.Code == nil {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentCreateResponse) GetCodeOk() (*int32, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *EquipmentCreateResponse) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *EquipmentCreateResponse) SetCode(v int32) {
	o.Code = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EquipmentCreateResponse) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentCreateResponse) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EquipmentCreateResponse) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EquipmentCreateResponse) SetType(v string) {
	o.Type = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *EquipmentCreateResponse) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentCreateResponse) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *EquipmentCreateResponse) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *EquipmentCreateResponse) SetMessage(v string) {
	o.Message = &v
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *EquipmentCreateResponse) GetAdditionalData() map[string]interface{} {
	if o == nil || o.AdditionalData == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentCreateResponse) GetAdditionalDataOk() (*map[string]interface{}, bool) {
	if o == nil || o.AdditionalData == nil {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *EquipmentCreateResponse) HasAdditionalData() bool {
	if o != nil && o.AdditionalData != nil {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]interface{} and assigns it to the AdditionalData field.
func (o *EquipmentCreateResponse) SetAdditionalData(v map[string]interface{}) {
	o.AdditionalData = &v
}

func (o EquipmentCreateResponse) MarshalJSON() ([]byte, error) {
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

type NullableEquipmentCreateResponse struct {
	value *EquipmentCreateResponse
	isSet bool
}

func (v NullableEquipmentCreateResponse) Get() *EquipmentCreateResponse {
	return v.value
}

func (v *NullableEquipmentCreateResponse) Set(val *EquipmentCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentCreateResponse(val *EquipmentCreateResponse) *NullableEquipmentCreateResponse {
	return &NullableEquipmentCreateResponse{value: val, isSet: true}
}

func (v NullableEquipmentCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


