# RouteOptionsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyAddressBundling** | Pointer to **bool** |  | [optional] 
**IncludeAddress** | Pointer to **bool** |  | [optional] 
**IncludeAddressObject** | Pointer to **bool** |  | [optional] 
**IncludeRouteStatus** | Pointer to **bool** |  | [optional] 
**IncludeStatusName** | Pointer to **bool** |  | [optional] 
**IncludeRouteTags** | Pointer to **bool** | Deprecated! use include_tags | [optional] 
**IncludeFiles** | Pointer to **bool** |  | [optional] 
**IncludeTags** | Pointer to **bool** |  | [optional] 
**IncludeTagNames** | Pointer to **bool** |  | [optional] 
**IncludeDriver** | Pointer to **bool** |  | [optional] 
**IncludeDriverLinks** | Pointer to **bool** |  | [optional] 
**IncludeCoDrivers** | Pointer to **bool** |  | [optional] 
**IncludeCoDriverIds** | Pointer to **bool** |  | [optional] 
**IncludeCar** | Pointer to **bool** |  | [optional] 
**IncludeCarLinks** | Pointer to **bool** |  | [optional] 
**IncludeVehicle** | Pointer to **bool** |  | [optional] 
**IncludeVehicleLinks** | Pointer to **bool** |  | [optional] 
**IncludeTrailer** | Pointer to **bool** |  | [optional] 
**IncludeTrailerLinks** | Pointer to **bool** |  | [optional] 
**IncludeDriverInfo** | Pointer to **bool** |  | [optional] 
**IncludeEquipmentInfoCar** | Pointer to **bool** |  | [optional] 
**IncludeEquipment** | Pointer to **bool** |  | [optional] 
**IncludeEquipmentIds** | Pointer to **bool** |  | [optional] 
**IncludeGpsLocations** | Pointer to **bool** |  | [optional] 
**IncludePause** | Pointer to **bool** |  | [optional] 
**IncludeActivityIds** | Pointer to **bool** |  | [optional] 
**IncludeLatestPosition** | Pointer to **bool** |  | [optional] 
**IncludeZones** | Pointer to **bool** |  | [optional] 
**IncludeZoneNames** | Pointer to **bool** |  | [optional] 
**IncludeNotes** | Pointer to **bool** |  | [optional] 
**IncludeLinks** | Pointer to **bool** |  | [optional] 
**IncludeMetaData** | Pointer to **bool** |  | [optional] 
**IncludeAppliedCapacities** | Pointer to **bool** |  | [optional] 
**IncludeAppliedCapacitiesV2** | Pointer to **bool** |  | [optional] 
**IncludeCapacities** | Pointer to **bool** |  | [optional] 
**IncludeRecurrence** | Pointer to **bool** |  | [optional] 

## Methods

### NewRouteOptionsModel

`func NewRouteOptionsModel() *RouteOptionsModel`

NewRouteOptionsModel instantiates a new RouteOptionsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteOptionsModelWithDefaults

`func NewRouteOptionsModelWithDefaults() *RouteOptionsModel`

NewRouteOptionsModelWithDefaults instantiates a new RouteOptionsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplyAddressBundling

`func (o *RouteOptionsModel) GetApplyAddressBundling() bool`

GetApplyAddressBundling returns the ApplyAddressBundling field if non-nil, zero value otherwise.

### GetApplyAddressBundlingOk

`func (o *RouteOptionsModel) GetApplyAddressBundlingOk() (*bool, bool)`

GetApplyAddressBundlingOk returns a tuple with the ApplyAddressBundling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyAddressBundling

`func (o *RouteOptionsModel) SetApplyAddressBundling(v bool)`

SetApplyAddressBundling sets ApplyAddressBundling field to given value.

### HasApplyAddressBundling

`func (o *RouteOptionsModel) HasApplyAddressBundling() bool`

HasApplyAddressBundling returns a boolean if a field has been set.

### GetIncludeAddress

`func (o *RouteOptionsModel) GetIncludeAddress() bool`

GetIncludeAddress returns the IncludeAddress field if non-nil, zero value otherwise.

### GetIncludeAddressOk

`func (o *RouteOptionsModel) GetIncludeAddressOk() (*bool, bool)`

GetIncludeAddressOk returns a tuple with the IncludeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAddress

`func (o *RouteOptionsModel) SetIncludeAddress(v bool)`

SetIncludeAddress sets IncludeAddress field to given value.

