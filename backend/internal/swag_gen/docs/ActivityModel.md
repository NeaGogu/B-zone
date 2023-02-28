# ActivityModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique Identifier | [optional] [readonly] 
**Uuid** | Pointer to **string** | unique per activity | [optional] [readonly] 
**BundleActivityId** | Pointer to **int64** | Unique Identifier for partner shipment activity | [optional] [readonly] 
**ShipmentActivityId** | Pointer to **int64** | Unique Identifier for partner shipment activity | [optional] [readonly] 
**ShipmentActivityNr** | Pointer to **string** | Number for partner shipment activity | [optional] [readonly] 
**BundleShipmentActivityNr** | Pointer to **string** | Number for partner bundle shipment activity | [optional] [readonly] 
**Nr** | Pointer to **string** | Number of this Activity | [optional] 
**ActivityTypeName** | Pointer to **string** | Activity Type Name. breakdown (:20), bundled (:32), car_end (:6), car_start (:5), car_wash (:27), combi (:29), depot (:28), driver_end (:10), driver_start (:9), dropoff (:2), gas_refill (:14), interior_wash (:30), maintenance (:26), maintenance_planned (:15), maintenance_unplanned (:16), other (:31), pause (:13), pickup (:1), route_end (:4), route_start (:3), sanitary (:21), stop (:11), unplanned_stop (:18), user_other (:17), wait (:19) | [optional] [readonly] 
**ActivityTypeId** | Pointer to **int64** | Unique Activity type ID. 1 (:pickup), 2 (:dropoff), 3 (:route_start), 4 (:route_end), 5 (:car_start), 6 (:car_end), 9 (:driver_start), 10 (:driver_end), 11 (:stop), 13 (:pause), 14 (:gas_refill), 15 (:maintenance_planned), 16 (:maintenance_unplanned), 17 (:user_other), 18 (:unplanned_stop), 19 (:wait), 20 (:breakdown), 21 (:sanitary), 26 (:maintenance), 27 (:car_wash), 28 (:depot), 29 (:combi), 30 (:interior_wash), 31 (:other), 32 (:bundled) | [optional] [readonly] 
**StatusId** | Pointer to **int32** | StatusId of this Activity, 28: activity_cancelled, 20: activity_incomplete, 21: activity_new, 39: activity_awaiting, 22: activity_accepted, 3: activity_planned, 4: activity_in_progress, 9: activity_executed | [optional] 
**StatusName** | Pointer to **string** | Activity Status | [optional] 
**AssignmentId** | Pointer to **int32** | Assignment ID | [optional] 
**AssignmentNr** | Pointer to **string** | Assignment NR | [optional] 
**Assignment** | Pointer to [**AssignmentModel**](AssignmentModel.md) |  | [optional] 
**AssignmentSequenceNr** | Pointer to **int64** | Assignment sequence number for multi day assignments | [optional] 
**Remarks** | Pointer to **string** | Remarks about this activity | [optional] 
**Locked** | Pointer to **int32** | Activity locked status. 0 &#x3D; not locked. 1 &#x3D; locked on route and time. 2 &#x3D; locked on route. | [optional] 
**Optimized** | Pointer to **bool** | Activity optimized status within route. | [optional] 
**Priority** | Pointer to **int32** | Priority level. 1 for highest priority, 3 for lowest priority. Default &#x3D; 2 | [optional] 
**TopPriority** | Pointer to **bool** | Activity has top priority | [optional] 
**SequenceNr** | Pointer to **int64** | Sequence number on Route | [optional] 
**PlannedDrivingTime** | Pointer to **int64** | Deprecated! Driving time from the activity before this one | [optional] 
**PlannedDrivingDuration** | Pointer to **int64** | Driving time from the activity before this one | [optional] 
**PlannedDrivingDistance** | Pointer to **int64** | Driving distance from the activity before this one | [optional] 
**Reference** | Pointer to **string** | Reference of this activity | [optional] 
**Description** | Pointer to **string** | description of this activity | [optional] 
**EarliestDeliveryDateTime** | Pointer to **string** |  | [optional] 
**LatestDeliveryDateTime** | Pointer to **string** |  | [optional] 
**EarliestDeliveryDate** | Pointer to **string** |  | [optional] 
**LatestDeliveryDate** | Pointer to **string** |  | [optional] 
**DateTimeFrom** | Pointer to **time.Time** | Earliest date-time | [optional] [readonly] 
**DateTimeTo** | Pointer to **time.Time** | latest date-time | [optional] [readonly] 
**PlannedDateTimeFrom** | Pointer to **time.Time** | planned date-time from (only filled for planned activities) | [optional] [readonly] 
**PlannedDateTimeTo** | Pointer to **time.Time** | planned date-time to (only filled for planned activities) | [optional] [readonly] 
**ExecutedDateTimeFrom** | Pointer to **time.Time** | executed date-time from (only filled for executed activities) | [optional] [readonly] 
**ExecutedDateTimeTo** | Pointer to **time.Time** | executed date-time to (only filled for executed activities) | [optional] [readonly] 
**Duration** | Pointer to **int32** | Duration of this activity in minutes | [optional] 
**DepotDuration** | Pointer to **int32** | Duration of the depot activity in minutes | [optional] 
**Active** | Pointer to **bool** | if active&#x3D;0: activity has been removed and is no longer visible in any bumbal interface | [optional] 
**RouteId** | Pointer to **string** | Route ID (unique) | [optional] [readonly] 
**RouteNr** | Pointer to **string** | Route Number (not unique) | [optional] [readonly] 
**RouteName** | Pointer to **string** | Route name | [optional] [readonly] 
**RouteDateTimeFrom** | Pointer to **time.Time** | Start date-time of route | [optional] [readonly] 
**RouteDateTimeTo** | Pointer to **time.Time** | End date-time of route | [optional] [readonly] 
**RouteEarliestDateTime** | Pointer to **time.Time** | Earliest date-time of route | [optional] [readonly] 
**RouteLatestDateTime** | Pointer to **time.Time** | Latest date-time of route | [optional] [readonly] 
**RouteStatusId** | Pointer to **int32** | Status ID of this Route | [optional] [readonly] 
**RouteStatusName** | Pointer to **string** | Status Name of this Route | [optional] [readonly] 
**RouteNrOfStops** | Pointer to **int32** | Number of stops on this route | [optional] [readonly] 
**RouteOverdue** | Pointer to **bool** | Route is overdue | [optional] [readonly] 
**CoDriverIds** | Pointer to **[]int32** | Unique Identifier(s) for co-drivers on route | [optional] 
**MatchingRouteIds** | Pointer to **[]int32** | Matching route ID&#39;s | [optional] 
**CoDrivers** | Pointer to [**[]UsersModel**](UsersModel.md) | list of co-drivers on route | [optional] [readonly] 
**UserId** | Pointer to **int32** | ID of the user who will execute this activity | [optional] 
**DriverId** | Pointer to **int32** | Driver ID connected to this route/activity | [optional] 
**DriverFirstName** | Pointer to **string** | Driver first name | [optional] [readonly] 
**DriverLastName** | Pointer to **string** | Driver last name | [optional] [readonly] 
**DriverNamePrefix** | Pointer to **string** | Driver prefix | [optional] [readonly] 
**DriverFullName** | Pointer to **string** | Driver full name | [optional] [readonly] 
**AddressappliedId** | Pointer to **int32** | Address Applied ID | [optional] 
**AddressId** | Pointer to **int32** | Address ID | [optional] 
**Address** | Pointer to [**AddressModel**](AddressModel.md) |  | [optional] 
**AddressApplied** | Pointer to [**AddressAppliedModel**](AddressAppliedModel.md) |  | [optional] 
**DepotAddress** | Pointer to [**AddressModel**](AddressModel.md) |  | [optional] 
**DepotAddressId** | Pointer to **int32** | Depot Address ID | [optional] 
**DepotActivity** | Pointer to [**ActivityModel**](ActivityModel.md) |  | [optional] 
**AllowedDriverIds** | Pointer to **[]int32** | Unique Identifier(s) for allowed drivers in activity | [optional] 
**AllowedDrivers** | Pointer to [**[]AllowedDriverModel**](AllowedDriverModel.md) |  | [optional] 
**AllowedDriversLinks** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**AppliedCapacities** | Pointer to [**AppliedCapacitiesModel**](AppliedCapacitiesModel.md) |  | [optional] 
**Capacities** | Pointer to [**[]CapacityModel**](CapacityModel.md) |  | [optional] 
**PackageLines** | Pointer to [**[]PackageLineModel**](PackageLineModel.md) |  | [optional] 
**Payments** | Pointer to [**[]PaymentModel**](PaymentModel.md) |  | [optional] 
**TimeSlots** | Pointer to [**[]TimeSlotModel**](TimeSlotModel.md) |  | [optional] 
**BrandId** | Pointer to **int32** | Brand ID | [optional] 
**BrandName** | Pointer to **string** | Brand name (must be unique) | [optional] 
**Brand** | Pointer to [**BrandModel**](BrandModel.md) |  | [optional] 
**Communication** | Pointer to [**CommunicationModel**](CommunicationModel.md) |  | [optional] 
**AssignmentLink** | Pointer to [**LinkModel**](LinkModel.md) |  | [optional] 
**RouteLink** | Pointer to [**LinkModel**](LinkModel.md) |  | [optional] 
**Route** | Pointer to [**RouteModel**](RouteModel.md) |  | [optional] 
**Driver** | Pointer to [**DriverModel**](DriverModel.md) |  | [optional] 
**DriverLinks** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**Car** | Pointer to [**VehicleModel**](VehicleModel.md) |  | [optional] 
**Vehicle** | Pointer to [**VehicleModel**](VehicleModel.md) |  | [optional] 
**Trailer** | Pointer to [**TrailerModel**](TrailerModel.md) |  | [optional] 
**ActivityIdBefore** | Pointer to **string** | Activity ID of the activity which must be performed before this activity | [optional] 
**ActivityIdAfter** | Pointer to **string** | Activity ID of the activity which must be performed after this activity | [optional] 
**BundledActivityIds** | Pointer to **[]int32** | Ids of activities within bundle activity | [optional] 
**Tags** | Pointer to [**[]TagModel**](TagModel.md) |  | [optional] 
**Recurrence** | Pointer to [**RecurrenceModel**](RecurrenceModel.md) |  | [optional] 
**RecurrenceId** | Pointer to **int64** | Unique Recurrence Identifier | [optional] [readonly] 
**TagNames** | Pointer to **[]string** | Tag names | [optional] 
**TagIds** | Pointer to **[]int32** | Tag type ids | [optional] 
**Zones** | Pointer to [**[]ZoneModel**](ZoneModel.md) |  | [optional] 
**ZoneNames** | Pointer to **[]string** | Zone names | [optional] 
**Links** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**MetaData** | Pointer to [**[]MetaDataModel**](MetaDataModel.md) |  | [optional] 
**Stats** | Pointer to [**ActivityStatsModel**](ActivityStatsModel.md) |  | [optional] 
**Notes** | Pointer to [**[]NoteModel**](NoteModel.md) |  | [optional] 
**Files** | Pointer to [**[]FileModel**](FileModel.md) |  | [optional] 
**ActivityCreatedAt** | Pointer to **time.Time** | created_at date time | [optional] [readonly] 
**ActivityUpdatedAt** | Pointer to **time.Time** | updated_at date time | [optional] [readonly] 
**ActivityCreatedBy** | Pointer to **int32** | created_by user id | [optional] [readonly] 
**ActivityUpdatedBy** | Pointer to **int32** | updated_by user id | [optional] [readonly] 
**ActivityCreatedByUser** | Pointer to [**UsersModel**](UsersModel.md) |  | [optional] 
**ActivityUpdatedByUser** | Pointer to [**UsersModel**](UsersModel.md) |  | [optional] 
**ActivityActive** | Pointer to **bool** | Activity is active (&#x3D;true). Inactive activities are not automatically considered in any of the application algorithms and will not be shown in the Bumbal Gui. | [optional] 
**ActivityRemoved** | Pointer to **bool** | Activity is removed (&#x3D;true). Removed activities are not automatically considered in any of the application algorithms and will not be shown in the Bumbal Gui. Removed activities are usually irrepairable. | [optional] 
**PaymentTotal** | Pointer to **int32** | Total to be paid in cents. readonly | [optional] 
**TransactionTotal** | Pointer to **int32** | Total has been paid in cents. readonly | [optional] 
**DueTotal** | Pointer to **int32** | Amount which hasn&#39;t been paid yet in cents. readonly | [optional] 
**Transactions** | Pointer to [**[]TransactionModel**](TransactionModel.md) |  | [optional] 

## Methods

### NewActivityModel

`func NewActivityModel() *ActivityModel`

NewActivityModel instantiates a new ActivityModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityModelWithDefaults

`func NewActivityModelWithDefaults() *ActivityModel`

NewActivityModelWithDefaults instantiates a new ActivityModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActivityModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivityModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivityModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ActivityModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ActivityModel) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ActivityModel) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ActivityModel) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ActivityModel) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetBundleActivityId

`func (o *ActivityModel) GetBundleActivityId() int64`

GetBundleActivityId returns the BundleActivityId field if non-nil, zero value otherwise.

### GetBundleActivityIdOk

`func (o *ActivityModel) GetBundleActivityIdOk() (*int64, bool)`

GetBundleActivityIdOk returns a tuple with the BundleActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleActivityId

`func (o *ActivityModel) SetBundleActivityId(v int64)`

SetBundleActivityId sets BundleActivityId field to given value.

### HasBundleActivityId

`func (o *ActivityModel) HasBundleActivityId() bool`

HasBundleActivityId returns a boolean if a field has been set.

### GetShipmentActivityId

`func (o *ActivityModel) GetShipmentActivityId() int64`

GetShipmentActivityId returns the ShipmentActivityId field if non-nil, zero value otherwise.

### GetShipmentActivityIdOk

`func (o *ActivityModel) GetShipmentActivityIdOk() (*int64, bool)`

GetShipmentActivityIdOk returns a tuple with the ShipmentActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentActivityId

`func (o *ActivityModel) SetShipmentActivityId(v int64)`

SetShipmentActivityId sets ShipmentActivityId field to given value.

### HasShipmentActivityId

`func (o *ActivityModel) HasShipmentActivityId() bool`

HasShipmentActivityId returns a boolean if a field has been set.

### GetShipmentActivityNr

`func (o *ActivityModel) GetShipmentActivityNr() string`

GetShipmentActivityNr returns the ShipmentActivityNr field if non-nil, zero value otherwise.

### GetShipmentActivityNrOk

`func (o *ActivityModel) GetShipmentActivityNrOk() (*string, bool)`

GetShipmentActivityNrOk returns a tuple with the ShipmentActivityNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentActivityNr

`func (o *ActivityModel) SetShipmentActivityNr(v string)`

SetShipmentActivityNr sets ShipmentActivityNr field to given value.

### HasShipmentActivityNr

`func (o *ActivityModel) HasShipmentActivityNr() bool`

HasShipmentActivityNr returns a boolean if a field has been set.

### GetBundleShipmentActivityNr

`func (o *ActivityModel) GetBundleShipmentActivityNr() string`

GetBundleShipmentActivityNr returns the BundleShipmentActivityNr field if non-nil, zero value otherwise.

### GetBundleShipmentActivityNrOk

`func (o *ActivityModel) GetBundleShipmentActivityNrOk() (*string, bool)`

GetBundleShipmentActivityNrOk returns a tuple with the BundleShipmentActivityNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleShipmentActivityNr

`func (o *ActivityModel) SetBundleShipmentActivityNr(v string)`

SetBundleShipmentActivityNr sets BundleShipmentActivityNr field to given value.

### HasBundleShipmentActivityNr

`func (o *ActivityModel) HasBundleShipmentActivityNr() bool`

HasBundleShipmentActivityNr returns a boolean if a field has been set.

### GetNr

`func (o *ActivityModel) GetNr() string`

GetNr returns the Nr field if non-nil, zero value otherwise.

### GetNrOk

`func (o *ActivityModel) GetNrOk() (*string, bool)`

GetNrOk returns a tuple with the Nr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNr

`func (o *ActivityModel) SetNr(v string)`

SetNr sets Nr field to given value.

### HasNr

`func (o *ActivityModel) HasNr() bool`

HasNr returns a boolean if a field has been set.

### GetActivityTypeName

`func (o *ActivityModel) GetActivityTypeName() string`

GetActivityTypeName returns the ActivityTypeName field if non-nil, zero value otherwise.

### GetActivityTypeNameOk

`func (o *ActivityModel) GetActivityTypeNameOk() (*string, bool)`

GetActivityTypeNameOk returns a tuple with the ActivityTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityTypeName

`func (o *ActivityModel) SetActivityTypeName(v string)`

SetActivityTypeName sets ActivityTypeName field to given value.

### HasActivityTypeName

`func (o *ActivityModel) HasActivityTypeName() bool`

HasActivityTypeName returns a boolean if a field has been set.

### GetActivityTypeId

`func (o *ActivityModel) GetActivityTypeId() int64`

GetActivityTypeId returns the ActivityTypeId field if non-nil, zero value otherwise.

### GetActivityTypeIdOk

`func (o *ActivityModel) GetActivityTypeIdOk() (*int64, bool)`

GetActivityTypeIdOk returns a tuple with the ActivityTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityTypeId

`func (o *ActivityModel) SetActivityTypeId(v int64)`

SetActivityTypeId sets ActivityTypeId field to given value.

### HasActivityTypeId

`func (o *ActivityModel) HasActivityTypeId() bool`

HasActivityTypeId returns a boolean if a field has been set.

### GetStatusId

`func (o *ActivityModel) GetStatusId() int32`

GetStatusId returns the StatusId field if non-nil, zero value otherwise.

### GetStatusIdOk

`func (o *ActivityModel) GetStatusIdOk() (*int32, bool)`

GetStatusIdOk returns a tuple with the StatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusId

`func (o *ActivityModel) SetStatusId(v int32)`

SetStatusId sets StatusId field to given value.

### HasStatusId

`func (o *ActivityModel) HasStatusId() bool`

HasStatusId returns a boolean if a field has been set.

### GetStatusName

`func (o *ActivityModel) GetStatusName() string`

GetStatusName returns the StatusName field if non-nil, zero value otherwise.

### GetStatusNameOk

`func (o *ActivityModel) GetStatusNameOk() (*string, bool)`

GetStatusNameOk returns a tuple with the StatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusName

`func (o *ActivityModel) SetStatusName(v string)`

SetStatusName sets StatusName field to given value.

### HasStatusName

`func (o *ActivityModel) HasStatusName() bool`

HasStatusName returns a boolean if a field has been set.

### GetAssignmentId

`func (o *ActivityModel) GetAssignmentId() int32`

GetAssignmentId returns the AssignmentId field if non-nil, zero value otherwise.

### GetAssignmentIdOk

`func (o *ActivityModel) GetAssignmentIdOk() (*int32, bool)`

