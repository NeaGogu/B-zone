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

// AddressFiltersModel struct for AddressFiltersModel
type AddressFiltersModel struct {
	// Unique Identifier(s)
	Id *[]int32 `json:"id,omitempty"`
	// Address Type names
	AddressTypeNames *[]string `json:"address_type_names,omitempty"`
	// Address Tags
	AddressTags *[]string `json:"address_tags,omitempty"`
	// Party ID
	PartyId *int32 `json:"party_id,omitempty"`
	// 
	Code *string `json:"code,omitempty"`
	// 
	UserId *int32 `json:"user_id,omitempty"`
}

// NewAddressFiltersModel instantiates a new AddressFiltersModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressFiltersModel() *AddressFiltersModel {
	this := AddressFiltersModel{}
	return &this
}

// NewAddressFiltersModelWithDefaults instantiates a new AddressFiltersModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressFiltersModelWithDefaults() *AddressFiltersModel {
	this := AddressFiltersModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AddressFiltersModel) GetId() []int32 {
	if o == nil || o.Id == nil {
		var ret []int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressFiltersModel) GetIdOk() (*[]int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AddressFiltersModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given []int32 and assigns it to the Id field.
func (o *AddressFiltersModel) SetId(v []int32) {
	o.Id = &v
}

// GetAddressTypeNames returns the AddressTypeNames field value if set, zero value otherwise.
func (o *AddressFiltersModel) GetAddressTypeNames() []string {
	if o == nil || o.AddressTypeNames == nil {
		var ret []string
		return ret
	}
	return *o.AddressTypeNames
}

// GetAddressTypeNamesOk returns a tuple with the AddressTypeNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressFiltersModel) GetAddressTypeNamesOk() (*[]string, bool) {
	if o == nil || o.AddressTypeNames == nil {
		return nil, false
	}
	return o.AddressTypeNames, true
}

// HasAddressTypeNames returns a boolean if a field has been set.
func (o *AddressFiltersModel) HasAddressTypeNames() bool {
	if o != nil && o.AddressTypeNames != nil {
		return true
	}

	return false
}

// SetAddressTypeNames gets a reference to the given []string and assigns it to the AddressTypeNames field.
func (o *AddressFiltersModel) SetAddressTypeNames(v []string) {
	o.AddressTypeNames = &v
}

// GetAddressTags returns the AddressTags field value if set, zero value otherwise.
func (o *AddressFiltersModel) GetAddressTags() []string {
	if o == nil || o.AddressTags == nil {
		var ret []string
		return ret
	}
	return *o.AddressTags
}

// GetAddressTagsOk returns a tuple with the AddressTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressFiltersModel) GetAddressTagsOk() (*[]string, bool) {
	if o == nil || o.AddressTags == nil {
		return nil, false
	}
	return o.AddressTags, true
}

// HasAddressTags returns a boolean if a field has been set.
func (o *AddressFiltersModel) HasAddressTags() bool {
	if o != nil && o.AddressTags != nil {
		return true
	}

	return false
}

// SetAddressTags gets a reference to the given []string and assigns it to the AddressTags field.
func (o *AddressFiltersModel) SetAddressTags(v []string) {
	o.AddressTags = &v
}

// GetPartyId returns the PartyId field value if set, zero value otherwise.
func (o *AddressFiltersModel) GetPartyId() int32 {
	if o == nil || o.PartyId == nil {
		var ret int32
		return ret
	}
	return *o.PartyId
}

// GetPartyIdOk returns a tuple with the PartyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressFiltersModel) GetPartyIdOk() (*int32, bool) {
	if o == nil || o.PartyId == nil {
		return nil, false
	}
	return o.PartyId, true
}

// HasPartyId returns a boolean if a field has been set.
func (o *AddressFiltersModel) HasPartyId() bool {
	if o != nil && o.PartyId != nil {
		return true
	}

	return false
}

// SetPartyId gets a reference to the given int32 and assigns it to the PartyId field.
func (o *AddressFiltersModel) SetPartyId(v int32) {
	o.PartyId = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AddressFiltersModel) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressFiltersModel) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AddressFiltersModel) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *AddressFiltersModel) SetCode(v string) {
	o.Code = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AddressFiltersModel) GetUserId() int32 {
	if o == nil || o.UserId == nil {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressFiltersModel) GetUserIdOk() (*int32, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AddressFiltersModel) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *AddressFiltersModel) SetUserId(v int32) {
	o.UserId = &v
}

func (o AddressFiltersModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.AddressTypeNames != nil {
		toSerialize["address_type_names"] = o.AddressTypeNames
	}
	if o.AddressTags != nil {
		toSerialize["address_tags"] = o.AddressTags
	}
	if o.PartyId != nil {
		toSerialize["party_id"] = o.PartyId
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.UserId != nil {
		toSerialize["user_id"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullableAddressFiltersModel struct {
	value *AddressFiltersModel
	isSet bool
}

func (v NullableAddressFiltersModel) Get() *AddressFiltersModel {
	return v.value
}

func (v *NullableAddressFiltersModel) Set(val *AddressFiltersModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressFiltersModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressFiltersModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressFiltersModel(val *AddressFiltersModel) *NullableAddressFiltersModel {
	return &NullableAddressFiltersModel{value: val, isSet: true}
}

func (v NullableAddressFiltersModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressFiltersModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