### HasIncludeAddress

`func (o *RouteOptionsModel) HasIncludeAddress() bool`

HasIncludeAddress returns a boolean if a field has been set.

### GetIncludeAddressObject

`func (o *RouteOptionsModel) GetIncludeAddressObject() bool`

GetIncludeAddressObject returns the IncludeAddressObject field if non-nil, zero value otherwise.

### GetIncludeAddressObjectOk

`func (o *RouteOptionsModel) GetIncludeAddressObjectOk() (*bool, bool)`

GetIncludeAddressObjectOk returns a tuple with the IncludeAddressObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAddressObject

`func (o *RouteOptionsModel) SetIncludeAddressObject(v bool)`

SetIncludeAddressObject sets IncludeAddressObject field to given value.

### HasIncludeAddressObject

`func (o *RouteOptionsModel) HasIncludeAddressObject() bool`

HasIncludeAddressObject returns a boolean if a field has been set.

### GetIncludeRouteStatus

`func (o *RouteOptionsModel) GetIncludeRouteStatus() bool`

GetIncludeRouteStatus returns the IncludeRouteStatus field if non-nil, zero value otherwise.

### GetIncludeRouteStatusOk

`func (o *RouteOptionsModel) GetIncludeRouteStatusOk() (*bool, bool)`

GetIncludeRouteStatusOk returns a tuple with the IncludeRouteStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRouteStatus

`func (o *RouteOptionsModel) SetIncludeRouteStatus(v bool)`

SetIncludeRouteStatus sets IncludeRouteStatus field to given value.

### HasIncludeRouteStatus

`func (o *RouteOptionsModel) HasIncludeRouteStatus() bool`

HasIncludeRouteStatus returns a boolean if a field has been set.

### GetIncludeStatusName

`func (o *RouteOptionsModel) GetIncludeStatusName() bool`

GetIncludeStatusName returns the IncludeStatusName field if non-nil, zero value otherwise.

### GetIncludeStatusNameOk

`func (o *RouteOptionsModel) GetIncludeStatusNameOk() (*bool, bool)`

GetIncludeStatusNameOk returns a tuple with the IncludeStatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeStatusName

`func (o *RouteOptionsModel) SetIncludeStatusName(v bool)`

SetIncludeStatusName sets IncludeStatusName field to given value.

### HasIncludeStatusName

`func (o *RouteOptionsModel) HasIncludeStatusName() bool`

HasIncludeStatusName returns a boolean if a field has been set.

### GetIncludeRouteTags

`func (o *RouteOptionsModel) GetIncludeRouteTags() bool`

GetIncludeRouteTags returns the IncludeRouteTags field if non-nil, zero value otherwise.

### GetIncludeRouteTagsOk

`func (o *RouteOptionsModel) GetIncludeRouteTagsOk() (*bool, bool)`

GetIncludeRouteTagsOk returns a tuple with the IncludeRouteTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRouteTags

`func (o *RouteOptionsModel) SetIncludeRouteTags(v bool)`

SetIncludeRouteTags sets IncludeRouteTags field to given value.

### HasIncludeRouteTags

`func (o *RouteOptionsModel) HasIncludeRouteTags() bool`

HasIncludeRouteTags returns a boolean if a field has been set.

### GetIncludeFiles

`func (o *RouteOptionsModel) GetIncludeFiles() bool`

GetIncludeFiles returns the IncludeFiles field if non-nil, zero value otherwise.

### GetIncludeFilesOk

`func (o *RouteOptionsModel) GetIncludeFilesOk() (*bool, bool)`

GetIncludeFilesOk returns a tuple with the IncludeFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFiles

`func (o *RouteOptionsModel) SetIncludeFiles(v bool)`

SetIncludeFiles sets IncludeFiles field to given value.

### HasIncludeFiles

`func (o *RouteOptionsModel) HasIncludeFiles() bool`

HasIncludeFiles returns a boolean if a field has been set.

### GetIncludeTags

`func (o *RouteOptionsModel) GetIncludeTags() bool`

GetIncludeTags returns the IncludeTags field if non-nil, zero value otherwise.

### GetIncludeTagsOk

`func (o *RouteOptionsModel) GetIncludeTagsOk() (*bool, bool)`

GetIncludeTagsOk returns a tuple with the IncludeTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTags

`func (o *RouteOptionsModel) SetIncludeTags(v bool)`

