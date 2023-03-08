# RouteEtaModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivitiesEta** | Pointer to [**[]ActivityEtaModel**](ActivityEtaModel.md) |  | [optional] 
**Id** | Pointer to **int64** | Unique ID of Route | [optional] [readonly] 
**Nr** | Pointer to **string** | Non-Unique number of Route | [optional] 
**Name** | Pointer to **string** | Description | [optional] 
**StatusName** | Pointer to **string** | Route Status | [optional] 
**StatusId** | Pointer to **int64** | Status ID of Route, 29:route_cancelled, 1:route_planned, 2:route_in_progress, 8:route_executed | [optional] 
**NrOfStops** | Pointer to **int32** | number of stops on this route (excluding start_route and end_route activities) | [optional] [readonly] 
**PauseId** | Pointer to **int32** | id of pause scheme to apply | [optional] 
**Pause** | Pointer to [**PauseModel**](PauseModel.md) |  | [optional] 
**DriverId** | Pointer to **int32** | id of executing driver | [optional] 
**DriverLink** | Pointer to [**LinkModel**](LinkModel.md) |  | [optional] 
**DriverLinks** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**Driver** | Pointer to [**UsersModel**](UsersModel.md) |  | [optional] 
**UserLink** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**CoDriverIds** | Pointer to **[]int32** | Unique Identifier(s) for co-drivers on route | [optional] 
**CoDrivers** | Pointer to [**[]UsersModel**](UsersModel.md) | list of co-drivers on route | [optional] [readonly] 
**CarId** | Pointer to **int32** |  | [optional] 
**VehicleId** | Pointer to **int32** | Bumbal internal id for vehicle associated with this route | [optional] 
**VehicleLink** | Pointer to [**[]LinkModel**](LinkModel.md) | link object to identify a vehicle uniquely by an external id and provider name | [optional] 
**Vehicle** | Pointer to [**VehicleModel**](VehicleModel.md) |  | [optional] 
**CarLink** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**Car** | Pointer to [**VehicleModel**](VehicleModel.md) |  | [optional] 
**TrailerId** | Pointer to **int32** |  | [optional] 
**TrailerLink** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**Trailer** | Pointer to [**TrailerModel**](TrailerModel.md) |  | [optional] 
**Equipment** | Pointer to [**[]EquipmentModel**](EquipmentModel.md) |  | [optional] [readonly] 
**EquipmentIds** | Pointer to **[]int32** | Unique Identifier(s) for equipment on route | [optional] 
**PlannedStartDuration** | Pointer to **int32** | Duration for start activity | [optional] 
**PlannedEndDuration** | Pointer to **int32** | Duration for end activity | [optional] 
**MaxSpeed** | Pointer to **int64** | Max Speed in km/h | [optional] 
**SpeedFactor** | Pointer to **float64** | Speed Factor | [optional] 
**DurationFactor** | Pointer to **float64** | Duration Factor | [optional] 
**CostPerMeter** | Pointer to **float64** | Cost per meter | [optional] 
**CostPerRoute** | Pointer to **float64** | Cost per route | [optional] 
**CostPerDrivingMinute** | Pointer to **float64** | Cost per driving minute | [optional] 
**CostPerWaitingMinute** | Pointer to **float64** | Cost per waiting minute | [optional] 
**CostPerServiceMinute** | Pointer to **float64** | Cost per service minute | [optional] 
**EarliestDate** | Pointer to **string** | Write only! | [optional] 
**LatestDate** | Pointer to **string** | Write only! | [optional] 
**EarliestDateTime** | Pointer to **time.Time** |  | [optional] 
**LatestDateTime** | Pointer to **time.Time** |  | [optional] 
**PlannedDateTimeFrom** | Pointer to **time.Time** | planned date-time from | [optional] [readonly] 
**PlannedDateTimeTo** | Pointer to **time.Time** | planned date-time to | [optional] [readonly] 
**ExecutedDateTimeFrom** | Pointer to **time.Time** | executed date-time from | [optional] [readonly] 
**ExecutedDateTimeTo** | Pointer to **time.Time** | executed date-time to | [optional] [readonly] 
**PlannedDrivingDistance** | Pointer to **int32** | Planned driving distance of this route in meters | [optional] 
**PlannedDrivingDuration** | Pointer to **int32** | Planned driving duration of this route in minutes | [optional] 
**PlannedWaitingDuration** | Pointer to **int32** | Planned waiting duration of this route in minutes | [optional] 
**PlannedActivityDuration** | Pointer to **int32** | Planned duration for all activities in this route in minutes | [optional] 
**PlannedTotalDuration** | Pointer to **int32** | Total planned duration of this route in minutes | [optional] 
**GpsLocations** | Pointer to [**[]GPSLocationModel**](GPSLocationModel.md) |  | [optional] [readonly] 
**LatestKnownPosition** | Pointer to [**GPSLocationModel**](GPSLocationModel.md) |  | [optional] 
**RecurrenceId** | Pointer to **int32** | id of recurrence where route belongs to | [optional] [readonly] 
**RecurrenceNr** | Pointer to **int32** | nr within recurrence where route belongs to | [optional] [readonly] 
**Recurrence** | Pointer to [**RecurrenceModel**](RecurrenceModel.md) |  | [optional] 
**Overdue** | Pointer to **bool** | whether any activity on route is overdue | [optional] [readonly] 
**Optimized** | Pointer to **bool** | Activity optimized status within route. | [optional] 
**Blocked** | Pointer to **bool** | a blocked route can not be auto-filled by customer calendars | [optional] [readonly] 
**Active** | Pointer to **bool** | if active&#x3D;0: route has been removed and is no longer visible in any bumbal interface | [optional] 
**StartAddress** | Pointer to [**AddressModel**](AddressModel.md) |  | [optional] 
**EndAddress** | Pointer to [**AddressModel**](AddressModel.md) |  | [optional] 
**PlannedCapacities** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**AppliedCapacities** | Pointer to [**AppliedCapacitiesModel**](AppliedCapacitiesModel.md) |  | [optional] 
**Capacities** | Pointer to [**[]CapacityModel**](CapacityModel.md) |  | [optional] 
**ActivityIds** | Pointer to **[]int32** | activity ids on route in order of execution | [optional] 
**Links** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**MetaData** | Pointer to [**[]MetaDataModel**](MetaDataModel.md) |  | [optional] 
**Notes** | Pointer to [**[]NoteModel**](NoteModel.md) |  | [optional] 
**Files** | Pointer to [**[]FileModel**](FileModel.md) |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | created_at date time | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | updated_at date time | [optional] [readonly] 
**Tags** | Pointer to [**[]TagModel**](TagModel.md) |  | [optional] 
**TagNames** | Pointer to **[]string** | Tag names | [optional] 
**Zones** | Pointer to [**[]ZoneModel**](ZoneModel.md) |  | [optional] 
**ZoneNames** | Pointer to **[]string** | Zone names | [optional] 

