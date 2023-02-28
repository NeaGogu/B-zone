# RouteModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewRouteModel

`func NewRouteModel() *RouteModel`

NewRouteModel instantiates a new RouteModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteModelWithDefaults

`func NewRouteModelWithDefaults() *RouteModel`

NewRouteModelWithDefaults instantiates a new RouteModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RouteModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RouteModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RouteModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RouteModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNr

`func (o *RouteModel) GetNr() string`

GetNr returns the Nr field if non-nil, zero value otherwise.

### GetNrOk

`func (o *RouteModel) GetNrOk() (*string, bool)`

GetNrOk returns a tuple with the Nr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNr

`func (o *RouteModel) SetNr(v string)`

SetNr sets Nr field to given value.

### HasNr

`func (o *RouteModel) HasNr() bool`

HasNr returns a boolean if a field has been set.

### GetName

`func (o *RouteModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RouteModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RouteModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RouteModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatusName

`func (o *RouteModel) GetStatusName() string`

GetStatusName returns the StatusName field if non-nil, zero value otherwise.

### GetStatusNameOk

`func (o *RouteModel) GetStatusNameOk() (*string, bool)`

GetStatusNameOk returns a tuple with the StatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusName

`func (o *RouteModel) SetStatusName(v string)`

SetStatusName sets StatusName field to given value.

### HasStatusName

`func (o *RouteModel) HasStatusName() bool`

HasStatusName returns a boolean if a field has been set.

### GetStatusId

`func (o *RouteModel) GetStatusId() int64`

GetStatusId returns the StatusId field if non-nil, zero value otherwise.

### GetStatusIdOk

`func (o *RouteModel) GetStatusIdOk() (*int64, bool)`

GetStatusIdOk returns a tuple with the StatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusId

`func (o *RouteModel) SetStatusId(v int64)`

SetStatusId sets StatusId field to given value.

### HasStatusId

`func (o *RouteModel) HasStatusId() bool`

HasStatusId returns a boolean if a field has been set.

### GetNrOfStops

`func (o *RouteModel) GetNrOfStops() int32`

GetNrOfStops returns the NrOfStops field if non-nil, zero value otherwise.

### GetNrOfStopsOk

`func (o *RouteModel) GetNrOfStopsOk() (*int32, bool)`

GetNrOfStopsOk returns a tuple with the NrOfStops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrOfStops

`func (o *RouteModel) SetNrOfStops(v int32)`

SetNrOfStops sets NrOfStops field to given value.

### HasNrOfStops

`func (o *RouteModel) HasNrOfStops() bool`

HasNrOfStops returns a boolean if a field has been set.

### GetPauseId

`func (o *RouteModel) GetPauseId() int32`

GetPauseId returns the PauseId field if non-nil, zero value otherwise.

### GetPauseIdOk

`func (o *RouteModel) GetPauseIdOk() (*int32, bool)`

GetPauseIdOk returns a tuple with the PauseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseId

`func (o *RouteModel) SetPauseId(v int32)`

SetPauseId sets PauseId field to given value.

### HasPauseId

`func (o *RouteModel) HasPauseId() bool`

HasPauseId returns a boolean if a field has been set.

### GetPause

`func (o *RouteModel) GetPause() PauseModel`

GetPause returns the Pause field if non-nil, zero value otherwise.

### GetPauseOk

`func (o *RouteModel) GetPauseOk() (*PauseModel, bool)`

GetPauseOk returns a tuple with the Pause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPause

`func (o *RouteModel) SetPause(v PauseModel)`

SetPause sets Pause field to given value.

### HasPause

`func (o *RouteModel) HasPause() bool`

HasPause returns a boolean if a field has been set.

### GetDriverId

`func (o *RouteModel) GetDriverId() int32`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *RouteModel) GetDriverIdOk() (*int32, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *RouteModel) SetDriverId(v int32)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *RouteModel) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### GetDriverLink

`func (o *RouteModel) GetDriverLink() LinkModel`

GetDriverLink returns the DriverLink field if non-nil, zero value otherwise.

### GetDriverLinkOk

`func (o *RouteModel) GetDriverLinkOk() (*LinkModel, bool)`

GetDriverLinkOk returns a tuple with the DriverLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverLink

`func (o *RouteModel) SetDriverLink(v LinkModel)`

SetDriverLink sets DriverLink field to given value.

### HasDriverLink

`func (o *RouteModel) HasDriverLink() bool`

HasDriverLink returns a boolean if a field has been set.

### GetDriverLinks

`func (o *RouteModel) GetDriverLinks() []LinkModel`

GetDriverLinks returns the DriverLinks field if non-nil, zero value otherwise.

### GetDriverLinksOk

`func (o *RouteModel) GetDriverLinksOk() (*[]LinkModel, bool)`

GetDriverLinksOk returns a tuple with the DriverLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverLinks

`func (o *RouteModel) SetDriverLinks(v []LinkModel)`

SetDriverLinks sets DriverLinks field to given value.

### HasDriverLinks

`func (o *RouteModel) HasDriverLinks() bool`

HasDriverLinks returns a boolean if a field has been set.

### GetDriver

`func (o *RouteModel) GetDriver() UsersModel`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *RouteModel) GetDriverOk() (*UsersModel, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *RouteModel) SetDriver(v UsersModel)`

