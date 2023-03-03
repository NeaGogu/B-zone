# VehicleTypeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique ID | [optional] 
**Name** | Pointer to **string** | The name of the VehicleType | [optional] 
**MaxSpeed** | Pointer to **int64** | Max Speed in km/h, Bikes (id 4) ignore max_speed | [optional] 
**SpeedFactor** | Pointer to **float64** | Speed Factor | [optional] 
**DurationFactor** | Pointer to **float64** | Duration Factor | [optional] 
**CostPerMeter** | Pointer to **float64** | Cost per meter | [optional] 
**CostPerRoute** | Pointer to **float64** | Cost per route | [optional] 
**CostPerDrivingMinute** | Pointer to **float64** | Cost per driving minute | [optional] 
**CostPerWaitingMinute** | Pointer to **float64** | Cost per waiting minute | [optional] 
**CostPerServiceMinute** | Pointer to **float64** | Cost per service minute | [optional] 
**CreatedAt** | Pointer to **time.Time** | created_at date time | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | updated_at date time | [optional] [readonly] 

## Methods

### NewVehicleTypeModel

`func NewVehicleTypeModel() *VehicleTypeModel`

NewVehicleTypeModel instantiates a new VehicleTypeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleTypeModelWithDefaults

`func NewVehicleTypeModelWithDefaults() *VehicleTypeModel`

NewVehicleTypeModelWithDefaults instantiates a new VehicleTypeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VehicleTypeModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VehicleTypeModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VehicleTypeModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *VehicleTypeModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VehicleTypeModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VehicleTypeModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VehicleTypeModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VehicleTypeModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMaxSpeed

`func (o *VehicleTypeModel) GetMaxSpeed() int64`

GetMaxSpeed returns the MaxSpeed field if non-nil, zero value otherwise.

### GetMaxSpeedOk

`func (o *VehicleTypeModel) GetMaxSpeedOk() (*int64, bool)`

GetMaxSpeedOk returns a tuple with the MaxSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpeed

`func (o *VehicleTypeModel) SetMaxSpeed(v int64)`

SetMaxSpeed sets MaxSpeed field to given value.

### HasMaxSpeed

`func (o *VehicleTypeModel) HasMaxSpeed() bool`

HasMaxSpeed returns a boolean if a field has been set.

### GetSpeedFactor

`func (o *VehicleTypeModel) GetSpeedFactor() float64`

GetSpeedFactor returns the SpeedFactor field if non-nil, zero value otherwise.

### GetSpeedFactorOk

`func (o *VehicleTypeModel) GetSpeedFactorOk() (*float64, bool)`

GetSpeedFactorOk returns a tuple with the SpeedFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedFactor

`func (o *VehicleTypeModel) SetSpeedFactor(v float64)`

SetSpeedFactor sets SpeedFactor field to given value.

### HasSpeedFactor

`func (o *VehicleTypeModel) HasSpeedFactor() bool`

HasSpeedFactor returns a boolean if a field has been set.

### GetDurationFactor

`func (o *VehicleTypeModel) GetDurationFactor() float64`

GetDurationFactor returns the DurationFactor field if non-nil, zero value otherwise.

### GetDurationFactorOk

`func (o *VehicleTypeModel) GetDurationFactorOk() (*float64, bool)`

GetDurationFactorOk returns a tuple with the DurationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationFactor

`func (o *VehicleTypeModel) SetDurationFactor(v float64)`

SetDurationFactor sets DurationFactor field to given value.

### HasDurationFactor

`func (o *VehicleTypeModel) HasDurationFactor() bool`

HasDurationFactor returns a boolean if a field has been set.

### GetCostPerMeter

`func (o *VehicleTypeModel) GetCostPerMeter() float64`

GetCostPerMeter returns the CostPerMeter field if non-nil, zero value otherwise.

### GetCostPerMeterOk

`func (o *VehicleTypeModel) GetCostPerMeterOk() (*float64, bool)`

GetCostPerMeterOk returns a tuple with the CostPerMeter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerMeter

`func (o *VehicleTypeModel) SetCostPerMeter(v float64)`

SetCostPerMeter sets CostPerMeter field to given value.

### HasCostPerMeter

`func (o *VehicleTypeModel) HasCostPerMeter() bool`

HasCostPerMeter returns a boolean if a field has been set.

### GetCostPerRoute

`func (o *VehicleTypeModel) GetCostPerRoute() float64`

GetCostPerRoute returns the CostPerRoute field if non-nil, zero value otherwise.

### GetCostPerRouteOk

`func (o *VehicleTypeModel) GetCostPerRouteOk() (*float64, bool)`

GetCostPerRouteOk returns a tuple with the CostPerRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerRoute

`func (o *VehicleTypeModel) SetCostPerRoute(v float64)`

SetCostPerRoute sets CostPerRoute field to given value.

### HasCostPerRoute

`func (o *VehicleTypeModel) HasCostPerRoute() bool`

HasCostPerRoute returns a boolean if a field has been set.

### GetCostPerDrivingMinute

`func (o *VehicleTypeModel) GetCostPerDrivingMinute() float64`

GetCostPerDrivingMinute returns the CostPerDrivingMinute field if non-nil, zero value otherwise.

### GetCostPerDrivingMinuteOk

`func (o *VehicleTypeModel) GetCostPerDrivingMinuteOk() (*float64, bool)`

GetCostPerDrivingMinuteOk returns a tuple with the CostPerDrivingMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerDrivingMinute

`func (o *VehicleTypeModel) SetCostPerDrivingMinute(v float64)`

SetCostPerDrivingMinute sets CostPerDrivingMinute field to given value.

### HasCostPerDrivingMinute

`func (o *VehicleTypeModel) HasCostPerDrivingMinute() bool`

HasCostPerDrivingMinute returns a boolean if a field has been set.

### GetCostPerWaitingMinute

`func (o *VehicleTypeModel) GetCostPerWaitingMinute() float64`

GetCostPerWaitingMinute returns the CostPerWaitingMinute field if non-nil, zero value otherwise.

### GetCostPerWaitingMinuteOk

`func (o *VehicleTypeModel) GetCostPerWaitingMinuteOk() (*float64, bool)`

GetCostPerWaitingMinuteOk returns a tuple with the CostPerWaitingMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerWaitingMinute

`func (o *VehicleTypeModel) SetCostPerWaitingMinute(v float64)`

SetCostPerWaitingMinute sets CostPerWaitingMinute field to given value.

### HasCostPerWaitingMinute

`func (o *VehicleTypeModel) HasCostPerWaitingMinute() bool`

HasCostPerWaitingMinute returns a boolean if a field has been set.

### GetCostPerServiceMinute

`func (o *VehicleTypeModel) GetCostPerServiceMinute() float64`

GetCostPerServiceMinute returns the CostPerServiceMinute field if non-nil, zero value otherwise.

### GetCostPerServiceMinuteOk

`func (o *VehicleTypeModel) GetCostPerServiceMinuteOk() (*float64, bool)`

GetCostPerServiceMinuteOk returns a tuple with the CostPerServiceMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerServiceMinute

`func (o *VehicleTypeModel) SetCostPerServiceMinute(v float64)`

SetCostPerServiceMinute sets CostPerServiceMinute field to given value.

### HasCostPerServiceMinute

`func (o *VehicleTypeModel) HasCostPerServiceMinute() bool`

HasCostPerServiceMinute returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VehicleTypeModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VehicleTypeModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VehicleTypeModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VehicleTypeModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VehicleTypeModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VehicleTypeModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VehicleTypeModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VehicleTypeModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


