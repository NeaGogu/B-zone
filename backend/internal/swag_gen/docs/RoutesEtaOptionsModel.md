# RoutesEtaOptionsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeAddress** | Pointer to **bool** |  | [optional] 
**IncludeAddressObject** | Pointer to **bool** |  | [optional] 
**IncludeRouteStatus** | Pointer to **bool** |  | [optional] 
**IncludeRouteTags** | Pointer to **bool** | Deprecated! use include_tags | [optional] 
**IncludeTags** | Pointer to **bool** |  | [optional] 
**IncludeTagNames** | Pointer to **bool** |  | [optional] 
**IncludeDriver** | Pointer to **bool** |  | [optional] 
**IncludeDriverLinks** | Pointer to **bool** |  | [optional] 
**IncludeCar** | Pointer to **bool** |  | [optional] 
**IncludeCarLinks** | Pointer to **bool** |  | [optional] 
**IncludeVehicle** | Pointer to **bool** |  | [optional] 
**IncludeVehicleLinks** | Pointer to **bool** |  | [optional] 
**IncludeTrailer** | Pointer to **bool** |  | [optional] 
**IncludeTrailerLinks** | Pointer to **bool** |  | [optional] 
**IncludeDriverInfo** | Pointer to **bool** |  | [optional] 
**IncludeEquipmentInfoCar** | Pointer to **bool** |  | [optional] 
**IncludeEquipment** | Pointer to **bool** |  | [optional] 
**IncludeGpsLocations** | Pointer to **bool** |  | [optional] 
**IncludePause** | Pointer to **bool** |  | [optional] 
**IncludeActivityIds** | Pointer to **bool** |  | [optional] 
**IncludeLatestPosition** | Pointer to **bool** |  | [optional] 
**IncludeZones** | Pointer to **bool** |  | [optional] 
**IncludeZoneNames** | Pointer to **bool** |  | [optional] 
**IncludeNotes** | Pointer to **bool** |  | [optional] 
**IncludeMetaData** | Pointer to **bool** |  | [optional] 

## Methods

### NewRoutesEtaOptionsModel

`func NewRoutesEtaOptionsModel() *RoutesEtaOptionsModel`

NewRoutesEtaOptionsModel instantiates a new RoutesEtaOptionsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutesEtaOptionsModelWithDefaults

`func NewRoutesEtaOptionsModelWithDefaults() *RoutesEtaOptionsModel`

NewRoutesEtaOptionsModelWithDefaults instantiates a new RoutesEtaOptionsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeAddress

`func (o *RoutesEtaOptionsModel) GetIncludeAddress() bool`

GetIncludeAddress returns the IncludeAddress field if non-nil, zero value otherwise.

### GetIncludeAddressOk

`func (o *RoutesEtaOptionsModel) GetIncludeAddressOk() (*bool, bool)`

GetIncludeAddressOk returns a tuple with the IncludeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAddress

`func (o *RoutesEtaOptionsModel) SetIncludeAddress(v bool)`

SetIncludeAddress sets IncludeAddress field to given value.

### HasIncludeAddress

`func (o *RoutesEtaOptionsModel) HasIncludeAddress() bool`

HasIncludeAddress returns a boolean if a field has been set.

### GetIncludeAddressObject

`func (o *RoutesEtaOptionsModel) GetIncludeAddressObject() bool`

GetIncludeAddressObject returns the IncludeAddressObject field if non-nil, zero value otherwise.

### GetIncludeAddressObjectOk

`func (o *RoutesEtaOptionsModel) GetIncludeAddressObjectOk() (*bool, bool)`

GetIncludeAddressObjectOk returns a tuple with the IncludeAddressObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAddressObject

`func (o *RoutesEtaOptionsModel) SetIncludeAddressObject(v bool)`

SetIncludeAddressObject sets IncludeAddressObject field to given value.

### HasIncludeAddressObject

`func (o *RoutesEtaOptionsModel) HasIncludeAddressObject() bool`

HasIncludeAddressObject returns a boolean if a field has been set.

### GetIncludeRouteStatus

`func (o *RoutesEtaOptionsModel) GetIncludeRouteStatus() bool`

GetIncludeRouteStatus returns the IncludeRouteStatus field if non-nil, zero value otherwise.

### GetIncludeRouteStatusOk

`func (o *RoutesEtaOptionsModel) GetIncludeRouteStatusOk() (*bool, bool)`

GetIncludeRouteStatusOk returns a tuple with the IncludeRouteStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRouteStatus

