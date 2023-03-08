# ActivityFiltersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **[]int32** | Unique Identifier(s) | [optional] 
**Nr** | Pointer to **[]string** | Activity nr | [optional] 
**AssignmentId** | Pointer to **[]int32** | Identifier(s) assignment(s) for activities | [optional] 
**ShipmentActivityId** | Pointer to **[]int32** | Unique Identifier(s) partner shipment activity | [optional] 
**EarliestDeliveryDateTimeSince** | Pointer to **time.Time** | Show activities with a earliest_delivery_date_time younger than filter value | [optional] 
**EarliestDeliveryDateTimeTill** | Pointer to **time.Time** | Show activities with a earliest_delivery_date_time older than filter value | [optional] 
**LatestDeliveryDateTimeSince** | Pointer to **time.Time** | Show activities with a latest_delivery_date_time younger than filter value | [optional] 
**LatestDeliveryDateTimeTill** | Pointer to **time.Time** | Show activities with a latest_delivery_date_time older than filter value | [optional] 
**EarliestDeliveryDateSince** | Pointer to **string** | Show activities with a earliest_delivery_date younger than filter value | [optional] 
**EarliestDeliveryDateTill** | Pointer to **string** | Show activities with a earliest_delivery_date older than filter value | [optional] 
**LatestDeliveryDateSince** | Pointer to **string** | Show activities with a latest_delivery_date younger than filter value | [optional] 
**LatestDeliveryDateTill** | Pointer to **string** | Show activities with a latest_delivery_date older than filter value | [optional] 
**DateTimeFrom** | Pointer to **time.Time** | DateTime From | [optional] 
**DateTimeTo** | Pointer to **time.Time** | DateTime To | [optional] 
**DateTimeFromSince** | Pointer to **time.Time** | filter activities with a DateTime From since this input | [optional] 
**DateTimeFromTill** | Pointer to **time.Time** | filter activities with a DateTime From till this input | [optional] 
**DateTimeToSince** | Pointer to **time.Time** | filter activities with a DateTime To since this input | [optional] 
**DateTimeToTill** | Pointer to **time.Time** | filter activities with a DateTime To till this input | [optional] 
**PlannedDateTimeFrom** | Pointer to **time.Time** | Planned DateTime From | [optional] 
**PlannedDateTimeTo** | Pointer to **time.Time** | Planned DateTime To | [optional] 
**PlannedDateTimeFromSince** | Pointer to **time.Time** | filter activities with a planned DateTime From since this input | [optional] 
**PlannedDateTimeFromTill** | Pointer to **time.Time** | filter activities with a planned DateTime From till this input | [optional] 
**PlannedDateTimeToSince** | Pointer to **time.Time** | filter activities with a planned DateTime To since this input | [optional] 
**PlannedDateTimeToTill** | Pointer to **time.Time** | filter activities with a planned DateTime To till this input | [optional] 
**UpdatedAtSince** | Pointer to **time.Time** | Show updated since | [optional] 
**UpdatedAtTill** | Pointer to **time.Time** | Show updated till | [optional] 
**Status** | Pointer to **[]string** | Activity Status | [optional] 
**StatusName** | Pointer to **[]string** | Activity Status name | [optional] 
**Active** | Pointer to **[]int32** | Active status of Activity, 0 values represent deleted activities | [optional] 
**StatusId** | Pointer to **[]int32** | Activity Status id | [optional] 
**ActivityTypeId** | Pointer to **[]int32** | Activity type id | [optional] 
**Locked** | Pointer to **[]int32** | Activity locked status. 0: not locked, 1: locked on route and time, 2: only locked on route | [optional] 
**ActivityTypeName** | Pointer to **[]string** | Activity type name | [optional] 
**BundleActivityId** | Pointer to **[]int32** | bundle activity id(s) | [optional] 
**Description** | Pointer to **[]string** | Activity description (not visible in Bumbal interface) | [optional] 
**Reference** | Pointer to **[]string** | Activity reference | [optional] 
**Priority** | Pointer to **[]int32** | Priority. 1: low, 2: medium, 3: high | [optional] 
**TopPriority** | Pointer to **bool** | top priority | [optional] 
**HasGeolocation** | Pointer to **bool** | has a properly geocoded address | [optional] 
**HasRouteAssigned** | Pointer to **bool** | has a route assigned | [optional] 
**SearchText** | Pointer to **string** | free search through text and numeric type columns | [optional] 
**RouteId** | Pointer to **[]int32** | Route id | [optional] 
**UserId** | Pointer to **int32** | ID of the user who will execute this activity | [optional] 
**PartyId** | Pointer to **int32** | Party ID | [optional] 
**RecurrenceId** | Pointer to **int32** | Recurrence ID | [optional] 
**DepotAddressLinks** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**TagNames** | Pointer to **[]string** | Tag names | [optional] 
**ZoneNames** | Pointer to **[]string** | Zone names | [optional] 
**Link** | Pointer to **[]map[string]interface{}** | Activity Link ids | [optional] 
**Links** | Pointer to **[]map[string]interface{}** | Activity Link ids | [optional] 
**CreatedAtSince** | Pointer to **time.Time** | Filter by activity created at since | [optional] 
**CreatedAtTill** | Pointer to **time.Time** | Filter by activity created at till | [optional] 
**TimeSlotsCreatedAtSince** | Pointer to **time.Time** | Filter by time_slots created at since | [optional] 
**TimeSlotsCreatedAtTill** | Pointer to **time.Time** | Filter by time_slots created at till | [optional] 
**DepotAddressId** | Pointer to **[]int32** | Depot address ID(s) | [optional] 
**AddressId** | Pointer to **[]int32** | address ID(s) (from stored masterdata) | [optional] 

