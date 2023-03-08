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

// AddressModel struct for AddressModel
type AddressModel struct {
	// 
	Id *int64 `json:"id,omitempty"`
	// 
	AddressId *int64 `json:"address_id,omitempty"`
	// 
	PartyId *int32 `json:"party_id,omitempty"`
	// 
	PartyName *string `json:"party_name,omitempty"`
	// 
	Code *string `json:"code,omitempty"`
	// 
	Summary *string `json:"summary,omitempty"`
	// 
	FullName *string `json:"full_name,omitempty"`
	// 
	Name1 *string `json:"name_1,omitempty"`
	// 
	Name2 *string `json:"name_2,omitempty"`
	// 
	Street1 *string `json:"street_1,omitempty"`
	// 
	Street2 *string `json:"street_2,omitempty"`
	// 
	FullAddressline *string `json:"full_addressline,omitempty"`
	// 
	HouseNr *string `json:"house_nr,omitempty"`
	// 
	HouseNrAddendum *string `json:"house_nr_addendum,omitempty"`
	// 
	Zipcode *string `json:"zipcode,omitempty"`
	// 
	City *string `json:"city,omitempty"`
	// 
	State *string `json:"state,omitempty"`
	// 
	IsoCountry *string `json:"iso_country,omitempty"`
	// 
	CountryName *string `json:"country_name,omitempty"`
	// 
	TimeFrom *string `json:"time_from,omitempty"`
	// 
	TimeTo *string `json:"time_to,omitempty"`
	// 
	OpeningHours *[]OpeningHoursRuleModel `json:"opening_hours,omitempty"`
	// Default duration for activities on this address in minutes
	Duration *int32 `json:"duration,omitempty"`
	// Address Type names
	AddressTypeNames *[]string `json:"address_type_names,omitempty"`
	// 
	Emails *[]EmailModel `json:"emails,omitempty"`
	// 
	PhoneNrs *[]PhoneNrModel `json:"phone_nrs,omitempty"`
	// 
	Latitude *string `json:"latitude,omitempty"`
	// 
	Longitude *string `json:"longitude,omitempty"`
	// 
	ContactPerson *string `json:"contact_person,omitempty"`
	// 
	UserId *int32 `json:"user_id,omitempty"`
	// 
	Links *[]LinkModel `json:"links,omitempty"`
	// 
	MetaData *[]MetaDataModel `json:"meta_data,omitempty"`
	// 
	Notes *[]NoteModel `json:"notes,omitempty"`
	// 
	Files *[]FileModel `json:"files,omitempty"`
	// Tag names
	TagNames *[]string `json:"tag_names,omitempty"`
	// 
	Tags *[]TagModel `json:"tags,omitempty"`
}

// NewAddressModel instantiates a new AddressModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressModel() *AddressModel {
	this := AddressModel{}
	return &this
}

// NewAddressModelWithDefaults instantiates a new AddressModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressModelWithDefaults() *AddressModel {
	this := AddressModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AddressModel) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AddressModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *AddressModel) SetId(v int64) {
	o.Id = &v
}

// GetAddressId returns the AddressId field value if set, zero value otherwise.
func (o *AddressModel) GetAddressId() int64 {
	if o == nil || o.AddressId == nil {
		var ret int64
		return ret
	}
	return *o.AddressId
}

// GetAddressIdOk returns a tuple with the AddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetAddressIdOk() (*int64, bool) {
	if o == nil || o.AddressId == nil {
		return nil, false
	}
	return o.AddressId, true
}

// HasAddressId returns a boolean if a field has been set.
func (o *AddressModel) HasAddressId() bool {
	if o != nil && o.AddressId != nil {
		return true
	}

	return false
}

// SetAddressId gets a reference to the given int64 and assigns it to the AddressId field.
func (o *AddressModel) SetAddressId(v int64) {
	o.AddressId = &v
}