## Methods

### NewRouteEtaModel

`func NewRouteEtaModel() *RouteEtaModel`

NewRouteEtaModel instantiates a new RouteEtaModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteEtaModelWithDefaults

`func NewRouteEtaModelWithDefaults() *RouteEtaModel`

NewRouteEtaModelWithDefaults instantiates a new RouteEtaModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivitiesEta

`func (o *RouteEtaModel) GetActivitiesEta() []ActivityEtaModel`

GetActivitiesEta returns the ActivitiesEta field if non-nil, zero value otherwise.

### GetActivitiesEtaOk

`func (o *RouteEtaModel) GetActivitiesEtaOk() (*[]ActivityEtaModel, bool)`

GetActivitiesEtaOk returns a tuple with the ActivitiesEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivitiesEta

`func (o *RouteEtaModel) SetActivitiesEta(v []ActivityEtaModel)`

SetActivitiesEta sets ActivitiesEta field to given value.

### HasActivitiesEta

`func (o *RouteEtaModel) HasActivitiesEta() bool`

HasActivitiesEta returns a boolean if a field has been set.

### GetId

`func (o *RouteEtaModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RouteEtaModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RouteEtaModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RouteEtaModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNr

`func (o *RouteEtaModel) GetNr() string`

GetNr returns the Nr field if non-nil, zero value otherwise.

### GetNrOk

`func (o *RouteEtaModel) GetNrOk() (*string, bool)`

GetNrOk returns a tuple with the Nr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNr

`func (o *RouteEtaModel) SetNr(v string)`

SetNr sets Nr field to given value.

### HasNr

`func (o *RouteEtaModel) HasNr() bool`

HasNr returns a boolean if a field has been set.

### GetName

`func (o *RouteEtaModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RouteEtaModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RouteEtaModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RouteEtaModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatusName

`func (o *RouteEtaModel) GetStatusName() string`

GetStatusName returns the StatusName field if non-nil, zero value otherwise.

### GetStatusNameOk

`func (o *RouteEtaModel) GetStatusNameOk() (*string, bool)`

GetStatusNameOk returns a tuple with the StatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusName

`func (o *RouteEtaModel) SetStatusName(v string)`

SetStatusName sets StatusName field to given value.

### HasStatusName

`func (o *RouteEtaModel) HasStatusName() bool`

HasStatusName returns a boolean if a field has been set.

### GetStatusId

`func (o *RouteEtaModel) GetStatusId() int64`

GetStatusId returns the StatusId field if non-nil, zero value otherwise.

### GetStatusIdOk

`func (o *RouteEtaModel) GetStatusIdOk() (*int64, bool)`

GetStatusIdOk returns a tuple with the StatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusId

`func (o *RouteEtaModel) SetStatusId(v int64)`

SetStatusId sets StatusId field to given value.

### HasStatusId

`func (o *RouteEtaModel) HasStatusId() bool`

HasStatusId returns a boolean if a field has been set.

### GetNrOfStops

`func (o *RouteEtaModel) GetNrOfStops() int32`

GetNrOfStops returns the NrOfStops field if non-nil, zero value otherwise.

### GetNrOfStopsOk

`func (o *RouteEtaModel) GetNrOfStopsOk() (*int32, bool)`

GetNrOfStopsOk returns a tuple with the NrOfStops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrOfStops

`func (o *RouteEtaModel) SetNrOfStops(v int32)`

SetNrOfStops sets NrOfStops field to given value.

### HasNrOfStops

`func (o *RouteEtaModel) HasNrOfStops() bool`

HasNrOfStops returns a boolean if a field has been set.

### GetPauseId

`func (o *RouteEtaModel) GetPauseId() int32`

GetPauseId returns the PauseId field if non-nil, zero value otherwise.

### GetPauseIdOk

`func (o *RouteEtaModel) GetPauseIdOk() (*int32, bool)`

GetPauseIdOk returns a tuple with the PauseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseId

`func (o *RouteEtaModel) SetPauseId(v int32)`

SetPauseId sets PauseId field to given value.

### HasPauseId

`func (o *RouteEtaModel) HasPauseId() bool`

HasPauseId returns a boolean if a field has been set.

### GetPause

`func (o *RouteEtaModel) GetPause() PauseModel`

GetPause returns the Pause field if non-nil, zero value otherwise.

### GetPauseOk

`func (o *RouteEtaModel) GetPauseOk() (*PauseModel, bool)`

GetPauseOk returns a tuple with the Pause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPause

`func (o *RouteEtaModel) SetPause(v PauseModel)`

SetPause sets Pause field to given value.

### HasPause

`func (o *RouteEtaModel) HasPause() bool`

HasPause returns a boolean if a field has been set.

### GetDriverId

`func (o *RouteEtaModel) GetDriverId() int32`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *RouteEtaModel) GetDriverIdOk() (*int32, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *RouteEtaModel) SetDriverId(v int32)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *RouteEtaModel) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### GetDriverLink

`func (o *RouteEtaModel) GetDriverLink() LinkModel`

GetDriverLink returns the DriverLink field if non-nil, zero value otherwise.

### GetDriverLinkOk

`func (o *RouteEtaModel) GetDriverLinkOk() (*LinkModel, bool)`

GetDriverLinkOk returns a tuple with the DriverLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverLink

`func (o *RouteEtaModel) SetDriverLink(v LinkModel)`

SetDriverLink sets DriverLink field to given value.

### HasDriverLink

`func (o *RouteEtaModel) HasDriverLink() bool`

HasDriverLink returns a boolean if a field has been set.

### GetDriverLinks

`func (o *RouteEtaModel) GetDriverLinks() []LinkModel`

GetDriverLinks returns the DriverLinks field if non-nil, zero value otherwise.

