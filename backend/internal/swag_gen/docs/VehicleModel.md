# VehicleModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique Identifier | [optional] [readonly] 
**VehicleTypeId** | Pointer to **int64** | Bumbal id for vehicle_type | [optional] [readonly] 
**VehicleTypeName** | Pointer to **string** | Bumbal id for vehicle_type | [optional] 
**Name** | Pointer to **string** | name | [optional] 
**RegistrationNr** | Pointer to **string** | registration_nr | [optional] 
**MaxSpeed** | Pointer to **int64** | Max Speed in km/h | [optional] 
**SpeedFactor** | Pointer to **float64** | Speed Factor | [optional] 
**StartDuration** | Pointer to **int64** | Start duration of using this vehicle in minutes | [optional] 
**EndDuration** | Pointer to **int64** | End duration of using this vehicle in minutes | [optional] 
**DurationFactor** | Pointer to **float64** | Duration Factor | [optional] 
**CostPerMeter** | Pointer to **float64** | Cost per meter | [optional] 
**CostPerRoute** | Pointer to **float64** | Cost per route | [optional] 
**CostPerDrivingMinute** | Pointer to **float64** | Cost per driving minute | [optional] 
**CostPerWaitingMinute** | Pointer to **float64** | Cost per waiting minute | [optional] 
**CostPerServiceMinute** | Pointer to **float64** | Cost per service minute | [optional] 
**AppliedCapacities** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**Capacities** | Pointer to [**[]CapacityModel**](CapacityModel.md) |  | [optional] 
**Tags** | Pointer to [**[]TagModel**](TagModel.md) |  | [optional] 
**MetaData** | Pointer to [**[]MetaDataModel**](MetaDataModel.md) |  | [optional] 
**Links** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**Files** | Pointer to [**[]FileModel**](FileModel.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | created_at date time | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | updated_at date time | [optional] [readonly] 
**UpdatedByName** | Pointer to **string** | Vehicle updated by user full name | [optional] 

## Methods

### NewVehicleModel

`func NewVehicleModel() *VehicleModel`

NewVehicleModel instantiates a new VehicleModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleModelWithDefaults

`func NewVehicleModelWithDefaults() *VehicleModel`

NewVehicleModelWithDefaults instantiates a new VehicleModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VehicleModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VehicleModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VehicleModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *VehicleModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVehicleTypeId

`func (o *VehicleModel) GetVehicleTypeId() int64`

GetVehicleTypeId returns the VehicleTypeId field if non-nil, zero value otherwise.

### GetVehicleTypeIdOk

`func (o *VehicleModel) GetVehicleTypeIdOk() (*int64, bool)`

GetVehicleTypeIdOk returns a tuple with the VehicleTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleTypeId

`func (o *VehicleModel) SetVehicleTypeId(v int64)`

SetVehicleTypeId sets VehicleTypeId field to given value.

### HasVehicleTypeId

`func (o *VehicleModel) HasVehicleTypeId() bool`

HasVehicleTypeId returns a boolean if a field has been set.

### GetVehicleTypeName

`func (o *VehicleModel) GetVehicleTypeName() string`

GetVehicleTypeName returns the VehicleTypeName field if non-nil, zero value otherwise.

### GetVehicleTypeNameOk

`func (o *VehicleModel) GetVehicleTypeNameOk() (*string, bool)`

GetVehicleTypeNameOk returns a tuple with the VehicleTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleTypeName

`func (o *VehicleModel) SetVehicleTypeName(v string)`

SetVehicleTypeName sets VehicleTypeName field to given value.

### HasVehicleTypeName

`func (o *VehicleModel) HasVehicleTypeName() bool`

HasVehicleTypeName returns a boolean if a field has been set.

### GetName

`func (o *VehicleModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VehicleModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VehicleModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VehicleModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegistrationNr

`func (o *VehicleModel) GetRegistrationNr() string`

GetRegistrationNr returns the RegistrationNr field if non-nil, zero value otherwise.

### GetRegistrationNrOk

`func (o *VehicleModel) GetRegistrationNrOk() (*string, bool)`

GetRegistrationNrOk returns a tuple with the RegistrationNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationNr

`func (o *VehicleModel) SetRegistrationNr(v string)`

SetRegistrationNr sets RegistrationNr field to given value.

### HasRegistrationNr

`func (o *VehicleModel) HasRegistrationNr() bool`

HasRegistrationNr returns a boolean if a field has been set.

### GetMaxSpeed

`func (o *VehicleModel) GetMaxSpeed() int64`

GetMaxSpeed returns the MaxSpeed field if non-nil, zero value otherwise.

### GetMaxSpeedOk

`func (o *VehicleModel) GetMaxSpeedOk() (*int64, bool)`

GetMaxSpeedOk returns a tuple with the MaxSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpeed

`func (o *VehicleModel) SetMaxSpeed(v int64)`

SetMaxSpeed sets MaxSpeed field to given value.

### HasMaxSpeed

`func (o *VehicleModel) HasMaxSpeed() bool`

HasMaxSpeed returns a boolean if a field has been set.

### GetSpeedFactor

`func (o *VehicleModel) GetSpeedFactor() float64`

GetSpeedFactor returns the SpeedFactor field if non-nil, zero value otherwise.

### GetSpeedFactorOk

`func (o *VehicleModel) GetSpeedFactorOk() (*float64, bool)`

GetSpeedFactorOk returns a tuple with the SpeedFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedFactor

`func (o *VehicleModel) SetSpeedFactor(v float64)`

SetSpeedFactor sets SpeedFactor field to given value.

### HasSpeedFactor

`func (o *VehicleModel) HasSpeedFactor() bool`

HasSpeedFactor returns a boolean if a field has been set.

### GetStartDuration

`func (o *VehicleModel) GetStartDuration() int64`

GetStartDuration returns the StartDuration field if non-nil, zero value otherwise.

### GetStartDurationOk

`func (o *VehicleModel) GetStartDurationOk() (*int64, bool)`

GetStartDurationOk returns a tuple with the StartDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDuration

`func (o *VehicleModel) SetStartDuration(v int64)`

SetStartDuration sets StartDuration field to given value.

### HasStartDuration

`func (o *VehicleModel) HasStartDuration() bool`

HasStartDuration returns a boolean if a field has been set.

### GetEndDuration

`func (o *VehicleModel) GetEndDuration() int64`

GetEndDuration returns the EndDuration field if non-nil, zero value otherwise.

### GetEndDurationOk

`func (o *VehicleModel) GetEndDurationOk() (*int64, bool)`

GetEndDurationOk returns a tuple with the EndDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDuration

`func (o *VehicleModel) SetEndDuration(v int64)`

SetEndDuration sets EndDuration field to given value.

### HasEndDuration

`func (o *VehicleModel) HasEndDuration() bool`

HasEndDuration returns a boolean if a field has been set.

### GetDurationFactor

`func (o *VehicleModel) GetDurationFactor() float64`

GetDurationFactor returns the DurationFactor field if non-nil, zero value otherwise.

### GetDurationFactorOk

`func (o *VehicleModel) GetDurationFactorOk() (*float64, bool)`

GetDurationFactorOk returns a tuple with the DurationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationFactor

`func (o *VehicleModel) SetDurationFactor(v float64)`

SetDurationFactor sets DurationFactor field to given value.

### HasDurationFactor

`func (o *VehicleModel) HasDurationFactor() bool`

HasDurationFactor returns a boolean if a field has been set.

### GetCostPerMeter

`func (o *VehicleModel) GetCostPerMeter() float64`

GetCostPerMeter returns the CostPerMeter field if non-nil, zero value otherwise.

### GetCostPerMeterOk

`func (o *VehicleModel) GetCostPerMeterOk() (*float64, bool)`

GetCostPerMeterOk returns a tuple with the CostPerMeter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerMeter

`func (o *VehicleModel) SetCostPerMeter(v float64)`

SetCostPerMeter sets CostPerMeter field to given value.

### HasCostPerMeter

`func (o *VehicleModel) HasCostPerMeter() bool`

HasCostPerMeter returns a boolean if a field has been set.

### GetCostPerRoute

`func (o *VehicleModel) GetCostPerRoute() float64`

GetCostPerRoute returns the CostPerRoute field if non-nil, zero value otherwise.

### GetCostPerRouteOk

`func (o *VehicleModel) GetCostPerRouteOk() (*float64, bool)`

GetCostPerRouteOk returns a tuple with the CostPerRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerRoute

`func (o *VehicleModel) SetCostPerRoute(v float64)`

SetCostPerRoute sets CostPerRoute field to given value.

### HasCostPerRoute

`func (o *VehicleModel) HasCostPerRoute() bool`

HasCostPerRoute returns a boolean if a field has been set.

### GetCostPerDrivingMinute

`func (o *VehicleModel) GetCostPerDrivingMinute() float64`

GetCostPerDrivingMinute returns the CostPerDrivingMinute field if non-nil, zero value otherwise.

### GetCostPerDrivingMinuteOk

`func (o *VehicleModel) GetCostPerDrivingMinuteOk() (*float64, bool)`

GetCostPerDrivingMinuteOk returns a tuple with the CostPerDrivingMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerDrivingMinute

`func (o *VehicleModel) SetCostPerDrivingMinute(v float64)`

SetCostPerDrivingMinute sets CostPerDrivingMinute field to given value.

### HasCostPerDrivingMinute

`func (o *VehicleModel) HasCostPerDrivingMinute() bool`

HasCostPerDrivingMinute returns a boolean if a field has been set.

### GetCostPerWaitingMinute

`func (o *VehicleModel) GetCostPerWaitingMinute() float64`

GetCostPerWaitingMinute returns the CostPerWaitingMinute field if non-nil, zero value otherwise.

### GetCostPerWaitingMinuteOk

`func (o *VehicleModel) GetCostPerWaitingMinuteOk() (*float64, bool)`

GetCostPerWaitingMinuteOk returns a tuple with the CostPerWaitingMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerWaitingMinute

`func (o *VehicleModel) SetCostPerWaitingMinute(v float64)`

SetCostPerWaitingMinute sets CostPerWaitingMinute field to given value.

### HasCostPerWaitingMinute

`func (o *VehicleModel) HasCostPerWaitingMinute() bool`

HasCostPerWaitingMinute returns a boolean if a field has been set.

### GetCostPerServiceMinute

`func (o *VehicleModel) GetCostPerServiceMinute() float64`

GetCostPerServiceMinute returns the CostPerServiceMinute field if non-nil, zero value otherwise.

### GetCostPerServiceMinuteOk

`func (o *VehicleModel) GetCostPerServiceMinuteOk() (*float64, bool)`

GetCostPerServiceMinuteOk returns a tuple with the CostPerServiceMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerServiceMinute

`func (o *VehicleModel) SetCostPerServiceMinute(v float64)`

SetCostPerServiceMinute sets CostPerServiceMinute field to given value.

### HasCostPerServiceMinute

`func (o *VehicleModel) HasCostPerServiceMinute() bool`

HasCostPerServiceMinute returns a boolean if a field has been set.

### GetAppliedCapacities

`func (o *VehicleModel) GetAppliedCapacities() map[string]interface{}`

GetAppliedCapacities returns the AppliedCapacities field if non-nil, zero value otherwise.

### GetAppliedCapacitiesOk

`func (o *VehicleModel) GetAppliedCapacitiesOk() (*map[string]interface{}, bool)`

GetAppliedCapacitiesOk returns a tuple with the AppliedCapacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedCapacities

`func (o *VehicleModel) SetAppliedCapacities(v map[string]interface{})`

SetAppliedCapacities sets AppliedCapacities field to given value.

### HasAppliedCapacities

`func (o *VehicleModel) HasAppliedCapacities() bool`

HasAppliedCapacities returns a boolean if a field has been set.

### GetCapacities

`func (o *VehicleModel) GetCapacities() []CapacityModel`

GetCapacities returns the Capacities field if non-nil, zero value otherwise.

### GetCapacitiesOk

`func (o *VehicleModel) GetCapacitiesOk() (*[]CapacityModel, bool)`

GetCapacitiesOk returns a tuple with the Capacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacities

`func (o *VehicleModel) SetCapacities(v []CapacityModel)`

SetCapacities sets Capacities field to given value.

### HasCapacities

`func (o *VehicleModel) HasCapacities() bool`

HasCapacities returns a boolean if a field has been set.

### GetTags

`func (o *VehicleModel) GetTags() []TagModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VehicleModel) GetTagsOk() (*[]TagModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VehicleModel) SetTags(v []TagModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VehicleModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMetaData

`func (o *VehicleModel) GetMetaData() []MetaDataModel`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *VehicleModel) GetMetaDataOk() (*[]MetaDataModel, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *VehicleModel) SetMetaData(v []MetaDataModel)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *VehicleModel) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetLinks

`func (o *VehicleModel) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VehicleModel) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VehicleModel) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *VehicleModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetFiles

`func (o *VehicleModel) GetFiles() []FileModel`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *VehicleModel) GetFilesOk() (*[]FileModel, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *VehicleModel) SetFiles(v []FileModel)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *VehicleModel) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VehicleModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VehicleModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VehicleModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VehicleModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VehicleModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VehicleModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VehicleModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VehicleModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedByName

`func (o *VehicleModel) GetUpdatedByName() string`

GetUpdatedByName returns the UpdatedByName field if non-nil, zero value otherwise.

### GetUpdatedByNameOk

`func (o *VehicleModel) GetUpdatedByNameOk() (*string, bool)`

GetUpdatedByNameOk returns a tuple with the UpdatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByName

`func (o *VehicleModel) SetUpdatedByName(v string)`

SetUpdatedByName sets UpdatedByName field to given value.

### HasUpdatedByName

`func (o *VehicleModel) HasUpdatedByName() bool`

HasUpdatedByName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


