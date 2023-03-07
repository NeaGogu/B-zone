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
	"time"
)

// CompartmentModel struct for CompartmentModel
type CompartmentModel struct {
	// Unique Identifier
	Id *int64 `json:"id,omitempty"`
	// Nr of compartment
	Nr *string `json:"nr,omitempty"`
	// 
	TrailerId *int32 `json:"trailer_id,omitempty"`
	// 
	FilledCapacities *[]FilledCapacityModel `json:"filled_capacities,omitempty"`
	// 
	TrailerLink *[]LinkModel `json:"trailer_link,omitempty"`
	// 
	Links *[]LinkModel `json:"links,omitempty"`
	// 
	MetaData *[]MetaDataModel `json:"meta_data,omitempty"`
	CompartmentCreatedBy *UsersModel `json:"compartment_created_by,omitempty"`
	CompartmentUpdatedBy *UsersModel `json:"compartment_updated_by,omitempty"`
	// created_at date time
	CompartmentCreatedAt *time.Time `json:"compartment_created_at,omitempty"`
	// updated_at date time
	CompartmentUpdatedAt *time.Time `json:"compartment_updated_at,omitempty"`
}

// NewCompartmentModel instantiates a new CompartmentModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompartmentModel() *CompartmentModel {
	this := CompartmentModel{}
	return &this
}

// NewCompartmentModelWithDefaults instantiates a new CompartmentModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompartmentModelWithDefaults() *CompartmentModel {
	this := CompartmentModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CompartmentModel) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompartmentModel) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CompartmentModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *CompartmentModel) SetId(v int64) {
	o.Id = &v
}

// GetNr returns the Nr field value if set, zero value otherwise.
func (o *CompartmentModel) GetNr() string {
	if o == nil || o.Nr == nil {
		var ret string
		return ret
	}
	return *o.Nr
}

// GetNrOk returns a tuple with the Nr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompartmentModel) GetNrOk() (*string, bool) {
	if o == nil || o.Nr == nil {
		return nil, false
	}
	return o.Nr, true
}

// HasNr returns a boolean if a field has been set.
func (o *CompartmentModel) HasNr() bool {
	if o != nil && o.Nr != nil {
		return true
	}

	return false
}

// SetNr gets a reference to the given string and assigns it to the Nr field.
func (o *CompartmentModel) SetNr(v string) {
	o.Nr = &v
}

// GetTrailerId returns the TrailerId field value if set, zero value otherwise.
func (o *CompartmentModel) GetTrailerId() int32 {
	if o == nil || o.TrailerId == nil {
		var ret int32
		return ret
	}
	return *o.TrailerId
}

// GetTrailerIdOk returns a tuple with the TrailerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompartmentModel) GetTrailerIdOk() (*int32, bool) {
	if o == nil || o.TrailerId == nil {
		return nil, false
	}
	return o.TrailerId, true
}

// HasTrailerId returns a boolean if a field has been set.
func (o *CompartmentModel) HasTrailerId() bool {
	if o != nil && o.TrailerId != nil {
		return true
	}

	return false
}

// SetTrailerId gets a reference to the given int32 and assigns it to the TrailerId field.
func (o *CompartmentModel) SetTrailerId(v int32) {
	o.TrailerId = &v
}

// GetFilledCapacities returns the FilledCapacities field value if set, zero value otherwise.
func (o *CompartmentModel) GetFilledCapacities() []FilledCapacityModel {
	if o == nil || o.FilledCapacities == nil {
		var ret []FilledCapacityModel
		return ret
	}
	return *o.FilledCapacities
}

// GetFilledCapacitiesOk returns a tuple with the FilledCapacities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompartmentModel) GetFilledCapacitiesOk() (*[]FilledCapacityModel, bool) {
	if o == nil || o.FilledCapacities == nil {
		return nil, false
	}
	return o.FilledCapacities, true
}

// HasFilledCapacities returns a boolean if a field has been set.
func (o *CompartmentModel) HasFilledCapacities() bool {
	if o != nil && o.FilledCapacities != nil {
		return true
	}

	return false
}