SetIncludeTags sets IncludeTags field to given value.

### HasIncludeTags

`func (o *RouteOptionsModel) HasIncludeTags() bool`

HasIncludeTags returns a boolean if a field has been set.

### GetIncludeTagNames

`func (o *RouteOptionsModel) GetIncludeTagNames() bool`

GetIncludeTagNames returns the IncludeTagNames field if non-nil, zero value otherwise.

### GetIncludeTagNamesOk

`func (o *RouteOptionsModel) GetIncludeTagNamesOk() (*bool, bool)`

GetIncludeTagNamesOk returns a tuple with the IncludeTagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTagNames

`func (o *RouteOptionsModel) SetIncludeTagNames(v bool)`

SetIncludeTagNames sets IncludeTagNames field to given value.

### HasIncludeTagNames

`func (o *RouteOptionsModel) HasIncludeTagNames() bool`

HasIncludeTagNames returns a boolean if a field has been set.

### GetIncludeDriver

`func (o *RouteOptionsModel) GetIncludeDriver() bool`

GetIncludeDriver returns the IncludeDriver field if non-nil, zero value otherwise.

### GetIncludeDriverOk

`func (o *RouteOptionsModel) GetIncludeDriverOk() (*bool, bool)`

GetIncludeDriverOk returns a tuple with the IncludeDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDriver

`func (o *RouteOptionsModel) SetIncludeDriver(v bool)`

SetIncludeDriver sets IncludeDriver field to given value.

### HasIncludeDriver

`func (o *RouteOptionsModel) HasIncludeDriver() bool`

HasIncludeDriver returns a boolean if a field has been set.

### GetIncludeDriverLinks

`func (o *RouteOptionsModel) GetIncludeDriverLinks() bool`

GetIncludeDriverLinks returns the IncludeDriverLinks field if non-nil, zero value otherwise.

### GetIncludeDriverLinksOk

`func (o *RouteOptionsModel) GetIncludeDriverLinksOk() (*bool, bool)`

GetIncludeDriverLinksOk returns a tuple with the IncludeDriverLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDriverLinks

`func (o *RouteOptionsModel) SetIncludeDriverLinks(v bool)`

SetIncludeDriverLinks sets IncludeDriverLinks field to given value.

### HasIncludeDriverLinks

`func (o *RouteOptionsModel) HasIncludeDriverLinks() bool`

HasIncludeDriverLinks returns a boolean if a field has been set.

### GetIncludeCoDrivers

`func (o *RouteOptionsModel) GetIncludeCoDrivers() bool`

GetIncludeCoDrivers returns the IncludeCoDrivers field if non-nil, zero value otherwise.

### GetIncludeCoDriversOk

`func (o *RouteOptionsModel) GetIncludeCoDriversOk() (*bool, bool)`

GetIncludeCoDriversOk returns a tuple with the IncludeCoDrivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCoDrivers

`func (o *RouteOptionsModel) SetIncludeCoDrivers(v bool)`

SetIncludeCoDrivers sets IncludeCoDrivers field to given value.

### HasIncludeCoDrivers

`func (o *RouteOptionsModel) HasIncludeCoDrivers() bool`

HasIncludeCoDrivers returns a boolean if a field has been set.

### GetIncludeCoDriverIds

`func (o *RouteOptionsModel) GetIncludeCoDriverIds() bool`

GetIncludeCoDriverIds returns the IncludeCoDriverIds field if non-nil, zero value otherwise.

### GetIncludeCoDriverIdsOk

`func (o *RouteOptionsModel) GetIncludeCoDriverIdsOk() (*bool, bool)`

GetIncludeCoDriverIdsOk returns a tuple with the IncludeCoDriverIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCoDriverIds

`func (o *RouteOptionsModel) SetIncludeCoDriverIds(v bool)`

SetIncludeCoDriverIds sets IncludeCoDriverIds field to given value.

### HasIncludeCoDriverIds

`func (o *RouteOptionsModel) HasIncludeCoDriverIds() bool`

HasIncludeCoDriverIds returns a boolean if a field has been set.

### GetIncludeCar

`func (o *RouteOptionsModel) GetIncludeCar() bool`

GetIncludeCar returns the IncludeCar field if non-nil, zero value otherwise.

### GetIncludeCarOk

`func (o *RouteOptionsModel) GetIncludeCarOk() (*bool, bool)`