// GetPartyId returns the PartyId field value if set, zero value otherwise.
func (o *AddressModel) GetPartyId() int32 {
	if o == nil || o.PartyId == nil {
		var ret int32
		return ret
	}
	return *o.PartyId
}

// GetPartyIdOk returns a tuple with the PartyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetPartyIdOk() (*int32, bool) {
	if o == nil || o.PartyId == nil {
		return nil, false
	}
	return o.PartyId, true
}

// HasPartyId returns a boolean if a field has been set.
func (o *AddressModel) HasPartyId() bool {
	if o != nil && o.PartyId != nil {
		return true
	}

	return false
}

// SetPartyId gets a reference to the given int32 and assigns it to the PartyId field.
func (o *AddressModel) SetPartyId(v int32) {
	o.PartyId = &v
}

// GetPartyName returns the PartyName field value if set, zero value otherwise.
func (o *AddressModel) GetPartyName() string {
	if o == nil || o.PartyName == nil {
		var ret string
		return ret
	}
	return *o.PartyName
}

// GetPartyNameOk returns a tuple with the PartyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetPartyNameOk() (*string, bool) {
	if o == nil || o.PartyName == nil {
		return nil, false
	}
	return o.PartyName, true
}

// HasPartyName returns a boolean if a field has been set.
func (o *AddressModel) HasPartyName() bool {
	if o != nil && o.PartyName != nil {
		return true
	}

	return false
}

// SetPartyName gets a reference to the given string and assigns it to the PartyName field.
func (o *AddressModel) SetPartyName(v string) {
	o.PartyName = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AddressModel) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AddressModel) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *AddressModel) SetCode(v string) {
	o.Code = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *AddressModel) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *AddressModel) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *AddressModel) SetSummary(v string) {
	o.Summary = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *AddressModel) GetFullName() string {
	if o == nil || o.FullName == nil {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetFullNameOk() (*string, bool) {
	if o == nil || o.FullName == nil {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *AddressModel) HasFullName() bool {
	if o != nil && o.FullName != nil {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *AddressModel) SetFullName(v string) {
	o.FullName = &v
}

// GetName1 returns the Name1 field value if set, zero value otherwise.
func (o *AddressModel) GetName1() string {
	if o == nil || o.Name1 == nil {
		var ret string
		return ret
	}
	return *o.Name1
}

// GetName1Ok returns a tuple with the Name1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetName1Ok() (*string, bool) {
	if o == nil || o.Name1 == nil {
		return nil, false
	}
	return o.Name1, true
}

// HasName1 returns a boolean if a field has been set.
func (o *AddressModel) HasName1() bool {
	if o != nil && o.Name1 != nil {
		return true
	}

	return false
}

// SetName1 gets a reference to the given string and assigns it to the Name1 field.
func (o *AddressModel) SetName1(v string) {
	o.Name1 = &v
}

// GetName2 returns the Name2 field value if set, zero value otherwise.
func (o *AddressModel) GetName2() string {
	if o == nil || o.Name2 == nil {
		var ret string
		return ret
	}
	return *o.Name2
}

// GetName2Ok returns a tuple with the Name2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetName2Ok() (*string, bool) {
	if o == nil || o.Name2 == nil {
		return nil, false
	}
	return o.Name2, true
}

// HasName2 returns a boolean if a field has been set.
func (o *AddressModel) HasName2() bool {
	if o != nil && o.Name2 != nil {
		return true
	}

	return false
}

// SetName2 gets a reference to the given string and assigns it to the Name2 field.
func (o *AddressModel) SetName2(v string) {
	o.Name2 = &v
}

// GetStreet1 returns the Street1 field value if set, zero value otherwise.
func (o *AddressModel) GetStreet1() string {
	if o == nil || o.Street1 == nil {
		var ret string
		return ret
	}
	return *o.Street1
}

// GetStreet1Ok returns a tuple with the Street1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetStreet1Ok() (*string, bool) {
	if o == nil || o.Street1 == nil {
		return nil, false
	}
	return o.Street1, true
}

// HasStreet1 returns a boolean if a field has been set.
func (o *AddressModel) HasStreet1() bool {
	if o != nil && o.Street1 != nil {
		return true
	}

	return false
}

// SetStreet1 gets a reference to the given string and assigns it to the Street1 field.
func (o *AddressModel) SetStreet1(v string) {
	o.Street1 = &v
}

// GetStreet2 returns the Street2 field value if set, zero value otherwise.
func (o *AddressModel) GetStreet2() string {
	if o == nil || o.Street2 == nil {
		var ret string
		return ret
	}
	return *o.Street2
}

// GetStreet2Ok returns a tuple with the Street2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetStreet2Ok() (*string, bool) {
	if o == nil || o.Street2 == nil {
		return nil, false
	}
	return o.Street2, true
}

