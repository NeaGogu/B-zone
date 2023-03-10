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

// PayloadItem struct for PayloadItem
type PayloadItem struct {
	PayloadKey *string `json:"payloadKey,omitempty"`
	PayloadValue *string `json:"payloadValue,omitempty"`
}

// NewPayloadItem instantiates a new PayloadItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayloadItem() *PayloadItem {
	this := PayloadItem{}
	return &this
}

// NewPayloadItemWithDefaults instantiates a new PayloadItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayloadItemWithDefaults() *PayloadItem {
	this := PayloadItem{}
	return &this
}

// GetPayloadKey returns the PayloadKey field value if set, zero value otherwise.
func (o *PayloadItem) GetPayloadKey() string {
	if o == nil || o.PayloadKey == nil {
		var ret string
		return ret
	}
	return *o.PayloadKey
}

// GetPayloadKeyOk returns a tuple with the PayloadKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayloadItem) GetPayloadKeyOk() (*string, bool) {
	if o == nil || o.PayloadKey == nil {
		return nil, false
	}
	return o.PayloadKey, true
}

// HasPayloadKey returns a boolean if a field has been set.
func (o *PayloadItem) HasPayloadKey() bool {
	if o != nil && o.PayloadKey != nil {
		return true
	}

	return false
}

// SetPayloadKey gets a reference to the given string and assigns it to the PayloadKey field.
func (o *PayloadItem) SetPayloadKey(v string) {
	o.PayloadKey = &v
}

// GetPayloadValue returns the PayloadValue field value if set, zero value otherwise.
func (o *PayloadItem) GetPayloadValue() string {
	if o == nil || o.PayloadValue == nil {
		var ret string
		return ret
	}
	return *o.PayloadValue
}

// GetPayloadValueOk returns a tuple with the PayloadValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PayloadItem) GetPayloadValueOk() (*string, bool) {
	if o == nil || o.PayloadValue == nil {
		return nil, false
	}
	return o.PayloadValue, true
}

// HasPayloadValue returns a boolean if a field has been set.
func (o *PayloadItem) HasPayloadValue() bool {
	if o != nil && o.PayloadValue != nil {
		return true
	}

	return false
}

// SetPayloadValue gets a reference to the given string and assigns it to the PayloadValue field.
func (o *PayloadItem) SetPayloadValue(v string) {
	o.PayloadValue = &v
}

func (o PayloadItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PayloadKey != nil {
		toSerialize["payloadKey"] = o.PayloadKey
	}
	if o.PayloadValue != nil {
		toSerialize["payloadValue"] = o.PayloadValue
	}
	return json.Marshal(toSerialize)
}

type NullablePayloadItem struct {
	value *PayloadItem
	isSet bool
}

func (v NullablePayloadItem) Get() *PayloadItem {
	return v.value
}

func (v *NullablePayloadItem) Set(val *PayloadItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePayloadItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePayloadItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayloadItem(val *PayloadItem) *NullablePayloadItem {
	return &NullablePayloadItem{value: val, isSet: true}
}

func (v NullablePayloadItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayloadItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

