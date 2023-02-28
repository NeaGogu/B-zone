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

// EquipmentModel struct for EquipmentModel
type EquipmentModel struct {
	// 
	Id *int64 `json:"id,omitempty"`
	// Equipment Type ID
	EquipmentTypeId *int32 `json:"equipment_type_id,omitempty"`
	// Equipment Type Name
	EquipmentTypeName *string `json:"equipment_type_name,omitempty"`
	// 
	Info *string `json:"info,omitempty"`
	// 
	Name *string `json:"name,omitempty"`
	// 
	RegistrationNr *string `json:"registration_nr,omitempty"`
	// 
	MaxSpeed *int32 `json:"max_speed,omitempty"`
	// 
	Tags *[]TagModel `json:"tags,omitempty"`
	// 
	AppliedCapacities *map[string]interface{} `json:"applied_capacities,omitempty"`
	// 
	Capacities *[]CapacityModel `json:"capacities,omitempty"`
	// 
	Links *[]LinkModel `json:"links,omitempty"`
	// 
	MetaData *[]MetaDataModel `json:"meta_data,omitempty"`
}

// NewEquipmentModel instantiates a new EquipmentModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentModel() *EquipmentModel {
	this := EquipmentModel{}
	return &this
}

// NewEquipmentModelWithDefaults instantiates a new EquipmentModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentModelWithDefaults() *EquipmentModel {
	this := EquipmentModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EquipmentModel) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentModel) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EquipmentModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *EquipmentModel) SetId(v int64) {
	o.Id = &v
}

// GetEquipmentTypeId returns the EquipmentTypeId field value if set, zero value otherwise.
func (o *EquipmentModel) GetEquipmentTypeId() int32 {
	if o == nil || o.EquipmentTypeId == nil {
		var ret int32
		return ret
	}
	return *o.EquipmentTypeId
}

// GetEquipmentTypeIdOk returns a tuple with the EquipmentTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentModel) GetEquipmentTypeIdOk() (*int32, bool) {
	if o == nil || o.EquipmentTypeId == nil {
		return nil, false
	}
	return o.EquipmentTypeId, true
}

// HasEquipmentTypeId returns a boolean if a field has been set.
func (o *EquipmentModel) HasEquipmentTypeId() bool {
	if o != nil && o.EquipmentTypeId != nil {
		return true
	}

	return false
}

// SetEquipmentTypeId gets a reference to the given int32 and assigns it to the EquipmentTypeId field.
func (o *EquipmentModel) SetEquipmentTypeId(v int32) {
	o.EquipmentTypeId = &v
}

// GetEquipmentTypeName returns the EquipmentTypeName field value if set, zero value otherwise.
func (o *EquipmentModel) GetEquipmentTypeName() string {
	if o == nil || o.EquipmentTypeName == nil {
		var ret string
		return ret
	}
	return *o.EquipmentTypeName
}

// GetEquipmentTypeNameOk returns a tuple with the EquipmentTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentModel) GetEquipmentTypeNameOk() (*string, bool) {
	if o == nil || o.EquipmentTypeName == nil {
		return nil, false
	}
	return o.EquipmentTypeName, true
}

// HasEquipmentTypeName returns a boolean if a field has been set.
func (o *EquipmentModel) HasEquipmentTypeName() bool {
	if o != nil && o.EquipmentTypeName != nil {
		return true
	}

	return false
}

// SetEquipmentTypeName gets a reference to the given string and assigns it to the EquipmentTypeName field.
func (o *EquipmentModel) SetEquipmentTypeName(v string) {
	o.EquipmentTypeName = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *EquipmentModel) GetInfo() string {
	if o == nil || o.Info == nil {
		var ret string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentModel) GetInfoOk() (*string, bool) {
	if o == nil || o.Info == nil {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *EquipmentModel) HasInfo() bool {
	if o != nil && o.Info != nil {
		return true
	}

	return false
}

// SetInfo gets a reference to the given string and assigns it to the Info field.
func (o *EquipmentModel) SetInfo(v string) {
	o.Info = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EquipmentModel) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentModel) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EquipmentModel) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EquipmentModel) SetName(v string) {
	o.Name = &v
}

// GetRegistrationNr returns the RegistrationNr field value if set, zero value otherwise.
func (o *EquipmentModel) GetRegistrationNr() string {
	if o == nil || o.RegistrationNr == nil {
		var ret string
		return ret
	}
	return *o.RegistrationNr
}

// GetRegistrationNrOk returns a tuple with the RegistrationNr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentModel) GetRegistrationNrOk() (*string, bool) {
	if o == nil || o.RegistrationNr == nil {
		return nil, false
	}
	return o.RegistrationNr, true
}

// HasRegistrationNr returns a boolean if a field has been set.
func (o *EquipmentModel) HasRegistrationNr() bool {
	if o != nil && o.RegistrationNr != nil {
		return true
	}

	return false
}

// SetRegistrationNr gets a reference to the given string and assigns it to the RegistrationNr field.
func (o *EquipmentModel) SetRegistrationNr(v string) {
	o.RegistrationNr = &v
}

// GetMaxSpeed returns the MaxSpeed field value if set, zero value otherwise.
func (o *EquipmentModel) GetMaxSpeed() int32 {
	if o == nil || o.MaxSpeed == nil {
		var ret int32
		return ret
	}
	return *o.MaxSpeed
}

