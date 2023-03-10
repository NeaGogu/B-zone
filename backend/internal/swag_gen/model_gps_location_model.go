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

// GPSLocationModel struct for GPSLocationModel
type GPSLocationModel struct {
	// 
	Longitude *float32 `json:"longitude,omitempty"`
	// 
	Altitude *float32 `json:"altitude,omitempty"`
	// 
	Latitude *float32 `json:"latitude,omitempty"`
	// 
	Heading *int32 `json:"heading,omitempty"`
	// 
	Timestamp *int32 `json:"timestamp,omitempty"`
}

// NewGPSLocationModel instantiates a new GPSLocationModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGPSLocationModel() *GPSLocationModel {
	this := GPSLocationModel{}
	return &this
}

// NewGPSLocationModelWithDefaults instantiates a new GPSLocationModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGPSLocationModelWithDefaults() *GPSLocationModel {
	this := GPSLocationModel{}
	return &this
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *GPSLocationModel) GetLongitude() float32 {
	if o == nil || o.Longitude == nil {
		var ret float32
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GPSLocationModel) GetLongitudeOk() (*float32, bool) {
	if o == nil || o.Longitude == nil {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *GPSLocationModel) HasLongitude() bool {
	if o != nil && o.Longitude != nil {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given float32 and assigns it to the Longitude field.
func (o *GPSLocationModel) SetLongitude(v float32) {
	o.Longitude = &v
}

// GetAltitude returns the Altitude field value if set, zero value otherwise.
func (o *GPSLocationModel) GetAltitude() float32 {
	if o == nil || o.Altitude == nil {
		var ret float32
		return ret
	}
	return *o.Altitude
}

// GetAltitudeOk returns a tuple with the Altitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GPSLocationModel) GetAltitudeOk() (*float32, bool) {
	if o == nil || o.Altitude == nil {
		return nil, false
	}
	return o.Altitude, true
}

// HasAltitude returns a boolean if a field has been set.
func (o *GPSLocationModel) HasAltitude() bool {
	if o != nil && o.Altitude != nil {
		return true
	}

	return false
}

// SetAltitude gets a reference to the given float32 and assigns it to the Altitude field.
func (o *GPSLocationModel) SetAltitude(v float32) {
	o.Altitude = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *GPSLocationModel) GetLatitude() float32 {
	if o == nil || o.Latitude == nil {
		var ret float32
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GPSLocationModel) GetLatitudeOk() (*float32, bool) {
	if o == nil || o.Latitude == nil {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *GPSLocationModel) HasLatitude() bool {
	if o != nil && o.Latitude != nil {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given float32 and assigns it to the Latitude field.
func (o *GPSLocationModel) SetLatitude(v float32) {
	o.Latitude = &v
}

// GetHeading returns the Heading field value if set, zero value otherwise.
func (o *GPSLocationModel) GetHeading() int32 {
	if o == nil || o.Heading == nil {
		var ret int32
		return ret
	}
	return *o.Heading
}

// GetHeadingOk returns a tuple with the Heading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GPSLocationModel) GetHeadingOk() (*int32, bool) {
	if o == nil || o.Heading == nil {
		return nil, false
	}
	return o.Heading, true
}

// HasHeading returns a boolean if a field has been set.
func (o *GPSLocationModel) HasHeading() bool {
	if o != nil && o.Heading != nil {
		return true
	}

	return false
}

// SetHeading gets a reference to the given int32 and assigns it to the Heading field.
func (o *GPSLocationModel) SetHeading(v int32) {
	o.Heading = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *GPSLocationModel) GetTimestamp() int32 {
	if o == nil || o.Timestamp == nil {
		var ret int32
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GPSLocationModel) GetTimestampOk() (*int32, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *GPSLocationModel) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int32 and assigns it to the Timestamp field.
func (o *GPSLocationModel) SetTimestamp(v int32) {
	o.Timestamp = &v
}

func (o GPSLocationModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Longitude != nil {
		toSerialize["longitude"] = o.Longitude
	}
	if o.Altitude != nil {
		toSerialize["altitude"] = o.Altitude
	}
	if o.Latitude != nil {
		toSerialize["latitude"] = o.Latitude
	}
	if o.Heading != nil {
		toSerialize["heading"] = o.Heading
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

type NullableGPSLocationModel struct {
	value *GPSLocationModel
	isSet bool
}

func (v NullableGPSLocationModel) Get() *GPSLocationModel {
	return v.value
}

func (v *NullableGPSLocationModel) Set(val *GPSLocationModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGPSLocationModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGPSLocationModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGPSLocationModel(val *GPSLocationModel) *NullableGPSLocationModel {
	return &NullableGPSLocationModel{value: val, isSet: true}
}

func (v NullableGPSLocationModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGPSLocationModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