GetAssignmentIdOk returns a tuple with the AssignmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentId

`func (o *ActivityModel) SetAssignmentId(v int32)`

SetAssignmentId sets AssignmentId field to given value.

### HasAssignmentId

`func (o *ActivityModel) HasAssignmentId() bool`

HasAssignmentId returns a boolean if a field has been set.

### GetAssignmentNr

`func (o *ActivityModel) GetAssignmentNr() string`

GetAssignmentNr returns the AssignmentNr field if non-nil, zero value otherwise.

### GetAssignmentNrOk

`func (o *ActivityModel) GetAssignmentNrOk() (*string, bool)`

GetAssignmentNrOk returns a tuple with the AssignmentNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentNr

`func (o *ActivityModel) SetAssignmentNr(v string)`

SetAssignmentNr sets AssignmentNr field to given value.

### HasAssignmentNr

`func (o *ActivityModel) HasAssignmentNr() bool`

HasAssignmentNr returns a boolean if a field has been set.

### GetAssignment

`func (o *ActivityModel) GetAssignment() AssignmentModel`

GetAssignment returns the Assignment field if non-nil, zero value otherwise.

### GetAssignmentOk

`func (o *ActivityModel) GetAssignmentOk() (*AssignmentModel, bool)`

GetAssignmentOk returns a tuple with the Assignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignment

`func (o *ActivityModel) SetAssignment(v AssignmentModel)`

SetAssignment sets Assignment field to given value.

### HasAssignment

`func (o *ActivityModel) HasAssignment() bool`

HasAssignment returns a boolean if a field has been set.

### GetAssignmentSequenceNr

`func (o *ActivityModel) GetAssignmentSequenceNr() int64`

GetAssignmentSequenceNr returns the AssignmentSequenceNr field if non-nil, zero value otherwise.

### GetAssignmentSequenceNrOk

`func (o *ActivityModel) GetAssignmentSequenceNrOk() (*int64, bool)`

GetAssignmentSequenceNrOk returns a tuple with the AssignmentSequenceNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentSequenceNr

`func (o *ActivityModel) SetAssignmentSequenceNr(v int64)`

SetAssignmentSequenceNr sets AssignmentSequenceNr field to given value.

### HasAssignmentSequenceNr

`func (o *ActivityModel) HasAssignmentSequenceNr() bool`

HasAssignmentSequenceNr returns a boolean if a field has been set.

### GetRemarks

`func (o *ActivityModel) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *ActivityModel) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *ActivityModel) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.

### HasRemarks

`func (o *ActivityModel) HasRemarks() bool`

HasRemarks returns a boolean if a field has been set.

### GetLocked

`func (o *ActivityModel) GetLocked() int32`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *ActivityModel) GetLockedOk() (*int32, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *ActivityModel) SetLocked(v int32)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *ActivityModel) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetOptimized

`func (o *ActivityModel) GetOptimized() bool`

GetOptimized returns the Optimized field if non-nil, zero value otherwise.

### GetOptimizedOk

`func (o *ActivityModel) GetOptimizedOk() (*bool, bool)`

GetOptimizedOk returns a tuple with the Optimized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimized

`func (o *ActivityModel) SetOptimized(v bool)`

SetOptimized sets Optimized field to given value.

### HasOptimized

`func (o *ActivityModel) HasOptimized() bool`

HasOptimized returns a boolean if a field has been set.

### GetPriority

`func (o *ActivityModel) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ActivityModel) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ActivityModel) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ActivityModel) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetTopPriority

`func (o *ActivityModel) GetTopPriority() bool`

GetTopPriority returns the TopPriority field if non-nil, zero value otherwise.

### GetTopPriorityOk

`func (o *ActivityModel) GetTopPriorityOk() (*bool, bool)`

GetTopPriorityOk returns a tuple with the TopPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopPriority

`func (o *ActivityModel) SetTopPriority(v bool)`

SetTopPriority sets TopPriority field to given value.

### HasTopPriority

`func (o *ActivityModel) HasTopPriority() bool`

HasTopPriority returns a boolean if a field has been set.

### GetSequenceNr

`func (o *ActivityModel) GetSequenceNr() int64`

GetSequenceNr returns the SequenceNr field if non-nil, zero value otherwise.

### GetSequenceNrOk

`func (o *ActivityModel) GetSequenceNrOk() (*int64, bool)`

GetSequenceNrOk returns a tuple with the SequenceNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNr

`func (o *ActivityModel) SetSequenceNr(v int64)`

SetSequenceNr sets SequenceNr field to given value.

### HasSequenceNr

`func (o *ActivityModel) HasSequenceNr() bool`

HasSequenceNr returns a boolean if a field has been set.

### GetPlannedDrivingTime

`func (o *ActivityModel) GetPlannedDrivingTime() int64`

GetPlannedDrivingTime returns the PlannedDrivingTime field if non-nil, zero value otherwise.

### GetPlannedDrivingTimeOk

`func (o *ActivityModel) GetPlannedDrivingTimeOk() (*int64, bool)`

GetPlannedDrivingTimeOk returns a tuple with the PlannedDrivingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDrivingTime

`func (o *ActivityModel) SetPlannedDrivingTime(v int64)`

SetPlannedDrivingTime sets PlannedDrivingTime field to given value.

### HasPlannedDrivingTime

`func (o *ActivityModel) HasPlannedDrivingTime() bool`

HasPlannedDrivingTime returns a boolean if a field has been set.

### GetPlannedDrivingDuration

`func (o *ActivityModel) GetPlannedDrivingDuration() int64`

GetPlannedDrivingDuration returns the PlannedDrivingDuration field if non-nil, zero value otherwise.

### GetPlannedDrivingDurationOk

`func (o *ActivityModel) GetPlannedDrivingDurationOk() (*int64, bool)`

GetPlannedDrivingDurationOk returns a tuple with the PlannedDrivingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDrivingDuration

`func (o *ActivityModel) SetPlannedDrivingDuration(v int64)`

SetPlannedDrivingDuration sets PlannedDrivingDuration field to given value.

### HasPlannedDrivingDuration

`func (o *ActivityModel) HasPlannedDrivingDuration() bool`

HasPlannedDrivingDuration returns a boolean if a field has been set.

### GetPlannedDrivingDistance

`func (o *ActivityModel) GetPlannedDrivingDistance() int64`

GetPlannedDrivingDistance returns the PlannedDrivingDistance field if non-nil, zero value otherwise.

### GetPlannedDrivingDistanceOk

`func (o *ActivityModel) GetPlannedDrivingDistanceOk() (*int64, bool)`

GetPlannedDrivingDistanceOk returns a tuple with the PlannedDrivingDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDrivingDistance

`func (o *ActivityModel) SetPlannedDrivingDistance(v int64)`

SetPlannedDrivingDistance sets PlannedDrivingDistance field to given value.

### HasPlannedDrivingDistance

`func (o *ActivityModel) HasPlannedDrivingDistance() bool`

HasPlannedDrivingDistance returns a boolean if a field has been set.

### GetReference