// HasStreet2 returns a boolean if a field has been set.
func (o *AddressModel) HasStreet2() bool {
	if o != nil && o.Street2 != nil {
		return true
	}

	return false
}

// SetStreet2 gets a reference to the given string and assigns it to the Street2 field.
func (o *AddressModel) SetStreet2(v string) {
	o.Street2 = &v
}

// GetFullAddressline returns the FullAddressline field value if set, zero value otherwise.
func (o *AddressModel) GetFullAddressline() string {
	if o == nil || o.FullAddressline == nil {
		var ret string
		return ret
	}
	return *o.FullAddressline
}

// GetFullAddresslineOk returns a tuple with the FullAddressline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetFullAddresslineOk() (*string, bool) {
	if o == nil || o.FullAddressline == nil {
		return nil, false
	}
	return o.FullAddressline, true
}

// HasFullAddressline returns a boolean if a field has been set.
func (o *AddressModel) HasFullAddressline() bool {
	if o != nil && o.FullAddressline != nil {
		return true
	}

	return false
}

// SetFullAddressline gets a reference to the given string and assigns it to the FullAddressline field.
func (o *AddressModel) SetFullAddressline(v string) {
	o.FullAddressline = &v
}

// GetHouseNr returns the HouseNr field value if set, zero value otherwise.
func (o *AddressModel) GetHouseNr() string {
	if o == nil || o.HouseNr == nil {
		var ret string
		return ret
	}
	return *o.HouseNr
}

// GetHouseNrOk returns a tuple with the HouseNr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetHouseNrOk() (*string, bool) {
	if o == nil || o.HouseNr == nil {
		return nil, false
	}
	return o.HouseNr, true
}

// HasHouseNr returns a boolean if a field has been set.
func (o *AddressModel) HasHouseNr() bool {
	if o != nil && o.HouseNr != nil {
		return true
	}

	return false
}

// SetHouseNr gets a reference to the given string and assigns it to the HouseNr field.
func (o *AddressModel) SetHouseNr(v string) {
	o.HouseNr = &v
}

// GetHouseNrAddendum returns the HouseNrAddendum field value if set, zero value otherwise.
func (o *AddressModel) GetHouseNrAddendum() string {
	if o == nil || o.HouseNrAddendum == nil {
		var ret string
		return ret
	}
	return *o.HouseNrAddendum
}

// GetHouseNrAddendumOk returns a tuple with the HouseNrAddendum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetHouseNrAddendumOk() (*string, bool) {
	if o == nil || o.HouseNrAddendum == nil {
		return nil, false
	}
	return o.HouseNrAddendum, true
}

// HasHouseNrAddendum returns a boolean if a field has been set.
func (o *AddressModel) HasHouseNrAddendum() bool {
	if o != nil && o.HouseNrAddendum != nil {
		return true
	}

	return false
}

// SetHouseNrAddendum gets a reference to the given string and assigns it to the HouseNrAddendum field.
func (o *AddressModel) SetHouseNrAddendum(v string) {
	o.HouseNrAddendum = &v
}

// GetZipcode returns the Zipcode field value if set, zero value otherwise.
func (o *AddressModel) GetZipcode() string {
	if o == nil || o.Zipcode == nil {
		var ret string
		return ret
	}
	return *o.Zipcode
}