`func (o *RoutesEtaOptionsModel) SetIncludeRouteStatus(v bool)`

SetIncludeRouteStatus sets IncludeRouteStatus field to given value.

### HasIncludeRouteStatus

`func (o *RoutesEtaOptionsModel) HasIncludeRouteStatus() bool`

HasIncludeRouteStatus returns a boolean if a field has been set.

### GetIncludeRouteTags

`func (o *RoutesEtaOptionsModel) GetIncludeRouteTags() bool`

GetIncludeRouteTags returns the IncludeRouteTags field if non-nil, zero value otherwise.

### GetIncludeRouteTagsOk

`func (o *RoutesEtaOptionsModel) GetIncludeRouteTagsOk() (*bool, bool)`

GetIncludeRouteTagsOk returns a tuple with the IncludeRouteTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRouteTags

`func (o *RoutesEtaOptionsModel) SetIncludeRouteTags(v bool)`

SetIncludeRouteTags sets IncludeRouteTags field to given value.

### HasIncludeRouteTags

`func (o *RoutesEtaOptionsModel) HasIncludeRouteTags() bool`

HasIncludeRouteTags returns a boolean if a field has been set.

### GetIncludeTags

`func (o *RoutesEtaOptionsModel) GetIncludeTags() bool`

GetIncludeTags returns the IncludeTags field if non-nil, zero value otherwise.

### GetIncludeTagsOk

`func (o *RoutesEtaOptionsModel) GetIncludeTagsOk() (*bool, bool)`

GetIncludeTagsOk returns a tuple with the IncludeTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTags

`func (o *RoutesEtaOptionsModel) SetIncludeTags(v bool)`

SetIncludeTags sets IncludeTags field to given value.

### HasIncludeTags

`func (o *RoutesEtaOptionsModel) HasIncludeTags() bool`

HasIncludeTags returns a boolean if a field has been set.

### GetIncludeTagNames

`func (o *RoutesEtaOptionsModel) GetIncludeTagNames() bool`

GetIncludeTagNames returns the IncludeTagNames field if non-nil, zero value otherwise.

### GetIncludeTagNamesOk

`func (o *RoutesEtaOptionsModel) GetIncludeTagNamesOk() (*bool, bool)`

GetIncludeTagNamesOk returns a tuple with the IncludeTagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTagNames

`func (o *RoutesEtaOptionsModel) SetIncludeTagNames(v bool)`

SetIncludeTagNames sets IncludeTagNames field to given value.

### HasIncludeTagNames

`func (o *RoutesEtaOptionsModel) HasIncludeTagNames() bool`

HasIncludeTagNames returns a boolean if a field has been set.

### GetIncludeDriver

`func (o *RoutesEtaOptionsModel) GetIncludeDriver() bool`

GetIncludeDriver returns the IncludeDriver field if non-nil, zero value otherwise.

### GetIncludeDriverOk

`func (o *RoutesEtaOptionsModel) GetIncludeDriverOk() (*bool, bool)`

GetIncludeDriverOk returns a tuple with the IncludeDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDriver

`func (o *RoutesEtaOptionsModel) SetIncludeDriver(v bool)`

SetIncludeDriver sets IncludeDriver field to given value.

### HasIncludeDriver

`func (o *RoutesEtaOptionsModel) HasIncludeDriver() bool`

HasIncludeDriver returns a boolean if a field has been set.

### GetIncludeDriverLinks

`func (o *RoutesEtaOptionsModel) GetIncludeDriverLinks() bool`

GetIncludeDriverLinks returns the IncludeDriverLinks field if non-nil, zero value otherwise.

### GetIncludeDriverLinksOk

`func (o *RoutesEtaOptionsModel) GetIncludeDriverLinksOk() (*bool, bool)`

GetIncludeDriverLinksOk returns a tuple with the IncludeDriverLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDriverLinks

`func (o *RoutesEtaOptionsModel) SetIncludeDriverLinks(v bool)`

SetIncludeDriverLinks sets IncludeDriverLinks field to given value.

### HasIncludeDriverLinks

`func (o *RoutesEtaOptionsModel) HasIncludeDriverLinks() bool`

HasIncludeDriverLinks returns a boolean if a field has been set.

### GetIncludeCar

`func (o *RoutesEtaOptionsModel) GetIncludeCar() bool`

GetIncludeCar returns the IncludeCar field if non-nil, zero value otherwise.

### GetIncludeCarOk

