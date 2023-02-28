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

// CapacityModel struct for CapacityModel
type CapacityModel struct {
	// Unique ID
	Id *int64 `json:"id,omitempty"`
	// id for capacity type
	CapacityTypeId *int32 `json:"capacity_type_id,omitempty"`
	// name of capacity type
	CapacityTypeName *string `json:"capacity_type_name,omitempty"`
	CapacityType *CapacityTypeModel `json:"capacity_type,omitempty"`
	// Capacity value
	CapacityValue *float64 `json:"capacity_value,omitempty"`
	// 
	UnitValues *[]UnitValueModel `json:"unit_values,omitempty"`
	// Name of used unit of measurement for values provided in capacity_value
	CapacityValueUomName *string `json:"capacity_value_uom_name,omitempty"`
	// Name of used unit of measurement for values provided in unit_values
	UnitValuesUomName *string `json:"unit_values_uom_name,omitempty"`
}

// NewCapacityModel instantiates a new CapacityModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapacityModel() *CapacityModel {
	this := CapacityModel{}
	return &this
}

// NewCapacityModelWithDefaults instantiates a new CapacityModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapacityModelWithDefaults() *CapacityModel {
	this := CapacityModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CapacityModel) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityModel) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CapacityModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *CapacityModel) SetId(v int64) {
	o.Id = &v
}

// GetCapacityTypeId returns the CapacityTypeId field value if set, zero value otherwise.
func (o *CapacityModel) GetCapacityTypeId() int32 {
	if o == nil || o.CapacityTypeId == nil {
		var ret int32
		return ret
	}
	return *o.CapacityTypeId
}

// GetCapacityTypeIdOk returns a tuple with the CapacityTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityModel) GetCapacityTypeIdOk() (*int32, bool) {
	if o == nil || o.CapacityTypeId == nil {
		return nil, false
	}
	return o.CapacityTypeId, true
}

// HasCapacityTypeId returns a boolean if a field has been set.
func (o *CapacityModel) HasCapacityTypeId() bool {
	if o != nil && o.CapacityTypeId != nil {
		return true
	}

	return false
}

// SetCapacityTypeId gets a reference to the given int32 and assigns it to the CapacityTypeId field.
func (o *CapacityModel) SetCapacityTypeId(v int32) {
	o.CapacityTypeId = &v
}

// GetCapacityTypeName returns the CapacityTypeName field value if set, zero value otherwise.
func (o *CapacityModel) GetCapacityTypeName() string {
	if o == nil || o.CapacityTypeName == nil {
		var ret string
		return ret
	}
	return *o.CapacityTypeName
}

// GetCapacityTypeNameOk returns a tuple with the CapacityTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityModel) GetCapacityTypeNameOk() (*string, bool) {
	if o == nil || o.CapacityTypeName == nil {
		return nil, false
	}
	return o.CapacityTypeName, true
}

// HasCapacityTypeName returns a boolean if a field has been set.
func (o *CapacityModel) HasCapacityTypeName() bool {
	if o != nil && o.CapacityTypeName != nil {
		return true
	}

	return false
}

// SetCapacityTypeName gets a reference to the given string and assigns it to the CapacityTypeName field.
func (o *CapacityModel) SetCapacityTypeName(v string) {
	o.CapacityTypeName = &v
}

// GetCapacityType returns the CapacityType field value if set, zero value otherwise.
func (o *CapacityModel) GetCapacityType() CapacityTypeModel {
	if o == nil || o.CapacityType == nil {
		var ret CapacityTypeModel
		return ret
	}
	return *o.CapacityType
}

// GetCapacityTypeOk returns a tuple with the CapacityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityModel) GetCapacityTypeOk() (*CapacityTypeModel, bool) {
	if o == nil || o.CapacityType == nil {
		return nil, false
	}
	return o.CapacityType, true
}

// HasCapacityType returns a boolean if a field has been set.
func (o *CapacityModel) HasCapacityType() bool {
	if o != nil && o.CapacityType != nil {
		return true
	}

	return false
}

// SetCapacityType gets a reference to the given CapacityTypeModel and assigns it to the CapacityType field.
func (o *CapacityModel) SetCapacityType(v CapacityTypeModel) {
	o.CapacityType = &v
}

// GetCapacityValue returns the CapacityValue field value if set, zero value otherwise.
func (o *CapacityModel) GetCapacityValue() float64 {
	if o == nil || o.CapacityValue == nil {
		var ret float64
		return ret
	}
	return *o.CapacityValue
}

// GetCapacityValueOk returns a tuple with the CapacityValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityModel) GetCapacityValueOk() (*float64, bool) {
	if o == nil || o.CapacityValue == nil {
		return nil, false
	}
	return o.CapacityValue, true
}

