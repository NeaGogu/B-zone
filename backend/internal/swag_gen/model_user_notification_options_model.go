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

// UserNotificationOptionsModel struct for UserNotificationOptionsModel
type UserNotificationOptionsModel struct {
	// 
	IncludeNotification *bool `json:"include_notification,omitempty"`
}

// NewUserNotificationOptionsModel instantiates a new UserNotificationOptionsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserNotificationOptionsModel() *UserNotificationOptionsModel {
	this := UserNotificationOptionsModel{}
	return &this
}

// NewUserNotificationOptionsModelWithDefaults instantiates a new UserNotificationOptionsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserNotificationOptionsModelWithDefaults() *UserNotificationOptionsModel {
	this := UserNotificationOptionsModel{}
	return &this
}

// GetIncludeNotification returns the IncludeNotification field value if set, zero value otherwise.
func (o *UserNotificationOptionsModel) GetIncludeNotification() bool {
	if o == nil || o.IncludeNotification == nil {
		var ret bool
		return ret
	}
	return *o.IncludeNotification
}

// GetIncludeNotificationOk returns a tuple with the IncludeNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserNotificationOptionsModel) GetIncludeNotificationOk() (*bool, bool) {
	if o == nil || o.IncludeNotification == nil {
		return nil, false
	}
	return o.IncludeNotification, true
}

// HasIncludeNotification returns a boolean if a field has been set.
func (o *UserNotificationOptionsModel) HasIncludeNotification() bool {
	if o != nil && o.IncludeNotification != nil {
		return true
	}

	return false
}

// SetIncludeNotification gets a reference to the given bool and assigns it to the IncludeNotification field.
func (o *UserNotificationOptionsModel) SetIncludeNotification(v bool) {
	o.IncludeNotification = &v
}

func (o UserNotificationOptionsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IncludeNotification != nil {
		toSerialize["include_notification"] = o.IncludeNotification
	}
	return json.Marshal(toSerialize)
}

type NullableUserNotificationOptionsModel struct {
	value *UserNotificationOptionsModel
	isSet bool
}

func (v NullableUserNotificationOptionsModel) Get() *UserNotificationOptionsModel {
	return v.value
}

func (v *NullableUserNotificationOptionsModel) Set(val *UserNotificationOptionsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableUserNotificationOptionsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableUserNotificationOptionsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserNotificationOptionsModel(val *UserNotificationOptionsModel) *NullableUserNotificationOptionsModel {
	return &NullableUserNotificationOptionsModel{value: val, isSet: true}
}

func (v NullableUserNotificationOptionsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserNotificationOptionsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