`func (o *RoutesEtaOptionsModel) GetIncludeCarOk() (*bool, bool)`

GetIncludeCarOk returns a tuple with the IncludeCar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCar

`func (o *RoutesEtaOptionsModel) SetIncludeCar(v bool)`

SetIncludeCar sets IncludeCar field to given value.

### HasIncludeCar

`func (o *RoutesEtaOptionsModel) HasIncludeCar() bool`

HasIncludeCar returns a boolean if a field has been set.

### GetIncludeCarLinks

`func (o *RoutesEtaOptionsModel) GetIncludeCarLinks() bool`

GetIncludeCarLinks returns the IncludeCarLinks field if non-nil, zero value otherwise.

### GetIncludeCarLinksOk

`func (o *RoutesEtaOptionsModel) GetIncludeCarLinksOk() (*bool, bool)`

GetIncludeCarLinksOk returns a tuple with the IncludeCarLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCarLinks

`func (o *RoutesEtaOptionsModel) SetIncludeCarLinks(v bool)`

SetIncludeCarLinks sets IncludeCarLinks field to given value.

### HasIncludeCarLinks

`func (o *RoutesEtaOptionsModel) HasIncludeCarLinks() bool`

HasIncludeCarLinks returns a boolean if a field has been set.

### GetIncludeVehicle

`func (o *RoutesEtaOptionsModel) GetIncludeVehicle() bool`

GetIncludeVehicle returns the IncludeVehicle field if non-nil, zero value otherwise.

### GetIncludeVehicleOk

`func (o *RoutesEtaOptionsModel) GetIncludeVehicleOk() (*bool, bool)`

GetIncludeVehicleOk returns a tuple with the IncludeVehicle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeVehicle

`func (o *RoutesEtaOptionsModel) SetIncludeVehicle(v bool)`

SetIncludeVehicle sets IncludeVehicle field to given value.

### HasIncludeVehicle

`func (o *RoutesEtaOptionsModel) HasIncludeVehicle() bool`

HasIncludeVehicle returns a boolean if a field has been set.

### GetIncludeVehicleLinks

`func (o *RoutesEtaOptionsModel) GetIncludeVehicleLinks() bool`

GetIncludeVehicleLinks returns the IncludeVehicleLinks field if non-nil, zero value otherwise.

### GetIncludeVehicleLinksOk

`func (o *RoutesEtaOptionsModel) GetIncludeVehicleLinksOk() (*bool, bool)`

GetIncludeVehicleLinksOk returns a tuple with the IncludeVehicleLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeVehicleLinks

`func (o *RoutesEtaOptionsModel) SetIncludeVehicleLinks(v bool)`

SetIncludeVehicleLinks sets IncludeVehicleLinks field to given value.

### HasIncludeVehicleLinks

`func (o *RoutesEtaOptionsModel) HasIncludeVehicleLinks() bool`

HasIncludeVehicleLinks returns a boolean if a field has been set.

### GetIncludeTrailer

`func (o *RoutesEtaOptionsModel) GetIncludeTrailer() bool`

GetIncludeTrailer returns the IncludeTrailer field if non-nil, zero value otherwise.

### GetIncludeTrailerOk

`func (o *RoutesEtaOptionsModel) GetIncludeTrailerOk() (*bool, bool)`

GetIncludeTrailerOk returns a tuple with the IncludeTrailer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTrailer

`func (o *RoutesEtaOptionsModel) SetIncludeTrailer(v bool)`

SetIncludeTrailer sets IncludeTrailer field to given value.

### HasIncludeTrailer

`func (o *RoutesEtaOptionsModel) HasIncludeTrailer() bool`

HasIncludeTrailer returns a boolean if a field has been set.

### GetIncludeTrailerLinks

`func (o *RoutesEtaOptionsModel) GetIncludeTrailerLinks() bool`

GetIncludeTrailerLinks returns the IncludeTrailerLinks field if non-nil, zero value otherwise.

### GetIncludeTrailerLinksOk

`func (o *RoutesEtaOptionsModel) GetIncludeTrailerLinksOk() (*bool, bool)`

GetIncludeTrailerLinksOk returns a tuple with the IncludeTrailerLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTrailerLinks

`func (o *RoutesEtaOptionsModel) SetIncludeTrailerLinks(v bool)`

SetIncludeTrailerLinks sets IncludeTrailerLinks field to given value.

### HasIncludeTrailerLinks

