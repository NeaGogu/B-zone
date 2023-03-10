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

// TimeSlotModel struct for TimeSlotModel
type TimeSlotModel struct {
	// 
	Id *int64 `json:"id,omitempty"`
	// Time Slot Type. first-entry (:1), planned (:2), actual (:3)
	TimeSlotType *string `json:"time_slot_type,omitempty"`
	// Time Slot Type ID, by default 1 if left out of the request. 1: first-entry, 2: planned, 3: actual
	TimeSlotTypeId *int64 `json:"time_slot_type_id,omitempty"`
	// Activity ID to which this TimeSlot belongs
	ActivityId *int64 `json:"activity_id,omitempty"`
	// 
	DateFrom *string `json:"date_from,omitempty"`
	// 
	TimeFrom *string `json:"time_from,omitempty"`
	// 
	DateTimeFrom *time.Time `json:"date_time_from,omitempty"`
	// 
	DateTo *string `json:"date_to,omitempty"`
	// 
	TimeTo *string `json:"time_to,omitempty"`
	// 
	DateTimeTo *time.Time `json:"date_time_to,omitempty"`
	// true if this time_slot was used to plan the activity within
	Planned *bool `json:"planned,omitempty"`
}

// NewTimeSlotModel instantiates a new TimeSlotModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeSlotModel() *TimeSlotModel {
	this := TimeSlotModel{}
	return &this
}

// NewTimeSlotModelWithDefaults instantiates a new TimeSlotModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeSlotModelWithDefaults() *TimeSlotModel {
	this := TimeSlotModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TimeSlotModel) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSlotModel) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TimeSlotModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *TimeSlotModel) SetId(v int64) {
	o.Id = &v
}

// GetTimeSlotType returns the TimeSlotType field value if set, zero value otherwise.
func (o *TimeSlotModel) GetTimeSlotType() string {
	if o == nil || o.TimeSlotType == nil {
		var ret string
		return ret
	}
	return *o.TimeSlotType
}

// GetTimeSlotTypeOk returns a tuple with the TimeSlotType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSlotModel) GetTimeSlotTypeOk() (*string, bool) {
	if o == nil || o.TimeSlotType == nil {
		return nil, false
	}
	return o.TimeSlotType, true
}

// HasTimeSlotType returns a boolean if a field has been set.
func (o *TimeSlotModel) HasTimeSlotType() bool {
	if o != nil && o.TimeSlotType != nil {
		return true
	}

	return false
}

// SetTimeSlotType gets a reference to the given string and assigns it to the TimeSlotType field.
func (o *TimeSlotModel) SetTimeSlotType(v string) {
	o.TimeSlotType = &v
}

// GetTimeSlotTypeId returns the TimeSlotTypeId field value if set, zero value otherwise.
func (o *TimeSlotModel) GetTimeSlotTypeId() int64 {
	if o == nil || o.TimeSlotTypeId == nil {
		var ret int64
		return ret
	}
	return *o.TimeSlotTypeId
}

// GetTimeSlotTypeIdOk returns a tuple with the TimeSlotTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSlotModel) GetTimeSlotTypeIdOk() (*int64, bool) {
	if o == nil || o.TimeSlotTypeId == nil {
		return nil, false
	}
	return o.TimeSlotTypeId, true
}

// HasTimeSlotTypeId returns a boolean if a field has been set.
func (o *TimeSlotModel) HasTimeSlotTypeId() bool {
	if o != nil && o.TimeSlotTypeId != nil {
		return true
	}

	return false
}

// SetTimeSlotTypeId gets a reference to the given int64 and assigns it to the TimeSlotTypeId field.
func (o *TimeSlotModel) SetTimeSlotTypeId(v int64) {
	o.TimeSlotTypeId = &v
}

// GetActivityId returns the ActivityId field value if set, zero value otherwise.
func (o *TimeSlotModel) GetActivityId() int64 {
	if o == nil || o.ActivityId == nil {
		var ret int64
		return ret
	}
	return *o.ActivityId
}

// GetActivityIdOk returns a tuple with the ActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSlotModel) GetActivityIdOk() (*int64, bool) {
	if o == nil || o.ActivityId == nil {
		return nil, false
	}
	return o.ActivityId, true
}

