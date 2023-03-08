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

// FileCopyArguments struct for FileCopyArguments
type FileCopyArguments struct {
	// 
	SourceFileId *int64 `json:"source_file_id,omitempty"`
	// 
	TargetObjectId *int64 `json:"target_object_id,omitempty"`
	// 
	TargetObjectType *int64 `json:"target_object_type,omitempty"`
}

// NewFileCopyArguments instantiates a new FileCopyArguments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileCopyArguments() *FileCopyArguments {
	this := FileCopyArguments{}
	return &this
}

// NewFileCopyArgumentsWithDefaults instantiates a new FileCopyArguments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileCopyArgumentsWithDefaults() *FileCopyArguments {
	this := FileCopyArguments{}
	return &this
}

// GetSourceFileId returns the SourceFileId field value if set, zero value otherwise.
func (o *FileCopyArguments) GetSourceFileId() int64 {
	if o == nil || o.SourceFileId == nil {
		var ret int64
		return ret
	}
	return *o.SourceFileId
}

// GetSourceFileIdOk returns a tuple with the SourceFileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileCopyArguments) GetSourceFileIdOk() (*int64, bool) {
	if o == nil || o.SourceFileId == nil {
		return nil, false
	}
	return o.SourceFileId, true
}

// HasSourceFileId returns a boolean if a field has been set.
func (o *FileCopyArguments) HasSourceFileId() bool {
	if o != nil && o.SourceFileId != nil {
		return true
	}

	return false
}

// SetSourceFileId gets a reference to the given int64 and assigns it to the SourceFileId field.
func (o *FileCopyArguments) SetSourceFileId(v int64) {
	o.SourceFileId = &v
}

// GetTargetObjectId returns the TargetObjectId field value if set, zero value otherwise.
func (o *FileCopyArguments) GetTargetObjectId() int64 {
	if o == nil || o.TargetObjectId == nil {
		var ret int64
		return ret
	}
	return *o.TargetObjectId
}

// GetTargetObjectIdOk returns a tuple with the TargetObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileCopyArguments) GetTargetObjectIdOk() (*int64, bool) {
	if o == nil || o.TargetObjectId == nil {
		return nil, false
	}
	return o.TargetObjectId, true
}

// HasTargetObjectId returns a boolean if a field has been set.
func (o *FileCopyArguments) HasTargetObjectId() bool {
	if o != nil && o.TargetObjectId != nil {
		return true
	}

	return false
}

// SetTargetObjectId gets a reference to the given int64 and assigns it to the TargetObjectId field.
func (o *FileCopyArguments) SetTargetObjectId(v int64) {
	o.TargetObjectId = &v
}

// GetTargetObjectType returns the TargetObjectType field value if set, zero value otherwise.
func (o *FileCopyArguments) GetTargetObjectType() int64 {
	if o == nil || o.TargetObjectType == nil {
		var ret int64
		return ret
	}
	return *o.TargetObjectType
}

// GetTargetObjectTypeOk returns a tuple with the TargetObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileCopyArguments) GetTargetObjectTypeOk() (*int64, bool) {
	if o == nil || o.TargetObjectType == nil {
		return nil, false
	}
	return o.TargetObjectType, true
}

// HasTargetObjectType returns a boolean if a field has been set.
func (o *FileCopyArguments) HasTargetObjectType() bool {
	if o != nil && o.TargetObjectType != nil {
		return true
	}

	return false
}

// SetTargetObjectType gets a reference to the given int64 and assigns it to the TargetObjectType field.
func (o *FileCopyArguments) SetTargetObjectType(v int64) {
	o.TargetObjectType = &v
}

func (o FileCopyArguments) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SourceFileId != nil {
		toSerialize["source_file_id"] = o.SourceFileId
	}
	if o.TargetObjectId != nil {
		toSerialize["target_object_id"] = o.TargetObjectId
	}
	if o.TargetObjectType != nil {
		toSerialize["target_object_type"] = o.TargetObjectType
	}
	return json.Marshal(toSerialize)
}

type NullableFileCopyArguments struct {
	value *FileCopyArguments
	isSet bool
}

func (v NullableFileCopyArguments) Get() *FileCopyArguments {
	return v.value
}

func (v *NullableFileCopyArguments) Set(val *FileCopyArguments) {
	v.value = val
	v.isSet = true
}

func (v NullableFileCopyArguments) IsSet() bool {
	return v.isSet
}

func (v *NullableFileCopyArguments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileCopyArguments(val *FileCopyArguments) *NullableFileCopyArguments {
	return &NullableFileCopyArguments{value: val, isSet: true}
}

func (v NullableFileCopyArguments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileCopyArguments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


