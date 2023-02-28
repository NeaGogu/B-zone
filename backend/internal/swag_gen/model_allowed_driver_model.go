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

// AllowedDriverModel struct for AllowedDriverModel
type AllowedDriverModel struct {
	// User identifier
	Id *int64 `json:"id,omitempty"`
	// User Full Name
	FullName *string `json:"full_name,omitempty"`
	// User Initials
	Initials *string `json:"initials,omitempty"`
}

// NewAllowedDriverModel instantiates a new AllowedDriverModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllowedDriverModel() *AllowedDriverModel {
	this := AllowedDriverModel{}
	return &this
}

// NewAllowedDriverModelWithDefaults instantiates a new AllowedDriverModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllowedDriverModelWithDefaults() *AllowedDriverModel {
	this := AllowedDriverModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AllowedDriverModel) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedDriverModel) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AllowedDriverModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *AllowedDriverModel) SetId(v int64) {
	o.Id = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *AllowedDriverModel) GetFullName() string {
	if o == nil || o.FullName == nil {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedDriverModel) GetFullNameOk() (*string, bool) {
	if o == nil || o.FullName == nil {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *AllowedDriverModel) HasFullName() bool {
	if o != nil && o.FullName != nil {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *AllowedDriverModel) SetFullName(v string) {
	o.FullName = &v
}

// GetInitials returns the Initials field value if set, zero value otherwise.
func (o *AllowedDriverModel) GetInitials() string {
	if o == nil || o.Initials == nil {
		var ret string
		return ret
	}
	return *o.Initials
}

// GetInitialsOk returns a tuple with the Initials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedDriverModel) GetInitialsOk() (*string, bool) {
	if o == nil || o.Initials == nil {
		return nil, false
	}
	return o.Initials, true
}

// HasInitials returns a boolean if a field has been set.
func (o *AllowedDriverModel) HasInitials() bool {
	if o != nil && o.Initials != nil {
		return true
	}

	return false
}

// SetInitials gets a reference to the given string and assigns it to the Initials field.
func (o *AllowedDriverModel) SetInitials(v string) {
	o.Initials = &v
}

func (o AllowedDriverModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.FullName != nil {
		toSerialize["full_name"] = o.FullName
	}
	if o.Initials != nil {
		toSerialize["initials"] = o.Initials
	}
	return json.Marshal(toSerialize)
}

type NullableAllowedDriverModel struct {
	value *AllowedDriverModel
	isSet bool
}

func (v NullableAllowedDriverModel) Get() *AllowedDriverModel {
	return v.value
}

func (v *NullableAllowedDriverModel) Set(val *AllowedDriverModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAllowedDriverModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAllowedDriverModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllowedDriverModel(val *AllowedDriverModel) *NullableAllowedDriverModel {
	return &NullableAllowedDriverModel{value: val, isSet: true}
}

func (v NullableAllowedDriverModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllowedDriverModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