// HasActivityId returns a boolean if a field has been set.
func (o *TimeSlotModel) HasActivityId() bool {
	if o != nil && o.ActivityId != nil {
		return true
	}

	return false
}

// SetActivityId gets a reference to the given int64 and assigns it to the ActivityId field.
func (o *TimeSlotModel) SetActivityId(v int64) {
	o.ActivityId = &v
}

// GetDateFrom returns the DateFrom field value if set, zero value otherwise.
func (o *TimeSlotModel) GetDateFrom() string {
	if o == nil || o.DateFrom == nil {
		var ret string
		return ret
	}
	return *o.DateFrom
}

// GetDateFromOk returns a tuple with the DateFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSlotModel) GetDateFromOk() (*string, bool) {
	if o == nil || o.DateFrom == nil {
		return nil, false
	}
	return o.DateFrom, true
}

// HasDateFrom returns a boolean if a field has been set.
func (o *TimeSlotModel) HasDateFrom() bool {
	if o != nil && o.DateFrom != nil {
		return true
	}

	return false
}

// SetDateFrom gets a reference to the given string and assigns it to the DateFrom field.
func (o *TimeSlotModel) SetDateFrom(v string) {
	o.DateFrom = &v
}

// GetTimeFrom returns the TimeFrom field value if set, zero value otherwise.
func (o *TimeSlotModel) GetTimeFrom() string {
	if o == nil || o.TimeFrom == nil {
		var ret string
		return ret
	}
	return *o.TimeFrom
}

// GetTimeFromOk returns a tuple with the TimeFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSlotModel) GetTimeFromOk() (*string, bool) {
	if o == nil || o.TimeFrom == nil {
		return nil, false
	}
	return o.TimeFrom, true
}

// HasTimeFrom returns a boolean if a field has been set.
func (o *TimeSlotModel) HasTimeFrom() bool {
	if o != nil && o.TimeFrom != nil {
		return true
	}

	return false
}

// SetTimeFrom gets a reference to the given string and assigns it to the TimeFrom field.
func (o *TimeSlotModel) SetTimeFrom(v string) {
	o.TimeFrom = &v
}

// GetDateTimeFrom returns the DateTimeFrom field value if set, zero value otherwise.
func (o *TimeSlotModel) GetDateTimeFrom() time.Time {
	if o == nil || o.DateTimeFrom == nil {
		var ret time.Time
		return ret
	}
	return *o.DateTimeFrom
}

// GetDateTimeFromOk returns a tuple with the DateTimeFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSlotModel) GetDateTimeFromOk() (*time.Time, bool) {
	if o == nil || o.DateTimeFrom == nil {
		return nil, false
	}
	return o.DateTimeFrom, true
}

// HasDateTimeFrom returns a boolean if a field has been set.
func (o *TimeSlotModel) HasDateTimeFrom() bool {
	if o != nil && o.DateTimeFrom != nil {
		return true
	}

	return false
}

// SetDateTimeFrom gets a reference to the given time.Time and assigns it to the DateTimeFrom field.
func (o *TimeSlotModel) SetDateTimeFrom(v time.Time) {
	o.DateTimeFrom = &v
}

// GetDateTo returns the DateTo field value if set, zero value otherwise.
func (o *TimeSlotModel) GetDateTo() string {
	if o == nil || o.DateTo == nil {
		var ret string
		return ret
	}
	return *o.DateTo
}

// GetDateToOk returns a tuple with the DateTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSlotModel) GetDateToOk() (*string, bool) {
	if o == nil || o.DateTo == nil {
		return nil, false
	}
	return o.DateTo, true
}

// HasDateTo returns a boolean if a field has been set.
func (o *TimeSlotModel) HasDateTo() bool {
	if o != nil && o.DateTo != nil {
		return true
	}

	return false
}

// SetDateTo gets a reference to the given string and assigns it to the DateTo field.
func (o *TimeSlotModel) SetDateTo(v string) {
	o.DateTo = &v
}

