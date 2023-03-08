# TimeSlotModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**TimeSlotType** | Pointer to **string** | Time Slot Type. first-entry (:1), planned (:2), actual (:3) | [optional] 
**TimeSlotTypeId** | Pointer to **int64** | Time Slot Type ID, by default 1 if left out of the request. 1: first-entry, 2: planned, 3: actual | [optional] 
**ActivityId** | Pointer to **int64** | Activity ID to which this TimeSlot belongs | [optional] 
**DateFrom** | Pointer to **string** |  | [optional] 
**TimeFrom** | Pointer to **string** |  | [optional] 
**DateTimeFrom** | Pointer to **time.Time** |  | [optional] 
**DateTo** | Pointer to **string** |  | [optional] 
**TimeTo** | Pointer to **string** |  | [optional] 
**DateTimeTo** | Pointer to **time.Time** |  | [optional] 
**Planned** | Pointer to **bool** | true if this time_slot was used to plan the activity within | [optional] 

## Methods

### NewTimeSlotModel

`func NewTimeSlotModel() *TimeSlotModel`

NewTimeSlotModel instantiates a new TimeSlotModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSlotModelWithDefaults

`func NewTimeSlotModelWithDefaults() *TimeSlotModel`

NewTimeSlotModelWithDefaults instantiates a new TimeSlotModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TimeSlotModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimeSlotModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimeSlotModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *TimeSlotModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimeSlotType

`func (o *TimeSlotModel) GetTimeSlotType() string`

GetTimeSlotType returns the TimeSlotType field if non-nil, zero value otherwise.

### GetTimeSlotTypeOk

`func (o *TimeSlotModel) GetTimeSlotTypeOk() (*string, bool)`

GetTimeSlotTypeOk returns a tuple with the TimeSlotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSlotType

`func (o *TimeSlotModel) SetTimeSlotType(v string)`

SetTimeSlotType sets TimeSlotType field to given value.

### HasTimeSlotType

`func (o *TimeSlotModel) HasTimeSlotType() bool`

HasTimeSlotType returns a boolean if a field has been set.

### GetTimeSlotTypeId

`func (o *TimeSlotModel) GetTimeSlotTypeId() int64`

GetTimeSlotTypeId returns the TimeSlotTypeId field if non-nil, zero value otherwise.

### GetTimeSlotTypeIdOk

`func (o *TimeSlotModel) GetTimeSlotTypeIdOk() (*int64, bool)`

GetTimeSlotTypeIdOk returns a tuple with the TimeSlotTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSlotTypeId

`func (o *TimeSlotModel) SetTimeSlotTypeId(v int64)`

SetTimeSlotTypeId sets TimeSlotTypeId field to given value.

### HasTimeSlotTypeId

`func (o *TimeSlotModel) HasTimeSlotTypeId() bool`

HasTimeSlotTypeId returns a boolean if a field has been set.

### GetActivityId

`func (o *TimeSlotModel) GetActivityId() int64`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *TimeSlotModel) GetActivityIdOk() (*int64, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *TimeSlotModel) SetActivityId(v int64)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *TimeSlotModel) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### GetDateFrom

`func (o *TimeSlotModel) GetDateFrom() string`

GetDateFrom returns the DateFrom field if non-nil, zero value otherwise.

### GetDateFromOk

`func (o *TimeSlotModel) GetDateFromOk() (*string, bool)`

GetDateFromOk returns a tuple with the DateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFrom

`func (o *TimeSlotModel) SetDateFrom(v string)`

SetDateFrom sets DateFrom field to given value.

### HasDateFrom

`func (o *TimeSlotModel) HasDateFrom() bool`

HasDateFrom returns a boolean if a field has been set.

### GetTimeFrom

`func (o *TimeSlotModel) GetTimeFrom() string`

GetTimeFrom returns the TimeFrom field if non-nil, zero value otherwise.

### GetTimeFromOk

`func (o *TimeSlotModel) GetTimeFromOk() (*string, bool)`

GetTimeFromOk returns a tuple with the TimeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrom

`func (o *TimeSlotModel) SetTimeFrom(v string)`

SetTimeFrom sets TimeFrom field to given value.

### HasTimeFrom

`func (o *TimeSlotModel) HasTimeFrom() bool`

HasTimeFrom returns a boolean if a field has been set.

### GetDateTimeFrom

`func (o *TimeSlotModel) GetDateTimeFrom() time.Time`

GetDateTimeFrom returns the DateTimeFrom field if non-nil, zero value otherwise.

### GetDateTimeFromOk

`func (o *TimeSlotModel) GetDateTimeFromOk() (*time.Time, bool)`

GetDateTimeFromOk returns a tuple with the DateTimeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFrom

`func (o *TimeSlotModel) SetDateTimeFrom(v time.Time)`

SetDateTimeFrom sets DateTimeFrom field to given value.

### HasDateTimeFrom

`func (o *TimeSlotModel) HasDateTimeFrom() bool`

HasDateTimeFrom returns a boolean if a field has been set.

### GetDateTo

`func (o *TimeSlotModel) GetDateTo() string`

GetDateTo returns the DateTo field if non-nil, zero value otherwise.

### GetDateToOk

`func (o *TimeSlotModel) GetDateToOk() (*string, bool)`

GetDateToOk returns a tuple with the DateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTo

`func (o *TimeSlotModel) SetDateTo(v string)`

SetDateTo sets DateTo field to given value.

### HasDateTo

`func (o *TimeSlotModel) HasDateTo() bool`

HasDateTo returns a boolean if a field has been set.

### GetTimeTo

`func (o *TimeSlotModel) GetTimeTo() string`

GetTimeTo returns the TimeTo field if non-nil, zero value otherwise.

### GetTimeToOk

`func (o *TimeSlotModel) GetTimeToOk() (*string, bool)`

GetTimeToOk returns a tuple with the TimeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeTo

`func (o *TimeSlotModel) SetTimeTo(v string)`

SetTimeTo sets TimeTo field to given value.

### HasTimeTo

`func (o *TimeSlotModel) HasTimeTo() bool`

HasTimeTo returns a boolean if a field has been set.

### GetDateTimeTo

`func (o *TimeSlotModel) GetDateTimeTo() time.Time`

GetDateTimeTo returns the DateTimeTo field if non-nil, zero value otherwise.

### GetDateTimeToOk

`func (o *TimeSlotModel) GetDateTimeToOk() (*time.Time, bool)`

GetDateTimeToOk returns a tuple with the DateTimeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeTo

`func (o *TimeSlotModel) SetDateTimeTo(v time.Time)`

SetDateTimeTo sets DateTimeTo field to given value.

### HasDateTimeTo

`func (o *TimeSlotModel) HasDateTimeTo() bool`

HasDateTimeTo returns a boolean if a field has been set.

### GetPlanned

`func (o *TimeSlotModel) GetPlanned() bool`

GetPlanned returns the Planned field if non-nil, zero value otherwise.

### GetPlannedOk

`func (o *TimeSlotModel) GetPlannedOk() (*bool, bool)`

GetPlannedOk returns a tuple with the Planned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanned

`func (o *TimeSlotModel) SetPlanned(v bool)`

SetPlanned sets Planned field to given value.

### HasPlanned

`func (o *TimeSlotModel) HasPlanned() bool`

HasPlanned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