`func (o *RoutesEtaOptionsModel) HasIncludeTrailerLinks() bool`

HasIncludeTrailerLinks returns a boolean if a field has been set.

### GetIncludeDriverInfo

`func (o *RoutesEtaOptionsModel) GetIncludeDriverInfo() bool`

GetIncludeDriverInfo returns the IncludeDriverInfo field if non-nil, zero value otherwise.

### GetIncludeDriverInfoOk

`func (o *RoutesEtaOptionsModel) GetIncludeDriverInfoOk() (*bool, bool)`

GetIncludeDriverInfoOk returns a tuple with the IncludeDriverInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDriverInfo

`func (o *RoutesEtaOptionsModel) SetIncludeDriverInfo(v bool)`

SetIncludeDriverInfo sets IncludeDriverInfo field to given value.

### HasIncludeDriverInfo

`func (o *RoutesEtaOptionsModel) HasIncludeDriverInfo() bool`

HasIncludeDriverInfo returns a boolean if a field has been set.

### GetIncludeEquipmentInfoCar

`func (o *RoutesEtaOptionsModel) GetIncludeEquipmentInfoCar() bool`

GetIncludeEquipmentInfoCar returns the IncludeEquipmentInfoCar field if non-nil, zero value otherwise.

### GetIncludeEquipmentInfoCarOk

`func (o *RoutesEtaOptionsModel) GetIncludeEquipmentInfoCarOk() (*bool, bool)`

GetIncludeEquipmentInfoCarOk returns a tuple with the IncludeEquipmentInfoCar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeEquipmentInfoCar

`func (o *RoutesEtaOptionsModel) SetIncludeEquipmentInfoCar(v bool)`

SetIncludeEquipmentInfoCar sets IncludeEquipmentInfoCar field to given value.

### HasIncludeEquipmentInfoCar

`func (o *RoutesEtaOptionsModel) HasIncludeEquipmentInfoCar() bool`

HasIncludeEquipmentInfoCar returns a boolean if a field has been set.

### GetIncludeEquipment

`func (o *RoutesEtaOptionsModel) GetIncludeEquipment() bool`

GetIncludeEquipment returns the IncludeEquipment field if non-nil, zero value otherwise.

### GetIncludeEquipmentOk

`func (o *RoutesEtaOptionsModel) GetIncludeEquipmentOk() (*bool, bool)`

GetIncludeEquipmentOk returns a tuple with the IncludeEquipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeEquipment

`func (o *RoutesEtaOptionsModel) SetIncludeEquipment(v bool)`

SetIncludeEquipment sets IncludeEquipment field to given value.

### HasIncludeEquipment

`func (o *RoutesEtaOptionsModel) HasIncludeEquipment() bool`

HasIncludeEquipment returns a boolean if a field has been set.

### GetIncludeGpsLocations

`func (o *RoutesEtaOptionsModel) GetIncludeGpsLocations() bool`

GetIncludeGpsLocations returns the IncludeGpsLocations field if non-nil, zero value otherwise.

### GetIncludeGpsLocationsOk

`func (o *RoutesEtaOptionsModel) GetIncludeGpsLocationsOk() (*bool, bool)`

GetIncludeGpsLocationsOk returns a tuple with the IncludeGpsLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeGpsLocations

`func (o *RoutesEtaOptionsModel) SetIncludeGpsLocations(v bool)`

SetIncludeGpsLocations sets IncludeGpsLocations field to given value.

### HasIncludeGpsLocations

`func (o *RoutesEtaOptionsModel) HasIncludeGpsLocations() bool`

HasIncludeGpsLocations returns a boolean if a field has been set.

### GetIncludePause

`func (o *RoutesEtaOptionsModel) GetIncludePause() bool`

GetIncludePause returns the IncludePause field if non-nil, zero value otherwise.

### GetIncludePauseOk

`func (o *RoutesEtaOptionsModel) GetIncludePauseOk() (*bool, bool)`

GetIncludePauseOk returns a tuple with the IncludePause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePause

`func (o *RoutesEtaOptionsModel) SetIncludePause(v bool)`

SetIncludePause sets IncludePause field to given value.

### HasIncludePause

`func (o *RoutesEtaOptionsModel) HasIncludePause() bool`

HasIncludePause returns a boolean if a field has been set.

### GetIncludeActivityIds

`func (o *RoutesEtaOptionsModel) GetIncludeActivityIds() bool`

GetIncludeActivityIds returns the IncludeActivityIds field if non-nil, zero value otherwise.