// SetFilledCapacities gets a reference to the given []FilledCapacityModel and assigns it to the FilledCapacities field.
func (o *CompartmentModel) SetFilledCapacities(v []FilledCapacityModel) {
	o.FilledCapacities = &v
}

// GetTrailerLink returns the TrailerLink field value if set, zero value otherwise.
func (o *CompartmentModel) GetTrailerLink() []LinkModel {
	if o == nil || o.TrailerLink == nil {
		var ret []LinkModel
		return ret
	}
	return *o.TrailerLink
}

// GetTrailerLinkOk returns a tuple with the TrailerLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompartmentModel) GetTrailerLinkOk() (*[]LinkModel, bool) {
	if o == nil || o.TrailerLink == nil {
		return nil, false
	}
	return o.TrailerLink, true
}

// HasTrailerLink returns a boolean if a field has been set.
func (o *CompartmentModel) HasTrailerLink() bool {
	if o != nil && o.TrailerLink != nil {
		return true
	}

	return false
}

// SetTrailerLink gets a reference to the given []LinkModel and assigns it to the TrailerLink field.
func (o *CompartmentModel) SetTrailerLink(v []LinkModel) {
	o.TrailerLink = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CompartmentModel) GetLinks() []LinkModel {
	if o == nil || o.Links == nil {
		var ret []LinkModel
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompartmentModel) GetLinksOk() (*[]LinkModel, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CompartmentModel) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []LinkModel and assigns it to the Links field.
func (o *CompartmentModel) SetLinks(v []LinkModel) {
	o.Links = &v
}

// GetMetaData returns the MetaData field value if set, zero value otherwise.
func (o *CompartmentModel) GetMetaData() []MetaDataModel {
	if o == nil || o.MetaData == nil {
		var ret []MetaDataModel
		return ret
	}
	return *o.MetaData
}

// GetMetaDataOk returns a tuple with the MetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompartmentModel) GetMetaDataOk() (*[]MetaDataModel, bool) {
	if o == nil || o.MetaData == nil {
		return nil, false
	}
	return o.MetaData, true
}

// HasMetaData returns a boolean if a field has been set.
func (o *CompartmentModel) HasMetaData() bool {
	if o != nil && o.MetaData != nil {
		return true
	}

	return false
}

// SetMetaData gets a reference to the given []MetaDataModel and assigns it to the MetaData field.
func (o *CompartmentModel) SetMetaData(v []MetaDataModel) {
	o.MetaData = &v
}

// GetCompartmentCreatedBy returns the CompartmentCreatedBy field value if set, zero value otherwise.
func (o *CompartmentModel) GetCompartmentCreatedBy() UsersModel {
	if o == nil || o.CompartmentCreatedBy == nil {
		var ret UsersModel
		return ret
	}
	return *o.CompartmentCreatedBy
}

// GetCompartmentCreatedByOk returns a tuple with the CompartmentCreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompartmentModel) GetCompartmentCreatedByOk() (*UsersModel, bool) {
	if o == nil || o.CompartmentCreatedBy == nil {
		return nil, false
	}
	return o.CompartmentCreatedBy, true
}

// HasCompartmentCreatedBy returns a boolean if a field has been set.
func (o *CompartmentModel) HasCompartmentCreatedBy() bool {
	if o != nil && o.CompartmentCreatedBy != nil {
		return true
	}

	return false
}

// SetCompartmentCreatedBy gets a reference to the given UsersModel and assigns it to the CompartmentCreatedBy field.
func (o *CompartmentModel) SetCompartmentCreatedBy(v UsersModel) {
	o.CompartmentCreatedBy = &v
}

// GetCompartmentUpdatedBy returns the CompartmentUpdatedBy field value if set, zero value otherwise.
func (o *CompartmentModel) GetCompartmentUpdatedBy() UsersModel {
	if o == nil || o.CompartmentUpdatedBy == nil {
		var ret UsersModel
		return ret
	}
	return *o.CompartmentUpdatedBy
}

// GetCompartmentUpdatedByOk returns a tuple with the CompartmentUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompartmentModel) GetCompartmentUpdatedByOk() (*UsersModel, bool) {
	if o == nil || o.CompartmentUpdatedBy == nil {
		return nil, false
	}
	return o.CompartmentUpdatedBy, true
}

