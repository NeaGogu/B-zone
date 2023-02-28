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

// TrailerModel struct for TrailerModel
type TrailerModel struct {
	// Unique Identifier
	Id *int64 `json:"id,omitempty"`
	// name
	Name *string `json:"name,omitempty"`
	// registration_nr
	RegistrationNr *string `json:"registration_nr,omitempty"`
	// name
	Info *string `json:"info,omitempty"`
	// 
	AppliedCapacities *map[string]interface{} `json:"applied_capacities,omitempty"`
	// 
	Capacities *[]CapacityModel `json:"capacities,omitempty"`
	// 
	Compartments *[]CompartmentModel `json:"compartments,omitempty"`
	// 
	Tags *[]TagModel `json:"tags,omitempty"`
	// Tag names
	TagNames *[]string `json:"tag_names,omitempty"`
	// 
	Links *[]LinkModel `json:"links,omitempty"`
	// 
	MetaData *[]MetaDataModel `json:"meta_data,omitempty"`
	// 
	Files *[]FileModel `json:"files,omitempty"`
	// created_at date time
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// updated_at date time
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// created_at date time
	CreatedBy *time.Time `json:"created_by,omitempty"`
	// updated_at date time
	UpdatedBy *time.Time `json:"updated_by,omitempty"`
	CreatedByUser *UsersModel `json:"created_by_user,omitempty"`
	UpdatedByUser *UsersModel `json:"updated_by_user,omitempty"`
	// Trailer updated by user full name
	UpdatedByName *string `json:"updated_by_name,omitempty"`
}

// NewTrailerModel instantiates a new TrailerModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrailerModel() *TrailerModel {
	this := TrailerModel{}
	return &this
}

// NewTrailerModelWithDefaults instantiates a new TrailerModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrailerModelWithDefaults() *TrailerModel {
	this := TrailerModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TrailerModel) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrailerModel) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TrailerModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *TrailerModel) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TrailerModel) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrailerModel) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TrailerModel) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TrailerModel) SetName(v string) {
	o.Name = &v
}

// GetRegistrationNr returns the RegistrationNr field value if set, zero value otherwise.
func (o *TrailerModel) GetRegistrationNr() string {
	if o == nil || o.RegistrationNr == nil {
		var ret string
		return ret
	}
	return *o.RegistrationNr
}

// GetRegistrationNrOk returns a tuple with the RegistrationNr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrailerModel) GetRegistrationNrOk() (*string, bool) {
	if o == nil || o.RegistrationNr == nil {
		return nil, false
	}
	return o.RegistrationNr, true
}

// HasRegistrationNr returns a boolean if a field has been set.
func (o *TrailerModel) HasRegistrationNr() bool {
	if o != nil && o.RegistrationNr != nil {
		return true
	}

	return false
}

// SetRegistrationNr gets a reference to the given string and assigns it to the RegistrationNr field.
func (o *TrailerModel) SetRegistrationNr(v string) {
	o.RegistrationNr = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *TrailerModel) GetInfo() string {
	if o == nil || o.Info == nil {
		var ret string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrailerModel) GetInfoOk() (*string, bool) {
	if o == nil || o.Info == nil {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *TrailerModel) HasInfo() bool {
	if o != nil && o.Info != nil {
		return true
	}

	return false
}

// SetInfo gets a reference to the given string and assigns it to the Info field.
func (o *TrailerModel) SetInfo(v string) {
	o.Info = &v
}

// GetAppliedCapacities returns the AppliedCapacities field value if set, zero value otherwise.
func (o *TrailerModel) GetAppliedCapacities() map[string]interface{} {
	if o == nil || o.AppliedCapacities == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.AppliedCapacities
}

// GetAppliedCapacitiesOk returns a tuple with the AppliedCapacities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrailerModel) GetAppliedCapacitiesOk() (*map[string]interface{}, bool) {
	if o == nil || o.AppliedCapacities == nil {
		return nil, false
	}
	return o.AppliedCapacities, true
}

// HasAppliedCapacities returns a boolean if a field has been set.
func (o *TrailerModel) HasAppliedCapacities() bool {
	if o != nil && o.AppliedCapacities != nil {
		return true
	}

	return false
}

// SetAppliedCapacities gets a reference to the given map[string]interface{} and assigns it to the AppliedCapacities field.
func (o *TrailerModel) SetAppliedCapacities(v map[string]interface{}) {
	o.AppliedCapacities = &v
}

// GetCapacities returns the Capacities field value if set, zero value otherwise.
func (o *TrailerModel) GetCapacities() []CapacityModel {
	if o == nil || o.Capacities == nil {
		var ret []CapacityModel
		return ret
	}
	return *o.Capacities
}

// GetCapacitiesOk returns a tuple with the Capacities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrailerModel) GetCapacitiesOk() (*[]CapacityModel, bool) {
	if o == nil || o.Capacities == nil {
		return nil, false
	}
	return o.Capacities, true
}