// HasCapacityValue returns a boolean if a field has been set.
func (o *CapacityModel) HasCapacityValue() bool {
	if o != nil && o.CapacityValue != nil {
		return true
	}

	return false
}

// SetCapacityValue gets a reference to the given float64 and assigns it to the CapacityValue field.
func (o *CapacityModel) SetCapacityValue(v float64) {
	o.CapacityValue = &v
}

// GetUnitValues returns the UnitValues field value if set, zero value otherwise.
func (o *CapacityModel) GetUnitValues() []UnitValueModel {
	if o == nil || o.UnitValues == nil {
		var ret []UnitValueModel
		return ret
	}
	return *o.UnitValues
}

// GetUnitValuesOk returns a tuple with the UnitValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityModel) GetUnitValuesOk() (*[]UnitValueModel, bool) {
	if o == nil || o.UnitValues == nil {
		return nil, false
	}
	return o.UnitValues, true
}

// HasUnitValues returns a boolean if a field has been set.
func (o *CapacityModel) HasUnitValues() bool {
	if o != nil && o.UnitValues != nil {
		return true
	}

	return false
}

// SetUnitValues gets a reference to the given []UnitValueModel and assigns it to the UnitValues field.
func (o *CapacityModel) SetUnitValues(v []UnitValueModel) {
	o.UnitValues = &v
}

// GetCapacityValueUomName returns the CapacityValueUomName field value if set, zero value otherwise.
func (o *CapacityModel) GetCapacityValueUomName() string {
	if o == nil || o.CapacityValueUomName == nil {
		var ret string
		return ret
	}
	return *o.CapacityValueUomName
}

// GetCapacityValueUomNameOk returns a tuple with the CapacityValueUomName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityModel) GetCapacityValueUomNameOk() (*string, bool) {
	if o == nil || o.CapacityValueUomName == nil {
		return nil, false
	}
	return o.CapacityValueUomName, true
}

// HasCapacityValueUomName returns a boolean if a field has been set.
func (o *CapacityModel) HasCapacityValueUomName() bool {
	if o != nil && o.CapacityValueUomName != nil {
		return true
	}

	return false
}

// SetCapacityValueUomName gets a reference to the given string and assigns it to the CapacityValueUomName field.
func (o *CapacityModel) SetCapacityValueUomName(v string) {
	o.CapacityValueUomName = &v
}

// GetUnitValuesUomName returns the UnitValuesUomName field value if set, zero value otherwise.
func (o *CapacityModel) GetUnitValuesUomName() string {
	if o == nil || o.UnitValuesUomName == nil {
		var ret string
		return ret
	}
	return *o.UnitValuesUomName
}

// GetUnitValuesUomNameOk returns a tuple with the UnitValuesUomName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CapacityModel) GetUnitValuesUomNameOk() (*string, bool) {
	if o == nil || o.UnitValuesUomName == nil {
		return nil, false
	}
	return o.UnitValuesUomName, true
}

// HasUnitValuesUomName returns a boolean if a field has been set.
func (o *CapacityModel) HasUnitValuesUomName() bool {
	if o != nil && o.UnitValuesUomName != nil {
		return true
	}

	return false
}

// SetUnitValuesUomName gets a reference to the given string and assigns it to the UnitValuesUomName field.
func (o *CapacityModel) SetUnitValuesUomName(v string) {
	o.UnitValuesUomName = &v
}

func (o CapacityModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CapacityTypeId != nil {
		toSerialize["capacity_type_id"] = o.CapacityTypeId
	}
	if o.CapacityTypeName != nil {
		toSerialize["capacity_type_name"] = o.CapacityTypeName
	}
	if o.CapacityType != nil {
		toSerialize["capacity_type"] = o.CapacityType
	}
	if o.CapacityValue != nil {
		toSerialize["capacity_value"] = o.CapacityValue
	}
	if o.UnitValues != nil {
		toSerialize["unit_values"] = o.UnitValues
	}
	if o.CapacityValueUomName != nil {
		toSerialize["capacity_value_uom_name"] = o.CapacityValueUomName
	}
	if o.UnitValuesUomName != nil {
		toSerialize["unit_values_uom_name"] = o.UnitValuesUomName
	}
	return json.Marshal(toSerialize)
}

type NullableCapacityModel struct {
	value *CapacityModel
	isSet bool
}

func (v NullableCapacityModel) Get() *CapacityModel {
	return v.value
}

func (v *NullableCapacityModel) Set(val *CapacityModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCapacityModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCapacityModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapacityModel(val *CapacityModel) *NullableCapacityModel {
	return &NullableCapacityModel{value: val, isSet: true}
}

func (v NullableCapacityModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapacityModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


