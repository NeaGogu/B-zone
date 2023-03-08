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

// CheckAvailabilityCacheDataModelActivityAddress 
type CheckAvailabilityCacheDataModelActivityAddress struct {
	// 
	Zipcode string `json:"zipcode"`
	// 
	IsoCountry string `json:"iso_country"`
}

// NewCheckAvailabilityCacheDataModelActivityAddress instantiates a new CheckAvailabilityCacheDataModelActivityAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckAvailabilityCacheDataModelActivityAddress(zipcode string, isoCountry string, ) *CheckAvailabilityCacheDataModelActivityAddress {
	this := CheckAvailabilityCacheDataModelActivityAddress{}
	this.Zipcode = zipcode
	this.IsoCountry = isoCountry
	return &this
}

// NewCheckAvailabilityCacheDataModelActivityAddressWithDefaults instantiates a new CheckAvailabilityCacheDataModelActivityAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckAvailabilityCacheDataModelActivityAddressWithDefaults() *CheckAvailabilityCacheDataModelActivityAddress {
	this := CheckAvailabilityCacheDataModelActivityAddress{}
	return &this
}

// GetZipcode returns the Zipcode field value
func (o *CheckAvailabilityCacheDataModelActivityAddress) GetZipcode() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Zipcode
}

// GetZipcodeOk returns a tuple with the Zipcode field value
// and a boolean to check if the value has been set.
func (o *CheckAvailabilityCacheDataModelActivityAddress) GetZipcodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Zipcode, true
}

// SetZipcode sets field value
func (o *CheckAvailabilityCacheDataModelActivityAddress) SetZipcode(v string) {
	o.Zipcode = v
}

// GetIsoCountry returns the IsoCountry field value
func (o *CheckAvailabilityCacheDataModelActivityAddress) GetIsoCountry() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.IsoCountry
}

// GetIsoCountryOk returns a tuple with the IsoCountry field value
// and a boolean to check if the value has been set.
func (o *CheckAvailabilityCacheDataModelActivityAddress) GetIsoCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsoCountry, true
}

// SetIsoCountry sets field value
func (o *CheckAvailabilityCacheDataModelActivityAddress) SetIsoCountry(v string) {
	o.IsoCountry = v
}

func (o CheckAvailabilityCacheDataModelActivityAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["zipcode"] = o.Zipcode
	}
	if true {
		toSerialize["iso_country"] = o.IsoCountry
	}
	return json.Marshal(toSerialize)
}

type NullableCheckAvailabilityCacheDataModelActivityAddress struct {
	value *CheckAvailabilityCacheDataModelActivityAddress
	isSet bool
}

func (v NullableCheckAvailabilityCacheDataModelActivityAddress) Get() *CheckAvailabilityCacheDataModelActivityAddress {
	return v.value
}

func (v *NullableCheckAvailabilityCacheDataModelActivityAddress) Set(val *CheckAvailabilityCacheDataModelActivityAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckAvailabilityCacheDataModelActivityAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckAvailabilityCacheDataModelActivityAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckAvailabilityCacheDataModelActivityAddress(val *CheckAvailabilityCacheDataModelActivityAddress) *NullableCheckAvailabilityCacheDataModelActivityAddress {
	return &NullableCheckAvailabilityCacheDataModelActivityAddress{value: val, isSet: true}
}

func (v NullableCheckAvailabilityCacheDataModelActivityAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckAvailabilityCacheDataModelActivityAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