// GetZipcodeOk returns a tuple with the Zipcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetZipcodeOk() (*string, bool) {
	if o == nil || o.Zipcode == nil {
		return nil, false
	}
	return o.Zipcode, true
}

// HasZipcode returns a boolean if a field has been set.
func (o *AddressModel) HasZipcode() bool {
	if o != nil && o.Zipcode != nil {
		return true
	}

	return false
}

// SetZipcode gets a reference to the given string and assigns it to the Zipcode field.
func (o *AddressModel) SetZipcode(v string) {
	o.Zipcode = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *AddressModel) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *AddressModel) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *AddressModel) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AddressModel) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AddressModel) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *AddressModel) SetState(v string) {
	o.State = &v
}

// GetIsoCountry returns the IsoCountry field value if set, zero value otherwise.
func (o *AddressModel) GetIsoCountry() string {
	if o == nil || o.IsoCountry == nil {
		var ret string
		return ret
	}
	return *o.IsoCountry
}

// GetIsoCountryOk returns a tuple with the IsoCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetIsoCountryOk() (*string, bool) {
	if o == nil || o.IsoCountry == nil {
		return nil, false
	}
	return o.IsoCountry, true
}

// HasIsoCountry returns a boolean if a field has been set.
func (o *AddressModel) HasIsoCountry() bool {
	if o != nil && o.IsoCountry != nil {
		return true
	}

	return false
}

// SetIsoCountry gets a reference to the given string and assigns it to the IsoCountry field.
func (o *AddressModel) SetIsoCountry(v string) {
	o.IsoCountry = &v
}

// GetCountryName returns the CountryName field value if set, zero value otherwise.
func (o *AddressModel) GetCountryName() string {
	if o == nil || o.CountryName == nil {
		var ret string
		return ret
	}
	return *o.CountryName
}

// GetCountryNameOk returns a tuple with the CountryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetCountryNameOk() (*string, bool) {
	if o == nil || o.CountryName == nil {
		return nil, false
	}
	return o.CountryName, true
}

// HasCountryName returns a boolean if a field has been set.
func (o *AddressModel) HasCountryName() bool {
	if o != nil && o.CountryName != nil {
		return true
	}

	return false
}

// SetCountryName gets a reference to the given string and assigns it to the CountryName field.
func (o *AddressModel) SetCountryName(v string) {
	o.CountryName = &v
}

// GetTimeFrom returns the TimeFrom field value if set, zero value otherwise.
func (o *AddressModel) GetTimeFrom() string {
	if o == nil || o.TimeFrom == nil {
		var ret string
		return ret
	}
	return *o.TimeFrom
}

// GetTimeFromOk returns a tuple with the TimeFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetTimeFromOk() (*string, bool) {
	if o == nil || o.TimeFrom == nil {
		return nil, false
	}
	return o.TimeFrom, true
}

// HasTimeFrom returns a boolean if a field has been set.
func (o *AddressModel) HasTimeFrom() bool {
	if o != nil && o.TimeFrom != nil {
		return true
	}

	return false
}

// SetTimeFrom gets a reference to the given string and assigns it to the TimeFrom field.
func (o *AddressModel) SetTimeFrom(v string) {
	o.TimeFrom = &v
}

// GetTimeTo returns the TimeTo field value if set, zero value otherwise.
func (o *AddressModel) GetTimeTo() string {
	if o == nil || o.TimeTo == nil {
		var ret string
		return ret
	}
	return *o.TimeTo
}

// GetTimeToOk returns a tuple with the TimeTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetTimeToOk() (*string, bool) {
	if o == nil || o.TimeTo == nil {
		return nil, false
	}
	return o.TimeTo, true
}

// HasTimeTo returns a boolean if a field has been set.
func (o *AddressModel) HasTimeTo() bool {
	if o != nil && o.TimeTo != nil {
		return true
	}

	return false
}