## Methods

### NewActivityFiltersModel

`func NewActivityFiltersModel() *ActivityFiltersModel`

NewActivityFiltersModel instantiates a new ActivityFiltersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityFiltersModelWithDefaults

`func NewActivityFiltersModelWithDefaults() *ActivityFiltersModel`

NewActivityFiltersModelWithDefaults instantiates a new ActivityFiltersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActivityFiltersModel) GetId() []int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivityFiltersModel) GetIdOk() (*[]int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivityFiltersModel) SetId(v []int32)`

SetId sets Id field to given value.

### HasId

`func (o *ActivityFiltersModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNr

`func (o *ActivityFiltersModel) GetNr() []string`

GetNr returns the Nr field if non-nil, zero value otherwise.

### GetNrOk

`func (o *ActivityFiltersModel) GetNrOk() (*[]string, bool)`

GetNrOk returns a tuple with the Nr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNr

`func (o *ActivityFiltersModel) SetNr(v []string)`

SetNr sets Nr field to given value.

### HasNr

`func (o *ActivityFiltersModel) HasNr() bool`

HasNr returns a boolean if a field has been set.

### GetAssignmentId

`func (o *ActivityFiltersModel) GetAssignmentId() []int32`

GetAssignmentId returns the AssignmentId field if non-nil, zero value otherwise.

### GetAssignmentIdOk

`func (o *ActivityFiltersModel) GetAssignmentIdOk() (*[]int32, bool)`

GetAssignmentIdOk returns a tuple with the AssignmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentId

`func (o *ActivityFiltersModel) SetAssignmentId(v []int32)`

SetAssignmentId sets AssignmentId field to given value.

### HasAssignmentId

`func (o *ActivityFiltersModel) HasAssignmentId() bool`

HasAssignmentId returns a boolean if a field has been set.

### GetShipmentActivityId

`func (o *ActivityFiltersModel) GetShipmentActivityId() []int32`

GetShipmentActivityId returns the ShipmentActivityId field if non-nil, zero value otherwise.

### GetShipmentActivityIdOk

`func (o *ActivityFiltersModel) GetShipmentActivityIdOk() (*[]int32, bool)`

GetShipmentActivityIdOk returns a tuple with the ShipmentActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentActivityId

`func (o *ActivityFiltersModel) SetShipmentActivityId(v []int32)`

SetShipmentActivityId sets ShipmentActivityId field to given value.

### HasShipmentActivityId

`func (o *ActivityFiltersModel) HasShipmentActivityId() bool`

HasShipmentActivityId returns a boolean if a field has been set.

### GetEarliestDeliveryDateTimeSince

`func (o *ActivityFiltersModel) GetEarliestDeliveryDateTimeSince() time.Time`

GetEarliestDeliveryDateTimeSince returns the EarliestDeliveryDateTimeSince field if non-nil, zero value otherwise.

### GetEarliestDeliveryDateTimeSinceOk

`func (o *ActivityFiltersModel) GetEarliestDeliveryDateTimeSinceOk() (*time.Time, bool)`

GetEarliestDeliveryDateTimeSinceOk returns a tuple with the EarliestDeliveryDateTimeSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestDeliveryDateTimeSince

`func (o *ActivityFiltersModel) SetEarliestDeliveryDateTimeSince(v time.Time)`

SetEarliestDeliveryDateTimeSince sets EarliestDeliveryDateTimeSince field to given value.

### HasEarliestDeliveryDateTimeSince

`func (o *ActivityFiltersModel) HasEarliestDeliveryDateTimeSince() bool`

HasEarliestDeliveryDateTimeSince returns a boolean if a field has been set.

### GetEarliestDeliveryDateTimeTill

`func (o *ActivityFiltersModel) GetEarliestDeliveryDateTimeTill() time.Time`

GetEarliestDeliveryDateTimeTill returns the EarliestDeliveryDateTimeTill field if non-nil, zero value otherwise.

### GetEarliestDeliveryDateTimeTillOk

`func (o *ActivityFiltersModel) GetEarliestDeliveryDateTimeTillOk() (*time.Time, bool)`

GetEarliestDeliveryDateTimeTillOk returns a tuple with the EarliestDeliveryDateTimeTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestDeliveryDateTimeTill

`func (o *ActivityFiltersModel) SetEarliestDeliveryDateTimeTill(v time.Time)`

SetEarliestDeliveryDateTimeTill sets EarliestDeliveryDateTimeTill field to given value.

### HasEarliestDeliveryDateTimeTill

`func (o *ActivityFiltersModel) HasEarliestDeliveryDateTimeTill() bool`

HasEarliestDeliveryDateTimeTill returns a boolean if a field has been set.

### GetLatestDeliveryDateTimeSince

`func (o *ActivityFiltersModel) GetLatestDeliveryDateTimeSince() time.Time`

GetLatestDeliveryDateTimeSince returns the LatestDeliveryDateTimeSince field if non-nil, zero value otherwise.

### GetLatestDeliveryDateTimeSinceOk

`func (o *ActivityFiltersModel) GetLatestDeliveryDateTimeSinceOk() (*time.Time, bool)`

GetLatestDeliveryDateTimeSinceOk returns a tuple with the LatestDeliveryDateTimeSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDeliveryDateTimeSince

`func (o *ActivityFiltersModel) SetLatestDeliveryDateTimeSince(v time.Time)`

SetLatestDeliveryDateTimeSince sets LatestDeliveryDateTimeSince field to given value.

### HasLatestDeliveryDateTimeSince

`func (o *ActivityFiltersModel) HasLatestDeliveryDateTimeSince() bool`

HasLatestDeliveryDateTimeSince returns a boolean if a field has been set.

### GetLatestDeliveryDateTimeTill

`func (o *ActivityFiltersModel) GetLatestDeliveryDateTimeTill() time.Time`

GetLatestDeliveryDateTimeTill returns the LatestDeliveryDateTimeTill field if non-nil, zero value otherwise.

### GetLatestDeliveryDateTimeTillOk

`func (o *ActivityFiltersModel) GetLatestDeliveryDateTimeTillOk() (*time.Time, bool)`

GetLatestDeliveryDateTimeTillOk returns a tuple with the LatestDeliveryDateTimeTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDeliveryDateTimeTill

`func (o *ActivityFiltersModel) SetLatestDeliveryDateTimeTill(v time.Time)`

SetLatestDeliveryDateTimeTill sets LatestDeliveryDateTimeTill field to given value.

### HasLatestDeliveryDateTimeTill

`func (o *ActivityFiltersModel) HasLatestDeliveryDateTimeTill() bool`

HasLatestDeliveryDateTimeTill returns a boolean if a field has been set.

### GetEarliestDeliveryDateSince

`func (o *ActivityFiltersModel) GetEarliestDeliveryDateSince() string`

GetEarliestDeliveryDateSince returns the EarliestDeliveryDateSince field if non-nil, zero value otherwise.

### GetEarliestDeliveryDateSinceOk

`func (o *ActivityFiltersModel) GetEarliestDeliveryDateSinceOk() (*string, bool)`

GetEarliestDeliveryDateSinceOk returns a tuple with the EarliestDeliveryDateSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestDeliveryDateSince

`func (o *ActivityFiltersModel) SetEarliestDeliveryDateSince(v string)`

SetEarliestDeliveryDateSince sets EarliestDeliveryDateSince field to given value.

### HasEarliestDeliveryDateSince

`func (o *ActivityFiltersModel) HasEarliestDeliveryDateSince() bool`

HasEarliestDeliveryDateSince returns a boolean if a field has been set.

### GetEarliestDeliveryDateTill

`func (o *ActivityFiltersModel) GetEarliestDeliveryDateTill() string`

GetEarliestDeliveryDateTill returns the EarliestDeliveryDateTill field if non-nil, zero value otherwise.

### GetEarliestDeliveryDateTillOk

`func (o *ActivityFiltersModel) GetEarliestDeliveryDateTillOk() (*string, bool)`

GetEarliestDeliveryDateTillOk returns a tuple with the EarliestDeliveryDateTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestDeliveryDateTill

`func (o *ActivityFiltersModel) SetEarliestDeliveryDateTill(v string)`

SetEarliestDeliveryDateTill sets EarliestDeliveryDateTill field to given value.

### HasEarliestDeliveryDateTill

`func (o *ActivityFiltersModel) HasEarliestDeliveryDateTill() bool`

HasEarliestDeliveryDateTill returns a boolean if a field has been set.

### GetLatestDeliveryDateSince

`func (o *ActivityFiltersModel) GetLatestDeliveryDateSince() string`

GetLatestDeliveryDateSince returns the LatestDeliveryDateSince field if non-nil, zero value otherwise.

### GetLatestDeliveryDateSinceOk

`func (o *ActivityFiltersModel) GetLatestDeliveryDateSinceOk() (*string, bool)`

GetLatestDeliveryDateSinceOk returns a tuple with the LatestDeliveryDateSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDeliveryDateSince

`func (o *ActivityFiltersModel) SetLatestDeliveryDateSince(v string)`

SetLatestDeliveryDateSince sets LatestDeliveryDateSince field to given value.

### HasLatestDeliveryDateSince

`func (o *ActivityFiltersModel) HasLatestDeliveryDateSince() bool`

HasLatestDeliveryDateSince returns a boolean if a field has been set.

### GetLatestDeliveryDateTill

`func (o *ActivityFiltersModel) GetLatestDeliveryDateTill() string`

GetLatestDeliveryDateTill returns the LatestDeliveryDateTill field if non-nil, zero value otherwise.

### GetLatestDeliveryDateTillOk

`func (o *ActivityFiltersModel) GetLatestDeliveryDateTillOk() (*string, bool)`

GetLatestDeliveryDateTillOk returns a tuple with the LatestDeliveryDateTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDeliveryDateTill

`func (o *ActivityFiltersModel) SetLatestDeliveryDateTill(v string)`

SetLatestDeliveryDateTill sets LatestDeliveryDateTill field to given value.

### HasLatestDeliveryDateTill

`func (o *ActivityFiltersModel) HasLatestDeliveryDateTill() bool`

HasLatestDeliveryDateTill returns a boolean if a field has been set.

### GetDateTimeFrom

`func (o *ActivityFiltersModel) GetDateTimeFrom() time.Time`

GetDateTimeFrom returns the DateTimeFrom field if non-nil, zero value otherwise.

### GetDateTimeFromOk

`func (o *ActivityFiltersModel) GetDateTimeFromOk() (*time.Time, bool)`

GetDateTimeFromOk returns a tuple with the DateTimeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFrom

`func (o *ActivityFiltersModel) SetDateTimeFrom(v time.Time)`

SetDateTimeFrom sets DateTimeFrom field to given value.

### HasDateTimeFrom

`func (o *ActivityFiltersModel) HasDateTimeFrom() bool`

HasDateTimeFrom returns a boolean if a field has been set.

### GetDateTimeTo

`func (o *ActivityFiltersModel) GetDateTimeTo() time.Time`

GetDateTimeTo returns the DateTimeTo field if non-nil, zero value otherwise.

### GetDateTimeToOk

`func (o *ActivityFiltersModel) GetDateTimeToOk() (*time.Time, bool)`

GetDateTimeToOk returns a tuple with the DateTimeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeTo

`func (o *ActivityFiltersModel) SetDateTimeTo(v time.Time)`

SetDateTimeTo sets DateTimeTo field to given value.

### HasDateTimeTo

`func (o *ActivityFiltersModel) HasDateTimeTo() bool`

HasDateTimeTo returns a boolean if a field has been set.

### GetDateTimeFromSince

`func (o *ActivityFiltersModel) GetDateTimeFromSince() time.Time`

GetDateTimeFromSince returns the DateTimeFromSince field if non-nil, zero value otherwise.

### GetDateTimeFromSinceOk

`func (o *ActivityFiltersModel) GetDateTimeFromSinceOk() (*time.Time, bool)`

GetDateTimeFromSinceOk returns a tuple with the DateTimeFromSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFromSince

`func (o *ActivityFiltersModel) SetDateTimeFromSince(v time.Time)`

SetDateTimeFromSince sets DateTimeFromSince field to given value.

### HasDateTimeFromSince

`func (o *ActivityFiltersModel) HasDateTimeFromSince() bool`

HasDateTimeFromSince returns a boolean if a field has been set.

### GetDateTimeFromTill

`func (o *ActivityFiltersModel) GetDateTimeFromTill() time.Time`

GetDateTimeFromTill returns the DateTimeFromTill field if non-nil, zero value otherwise.

### GetDateTimeFromTillOk

`func (o *ActivityFiltersModel) GetDateTimeFromTillOk() (*time.Time, bool)`

GetDateTimeFromTillOk returns a tuple with the DateTimeFromTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFromTill

`func (o *ActivityFiltersModel) SetDateTimeFromTill(v time.Time)`

SetDateTimeFromTill sets DateTimeFromTill field to given value.

### HasDateTimeFromTill

`func (o *ActivityFiltersModel) HasDateTimeFromTill() bool`

HasDateTimeFromTill returns a boolean if a field has been set.

### GetDateTimeToSince

`func (o *ActivityFiltersModel) GetDateTimeToSince() time.Time`

GetDateTimeToSince returns the DateTimeToSince field if non-nil, zero value otherwise.

### GetDateTimeToSinceOk

`func (o *ActivityFiltersModel) GetDateTimeToSinceOk() (*time.Time, bool)`

GetDateTimeToSinceOk returns a tuple with the DateTimeToSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeToSince

`func (o *ActivityFiltersModel) SetDateTimeToSince(v time.Time)`

SetDateTimeToSince sets DateTimeToSince field to given value.

### HasDateTimeToSince

`func (o *ActivityFiltersModel) HasDateTimeToSince() bool`

HasDateTimeToSince returns a boolean if a field has been set.

### GetDateTimeToTill

`func (o *ActivityFiltersModel) GetDateTimeToTill() time.Time`

GetDateTimeToTill returns the DateTimeToTill field if non-nil, zero value otherwise.

### GetDateTimeToTillOk

`func (o *ActivityFiltersModel) GetDateTimeToTillOk() (*time.Time, bool)`

GetDateTimeToTillOk returns a tuple with the DateTimeToTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeToTill

`func (o *ActivityFiltersModel) SetDateTimeToTill(v time.Time)`

SetDateTimeToTill sets DateTimeToTill field to given value.

### HasDateTimeToTill

`func (o *ActivityFiltersModel) HasDateTimeToTill() bool`

HasDateTimeToTill returns a boolean if a field has been set.

### GetPlannedDateTimeFrom

`func (o *ActivityFiltersModel) GetPlannedDateTimeFrom() time.Time`

GetPlannedDateTimeFrom returns the PlannedDateTimeFrom field if non-nil, zero value otherwise.

### GetPlannedDateTimeFromOk

`func (o *ActivityFiltersModel) GetPlannedDateTimeFromOk() (*time.Time, bool)`

GetPlannedDateTimeFromOk returns a tuple with the PlannedDateTimeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDateTimeFrom

`func (o *ActivityFiltersModel) SetPlannedDateTimeFrom(v time.Time)`

SetPlannedDateTimeFrom sets PlannedDateTimeFrom field to given value.

### HasPlannedDateTimeFrom

`func (o *ActivityFiltersModel) HasPlannedDateTimeFrom() bool`

HasPlannedDateTimeFrom returns a boolean if a field has been set.

### GetPlannedDateTimeTo

`func (o *ActivityFiltersModel) GetPlannedDateTimeTo() time.Time`

GetPlannedDateTimeTo returns the PlannedDateTimeTo field if non-nil, zero value otherwise.

### GetPlannedDateTimeToOk

`func (o *ActivityFiltersModel) GetPlannedDateTimeToOk() (*time.Time, bool)`

GetPlannedDateTimeToOk returns a tuple with the PlannedDateTimeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDateTimeTo

`func (o *ActivityFiltersModel) SetPlannedDateTimeTo(v time.Time)`

SetPlannedDateTimeTo sets PlannedDateTimeTo field to given value.

### HasPlannedDateTimeTo

`func (o *ActivityFiltersModel) HasPlannedDateTimeTo() bool`

HasPlannedDateTimeTo returns a boolean if a field has been set.

### GetPlannedDateTimeFromSince

`func (o *ActivityFiltersModel) GetPlannedDateTimeFromSince() time.Time`

GetPlannedDateTimeFromSince returns the PlannedDateTimeFromSince field if non-nil, zero value otherwise.

### GetPlannedDateTimeFromSinceOk

`func (o *ActivityFiltersModel) GetPlannedDateTimeFromSinceOk() (*time.Time, bool)`

GetPlannedDateTimeFromSinceOk returns a tuple with the PlannedDateTimeFromSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDateTimeFromSince

`func (o *ActivityFiltersModel) SetPlannedDateTimeFromSince(v time.Time)`

SetPlannedDateTimeFromSince sets PlannedDateTimeFromSince field to given value.

### HasPlannedDateTimeFromSince

`func (o *ActivityFiltersModel) HasPlannedDateTimeFromSince() bool`

HasPlannedDateTimeFromSince returns a boolean if a field has been set.

### GetPlannedDateTimeFromTill

`func (o *ActivityFiltersModel) GetPlannedDateTimeFromTill() time.Time`

GetPlannedDateTimeFromTill returns the PlannedDateTimeFromTill field if non-nil, zero value otherwise.

### GetPlannedDateTimeFromTillOk

`func (o *ActivityFiltersModel) GetPlannedDateTimeFromTillOk() (*time.Time, bool)`

GetPlannedDateTimeFromTillOk returns a tuple with the PlannedDateTimeFromTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDateTimeFromTill

`func (o *ActivityFiltersModel) SetPlannedDateTimeFromTill(v time.Time)`

SetPlannedDateTimeFromTill sets PlannedDateTimeFromTill field to given value.

### HasPlannedDateTimeFromTill

`func (o *ActivityFiltersModel) HasPlannedDateTimeFromTill() bool`

HasPlannedDateTimeFromTill returns a boolean if a field has been set.

### GetPlannedDateTimeToSince

`func (o *ActivityFiltersModel) GetPlannedDateTimeToSince() time.Time`

GetPlannedDateTimeToSince returns the PlannedDateTimeToSince field if non-nil, zero value otherwise.

### GetPlannedDateTimeToSinceOk

`func (o *ActivityFiltersModel) GetPlannedDateTimeToSinceOk() (*time.Time, bool)`

GetPlannedDateTimeToSinceOk returns a tuple with the PlannedDateTimeToSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDateTimeToSince

`func (o *ActivityFiltersModel) SetPlannedDateTimeToSince(v time.Time)`

SetPlannedDateTimeToSince sets PlannedDateTimeToSince field to given value.

### HasPlannedDateTimeToSince

`func (o *ActivityFiltersModel) HasPlannedDateTimeToSince() bool`

HasPlannedDateTimeToSince returns a boolean if a field has been set.

### GetPlannedDateTimeToTill

`func (o *ActivityFiltersModel) GetPlannedDateTimeToTill() time.Time`

GetPlannedDateTimeToTill returns the PlannedDateTimeToTill field if non-nil, zero value otherwise.

### GetPlannedDateTimeToTillOk

`func (o *ActivityFiltersModel) GetPlannedDateTimeToTillOk() (*time.Time, bool)`

GetPlannedDateTimeToTillOk returns a tuple with the PlannedDateTimeToTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDateTimeToTill

`func (o *ActivityFiltersModel) SetPlannedDateTimeToTill(v time.Time)`

SetPlannedDateTimeToTill sets PlannedDateTimeToTill field to given value.

### HasPlannedDateTimeToTill

`func (o *ActivityFiltersModel) HasPlannedDateTimeToTill() bool`

HasPlannedDateTimeToTill returns a boolean if a field has been set.

### GetUpdatedAtSince

`func (o *ActivityFiltersModel) GetUpdatedAtSince() time.Time`

GetUpdatedAtSince returns the UpdatedAtSince field if non-nil, zero value otherwise.

### GetUpdatedAtSinceOk

`func (o *ActivityFiltersModel) GetUpdatedAtSinceOk() (*time.Time, bool)`

GetUpdatedAtSinceOk returns a tuple with the UpdatedAtSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtSince

`func (o *ActivityFiltersModel) SetUpdatedAtSince(v time.Time)`

SetUpdatedAtSince sets UpdatedAtSince field to given value.

### HasUpdatedAtSince

`func (o *ActivityFiltersModel) HasUpdatedAtSince() bool`

HasUpdatedAtSince returns a boolean if a field has been set.

### GetUpdatedAtTill

`func (o *ActivityFiltersModel) GetUpdatedAtTill() time.Time`

GetUpdatedAtTill returns the UpdatedAtTill field if non-nil, zero value otherwise.

### GetUpdatedAtTillOk

`func (o *ActivityFiltersModel) GetUpdatedAtTillOk() (*time.Time, bool)`

GetUpdatedAtTillOk returns a tuple with the UpdatedAtTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtTill

`func (o *ActivityFiltersModel) SetUpdatedAtTill(v time.Time)`

SetUpdatedAtTill sets UpdatedAtTill field to given value.

### HasUpdatedAtTill

`func (o *ActivityFiltersModel) HasUpdatedAtTill() bool`

HasUpdatedAtTill returns a boolean if a field has been set.

### GetStatus

`func (o *ActivityFiltersModel) GetStatus() []string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActivityFiltersModel) GetStatusOk() (*[]string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActivityFiltersModel) SetStatus(v []string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ActivityFiltersModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusName

`func (o *ActivityFiltersModel) GetStatusName() []string`

GetStatusName returns the StatusName field if non-nil, zero value otherwise.

### GetStatusNameOk

`func (o *ActivityFiltersModel) GetStatusNameOk() (*[]string, bool)`

GetStatusNameOk returns a tuple with the StatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusName

`func (o *ActivityFiltersModel) SetStatusName(v []string)`

SetStatusName sets StatusName field to given value.

### HasStatusName

`func (o *ActivityFiltersModel) HasStatusName() bool`

HasStatusName returns a boolean if a field has been set.

### GetActive

`func (o *ActivityFiltersModel) GetActive() []int32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ActivityFiltersModel) GetActiveOk() (*[]int32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ActivityFiltersModel) SetActive(v []int32)`

SetActive sets Active field to given value.

### HasActive

`func (o *ActivityFiltersModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetStatusId

`func (o *ActivityFiltersModel) GetStatusId() []int32`

GetStatusId returns the StatusId field if non-nil, zero value otherwise.

### GetStatusIdOk

`func (o *ActivityFiltersModel) GetStatusIdOk() (*[]int32, bool)`

GetStatusIdOk returns a tuple with the StatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusId

`func (o *ActivityFiltersModel) SetStatusId(v []int32)`

SetStatusId sets StatusId field to given value.

### HasStatusId

`func (o *ActivityFiltersModel) HasStatusId() bool`

HasStatusId returns a boolean if a field has been set.

### GetActivityTypeId

`func (o *ActivityFiltersModel) GetActivityTypeId() []int32`

GetActivityTypeId returns the ActivityTypeId field if non-nil, zero value otherwise.

### GetActivityTypeIdOk

`func (o *ActivityFiltersModel) GetActivityTypeIdOk() (*[]int32, bool)`

GetActivityTypeIdOk returns a tuple with the ActivityTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityTypeId

`func (o *ActivityFiltersModel) SetActivityTypeId(v []int32)`

SetActivityTypeId sets ActivityTypeId field to given value.

### HasActivityTypeId

`func (o *ActivityFiltersModel) HasActivityTypeId() bool`

HasActivityTypeId returns a boolean if a field has been set.

### GetLocked

`func (o *ActivityFiltersModel) GetLocked() []int32`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *ActivityFiltersModel) GetLockedOk() (*[]int32, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *ActivityFiltersModel) SetLocked(v []int32)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *ActivityFiltersModel) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetActivityTypeName

`func (o *ActivityFiltersModel) GetActivityTypeName() []string`

GetActivityTypeName returns the ActivityTypeName field if non-nil, zero value otherwise.

### GetActivityTypeNameOk

`func (o *ActivityFiltersModel) GetActivityTypeNameOk() (*[]string, bool)`

GetActivityTypeNameOk returns a tuple with the ActivityTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityTypeName

`func (o *ActivityFiltersModel) SetActivityTypeName(v []string)`

SetActivityTypeName sets ActivityTypeName field to given value.

### HasActivityTypeName

`func (o *ActivityFiltersModel) HasActivityTypeName() bool`

HasActivityTypeName returns a boolean if a field has been set.

### GetBundleActivityId

`func (o *ActivityFiltersModel) GetBundleActivityId() []int32`

GetBundleActivityId returns the BundleActivityId field if non-nil, zero value otherwise.

### GetBundleActivityIdOk

`func (o *ActivityFiltersModel) GetBundleActivityIdOk() (*[]int32, bool)`

GetBundleActivityIdOk returns a tuple with the BundleActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleActivityId

`func (o *ActivityFiltersModel) SetBundleActivityId(v []int32)`

SetBundleActivityId sets BundleActivityId field to given value.

### HasBundleActivityId

`func (o *ActivityFiltersModel) HasBundleActivityId() bool`