// HasCapacities returns a boolean if a field has been set.
func (o *TrailerModel) HasCapacities() bool {
	if o != nil && o.Capacities != nil {
		return true
	}

	return false
}

// SetCapacities gets a reference to the given []CapacityModel and assigns it to the Capacities field.
func (o *TrailerModel) SetCapacities(v []CapacityModel) {
	o.Capacities = &v
}

// GetCompartments returns the Compartments field value if set, zero value otherwise.
func (o *TrailerModel) GetCompartments() []CompartmentModel {
	if o == nil || o.Compartments == nil {
		var ret []CompartmentModel
		return ret
	}
	return *o.Compartments
}

// GetCompartmentsOk returns a tuple with the Compartments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrailerModel) GetCompartmentsOk() (*[]CompartmentModel, bool) {
	if o == nil || o.Compartments == nil {
		return nil, false
	}
	return o.Compartments, true
}

// HasCompartments returns a boolean if a field has been set.
func (o *TrailerModel) HasCompartments() bool {
	if o != nil && o.Compartments != nil {
		return true
	}

	return false
}

// SetCompartments gets a reference to the given []CompartmentModel and assigns it to the Compartments field.
func (o *TrailerModel) SetCompartments(v []CompartmentModel) {
	o.Compartments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *TrailerModel) GetTags() []TagModel {
	if o == nil || o.Tags == nil {
		var ret []TagModel
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrailerModel) GetTagsOk() (*[]TagModel, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TrailerModel) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagModel and assigns it to the Tags field.
func (o *TrailerModel) SetTags(v []TagModel) {
	o.Tags = &v
}

// GetTagNames returns the TagNames field value if set, zero value otherwise.
func (o *TrailerModel) GetTagNames() []string {
	if o == nil || o.TagNames == nil {
		var ret []string
		return ret
	}
	return *o.TagNames
}

// GetTagNamesOk returns a tuple with the TagNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrailerModel) GetTagNamesOk() (*[]string, bool) {
	if o == nil || o.TagNames == nil {
		return nil, false
	}
	return o.TagNames, true
}

// HasTagNames returns a boolean if a field has been set.
func (o *TrailerModel) HasTagNames() bool {
	if o != nil && o.TagNames != nil {
		return true
	}

	return false
}

// SetTagNames gets a reference to the given []string and assigns it to the TagNames field.
func (o *TrailerModel) SetTagNames(v []string) {
	o.TagNames = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TrailerModel) GetLinks() []LinkModel {
	if o == nil || o.Links == nil {
		var ret []LinkModel
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrailerModel) GetLinksOk() (*[]LinkModel, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TrailerModel) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []LinkModel and assigns it to the Links field.
func (o *TrailerModel) SetLinks(v []LinkModel) {
	o.Links = &v
}

// GetMetaData returns the MetaData field value if set, zero value otherwise.
func (o *TrailerModel) GetMetaData() []MetaDataModel {
	if o == nil || o.MetaData == nil {
		var ret []MetaDataModel
		return ret
	}
	return *o.MetaData
}

// GetMetaDataOk returns a tuple with the MetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrailerModel) GetMetaDataOk() (*[]MetaDataModel, bool) {
	if o == nil || o.MetaData == nil {
		return nil, false
	}
	return o.MetaData, true
}

// HasMetaData returns a boolean if a field has been set.
func (o *TrailerModel) HasMetaData() bool {
	if o != nil && o.MetaData != nil {
		return true
	}

	return false
}

// SetMetaData gets a reference to the given []MetaDataModel and assigns it to the MetaData field.
func (o *TrailerModel) SetMetaData(v []MetaDataModel) {
	o.MetaData = &v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *TrailerModel) GetFiles() []FileModel {
	if o == nil || o.Files == nil {
		var ret []FileModel
		return ret
	}
	return *o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrailerModel) GetFilesOk() (*[]FileModel, bool) {
	if o == nil || o.Files == nil {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *TrailerModel) HasFiles() bool {
	if o != nil && o.Files != nil {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []FileModel and assigns it to the Files field.
func (o *TrailerModel) SetFiles(v []FileModel) {
	o.Files = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TrailerModel) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrailerModel) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TrailerModel) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *TrailerModel) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TrailerModel) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrailerModel) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TrailerModel) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *TrailerModel) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *TrailerModel) GetCreatedBy() time.Time {
	if o == nil || o.CreatedBy == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrailerModel) GetCreatedByOk() (*time.Time, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *TrailerModel) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given time.Time and assigns it to the CreatedBy field.