`func (o *ActivityModel) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ActivityModel) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ActivityModel) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ActivityModel) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetDescription

`func (o *ActivityModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ActivityModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ActivityModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ActivityModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEarliestDeliveryDateTime

`func (o *ActivityModel) GetEarliestDeliveryDateTime() string`

GetEarliestDeliveryDateTime returns the EarliestDeliveryDateTime field if non-nil, zero value otherwise.

### GetEarliestDeliveryDateTimeOk

`func (o *ActivityModel) GetEarliestDeliveryDateTimeOk() (*string, bool)`

GetEarliestDeliveryDateTimeOk returns a tuple with the EarliestDeliveryDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestDeliveryDateTime

`func (o *ActivityModel) SetEarliestDeliveryDateTime(v string)`

SetEarliestDeliveryDateTime sets EarliestDeliveryDateTime field to given value.

### HasEarliestDeliveryDateTime

`func (o *ActivityModel) HasEarliestDeliveryDateTime() bool`

HasEarliestDeliveryDateTime returns a boolean if a field has been set.

### GetLatestDeliveryDateTime

`func (o *ActivityModel) GetLatestDeliveryDateTime() string`

GetLatestDeliveryDateTime returns the LatestDeliveryDateTime field if non-nil, zero value otherwise.

### GetLatestDeliveryDateTimeOk

`func (o *ActivityModel) GetLatestDeliveryDateTimeOk() (*string, bool)`

GetLatestDeliveryDateTimeOk returns a tuple with the LatestDeliveryDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDeliveryDateTime

`func (o *ActivityModel) SetLatestDeliveryDateTime(v string)`

SetLatestDeliveryDateTime sets LatestDeliveryDateTime field to given value.

### HasLatestDeliveryDateTime

`func (o *ActivityModel) HasLatestDeliveryDateTime() bool`

HasLatestDeliveryDateTime returns a boolean if a field has been set.

### GetEarliestDeliveryDate

`func (o *ActivityModel) GetEarliestDeliveryDate() string`

GetEarliestDeliveryDate returns the EarliestDeliveryDate field if non-nil, zero value otherwise.

### GetEarliestDeliveryDateOk

`func (o *ActivityModel) GetEarliestDeliveryDateOk() (*string, bool)`

GetEarliestDeliveryDateOk returns a tuple with the EarliestDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestDeliveryDate

`func (o *ActivityModel) SetEarliestDeliveryDate(v string)`

SetEarliestDeliveryDate sets EarliestDeliveryDate field to given value.

### HasEarliestDeliveryDate

`func (o *ActivityModel) HasEarliestDeliveryDate() bool`

HasEarliestDeliveryDate returns a boolean if a field has been set.

### GetLatestDeliveryDate

`func (o *ActivityModel) GetLatestDeliveryDate() string`

GetLatestDeliveryDate returns the LatestDeliveryDate field if non-nil, zero value otherwise.

### GetLatestDeliveryDateOk

`func (o *ActivityModel) GetLatestDeliveryDateOk() (*string, bool)`

GetLatestDeliveryDateOk returns a tuple with the LatestDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDeliveryDate

`func (o *ActivityModel) SetLatestDeliveryDate(v string)`

SetLatestDeliveryDate sets LatestDeliveryDate field to given value.

### HasLatestDeliveryDate

`func (o *ActivityModel) HasLatestDeliveryDate() bool`

HasLatestDeliveryDate returns a boolean if a field has been set.

### GetDateTimeFrom

`func (o *ActivityModel) GetDateTimeFrom() time.Time`

GetDateTimeFrom returns the DateTimeFrom field if non-nil, zero value otherwise.

### GetDateTimeFromOk

`func (o *ActivityModel) GetDateTimeFromOk() (*time.Time, bool)`

GetDateTimeFromOk returns a tuple with the DateTimeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFrom

`func (o *ActivityModel) SetDateTimeFrom(v time.Time)`

SetDateTimeFrom sets DateTimeFrom field to given value.

### HasDateTimeFrom

`func (o *ActivityModel) HasDateTimeFrom() bool`

HasDateTimeFrom returns a boolean if a field has been set.

### GetDateTimeTo

`func (o *ActivityModel) GetDateTimeTo() time.Time`

GetDateTimeTo returns the DateTimeTo field if non-nil, zero value otherwise.

### GetDateTimeToOk

`func (o *ActivityModel) GetDateTimeToOk() (*time.Time, bool)`

GetDateTimeToOk returns a tuple with the DateTimeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeTo

`func (o *ActivityModel) SetDateTimeTo(v time.Time)`

SetDateTimeTo sets DateTimeTo field to given value.

### HasDateTimeTo

`func (o *ActivityModel) HasDateTimeTo() bool`

HasDateTimeTo returns a boolean if a field has been set.

### GetPlannedDateTimeFrom

`func (o *ActivityModel) GetPlannedDateTimeFrom() time.Time`

GetPlannedDateTimeFrom returns the PlannedDateTimeFrom field if non-nil, zero value otherwise.

### GetPlannedDateTimeFromOk

`func (o *ActivityModel) GetPlannedDateTimeFromOk() (*time.Time, bool)`

GetPlannedDateTimeFromOk returns a tuple with the PlannedDateTimeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDateTimeFrom

`func (o *ActivityModel) SetPlannedDateTimeFrom(v time.Time)`

SetPlannedDateTimeFrom sets PlannedDateTimeFrom field to given value.

### HasPlannedDateTimeFrom

`func (o *ActivityModel) HasPlannedDateTimeFrom() bool`

HasPlannedDateTimeFrom returns a boolean if a field has been set.

### GetPlannedDateTimeTo

`func (o *ActivityModel) GetPlannedDateTimeTo() time.Time`

GetPlannedDateTimeTo returns the PlannedDateTimeTo field if non-nil, zero value otherwise.

### GetPlannedDateTimeToOk

`func (o *ActivityModel) GetPlannedDateTimeToOk() (*time.Time, bool)`

GetPlannedDateTimeToOk returns a tuple with the PlannedDateTimeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDateTimeTo

`func (o *ActivityModel) SetPlannedDateTimeTo(v time.Time)`

SetPlannedDateTimeTo sets PlannedDateTimeTo field to given value.

### HasPlannedDateTimeTo

`func (o *ActivityModel) HasPlannedDateTimeTo() bool`

HasPlannedDateTimeTo returns a boolean if a field has been set.

### GetExecutedDateTimeFrom

`func (o *ActivityModel) GetExecutedDateTimeFrom() time.Time`

GetExecutedDateTimeFrom returns the ExecutedDateTimeFrom field if non-nil, zero value otherwise.

### GetExecutedDateTimeFromOk

`func (o *ActivityModel) GetExecutedDateTimeFromOk() (*time.Time, bool)`

GetExecutedDateTimeFromOk returns a tuple with the ExecutedDateTimeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedDateTimeFrom

`func (o *ActivityModel) SetExecutedDateTimeFrom(v time.Time)`

SetExecutedDateTimeFrom sets ExecutedDateTimeFrom field to given value.

### HasExecutedDateTimeFrom

`func (o *ActivityModel) HasExecutedDateTimeFrom() bool`

HasExecutedDateTimeFrom returns a boolean if a field has been set.

### GetExecutedDateTimeTo

`func (o *ActivityModel) GetExecutedDateTimeTo() time.Time`

GetExecutedDateTimeTo returns the ExecutedDateTimeTo field if non-nil, zero value otherwise.

### GetExecutedDateTimeToOk

`func (o *ActivityModel) GetExecutedDateTimeToOk() (*time.Time, bool)`

GetExecutedDateTimeToOk returns a tuple with the ExecutedDateTimeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedDateTimeTo

`func (o *ActivityModel) SetExecutedDateTimeTo(v time.Time)`

SetExecutedDateTimeTo sets ExecutedDateTimeTo field to given value.

### HasExecutedDateTimeTo

`func (o *ActivityModel) HasExecutedDateTimeTo() bool`

HasExecutedDateTimeTo returns a boolean if a field has been set.

### GetDuration

`func (o *ActivityModel) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ActivityModel) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ActivityModel) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ActivityModel) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetDepotDuration

`func (o *ActivityModel) GetDepotDuration() int32`

GetDepotDuration returns the DepotDuration field if non-nil, zero value otherwise.

### GetDepotDurationOk

`func (o *ActivityModel) GetDepotDurationOk() (*int32, bool)`

GetDepotDurationOk returns a tuple with the DepotDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotDuration

`func (o *ActivityModel) SetDepotDuration(v int32)`

SetDepotDuration sets DepotDuration field to given value.

### HasDepotDuration

`func (o *ActivityModel) HasDepotDuration() bool`

HasDepotDuration returns a boolean if a field has been set.

### GetActive

`func (o *ActivityModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ActivityModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ActivityModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ActivityModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetRouteId

`func (o *ActivityModel) GetRouteId() string`

GetRouteId returns the RouteId field if non-nil, zero value otherwise.

### GetRouteIdOk

`func (o *ActivityModel) GetRouteIdOk() (*string, bool)`

GetRouteIdOk returns a tuple with the RouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteId

`func (o *ActivityModel) SetRouteId(v string)`

SetRouteId sets RouteId field to given value.

### HasRouteId

`func (o *ActivityModel) HasRouteId() bool`

HasRouteId returns a boolean if a field has been set.

### GetRouteNr

`func (o *ActivityModel) GetRouteNr() string`

GetRouteNr returns the RouteNr field if non-nil, zero value otherwise.

### GetRouteNrOk

`func (o *ActivityModel) GetRouteNrOk() (*string, bool)`

GetRouteNrOk returns a tuple with the RouteNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteNr

`func (o *ActivityModel) SetRouteNr(v string)`

SetRouteNr sets RouteNr field to given value.

### HasRouteNr

`func (o *ActivityModel) HasRouteNr() bool`

HasRouteNr returns a boolean if a field has been set.

### GetRouteName

`func (o *ActivityModel) GetRouteName() string`

GetRouteName returns the RouteName field if non-nil, zero value otherwise.

### GetRouteNameOk

`func (o *ActivityModel) GetRouteNameOk() (*string, bool)`

GetRouteNameOk returns a tuple with the RouteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteName

`func (o *ActivityModel) SetRouteName(v string)`

SetRouteName sets RouteName field to given value.

### HasRouteName

`func (o *ActivityModel) HasRouteName() bool`

HasRouteName returns a boolean if a field has been set.

### GetRouteDateTimeFrom

`func (o *ActivityModel) GetRouteDateTimeFrom() time.Time`

GetRouteDateTimeFrom returns the RouteDateTimeFrom field if non-nil, zero value otherwise.

### GetRouteDateTimeFromOk

`func (o *ActivityModel) GetRouteDateTimeFromOk() (*time.Time, bool)`

GetRouteDateTimeFromOk returns a tuple with the RouteDateTimeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteDateTimeFrom

`func (o *ActivityModel) SetRouteDateTimeFrom(v time.Time)`

SetRouteDateTimeFrom sets RouteDateTimeFrom field to given value.

### HasRouteDateTimeFrom

`func (o *ActivityModel) HasRouteDateTimeFrom() bool`

HasRouteDateTimeFrom returns a boolean if a field has been set.

### GetRouteDateTimeTo

`func (o *ActivityModel) GetRouteDateTimeTo() time.Time`

GetRouteDateTimeTo returns the RouteDateTimeTo field if non-nil, zero value otherwise.

### GetRouteDateTimeToOk

`func (o *ActivityModel) GetRouteDateTimeToOk() (*time.Time, bool)`

GetRouteDateTimeToOk returns a tuple with the RouteDateTimeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteDateTimeTo

`func (o *ActivityModel) SetRouteDateTimeTo(v time.Time)`

SetRouteDateTimeTo sets RouteDateTimeTo field to given value.

### HasRouteDateTimeTo

`func (o *ActivityModel) HasRouteDateTimeTo() bool`

HasRouteDateTimeTo returns a boolean if a field has been set.

### GetRouteEarliestDateTime

`func (o *ActivityModel) GetRouteEarliestDateTime() time.Time`

GetRouteEarliestDateTime returns the RouteEarliestDateTime field if non-nil, zero value otherwise.

### GetRouteEarliestDateTimeOk

`func (o *ActivityModel) GetRouteEarliestDateTimeOk() (*time.Time, bool)`

GetRouteEarliestDateTimeOk returns a tuple with the RouteEarliestDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteEarliestDateTime

`func (o *ActivityModel) SetRouteEarliestDateTime(v time.Time)`

SetRouteEarliestDateTime sets RouteEarliestDateTime field to given value.

### HasRouteEarliestDateTime

`func (o *ActivityModel) HasRouteEarliestDateTime() bool`

HasRouteEarliestDateTime returns a boolean if a field has been set.

### GetRouteLatestDateTime

`func (o *ActivityModel) GetRouteLatestDateTime() time.Time`

GetRouteLatestDateTime returns the RouteLatestDateTime field if non-nil, zero value otherwise.

### GetRouteLatestDateTimeOk

`func (o *ActivityModel) GetRouteLatestDateTimeOk() (*time.Time, bool)`

GetRouteLatestDateTimeOk returns a tuple with the RouteLatestDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteLatestDateTime

`func (o *ActivityModel) SetRouteLatestDateTime(v time.Time)`

SetRouteLatestDateTime sets RouteLatestDateTime field to given value.

### HasRouteLatestDateTime

`func (o *ActivityModel) HasRouteLatestDateTime() bool`

HasRouteLatestDateTime returns a boolean if a field has been set.

### GetRouteStatusId

`func (o *ActivityModel) GetRouteStatusId() int32`

GetRouteStatusId returns the RouteStatusId field if non-nil, zero value otherwise.

### GetRouteStatusIdOk

`func (o *ActivityModel) GetRouteStatusIdOk() (*int32, bool)`

GetRouteStatusIdOk returns a tuple with the RouteStatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteStatusId

`func (o *ActivityModel) SetRouteStatusId(v int32)`

SetRouteStatusId sets RouteStatusId field to given value.

### HasRouteStatusId

`func (o *ActivityModel) HasRouteStatusId() bool`

HasRouteStatusId returns a boolean if a field has been set.

### GetRouteStatusName

`func (o *ActivityModel) GetRouteStatusName() string`

GetRouteStatusName returns the RouteStatusName field if non-nil, zero value otherwise.

### GetRouteStatusNameOk

`func (o *ActivityModel) GetRouteStatusNameOk() (*string, bool)`

GetRouteStatusNameOk returns a tuple with the RouteStatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteStatusName

`func (o *ActivityModel) SetRouteStatusName(v string)`

SetRouteStatusName sets RouteStatusName field to given value.

### HasRouteStatusName

`func (o *ActivityModel) HasRouteStatusName() bool`

HasRouteStatusName returns a boolean if a field has been set.

### GetRouteNrOfStops

`func (o *ActivityModel) GetRouteNrOfStops() int32`

GetRouteNrOfStops returns the RouteNrOfStops field if non-nil, zero value otherwise.

### GetRouteNrOfStopsOk

`func (o *ActivityModel) GetRouteNrOfStopsOk() (*int32, bool)`

GetRouteNrOfStopsOk returns a tuple with the RouteNrOfStops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteNrOfStops

`func (o *ActivityModel) SetRouteNrOfStops(v int32)`

SetRouteNrOfStops sets RouteNrOfStops field to given value.

### HasRouteNrOfStops

`func (o *ActivityModel) HasRouteNrOfStops() bool`

HasRouteNrOfStops returns a boolean if a field has been set.

### GetRouteOverdue

`func (o *ActivityModel) GetRouteOverdue() bool`

GetRouteOverdue returns the RouteOverdue field if non-nil, zero value otherwise.

### GetRouteOverdueOk

`func (o *ActivityModel) GetRouteOverdueOk() (*bool, bool)`

GetRouteOverdueOk returns a tuple with the RouteOverdue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteOverdue

`func (o *ActivityModel) SetRouteOverdue(v bool)`

SetRouteOverdue sets RouteOverdue field to given value.

### HasRouteOverdue

`func (o *ActivityModel) HasRouteOverdue() bool`

HasRouteOverdue returns a boolean if a field has been set.

### GetCoDriverIds

`func (o *ActivityModel) GetCoDriverIds() []int32`

GetCoDriverIds returns the CoDriverIds field if non-nil, zero value otherwise.

### GetCoDriverIdsOk

`func (o *ActivityModel) GetCoDriverIdsOk() (*[]int32, bool)`

GetCoDriverIdsOk returns a tuple with the CoDriverIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoDriverIds

`func (o *ActivityModel) SetCoDriverIds(v []int32)`

SetCoDriverIds sets CoDriverIds field to given value.

### HasCoDriverIds

`func (o *ActivityModel) HasCoDriverIds() bool`

HasCoDriverIds returns a boolean if a field has been set.

### GetMatchingRouteIds

`func (o *ActivityModel) GetMatchingRouteIds() []int32`

GetMatchingRouteIds returns the MatchingRouteIds field if non-nil, zero value otherwise.

### GetMatchingRouteIdsOk

`func (o *ActivityModel) GetMatchingRouteIdsOk() (*[]int32, bool)`

GetMatchingRouteIdsOk returns a tuple with the MatchingRouteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingRouteIds

`func (o *ActivityModel) SetMatchingRouteIds(v []int32)`

SetMatchingRouteIds sets MatchingRouteIds field to given value.

### HasMatchingRouteIds

`func (o *ActivityModel) HasMatchingRouteIds() bool`

HasMatchingRouteIds returns a boolean if a field has been set.

### GetCoDrivers

`func (o *ActivityModel) GetCoDrivers() []UsersModel`

GetCoDrivers returns the CoDrivers field if non-nil, zero value otherwise.

### GetCoDriversOk

`func (o *ActivityModel) GetCoDriversOk() (*[]UsersModel, bool)`

GetCoDriversOk returns a tuple with the CoDrivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoDrivers

`func (o *ActivityModel) SetCoDrivers(v []UsersModel)`

SetCoDrivers sets CoDrivers field to given value.

### HasCoDrivers

`func (o *ActivityModel) HasCoDrivers() bool`

HasCoDrivers returns a boolean if a field has been set.

### GetUserId

`func (o *ActivityModel) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ActivityModel) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ActivityModel) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ActivityModel) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetDriverId

`func (o *ActivityModel) GetDriverId() int32`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *ActivityModel) GetDriverIdOk() (*int32, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *ActivityModel) SetDriverId(v int32)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *ActivityModel) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### GetDriverFirstName

`func (o *ActivityModel) GetDriverFirstName() string`

GetDriverFirstName returns the DriverFirstName field if non-nil, zero value otherwise.

### GetDriverFirstNameOk

`func (o *ActivityModel) GetDriverFirstNameOk() (*string, bool)`

GetDriverFirstNameOk returns a tuple with the DriverFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverFirstName

`func (o *ActivityModel) SetDriverFirstName(v string)`

SetDriverFirstName sets DriverFirstName field to given value.

### HasDriverFirstName

`func (o *ActivityModel) HasDriverFirstName() bool`

HasDriverFirstName returns a boolean if a field has been set.

### GetDriverLastName

`func (o *ActivityModel) GetDriverLastName() string`

GetDriverLastName returns the DriverLastName field if non-nil, zero value otherwise.

### GetDriverLastNameOk

`func (o *ActivityModel) GetDriverLastNameOk() (*string, bool)`

GetDriverLastNameOk returns a tuple with the DriverLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverLastName

`func (o *ActivityModel) SetDriverLastName(v string)`

SetDriverLastName sets DriverLastName field to given value.

### HasDriverLastName

`func (o *ActivityModel) HasDriverLastName() bool`

HasDriverLastName returns a boolean if a field has been set.

### GetDriverNamePrefix

`func (o *ActivityModel) GetDriverNamePrefix() string`

GetDriverNamePrefix returns the DriverNamePrefix field if non-nil, zero value otherwise.

### GetDriverNamePrefixOk

`func (o *ActivityModel) GetDriverNamePrefixOk() (*string, bool)`

GetDriverNamePrefixOk returns a tuple with the DriverNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverNamePrefix

`func (o *ActivityModel) SetDriverNamePrefix(v string)`

SetDriverNamePrefix sets DriverNamePrefix field to given value.

### HasDriverNamePrefix

`func (o *ActivityModel) HasDriverNamePrefix() bool`

HasDriverNamePrefix returns a boolean if a field has been set.

### GetDriverFullName

`func (o *ActivityModel) GetDriverFullName() string`

GetDriverFullName returns the DriverFullName field if non-nil, zero value otherwise.

### GetDriverFullNameOk

`func (o *ActivityModel) GetDriverFullNameOk() (*string, bool)`

GetDriverFullNameOk returns a tuple with the DriverFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverFullName

`func (o *ActivityModel) SetDriverFullName(v string)`

SetDriverFullName sets DriverFullName field to given value.

### HasDriverFullName

`func (o *ActivityModel) HasDriverFullName() bool`

HasDriverFullName returns a boolean if a field has been set.

### GetAddressappliedId

`func (o *ActivityModel) GetAddressappliedId() int32`

GetAddressappliedId returns the AddressappliedId field if non-nil, zero value otherwise.

### GetAddressappliedIdOk

`func (o *ActivityModel) GetAddressappliedIdOk() (*int32, bool)`

GetAddressappliedIdOk returns a tuple with the AddressappliedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressappliedId

`func (o *ActivityModel) SetAddressappliedId(v int32)`

SetAddressappliedId sets AddressappliedId field to given value.

### HasAddressappliedId

`func (o *ActivityModel) HasAddressappliedId() bool`

HasAddressappliedId returns a boolean if a field has been set.

### GetAddressId

`func (o *ActivityModel) GetAddressId() int32`

GetAddressId returns the AddressId field if non-nil, zero value otherwise.

### GetAddressIdOk

`func (o *ActivityModel) GetAddressIdOk() (*int32, bool)`

GetAddressIdOk returns a tuple with the AddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressId

`func (o *ActivityModel) SetAddressId(v int32)`

SetAddressId sets AddressId field to given value.

### HasAddressId

`func (o *ActivityModel) HasAddressId() bool`

HasAddressId returns a boolean if a field has been set.

### GetAddress

`func (o *ActivityModel) GetAddress() AddressModel`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ActivityModel) GetAddressOk() (*AddressModel, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ActivityModel) SetAddress(v AddressModel)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ActivityModel) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddressApplied

`func (o *ActivityModel) GetAddressApplied() AddressAppliedModel`

GetAddressApplied returns the AddressApplied field if non-nil, zero value otherwise.

### GetAddressAppliedOk

`func (o *ActivityModel) GetAddressAppliedOk() (*AddressAppliedModel, bool)`

GetAddressAppliedOk returns a tuple with the AddressApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressApplied

`func (o *ActivityModel) SetAddressApplied(v AddressAppliedModel)`

SetAddressApplied sets AddressApplied field to given value.

### HasAddressApplied

`func (o *ActivityModel) HasAddressApplied() bool`

HasAddressApplied returns a boolean if a field has been set.

### GetDepotAddress

`func (o *ActivityModel) GetDepotAddress() AddressModel`

GetDepotAddress returns the DepotAddress field if non-nil, zero value otherwise.

### GetDepotAddressOk

`func (o *ActivityModel) GetDepotAddressOk() (*AddressModel, bool)`

GetDepotAddressOk returns a tuple with the DepotAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotAddress

`func (o *ActivityModel) SetDepotAddress(v AddressModel)`

SetDepotAddress sets DepotAddress field to given value.

### HasDepotAddress

`func (o *ActivityModel) HasDepotAddress() bool`

HasDepotAddress returns a boolean if a field has been set.

### GetDepotAddressId

`func (o *ActivityModel) GetDepotAddressId() int32`

GetDepotAddressId returns the DepotAddressId field if non-nil, zero value otherwise.

### GetDepotAddressIdOk

`func (o *ActivityModel) GetDepotAddressIdOk() (*int32, bool)`

GetDepotAddressIdOk returns a tuple with the DepotAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotAddressId

`func (o *ActivityModel) SetDepotAddressId(v int32)`

SetDepotAddressId sets DepotAddressId field to given value.

### HasDepotAddressId

`func (o *ActivityModel) HasDepotAddressId() bool`

HasDepotAddressId returns a boolean if a field has been set.

### GetDepotActivity

`func (o *ActivityModel) GetDepotActivity() ActivityModel`

GetDepotActivity returns the DepotActivity field if non-nil, zero value otherwise.

### GetDepotActivityOk

`func (o *ActivityModel) GetDepotActivityOk() (*ActivityModel, bool)`

GetDepotActivityOk returns a tuple with the DepotActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotActivity

`func (o *ActivityModel) SetDepotActivity(v ActivityModel)`

SetDepotActivity sets DepotActivity field to given value.

### HasDepotActivity

`func (o *ActivityModel) HasDepotActivity() bool`

HasDepotActivity returns a boolean if a field has been set.

### GetAllowedDriverIds

`func (o *ActivityModel) GetAllowedDriverIds() []int32`

GetAllowedDriverIds returns the AllowedDriverIds field if non-nil, zero value otherwise.

### GetAllowedDriverIdsOk

`func (o *ActivityModel) GetAllowedDriverIdsOk() (*[]int32, bool)`

GetAllowedDriverIdsOk returns a tuple with the AllowedDriverIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDriverIds

`func (o *ActivityModel) SetAllowedDriverIds(v []int32)`

SetAllowedDriverIds sets AllowedDriverIds field to given value.

### HasAllowedDriverIds

`func (o *ActivityModel) HasAllowedDriverIds() bool`

HasAllowedDriverIds returns a boolean if a field has been set.

### GetAllowedDrivers

`func (o *ActivityModel) GetAllowedDrivers() []AllowedDriverModel`

GetAllowedDrivers returns the AllowedDrivers field if non-nil, zero value otherwise.

### GetAllowedDriversOk

`func (o *ActivityModel) GetAllowedDriversOk() (*[]AllowedDriverModel, bool)`

GetAllowedDriversOk returns a tuple with the AllowedDrivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDrivers

`func (o *ActivityModel) SetAllowedDrivers(v []AllowedDriverModel)`

SetAllowedDrivers sets AllowedDrivers field to given value.

### HasAllowedDrivers

`func (o *ActivityModel) HasAllowedDrivers() bool`

HasAllowedDrivers returns a boolean if a field has been set.

### GetAllowedDriversLinks

`func (o *ActivityModel) GetAllowedDriversLinks() []LinkModel`

GetAllowedDriversLinks returns the AllowedDriversLinks field if non-nil, zero value otherwise.

### GetAllowedDriversLinksOk

`func (o *ActivityModel) GetAllowedDriversLinksOk() (*[]LinkModel, bool)`

GetAllowedDriversLinksOk returns a tuple with the AllowedDriversLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDriversLinks

`func (o *ActivityModel) SetAllowedDriversLinks(v []LinkModel)`

SetAllowedDriversLinks sets AllowedDriversLinks field to given value.

### HasAllowedDriversLinks

`func (o *ActivityModel) HasAllowedDriversLinks() bool`

HasAllowedDriversLinks returns a boolean if a field has been set.

### GetAppliedCapacities

`func (o *ActivityModel) GetAppliedCapacities() AppliedCapacitiesModel`

GetAppliedCapacities returns the AppliedCapacities field if non-nil, zero value otherwise.

### GetAppliedCapacitiesOk

`func (o *ActivityModel) GetAppliedCapacitiesOk() (*AppliedCapacitiesModel, bool)`

GetAppliedCapacitiesOk returns a tuple with the AppliedCapacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedCapacities

`func (o *ActivityModel) SetAppliedCapacities(v AppliedCapacitiesModel)`

SetAppliedCapacities sets AppliedCapacities field to given value.

### HasAppliedCapacities

`func (o *ActivityModel) HasAppliedCapacities() bool`

HasAppliedCapacities returns a boolean if a field has been set.

### GetCapacities

`func (o *ActivityModel) GetCapacities() []CapacityModel`

GetCapacities returns the Capacities field if non-nil, zero value otherwise.

### GetCapacitiesOk

`func (o *ActivityModel) GetCapacitiesOk() (*[]CapacityModel, bool)`

GetCapacitiesOk returns a tuple with the Capacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacities

`func (o *ActivityModel) SetCapacities(v []CapacityModel)`

SetCapacities sets Capacities field to given value.

### HasCapacities

`func (o *ActivityModel) HasCapacities() bool`

HasCapacities returns a boolean if a field has been set.

### GetPackageLines

`func (o *ActivityModel) GetPackageLines() []PackageLineModel`

GetPackageLines returns the PackageLines field if non-nil, zero value otherwise.

### GetPackageLinesOk

`func (o *ActivityModel) GetPackageLinesOk() (*[]PackageLineModel, bool)`

GetPackageLinesOk returns a tuple with the PackageLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLines

`func (o *ActivityModel) SetPackageLines(v []PackageLineModel)`

SetPackageLines sets PackageLines field to given value.

### HasPackageLines

`func (o *ActivityModel) HasPackageLines() bool`

HasPackageLines returns a boolean if a field has been set.

### GetPayments

`func (o *ActivityModel) GetPayments() []PaymentModel`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *ActivityModel) GetPaymentsOk() (*[]PaymentModel, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *ActivityModel) SetPayments(v []PaymentModel)`

SetPayments sets Payments field to given value.

### HasPayments

`func (o *ActivityModel) HasPayments() bool`

HasPayments returns a boolean if a field has been set.

### GetTimeSlots

`func (o *ActivityModel) GetTimeSlots() []TimeSlotModel`

GetTimeSlots returns the TimeSlots field if non-nil, zero value otherwise.

### GetTimeSlotsOk

`func (o *ActivityModel) GetTimeSlotsOk() (*[]TimeSlotModel, bool)`

GetTimeSlotsOk returns a tuple with the TimeSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSlots

`func (o *ActivityModel) SetTimeSlots(v []TimeSlotModel)`

SetTimeSlots sets TimeSlots field to given value.

### HasTimeSlots

`func (o *ActivityModel) HasTimeSlots() bool`

HasTimeSlots returns a boolean if a field has been set.

### GetBrandId

`func (o *ActivityModel) GetBrandId() int32`

GetBrandId returns the BrandId field if non-nil, zero value otherwise.

### GetBrandIdOk

`func (o *ActivityModel) GetBrandIdOk() (*int32, bool)`

GetBrandIdOk returns a tuple with the BrandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandId

`func (o *ActivityModel) SetBrandId(v int32)`

SetBrandId sets BrandId field to given value.

### HasBrandId

`func (o *ActivityModel) HasBrandId() bool`

HasBrandId returns a boolean if a field has been set.

### GetBrandName

`func (o *ActivityModel) GetBrandName() string`

GetBrandName returns the BrandName field if non-nil, zero value otherwise.

### GetBrandNameOk

`func (o *ActivityModel) GetBrandNameOk() (*string, bool)`

GetBrandNameOk returns a tuple with the BrandName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandName

`func (o *ActivityModel) SetBrandName(v string)`

SetBrandName sets BrandName field to given value.

### HasBrandName

`func (o *ActivityModel) HasBrandName() bool`

HasBrandName returns a boolean if a field has been set.

### GetBrand

`func (o *ActivityModel) GetBrand() BrandModel`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *ActivityModel) GetBrandOk() (*BrandModel, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *ActivityModel) SetBrand(v BrandModel)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *ActivityModel) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetCommunication

`func (o *ActivityModel) GetCommunication() CommunicationModel`

GetCommunication returns the Communication field if non-nil, zero value otherwise.

### GetCommunicationOk

`func (o *ActivityModel) GetCommunicationOk() (*CommunicationModel, bool)`

GetCommunicationOk returns a tuple with the Communication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunication

`func (o *ActivityModel) SetCommunication(v CommunicationModel)`

SetCommunication sets Communication field to given value.

### HasCommunication

`func (o *ActivityModel) HasCommunication() bool`

HasCommunication returns a boolean if a field has been set.

### GetAssignmentLink

`func (o *ActivityModel) GetAssignmentLink() LinkModel`

GetAssignmentLink returns the AssignmentLink field if non-nil, zero value otherwise.

### GetAssignmentLinkOk

`func (o *ActivityModel) GetAssignmentLinkOk() (*LinkModel, bool)`

GetAssignmentLinkOk returns a tuple with the AssignmentLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentLink

`func (o *ActivityModel) SetAssignmentLink(v LinkModel)`

SetAssignmentLink sets AssignmentLink field to given value.

### HasAssignmentLink

`func (o *ActivityModel) HasAssignmentLink() bool`

HasAssignmentLink returns a boolean if a field has been set.

### GetRouteLink

`func (o *ActivityModel) GetRouteLink() LinkModel`