// GetMaxSpeedOk returns a tuple with the MaxSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentModel) GetMaxSpeedOk() (*int32, bool) {
	if o == nil || o.MaxSpeed == nil {
		return nil, false
	}
	return o.MaxSpeed, true
}

// HasMaxSpeed returns a boolean if a field has been set.
func (o *EquipmentModel) HasMaxSpeed() bool {
	if o != nil && o.MaxSpeed != nil {
		return true
	}

	return false
}

// SetMaxSpeed gets a reference to the given int32 and assigns it to the MaxSpeed field.
func (o *EquipmentModel) SetMaxSpeed(v int32) {
	o.MaxSpeed = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *EquipmentModel) GetTags() []TagModel {
	if o == nil || o.Tags == nil {
		var ret []TagModel
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentModel) GetTagsOk() (*[]TagModel, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *EquipmentModel) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagModel and assigns it to the Tags field.
func (o *EquipmentModel) SetTags(v []TagModel) {
	o.Tags = &v
}

// GetAppliedCapacities returns the AppliedCapacities field value if set, zero value otherwise.
func (o *EquipmentModel) GetAppliedCapacities() map[string]interface{} {
	if o == nil || o.AppliedCapacities == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.AppliedCapacities
}

// GetAppliedCapacitiesOk returns a tuple with the AppliedCapacities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentModel) GetAppliedCapacitiesOk() (*map[string]interface{}, bool) {
	if o == nil || o.AppliedCapacities == nil {
		return nil, false
	}
	return o.AppliedCapacities, true
}

// HasAppliedCapacities returns a boolean if a field has been set.
func (o *EquipmentModel) HasAppliedCapacities() bool {
	if o != nil && o.AppliedCapacities != nil {
		return true
	}

	return false
}

// SetAppliedCapacities gets a reference to the given map[string]interface{} and assigns it to the AppliedCapacities field.
func (o *EquipmentModel) SetAppliedCapacities(v map[string]interface{}) {
	o.AppliedCapacities = &v
}

// GetCapacities returns the Capacities field value if set, zero value otherwise.
func (o *EquipmentModel) GetCapacities() []CapacityModel {
	if o == nil || o.Capacities == nil {
		var ret []CapacityModel
		return ret
	}
	return *o.Capacities
}

// GetCapacitiesOk returns a tuple with the Capacities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentModel) GetCapacitiesOk() (*[]CapacityModel, bool) {
	if o == nil || o.Capacities == nil {
		return nil, false
	}
	return o.Capacities, true
}

// HasCapacities returns a boolean if a field has been set.
func (o *EquipmentModel) HasCapacities() bool {
	if o != nil && o.Capacities != nil {
		return true
	}

	return false
}

// SetCapacities gets a reference to the given []CapacityModel and assigns it to the Capacities field.
func (o *EquipmentModel) SetCapacities(v []CapacityModel) {
	o.Capacities = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EquipmentModel) GetLinks() []LinkModel {
	if o == nil || o.Links == nil {
		var ret []LinkModel
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentModel) GetLinksOk() (*[]LinkModel, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EquipmentModel) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []LinkModel and assigns it to the Links field.
func (o *EquipmentModel) SetLinks(v []LinkModel) {
	o.Links = &v
}

// GetMetaData returns the MetaData field value if set, zero value otherwise.
func (o *EquipmentModel) GetMetaData() []MetaDataModel {
	if o == nil || o.MetaData == nil {
		var ret []MetaDataModel
		return ret
	}
	return *o.MetaData
}

// GetMetaDataOk returns a tuple with the MetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentModel) GetMetaDataOk() (*[]MetaDataModel, bool) {
	if o == nil || o.MetaData == nil {
		return nil, false
	}
	return o.MetaData, true
}

// HasMetaData returns a boolean if a field has been set.
func (o *EquipmentModel) HasMetaData() bool {
	if o != nil && o.MetaData != nil {
		return true
	}

	return false
}

// SetMetaData gets a reference to the given []MetaDataModel and assigns it to the MetaData field.
func (o *EquipmentModel) SetMetaData(v []MetaDataModel) {
	o.MetaData = &v
}

func (o EquipmentModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.EquipmentTypeId != nil {
		toSerialize["equipment_type_id"] = o.EquipmentTypeId
	}
	if o.EquipmentTypeName != nil {
		toSerialize["equipment_type_name"] = o.EquipmentTypeName
	}
	if o.Info != nil {
		toSerialize["info"] = o.Info
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.RegistrationNr != nil {
		toSerialize["registration_nr"] = o.RegistrationNr
	}
	if o.MaxSpeed != nil {
		toSerialize["max_speed"] = o.MaxSpeed
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.AppliedCapacities != nil {
		toSerialize["applied_capacities"] = o.AppliedCapacities
	}
	if o.Capacities != nil {
		toSerialize["capacities"] = o.Capacities
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.MetaData != nil {
		toSerialize["meta_data"] = o.MetaData
	}
	return json.Marshal(toSerialize)
}

type NullableEquipmentModel struct {
	value *EquipmentModel
	isSet bool
}

func (v NullableEquipmentModel) Get() *EquipmentModel {
	return v.value
}

func (v *NullableEquipmentModel) Set(val *EquipmentModel) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentModel) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentModel(val *EquipmentModel) *NullableEquipmentModel {
	return &NullableEquipmentModel{value: val, isSet: true}
}

func (v NullableEquipmentModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


