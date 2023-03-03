# RouteFiltersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **[]int32** | Unique Identifier | [optional] 
**CoDriverIds** | Pointer to **[]int32** | Unique Identifier | [optional] 
**Nr** | Pointer to **[]string** | Route nr | [optional] 
**DateTimeFrom** | Pointer to **time.Time** |  | [optional] 
**DateTimeTo** | Pointer to **time.Time** |  | [optional] 
**EarliestDateTimeSince** | Pointer to **time.Time** | filter routes with an Earliest DateTime To since this input | [optional] 
**EarliestDateTimeTill** | Pointer to **time.Time** | filter routes with an Earliest DateTime To till this input | [optional] 
**LatestDateTimeSince** | Pointer to **time.Time** | filter routes with an Latest DateTime To since this input | [optional] 
**LatestDateTimeTill** | Pointer to **time.Time** | filter routes with an Latest DateTime To till this input | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAtSince** | Pointer to **time.Time** | filter routes with an updated at date-time since this input | [optional] 
**UpdatedAtTill** | Pointer to **time.Time** | filter routes with an updated at date-time till this input | [optional] 
**Active** | Pointer to **[]int32** | Active status of route, 0 values represent deleted routes | [optional] 
**Status** | Pointer to **[]string** | Route Status | [optional] 
**StatusId** | Pointer to **[]int32** |  | [optional] 
**DriverId** | Pointer to **[]int32** |  | [optional] 
**RecurrenceId** | Pointer to **int32** | Recurrence ID | [optional] 
**TagNames** | Pointer to **[]string** | Tag names | [optional] 
**ZoneNames** | Pointer to **[]string** | Zone names | [optional] 
**Optimized** | Pointer to **[]bool** | Optimized status of Route. | [optional] 
**Blocked** | Pointer to **[]bool** | Blocked status of Route | [optional] 
**NrOfStops** | Pointer to **[]int32** | Number of stops | [optional] 
**SearchText** | Pointer to **string** | free search through text and numeric type columns | [optional] 

## Methods

### NewRouteFiltersModel

`func NewRouteFiltersModel() *RouteFiltersModel`

NewRouteFiltersModel instantiates a new RouteFiltersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteFiltersModelWithDefaults

`func NewRouteFiltersModelWithDefaults() *RouteFiltersModel`

