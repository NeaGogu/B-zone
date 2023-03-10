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

// NotificationCategoryOptionsModel struct for NotificationCategoryOptionsModel
type NotificationCategoryOptionsModel struct {
	// 
	IncludeRecordInfo *bool `json:"include_record_info,omitempty"`
}

// NewNotificationCategoryOptionsModel instantiates a new NotificationCategoryOptionsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationCategoryOptionsModel() *NotificationCategoryOptionsModel {
	this := NotificationCategoryOptionsModel{}
	return &this
}

// NewNotificationCategoryOptionsModelWithDefaults instantiates a new NotificationCategoryOptionsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationCategoryOptionsModelWithDefaults() *NotificationCategoryOptionsModel {
	this := NotificationCategoryOptionsModel{}
	return &this
}

// GetIncludeRecordInfo returns the IncludeRecordInfo field value if set, zero value otherwise.
func (o *NotificationCategoryOptionsModel) GetIncludeRecordInfo() bool {
	if o == nil || o.IncludeRecordInfo == nil {
		var ret bool
		return ret
	}
	return *o.IncludeRecordInfo
}

// GetIncludeRecordInfoOk returns a tuple with the IncludeRecordInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationCategoryOptionsModel) GetIncludeRecordInfoOk() (*bool, bool) {
	if o == nil || o.IncludeRecordInfo == nil {
		return nil, false
	}
	return o.IncludeRecordInfo, true
}

// HasIncludeRecordInfo returns a boolean if a field has been set.
func (o *NotificationCategoryOptionsModel) HasIncludeRecordInfo() bool {
	if o != nil && o.IncludeRecordInfo != nil {
		return true
	}

	return false
}

// SetIncludeRecordInfo gets a reference to the given bool and assigns it to the IncludeRecordInfo field.
func (o *NotificationCategoryOptionsModel) SetIncludeRecordInfo(v bool) {
	o.IncludeRecordInfo = &v
}

func (o NotificationCategoryOptionsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IncludeRecordInfo != nil {
		toSerialize["include_record_info"] = o.IncludeRecordInfo
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationCategoryOptionsModel struct {
	value *NotificationCategoryOptionsModel
	isSet bool
}

func (v NullableNotificationCategoryOptionsModel) Get() *NotificationCategoryOptionsModel {
	return v.value
}

func (v *NullableNotificationCategoryOptionsModel) Set(val *NotificationCategoryOptionsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationCategoryOptionsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationCategoryOptionsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationCategoryOptionsModel(val *NotificationCategoryOptionsModel) *NullableNotificationCategoryOptionsModel {
	return &NullableNotificationCategoryOptionsModel{value: val, isSet: true}
}

func (v NullableNotificationCategoryOptionsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationCategoryOptionsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