GetIncludeCarOk returns a tuple with the IncludeCar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCar

`func (o *RouteOptionsModel) SetIncludeCar(v bool)`

SetIncludeCar sets IncludeCar field to given value.

### HasIncludeCar

`func (o *RouteOptionsModel) HasIncludeCar() bool`

HasIncludeCar returns a boolean if a field has been set.

### GetIncludeCarLinks

`func (o *RouteOptionsModel) GetIncludeCarLinks() bool`

GetIncludeCarLinks returns the IncludeCarLinks field if non-nil, zero value otherwise.

### GetIncludeCarLinksOk

`func (o *RouteOptionsModel) GetIncludeCarLinksOk() (*bool, bool)`

GetIncludeCarLinksOk returns a tuple with the IncludeCarLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCarLinks

`func (o *RouteOptionsModel) SetIncludeCarLinks(v bool)`

SetIncludeCarLinks sets IncludeCarLinks field to given value.

### HasIncludeCarLinks

`func (o *RouteOptionsModel) HasIncludeCarLinks() bool`

HasIncludeCarLinks returns a boolean if a field has been set.

### GetIncludeVehicle

`func (o *RouteOptionsModel) GetIncludeVehicle() bool`

GetIncludeVehicle returns the IncludeVehicle field if non-nil, zero value otherwise.

### GetIncludeVehicleOk

`func (o *RouteOptionsModel) GetIncludeVehicleOk() (*bool, bool)`

GetIncludeVehicleOk returns a tuple with the IncludeVehicle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeVehicle

`func (o *RouteOptionsModel) SetIncludeVehicle(v bool)`

SetIncludeVehicle sets IncludeVehicle field to given value.

### HasIncludeVehicle

`func (o *RouteOptionsModel) HasIncludeVehicle() bool`

HasIncludeVehicle returns a boolean if a field has been set.

### GetIncludeVehicleLinks

`func (o *RouteOptionsModel) GetIncludeVehicleLinks() bool`

GetIncludeVehicleLinks returns the IncludeVehicleLinks field if non-nil, zero value otherwise.

### GetIncludeVehicleLinksOk

`func (o *RouteOptionsModel) GetIncludeVehicleLinksOk() (*bool, bool)`

GetIncludeVehicleLinksOk returns a tuple with the IncludeVehicleLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeVehicleLinks

`func (o *RouteOptionsModel) SetIncludeVehicleLinks(v bool)`

SetIncludeVehicleLinks sets IncludeVehicleLinks field to given value.

### HasIncludeVehicleLinks

`func (o *RouteOptionsModel) HasIncludeVehicleLinks() bool`

HasIncludeVehicleLinks returns a boolean if a field has been set.

### GetIncludeTrailer

`func (o *RouteOptionsModel) GetIncludeTrailer() bool`

GetIncludeTrailer returns the IncludeTrailer field if non-nil, zero value otherwise.

### GetIncludeTrailerOk

`func (o *RouteOptionsModel) GetIncludeTrailerOk() (*bool, bool)`

GetIncludeTrailerOk returns a tuple with the IncludeTrailer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTrailer

`func (o *RouteOptionsModel) SetIncludeTrailer(v bool)`

SetIncludeTrailer sets IncludeTrailer field to given value.

### HasIncludeTrailer

`func (o *RouteOptionsModel) HasIncludeTrailer() bool`

HasIncludeTrailer returns a boolean if a field has been set.

### GetIncludeTrailerLinks

`func (o *RouteOptionsModel) GetIncludeTrailerLinks() bool`

GetIncludeTrailerLinks returns the IncludeTrailerLinks field if non-nil, zero value otherwise.

### GetIncludeTrailerLinksOk

`func (o *RouteOptionsModel) GetIncludeTrailerLinksOk() (*bool, bool)`

GetIncludeTrailerLinksOk returns a tuple with the IncludeTrailerLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTrailerLinks

`func (o *RouteOptionsModel) SetIncludeTrailerLinks(v bool)`

SetIncludeTrailerLinks sets IncludeTrailerLinks field to given value.

### HasIncludeTrailerLinks

`func (o *RouteOptionsModel) HasIncludeTrailerLinks() bool`

HasIncludeTrailerLinks returns a boolean if a field has been set.

### GetIncludeDriverInfo

`func (o *RouteOptionsModel) GetIncludeDriverInfo() bool`