NewRouteFiltersModelWithDefaults instantiates a new RouteFiltersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RouteFiltersModel) GetId() []int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RouteFiltersModel) GetIdOk() (*[]int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RouteFiltersModel) SetId(v []int32)`

SetId sets Id field to given value.

### HasId

`func (o *RouteFiltersModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCoDriverIds

`func (o *RouteFiltersModel) GetCoDriverIds() []int32`

GetCoDriverIds returns the CoDriverIds field if non-nil, zero value otherwise.

### GetCoDriverIdsOk

`func (o *RouteFiltersModel) GetCoDriverIdsOk() (*[]int32, bool)`

GetCoDriverIdsOk returns a tuple with the CoDriverIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoDriverIds

`func (o *RouteFiltersModel) SetCoDriverIds(v []int32)`

SetCoDriverIds sets CoDriverIds field to given value.

### HasCoDriverIds

`func (o *RouteFiltersModel) HasCoDriverIds() bool`

HasCoDriverIds returns a boolean if a field has been set.

### GetNr

`func (o *RouteFiltersModel) GetNr() []string`

GetNr returns the Nr field if non-nil, zero value otherwise.

### GetNrOk

`func (o *RouteFiltersModel) GetNrOk() (*[]string, bool)`

GetNrOk returns a tuple with the Nr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNr

`func (o *RouteFiltersModel) SetNr(v []string)`

SetNr sets Nr field to given value.

### HasNr

`func (o *RouteFiltersModel) HasNr() bool`

HasNr returns a boolean if a field has been set.

### GetDateTimeFrom

`func (o *RouteFiltersModel) GetDateTimeFrom() time.Time`

GetDateTimeFrom returns the DateTimeFrom field if non-nil, zero value otherwise.

### GetDateTimeFromOk

`func (o *RouteFiltersModel) GetDateTimeFromOk() (*time.Time, bool)`

GetDateTimeFromOk returns a tuple with the DateTimeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFrom

`func (o *RouteFiltersModel) SetDateTimeFrom(v time.Time)`

SetDateTimeFrom sets DateTimeFrom field to given value.

### HasDateTimeFrom

`func (o *RouteFiltersModel) HasDateTimeFrom() bool`

HasDateTimeFrom returns a boolean if a field has been set.

### GetDateTimeTo

`func (o *RouteFiltersModel) GetDateTimeTo() time.Time`

GetDateTimeTo returns the DateTimeTo field if non-nil, zero value otherwise.

### GetDateTimeToOk

`func (o *RouteFiltersModel) GetDateTimeToOk() (*time.Time, bool)`

GetDateTimeToOk returns a tuple with the DateTimeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeTo

`func (o *RouteFiltersModel) SetDateTimeTo(v time.Time)`

SetDateTimeTo sets DateTimeTo field to given value.

### HasDateTimeTo

`func (o *RouteFiltersModel) HasDateTimeTo() bool`

HasDateTimeTo returns a boolean if a field has been set.

### GetEarliestDateTimeSince

`func (o *RouteFiltersModel) GetEarliestDateTimeSince() time.Time`

GetEarliestDateTimeSince returns the EarliestDateTimeSince field if non-nil, zero value otherwise.

### GetEarliestDateTimeSinceOk

`func (o *RouteFiltersModel) GetEarliestDateTimeSinceOk() (*time.Time, bool)`

GetEarliestDateTimeSinceOk returns a tuple with the EarliestDateTimeSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestDateTimeSince

`func (o *RouteFiltersModel) SetEarliestDateTimeSince(v time.Time)`

SetEarliestDateTimeSince sets EarliestDateTimeSince field to given value.

### HasEarliestDateTimeSince

`func (o *RouteFiltersModel) HasEarliestDateTimeSince() bool`

HasEarliestDateTimeSince returns a boolean if a field has been set.

### GetEarliestDateTimeTill

`func (o *RouteFiltersModel) GetEarliestDateTimeTill() time.Time`

GetEarliestDateTimeTill returns the EarliestDateTimeTill field if non-nil, zero value otherwise.

### GetEarliestDateTimeTillOk

`func (o *RouteFiltersModel) GetEarliestDateTimeTillOk() (*time.Time, bool)`

GetEarliestDateTimeTillOk returns a tuple with the EarliestDateTimeTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestDateTimeTill

`func (o *RouteFiltersModel) SetEarliestDateTimeTill(v time.Time)`

SetEarliestDateTimeTill sets EarliestDateTimeTill field to given value.

### HasEarliestDateTimeTill

`func (o *RouteFiltersModel) HasEarliestDateTimeTill() bool`

HasEarliestDateTimeTill returns a boolean if a field has been set.

### GetLatestDateTimeSince

`func (o *RouteFiltersModel) GetLatestDateTimeSince() time.Time`

GetLatestDateTimeSince returns the LatestDateTimeSince field if non-nil, zero value otherwise.

### GetLatestDateTimeSinceOk

`func (o *RouteFiltersModel) GetLatestDateTimeSinceOk() (*time.Time, bool)`

GetLatestDateTimeSinceOk returns a tuple with the LatestDateTimeSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDateTimeSince

`func (o *RouteFiltersModel) SetLatestDateTimeSince(v time.Time)`

SetLatestDateTimeSince sets LatestDateTimeSince field to given value.

### HasLatestDateTimeSince

`func (o *RouteFiltersModel) HasLatestDateTimeSince() bool`

HasLatestDateTimeSince returns a boolean if a field has been set.

### GetLatestDateTimeTill

`func (o *RouteFiltersModel) GetLatestDateTimeTill() time.Time`

GetLatestDateTimeTill returns the LatestDateTimeTill field if non-nil, zero value otherwise.

### GetLatestDateTimeTillOk

`func (o *RouteFiltersModel) GetLatestDateTimeTillOk() (*time.Time, bool)`

GetLatestDateTimeTillOk returns a tuple with the LatestDateTimeTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDateTimeTill

`func (o *RouteFiltersModel) SetLatestDateTimeTill(v time.Time)`

SetLatestDateTimeTill sets LatestDateTimeTill field to given value.

### HasLatestDateTimeTill

`func (o *RouteFiltersModel) HasLatestDateTimeTill() bool`

HasLatestDateTimeTill returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RouteFiltersModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RouteFiltersModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RouteFiltersModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RouteFiltersModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedAtSince

`func (o *RouteFiltersModel) GetUpdatedAtSince() time.Time`

GetUpdatedAtSince returns the UpdatedAtSince field if non-nil, zero value otherwise.

### GetUpdatedAtSinceOk

`func (o *RouteFiltersModel) GetUpdatedAtSinceOk() (*time.Time, bool)`

GetUpdatedAtSinceOk returns a tuple with the UpdatedAtSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtSince

`func (o *RouteFiltersModel) SetUpdatedAtSince(v time.Time)`

SetUpdatedAtSince sets UpdatedAtSince field to given value.

### HasUpdatedAtSince

`func (o *RouteFiltersModel) HasUpdatedAtSince() bool`

HasUpdatedAtSince returns a boolean if a field has been set.

### GetUpdatedAtTill

`func (o *RouteFiltersModel) GetUpdatedAtTill() time.Time`

GetUpdatedAtTill returns the UpdatedAtTill field if non-nil, zero value otherwise.

### GetUpdatedAtTillOk

`func (o *RouteFiltersModel) GetUpdatedAtTillOk() (*time.Time, bool)`

GetUpdatedAtTillOk returns a tuple with the UpdatedAtTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtTill

`func (o *RouteFiltersModel) SetUpdatedAtTill(v time.Time)`

SetUpdatedAtTill sets UpdatedAtTill field to given value.

### HasUpdatedAtTill

`func (o *RouteFiltersModel) HasUpdatedAtTill() bool`

HasUpdatedAtTill returns a boolean if a field has been set.

### GetActive

`func (o *RouteFiltersModel) GetActive() []int32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RouteFiltersModel) GetActiveOk() (*[]int32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RouteFiltersModel) SetActive(v []int32)`

SetActive sets Active field to given value.

### HasActive

`func (o *RouteFiltersModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetStatus

`func (o *RouteFiltersModel) GetStatus() []string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RouteFiltersModel) GetStatusOk() (*[]string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RouteFiltersModel) SetStatus(v []string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RouteFiltersModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusId

`func (o *RouteFiltersModel) GetStatusId() []int32`

GetStatusId returns the StatusId field if non-nil, zero value otherwise.

### GetStatusIdOk

`func (o *RouteFiltersModel) GetStatusIdOk() (*[]int32, bool)`

GetStatusIdOk returns a tuple with the StatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusId

`func (o *RouteFiltersModel) SetStatusId(v []int32)`

SetStatusId sets StatusId field to given value.

### HasStatusId

`func (o *RouteFiltersModel) HasStatusId() bool`

HasStatusId returns a boolean if a field has been set.

### GetDriverId

`func (o *RouteFiltersModel) GetDriverId() []int32`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *RouteFiltersModel) GetDriverIdOk() (*[]int32, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *RouteFiltersModel) SetDriverId(v []int32)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *RouteFiltersModel) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### GetRecurrenceId