### GetDriverLinksOk

`func (o *RouteEtaModel) GetDriverLinksOk() (*[]LinkModel, bool)`

GetDriverLinksOk returns a tuple with the DriverLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverLinks

`func (o *RouteEtaModel) SetDriverLinks(v []LinkModel)`

SetDriverLinks sets DriverLinks field to given value.

### HasDriverLinks

`func (o *RouteEtaModel) HasDriverLinks() bool`

HasDriverLinks returns a boolean if a field has been set.

### GetDriver

`func (o *RouteEtaModel) GetDriver() UsersModel`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *RouteEtaModel) GetDriverOk() (*UsersModel, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *RouteEtaModel) SetDriver(v UsersModel)`

SetDriver sets Driver field to given value.

### HasDriver

`func (o *RouteEtaModel) HasDriver() bool`

HasDriver returns a boolean if a field has been set.

### GetUserLink

`func (o *RouteEtaModel) GetUserLink() []LinkModel`

GetUserLink returns the UserLink field if non-nil, zero value otherwise.

### GetUserLinkOk

`func (o *RouteEtaModel) GetUserLinkOk() (*[]LinkModel, bool)`

GetUserLinkOk returns a tuple with the UserLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLink

`func (o *RouteEtaModel) SetUserLink(v []LinkModel)`

SetUserLink sets UserLink field to given value.

### HasUserLink

`func (o *RouteEtaModel) HasUserLink() bool`

HasUserLink returns a boolean if a field has been set.

### GetCoDriverIds

`func (o *RouteEtaModel) GetCoDriverIds() []int32`

GetCoDriverIds returns the CoDriverIds field if non-nil, zero value otherwise.

### GetCoDriverIdsOk

`func (o *RouteEtaModel) GetCoDriverIdsOk() (*[]int32, bool)`

GetCoDriverIdsOk returns a tuple with the CoDriverIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoDriverIds

`func (o *RouteEtaModel) SetCoDriverIds(v []int32)`

SetCoDriverIds sets CoDriverIds field to given value.

### HasCoDriverIds

`func (o *RouteEtaModel) HasCoDriverIds() bool`

HasCoDriverIds returns a boolean if a field has been set.

### GetCoDrivers

`func (o *RouteEtaModel) GetCoDrivers() []UsersModel`

GetCoDrivers returns the CoDrivers field if non-nil, zero value otherwise.

### GetCoDriversOk

`func (o *RouteEtaModel) GetCoDriversOk() (*[]UsersModel, bool)`

GetCoDriversOk returns a tuple with the CoDrivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoDrivers

`func (o *RouteEtaModel) SetCoDrivers(v []UsersModel)`

SetCoDrivers sets CoDrivers field to given value.

### HasCoDrivers

`func (o *RouteEtaModel) HasCoDrivers() bool`

HasCoDrivers returns a boolean if a field has been set.

### GetCarId

`func (o *RouteEtaModel) GetCarId() int32`

GetCarId returns the CarId field if non-nil, zero value otherwise.

### GetCarIdOk

`func (o *RouteEtaModel) GetCarIdOk() (*int32, bool)`

GetCarIdOk returns a tuple with the CarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarId

`func (o *RouteEtaModel) SetCarId(v int32)`

SetCarId sets CarId field to given value.

### HasCarId

`func (o *RouteEtaModel) HasCarId() bool`

HasCarId returns a boolean if a field has been set.

### GetVehicleId

`func (o *RouteEtaModel) GetVehicleId() int32`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *RouteEtaModel) GetVehicleIdOk() (*int32, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *RouteEtaModel) SetVehicleId(v int32)`

SetVehicleId sets VehicleId field to given value.

### HasVehicleId

`func (o *RouteEtaModel) HasVehicleId() bool`

HasVehicleId returns a boolean if a field has been set.

### GetVehicleLink

`func (o *RouteEtaModel) GetVehicleLink() []LinkModel`

GetVehicleLink returns the VehicleLink field if non-nil, zero value otherwise.

### GetVehicleLinkOk

`func (o *RouteEtaModel) GetVehicleLinkOk() (*[]LinkModel, bool)`

GetVehicleLinkOk returns a tuple with the VehicleLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleLink

`func (o *RouteEtaModel) SetVehicleLink(v []LinkModel)`

SetVehicleLink sets VehicleLink field to given value.

### HasVehicleLink

`func (o *RouteEtaModel) HasVehicleLink() bool`

HasVehicleLink returns a boolean if a field has been set.

### GetVehicle

`func (o *RouteEtaModel) GetVehicle() VehicleModel`

GetVehicle returns the Vehicle field if non-nil, zero value otherwise.

### GetVehicleOk

`func (o *RouteEtaModel) GetVehicleOk() (*VehicleModel, bool)`

GetVehicleOk returns a tuple with the Vehicle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicle

`func (o *RouteEtaModel) SetVehicle(v VehicleModel)`

SetVehicle sets Vehicle field to given value.

### HasVehicle

`func (o *RouteEtaModel) HasVehicle() bool`

HasVehicle returns a boolean if a field has been set.

### GetCarLink

`func (o *RouteEtaModel) GetCarLink() []LinkModel`

GetCarLink returns the CarLink field if non-nil, zero value otherwise.

### GetCarLinkOk

`func (o *RouteEtaModel) GetCarLinkOk() (*[]LinkModel, bool)`

GetCarLinkOk returns a tuple with the CarLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarLink

`func (o *RouteEtaModel) SetCarLink(v []LinkModel)`

SetCarLink sets CarLink field to given value.

### HasCarLink

`func (o *RouteEtaModel) HasCarLink() bool`

HasCarLink returns a boolean if a field has been set.

### GetCar

`func (o *RouteEtaModel) GetCar() VehicleModel`

GetCar returns the Car field if non-nil, zero value otherwise.

### GetCarOk

`func (o *RouteEtaModel) GetCarOk() (*VehicleModel, bool)`

GetCarOk returns a tuple with the Car field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCar

`func (o *RouteEtaModel) SetCar(v VehicleModel)`

SetCar sets Car field to given value.

### HasCar

`func (o *RouteEtaModel) HasCar() bool`

HasCar returns a boolean if a field has been set.

### GetTrailerId

`func (o *RouteEtaModel) GetTrailerId() int32`

GetTrailerId returns the TrailerId field if non-nil, zero value otherwise.

