# RoutesEtaFiltersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **[]int32** | Unique Identifier | [optional] 
**DateTimeFrom** | Pointer to **time.Time** |  | [optional] 
**DateTimeTo** | Pointer to **time.Time** |  | [optional] 
**EarliestDateTimeSince** | Pointer to **time.Time** | filter routes with an Earliest DateTime To since this input | [optional] 
**EarliestDateTimeTill** | Pointer to **time.Time** | filter routes with an Earliest DateTime To till this input | [optional] 
**LatestDateTimeSince** | Pointer to **time.Time** | filter routes with an Latest DateTime To since this input | [optional] 
**LatestDateTimeTill** | Pointer to **time.Time** | filter routes with an Latest DateTime To till this input | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAtSince** | Pointer to **time.Time** | filter routes with an updated at date-time since this input | [optional] 
**UpdatedAtTill** | Pointer to **time.Time** | filter routes with an updated at date-time till this input | [optional] 
**StatusId** | Pointer to **[]int32** |  | [optional] 
**DriverId** | Pointer to **[]int32** |  | [optional] 
**RecurrenceId** | Pointer to **int32** | Recurrence ID | [optional] 
**TagNames** | Pointer to **[]string** | Tag names | [optional] 
**ZoneNames** | Pointer to **[]string** | Zone names | [optional] 
**Optimized** | Pointer to **[]bool** | Optimized status of Route. | [optional] 
**Blocked** | Pointer to **[]bool** | Blocked status of Route | [optional] 

## Methods

### NewRoutesEtaFiltersModel

`func NewRoutesEtaFiltersModel() *RoutesEtaFiltersModel`

NewRoutesEtaFiltersModel instantiates a new RoutesEtaFiltersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutesEtaFiltersModelWithDefaults

`func NewRoutesEtaFiltersModelWithDefaults() *RoutesEtaFiltersModel`