### GetIncludeActivityIdsOk

`func (o *RoutesEtaOptionsModel) GetIncludeActivityIdsOk() (*bool, bool)`

GetIncludeActivityIdsOk returns a tuple with the IncludeActivityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeActivityIds

`func (o *RoutesEtaOptionsModel) SetIncludeActivityIds(v bool)`

SetIncludeActivityIds sets IncludeActivityIds field to given value.

### HasIncludeActivityIds

`func (o *RoutesEtaOptionsModel) HasIncludeActivityIds() bool`

HasIncludeActivityIds returns a boolean if a field has been set.

### GetIncludeLatestPosition

`func (o *RoutesEtaOptionsModel) GetIncludeLatestPosition() bool`

GetIncludeLatestPosition returns the IncludeLatestPosition field if non-nil, zero value otherwise.

### GetIncludeLatestPositionOk

`func (o *RoutesEtaOptionsModel) GetIncludeLatestPositionOk() (*bool, bool)`

GetIncludeLatestPositionOk returns a tuple with the IncludeLatestPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLatestPosition

`func (o *RoutesEtaOptionsModel) SetIncludeLatestPosition(v bool)`

SetIncludeLatestPosition sets IncludeLatestPosition field to given value.

### HasIncludeLatestPosition

`func (o *RoutesEtaOptionsModel) HasIncludeLatestPosition() bool`

HasIncludeLatestPosition returns a boolean if a field has been set.

### GetIncludeZones

`func (o *RoutesEtaOptionsModel) GetIncludeZones() bool`

GetIncludeZones returns the IncludeZones field if non-nil, zero value otherwise.

### GetIncludeZonesOk

`func (o *RoutesEtaOptionsModel) GetIncludeZonesOk() (*bool, bool)`

GetIncludeZonesOk returns a tuple with the IncludeZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeZones

`func (o *RoutesEtaOptionsModel) SetIncludeZones(v bool)`

SetIncludeZones sets IncludeZones field to given value.

### HasIncludeZones

`func (o *RoutesEtaOptionsModel) HasIncludeZones() bool`

HasIncludeZones returns a boolean if a field has been set.

### GetIncludeZoneNames

`func (o *RoutesEtaOptionsModel) GetIncludeZoneNames() bool`

GetIncludeZoneNames returns the IncludeZoneNames field if non-nil, zero value otherwise.

### GetIncludeZoneNamesOk

`func (o *RoutesEtaOptionsModel) GetIncludeZoneNamesOk() (*bool, bool)`

GetIncludeZoneNamesOk returns a tuple with the IncludeZoneNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeZoneNames

`func (o *RoutesEtaOptionsModel) SetIncludeZoneNames(v bool)`

SetIncludeZoneNames sets IncludeZoneNames field to given value.

### HasIncludeZoneNames

`func (o *RoutesEtaOptionsModel) HasIncludeZoneNames() bool`

HasIncludeZoneNames returns a boolean if a field has been set.

### GetIncludeNotes

`func (o *RoutesEtaOptionsModel) GetIncludeNotes() bool`

GetIncludeNotes returns the IncludeNotes field if non-nil, zero value otherwise.

### GetIncludeNotesOk

`func (o *RoutesEtaOptionsModel) GetIncludeNotesOk() (*bool, bool)`

GetIncludeNotesOk returns a tuple with the IncludeNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeNotes

`func (o *RoutesEtaOptionsModel) SetIncludeNotes(v bool)`

SetIncludeNotes sets IncludeNotes field to given value.

### HasIncludeNotes

`func (o *RoutesEtaOptionsModel) HasIncludeNotes() bool`

HasIncludeNotes returns a boolean if a field has been set.

### GetIncludeMetaData

`func (o *RoutesEtaOptionsModel) GetIncludeMetaData() bool`

GetIncludeMetaData returns the IncludeMetaData field if non-nil, zero value otherwise.

### GetIncludeMetaDataOk

`func (o *RoutesEtaOptionsModel) GetIncludeMetaDataOk() (*bool, bool)`

GetIncludeMetaDataOk returns a tuple with the IncludeMetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMetaData

`func (o *RoutesEtaOptionsModel) SetIncludeMetaData(v bool)`

SetIncludeMetaData sets IncludeMetaData field to given value.

### HasIncludeMetaData

`func (o *RoutesEtaOptionsModel) HasIncludeMetaData() bool`

HasIncludeMetaData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


