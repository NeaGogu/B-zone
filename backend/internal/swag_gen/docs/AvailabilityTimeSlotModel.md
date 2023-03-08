# AvailabilityTimeSlotModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Key** | Pointer to **string** | unique key per analyzed time slot, uuid type | [optional] 
**DateTimeFrom** | Pointer to **time.Time** |  | [optional] 
**DateTimeTo** | Pointer to **time.Time** |  | [optional] 
**ProposedPlanDateTimeFrom** | Pointer to **time.Time** |  | [optional] 
**ProposedPlanDateTimeTo** | Pointer to **time.Time** |  | [optional] 
**ProposedDriver** | Pointer to [**DriverModel**](DriverModel.md) |  | [optional] 
**Impact** | Pointer to [**[]AvailabilityTimeSlotImpactModel**](AvailabilityTimeSlotImpactModel.md) |  | [optional] 
**FollowUpTimeSlots** | Pointer to [**[]AvailabilityFollowUpTimeSlotModel**](AvailabilityFollowUpTimeSlotModel.md) |  | [optional] 

## Methods

### NewAvailabilityTimeSlotModel

`func NewAvailabilityTimeSlotModel() *AvailabilityTimeSlotModel`

NewAvailabilityTimeSlotModel instantiates a new AvailabilityTimeSlotModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailabilityTimeSlotModelWithDefaults

`func NewAvailabilityTimeSlotModelWithDefaults() *AvailabilityTimeSlotModel`

NewAvailabilityTimeSlotModelWithDefaults instantiates a new AvailabilityTimeSlotModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AvailabilityTimeSlotModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AvailabilityTimeSlotModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AvailabilityTimeSlotModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AvailabilityTimeSlotModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *AvailabilityTimeSlotModel) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AvailabilityTimeSlotModel) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AvailabilityTimeSlotModel) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AvailabilityTimeSlotModel) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetDateTimeFrom

`func (o *AvailabilityTimeSlotModel) GetDateTimeFrom() time.Time`

GetDateTimeFrom returns the DateTimeFrom field if non-nil, zero value otherwise.

### GetDateTimeFromOk

`func (o *AvailabilityTimeSlotModel) GetDateTimeFromOk() (*time.Time, bool)`

GetDateTimeFromOk returns a tuple with the DateTimeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFrom

`func (o *AvailabilityTimeSlotModel) SetDateTimeFrom(v time.Time)`

SetDateTimeFrom sets DateTimeFrom field to given value.

### HasDateTimeFrom

`func (o *AvailabilityTimeSlotModel) HasDateTimeFrom() bool`

HasDateTimeFrom returns a boolean if a field has been set.

### GetDateTimeTo

`func (o *AvailabilityTimeSlotModel) GetDateTimeTo() time.Time`

GetDateTimeTo returns the DateTimeTo field if non-nil, zero value otherwise.

### GetDateTimeToOk

`func (o *AvailabilityTimeSlotModel) GetDateTimeToOk() (*time.Time, bool)`

GetDateTimeToOk returns a tuple with the DateTimeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeTo

`func (o *AvailabilityTimeSlotModel) SetDateTimeTo(v time.Time)`

SetDateTimeTo sets DateTimeTo field to given value.

### HasDateTimeTo

`func (o *AvailabilityTimeSlotModel) HasDateTimeTo() bool`

HasDateTimeTo returns a boolean if a field has been set.

### GetProposedPlanDateTimeFrom

`func (o *AvailabilityTimeSlotModel) GetProposedPlanDateTimeFrom() time.Time`

GetProposedPlanDateTimeFrom returns the ProposedPlanDateTimeFrom field if non-nil, zero value otherwise.

### GetProposedPlanDateTimeFromOk

`func (o *AvailabilityTimeSlotModel) GetProposedPlanDateTimeFromOk() (*time.Time, bool)`

GetProposedPlanDateTimeFromOk returns a tuple with the ProposedPlanDateTimeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedPlanDateTimeFrom

`func (o *AvailabilityTimeSlotModel) SetProposedPlanDateTimeFrom(v time.Time)`

SetProposedPlanDateTimeFrom sets ProposedPlanDateTimeFrom field to given value.

### HasProposedPlanDateTimeFrom

`func (o *AvailabilityTimeSlotModel) HasProposedPlanDateTimeFrom() bool`

HasProposedPlanDateTimeFrom returns a boolean if a field has been set.

### GetProposedPlanDateTimeTo

`func (o *AvailabilityTimeSlotModel) GetProposedPlanDateTimeTo() time.Time`

GetProposedPlanDateTimeTo returns the ProposedPlanDateTimeTo field if non-nil, zero value otherwise.

### GetProposedPlanDateTimeToOk

`func (o *AvailabilityTimeSlotModel) GetProposedPlanDateTimeToOk() (*time.Time, bool)`

GetProposedPlanDateTimeToOk returns a tuple with the ProposedPlanDateTimeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedPlanDateTimeTo

`func (o *AvailabilityTimeSlotModel) SetProposedPlanDateTimeTo(v time.Time)`

SetProposedPlanDateTimeTo sets ProposedPlanDateTimeTo field to given value.

### HasProposedPlanDateTimeTo

`func (o *AvailabilityTimeSlotModel) HasProposedPlanDateTimeTo() bool`

HasProposedPlanDateTimeTo returns a boolean if a field has been set.

### GetProposedDriver

`func (o *AvailabilityTimeSlotModel) GetProposedDriver() DriverModel`

GetProposedDriver returns the ProposedDriver field if non-nil, zero value otherwise.

### GetProposedDriverOk

`func (o *AvailabilityTimeSlotModel) GetProposedDriverOk() (*DriverModel, bool)`

GetProposedDriverOk returns a tuple with the ProposedDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedDriver

`func (o *AvailabilityTimeSlotModel) SetProposedDriver(v DriverModel)`

SetProposedDriver sets ProposedDriver field to given value.

### HasProposedDriver

`func (o *AvailabilityTimeSlotModel) HasProposedDriver() bool`

HasProposedDriver returns a boolean if a field has been set.

### GetImpact

`func (o *AvailabilityTimeSlotModel) GetImpact() []AvailabilityTimeSlotImpactModel`

GetImpact returns the Impact field if non-nil, zero value otherwise.

### GetImpactOk

`func (o *AvailabilityTimeSlotModel) GetImpactOk() (*[]AvailabilityTimeSlotImpactModel, bool)`

GetImpactOk returns a tuple with the Impact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpact

`func (o *AvailabilityTimeSlotModel) SetImpact(v []AvailabilityTimeSlotImpactModel)`

SetImpact sets Impact field to given value.

### HasImpact

`func (o *AvailabilityTimeSlotModel) HasImpact() bool`

HasImpact returns a boolean if a field has been set.

### GetFollowUpTimeSlots

`func (o *AvailabilityTimeSlotModel) GetFollowUpTimeSlots() []AvailabilityFollowUpTimeSlotModel`

GetFollowUpTimeSlots returns the FollowUpTimeSlots field if non-nil, zero value otherwise.

### GetFollowUpTimeSlotsOk

`func (o *AvailabilityTimeSlotModel) GetFollowUpTimeSlotsOk() (*[]AvailabilityFollowUpTimeSlotModel, bool)`

GetFollowUpTimeSlotsOk returns a tuple with the FollowUpTimeSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowUpTimeSlots

`func (o *AvailabilityTimeSlotModel) SetFollowUpTimeSlots(v []AvailabilityFollowUpTimeSlotModel)`

SetFollowUpTimeSlots sets FollowUpTimeSlots field to given value.

### HasFollowUpTimeSlots

`func (o *AvailabilityTimeSlotModel) HasFollowUpTimeSlots() bool`

HasFollowUpTimeSlots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