### GetTrailerIdOk

`func (o *RouteEtaModel) GetTrailerIdOk() (*int32, bool)`

GetTrailerIdOk returns a tuple with the TrailerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerId

`func (o *RouteEtaModel) SetTrailerId(v int32)`

SetTrailerId sets TrailerId field to given value.

### HasTrailerId

`func (o *RouteEtaModel) HasTrailerId() bool`

HasTrailerId returns a boolean if a field has been set.

### GetTrailerLink

`func (o *RouteEtaModel) GetTrailerLink() []LinkModel`

GetTrailerLink returns the TrailerLink field if non-nil, zero value otherwise.

### GetTrailerLinkOk

`func (o *RouteEtaModel) GetTrailerLinkOk() (*[]LinkModel, bool)`

GetTrailerLinkOk returns a tuple with the TrailerLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerLink

`func (o *RouteEtaModel) SetTrailerLink(v []LinkModel)`

SetTrailerLink sets TrailerLink field to given value.

### HasTrailerLink

`func (o *RouteEtaModel) HasTrailerLink() bool`

HasTrailerLink returns a boolean if a field has been set.

### GetTrailer

`func (o *RouteEtaModel) GetTrailer() TrailerModel`

GetTrailer returns the Trailer field if non-nil, zero value otherwise.

### GetTrailerOk

`func (o *RouteEtaModel) GetTrailerOk() (*TrailerModel, bool)`

GetTrailerOk returns a tuple with the Trailer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailer

`func (o *RouteEtaModel) SetTrailer(v TrailerModel)`

SetTrailer sets Trailer field to given value.

### HasTrailer

`func (o *RouteEtaModel) HasTrailer() bool`

HasTrailer returns a boolean if a field has been set.

### GetEquipment

`func (o *RouteEtaModel) GetEquipment() []EquipmentModel`

GetEquipment returns the Equipment field if non-nil, zero value otherwise.

### GetEquipmentOk

`func (o *RouteEtaModel) GetEquipmentOk() (*[]EquipmentModel, bool)`

GetEquipmentOk returns a tuple with the Equipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipment

`func (o *RouteEtaModel) SetEquipment(v []EquipmentModel)`

SetEquipment sets Equipment field to given value.

### HasEquipment

`func (o *RouteEtaModel) HasEquipment() bool`

HasEquipment returns a boolean if a field has been set.

### GetEquipmentIds

`func (o *RouteEtaModel) GetEquipmentIds() []int32`

GetEquipmentIds returns the EquipmentIds field if non-nil, zero value otherwise.

### GetEquipmentIdsOk

`func (o *RouteEtaModel) GetEquipmentIdsOk() (*[]int32, bool)`

GetEquipmentIdsOk returns a tuple with the EquipmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentIds

`func (o *RouteEtaModel) SetEquipmentIds(v []int32)`

SetEquipmentIds sets EquipmentIds field to given value.

### HasEquipmentIds

`func (o *RouteEtaModel) HasEquipmentIds() bool`

HasEquipmentIds returns a boolean if a field has been set.

### GetPlannedStartDuration

`func (o *RouteEtaModel) GetPlannedStartDuration() int32`

GetPlannedStartDuration returns the PlannedStartDuration field if non-nil, zero value otherwise.

### GetPlannedStartDurationOk

`func (o *RouteEtaModel) GetPlannedStartDurationOk() (*int32, bool)`

GetPlannedStartDurationOk returns a tuple with the PlannedStartDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedStartDuration

`func (o *RouteEtaModel) SetPlannedStartDuration(v int32)`

SetPlannedStartDuration sets PlannedStartDuration field to given value.

### HasPlannedStartDuration

`func (o *RouteEtaModel) HasPlannedStartDuration() bool`

HasPlannedStartDuration returns a boolean if a field has been set.

### GetPlannedEndDuration

`func (o *RouteEtaModel) GetPlannedEndDuration() int32`

GetPlannedEndDuration returns the PlannedEndDuration field if non-nil, zero value otherwise.

### GetPlannedEndDurationOk

`func (o *RouteEtaModel) GetPlannedEndDurationOk() (*int32, bool)`

GetPlannedEndDurationOk returns a tuple with the PlannedEndDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedEndDuration

`func (o *RouteEtaModel) SetPlannedEndDuration(v int32)`

SetPlannedEndDuration sets PlannedEndDuration field to given value.

### HasPlannedEndDuration

`func (o *RouteEtaModel) HasPlannedEndDuration() bool`

HasPlannedEndDuration returns a boolean if a field has been set.

### GetMaxSpeed

`func (o *RouteEtaModel) GetMaxSpeed() int64`

GetMaxSpeed returns the MaxSpeed field if non-nil, zero value otherwise.

### GetMaxSpeedOk

`func (o *RouteEtaModel) GetMaxSpeedOk() (*int64, bool)`

GetMaxSpeedOk returns a tuple with the MaxSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpeed

`func (o *RouteEtaModel) SetMaxSpeed(v int64)`

SetMaxSpeed sets MaxSpeed field to given value.

### HasMaxSpeed

`func (o *RouteEtaModel) HasMaxSpeed() bool`

HasMaxSpeed returns a boolean if a field has been set.

### GetSpeedFactor

`func (o *RouteEtaModel) GetSpeedFactor() float64`

GetSpeedFactor returns the SpeedFactor field if non-nil, zero value otherwise.

### GetSpeedFactorOk

`func (o *RouteEtaModel) GetSpeedFactorOk() (*float64, bool)`

GetSpeedFactorOk returns a tuple with the SpeedFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedFactor

`func (o *RouteEtaModel) SetSpeedFactor(v float64)`

SetSpeedFactor sets SpeedFactor field to given value.

### HasSpeedFactor

`func (o *RouteEtaModel) HasSpeedFactor() bool`

HasSpeedFactor returns a boolean if a field has been set.

### GetDurationFactor

`func (o *RouteEtaModel) GetDurationFactor() float64`

GetDurationFactor returns the DurationFactor field if non-nil, zero value otherwise.

### GetDurationFactorOk

`func (o *RouteEtaModel) GetDurationFactorOk() (*float64, bool)`

GetDurationFactorOk returns a tuple with the DurationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationFactor

`func (o *RouteEtaModel) SetDurationFactor(v float64)`

SetDurationFactor sets DurationFactor field to given value.

