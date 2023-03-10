package models

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
)

type ZoneRangeDBModel struct {
	DB *mongo.Database
}

// ZoneRangeModel struct for ZoneRangeModel
type ZoneRangeModel struct {
	// Unique Zone type ID
	Id *int64 `json:"zone_range_id,omitempty"`
	// Zipcode range start
	ZipcodeFrom *int64 `json:"zipcode_from,omitempty"`
	// Zipcode range end
	ZipcodeTo *int64 `json:"zipcode_to,omitempty"`
	//
	IsoCountry *string `json:"iso_country,omitempty"`
}

// NewZoneRangeModel instantiates a new ZoneRangeModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneRangeModel() *ZoneRangeModel {
	this := ZoneRangeModel{}
	return &this
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneRangeModel) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ZoneRangeModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ZoneRangeModel) SetId(v int64) {
	o.Id = &v
}

// GetZipcodeFromOk returns a tuple with the ZipcodeFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneRangeModel) GetZipcodeFromOk() (*int64, bool) {
	if o == nil || o.ZipcodeFrom == nil {
		return nil, false
	}
	return o.ZipcodeFrom, true
}

// HasZipcodeFrom returns a boolean if a field has been set.
func (o *ZoneRangeModel) HasZipcodeFrom() bool {
	if o != nil && o.ZipcodeFrom != nil {
		return true
	}

	return false
}

// SetZipcodeFrom gets a reference to the given int64 and assigns it to the ZipcodeFrom field.
func (o *ZoneRangeModel) SetZipcodeFrom(v int64) {
	o.ZipcodeFrom = &v
}

// GetZipcodeToOk returns a tuple with the ZipcodeTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneRangeModel) GetZipcodeToOk() (*int64, bool) {
	if o == nil || o.ZipcodeTo == nil {
		return nil, false
	}
	return o.ZipcodeTo, true
}

// HasZipcodeTo returns a boolean if a field has been set.
func (o *ZoneRangeModel) HasZipcodeTo() bool {
	if o != nil && o.ZipcodeTo != nil {
		return true
	}

	return false
}

// SetZipcodeTo gets a reference to the given int64 and assigns it to the ZipcodeTo field.
func (o *ZoneRangeModel) SetZipcodeTo(v int64) {
	o.ZipcodeTo = &v
}

// GetIsoCountryOk returns a tuple with the IsoCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneRangeModel) GetIsoCountryOk() (*string, bool) {
	if o == nil || o.IsoCountry == nil {
		return nil, false
	}
	return o.IsoCountry, true
}

// HasIsoCountry returns a boolean if a field has been set.
func (o *ZoneRangeModel) HasIsoCountry() bool {
	if o != nil && o.IsoCountry != nil {
		return true
	}

	return false
}

// SetIsoCountry gets a reference to the given string and assigns it to the IsoCountry field.
func (o *ZoneRangeModel) SetIsoCountry(v string) {
	o.IsoCountry = &v
}

func (o ZoneRangeModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ZipcodeFrom != nil {
		toSerialize["zipcode_from"] = o.ZipcodeFrom
	}
	if o.ZipcodeTo != nil {
		toSerialize["zipcode_to"] = o.ZipcodeTo
	}
	if o.IsoCountry != nil {
		toSerialize["iso_country"] = o.IsoCountry
	}
	return json.Marshal(toSerialize)
}

type NullableZoneRangeModel struct {
	value *ZoneRangeModel
	isSet bool
}

func (v *NullableZoneRangeModel) Set(val *ZoneRangeModel) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneRangeModel) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneRangeModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneRangeModel(val *ZoneRangeModel) *NullableZoneRangeModel {
	return &NullableZoneRangeModel{value: val, isSet: true}
}

func (v NullableZoneRangeModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneRangeModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
