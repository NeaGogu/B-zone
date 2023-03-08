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

// TransactionOptionsModel struct for TransactionOptionsModel
type TransactionOptionsModel struct {
	// 
	IncludeRecordInfo *bool `json:"include_record_info,omitempty"`
}

// NewTransactionOptionsModel instantiates a new TransactionOptionsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionOptionsModel() *TransactionOptionsModel {
	this := TransactionOptionsModel{}
	return &this
}

// NewTransactionOptionsModelWithDefaults instantiates a new TransactionOptionsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionOptionsModelWithDefaults() *TransactionOptionsModel {
	this := TransactionOptionsModel{}
	return &this
}

// GetIncludeRecordInfo returns the IncludeRecordInfo field value if set, zero value otherwise.
func (o *TransactionOptionsModel) GetIncludeRecordInfo() bool {
	if o == nil || o.IncludeRecordInfo == nil {
		var ret bool
		return ret
	}
	return *o.IncludeRecordInfo
}

// GetIncludeRecordInfoOk returns a tuple with the IncludeRecordInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionOptionsModel) GetIncludeRecordInfoOk() (*bool, bool) {
	if o == nil || o.IncludeRecordInfo == nil {
		return nil, false
	}
	return o.IncludeRecordInfo, true
}

// HasIncludeRecordInfo returns a boolean if a field has been set.
func (o *TransactionOptionsModel) HasIncludeRecordInfo() bool {
	if o != nil && o.IncludeRecordInfo != nil {
		return true
	}

	return false
}

// SetIncludeRecordInfo gets a reference to the given bool and assigns it to the IncludeRecordInfo field.
func (o *TransactionOptionsModel) SetIncludeRecordInfo(v bool) {
	o.IncludeRecordInfo = &v
}

func (o TransactionOptionsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IncludeRecordInfo != nil {
		toSerialize["include_record_info"] = o.IncludeRecordInfo
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionOptionsModel struct {
	value *TransactionOptionsModel
	isSet bool
}

func (v NullableTransactionOptionsModel) Get() *TransactionOptionsModel {
	return v.value
}

func (v *NullableTransactionOptionsModel) Set(val *TransactionOptionsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionOptionsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionOptionsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionOptionsModel(val *TransactionOptionsModel) *NullableTransactionOptionsModel {
	return &NullableTransactionOptionsModel{value: val, isSet: true}
}

func (v NullableTransactionOptionsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionOptionsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