SetDriver sets Driver field to given value.

### HasDriver

`func (o *RouteModel) HasDriver() bool`

HasDriver returns a boolean if a field has been set.

### GetUserLink

`func (o *RouteModel) GetUserLink() []LinkModel`

GetUserLink returns the UserLink field if non-nil, zero value otherwise.

### GetUserLinkOk

`func (o *RouteModel) GetUserLinkOk() (*[]LinkModel, bool)`

GetUserLinkOk returns a tuple with the UserLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLink

`func (o *RouteModel) SetUserLink(v []LinkModel)`

SetUserLink sets UserLink field to given value.

### HasUserLink

`func (o *RouteModel) HasUserLink() bool`

HasUserLink returns a boolean if a field has been set.

### GetCoDriverIds

`func (o *RouteModel) GetCoDriverIds() []int32`

GetCoDriverIds returns the CoDriverIds field if non-nil, zero value otherwise.

### GetCoDriverIdsOk

`func (o *RouteModel) GetCoDriverIdsOk() (*[]int32, bool)`

GetCoDriverIdsOk returns a tuple with the CoDriverIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoDriverIds

`func (o *RouteModel) SetCoDriverIds(v []int32)`

SetCoDriverIds sets CoDriverIds field to given value.

### HasCoDriverIds

`func (o *RouteModel) HasCoDriverIds() bool`

HasCoDriverIds returns a boolean if a field has been set.

### GetCoDrivers

`func (o *RouteModel) GetCoDrivers() []UsersModel`

GetCoDrivers returns the CoDrivers field if non-nil, zero value otherwise.

### GetCoDriversOk

`func (o *RouteModel) GetCoDriversOk() (*[]UsersModel, bool)`

GetCoDriversOk returns a tuple with the CoDrivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoDrivers

`func (o *RouteModel) SetCoDrivers(v []UsersModel)`

SetCoDrivers sets CoDrivers field to given value.

### HasCoDrivers

`func (o *RouteModel) HasCoDrivers() bool`

HasCoDrivers returns a boolean if a field has been set.

### GetCarId

`func (o *RouteModel) GetCarId() int32`

GetCarId returns the CarId field if non-nil, zero value otherwise.

### GetCarIdOk

`func (o *RouteModel) GetCarIdOk() (*int32, bool)`

GetCarIdOk returns a tuple with the CarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarId

`func (o *RouteModel) SetCarId(v int32)`

SetCarId sets CarId field to given value.

### HasCarId

`func (o *RouteModel) HasCarId() bool`

HasCarId returns a boolean if a field has been set.

### GetVehicleId