NewRoutesEtaFiltersModelWithDefaults instantiates a new RoutesEtaFiltersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RoutesEtaFiltersModel) GetId() []int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoutesEtaFiltersModel) GetIdOk() (*[]int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoutesEtaFiltersModel) SetId(v []int32)`

SetId sets Id field to given value.

### HasId

`func (o *RoutesEtaFiltersModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDateTimeFrom

`func (o *RoutesEtaFiltersModel) GetDateTimeFrom() time.Time`

GetDateTimeFrom returns the DateTimeFrom field if non-nil, zero value otherwise.

### GetDateTimeFromOk

`func (o *RoutesEtaFiltersModel) GetDateTimeFromOk() (*time.Time, bool)`

GetDateTimeFromOk returns a tuple with the DateTimeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFrom

`func (o *RoutesEtaFiltersModel) SetDateTimeFrom(v time.Time)`

SetDateTimeFrom sets DateTimeFrom field to given value.

### HasDateTimeFrom

`func (o *RoutesEtaFiltersModel) HasDateTimeFrom() bool`

HasDateTimeFrom returns a boolean if a field has been set.

### GetDateTimeTo

`func (o *RoutesEtaFiltersModel) GetDateTimeTo() time.Time`

GetDateTimeTo returns the DateTimeTo field if non-nil, zero value otherwise.

### GetDateTimeToOk

`func (o *RoutesEtaFiltersModel) GetDateTimeToOk() (*time.Time, bool)`

GetDateTimeToOk returns a tuple with the DateTimeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeTo

`func (o *RoutesEtaFiltersModel) SetDateTimeTo(v time.Time)`

SetDateTimeTo sets DateTimeTo field to given value.

### HasDateTimeTo

`func (o *RoutesEtaFiltersModel) HasDateTimeTo() bool`

HasDateTimeTo returns a boolean if a field has been set.

### GetEarliestDateTimeSince

`func (o *RoutesEtaFiltersModel) GetEarliestDateTimeSince() time.Time`

GetEarliestDateTimeSince returns the EarliestDateTimeSince field if non-nil, zero value otherwise.

### GetEarliestDateTimeSinceOk

`func (o *RoutesEtaFiltersModel) GetEarliestDateTimeSinceOk() (*time.Time, bool)`

GetEarliestDateTimeSinceOk returns a tuple with the EarliestDateTimeSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestDateTimeSince

`func (o *RoutesEtaFiltersModel) SetEarliestDateTimeSince(v time.Time)`

SetEarliestDateTimeSince sets EarliestDateTimeSince field to given value.

### HasEarliestDateTimeSince

`func (o *RoutesEtaFiltersModel) HasEarliestDateTimeSince() bool`

HasEarliestDateTimeSince returns a boolean if a field has been set.

### GetEarliestDateTimeTill

`func (o *RoutesEtaFiltersModel) GetEarliestDateTimeTill() time.Time`

GetEarliestDateTimeTill returns the EarliestDateTimeTill field if non-nil, zero value otherwise.

### GetEarliestDateTimeTillOk

`func (o *RoutesEtaFiltersModel) GetEarliestDateTimeTillOk() (*time.Time, bool)`

GetEarliestDateTimeTillOk returns a tuple with the EarliestDateTimeTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestDateTimeTill

`func (o *RoutesEtaFiltersModel) SetEarliestDateTimeTill(v time.Time)`

SetEarliestDateTimeTill sets EarliestDateTimeTill field to given value.

### HasEarliestDateTimeTill

`func (o *RoutesEtaFiltersModel) HasEarliestDateTimeTill() bool`

HasEarliestDateTimeTill returns a boolean if a field has been set.

### GetLatestDateTimeSince

`func (o *RoutesEtaFiltersModel) GetLatestDateTimeSince() time.Time`

GetLatestDateTimeSince returns the LatestDateTimeSince field if non-nil, zero value otherwise.

### GetLatestDateTimeSinceOk

`func (o *RoutesEtaFiltersModel) GetLatestDateTimeSinceOk() (*time.Time, bool)`

GetLatestDateTimeSinceOk returns a tuple with the LatestDateTimeSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDateTimeSince

`func (o *RoutesEtaFiltersModel) SetLatestDateTimeSince(v time.Time)`

SetLatestDateTimeSince sets LatestDateTimeSince field to given value.

### HasLatestDateTimeSince

`func (o *RoutesEtaFiltersModel) HasLatestDateTimeSince() bool`

HasLatestDateTimeSince returns a boolean if a field has been set.

### GetLatestDateTimeTill

`func (o *RoutesEtaFiltersModel) GetLatestDateTimeTill() time.Time`

GetLatestDateTimeTill returns the LatestDateTimeTill field if non-nil, zero value otherwise.

### GetLatestDateTimeTillOk

`func (o *RoutesEtaFiltersModel) GetLatestDateTimeTillOk() (*time.Time, bool)`

GetLatestDateTimeTillOk returns a tuple with the LatestDateTimeTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDateTimeTill

`func (o *RoutesEtaFiltersModel) SetLatestDateTimeTill(v time.Time)`

SetLatestDateTimeTill sets LatestDateTimeTill field to given value.

### HasLatestDateTimeTill

`func (o *RoutesEtaFiltersModel) HasLatestDateTimeTill() bool`

HasLatestDateTimeTill returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RoutesEtaFiltersModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RoutesEtaFiltersModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RoutesEtaFiltersModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RoutesEtaFiltersModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedAtSince

`func (o *RoutesEtaFiltersModel) GetUpdatedAtSince() time.Time`

GetUpdatedAtSince returns the UpdatedAtSince field if non-nil, zero value otherwise.

### GetUpdatedAtSinceOk

`func (o *RoutesEtaFiltersModel) GetUpdatedAtSinceOk() (*time.Time, bool)`

GetUpdatedAtSinceOk returns a tuple with the UpdatedAtSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtSince

`func (o *RoutesEtaFiltersModel) SetUpdatedAtSince(v time.Time)`

SetUpdatedAtSince sets UpdatedAtSince field to given value.

### HasUpdatedAtSince

`func (o *RoutesEtaFiltersModel) HasUpdatedAtSince() bool`

HasUpdatedAtSince returns a boolean if a field has been set.

### GetUpdatedAtTill

`func (o *RoutesEtaFiltersModel) GetUpdatedAtTill() time.Time`

GetUpdatedAtTill returns the UpdatedAtTill field if non-nil, zero value otherwise.

### GetUpdatedAtTillOk

`func (o *RoutesEtaFiltersModel) GetUpdatedAtTillOk() (*time.Time, bool)`

GetUpdatedAtTillOk returns a tuple with the UpdatedAtTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtTill

`func (o *RoutesEtaFiltersModel) SetUpdatedAtTill(v time.Time)`

SetUpdatedAtTill sets UpdatedAtTill field to given value.

### HasUpdatedAtTill

`func (o *RoutesEtaFiltersModel) HasUpdatedAtTill() bool`

HasUpdatedAtTill returns a boolean if a field has been set.

### GetStatusId

`func (o *RoutesEtaFiltersModel) GetStatusId() []int32`

GetStatusId returns the StatusId field if non-nil, zero value otherwise.

### GetStatusIdOk

`func (o *RoutesEtaFiltersModel) GetStatusIdOk() (*[]int32, bool)`

GetStatusIdOk returns a tuple with the StatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusId

`func (o *RoutesEtaFiltersModel) SetStatusId(v []int32)`

SetStatusId sets StatusId field to given value.

### HasStatusId

`func (o *RoutesEtaFiltersModel) HasStatusId() bool`

HasStatusId returns a boolean if a field has been set.

### GetDriverId

`func (o *RoutesEtaFiltersModel) GetDriverId() []int32`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *RoutesEtaFiltersModel) GetDriverIdOk() (*[]int32, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *RoutesEtaFiltersModel) SetDriverId(v []int32)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *RoutesEtaFiltersModel) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### GetRecurrenceId