GetIncludeDriverInfo returns the IncludeDriverInfo field if non-nil, zero value otherwise.

### GetIncludeDriverInfoOk

`func (o *RouteOptionsModel) GetIncludeDriverInfoOk() (*bool, bool)`

GetIncludeDriverInfoOk returns a tuple with the IncludeDriverInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDriverInfo

`func (o *RouteOptionsModel) SetIncludeDriverInfo(v bool)`

SetIncludeDriverInfo sets IncludeDriverInfo field to given value.

### HasIncludeDriverInfo

`func (o *RouteOptionsModel) HasIncludeDriverInfo() bool`

HasIncludeDriverInfo returns a boolean if a field has been set.

### GetIncludeEquipmentInfoCar

`func (o *RouteOptionsModel) GetIncludeEquipmentInfoCar() bool`

GetIncludeEquipmentInfoCar returns the IncludeEquipmentInfoCar field if non-nil, zero value otherwise.

### GetIncludeEquipmentInfoCarOk

`func (o *RouteOptionsModel) GetIncludeEquipmentInfoCarOk() (*bool, bool)`

GetIncludeEquipmentInfoCarOk returns a tuple with the IncludeEquipmentInfoCar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeEquipmentInfoCar

`func (o *RouteOptionsModel) SetIncludeEquipmentInfoCar(v bool)`

SetIncludeEquipmentInfoCar sets IncludeEquipmentInfoCar field to given value.

### HasIncludeEquipmentInfoCar

`func (o *RouteOptionsModel) HasIncludeEquipmentInfoCar() bool`

HasIncludeEquipmentInfoCar returns a boolean if a field has been set.

### GetIncludeEquipment

`func (o *RouteOptionsModel) GetIncludeEquipment() bool`

GetIncludeEquipment returns the IncludeEquipment field if non-nil, zero value otherwise.

### GetIncludeEquipmentOk

`func (o *RouteOptionsModel) GetIncludeEquipmentOk() (*bool, bool)`

GetIncludeEquipmentOk returns a tuple with the IncludeEquipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeEquipment

`func (o *RouteOptionsModel) SetIncludeEquipment(v bool)`

SetIncludeEquipment sets IncludeEquipment field to given value.

### HasIncludeEquipment

`func (o *RouteOptionsModel) HasIncludeEquipment() bool`

HasIncludeEquipment returns a boolean if a field has been set.

### GetIncludeEquipmentIds

`func (o *RouteOptionsModel) GetIncludeEquipmentIds() bool`

GetIncludeEquipmentIds returns the IncludeEquipmentIds field if non-nil, zero value otherwise.

### GetIncludeEquipmentIdsOk

`func (o *RouteOptionsModel) GetIncludeEquipmentIdsOk() (*bool, bool)`

GetIncludeEquipmentIdsOk returns a tuple with the IncludeEquipmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeEquipmentIds

`func (o *RouteOptionsModel) SetIncludeEquipmentIds(v bool)`

SetIncludeEquipmentIds sets IncludeEquipmentIds field to given value.

### HasIncludeEquipmentIds

`func (o *RouteOptionsModel) HasIncludeEquipmentIds() bool`

HasIncludeEquipmentIds returns a boolean if a field has been set.

### GetIncludeGpsLocations

`func (o *RouteOptionsModel) GetIncludeGpsLocations() bool`

GetIncludeGpsLocations returns the IncludeGpsLocations field if non-nil, zero value otherwise.

### GetIncludeGpsLocationsOk

`func (o *RouteOptionsModel) GetIncludeGpsLocationsOk() (*bool, bool)`

GetIncludeGpsLocationsOk returns a tuple with the IncludeGpsLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeGpsLocations

`func (o *RouteOptionsModel) SetIncludeGpsLocations(v bool)`

SetIncludeGpsLocations sets IncludeGpsLocations field to given value.

### HasIncludeGpsLocations

`func (o *RouteOptionsModel) HasIncludeGpsLocations() bool`

HasIncludeGpsLocations returns a boolean if a field has been set.

### GetIncludePause

`func (o *RouteOptionsModel) GetIncludePause() bool`

GetIncludePause returns the IncludePause field if non-nil, zero value otherwise.

### GetIncludePauseOk

`func (o *RouteOptionsModel) GetIncludePauseOk() (*bool, bool)`

GetIncludePauseOk returns a tuple with the IncludePause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePause

`func (o *RouteOptionsModel) SetIncludePause(v bool)`

SetIncludePause sets IncludePause field to given value.

### HasIncludePause

`func (o *RouteOptionsModel) HasIncludePause() bool`

HasIncludePause returns a boolean if a field has been set.

### GetIncludeActivityIds

`func (o *RouteOptionsModel) GetIncludeActivityIds() bool`

GetIncludeActivityIds returns the IncludeActivityIds field if non-nil, zero value otherwise.

### GetIncludeActivityIdsOk

`func (o *RouteOptionsModel) GetIncludeActivityIdsOk() (*bool, bool)`

GetIncludeActivityIdsOk returns a tuple with the IncludeActivityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeActivityIds

`func (o *RouteOptionsModel) SetIncludeActivityIds(v bool)`

SetIncludeActivityIds sets IncludeActivityIds field to given value.

### HasIncludeActivityIds

`func (o *RouteOptionsModel) HasIncludeActivityIds() bool`

HasIncludeActivityIds returns a boolean if a field has been set.

### GetIncludeLatestPosition

`func (o *RouteOptionsModel) GetIncludeLatestPosition() bool`

GetIncludeLatestPosition returns the IncludeLatestPosition field if non-nil, zero value otherwise.

### GetIncludeLatestPositionOk

`func (o *RouteOptionsModel) GetIncludeLatestPositionOk() (*bool, bool)`

GetIncludeLatestPositionOk returns a tuple with the IncludeLatestPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLatestPosition

`func (o *RouteOptionsModel) SetIncludeLatestPosition(v bool)`

SetIncludeLatestPosition sets IncludeLatestPosition field to given value.

### HasIncludeLatestPosition

`func (o *RouteOptionsModel) HasIncludeLatestPosition() bool`

HasIncludeLatestPosition returns a boolean if a field has been set.

### GetIncludeZones

`func (o *RouteOptionsModel) GetIncludeZones() bool`

GetIncludeZones returns the IncludeZones field if non-nil, zero value otherwise.

### GetIncludeZonesOk

`func (o *RouteOptionsModel) GetIncludeZonesOk() (*bool, bool)`

GetIncludeZonesOk returns a tuple with the IncludeZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeZones

`func (o *RouteOptionsModel) SetIncludeZones(v bool)`

SetIncludeZones sets IncludeZones field to given value.

### HasIncludeZones

`func (o *RouteOptionsModel) HasIncludeZones() bool`

HasIncludeZones returns a boolean if a field has been set.

### GetIncludeZoneNames

`func (o *RouteOptionsModel) GetIncludeZoneNames() bool`

GetIncludeZoneNames returns the IncludeZoneNames field if non-nil, zero value otherwise.

### GetIncludeZoneNamesOk

`func (o *RouteOptionsModel) GetIncludeZoneNamesOk() (*bool, bool)`

GetIncludeZoneNamesOk returns a tuple with the IncludeZoneNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeZoneNames

`func (o *RouteOptionsModel) SetIncludeZoneNames(v bool)`

SetIncludeZoneNames sets IncludeZoneNames field to given value.

### HasIncludeZoneNames

`func (o *RouteOptionsModel) HasIncludeZoneNames() bool`

HasIncludeZoneNames returns a boolean if a field has been set.

### GetIncludeNotes

`func (o *RouteOptionsModel) GetIncludeNotes() bool`

GetIncludeNotes returns the IncludeNotes field if non-nil, zero value otherwise.

### GetIncludeNotesOk

`func (o *RouteOptionsModel) GetIncludeNotesOk() (*bool, bool)`

GetIncludeNotesOk returns a tuple with the IncludeNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeNotes

`func (o *RouteOptionsModel) SetIncludeNotes(v bool)`

SetIncludeNotes sets IncludeNotes field to given value.

### HasIncludeNotes

`func (o *RouteOptionsModel) HasIncludeNotes() bool`

HasIncludeNotes returns a boolean if a field has been set.

### GetIncludeLinks

`func (o *RouteOptionsModel) GetIncludeLinks() bool`

GetIncludeLinks returns the IncludeLinks field if non-nil, zero value otherwise.

### GetIncludeLinksOk

`func (o *RouteOptionsModel) GetIncludeLinksOk() (*bool, bool)`

GetIncludeLinksOk returns a tuple with the IncludeLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLinks

`func (o *RouteOptionsModel) SetIncludeLinks(v bool)`

