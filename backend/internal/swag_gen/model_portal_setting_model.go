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

// PortalSettingModel struct for PortalSettingModel
type PortalSettingModel struct {
	// Unique Identifier
	Id *int64 `json:"id,omitempty"`
	// Name of portal setting
	Name *string `json:"name,omitempty"`
	// Value of portal setting
	Value *string `json:"value,omitempty"`
	// Name of portal settings group
	PortalSettingGroupName *string `json:"portal_setting_group_name,omitempty"`
	// Name of portal
	PortalName *string `json:"portal_name,omitempty"`
}

// NewPortalSettingModel instantiates a new PortalSettingModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortalSettingModel() *PortalSettingModel {
	this := PortalSettingModel{}
	return &this
}

// NewPortalSettingModelWithDefaults instantiates a new PortalSettingModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortalSettingModelWithDefaults() *PortalSettingModel {
	this := PortalSettingModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PortalSettingModel) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortalSettingModel) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PortalSettingModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *PortalSettingModel) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PortalSettingModel) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortalSettingModel) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PortalSettingModel) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PortalSettingModel) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PortalSettingModel) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortalSettingModel) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PortalSettingModel) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *PortalSettingModel) SetValue(v string) {
	o.Value = &v
}

// GetPortalSettingGroupName returns the PortalSettingGroupName field value if set, zero value otherwise.
func (o *PortalSettingModel) GetPortalSettingGroupName() string {
	if o == nil || o.PortalSettingGroupName == nil {
		var ret string
		return ret
	}
	return *o.PortalSettingGroupName
}

// GetPortalSettingGroupNameOk returns a tuple with the PortalSettingGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortalSettingModel) GetPortalSettingGroupNameOk() (*string, bool) {
	if o == nil || o.PortalSettingGroupName == nil {
		return nil, false
	}
	return o.PortalSettingGroupName, true
}

// HasPortalSettingGroupName returns a boolean if a field has been set.
func (o *PortalSettingModel) HasPortalSettingGroupName() bool {
	if o != nil && o.PortalSettingGroupName != nil {
		return true
	}

	return false
}

// SetPortalSettingGroupName gets a reference to the given string and assigns it to the PortalSettingGroupName field.
func (o *PortalSettingModel) SetPortalSettingGroupName(v string) {
	o.PortalSettingGroupName = &v
}

// GetPortalName returns the PortalName field value if set, zero value otherwise.
func (o *PortalSettingModel) GetPortalName() string {
	if o == nil || o.PortalName == nil {
		var ret string
		return ret
	}
	return *o.PortalName
}

// GetPortalNameOk returns a tuple with the PortalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortalSettingModel) GetPortalNameOk() (*string, bool) {
	if o == nil || o.PortalName == nil {
		return nil, false
	}
	return o.PortalName, true
}

// HasPortalName returns a boolean if a field has been set.
func (o *PortalSettingModel) HasPortalName() bool {
	if o != nil && o.PortalName != nil {
		return true
	}

	return false
}

// SetPortalName gets a reference to the given string and assigns it to the PortalName field.
func (o *PortalSettingModel) SetPortalName(v string) {
	o.PortalName = &v
}

func (o PortalSettingModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.PortalSettingGroupName != nil {
		toSerialize["portal_setting_group_name"] = o.PortalSettingGroupName
	}
	if o.PortalName != nil {
		toSerialize["portal_name"] = o.PortalName
	}
	return json.Marshal(toSerialize)
}

type NullablePortalSettingModel struct {
	value *PortalSettingModel
	isSet bool
}

func (v NullablePortalSettingModel) Get() *PortalSettingModel {
	return v.value
}

func (v *NullablePortalSettingModel) Set(val *PortalSettingModel) {
	v.value = val
	v.isSet = true
}

func (v NullablePortalSettingModel) IsSet() bool {
	return v.isSet
}

func (v *NullablePortalSettingModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortalSettingModel(val *PortalSettingModel) *NullablePortalSettingModel {
	return &NullablePortalSettingModel{value: val, isSet: true}
}

func (v NullablePortalSettingModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortalSettingModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