HasBundleActivityId returns a boolean if a field has been set.

### GetDescription

`func (o *ActivityFiltersModel) GetDescription() []string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ActivityFiltersModel) GetDescriptionOk() (*[]string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ActivityFiltersModel) SetDescription(v []string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ActivityFiltersModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetReference

`func (o *ActivityFiltersModel) GetReference() []string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ActivityFiltersModel) GetReferenceOk() (*[]string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ActivityFiltersModel) SetReference(v []string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ActivityFiltersModel) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetPriority

`func (o *ActivityFiltersModel) GetPriority() []int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ActivityFiltersModel) GetPriorityOk() (*[]int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ActivityFiltersModel) SetPriority(v []int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ActivityFiltersModel) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTopPriority

`func (o *ActivityFiltersModel) GetTopPriority() bool`

GetTopPriority returns the TopPriority field if non-nil, zero value otherwise.

### GetTopPriorityOk

`func (o *ActivityFiltersModel) GetTopPriorityOk() (*bool, bool)`

GetTopPriorityOk returns a tuple with the TopPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopPriority

`func (o *ActivityFiltersModel) SetTopPriority(v bool)`

SetTopPriority sets TopPriority field to given value.

### HasTopPriority

`func (o *ActivityFiltersModel) HasTopPriority() bool`

HasTopPriority returns a boolean if a field has been set.

### GetHasGeolocation

`func (o *ActivityFiltersModel) GetHasGeolocation() bool`

GetHasGeolocation returns the HasGeolocation field if non-nil, zero value otherwise.

### GetHasGeolocationOk

`func (o *ActivityFiltersModel) GetHasGeolocationOk() (*bool, bool)`

GetHasGeolocationOk returns a tuple with the HasGeolocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasGeolocation

`func (o *ActivityFiltersModel) SetHasGeolocation(v bool)`

SetHasGeolocation sets HasGeolocation field to given value.

### HasHasGeolocation

`func (o *ActivityFiltersModel) HasHasGeolocation() bool`

HasHasGeolocation returns a boolean if a field has been set.

### GetHasRouteAssigned

`func (o *ActivityFiltersModel) GetHasRouteAssigned() bool`

GetHasRouteAssigned returns the HasRouteAssigned field if non-nil, zero value otherwise.

### GetHasRouteAssignedOk

`func (o *ActivityFiltersModel) GetHasRouteAssignedOk() (*bool, bool)`

GetHasRouteAssignedOk returns a tuple with the HasRouteAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRouteAssigned

`func (o *ActivityFiltersModel) SetHasRouteAssigned(v bool)`

SetHasRouteAssigned sets HasRouteAssigned field to given value.

### HasHasRouteAssigned

`func (o *ActivityFiltersModel) HasHasRouteAssigned() bool`

HasHasRouteAssigned returns a boolean if a field has been set.

### GetSearchText

`func (o *ActivityFiltersModel) GetSearchText() string`

GetSearchText returns the SearchText field if non-nil, zero value otherwise.

### GetSearchTextOk

`func (o *ActivityFiltersModel) GetSearchTextOk() (*string, bool)`

GetSearchTextOk returns a tuple with the SearchText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchText

`func (o *ActivityFiltersModel) SetSearchText(v string)`

SetSearchText sets SearchText field to given value.

### HasSearchText

`func (o *ActivityFiltersModel) HasSearchText() bool`

HasSearchText returns a boolean if a field has been set.

### GetRouteId

`func (o *ActivityFiltersModel) GetRouteId() []int32`

GetRouteId returns the RouteId field if non-nil, zero value otherwise.

### GetRouteIdOk

`func (o *ActivityFiltersModel) GetRouteIdOk() (*[]int32, bool)`

GetRouteIdOk returns a tuple with the RouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteId

`func (o *ActivityFiltersModel) SetRouteId(v []int32)`

SetRouteId sets RouteId field to given value.

### HasRouteId

`func (o *ActivityFiltersModel) HasRouteId() bool`

HasRouteId returns a boolean if a field has been set.

### GetUserId

`func (o *ActivityFiltersModel) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ActivityFiltersModel) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ActivityFiltersModel) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ActivityFiltersModel) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetPartyId

`func (o *ActivityFiltersModel) GetPartyId() int32`

GetPartyId returns the PartyId field if non-nil, zero value otherwise.

### GetPartyIdOk

`func (o *ActivityFiltersModel) GetPartyIdOk() (*int32, bool)`

GetPartyIdOk returns a tuple with the PartyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartyId

`func (o *ActivityFiltersModel) SetPartyId(v int32)`

SetPartyId sets PartyId field to given value.

### HasPartyId

`func (o *ActivityFiltersModel) HasPartyId() bool`

HasPartyId returns a boolean if a field has been set.

### GetRecurrenceId

`func (o *ActivityFiltersModel) GetRecurrenceId() int32`

GetRecurrenceId returns the RecurrenceId field if non-nil, zero value otherwise.

### GetRecurrenceIdOk

`func (o *ActivityFiltersModel) GetRecurrenceIdOk() (*int32, bool)`

GetRecurrenceIdOk returns a tuple with the RecurrenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceId

`func (o *ActivityFiltersModel) SetRecurrenceId(v int32)`

SetRecurrenceId sets RecurrenceId field to given value.

### HasRecurrenceId

`func (o *ActivityFiltersModel) HasRecurrenceId() bool`

HasRecurrenceId returns a boolean if a field has been set.

### GetDepotAddressLinks