// SetTimeTo gets a reference to the given string and assigns it to the TimeTo field.
func (o *AddressModel) SetTimeTo(v string) {
	o.TimeTo = &v
}

// GetOpeningHours returns the OpeningHours field value if set, zero value otherwise.
func (o *AddressModel) GetOpeningHours() []OpeningHoursRuleModel {
	if o == nil || o.OpeningHours == nil {
		var ret []OpeningHoursRuleModel
		return ret
	}
	return *o.OpeningHours
}

// GetOpeningHoursOk returns a tuple with the OpeningHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetOpeningHoursOk() (*[]OpeningHoursRuleModel, bool) {
	if o == nil || o.OpeningHours == nil {
		return nil, false
	}
	return o.OpeningHours, true
}

// HasOpeningHours returns a boolean if a field has been set.
func (o *AddressModel) HasOpeningHours() bool {
	if o != nil && o.OpeningHours != nil {
		return true
	}

	return false
}

// SetOpeningHours gets a reference to the given []OpeningHoursRuleModel and assigns it to the OpeningHours field.
func (o *AddressModel) SetOpeningHours(v []OpeningHoursRuleModel) {
	o.OpeningHours = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *AddressModel) GetDuration() int32 {
	if o == nil || o.Duration == nil {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetDurationOk() (*int32, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *AddressModel) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *AddressModel) SetDuration(v int32) {
	o.Duration = &v
}

// GetAddressTypeNames returns the AddressTypeNames field value if set, zero value otherwise.
func (o *AddressModel) GetAddressTypeNames() []string {
	if o == nil || o.AddressTypeNames == nil {
		var ret []string
		return ret
	}
	return *o.AddressTypeNames
}

// GetAddressTypeNamesOk returns a tuple with the AddressTypeNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetAddressTypeNamesOk() (*[]string, bool) {
	if o == nil || o.AddressTypeNames == nil {
		return nil, false
	}
	return o.AddressTypeNames, true
}

// HasAddressTypeNames returns a boolean if a field has been set.
func (o *AddressModel) HasAddressTypeNames() bool {
	if o != nil && o.AddressTypeNames != nil {
		return true
	}

	return false
}