// GetTimeTo returns the TimeTo field value if set, zero value otherwise.
func (o *TimeSlotModel) GetTimeTo() string {
	if o == nil || o.TimeTo == nil {
		var ret string
		return ret
	}
	return *o.TimeTo
}

// GetTimeToOk returns a tuple with the TimeTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSlotModel) GetTimeToOk() (*string, bool) {
	if o == nil || o.TimeTo == nil {
		return nil, false
	}
	return o.TimeTo, true
}

// HasTimeTo returns a boolean if a field has been set.
func (o *TimeSlotModel) HasTimeTo() bool {
	if o != nil && o.TimeTo != nil {
		return true
	}

	return false
}

// SetTimeTo gets a reference to the given string and assigns it to the TimeTo field.
func (o *TimeSlotModel) SetTimeTo(v string) {
	o.TimeTo = &v
}

// GetDateTimeTo returns the DateTimeTo field value if set, zero value otherwise.
func (o *TimeSlotModel) GetDateTimeTo() time.Time {
	if o == nil || o.DateTimeTo == nil {
		var ret time.Time
		return ret
	}
	return *o.DateTimeTo
}

// GetDateTimeToOk returns a tuple with the DateTimeTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSlotModel) GetDateTimeToOk() (*time.Time, bool) {
	if o == nil || o.DateTimeTo == nil {
		return nil, false
	}
	return o.DateTimeTo, true
}

// HasDateTimeTo returns a boolean if a field has been set.
func (o *TimeSlotModel) HasDateTimeTo() bool {
	if o != nil && o.DateTimeTo != nil {
		return true
	}

	return false
}

// SetDateTimeTo gets a reference to the given time.Time and assigns it to the DateTimeTo field.
func (o *TimeSlotModel) SetDateTimeTo(v time.Time) {
	o.DateTimeTo = &v
}

// GetPlanned returns the Planned field value if set, zero value otherwise.
func (o *TimeSlotModel) GetPlanned() bool {
	if o == nil || o.Planned == nil {
		var ret bool
		return ret
	}
	return *o.Planned
}

// GetPlannedOk returns a tuple with the Planned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSlotModel) GetPlannedOk() (*bool, bool) {
	if o == nil || o.Planned == nil {
		return nil, false
	}
	return o.Planned, true
}

// HasPlanned returns a boolean if a field has been set.
func (o *TimeSlotModel) HasPlanned() bool {
	if o != nil && o.Planned != nil {
		return true
	}

	return false
}

// SetPlanned gets a reference to the given bool and assigns it to the Planned field.
func (o *TimeSlotModel) SetPlanned(v bool) {
	o.Planned = &v
}

func (o TimeSlotModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.TimeSlotType != nil {
		toSerialize["time_slot_type"] = o.TimeSlotType
	}
	if o.TimeSlotTypeId != nil {
		toSerialize["time_slot_type_id"] = o.TimeSlotTypeId
	}
	if o.ActivityId != nil {
		toSerialize["activity_id"] = o.ActivityId
	}
	if o.DateFrom != nil {
		toSerialize["date_from"] = o.DateFrom
	}
	if o.TimeFrom != nil {
		toSerialize["time_from"] = o.TimeFrom
	}
	if o.DateTimeFrom != nil {
		toSerialize["date_time_from"] = o.DateTimeFrom
	}
	if o.DateTo != nil {
		toSerialize["date_to"] = o.DateTo
	}
	if o.TimeTo != nil {
		toSerialize["time_to"] = o.TimeTo
	}
	if o.DateTimeTo != nil {
		toSerialize["date_time_to"] = o.DateTimeTo
	}
	if o.Planned != nil {
		toSerialize["planned"] = o.Planned
	}
	return json.Marshal(toSerialize)
}

type NullableTimeSlotModel struct {
	value *TimeSlotModel
	isSet bool
}

func (v NullableTimeSlotModel) Get() *TimeSlotModel {
	return v.value
}

func (v *NullableTimeSlotModel) Set(val *TimeSlotModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeSlotModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeSlotModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeSlotModel(val *TimeSlotModel) *NullableTimeSlotModel {
	return &NullableTimeSlotModel{value: val, isSet: true}
}

func (v NullableTimeSlotModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeSlotModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