### HasDurationFactor

`func (o *RouteEtaModel) HasDurationFactor() bool`

HasDurationFactor returns a boolean if a field has been set.

### GetCostPerMeter

`func (o *RouteEtaModel) GetCostPerMeter() float64`

GetCostPerMeter returns the CostPerMeter field if non-nil, zero value otherwise.

### GetCostPerMeterOk

`func (o *RouteEtaModel) GetCostPerMeterOk() (*float64, bool)`

GetCostPerMeterOk returns a tuple with the CostPerMeter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerMeter

`func (o *RouteEtaModel) SetCostPerMeter(v float64)`

SetCostPerMeter sets CostPerMeter field to given value.

### HasCostPerMeter

`func (o *RouteEtaModel) HasCostPerMeter() bool`

HasCostPerMeter returns a boolean if a field has been set.

### GetCostPerRoute

`func (o *RouteEtaModel) GetCostPerRoute() float64`

GetCostPerRoute returns the CostPerRoute field if non-nil, zero value otherwise.

### GetCostPerRouteOk

`func (o *RouteEtaModel) GetCostPerRouteOk() (*float64, bool)`

GetCostPerRouteOk returns a tuple with the CostPerRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerRoute

`func (o *RouteEtaModel) SetCostPerRoute(v float64)`

SetCostPerRoute sets CostPerRoute field to given value.

### HasCostPerRoute

`func (o *RouteEtaModel) HasCostPerRoute() bool`

HasCostPerRoute returns a boolean if a field has been set.

### GetCostPerDrivingMinute

`func (o *RouteEtaModel) GetCostPerDrivingMinute() float64`

GetCostPerDrivingMinute returns the CostPerDrivingMinute field if non-nil, zero value otherwise.

### GetCostPerDrivingMinuteOk

`func (o *RouteEtaModel) GetCostPerDrivingMinuteOk() (*float64, bool)`

GetCostPerDrivingMinuteOk returns a tuple with the CostPerDrivingMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerDrivingMinute

`func (o *RouteEtaModel) SetCostPerDrivingMinute(v float64)`

SetCostPerDrivingMinute sets CostPerDrivingMinute field to given value.

### HasCostPerDrivingMinute

`func (o *RouteEtaModel) HasCostPerDrivingMinute() bool`

HasCostPerDrivingMinute returns a boolean if a field has been set.

### GetCostPerWaitingMinute

`func (o *RouteEtaModel) GetCostPerWaitingMinute() float64`

GetCostPerWaitingMinute returns the CostPerWaitingMinute field if non-nil, zero value otherwise.

### GetCostPerWaitingMinuteOk

`func (o *RouteEtaModel) GetCostPerWaitingMinuteOk() (*float64, bool)`

GetCostPerWaitingMinuteOk returns a tuple with the CostPerWaitingMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerWaitingMinute

`func (o *RouteEtaModel) SetCostPerWaitingMinute(v float64)`

SetCostPerWaitingMinute sets CostPerWaitingMinute field to given value.

### HasCostPerWaitingMinute

`func (o *RouteEtaModel) HasCostPerWaitingMinute() bool`

HasCostPerWaitingMinute returns a boolean if a field has been set.

### GetCostPerServiceMinute

`func (o *RouteEtaModel) GetCostPerServiceMinute() float64`

GetCostPerServiceMinute returns the CostPerServiceMinute field if non-nil, zero value otherwise.

### GetCostPerServiceMinuteOk

`func (o *RouteEtaModel) GetCostPerServiceMinuteOk() (*float64, bool)`

GetCostPerServiceMinuteOk returns a tuple with the CostPerServiceMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerServiceMinute

`func (o *RouteEtaModel) SetCostPerServiceMinute(v float64)`

SetCostPerServiceMinute sets CostPerServiceMinute field to given value.

### HasCostPerServiceMinute

`func (o *RouteEtaModel) HasCostPerServiceMinute() bool`

HasCostPerServiceMinute returns a boolean if a field has been set.

### GetEarliestDate

`func (o *RouteEtaModel) GetEarliestDate() string`

GetEarliestDate returns the EarliestDate field if non-nil, zero value otherwise.

### GetEarliestDateOk

`func (o *RouteEtaModel) GetEarliestDateOk() (*string, bool)`

GetEarliestDateOk returns a tuple with the EarliestDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestDate

`func (o *RouteEtaModel) SetEarliestDate(v string)`

SetEarliestDate sets EarliestDate field to given value.

### HasEarliestDate

`func (o *RouteEtaModel) HasEarliestDate() bool`

HasEarliestDate returns a boolean if a field has been set.

### GetLatestDate

`func (o *RouteEtaModel) GetLatestDate() string`

GetLatestDate returns the LatestDate field if non-nil, zero value otherwise.

### GetLatestDateOk

`func (o *RouteEtaModel) GetLatestDateOk() (*string, bool)`

GetLatestDateOk returns a tuple with the LatestDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDate

`func (o *RouteEtaModel) SetLatestDate(v string)`

SetLatestDate sets LatestDate field to given value.

### HasLatestDate

`func (o *RouteEtaModel) HasLatestDate() bool`

HasLatestDate returns a boolean if a field has been set.

### GetEarliestDateTime

`func (o *RouteEtaModel) GetEarliestDateTime() time.Time`

GetEarliestDateTime returns the EarliestDateTime field if non-nil, zero value otherwise.

### GetEarliestDateTimeOk

`func (o *RouteEtaModel) GetEarliestDateTimeOk() (*time.Time, bool)`

GetEarliestDateTimeOk returns a tuple with the EarliestDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestDateTime

`func (o *RouteEtaModel) SetEarliestDateTime(v time.Time)`

SetEarliestDateTime sets EarliestDateTime field to given value.

### HasEarliestDateTime

`func (o *RouteEtaModel) HasEarliestDateTime() bool`

HasEarliestDateTime returns a boolean if a field has been set.

### GetLatestDateTime

`func (o *RouteEtaModel) GetLatestDateTime() time.Time`

GetLatestDateTime returns the LatestDateTime field if non-nil, zero value otherwise.

### GetLatestDateTimeOk

`func (o *RouteEtaModel) GetLatestDateTimeOk() (*time.Time, bool)`

GetLatestDateTimeOk returns a tuple with the LatestDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDateTime

`func (o *RouteEtaModel) SetLatestDateTime(v time.Time)`

