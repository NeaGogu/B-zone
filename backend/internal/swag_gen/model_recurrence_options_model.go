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

// RecurrenceOptionsModel struct for RecurrenceOptionsModel
type RecurrenceOptionsModel struct {
	// 
	IncludeRecurrenceMetaData *bool `json:"include_recurrence_meta_data,omitempty"`
	// 
	IncludeRecurrenceLinks *bool `json:"include_recurrence_links,omitempty"`
	// 
	IncludeRecurrenceBase *bool `json:"include_recurrence_base,omitempty"`
}

// NewRecurrenceOptionsModel instantiates a new RecurrenceOptionsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecurrenceOptionsModel() *RecurrenceOptionsModel {
	this := RecurrenceOptionsModel{}
	return &this
}

// NewRecurrenceOptionsModelWithDefaults instantiates a new RecurrenceOptionsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecurrenceOptionsModelWithDefaults() *RecurrenceOptionsModel {
	this := RecurrenceOptionsModel{}
	return &this
}

// GetIncludeRecurrenceMetaData returns the IncludeRecurrenceMetaData field value if set, zero value otherwise.
func (o *RecurrenceOptionsModel) GetIncludeRecurrenceMetaData() bool {
	if o == nil || o.IncludeRecurrenceMetaData == nil {
		var ret bool
		return ret
	}
	return *o.IncludeRecurrenceMetaData
}

// GetIncludeRecurrenceMetaDataOk returns a tuple with the IncludeRecurrenceMetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurrenceOptionsModel) GetIncludeRecurrenceMetaDataOk() (*bool, bool) {
	if o == nil || o.IncludeRecurrenceMetaData == nil {
		return nil, false
	}
	return o.IncludeRecurrenceMetaData, true
}

// HasIncludeRecurrenceMetaData returns a boolean if a field has been set.
func (o *RecurrenceOptionsModel) HasIncludeRecurrenceMetaData() bool {
	if o != nil && o.IncludeRecurrenceMetaData != nil {
		return true
	}

	return false
}

// SetIncludeRecurrenceMetaData gets a reference to the given bool and assigns it to the IncludeRecurrenceMetaData field.
func (o *RecurrenceOptionsModel) SetIncludeRecurrenceMetaData(v bool) {
	o.IncludeRecurrenceMetaData = &v
}

// GetIncludeRecurrenceLinks returns the IncludeRecurrenceLinks field value if set, zero value otherwise.
func (o *RecurrenceOptionsModel) GetIncludeRecurrenceLinks() bool {
	if o == nil || o.IncludeRecurrenceLinks == nil {
		var ret bool
		return ret
	}
	return *o.IncludeRecurrenceLinks
}

// GetIncludeRecurrenceLinksOk returns a tuple with the IncludeRecurrenceLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurrenceOptionsModel) GetIncludeRecurrenceLinksOk() (*bool, bool) {
	if o == nil || o.IncludeRecurrenceLinks == nil {
		return nil, false
	}
	return o.IncludeRecurrenceLinks, true
}

// HasIncludeRecurrenceLinks returns a boolean if a field has been set.
func (o *RecurrenceOptionsModel) HasIncludeRecurrenceLinks() bool {
	if o != nil && o.IncludeRecurrenceLinks != nil {
		return true
	}

	return false
}

// SetIncludeRecurrenceLinks gets a reference to the given bool and assigns it to the IncludeRecurrenceLinks field.
func (o *RecurrenceOptionsModel) SetIncludeRecurrenceLinks(v bool) {
	o.IncludeRecurrenceLinks = &v
}

// GetIncludeRecurrenceBase returns the IncludeRecurrenceBase field value if set, zero value otherwise.
func (o *RecurrenceOptionsModel) GetIncludeRecurrenceBase() bool {
	if o == nil || o.IncludeRecurrenceBase == nil {
		var ret bool
		return ret
	}
	return *o.IncludeRecurrenceBase
}

// GetIncludeRecurrenceBaseOk returns a tuple with the IncludeRecurrenceBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurrenceOptionsModel) GetIncludeRecurrenceBaseOk() (*bool, bool) {
	if o == nil || o.IncludeRecurrenceBase == nil {
		return nil, false
	}
	return o.IncludeRecurrenceBase, true
}

// HasIncludeRecurrenceBase returns a boolean if a field has been set.
func (o *RecurrenceOptionsModel) HasIncludeRecurrenceBase() bool {
	if o != nil && o.IncludeRecurrenceBase != nil {
		return true
	}

	return false
}

// SetIncludeRecurrenceBase gets a reference to the given bool and assigns it to the IncludeRecurrenceBase field.
func (o *RecurrenceOptionsModel) SetIncludeRecurrenceBase(v bool) {
	o.IncludeRecurrenceBase = &v
}

func (o RecurrenceOptionsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IncludeRecurrenceMetaData != nil {
		toSerialize["include_recurrence_meta_data"] = o.IncludeRecurrenceMetaData
	}
	if o.IncludeRecurrenceLinks != nil {
		toSerialize["include_recurrence_links"] = o.IncludeRecurrenceLinks
	}
	if o.IncludeRecurrenceBase != nil {
		toSerialize["include_recurrence_base"] = o.IncludeRecurrenceBase
	}
	return json.Marshal(toSerialize)
}

type NullableRecurrenceOptionsModel struct {
	value *RecurrenceOptionsModel
	isSet bool
}

func (v NullableRecurrenceOptionsModel) Get() *RecurrenceOptionsModel {
	return v.value
}

func (v *NullableRecurrenceOptionsModel) Set(val *RecurrenceOptionsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableRecurrenceOptionsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableRecurrenceOptionsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecurrenceOptionsModel(val *RecurrenceOptionsModel) *NullableRecurrenceOptionsModel {
	return &NullableRecurrenceOptionsModel{value: val, isSet: true}
}

func (v NullableRecurrenceOptionsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecurrenceOptionsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