// HasCompartmentUpdatedBy returns a boolean if a field has been set.
func (o *CompartmentModel) HasCompartmentUpdatedBy() bool {
	if o != nil && o.CompartmentUpdatedBy != nil {
		return true
	}

	return false
}

// SetCompartmentUpdatedBy gets a reference to the given UsersModel and assigns it to the CompartmentUpdatedBy field.
func (o *CompartmentModel) SetCompartmentUpdatedBy(v UsersModel) {
	o.CompartmentUpdatedBy = &v
}

// GetCompartmentCreatedAt returns the CompartmentCreatedAt field value if set, zero value otherwise.
func (o *CompartmentModel) GetCompartmentCreatedAt() time.Time {
	if o == nil || o.CompartmentCreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CompartmentCreatedAt
}

// GetCompartmentCreatedAtOk returns a tuple with the CompartmentCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompartmentModel) GetCompartmentCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CompartmentCreatedAt == nil {
		return nil, false
	}
	return o.CompartmentCreatedAt, true
}

// HasCompartmentCreatedAt returns a boolean if a field has been set.
func (o *CompartmentModel) HasCompartmentCreatedAt() bool {
	if o != nil && o.CompartmentCreatedAt != nil {
		return true
	}

	return false
}

// SetCompartmentCreatedAt gets a reference to the given time.Time and assigns it to the CompartmentCreatedAt field.
func (o *CompartmentModel) SetCompartmentCreatedAt(v time.Time) {
	o.CompartmentCreatedAt = &v
}

// GetCompartmentUpdatedAt returns the CompartmentUpdatedAt field value if set, zero value otherwise.
func (o *CompartmentModel) GetCompartmentUpdatedAt() time.Time {
	if o == nil || o.CompartmentUpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CompartmentUpdatedAt
}

// GetCompartmentUpdatedAtOk returns a tuple with the CompartmentUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompartmentModel) GetCompartmentUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.CompartmentUpdatedAt == nil {
		return nil, false
	}
	return o.CompartmentUpdatedAt, true
}

// HasCompartmentUpdatedAt returns a boolean if a field has been set.
func (o *CompartmentModel) HasCompartmentUpdatedAt() bool {
	if o != nil && o.CompartmentUpdatedAt != nil {
		return true
	}

	return false
}

// SetCompartmentUpdatedAt gets a reference to the given time.Time and assigns it to the CompartmentUpdatedAt field.
func (o *CompartmentModel) SetCompartmentUpdatedAt(v time.Time) {
	o.CompartmentUpdatedAt = &v
}

func (o CompartmentModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Nr != nil {
		toSerialize["nr"] = o.Nr
	}
	if o.TrailerId != nil {
		toSerialize["trailer_id"] = o.TrailerId
	}
	if o.FilledCapacities != nil {
		toSerialize["filled_capacities"] = o.FilledCapacities
	}
	if o.TrailerLink != nil {
		toSerialize["trailer_link"] = o.TrailerLink
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.MetaData != nil {
		toSerialize["meta_data"] = o.MetaData
	}
	if o.CompartmentCreatedBy != nil {
		toSerialize["compartment_created_by"] = o.CompartmentCreatedBy
	}
	if o.CompartmentUpdatedBy != nil {
		toSerialize["compartment_updated_by"] = o.CompartmentUpdatedBy
	}
	if o.CompartmentCreatedAt != nil {
		toSerialize["compartment_created_at"] = o.CompartmentCreatedAt
	}
	if o.CompartmentUpdatedAt != nil {
		toSerialize["compartment_updated_at"] = o.CompartmentUpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableCompartmentModel struct {
	value *CompartmentModel
	isSet bool
}

func (v NullableCompartmentModel) Get() *CompartmentModel {
	return v.value
}

func (v *NullableCompartmentModel) Set(val *CompartmentModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCompartmentModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCompartmentModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompartmentModel(val *CompartmentModel) *NullableCompartmentModel {
	return &NullableCompartmentModel{value: val, isSet: true}
}

func (v NullableCompartmentModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompartmentModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