GetRouteLink returns the RouteLink field if non-nil, zero value otherwise.

### GetRouteLinkOk

`func (o *ActivityModel) GetRouteLinkOk() (*LinkModel, bool)`

GetRouteLinkOk returns a tuple with the RouteLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteLink

`func (o *ActivityModel) SetRouteLink(v LinkModel)`

SetRouteLink sets RouteLink field to given value.

### HasRouteLink

`func (o *ActivityModel) HasRouteLink() bool`

HasRouteLink returns a boolean if a field has been set.

### GetRoute

`func (o *ActivityModel) GetRoute() RouteModel`

GetRoute returns the Route field if non-nil, zero value otherwise.

### GetRouteOk

`func (o *ActivityModel) GetRouteOk() (*RouteModel, bool)`

GetRouteOk returns a tuple with the Route field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoute

`func (o *ActivityModel) SetRoute(v RouteModel)`

SetRoute sets Route field to given value.

### HasRoute

`func (o *ActivityModel) HasRoute() bool`

HasRoute returns a boolean if a field has been set.

### GetDriver

`func (o *ActivityModel) GetDriver() DriverModel`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *ActivityModel) GetDriverOk() (*DriverModel, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *ActivityModel) SetDriver(v DriverModel)`

SetDriver sets Driver field to given value.

### HasDriver

`func (o *ActivityModel) HasDriver() bool`

HasDriver returns a boolean if a field has been set.

### GetDriverLinks

`func (o *ActivityModel) GetDriverLinks() []LinkModel`

GetDriverLinks returns the DriverLinks field if non-nil, zero value otherwise.

### GetDriverLinksOk

`func (o *ActivityModel) GetDriverLinksOk() (*[]LinkModel, bool)`

GetDriverLinksOk returns a tuple with the DriverLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverLinks

`func (o *ActivityModel) SetDriverLinks(v []LinkModel)`

SetDriverLinks sets DriverLinks field to given value.

### HasDriverLinks

`func (o *ActivityModel) HasDriverLinks() bool`

HasDriverLinks returns a boolean if a field has been set.

### GetCar

`func (o *ActivityModel) GetCar() VehicleModel`

GetCar returns the Car field if non-nil, zero value otherwise.

### GetCarOk

`func (o *ActivityModel) GetCarOk() (*VehicleModel, bool)`

GetCarOk returns a tuple with the Car field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCar

`func (o *ActivityModel) SetCar(v VehicleModel)`

SetCar sets Car field to given value.

### HasCar

`func (o *ActivityModel) HasCar() bool`

HasCar returns a boolean if a field has been set.

### GetVehicle

`func (o *ActivityModel) GetVehicle() VehicleModel`

GetVehicle returns the Vehicle field if non-nil, zero value otherwise.

### GetVehicleOk

`func (o *ActivityModel) GetVehicleOk() (*VehicleModel, bool)`

GetVehicleOk returns a tuple with the Vehicle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicle

`func (o *ActivityModel) SetVehicle(v VehicleModel)`

SetVehicle sets Vehicle field to given value.

### HasVehicle

`func (o *ActivityModel) HasVehicle() bool`

HasVehicle returns a boolean if a field has been set.

### GetTrailer

`func (o *ActivityModel) GetTrailer() TrailerModel`

GetTrailer returns the Trailer field if non-nil, zero value otherwise.

### GetTrailerOk

`func (o *ActivityModel) GetTrailerOk() (*TrailerModel, bool)`

GetTrailerOk returns a tuple with the Trailer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailer

`func (o *ActivityModel) SetTrailer(v TrailerModel)`

SetTrailer sets Trailer field to given value.

### HasTrailer

`func (o *ActivityModel) HasTrailer() bool`

HasTrailer returns a boolean if a field has been set.

### GetActivityIdBefore

`func (o *ActivityModel) GetActivityIdBefore() string`

GetActivityIdBefore returns the ActivityIdBefore field if non-nil, zero value otherwise.

### GetActivityIdBeforeOk

`func (o *ActivityModel) GetActivityIdBeforeOk() (*string, bool)`

GetActivityIdBeforeOk returns a tuple with the ActivityIdBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityIdBefore

`func (o *ActivityModel) SetActivityIdBefore(v string)`

SetActivityIdBefore sets ActivityIdBefore field to given value.

### HasActivityIdBefore

`func (o *ActivityModel) HasActivityIdBefore() bool`

HasActivityIdBefore returns a boolean if a field has been set.

### GetActivityIdAfter

`func (o *ActivityModel) GetActivityIdAfter() string`

GetActivityIdAfter returns the ActivityIdAfter field if non-nil, zero value otherwise.

### GetActivityIdAfterOk

`func (o *ActivityModel) GetActivityIdAfterOk() (*string, bool)`

GetActivityIdAfterOk returns a tuple with the ActivityIdAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityIdAfter

`func (o *ActivityModel) SetActivityIdAfter(v string)`

SetActivityIdAfter sets ActivityIdAfter field to given value.

### HasActivityIdAfter

`func (o *ActivityModel) HasActivityIdAfter() bool`

HasActivityIdAfter returns a boolean if a field has been set.

### GetBundledActivityIds

`func (o *ActivityModel) GetBundledActivityIds() []int32`

GetBundledActivityIds returns the BundledActivityIds field if non-nil, zero value otherwise.

### GetBundledActivityIdsOk

`func (o *ActivityModel) GetBundledActivityIdsOk() (*[]int32, bool)`

GetBundledActivityIdsOk returns a tuple with the BundledActivityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundledActivityIds

`func (o *ActivityModel) SetBundledActivityIds(v []int32)`

SetBundledActivityIds sets BundledActivityIds field to given value.

### HasBundledActivityIds

`func (o *ActivityModel) HasBundledActivityIds() bool`

HasBundledActivityIds returns a boolean if a field has been set.

### GetTags

`func (o *ActivityModel) GetTags() []TagModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ActivityModel) GetTagsOk() (*[]TagModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ActivityModel) SetTags(v []TagModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ActivityModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetRecurrence

`func (o *ActivityModel) GetRecurrence() RecurrenceModel`

GetRecurrence returns the Recurrence field if non-nil, zero value otherwise.

### GetRecurrenceOk

`func (o *ActivityModel) GetRecurrenceOk() (*RecurrenceModel, bool)`

GetRecurrenceOk returns a tuple with the Recurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrence

`func (o *ActivityModel) SetRecurrence(v RecurrenceModel)`

SetRecurrence sets Recurrence field to given value.

### HasRecurrence

`func (o *ActivityModel) HasRecurrence() bool`

HasRecurrence returns a boolean if a field has been set.

### GetRecurrenceId

`func (o *ActivityModel) GetRecurrenceId() int64`

GetRecurrenceId returns the RecurrenceId field if non-nil, zero value otherwise.

### GetRecurrenceIdOk

`func (o *ActivityModel) GetRecurrenceIdOk() (*int64, bool)`

GetRecurrenceIdOk returns a tuple with the RecurrenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceId

`func (o *ActivityModel) SetRecurrenceId(v int64)`

SetRecurrenceId sets RecurrenceId field to given value.

### HasRecurrenceId

`func (o *ActivityModel) HasRecurrenceId() bool`

HasRecurrenceId returns a boolean if a field has been set.

### GetTagNames

`func (o *ActivityModel) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *ActivityModel) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *ActivityModel) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.

### HasTagNames

`func (o *ActivityModel) HasTagNames() bool`

HasTagNames returns a boolean if a field has been set.

### GetTagIds

`func (o *ActivityModel) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *ActivityModel) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *ActivityModel) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *ActivityModel) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetZones

`func (o *ActivityModel) GetZones() []ZoneModel`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *ActivityModel) GetZonesOk() (*[]ZoneModel, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *ActivityModel) SetZones(v []ZoneModel)`

SetZones sets Zones field to given value.

### HasZones

`func (o *ActivityModel) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetZoneNames

`func (o *ActivityModel) GetZoneNames() []string`

GetZoneNames returns the ZoneNames field if non-nil, zero value otherwise.

### GetZoneNamesOk

`func (o *ActivityModel) GetZoneNamesOk() (*[]string, bool)`

GetZoneNamesOk returns a tuple with the ZoneNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneNames

`func (o *ActivityModel) SetZoneNames(v []string)`

SetZoneNames sets ZoneNames field to given value.

### HasZoneNames

`func (o *ActivityModel) HasZoneNames() bool`

HasZoneNames returns a boolean if a field has been set.

### GetLinks

`func (o *ActivityModel) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ActivityModel) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ActivityModel) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ActivityModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetaData

`func (o *ActivityModel) GetMetaData() []MetaDataModel`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *ActivityModel) GetMetaDataOk() (*[]MetaDataModel, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *ActivityModel) SetMetaData(v []MetaDataModel)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *ActivityModel) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetStats

`func (o *ActivityModel) GetStats() ActivityStatsModel`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ActivityModel) GetStatsOk() (*ActivityStatsModel, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ActivityModel) SetStats(v ActivityStatsModel)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ActivityModel) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetNotes

`func (o *ActivityModel) GetNotes() []NoteModel`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ActivityModel) GetNotesOk() (*[]NoteModel, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ActivityModel) SetNotes(v []NoteModel)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ActivityModel) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetFiles

`func (o *ActivityModel) GetFiles() []FileModel`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *ActivityModel) GetFilesOk() (*[]FileModel, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *ActivityModel) SetFiles(v []FileModel)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *ActivityModel) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetActivityCreatedAt

`func (o *ActivityModel) GetActivityCreatedAt() time.Time`

GetActivityCreatedAt returns the ActivityCreatedAt field if non-nil, zero value otherwise.

### GetActivityCreatedAtOk

`func (o *ActivityModel) GetActivityCreatedAtOk() (*time.Time, bool)`

GetActivityCreatedAtOk returns a tuple with the ActivityCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityCreatedAt

`func (o *ActivityModel) SetActivityCreatedAt(v time.Time)`

SetActivityCreatedAt sets ActivityCreatedAt field to given value.

### HasActivityCreatedAt

`func (o *ActivityModel) HasActivityCreatedAt() bool`

HasActivityCreatedAt returns a boolean if a field has been set.

### GetActivityUpdatedAt

`func (o *ActivityModel) GetActivityUpdatedAt() time.Time`

GetActivityUpdatedAt returns the ActivityUpdatedAt field if non-nil, zero value otherwise.

### GetActivityUpdatedAtOk

`func (o *ActivityModel) GetActivityUpdatedAtOk() (*time.Time, bool)`

GetActivityUpdatedAtOk returns a tuple with the ActivityUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityUpdatedAt

`func (o *ActivityModel) SetActivityUpdatedAt(v time.Time)`

SetActivityUpdatedAt sets ActivityUpdatedAt field to given value.

### HasActivityUpdatedAt

`func (o *ActivityModel) HasActivityUpdatedAt() bool`

HasActivityUpdatedAt returns a boolean if a field has been set.

### GetActivityCreatedBy

`func (o *ActivityModel) GetActivityCreatedBy() int32`

GetActivityCreatedBy returns the ActivityCreatedBy field if non-nil, zero value otherwise.

### GetActivityCreatedByOk

`func (o *ActivityModel) GetActivityCreatedByOk() (*int32, bool)`

GetActivityCreatedByOk returns a tuple with the ActivityCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityCreatedBy

`func (o *ActivityModel) SetActivityCreatedBy(v int32)`

SetActivityCreatedBy sets ActivityCreatedBy field to given value.

### HasActivityCreatedBy

`func (o *ActivityModel) HasActivityCreatedBy() bool`

HasActivityCreatedBy returns a boolean if a field has been set.

### GetActivityUpdatedBy

`func (o *ActivityModel) GetActivityUpdatedBy() int32`

GetActivityUpdatedBy returns the ActivityUpdatedBy field if non-nil, zero value otherwise.

### GetActivityUpdatedByOk

`func (o *ActivityModel) GetActivityUpdatedByOk() (*int32, bool)`

GetActivityUpdatedByOk returns a tuple with the ActivityUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityUpdatedBy

`func (o *ActivityModel) SetActivityUpdatedBy(v int32)`

SetActivityUpdatedBy sets ActivityUpdatedBy field to given value.

### HasActivityUpdatedBy

`func (o *ActivityModel) HasActivityUpdatedBy() bool`

HasActivityUpdatedBy returns a boolean if a field has been set.

### GetActivityCreatedByUser

`func (o *ActivityModel) GetActivityCreatedByUser() UsersModel`

GetActivityCreatedByUser returns the ActivityCreatedByUser field if non-nil, zero value otherwise.

### GetActivityCreatedByUserOk

`func (o *ActivityModel) GetActivityCreatedByUserOk() (*UsersModel, bool)`

GetActivityCreatedByUserOk returns a tuple with the ActivityCreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityCreatedByUser

`func (o *ActivityModel) SetActivityCreatedByUser(v UsersModel)`

SetActivityCreatedByUser sets ActivityCreatedByUser field to given value.

### HasActivityCreatedByUser

`func (o *ActivityModel) HasActivityCreatedByUser() bool`

HasActivityCreatedByUser returns a boolean if a field has been set.

### GetActivityUpdatedByUser

`func (o *ActivityModel) GetActivityUpdatedByUser() UsersModel`

GetActivityUpdatedByUser returns the ActivityUpdatedByUser field if non-nil, zero value otherwise.

### GetActivityUpdatedByUserOk

`func (o *ActivityModel) GetActivityUpdatedByUserOk() (*UsersModel, bool)`

GetActivityUpdatedByUserOk returns a tuple with the ActivityUpdatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityUpdatedByUser

`func (o *ActivityModel) SetActivityUpdatedByUser(v UsersModel)`

SetActivityUpdatedByUser sets ActivityUpdatedByUser field to given value.

### HasActivityUpdatedByUser

`func (o *ActivityModel) HasActivityUpdatedByUser() bool`

HasActivityUpdatedByUser returns a boolean if a field has been set.

### GetActivityActive

`func (o *ActivityModel) GetActivityActive() bool`

GetActivityActive returns the ActivityActive field if non-nil, zero value otherwise.

### GetActivityActiveOk

`func (o *ActivityModel) GetActivityActiveOk() (*bool, bool)`

GetActivityActiveOk returns a tuple with the ActivityActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityActive

`func (o *ActivityModel) SetActivityActive(v bool)`

SetActivityActive sets ActivityActive field to given value.

### HasActivityActive

`func (o *ActivityModel) HasActivityActive() bool`

HasActivityActive returns a boolean if a field has been set.

### GetActivityRemoved

`func (o *ActivityModel) GetActivityRemoved() bool`

GetActivityRemoved returns the ActivityRemoved field if non-nil, zero value otherwise.

### GetActivityRemovedOk

`func (o *ActivityModel) GetActivityRemovedOk() (*bool, bool)`

GetActivityRemovedOk returns a tuple with the ActivityRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityRemoved

`func (o *ActivityModel) SetActivityRemoved(v bool)`

SetActivityRemoved sets ActivityRemoved field to given value.

### HasActivityRemoved

`func (o *ActivityModel) HasActivityRemoved() bool`

HasActivityRemoved returns a boolean if a field has been set.

### GetPaymentTotal

`func (o *ActivityModel) GetPaymentTotal() int32`

GetPaymentTotal returns the PaymentTotal field if non-nil, zero value otherwise.

### GetPaymentTotalOk

`func (o *ActivityModel) GetPaymentTotalOk() (*int32, bool)`

GetPaymentTotalOk returns a tuple with the PaymentTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTotal

`func (o *ActivityModel) SetPaymentTotal(v int32)`

SetPaymentTotal sets PaymentTotal field to given value.

### HasPaymentTotal

`func (o *ActivityModel) HasPaymentTotal() bool`

HasPaymentTotal returns a boolean if a field has been set.

### GetTransactionTotal

`func (o *ActivityModel) GetTransactionTotal() int32`

GetTransactionTotal returns the TransactionTotal field if non-nil, zero value otherwise.

### GetTransactionTotalOk

`func (o *ActivityModel) GetTransactionTotalOk() (*int32, bool)`

GetTransactionTotalOk returns a tuple with the TransactionTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTotal

`func (o *ActivityModel) SetTransactionTotal(v int32)`

SetTransactionTotal sets TransactionTotal field to given value.

### HasTransactionTotal

`func (o *ActivityModel) HasTransactionTotal() bool`

HasTransactionTotal returns a boolean if a field has been set.

### GetDueTotal

`func (o *ActivityModel) GetDueTotal() int32`

GetDueTotal returns the DueTotal field if non-nil, zero value otherwise.

### GetDueTotalOk

`func (o *ActivityModel) GetDueTotalOk() (*int32, bool)`

GetDueTotalOk returns a tuple with the DueTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTotal

`func (o *ActivityModel) SetDueTotal(v int32)`

SetDueTotal sets DueTotal field to given value.

### HasDueTotal

`func (o *ActivityModel) HasDueTotal() bool`

HasDueTotal returns a boolean if a field has been set.

### GetTransactions

`func (o *ActivityModel) GetTransactions() []TransactionModel`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *ActivityModel) GetTransactionsOk() (*[]TransactionModel, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *ActivityModel) SetTransactions(v []TransactionModel)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *ActivityModel) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