`func (o *RoutesEtaFiltersModel) GetRecurrenceId() int32`

GetRecurrenceId returns the RecurrenceId field if non-nil, zero value otherwise.

### GetRecurrenceIdOk

`func (o *RoutesEtaFiltersModel) GetRecurrenceIdOk() (*int32, bool)`

GetRecurrenceIdOk returns a tuple with the RecurrenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceId

`func (o *RoutesEtaFiltersModel) SetRecurrenceId(v int32)`

SetRecurrenceId sets RecurrenceId field to given value.

### HasRecurrenceId

`func (o *RoutesEtaFiltersModel) HasRecurrenceId() bool`

HasRecurrenceId returns a boolean if a field has been set.

### GetTagNames

`func (o *RoutesEtaFiltersModel) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *RoutesEtaFiltersModel) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *RoutesEtaFiltersModel) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.

### HasTagNames

`func (o *RoutesEtaFiltersModel) HasTagNames() bool`

HasTagNames returns a boolean if a field has been set.

### GetZoneNames

`func (o *RoutesEtaFiltersModel) GetZoneNames() []string`

GetZoneNames returns the ZoneNames field if non-nil, zero value otherwise.

### GetZoneNamesOk

`func (o *RoutesEtaFiltersModel) GetZoneNamesOk() (*[]string, bool)`

GetZoneNamesOk returns a tuple with the ZoneNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneNames

`func (o *RoutesEtaFiltersModel) SetZoneNames(v []string)`

SetZoneNames sets ZoneNames field to given value.

### HasZoneNames

`func (o *RoutesEtaFiltersModel) HasZoneNames() bool`

HasZoneNames returns a boolean if a field has been set.

### GetOptimized

`func (o *RoutesEtaFiltersModel) GetOptimized() []bool`

GetOptimized returns the Optimized field if non-nil, zero value otherwise.

### GetOptimizedOk

`func (o *RoutesEtaFiltersModel) GetOptimizedOk() (*[]bool, bool)`

GetOptimizedOk returns a tuple with the Optimized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimized

`func (o *RoutesEtaFiltersModel) SetOptimized(v []bool)`

SetOptimized sets Optimized field to given value.

### HasOptimized

`func (o *RoutesEtaFiltersModel) HasOptimized() bool`

HasOptimized returns a boolean if a field has been set.

### GetBlocked

`func (o *RoutesEtaFiltersModel) GetBlocked() []bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *RoutesEtaFiltersModel) GetBlockedOk() (*[]bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *RoutesEtaFiltersModel) SetBlocked(v []bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *RoutesEtaFiltersModel) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