SetLatestDateTime sets LatestDateTime field to given value.

### HasLatestDateTime

`func (o *RouteEtaModel) HasLatestDateTime() bool`

HasLatestDateTime returns a boolean if a field has been set.

### GetPlannedDateTimeFrom

`func (o *RouteEtaModel) GetPlannedDateTimeFrom() time.Time`

GetPlannedDateTimeFrom returns the PlannedDateTimeFrom field if non-nil, zero value otherwise.

### GetPlannedDateTimeFromOk

`func (o *RouteEtaModel) GetPlannedDateTimeFromOk() (*time.Time, bool)`

GetPlannedDateTimeFromOk returns a tuple with the PlannedDateTimeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDateTimeFrom

`func (o *RouteEtaModel) SetPlannedDateTimeFrom(v time.Time)`

SetPlannedDateTimeFrom sets PlannedDateTimeFrom field to given value.

### HasPlannedDateTimeFrom

`func (o *RouteEtaModel) HasPlannedDateTimeFrom() bool`

HasPlannedDateTimeFrom returns a boolean if a field has been set.

### GetPlannedDateTimeTo

`func (o *RouteEtaModel) GetPlannedDateTimeTo() time.Time`

GetPlannedDateTimeTo returns the PlannedDateTimeTo field if non-nil, zero value otherwise.

### GetPlannedDateTimeToOk

`func (o *RouteEtaModel) GetPlannedDateTimeToOk() (*time.Time, bool)`

GetPlannedDateTimeToOk returns a tuple with the PlannedDateTimeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDateTimeTo

`func (o *RouteEtaModel) SetPlannedDateTimeTo(v time.Time)`

SetPlannedDateTimeTo sets PlannedDateTimeTo field to given value.

### HasPlannedDateTimeTo

`func (o *RouteEtaModel) HasPlannedDateTimeTo() bool`

HasPlannedDateTimeTo returns a boolean if a field has been set.

### GetExecutedDateTimeFrom

`func (o *RouteEtaModel) GetExecutedDateTimeFrom() time.Time`

GetExecutedDateTimeFrom returns the ExecutedDateTimeFrom field if non-nil, zero value otherwise.

### GetExecutedDateTimeFromOk

`func (o *RouteEtaModel) GetExecutedDateTimeFromOk() (*time.Time, bool)`

GetExecutedDateTimeFromOk returns a tuple with the ExecutedDateTimeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedDateTimeFrom

`func (o *RouteEtaModel) SetExecutedDateTimeFrom(v time.Time)`

SetExecutedDateTimeFrom sets ExecutedDateTimeFrom field to given value.

### HasExecutedDateTimeFrom

`func (o *RouteEtaModel) HasExecutedDateTimeFrom() bool`

HasExecutedDateTimeFrom returns a boolean if a field has been set.

### GetExecutedDateTimeTo

`func (o *RouteEtaModel) GetExecutedDateTimeTo() time.Time`

GetExecutedDateTimeTo returns the ExecutedDateTimeTo field if non-nil, zero value otherwise.

### GetExecutedDateTimeToOk

`func (o *RouteEtaModel) GetExecutedDateTimeToOk() (*time.Time, bool)`

GetExecutedDateTimeToOk returns a tuple with the ExecutedDateTimeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedDateTimeTo

`func (o *RouteEtaModel) SetExecutedDateTimeTo(v time.Time)`

SetExecutedDateTimeTo sets ExecutedDateTimeTo field to given value.

### HasExecutedDateTimeTo

`func (o *RouteEtaModel) HasExecutedDateTimeTo() bool`

HasExecutedDateTimeTo returns a boolean if a field has been set.

### GetPlannedDrivingDistance

`func (o *RouteEtaModel) GetPlannedDrivingDistance() int32`

GetPlannedDrivingDistance returns the PlannedDrivingDistance field if non-nil, zero value otherwise.

### GetPlannedDrivingDistanceOk

`func (o *RouteEtaModel) GetPlannedDrivingDistanceOk() (*int32, bool)`

GetPlannedDrivingDistanceOk returns a tuple with the PlannedDrivingDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDrivingDistance

`func (o *RouteEtaModel) SetPlannedDrivingDistance(v int32)`

SetPlannedDrivingDistance sets PlannedDrivingDistance field to given value.

### HasPlannedDrivingDistance

`func (o *RouteEtaModel) HasPlannedDrivingDistance() bool`

HasPlannedDrivingDistance returns a boolean if a field has been set.

### GetPlannedDrivingDuration

`func (o *RouteEtaModel) GetPlannedDrivingDuration() int32`

GetPlannedDrivingDuration returns the PlannedDrivingDuration field if non-nil, zero value otherwise.

### GetPlannedDrivingDurationOk

`func (o *RouteEtaModel) GetPlannedDrivingDurationOk() (*int32, bool)`

GetPlannedDrivingDurationOk returns a tuple with the PlannedDrivingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDrivingDuration

`func (o *RouteEtaModel) SetPlannedDrivingDuration(v int32)`

SetPlannedDrivingDuration sets PlannedDrivingDuration field to given value.

### HasPlannedDrivingDuration

`func (o *RouteEtaModel) HasPlannedDrivingDuration() bool`

HasPlannedDrivingDuration returns a boolean if a field has been set.

### GetPlannedWaitingDuration

`func (o *RouteEtaModel) GetPlannedWaitingDuration() int32`

GetPlannedWaitingDuration returns the PlannedWaitingDuration field if non-nil, zero value otherwise.

### GetPlannedWaitingDurationOk

`func (o *RouteEtaModel) GetPlannedWaitingDurationOk() (*int32, bool)`

GetPlannedWaitingDurationOk returns a tuple with the PlannedWaitingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedWaitingDuration

`func (o *RouteEtaModel) SetPlannedWaitingDuration(v int32)`

SetPlannedWaitingDuration sets PlannedWaitingDuration field to given value.

### HasPlannedWaitingDuration

`func (o *RouteEtaModel) HasPlannedWaitingDuration() bool`

HasPlannedWaitingDuration returns a boolean if a field has been set.

### GetPlannedActivityDuration

`func (o *RouteEtaModel) GetPlannedActivityDuration() int32`

GetPlannedActivityDuration returns the PlannedActivityDuration field if non-nil, zero value otherwise.

### GetPlannedActivityDurationOk