`func (o *RouteModel) GetVehicleId() int32`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *RouteModel) GetVehicleIdOk() (*int32, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *RouteModel) SetVehicleId(v int32)`

SetVehicleId sets VehicleId field to given value.

### HasVehicleId

`func (o *RouteModel) HasVehicleId() bool`

HasVehicleId returns a boolean if a field has been set.

### GetVehicleLink

`func (o *RouteModel) GetVehicleLink() []LinkModel`

GetVehicleLink returns the VehicleLink field if non-nil, zero value otherwise.

### GetVehicleLinkOk

`func (o *RouteModel) GetVehicleLinkOk() (*[]LinkModel, bool)`

GetVehicleLinkOk returns a tuple with the VehicleLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleLink

`func (o *RouteModel) SetVehicleLink(v []LinkModel)`

SetVehicleLink sets VehicleLink field to given value.

### HasVehicleLink

`func (o *RouteModel) HasVehicleLink() bool`

HasVehicleLink returns a boolean if a field has been set.

### GetVehicle

`func (o *RouteModel) GetVehicle() VehicleModel`

GetVehicle returns the Vehicle field if non-nil, zero value otherwise.

### GetVehicleOk

`func (o *RouteModel) GetVehicleOk() (*VehicleModel, bool)`

GetVehicleOk returns a tuple with the Vehicle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicle

`func (o *RouteModel) SetVehicle(v VehicleModel)`

SetVehicle sets Vehicle field to given value.

### HasVehicle

`func (o *RouteModel) HasVehicle() bool`

HasVehicle returns a boolean if a field has been set.

### GetCarLink

`func (o *RouteModel) GetCarLink() []LinkModel`

GetCarLink returns the CarLink field if non-nil, zero value otherwise.

### GetCarLinkOk

`func (o *RouteModel) GetCarLinkOk() (*[]LinkModel, bool)`

GetCarLinkOk returns a tuple with the CarLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarLink

`func (o *RouteModel) SetCarLink(v []LinkModel)`

SetCarLink sets CarLink field to given value.

### HasCarLink

`func (o *RouteModel) HasCarLink() bool`

HasCarLink returns a boolean if a field has been set.

### GetCar

`func (o *RouteModel) GetCar() VehicleModel`

GetCar returns the Car field if non-nil, zero value otherwise.

### GetCarOk

`func (o *RouteModel) GetCarOk() (*VehicleModel, bool)`

GetCarOk returns a tuple with the Car field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCar

`func (o *RouteModel) SetCar(v VehicleModel)`

SetCar sets Car field to given value.

### HasCar

`func (o *RouteModel) HasCar() bool`

HasCar returns a boolean if a field has been set.

### GetTrailerId

`func (o *RouteModel) GetTrailerId() int32`

GetTrailerId returns the TrailerId field if non-nil, zero value otherwise.

### GetTrailerIdOk

`func (o *RouteModel) GetTrailerIdOk() (*int32, bool)`

GetTrailerIdOk returns a tuple with the TrailerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerId

`func (o *RouteModel) SetTrailerId(v int32)`

SetTrailerId sets TrailerId field to given value.

### HasTrailerId

`func (o *RouteModel) HasTrailerId() bool`

HasTrailerId returns a boolean if a field has been set.

### GetTrailerLink

`func (o *RouteModel) GetTrailerLink() []LinkModel`

GetTrailerLink returns the TrailerLink field if non-nil, zero value otherwise.

### GetTrailerLinkOk

`func (o *RouteModel) GetTrailerLinkOk() (*[]LinkModel, bool)`

GetTrailerLinkOk returns a tuple with the TrailerLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerLink

`func (o *RouteModel) SetTrailerLink(v []LinkModel)`

SetTrailerLink sets TrailerLink field to given value.

### HasTrailerLink

`func (o *RouteModel) HasTrailerLink() bool`

HasTrailerLink returns a boolean if a field has been set.

### GetTrailer

`func (o *RouteModel) GetTrailer() TrailerModel`

GetTrailer returns the Trailer field if non-nil, zero value otherwise.

### GetTrailerOk

`func (o *RouteModel) GetTrailerOk() (*TrailerModel, bool)`

GetTrailerOk returns a tuple with the Trailer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailer

`func (o *RouteModel) SetTrailer(v TrailerModel)`

SetTrailer sets Trailer field to given value.

### HasTrailer

`func (o *RouteModel) HasTrailer() bool`

HasTrailer returns a boolean if a field has been set.

### GetEquipment

`func (o *RouteModel) GetEquipment() []EquipmentModel`

GetEquipment returns the Equipment field if non-nil, zero value otherwise.

### GetEquipmentOk

`func (o *RouteModel) GetEquipmentOk() (*[]EquipmentModel, bool)`

GetEquipmentOk returns a tuple with the Equipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipment

`func (o *RouteModel) SetEquipment(v []EquipmentModel)`

SetEquipment sets Equipment field to given value.

### HasEquipment

`func (o *RouteModel) HasEquipment() bool`

HasEquipment returns a boolean if a field has been set.

### GetEquipmentIds

`func (o *RouteModel) GetEquipmentIds() []int32`

GetEquipmentIds returns the EquipmentIds field if non-nil, zero value otherwise.

### GetEquipmentIdsOk

`func (o *RouteModel) GetEquipmentIdsOk() (*[]int32, bool)`

GetEquipmentIdsOk returns a tuple with the EquipmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentIds

`func (o *RouteModel) SetEquipmentIds(v []int32)`

SetEquipmentIds sets EquipmentIds field to given value.

### HasEquipmentIds

`func (o *RouteModel) HasEquipmentIds() bool`

HasEquipmentIds returns a boolean if a field has been set.

### GetPlannedStartDuration

`func (o *RouteModel) GetPlannedStartDuration() int32`

GetPlannedStartDuration returns the PlannedStartDuration field if non-nil, zero value otherwise.

### GetPlannedStartDurationOk

`func (o *RouteModel) GetPlannedStartDurationOk() (*int32, bool)`

GetPlannedStartDurationOk returns a tuple with the PlannedStartDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedStartDuration

`func (o *RouteModel) SetPlannedStartDuration(v int32)`

SetPlannedStartDuration sets PlannedStartDuration field to given value.

### HasPlannedStartDuration

`func (o *RouteModel) HasPlannedStartDuration() bool`

HasPlannedStartDuration returns a boolean if a field has been set.

### GetPlannedEndDuration

`func (o *RouteModel) GetPlannedEndDuration() int32`

GetPlannedEndDuration returns the PlannedEndDuration field if non-nil, zero value otherwise.

### GetPlannedEndDurationOk

`func (o *RouteModel) GetPlannedEndDurationOk() (*int32, bool)`

GetPlannedEndDurationOk returns a tuple with the PlannedEndDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedEndDuration

`func (o *RouteModel) SetPlannedEndDuration(v int32)`

SetPlannedEndDuration sets PlannedEndDuration field to given value.

### HasPlannedEndDuration

`func (o *RouteModel) HasPlannedEndDuration() bool`

HasPlannedEndDuration returns a boolean if a field has been set.

### GetMaxSpeed

`func (o *RouteModel) GetMaxSpeed() int64`

GetMaxSpeed returns the MaxSpeed field if non-nil, zero value otherwise.

### GetMaxSpeedOk

`func (o *RouteModel) GetMaxSpeedOk() (*int64, bool)`

GetMaxSpeedOk returns a tuple with the MaxSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpeed

`func (o *RouteModel) SetMaxSpeed(v int64)`

SetMaxSpeed sets MaxSpeed field to given value.

### HasMaxSpeed

`func (o *RouteModel) HasMaxSpeed() bool`

HasMaxSpeed returns a boolean if a field has been set.

### GetSpeedFactor

`func (o *RouteModel) GetSpeedFactor() float64`

GetSpeedFactor returns the SpeedFactor field if non-nil, zero value otherwise.

### GetSpeedFactorOk

`func (o *RouteModel) GetSpeedFactorOk() (*float64, bool)`

GetSpeedFactorOk returns a tuple with the SpeedFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedFactor

`func (o *RouteModel) SetSpeedFactor(v float64)`

SetSpeedFactor sets SpeedFactor field to given value.

### HasSpeedFactor

`func (o *RouteModel) HasSpeedFactor() bool`

HasSpeedFactor returns a boolean if a field has been set.

### GetDurationFactor

`func (o *RouteModel) GetDurationFactor() float64`

GetDurationFactor returns the DurationFactor field if non-nil, zero value otherwise.

### GetDurationFactorOk

`func (o *RouteModel) GetDurationFactorOk() (*float64, bool)`

GetDurationFactorOk returns a tuple with the DurationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationFactor

`func (o *RouteModel) SetDurationFactor(v float64)`

SetDurationFactor sets DurationFactor field to given value.

### HasDurationFactor

`func (o *RouteModel) HasDurationFactor() bool`

HasDurationFactor returns a boolean if a field has been set.

### GetCostPerMeter

`func (o *RouteModel) GetCostPerMeter() float64`

GetCostPerMeter returns the CostPerMeter field if non-nil, zero value otherwise.

### GetCostPerMeterOk

`func (o *RouteModel) GetCostPerMeterOk() (*float64, bool)`

GetCostPerMeterOk returns a tuple with the CostPerMeter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerMeter

`func (o *RouteModel) SetCostPerMeter(v float64)`

SetCostPerMeter sets CostPerMeter field to given value.

### HasCostPerMeter

`func (o *RouteModel) HasCostPerMeter() bool`

HasCostPerMeter returns a boolean if a field has been set.

### GetCostPerRoute

`func (o *RouteModel) GetCostPerRoute() float64`

GetCostPerRoute returns the CostPerRoute field if non-nil, zero value otherwise.

### GetCostPerRouteOk

`func (o *RouteModel) GetCostPerRouteOk() (*float64, bool)`

GetCostPerRouteOk returns a tuple with the CostPerRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerRoute

`func (o *RouteModel) SetCostPerRoute(v float64)`

SetCostPerRoute sets CostPerRoute field to given value.

### HasCostPerRoute

`func (o *RouteModel) HasCostPerRoute() bool`

HasCostPerRoute returns a boolean if a field has been set.

### GetCostPerDrivingMinute

`func (o *RouteModel) GetCostPerDrivingMinute() float64`

GetCostPerDrivingMinute returns the CostPerDrivingMinute field if non-nil, zero value otherwise.

### GetCostPerDrivingMinuteOk

`func (o *RouteModel) GetCostPerDrivingMinuteOk() (*float64, bool)`

GetCostPerDrivingMinuteOk returns a tuple with the CostPerDrivingMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerDrivingMinute

`func (o *RouteModel) SetCostPerDrivingMinute(v float64)`

SetCostPerDrivingMinute sets CostPerDrivingMinute field to given value.

### HasCostPerDrivingMinute

`func (o *RouteModel) HasCostPerDrivingMinute() bool`

HasCostPerDrivingMinute returns a boolean if a field has been set.

### GetCostPerWaitingMinute

`func (o *RouteModel) GetCostPerWaitingMinute() float64`

GetCostPerWaitingMinute returns the CostPerWaitingMinute field if non-nil, zero value otherwise.

### GetCostPerWaitingMinuteOk

`func (o *RouteModel) GetCostPerWaitingMinuteOk() (*float64, bool)`

GetCostPerWaitingMinuteOk returns a tuple with the CostPerWaitingMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerWaitingMinute

`func (o *RouteModel) SetCostPerWaitingMinute(v float64)`

SetCostPerWaitingMinute sets CostPerWaitingMinute field to given value.

### HasCostPerWaitingMinute

`func (o *RouteModel) HasCostPerWaitingMinute() bool`

HasCostPerWaitingMinute returns a boolean if a field has been set.

### GetCostPerServiceMinute

`func (o *RouteModel) GetCostPerServiceMinute() float64`

GetCostPerServiceMinute returns the CostPerServiceMinute field if non-nil, zero value otherwise.

### GetCostPerServiceMinuteOk

`func (o *RouteModel) GetCostPerServiceMinuteOk() (*float64, bool)`

GetCostPerServiceMinuteOk returns a tuple with the CostPerServiceMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerServiceMinute

`func (o *RouteModel) SetCostPerServiceMinute(v float64)`

SetCostPerServiceMinute sets CostPerServiceMinute field to given value.

### HasCostPerServiceMinute

`func (o *RouteModel) HasCostPerServiceMinute() bool`

HasCostPerServiceMinute returns a boolean if a field has been set.

### GetEarliestDate

`func (o *RouteModel) GetEarliestDate() string`

GetEarliestDate returns the EarliestDate field if non-nil, zero value otherwise.

### GetEarliestDateOk

`func (o *RouteModel) GetEarliestDateOk() (*string, bool)`

GetEarliestDateOk returns a tuple with the EarliestDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestDate

`func (o *RouteModel) SetEarliestDate(v string)`

SetEarliestDate sets EarliestDate field to given value.

### HasEarliestDate

`func (o *RouteModel) HasEarliestDate() bool`

HasEarliestDate returns a boolean if a field has been set.

### GetLatestDate

`func (o *RouteModel) GetLatestDate() string`

GetLatestDate returns the LatestDate field if non-nil, zero value otherwise.

### GetLatestDateOk

`func (o *RouteModel) GetLatestDateOk() (*string, bool)`

GetLatestDateOk returns a tuple with the LatestDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDate

`func (o *RouteModel) SetLatestDate(v string)`

SetLatestDate sets LatestDate field to given value.

### HasLatestDate

`func (o *RouteModel) HasLatestDate() bool`

HasLatestDate returns a boolean if a field has been set.

### GetEarliestDateTime

`func (o *RouteModel) GetEarliestDateTime() time.Time`

GetEarliestDateTime returns the EarliestDateTime field if non-nil, zero value otherwise.

### GetEarliestDateTimeOk

`func (o *RouteModel) GetEarliestDateTimeOk() (*time.Time, bool)`

GetEarliestDateTimeOk returns a tuple with the EarliestDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestDateTime

`func (o *RouteModel) SetEarliestDateTime(v time.Time)`

SetEarliestDateTime sets EarliestDateTime field to given value.

### HasEarliestDateTime

`func (o *RouteModel) HasEarliestDateTime() bool`

HasEarliestDateTime returns a boolean if a field has been set.

### GetLatestDateTime

`func (o *RouteModel) GetLatestDateTime() time.Time`

GetLatestDateTime returns the LatestDateTime field if non-nil, zero value otherwise.

### GetLatestDateTimeOk

`func (o *RouteModel) GetLatestDateTimeOk() (*time.Time, bool)`

GetLatestDateTimeOk returns a tuple with the LatestDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestDateTime

`func (o *RouteModel) SetLatestDateTime(v time.Time)`

SetLatestDateTime sets LatestDateTime field to given value.

### HasLatestDateTime

`func (o *RouteModel) HasLatestDateTime() bool`

HasLatestDateTime returns a boolean if a field has been set.

### GetPlannedDateTimeFrom

`func (o *RouteModel) GetPlannedDateTimeFrom() time.Time`

GetPlannedDateTimeFrom returns the PlannedDateTimeFrom field if non-nil, zero value otherwise.

### GetPlannedDateTimeFromOk

`func (o *RouteModel) GetPlannedDateTimeFromOk() (*time.Time, bool)`

GetPlannedDateTimeFromOk returns a tuple with the PlannedDateTimeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDateTimeFrom

`func (o *RouteModel) SetPlannedDateTimeFrom(v time.Time)`

SetPlannedDateTimeFrom sets PlannedDateTimeFrom field to given value.

### HasPlannedDateTimeFrom

`func (o *RouteModel) HasPlannedDateTimeFrom() bool`

HasPlannedDateTimeFrom returns a boolean if a field has been set.

### GetPlannedDateTimeTo

`func (o *RouteModel) GetPlannedDateTimeTo() time.Time`

GetPlannedDateTimeTo returns the PlannedDateTimeTo field if non-nil, zero value otherwise.

### GetPlannedDateTimeToOk

`func (o *RouteModel) GetPlannedDateTimeToOk() (*time.Time, bool)`

GetPlannedDateTimeToOk returns a tuple with the PlannedDateTimeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDateTimeTo

`func (o *RouteModel) SetPlannedDateTimeTo(v time.Time)`

SetPlannedDateTimeTo sets PlannedDateTimeTo field to given value.

### HasPlannedDateTimeTo

`func (o *RouteModel) HasPlannedDateTimeTo() bool`

HasPlannedDateTimeTo returns a boolean if a field has been set.

### GetExecutedDateTimeFrom

`func (o *RouteModel) GetExecutedDateTimeFrom() time.Time`

GetExecutedDateTimeFrom returns the ExecutedDateTimeFrom field if non-nil, zero value otherwise.

### GetExecutedDateTimeFromOk

`func (o *RouteModel) GetExecutedDateTimeFromOk() (*time.Time, bool)`

GetExecutedDateTimeFromOk returns a tuple with the ExecutedDateTimeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedDateTimeFrom

`func (o *RouteModel) SetExecutedDateTimeFrom(v time.Time)`

SetExecutedDateTimeFrom sets ExecutedDateTimeFrom field to given value.

### HasExecutedDateTimeFrom

`func (o *RouteModel) HasExecutedDateTimeFrom() bool`

HasExecutedDateTimeFrom returns a boolean if a field has been set.

### GetExecutedDateTimeTo

`func (o *RouteModel) GetExecutedDateTimeTo() time.Time`

GetExecutedDateTimeTo returns the ExecutedDateTimeTo field if non-nil, zero value otherwise.

### GetExecutedDateTimeToOk

`func (o *RouteModel) GetExecutedDateTimeToOk() (*time.Time, bool)`

GetExecutedDateTimeToOk returns a tuple with the ExecutedDateTimeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedDateTimeTo

`func (o *RouteModel) SetExecutedDateTimeTo(v time.Time)`

SetExecutedDateTimeTo sets ExecutedDateTimeTo field to given value.

### HasExecutedDateTimeTo

`func (o *RouteModel) HasExecutedDateTimeTo() bool`

HasExecutedDateTimeTo returns a boolean if a field has been set.

### GetPlannedDrivingDistance

`func (o *RouteModel) GetPlannedDrivingDistance() int32`

GetPlannedDrivingDistance returns the PlannedDrivingDistance field if non-nil, zero value otherwise.

### GetPlannedDrivingDistanceOk

`func (o *RouteModel) GetPlannedDrivingDistanceOk() (*int32, bool)`

GetPlannedDrivingDistanceOk returns a tuple with the PlannedDrivingDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDrivingDistance

`func (o *RouteModel) SetPlannedDrivingDistance(v int32)`

SetPlannedDrivingDistance sets PlannedDrivingDistance field to given value.

### HasPlannedDrivingDistance

`func (o *RouteModel) HasPlannedDrivingDistance() bool`

HasPlannedDrivingDistance returns a boolean if a field has been set.

### GetPlannedDrivingDuration

`func (o *RouteModel) GetPlannedDrivingDuration() int32`

GetPlannedDrivingDuration returns the PlannedDrivingDuration field if non-nil, zero value otherwise.

### GetPlannedDrivingDurationOk

`func (o *RouteModel) GetPlannedDrivingDurationOk() (*int32, bool)`

GetPlannedDrivingDurationOk returns a tuple with the PlannedDrivingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDrivingDuration

`func (o *RouteModel) SetPlannedDrivingDuration(v int32)`

SetPlannedDrivingDuration sets PlannedDrivingDuration field to given value.

### HasPlannedDrivingDuration

`func (o *RouteModel) HasPlannedDrivingDuration() bool`

HasPlannedDrivingDuration returns a boolean if a field has been set.

### GetPlannedWaitingDuration

`func (o *RouteModel) GetPlannedWaitingDuration() int32`

GetPlannedWaitingDuration returns the PlannedWaitingDuration field if non-nil, zero value otherwise.

### GetPlannedWaitingDurationOk

`func (o *RouteModel) GetPlannedWaitingDurationOk() (*int32, bool)`

GetPlannedWaitingDurationOk returns a tuple with the PlannedWaitingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedWaitingDuration

`func (o *RouteModel) SetPlannedWaitingDuration(v int32)`

SetPlannedWaitingDuration sets PlannedWaitingDuration field to given value.

### HasPlannedWaitingDuration

`func (o *RouteModel) HasPlannedWaitingDuration() bool`

HasPlannedWaitingDuration returns a boolean if a field has been set.

### GetPlannedActivityDuration

`func (o *RouteModel) GetPlannedActivityDuration() int32`

GetPlannedActivityDuration returns the PlannedActivityDuration field if non-nil, zero value otherwise.

### GetPlannedActivityDurationOk

`func (o *RouteModel) GetPlannedActivityDurationOk() (*int32, bool)`

GetPlannedActivityDurationOk returns a tuple with the PlannedActivityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedActivityDuration

`func (o *RouteModel) SetPlannedActivityDuration(v int32)`

SetPlannedActivityDuration sets PlannedActivityDuration field to given value.

### HasPlannedActivityDuration

`func (o *RouteModel) HasPlannedActivityDuration() bool`

HasPlannedActivityDuration returns a boolean if a field has been set.

### GetPlannedTotalDuration

`func (o *RouteModel) GetPlannedTotalDuration() int32`

GetPlannedTotalDuration returns the PlannedTotalDuration field if non-nil, zero value otherwise.

### GetPlannedTotalDurationOk

`func (o *RouteModel) GetPlannedTotalDurationOk() (*int32, bool)`

GetPlannedTotalDurationOk returns a tuple with the PlannedTotalDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedTotalDuration

`func (o *RouteModel) SetPlannedTotalDuration(v int32)`

SetPlannedTotalDuration sets PlannedTotalDuration field to given value.

### HasPlannedTotalDuration

`func (o *RouteModel) HasPlannedTotalDuration() bool`

HasPlannedTotalDuration returns a boolean if a field has been set.

### GetGpsLocations

`func (o *RouteModel) GetGpsLocations() []GPSLocationModel`

GetGpsLocations returns the GpsLocations field if non-nil, zero value otherwise.

### GetGpsLocationsOk

`func (o *RouteModel) GetGpsLocationsOk() (*[]GPSLocationModel, bool)`

GetGpsLocationsOk returns a tuple with the GpsLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsLocations

`func (o *RouteModel) SetGpsLocations(v []GPSLocationModel)`

SetGpsLocations sets GpsLocations field to given value.

### HasGpsLocations

`func (o *RouteModel) HasGpsLocations() bool`

HasGpsLocations returns a boolean if a field has been set.

### GetLatestKnownPosition

`func (o *RouteModel) GetLatestKnownPosition() GPSLocationModel`

GetLatestKnownPosition returns the LatestKnownPosition field if non-nil, zero value otherwise.

### GetLatestKnownPositionOk

`func (o *RouteModel) GetLatestKnownPositionOk() (*GPSLocationModel, bool)`

GetLatestKnownPositionOk returns a tuple with the LatestKnownPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestKnownPosition

`func (o *RouteModel) SetLatestKnownPosition(v GPSLocationModel)`

SetLatestKnownPosition sets LatestKnownPosition field to given value.

### HasLatestKnownPosition

`func (o *RouteModel) HasLatestKnownPosition() bool`

HasLatestKnownPosition returns a boolean if a field has been set.

### GetRecurrenceId

`func (o *RouteModel) GetRecurrenceId() int32`

GetRecurrenceId returns the RecurrenceId field if non-nil, zero value otherwise.

### GetRecurrenceIdOk

`func (o *RouteModel) GetRecurrenceIdOk() (*int32, bool)`

GetRecurrenceIdOk returns a tuple with the RecurrenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceId

`func (o *RouteModel) SetRecurrenceId(v int32)`

SetRecurrenceId sets RecurrenceId field to given value.

### HasRecurrenceId

`func (o *RouteModel) HasRecurrenceId() bool`

HasRecurrenceId returns a boolean if a field has been set.

### GetRecurrenceNr

`func (o *RouteModel) GetRecurrenceNr() int32`

GetRecurrenceNr returns the RecurrenceNr field if non-nil, zero value otherwise.

### GetRecurrenceNrOk

`func (o *RouteModel) GetRecurrenceNrOk() (*int32, bool)`

GetRecurrenceNrOk returns a tuple with the RecurrenceNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrenceNr

`func (o *RouteModel) SetRecurrenceNr(v int32)`

SetRecurrenceNr sets RecurrenceNr field to given value.

### HasRecurrenceNr

`func (o *RouteModel) HasRecurrenceNr() bool`

HasRecurrenceNr returns a boolean if a field has been set.

### GetRecurrence

`func (o *RouteModel) GetRecurrence() RecurrenceModel`

GetRecurrence returns the Recurrence field if non-nil, zero value otherwise.

### GetRecurrenceOk

`func (o *RouteModel) GetRecurrenceOk() (*RecurrenceModel, bool)`

GetRecurrenceOk returns a tuple with the Recurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrence

`func (o *RouteModel) SetRecurrence(v RecurrenceModel)`

SetRecurrence sets Recurrence field to given value.

### HasRecurrence

`func (o *RouteModel) HasRecurrence() bool`

HasRecurrence returns a boolean if a field has been set.

### GetOverdue

`func (o *RouteModel) GetOverdue() bool`

GetOverdue returns the Overdue field if non-nil, zero value otherwise.

### GetOverdueOk

`func (o *RouteModel) GetOverdueOk() (*bool, bool)`

GetOverdueOk returns a tuple with the Overdue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdue

`func (o *RouteModel) SetOverdue(v bool)`

SetOverdue sets Overdue field to given value.

### HasOverdue

`func (o *RouteModel) HasOverdue() bool`

HasOverdue returns a boolean if a field has been set.

### GetOptimized

`func (o *RouteModel) GetOptimized() bool`

GetOptimized returns the Optimized field if non-nil, zero value otherwise.

### GetOptimizedOk

`func (o *RouteModel) GetOptimizedOk() (*bool, bool)`

GetOptimizedOk returns a tuple with the Optimized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimized

`func (o *RouteModel) SetOptimized(v bool)`

SetOptimized sets Optimized field to given value.

### HasOptimized

`func (o *RouteModel) HasOptimized() bool`

HasOptimized returns a boolean if a field has been set.

### GetBlocked

`func (o *RouteModel) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *RouteModel) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *RouteModel) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *RouteModel) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetActive

`func (o *RouteModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RouteModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RouteModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *RouteModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetStartAddress

`func (o *RouteModel) GetStartAddress() AddressModel`

GetStartAddress returns the StartAddress field if non-nil, zero value otherwise.

### GetStartAddressOk

`func (o *RouteModel) GetStartAddressOk() (*AddressModel, bool)`

GetStartAddressOk returns a tuple with the StartAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddress

`func (o *RouteModel) SetStartAddress(v AddressModel)`

SetStartAddress sets StartAddress field to given value.

### HasStartAddress

`func (o *RouteModel) HasStartAddress() bool`

HasStartAddress returns a boolean if a field has been set.

### GetEndAddress

`func (o *RouteModel) GetEndAddress() AddressModel`

GetEndAddress returns the EndAddress field if non-nil, zero value otherwise.

### GetEndAddressOk

`func (o *RouteModel) GetEndAddressOk() (*AddressModel, bool)`

GetEndAddressOk returns a tuple with the EndAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddress

`func (o *RouteModel) SetEndAddress(v AddressModel)`

SetEndAddress sets EndAddress field to given value.

### HasEndAddress

`func (o *RouteModel) HasEndAddress() bool`

HasEndAddress returns a boolean if a field has been set.

### GetPlannedCapacities

`func (o *RouteModel) GetPlannedCapacities() map[string]interface{}`

GetPlannedCapacities returns the PlannedCapacities field if non-nil, zero value otherwise.

### GetPlannedCapacitiesOk

`func (o *RouteModel) GetPlannedCapacitiesOk() (*map[string]interface{}, bool)`

GetPlannedCapacitiesOk returns a tuple with the PlannedCapacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedCapacities

`func (o *RouteModel) SetPlannedCapacities(v map[string]interface{})`

SetPlannedCapacities sets PlannedCapacities field to given value.

### HasPlannedCapacities

`func (o *RouteModel) HasPlannedCapacities() bool`

HasPlannedCapacities returns a boolean if a field has been set.

### GetAppliedCapacities

`func (o *RouteModel) GetAppliedCapacities() AppliedCapacitiesModel`

GetAppliedCapacities returns the AppliedCapacities field if non-nil, zero value otherwise.

### GetAppliedCapacitiesOk

`func (o *RouteModel) GetAppliedCapacitiesOk() (*AppliedCapacitiesModel, bool)`

GetAppliedCapacitiesOk returns a tuple with the AppliedCapacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedCapacities

`func (o *RouteModel) SetAppliedCapacities(v AppliedCapacitiesModel)`

SetAppliedCapacities sets AppliedCapacities field to given value.

### HasAppliedCapacities

`func (o *RouteModel) HasAppliedCapacities() bool`

HasAppliedCapacities returns a boolean if a field has been set.

### GetCapacities

`func (o *RouteModel) GetCapacities() []CapacityModel`

GetCapacities returns the Capacities field if non-nil, zero value otherwise.

### GetCapacitiesOk

`func (o *RouteModel) GetCapacitiesOk() (*[]CapacityModel, bool)`

GetCapacitiesOk returns a tuple with the Capacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacities

`func (o *RouteModel) SetCapacities(v []CapacityModel)`

SetCapacities sets Capacities field to given value.

### HasCapacities

`func (o *RouteModel) HasCapacities() bool`

HasCapacities returns a boolean if a field has been set.

### GetActivityIds

`func (o *RouteModel) GetActivityIds() []int32`

GetActivityIds returns the ActivityIds field if non-nil, zero value otherwise.

### GetActivityIdsOk

`func (o *RouteModel) GetActivityIdsOk() (*[]int32, bool)`

GetActivityIdsOk returns a tuple with the ActivityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityIds

`func (o *RouteModel) SetActivityIds(v []int32)`

SetActivityIds sets ActivityIds field to given value.

### HasActivityIds

`func (o *RouteModel) HasActivityIds() bool`

HasActivityIds returns a boolean if a field has been set.

### GetLinks

`func (o *RouteModel) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RouteModel) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RouteModel) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RouteModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetaData

`func (o *RouteModel) GetMetaData() []MetaDataModel`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *RouteModel) GetMetaDataOk() (*[]MetaDataModel, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *RouteModel) SetMetaData(v []MetaDataModel)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *RouteModel) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetNotes

`func (o *RouteModel) GetNotes() []NoteModel`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *RouteModel) GetNotesOk() (*[]NoteModel, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *RouteModel) SetNotes(v []NoteModel)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *RouteModel) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetFiles

`func (o *RouteModel) GetFiles() []FileModel`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *RouteModel) GetFilesOk() (*[]FileModel, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *RouteModel) SetFiles(v []FileModel)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *RouteModel) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RouteModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RouteModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RouteModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RouteModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RouteModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RouteModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RouteModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RouteModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetTags

`func (o *RouteModel) GetTags() []TagModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RouteModel) GetTagsOk() (*[]TagModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RouteModel) SetTags(v []TagModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RouteModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTagNames

`func (o *RouteModel) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *RouteModel) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *RouteModel) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.

### HasTagNames

`func (o *RouteModel) HasTagNames() bool`

HasTagNames returns a boolean if a field has been set.

### GetZones

`func (o *RouteModel) GetZones() []ZoneModel`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *RouteModel) GetZonesOk() (*[]ZoneModel, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *RouteModel) SetZones(v []ZoneModel)`

SetZones sets Zones field to given value.

### HasZones

`func (o *RouteModel) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetZoneNames

`func (o *RouteModel) GetZoneNames() []string`

GetZoneNames returns the ZoneNames field if non-nil, zero value otherwise.

### GetZoneNamesOk

`func (o *RouteModel) GetZoneNamesOk() (*[]string, bool)`

GetZoneNamesOk returns a tuple with the ZoneNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneNames

`func (o *RouteModel) SetZoneNames(v []string)`

SetZoneNames sets ZoneNames field to given value.

### HasZoneNames

`func (o *RouteModel) HasZoneNames() bool`

HasZoneNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