`func (o *ActivityFiltersModel) GetDepotAddressLinks() []LinkModel`

GetDepotAddressLinks returns the DepotAddressLinks field if non-nil, zero value otherwise.

### GetDepotAddressLinksOk

`func (o *ActivityFiltersModel) GetDepotAddressLinksOk() (*[]LinkModel, bool)`

GetDepotAddressLinksOk returns a tuple with the DepotAddressLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotAddressLinks

`func (o *ActivityFiltersModel) SetDepotAddressLinks(v []LinkModel)`

SetDepotAddressLinks sets DepotAddressLinks field to given value.

### HasDepotAddressLinks

`func (o *ActivityFiltersModel) HasDepotAddressLinks() bool`

HasDepotAddressLinks returns a boolean if a field has been set.

### GetTagNames

`func (o *ActivityFiltersModel) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *ActivityFiltersModel) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *ActivityFiltersModel) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.

### HasTagNames

`func (o *ActivityFiltersModel) HasTagNames() bool`

HasTagNames returns a boolean if a field has been set.

### GetZoneNames

`func (o *ActivityFiltersModel) GetZoneNames() []string`

GetZoneNames returns the ZoneNames field if non-nil, zero value otherwise.

### GetZoneNamesOk

`func (o *ActivityFiltersModel) GetZoneNamesOk() (*[]string, bool)`

GetZoneNamesOk returns a tuple with the ZoneNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneNames

`func (o *ActivityFiltersModel) SetZoneNames(v []string)`

SetZoneNames sets ZoneNames field to given value.

### HasZoneNames

`func (o *ActivityFiltersModel) HasZoneNames() bool`

HasZoneNames returns a boolean if a field has been set.

### GetLink

`func (o *ActivityFiltersModel) GetLink() []map[string]interface{}`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *ActivityFiltersModel) GetLinkOk() (*[]map[string]interface{}, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *ActivityFiltersModel) SetLink(v []map[string]interface{})`