`func (o *RouteFiltersModel) GetRecurrenceId() int32`

GetRecurrenceId returns the RecurrenceId field if non-nil, zero value otherwise.

### GetRecurrenceIdOk

`func (o *RouteFiltersModel) GetRecurrenceIdOk() (*int32, bool)`

GetRecurrenceIdOk returns a tuple with the RecurrenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceId

`func (o *RouteFiltersModel) SetRecurrenceId(v int32)`

SetRecurrenceId sets RecurrenceId field to given value.

### HasRecurrenceId

`func (o *RouteFiltersModel) HasRecurrenceId() bool`

HasRecurrenceId returns a boolean if a field has been set.

### GetTagNames

`func (o *RouteFiltersModel) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *RouteFiltersModel) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *RouteFiltersModel) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.

### HasTagNames

`func (o *RouteFiltersModel) HasTagNames() bool`

HasTagNames returns a boolean if a field has been set.

### GetZoneNames

`func (o *RouteFiltersModel) GetZoneNames() []string`

GetZoneNames returns the ZoneNames field if non-nil, zero value otherwise.

### GetZoneNamesOk

`func (o *RouteFiltersModel) GetZoneNamesOk() (*[]string, bool)`

GetZoneNamesOk returns a tuple with the ZoneNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneNames

`func (o *RouteFiltersModel) SetZoneNames(v []string)`

SetZoneNames sets ZoneNames field to given value.

### HasZoneNames

`func (o *RouteFiltersModel) HasZoneNames() bool`

HasZoneNames returns a boolean if a field has been set.

### GetOptimized

`func (o *RouteFiltersModel) GetOptimized() []bool`

GetOptimized returns the Optimized field if non-nil, zero value otherwise.

### GetOptimizedOk

`func (o *RouteFiltersModel) GetOptimizedOk() (*[]bool, bool)`

GetOptimizedOk returns a tuple with the Optimized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimized

`func (o *RouteFiltersModel) SetOptimized(v []bool)`

SetOptimized sets Optimized field to given value.

### HasOptimized

`func (o *RouteFiltersModel) HasOptimized() bool`

HasOptimized returns a boolean if a field has been set.

### GetBlocked

`func (o *RouteFiltersModel) GetBlocked() []bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *RouteFiltersModel) GetBlockedOk() (*[]bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *RouteFiltersModel) SetBlocked(v []bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *RouteFiltersModel) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetNrOfStops

`func (o *RouteFiltersModel) GetNrOfStops() []int32`

GetNrOfStops returns the NrOfStops field if non-nil, zero value otherwise.

### GetNrOfStopsOk

`func (o *RouteFiltersModel) GetNrOfStopsOk() (*[]int32, bool)`

GetNrOfStopsOk returns a tuple with the NrOfStops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrOfStops

`func (o *RouteFiltersModel) SetNrOfStops(v []int32)`

SetNrOfStops sets NrOfStops field to given value.

### HasNrOfStops

`func (o *RouteFiltersModel) HasNrOfStops() bool`

HasNrOfStops returns a boolean if a field has been set.

### GetSearchText

`func (o *RouteFiltersModel) GetSearchText() string`

GetSearchText returns the SearchText field if non-nil, zero value otherwise.

### GetSearchTextOk

`func (o *RouteFiltersModel) GetSearchTextOk() (*string, bool)`

GetSearchTextOk returns a tuple with the SearchText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchText

`func (o *RouteFiltersModel) SetSearchText(v string)`

SetSearchText sets SearchText field to given value.

### HasSearchText

`func (o *RouteFiltersModel) HasSearchText() bool`

HasSearchText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