// SetAddressTypeNames gets a reference to the given []string and assigns it to the AddressTypeNames field.
func (o *AddressModel) SetAddressTypeNames(v []string) {
	o.AddressTypeNames = &v
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *AddressModel) GetEmails() []EmailModel {
	if o == nil || o.Emails == nil {
		var ret []EmailModel
		return ret
	}
	return *o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetEmailsOk() (*[]EmailModel, bool) {
	if o == nil || o.Emails == nil {
		return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *AddressModel) HasEmails() bool {
	if o != nil && o.Emails != nil {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []EmailModel and assigns it to the Emails field.
func (o *AddressModel) SetEmails(v []EmailModel) {
	o.Emails = &v
}

// GetPhoneNrs returns the PhoneNrs field value if set, zero value otherwise.
func (o *AddressModel) GetPhoneNrs() []PhoneNrModel {
	if o == nil || o.PhoneNrs == nil {
		var ret []PhoneNrModel
		return ret
	}
	return *o.PhoneNrs
}

// GetPhoneNrsOk returns a tuple with the PhoneNrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetPhoneNrsOk() (*[]PhoneNrModel, bool) {
	if o == nil || o.PhoneNrs == nil {
		return nil, false
	}
	return o.PhoneNrs, true
}

// HasPhoneNrs returns a boolean if a field has been set.
func (o *AddressModel) HasPhoneNrs() bool {
	if o != nil && o.PhoneNrs != nil {
		return true
	}

	return false
}

// SetPhoneNrs gets a reference to the given []PhoneNrModel and assigns it to the PhoneNrs field.
func (o *AddressModel) SetPhoneNrs(v []PhoneNrModel) {
	o.PhoneNrs = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *AddressModel) GetLatitude() string {
	if o == nil || o.Latitude == nil {
		var ret string
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetLatitudeOk() (*string, bool) {
	if o == nil || o.Latitude == nil {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *AddressModel) HasLatitude() bool {
	if o != nil && o.Latitude != nil {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given string and assigns it to the Latitude field.
func (o *AddressModel) SetLatitude(v string) {
	o.Latitude = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *AddressModel) GetLongitude() string {
	if o == nil || o.Longitude == nil {
		var ret string
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetLongitudeOk() (*string, bool) {
	if o == nil || o.Longitude == nil {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *AddressModel) HasLongitude() bool {
	if o != nil && o.Longitude != nil {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given string and assigns it to the Longitude field.
func (o *AddressModel) SetLongitude(v string) {
	o.Longitude = &v
}

// GetContactPerson returns the ContactPerson field value if set, zero value otherwise.
func (o *AddressModel) GetContactPerson() string {
	if o == nil || o.ContactPerson == nil {
		var ret string
		return ret
	}
	return *o.ContactPerson
}

// GetContactPersonOk returns a tuple with the ContactPerson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetContactPersonOk() (*string, bool) {
	if o == nil || o.ContactPerson == nil {
		return nil, false
	}
	return o.ContactPerson, true
}

// HasContactPerson returns a boolean if a field has been set.
func (o *AddressModel) HasContactPerson() bool {
	if o != nil && o.ContactPerson != nil {
		return true
	}

	return false
}

// SetContactPerson gets a reference to the given string and assigns it to the ContactPerson field.
func (o *AddressModel) SetContactPerson(v string) {
	o.ContactPerson = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AddressModel) GetUserId() int32 {
	if o == nil || o.UserId == nil {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetUserIdOk() (*int32, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AddressModel) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *AddressModel) SetUserId(v int32) {
	o.UserId = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AddressModel) GetLinks() []LinkModel {
	if o == nil || o.Links == nil {
		var ret []LinkModel
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetLinksOk() (*[]LinkModel, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AddressModel) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []LinkModel and assigns it to the Links field.
func (o *AddressModel) SetLinks(v []LinkModel) {
	o.Links = &v
}

// GetMetaData returns the MetaData field value if set, zero value otherwise.
func (o *AddressModel) GetMetaData() []MetaDataModel {
	if o == nil || o.MetaData == nil {
		var ret []MetaDataModel
		return ret
	}
	return *o.MetaData
}

// GetMetaDataOk returns a tuple with the MetaData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetMetaDataOk() (*[]MetaDataModel, bool) {
	if o == nil || o.MetaData == nil {
		return nil, false
	}
	return o.MetaData, true
}

// HasMetaData returns a boolean if a field has been set.
func (o *AddressModel) HasMetaData() bool {
	if o != nil && o.MetaData != nil {
		return true
	}

	return false
}

// SetMetaData gets a reference to the given []MetaDataModel and assigns it to the MetaData field.
func (o *AddressModel) SetMetaData(v []MetaDataModel) {
	o.MetaData = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *AddressModel) GetNotes() []NoteModel {
	if o == nil || o.Notes == nil {
		var ret []NoteModel
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetNotesOk() (*[]NoteModel, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *AddressModel) HasNotes() bool {
	if o != nil && o.Notes != nil {
		return true
	}

	return false
}

// SetNotes gets a reference to the given []NoteModel and assigns it to the Notes field.
func (o *AddressModel) SetNotes(v []NoteModel) {
	o.Notes = &v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *AddressModel) GetFiles() []FileModel {
	if o == nil || o.Files == nil {
		var ret []FileModel
		return ret
	}
	return *o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetFilesOk() (*[]FileModel, bool) {
	if o == nil || o.Files == nil {
		return nil, false
	}
	return o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *AddressModel) HasFiles() bool {
	if o != nil && o.Files != nil {
		return true
	}

	return false
}

// SetFiles gets a reference to the given []FileModel and assigns it to the Files field.
func (o *AddressModel) SetFiles(v []FileModel) {
	o.Files = &v
}

// GetTagNames returns the TagNames field value if set, zero value otherwise.
func (o *AddressModel) GetTagNames() []string {
	if o == nil || o.TagNames == nil {
		var ret []string
		return ret
	}
	return *o.TagNames
}

// GetTagNamesOk returns a tuple with the TagNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetTagNamesOk() (*[]string, bool) {
	if o == nil || o.TagNames == nil {
		return nil, false
	}
	return o.TagNames, true
}

// HasTagNames returns a boolean if a field has been set.
func (o *AddressModel) HasTagNames() bool {
	if o != nil && o.TagNames != nil {
		return true
	}

	return false
}

// SetTagNames gets a reference to the given []string and assigns it to the TagNames field.
func (o *AddressModel) SetTagNames(v []string) {
	o.TagNames = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AddressModel) GetTags() []TagModel {
	if o == nil || o.Tags == nil {
		var ret []TagModel
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressModel) GetTagsOk() (*[]TagModel, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AddressModel) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagModel and assigns it to the Tags field.
func (o *AddressModel) SetTags(v []TagModel) {
	o.Tags = &v
}

func (o AddressModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.AddressId != nil {
		toSerialize["address_id"] = o.AddressId
	}
	if o.PartyId != nil {
		toSerialize["party_id"] = o.PartyId
	}
	if o.PartyName != nil {
		toSerialize["party_name"] = o.PartyName
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.FullName != nil {
		toSerialize["full_name"] = o.FullName
	}
	if o.Name1 != nil {
		toSerialize["name_1"] = o.Name1
	}
	if o.Name2 != nil {
		toSerialize["name_2"] = o.Name2
	}
	if o.Street1 != nil {
		toSerialize["street_1"] = o.Street1
	}
	if o.Street2 != nil {
		toSerialize["street_2"] = o.Street2
	}
	if o.FullAddressline != nil {
		toSerialize["full_addressline"] = o.FullAddressline
	}
	if o.HouseNr != nil {
		toSerialize["house_nr"] = o.HouseNr
	}
	if o.HouseNrAddendum != nil {
		toSerialize["house_nr_addendum"] = o.HouseNrAddendum
	}
	if o.Zipcode != nil {
		toSerialize["zipcode"] = o.Zipcode
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.IsoCountry != nil {
		toSerialize["iso_country"] = o.IsoCountry
	}
	if o.CountryName != nil {
		toSerialize["country_name"] = o.CountryName
	}
	if o.TimeFrom != nil {
		toSerialize["time_from"] = o.TimeFrom
	}
	if o.TimeTo != nil {
		toSerialize["time_to"] = o.TimeTo
	}
	if o.OpeningHours != nil {
		toSerialize["opening_hours"] = o.OpeningHours
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.AddressTypeNames != nil {
		toSerialize["address_type_names"] = o.AddressTypeNames
	}
	if o.Emails != nil {
		toSerialize["emails"] = o.Emails
	}
	if o.PhoneNrs != nil {
		toSerialize["phone_nrs"] = o.PhoneNrs
	}
	if o.Latitude != nil {
		toSerialize["latitude"] = o.Latitude
	}
	if o.Longitude != nil {
		toSerialize["longitude"] = o.Longitude
	}
	if o.ContactPerson != nil {
		toSerialize["contact_person"] = o.ContactPerson
	}
	if o.UserId != nil {
		toSerialize["user_id"] = o.UserId
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.MetaData != nil {
		toSerialize["meta_data"] = o.MetaData
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	if o.TagNames != nil {
		toSerialize["tag_names"] = o.TagNames
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableAddressModel struct {
	value *AddressModel
	isSet bool
}

func (v NullableAddressModel) Get() *AddressModel {
	return v.value
}

func (v *NullableAddressModel) Set(val *AddressModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressModel(val *AddressModel) *NullableAddressModel {
	return &NullableAddressModel{value: val, isSet: true}
}

func (v NullableAddressModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