`func (o *RouteEtaModel) GetPlannedActivityDurationOk() (*int32, bool)`

GetPlannedActivityDurationOk returns a tuple with the PlannedActivityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedActivityDuration

`func (o *RouteEtaModel) SetPlannedActivityDuration(v int32)`

SetPlannedActivityDuration sets PlannedActivityDuration field to given value.

### HasPlannedActivityDuration

`func (o *RouteEtaModel) HasPlannedActivityDuration() bool`

HasPlannedActivityDuration returns a boolean if a field has been set.

### GetPlannedTotalDuration

`func (o *RouteEtaModel) GetPlannedTotalDuration() int32`

GetPlannedTotalDuration returns the PlannedTotalDuration field if non-nil, zero value otherwise.

### GetPlannedTotalDurationOk

`func (o *RouteEtaModel) GetPlannedTotalDurationOk() (*int32, bool)`

GetPlannedTotalDurationOk returns a tuple with the PlannedTotalDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedTotalDuration

`func (o *RouteEtaModel) SetPlannedTotalDuration(v int32)`

SetPlannedTotalDuration sets PlannedTotalDuration field to given value.

### HasPlannedTotalDuration

`func (o *RouteEtaModel) HasPlannedTotalDuration() bool`

HasPlannedTotalDuration returns a boolean if a field has been set.

### GetGpsLocations

`func (o *RouteEtaModel) GetGpsLocations() []GPSLocationModel`

GetGpsLocations returns the GpsLocations field if non-nil, zero value otherwise.

### GetGpsLocationsOk

`func (o *RouteEtaModel) GetGpsLocationsOk() (*[]GPSLocationModel, bool)`

GetGpsLocationsOk returns a tuple with the GpsLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsLocations

`func (o *RouteEtaModel) SetGpsLocations(v []GPSLocationModel)`

SetGpsLocations sets GpsLocations field to given value.

### HasGpsLocations

`func (o *RouteEtaModel) HasGpsLocations() bool`

HasGpsLocations returns a boolean if a field has been set.

### GetLatestKnownPosition

`func (o *RouteEtaModel) GetLatestKnownPosition() GPSLocationModel`

GetLatestKnownPosition returns the LatestKnownPosition field if non-nil, zero value otherwise.

### GetLatestKnownPositionOk

`func (o *RouteEtaModel) GetLatestKnownPositionOk() (*GPSLocationModel, bool)`

GetLatestKnownPositionOk returns a tuple with the LatestKnownPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestKnownPosition

`func (o *RouteEtaModel) SetLatestKnownPosition(v GPSLocationModel)`

SetLatestKnownPosition sets LatestKnownPosition field to given value.

### HasLatestKnownPosition

`func (o *RouteEtaModel) HasLatestKnownPosition() bool`

HasLatestKnownPosition returns a boolean if a field has been set.

### GetRecurrenceId

`func (o *RouteEtaModel) GetRecurrenceId() int32`

GetRecurrenceId returns the RecurrenceId field if non-nil, zero value otherwise.

### GetRecurrenceIdOk

`func (o *RouteEtaModel) GetRecurrenceIdOk() (*int32, bool)`

GetRecurrenceIdOk returns a tuple with the RecurrenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceId

`func (o *RouteEtaModel) SetRecurrenceId(v int32)`

SetRecurrenceId sets RecurrenceId field to given value.

### HasRecurrenceId

`func (o *RouteEtaModel) HasRecurrenceId() bool`

HasRecurrenceId returns a boolean if a field has been set.

### GetRecurrenceNr

`func (o *RouteEtaModel) GetRecurrenceNr() int32`

GetRecurrenceNr returns the RecurrenceNr field if non-nil, zero value otherwise.

### GetRecurrenceNrOk

`func (o *RouteEtaModel) GetRecurrenceNrOk() (*int32, bool)`

GetRecurrenceNrOk returns a tuple with the RecurrenceNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceNr

`func (o *RouteEtaModel) SetRecurrenceNr(v int32)`

SetRecurrenceNr sets RecurrenceNr field to given value.

### HasRecurrenceNr

`func (o *RouteEtaModel) HasRecurrenceNr() bool`

HasRecurrenceNr returns a boolean if a field has been set.

### GetRecurrence

`func (o *RouteEtaModel) GetRecurrence() RecurrenceModel`

GetRecurrence returns the Recurrence field if non-nil, zero value otherwise.

### GetRecurrenceOk

`func (o *RouteEtaModel) GetRecurrenceOk() (*RecurrenceModel, bool)`

GetRecurrenceOk returns a tuple with the Recurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrence

`func (o *RouteEtaModel) SetRecurrence(v RecurrenceModel)`

SetRecurrence sets Recurrence field to given value.

### HasRecurrence

`func (o *RouteEtaModel) HasRecurrence() bool`

HasRecurrence returns a boolean if a field has been set.

### GetOverdue

`func (o *RouteEtaModel) GetOverdue() bool`

GetOverdue returns the Overdue field if non-nil, zero value otherwise.

### GetOverdueOk

`func (o *RouteEtaModel) GetOverdueOk() (*bool, bool)`

GetOverdueOk returns a tuple with the Overdue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdue

`func (o *RouteEtaModel) SetOverdue(v bool)`

SetOverdue sets Overdue field to given value.

### HasOverdue

`func (o *RouteEtaModel) HasOverdue() bool`

HasOverdue returns a boolean if a field has been set.

### GetOptimized

`func (o *RouteEtaModel) GetOptimized() bool`

GetOptimized returns the Optimized field if non-nil, zero value otherwise.

### GetOptimizedOk

`func (o *RouteEtaModel) GetOptimizedOk() (*bool, bool)`

GetOptimizedOk returns a tuple with the Optimized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimized

`func (o *RouteEtaModel) SetOptimized(v bool)`

SetOptimized sets Optimized field to given value.

### HasOptimized

`func (o *RouteEtaModel) HasOptimized() bool`

HasOptimized returns a boolean if a field has been set.

### GetBlocked