func (o *TrailerModel) SetCreatedBy(v time.Time) {
	o.CreatedBy = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *TrailerModel) GetUpdatedBy() time.Time {
	if o == nil || o.UpdatedBy == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrailerModel) GetUpdatedByOk() (*time.Time, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *TrailerModel) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given time.Time and assigns it to the UpdatedBy field.
func (o *TrailerModel) SetUpdatedBy(v time.Time) {
	o.UpdatedBy = &v
}

// GetCreatedByUser returns the CreatedByUser field value if set, zero value otherwise.
func (o *TrailerModel) GetCreatedByUser() UsersModel {
	if o == nil || o.CreatedByUser == nil {
		var ret UsersModel
		return ret
	}
	return *o.CreatedByUser
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrailerModel) GetCreatedByUserOk() (*UsersModel, bool) {
	if o == nil || o.CreatedByUser == nil {
		return nil, false
	}
	return o.CreatedByUser, true
}

// HasCreatedByUser returns a boolean if a field has been set.
func (o *TrailerModel) HasCreatedByUser() bool {
	if o != nil && o.CreatedByUser != nil {
		return true
	}

	return false
}

// SetCreatedByUser gets a reference to the given UsersModel and assigns it to the CreatedByUser field.
func (o *TrailerModel) SetCreatedByUser(v UsersModel) {
	o.CreatedByUser = &v
}

// GetUpdatedByUser returns the UpdatedByUser field value if set, zero value otherwise.
func (o *TrailerModel) GetUpdatedByUser() UsersModel {
	if o == nil || o.UpdatedByUser == nil {
		var ret UsersModel
		return ret
	}
	return *o.UpdatedByUser
}

// GetUpdatedByUserOk returns a tuple with the UpdatedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrailerModel) GetUpdatedByUserOk() (*UsersModel, bool) {
	if o == nil || o.UpdatedByUser == nil {
		return nil, false
	}
	return o.UpdatedByUser, true
}

// HasUpdatedByUser returns a boolean if a field has been set.
func (o *TrailerModel) HasUpdatedByUser() bool {
	if o != nil && o.UpdatedByUser != nil {
		return true
	}

	return false
}

// SetUpdatedByUser gets a reference to the given UsersModel and assigns it to the UpdatedByUser field.
func (o *TrailerModel) SetUpdatedByUser(v UsersModel) {
	o.UpdatedByUser = &v
}

// GetUpdatedByName returns the UpdatedByName field value if set, zero value otherwise.
func (o *TrailerModel) GetUpdatedByName() string {
	if o == nil || o.UpdatedByName == nil {
		var ret string
		return ret
	}
	return *o.UpdatedByName
}

// GetUpdatedByNameOk returns a tuple with the UpdatedByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrailerModel) GetUpdatedByNameOk() (*string, bool) {
	if o == nil || o.UpdatedByName == nil {
		return nil, false
	}
	return o.UpdatedByName, true
}

// HasUpdatedByName returns a boolean if a field has been set.
func (o *TrailerModel) HasUpdatedByName() bool {
	if o != nil && o.UpdatedByName != nil {
		return true
	}

	return false
}

// SetUpdatedByName gets a reference to the given string and assigns it to the UpdatedByName field.
func (o *TrailerModel) SetUpdatedByName(v string) {
	o.UpdatedByName = &v
}

func (o TrailerModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.RegistrationNr != nil {
		toSerialize["registration_nr"] = o.RegistrationNr
	}
	if o.Info != nil {
		toSerialize["info"] = o.Info
	}
	if o.AppliedCapacities != nil {
		toSerialize["applied_capacities"] = o.AppliedCapacities
	}
	if o.Capacities != nil {
		toSerialize["capacities"] = o.Capacities
	}
	if o.Compartments != nil {
		toSerialize["compartments"] = o.Compartments
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TagNames != nil {
		toSerialize["tag_names"] = o.TagNames
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.MetaData != nil {
		toSerialize["meta_data"] = o.MetaData
	}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.UpdatedBy != nil {
		toSerialize["updated_by"] = o.UpdatedBy
	}
	if o.CreatedByUser != nil {
		toSerialize["created_by_user"] = o.CreatedByUser
	}
	if o.UpdatedByUser != nil {
		toSerialize["updated_by_user"] = o.UpdatedByUser
	}
	if o.UpdatedByName != nil {
		toSerialize["updated_by_name"] = o.UpdatedByName
	}
	return json.Marshal(toSerialize)
}

type NullableTrailerModel struct {
	value *TrailerModel
	isSet bool
}

func (v NullableTrailerModel) Get() *TrailerModel {
	return v.value
}

func (v *NullableTrailerModel) Set(val *TrailerModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTrailerModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTrailerModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrailerModel(val *TrailerModel) *NullableTrailerModel {
	return &NullableTrailerModel{value: val, isSet: true}
}

func (v NullableTrailerModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrailerModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


