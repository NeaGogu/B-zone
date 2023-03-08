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

// CommunicationDeliveryMethodModel struct for CommunicationDeliveryMethodModel
type CommunicationDeliveryMethodModel struct {
	// Identifier
	Id *int32 `json:"id,omitempty"`
	// Name of this delivery method
	Name *string `json:"name,omitempty"`
}

// NewCommunicationDeliveryMethodModel instantiates a new CommunicationDeliveryMethodModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommunicationDeliveryMethodModel() *CommunicationDeliveryMethodModel {
	this := CommunicationDeliveryMethodModel{}
	return &this
}

// NewCommunicationDeliveryMethodModelWithDefaults instantiates a new CommunicationDeliveryMethodModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommunicationDeliveryMethodModelWithDefaults() *CommunicationDeliveryMethodModel {
	this := CommunicationDeliveryMethodModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CommunicationDeliveryMethodModel) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationDeliveryMethodModel) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CommunicationDeliveryMethodModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CommunicationDeliveryMethodModel) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CommunicationDeliveryMethodModel) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationDeliveryMethodModel) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CommunicationDeliveryMethodModel) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CommunicationDeliveryMethodModel) SetName(v string) {
	o.Name = &v
}

func (o CommunicationDeliveryMethodModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableCommunicationDeliveryMethodModel struct {
	value *CommunicationDeliveryMethodModel
	isSet bool
}

func (v NullableCommunicationDeliveryMethodModel) Get() *CommunicationDeliveryMethodModel {
	return v.value
}

func (v *NullableCommunicationDeliveryMethodModel) Set(val *CommunicationDeliveryMethodModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCommunicationDeliveryMethodModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCommunicationDeliveryMethodModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommunicationDeliveryMethodModel(val *CommunicationDeliveryMethodModel) *NullableCommunicationDeliveryMethodModel {
	return &NullableCommunicationDeliveryMethodModel{value: val, isSet: true}
}

func (v NullableCommunicationDeliveryMethodModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommunicationDeliveryMethodModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


