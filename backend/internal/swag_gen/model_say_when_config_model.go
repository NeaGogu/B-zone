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

// SayWhenConfigModel struct for SayWhenConfigModel
type SayWhenConfigModel struct {
	// 
	PlannerReferenceMapping *string `json:"plannerReferenceMapping,omitempty"`
	// 
	MetaData *string `json:"metaData,omitempty"`
	// 
	KeyRing *string `json:"keyRing,omitempty"`
	// 
	VisitTypeMap *string `json:"visitTypeMap,omitempty"`
}

// NewSayWhenConfigModel instantiates a new SayWhenConfigModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSayWhenConfigModel() *SayWhenConfigModel {
	this := SayWhenConfigModel{}
	return &this
}

// NewSayWhenConfigModelWithDefaults instantiates a new SayWhenConfigModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSayWhenConfigModelWithDefaults() *SayWhenConfigModel {
	this := SayWhenConfigModel{}
	return &this
}

// GetPlannerReferenceMapping returns the PlannerReferenceMapping field value if set, zero value otherwise.
func (o *SayWhenConfigModel) GetPlannerReferenceMapping() string {
	if o == nil || o.PlannerReferenceMapping == nil {
		var ret string
		return ret
	}
	return *o.PlannerReferenceMapping
}

// GetPlannerReferenceMappingOk returns a tuple with the PlannerReferenceMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SayWhenConfigModel) GetPlannerReferenceMappingOk() (*string, bool) {
	if o == nil || o.PlannerReferenceMapping == nil {
		return nil, false
	}
	return o.PlannerReferenceMapping, true
}

// HasPlannerReferenceMapping returns a boolean if a field has been set.
func (o *SayWhenConfigModel) HasPlannerReferenceMapping() bool {
	if o != nil && o.PlannerReferenceMapping != nil {
		return true
	}

	return false
}

// SetPlannerReferenceMapping gets a reference to the given string and assigns it to the PlannerReferenceMapping field.
func (o *SayWhenConfigModel) SetPlannerReferenceMapping(v string) {
	o.PlannerReferenceMapping = &v
}

// GetMetaData returns the MetaData field value if set, zero value otherwise.
func (o *SayWhenConfigModel) GetMetaData() string {
	if o == nil || o.MetaData == nil {
		var ret string
		return ret
	}
	return *o.MetaData
}

// GetMetaDataOk returns a tuple with the MetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SayWhenConfigModel) GetMetaDataOk() (*string, bool) {
	if o == nil || o.MetaData == nil {
		return nil, false
	}
	return o.MetaData, true
}

// HasMetaData returns a boolean if a field has been set.
func (o *SayWhenConfigModel) HasMetaData() bool {
	if o != nil && o.MetaData != nil {
		return true
	}

	return false
}

// SetMetaData gets a reference to the given string and assigns it to the MetaData field.
func (o *SayWhenConfigModel) SetMetaData(v string) {
	o.MetaData = &v
}

// GetKeyRing returns the KeyRing field value if set, zero value otherwise.
func (o *SayWhenConfigModel) GetKeyRing() string {
	if o == nil || o.KeyRing == nil {
		var ret string
		return ret
	}
	return *o.KeyRing
}

// GetKeyRingOk returns a tuple with the KeyRing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SayWhenConfigModel) GetKeyRingOk() (*string, bool) {
	if o == nil || o.KeyRing == nil {
		return nil, false
	}
	return o.KeyRing, true
}

// HasKeyRing returns a boolean if a field has been set.
func (o *SayWhenConfigModel) HasKeyRing() bool {
	if o != nil && o.KeyRing != nil {
		return true
	}

	return false
}

// SetKeyRing gets a reference to the given string and assigns it to the KeyRing field.
func (o *SayWhenConfigModel) SetKeyRing(v string) {
	o.KeyRing = &v
}

// GetVisitTypeMap returns the VisitTypeMap field value if set, zero value otherwise.
func (o *SayWhenConfigModel) GetVisitTypeMap() string {
	if o == nil || o.VisitTypeMap == nil {
		var ret string
		return ret
	}
	return *o.VisitTypeMap
}

// GetVisitTypeMapOk returns a tuple with the VisitTypeMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SayWhenConfigModel) GetVisitTypeMapOk() (*string, bool) {
	if o == nil || o.VisitTypeMap == nil {
		return nil, false
	}
	return o.VisitTypeMap, true
}

// HasVisitTypeMap returns a boolean if a field has been set.
func (o *SayWhenConfigModel) HasVisitTypeMap() bool {
	if o != nil && o.VisitTypeMap != nil {
		return true
	}

	return false
}

// SetVisitTypeMap gets a reference to the given string and assigns it to the VisitTypeMap field.
func (o *SayWhenConfigModel) SetVisitTypeMap(v string) {
	o.VisitTypeMap = &v
}

func (o SayWhenConfigModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PlannerReferenceMapping != nil {
		toSerialize["plannerReferenceMapping"] = o.PlannerReferenceMapping
	}
	if o.MetaData != nil {
		toSerialize["metaData"] = o.MetaData
	}
	if o.KeyRing != nil {
		toSerialize["keyRing"] = o.KeyRing
	}
	if o.VisitTypeMap != nil {
		toSerialize["visitTypeMap"] = o.VisitTypeMap
	}
	return json.Marshal(toSerialize)
}

type NullableSayWhenConfigModel struct {
	value *SayWhenConfigModel
	isSet bool
}

func (v NullableSayWhenConfigModel) Get() *SayWhenConfigModel {
	return v.value
}

func (v *NullableSayWhenConfigModel) Set(val *SayWhenConfigModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSayWhenConfigModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSayWhenConfigModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSayWhenConfigModel(val *SayWhenConfigModel) *NullableSayWhenConfigModel {
	return &NullableSayWhenConfigModel{value: val, isSet: true}
}

func (v NullableSayWhenConfigModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSayWhenConfigModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