SetLink sets Link field to given value.

### HasLink

`func (o *ActivityFiltersModel) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetLinks

`func (o *ActivityFiltersModel) GetLinks() []map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ActivityFiltersModel) GetLinksOk() (*[]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ActivityFiltersModel) SetLinks(v []map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ActivityFiltersModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetCreatedAtSince

`func (o *ActivityFiltersModel) GetCreatedAtSince() time.Time`

GetCreatedAtSince returns the CreatedAtSince field if non-nil, zero value otherwise.

### GetCreatedAtSinceOk

`func (o *ActivityFiltersModel) GetCreatedAtSinceOk() (*time.Time, bool)`

GetCreatedAtSinceOk returns a tuple with the CreatedAtSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtSince

`func (o *ActivityFiltersModel) SetCreatedAtSince(v time.Time)`

SetCreatedAtSince sets CreatedAtSince field to given value.

### HasCreatedAtSince

`func (o *ActivityFiltersModel) HasCreatedAtSince() bool`

HasCreatedAtSince returns a boolean if a field has been set.

### GetCreatedAtTill

`func (o *ActivityFiltersModel) GetCreatedAtTill() time.Time`

GetCreatedAtTill returns the CreatedAtTill field if non-nil, zero value otherwise.

### GetCreatedAtTillOk

`func (o *ActivityFiltersModel) GetCreatedAtTillOk() (*time.Time, bool)`

GetCreatedAtTillOk returns a tuple with the CreatedAtTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAtTill

`func (o *ActivityFiltersModel) SetCreatedAtTill(v time.Time)`

SetCreatedAtTill sets CreatedAtTill field to given value.

### HasCreatedAtTill

`func (o *ActivityFiltersModel) HasCreatedAtTill() bool`

HasCreatedAtTill returns a boolean if a field has been set.

### GetTimeSlotsCreatedAtSince

`func (o *ActivityFiltersModel) GetTimeSlotsCreatedAtSince() time.Time`

GetTimeSlotsCreatedAtSince returns the TimeSlotsCreatedAtSince field if non-nil, zero value otherwise.

### GetTimeSlotsCreatedAtSinceOk

`func (o *ActivityFiltersModel) GetTimeSlotsCreatedAtSinceOk() (*time.Time, bool)`

GetTimeSlotsCreatedAtSinceOk returns a tuple with the TimeSlotsCreatedAtSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSlotsCreatedAtSince

`func (o *ActivityFiltersModel) SetTimeSlotsCreatedAtSince(v time.Time)`

SetTimeSlotsCreatedAtSince sets TimeSlotsCreatedAtSince field to given value.

### HasTimeSlotsCreatedAtSince

`func (o *ActivityFiltersModel) HasTimeSlotsCreatedAtSince() bool`

HasTimeSlotsCreatedAtSince returns a boolean if a field has been set.

### GetTimeSlotsCreatedAtTill

`func (o *ActivityFiltersModel) GetTimeSlotsCreatedAtTill() time.Time`

GetTimeSlotsCreatedAtTill returns the TimeSlotsCreatedAtTill field if non-nil, zero value otherwise.

### GetTimeSlotsCreatedAtTillOk

`func (o *ActivityFiltersModel) GetTimeSlotsCreatedAtTillOk() (*time.Time, bool)`

GetTimeSlotsCreatedAtTillOk returns a tuple with the TimeSlotsCreatedAtTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSlotsCreatedAtTill

`func (o *ActivityFiltersModel) SetTimeSlotsCreatedAtTill(v time.Time)`

SetTimeSlotsCreatedAtTill sets TimeSlotsCreatedAtTill field to given value.

### HasTimeSlotsCreatedAtTill

`func (o *ActivityFiltersModel) HasTimeSlotsCreatedAtTill() bool`

HasTimeSlotsCreatedAtTill returns a boolean if a field has been set.

### GetDepotAddressId

`func (o *ActivityFiltersModel) GetDepotAddressId() []int32`

GetDepotAddressId returns the DepotAddressId field if non-nil, zero value otherwise.

### GetDepotAddressIdOk

`func (o *ActivityFiltersModel) GetDepotAddressIdOk() (*[]int32, bool)`

GetDepotAddressIdOk returns a tuple with the DepotAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotAddressId

`func (o *ActivityFiltersModel) SetDepotAddressId(v []int32)`

SetDepotAddressId sets DepotAddressId field to given value.

### HasDepotAddressId

`func (o *ActivityFiltersModel) HasDepotAddressId() bool`

HasDepotAddressId returns a boolean if a field has been set.

### GetAddressId

`func (o *ActivityFiltersModel) GetAddressId() []int32`

GetAddressId returns the AddressId field if non-nil, zero value otherwise.

### GetAddressIdOk

`func (o *ActivityFiltersModel) GetAddressIdOk() (*[]int32, bool)`

GetAddressIdOk returns a tuple with the AddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressId

`func (o *ActivityFiltersModel) SetAddressId(v []int32)`

SetAddressId sets AddressId field to given value.

### HasAddressId

`func (o *ActivityFiltersModel) HasAddressId() bool`

HasAddressId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