SetIncludeLinks sets IncludeLinks field to given value.

### HasIncludeLinks

`func (o *RouteOptionsModel) HasIncludeLinks() bool`

HasIncludeLinks returns a boolean if a field has been set.

### GetIncludeMetaData

`func (o *RouteOptionsModel) GetIncludeMetaData() bool`

GetIncludeMetaData returns the IncludeMetaData field if non-nil, zero value otherwise.

### GetIncludeMetaDataOk

`func (o *RouteOptionsModel) GetIncludeMetaDataOk() (*bool, bool)`

GetIncludeMetaDataOk returns a tuple with the IncludeMetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMetaData

`func (o *RouteOptionsModel) SetIncludeMetaData(v bool)`

SetIncludeMetaData sets IncludeMetaData field to given value.

### HasIncludeMetaData

`func (o *RouteOptionsModel) HasIncludeMetaData() bool`

HasIncludeMetaData returns a boolean if a field has been set.

### GetIncludeAppliedCapacities

`func (o *RouteOptionsModel) GetIncludeAppliedCapacities() bool`

GetIncludeAppliedCapacities returns the IncludeAppliedCapacities field if non-nil, zero value otherwise.

### GetIncludeAppliedCapacitiesOk

`func (o *RouteOptionsModel) GetIncludeAppliedCapacitiesOk() (*bool, bool)`

GetIncludeAppliedCapacitiesOk returns a tuple with the IncludeAppliedCapacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAppliedCapacities

`func (o *RouteOptionsModel) SetIncludeAppliedCapacities(v bool)`

SetIncludeAppliedCapacities sets IncludeAppliedCapacities field to given value.

### HasIncludeAppliedCapacities

`func (o *RouteOptionsModel) HasIncludeAppliedCapacities() bool`

HasIncludeAppliedCapacities returns a boolean if a field has been set.

### GetIncludeAppliedCapacitiesV2

`func (o *RouteOptionsModel) GetIncludeAppliedCapacitiesV2() bool`

GetIncludeAppliedCapacitiesV2 returns the IncludeAppliedCapacitiesV2 field if non-nil, zero value otherwise.

### GetIncludeAppliedCapacitiesV2Ok

`func (o *RouteOptionsModel) GetIncludeAppliedCapacitiesV2Ok() (*bool, bool)`

GetIncludeAppliedCapacitiesV2Ok returns a tuple with the IncludeAppliedCapacitiesV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAppliedCapacitiesV2

`func (o *RouteOptionsModel) SetIncludeAppliedCapacitiesV2(v bool)`

SetIncludeAppliedCapacitiesV2 sets IncludeAppliedCapacitiesV2 field to given value.

### HasIncludeAppliedCapacitiesV2

`func (o *RouteOptionsModel) HasIncludeAppliedCapacitiesV2() bool`

HasIncludeAppliedCapacitiesV2 returns a boolean if a field has been set.

### GetIncludeCapacities

`func (o *RouteOptionsModel) GetIncludeCapacities() bool`

GetIncludeCapacities returns the IncludeCapacities field if non-nil, zero value otherwise.

### GetIncludeCapacitiesOk

`func (o *RouteOptionsModel) GetIncludeCapacitiesOk() (*bool, bool)`

GetIncludeCapacitiesOk returns a tuple with the IncludeCapacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCapacities

`func (o *RouteOptionsModel) SetIncludeCapacities(v bool)`

SetIncludeCapacities sets IncludeCapacities field to given value.

### HasIncludeCapacities

`func (o *RouteOptionsModel) HasIncludeCapacities() bool`

HasIncludeCapacities returns a boolean if a field has been set.

### GetIncludeRecurrence

`func (o *RouteOptionsModel) GetIncludeRecurrence() bool`

GetIncludeRecurrence returns the IncludeRecurrence field if non-nil, zero value otherwise.

### GetIncludeRecurrenceOk

`func (o *RouteOptionsModel) GetIncludeRecurrenceOk() (*bool, bool)`

GetIncludeRecurrenceOk returns a tuple with the IncludeRecurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRecurrence

`func (o *RouteOptionsModel) SetIncludeRecurrence(v bool)`

SetIncludeRecurrence sets IncludeRecurrence field to given value.

### HasIncludeRecurrence

`func (o *RouteOptionsModel) HasIncludeRecurrence() bool`

HasIncludeRecurrence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