`func (o *RouteEtaModel) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *RouteEtaModel) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *RouteEtaModel) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *RouteEtaModel) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetActive

`func (o *RouteEtaModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RouteEtaModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RouteEtaModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *RouteEtaModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetStartAddress

`func (o *RouteEtaModel) GetStartAddress() AddressModel`

GetStartAddress returns the StartAddress field if non-nil, zero value otherwise.

### GetStartAddressOk

`func (o *RouteEtaModel) GetStartAddressOk() (*AddressModel, bool)`

GetStartAddressOk returns a tuple with the StartAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddress

`func (o *RouteEtaModel) SetStartAddress(v AddressModel)`

SetStartAddress sets StartAddress field to given value.

### HasStartAddress

`func (o *RouteEtaModel) HasStartAddress() bool`

HasStartAddress returns a boolean if a field has been set.

### GetEndAddress

`func (o *RouteEtaModel) GetEndAddress() AddressModel`

GetEndAddress returns the EndAddress field if non-nil, zero value otherwise.

### GetEndAddressOk

`func (o *RouteEtaModel) GetEndAddressOk() (*AddressModel, bool)`

GetEndAddressOk returns a tuple with the EndAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddress

`func (o *RouteEtaModel) SetEndAddress(v AddressModel)`

SetEndAddress sets EndAddress field to given value.

### HasEndAddress

`func (o *RouteEtaModel) HasEndAddress() bool`

HasEndAddress returns a boolean if a field has been set.

### GetPlannedCapacities

`func (o *RouteEtaModel) GetPlannedCapacities() map[string]interface{}`

GetPlannedCapacities returns the PlannedCapacities field if non-nil, zero value otherwise.

### GetPlannedCapacitiesOk

`func (o *RouteEtaModel) GetPlannedCapacitiesOk() (*map[string]interface{}, bool)`

GetPlannedCapacitiesOk returns a tuple with the PlannedCapacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedCapacities

`func (o *RouteEtaModel) SetPlannedCapacities(v map[string]interface{})`

SetPlannedCapacities sets PlannedCapacities field to given value.

### HasPlannedCapacities

`func (o *RouteEtaModel) HasPlannedCapacities() bool`

HasPlannedCapacities returns a boolean if a field has been set.

### GetAppliedCapacities

`func (o *RouteEtaModel) GetAppliedCapacities() AppliedCapacitiesModel`

GetAppliedCapacities returns the AppliedCapacities field if non-nil, zero value otherwise.

### GetAppliedCapacitiesOk

`func (o *RouteEtaModel) GetAppliedCapacitiesOk() (*AppliedCapacitiesModel, bool)`

GetAppliedCapacitiesOk returns a tuple with the AppliedCapacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedCapacities

`func (o *RouteEtaModel) SetAppliedCapacities(v AppliedCapacitiesModel)`

SetAppliedCapacities sets AppliedCapacities field to given value.

### HasAppliedCapacities

`func (o *RouteEtaModel) HasAppliedCapacities() bool`

HasAppliedCapacities returns a boolean if a field has been set.

### GetCapacities

`func (o *RouteEtaModel) GetCapacities() []CapacityModel`

GetCapacities returns the Capacities field if non-nil, zero value otherwise.

### GetCapacitiesOk

`func (o *RouteEtaModel) GetCapacitiesOk() (*[]CapacityModel, bool)`

GetCapacitiesOk returns a tuple with the Capacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacities

`func (o *RouteEtaModel) SetCapacities(v []CapacityModel)`

SetCapacities sets Capacities field to given value.

### HasCapacities

`func (o *RouteEtaModel) HasCapacities() bool`

HasCapacities returns a boolean if a field has been set.

### GetActivityIds

`func (o *RouteEtaModel) GetActivityIds() []int32`

GetActivityIds returns the ActivityIds field if non-nil, zero value otherwise.

### GetActivityIdsOk

`func (o *RouteEtaModel) GetActivityIdsOk() (*[]int32, bool)`

GetActivityIdsOk returns a tuple with the ActivityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityIds

`func (o *RouteEtaModel) SetActivityIds(v []int32)`

SetActivityIds sets ActivityIds field to given value.

### HasActivityIds

`func (o *RouteEtaModel) HasActivityIds() bool`

HasActivityIds returns a boolean if a field has been set.

### GetLinks

`func (o *RouteEtaModel) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RouteEtaModel) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RouteEtaModel) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RouteEtaModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetaData

`func (o *RouteEtaModel) GetMetaData() []MetaDataModel`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *RouteEtaModel) GetMetaDataOk() (*[]MetaDataModel, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *RouteEtaModel) SetMetaData(v []MetaDataModel)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *RouteEtaModel) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetNotes

`func (o *RouteEtaModel) GetNotes() []NoteModel`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *RouteEtaModel) GetNotesOk() (*[]NoteModel, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *RouteEtaModel) SetNotes(v []NoteModel)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *RouteEtaModel) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetFiles

`func (o *RouteEtaModel) GetFiles() []FileModel`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *RouteEtaModel) GetFilesOk() (*[]FileModel, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *RouteEtaModel) SetFiles(v []FileModel)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *RouteEtaModel) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RouteEtaModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RouteEtaModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RouteEtaModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RouteEtaModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RouteEtaModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RouteEtaModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RouteEtaModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RouteEtaModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetTags

`func (o *RouteEtaModel) GetTags() []TagModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RouteEtaModel) GetTagsOk() (*[]TagModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RouteEtaModel) SetTags(v []TagModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RouteEtaModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTagNames

`func (o *RouteEtaModel) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *RouteEtaModel) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *RouteEtaModel) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.

### HasTagNames

`func (o *RouteEtaModel) HasTagNames() bool`

HasTagNames returns a boolean if a field has been set.

### GetZones

`func (o *RouteEtaModel) GetZones() []ZoneModel`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *RouteEtaModel) GetZonesOk() (*[]ZoneModel, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *RouteEtaModel) SetZones(v []ZoneModel)`

SetZones sets Zones field to given value.

### HasZones

`func (o *RouteEtaModel) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetZoneNames

`func (o *RouteEtaModel) GetZoneNames() []string`

GetZoneNames returns the ZoneNames field if non-nil, zero value otherwise.

### GetZoneNamesOk

`func (o *RouteEtaModel) GetZoneNamesOk() (*[]string, bool)`

GetZoneNamesOk returns a tuple with the ZoneNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneNames

`func (o *RouteEtaModel) SetZoneNames(v []string)`

SetZoneNames sets ZoneNames field to given value.

### HasZoneNames

`func (o *RouteEtaModel) HasZoneNames() bool`

HasZoneNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


