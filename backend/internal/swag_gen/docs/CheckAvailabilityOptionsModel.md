# CheckAvailabilityOptionsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeImpact** | Pointer to **bool** |  | [optional] 
**Synchronous** | Pointer to **bool** |  | [optional] 
**IncludeCapacities** | Pointer to **bool** |  | [optional] 
**IncludeProposedPlanTimes** | Pointer to **bool** |  | [optional] 
**IncludeProposedDriver** | Pointer to **bool** |  | [optional] 
**ActivityTimesMarginsLeading** | Pointer to **bool** | default: true | [optional] 
**ApplyCutOffTimes** | Pointer to **bool** | default: true | [optional] 
**ConsiderHistoricTrafficInfo** | Pointer to **bool** |  | [optional] 

## Methods

### NewCheckAvailabilityOptionsModel

`func NewCheckAvailabilityOptionsModel() *CheckAvailabilityOptionsModel`

NewCheckAvailabilityOptionsModel instantiates a new CheckAvailabilityOptionsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckAvailabilityOptionsModelWithDefaults

`func NewCheckAvailabilityOptionsModelWithDefaults() *CheckAvailabilityOptionsModel`

NewCheckAvailabilityOptionsModelWithDefaults instantiates a new CheckAvailabilityOptionsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeImpact

`func (o *CheckAvailabilityOptionsModel) GetIncludeImpact() bool`

GetIncludeImpact returns the IncludeImpact field if non-nil, zero value otherwise.

### GetIncludeImpactOk

`func (o *CheckAvailabilityOptionsModel) GetIncludeImpactOk() (*bool, bool)`

GetIncludeImpactOk returns a tuple with the IncludeImpact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeImpact

`func (o *CheckAvailabilityOptionsModel) SetIncludeImpact(v bool)`

SetIncludeImpact sets IncludeImpact field to given value.

### HasIncludeImpact

`func (o *CheckAvailabilityOptionsModel) HasIncludeImpact() bool`

HasIncludeImpact returns a boolean if a field has been set.

### GetSynchronous

`func (o *CheckAvailabilityOptionsModel) GetSynchronous() bool`

GetSynchronous returns the Synchronous field if non-nil, zero value otherwise.

### GetSynchronousOk

`func (o *CheckAvailabilityOptionsModel) GetSynchronousOk() (*bool, bool)`

GetSynchronousOk returns a tuple with the Synchronous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronous

`func (o *CheckAvailabilityOptionsModel) SetSynchronous(v bool)`

SetSynchronous sets Synchronous field to given value.

### HasSynchronous

`func (o *CheckAvailabilityOptionsModel) HasSynchronous() bool`

HasSynchronous returns a boolean if a field has been set.

### GetIncludeCapacities

`func (o *CheckAvailabilityOptionsModel) GetIncludeCapacities() bool`

GetIncludeCapacities returns the IncludeCapacities field if non-nil, zero value otherwise.

### GetIncludeCapacitiesOk

`func (o *CheckAvailabilityOptionsModel) GetIncludeCapacitiesOk() (*bool, bool)`

GetIncludeCapacitiesOk returns a tuple with the IncludeCapacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCapacities

`func (o *CheckAvailabilityOptionsModel) SetIncludeCapacities(v bool)`

SetIncludeCapacities sets IncludeCapacities field to given value.

### HasIncludeCapacities

`func (o *CheckAvailabilityOptionsModel) HasIncludeCapacities() bool`

HasIncludeCapacities returns a boolean if a field has been set.

### GetIncludeProposedPlanTimes

`func (o *CheckAvailabilityOptionsModel) GetIncludeProposedPlanTimes() bool`

GetIncludeProposedPlanTimes returns the IncludeProposedPlanTimes field if non-nil, zero value otherwise.

### GetIncludeProposedPlanTimesOk

`func (o *CheckAvailabilityOptionsModel) GetIncludeProposedPlanTimesOk() (*bool, bool)`

GetIncludeProposedPlanTimesOk returns a tuple with the IncludeProposedPlanTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeProposedPlanTimes

`func (o *CheckAvailabilityOptionsModel) SetIncludeProposedPlanTimes(v bool)`

SetIncludeProposedPlanTimes sets IncludeProposedPlanTimes field to given value.

### HasIncludeProposedPlanTimes

`func (o *CheckAvailabilityOptionsModel) HasIncludeProposedPlanTimes() bool`

HasIncludeProposedPlanTimes returns a boolean if a field has been set.

### GetIncludeProposedDriver

`func (o *CheckAvailabilityOptionsModel) GetIncludeProposedDriver() bool`

GetIncludeProposedDriver returns the IncludeProposedDriver field if non-nil, zero value otherwise.

### GetIncludeProposedDriverOk

`func (o *CheckAvailabilityOptionsModel) GetIncludeProposedDriverOk() (*bool, bool)`

GetIncludeProposedDriverOk returns a tuple with the IncludeProposedDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeProposedDriver

`func (o *CheckAvailabilityOptionsModel) SetIncludeProposedDriver(v bool)`

SetIncludeProposedDriver sets IncludeProposedDriver field to given value.

### HasIncludeProposedDriver

`func (o *CheckAvailabilityOptionsModel) HasIncludeProposedDriver() bool`

HasIncludeProposedDriver returns a boolean if a field has been set.

### GetActivityTimesMarginsLeading

`func (o *CheckAvailabilityOptionsModel) GetActivityTimesMarginsLeading() bool`

GetActivityTimesMarginsLeading returns the ActivityTimesMarginsLeading field if non-nil, zero value otherwise.

### GetActivityTimesMarginsLeadingOk

`func (o *CheckAvailabilityOptionsModel) GetActivityTimesMarginsLeadingOk() (*bool, bool)`

GetActivityTimesMarginsLeadingOk returns a tuple with the ActivityTimesMarginsLeading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityTimesMarginsLeading

`func (o *CheckAvailabilityOptionsModel) SetActivityTimesMarginsLeading(v bool)`

SetActivityTimesMarginsLeading sets ActivityTimesMarginsLeading field to given value.

### HasActivityTimesMarginsLeading

`func (o *CheckAvailabilityOptionsModel) HasActivityTimesMarginsLeading() bool`

HasActivityTimesMarginsLeading returns a boolean if a field has been set.

### GetApplyCutOffTimes

`func (o *CheckAvailabilityOptionsModel) GetApplyCutOffTimes() bool`

GetApplyCutOffTimes returns the ApplyCutOffTimes field if non-nil, zero value otherwise.

### GetApplyCutOffTimesOk

`func (o *CheckAvailabilityOptionsModel) GetApplyCutOffTimesOk() (*bool, bool)`

GetApplyCutOffTimesOk returns a tuple with the ApplyCutOffTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyCutOffTimes

`func (o *CheckAvailabilityOptionsModel) SetApplyCutOffTimes(v bool)`

SetApplyCutOffTimes sets ApplyCutOffTimes field to given value.

### HasApplyCutOffTimes

`func (o *CheckAvailabilityOptionsModel) HasApplyCutOffTimes() bool`

HasApplyCutOffTimes returns a boolean if a field has been set.

### GetConsiderHistoricTrafficInfo

`func (o *CheckAvailabilityOptionsModel) GetConsiderHistoricTrafficInfo() bool`

GetConsiderHistoricTrafficInfo returns the ConsiderHistoricTrafficInfo field if non-nil, zero value otherwise.

### GetConsiderHistoricTrafficInfoOk

`func (o *CheckAvailabilityOptionsModel) GetConsiderHistoricTrafficInfoOk() (*bool, bool)`

GetConsiderHistoricTrafficInfoOk returns a tuple with the ConsiderHistoricTrafficInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsiderHistoricTrafficInfo

`func (o *CheckAvailabilityOptionsModel) SetConsiderHistoricTrafficInfo(v bool)`

SetConsiderHistoricTrafficInfo sets ConsiderHistoricTrafficInfo field to given value.

### HasConsiderHistoricTrafficInfo

`func (o *CheckAvailabilityOptionsModel) HasConsiderHistoricTrafficInfo() bool`

HasConsiderHistoricTrafficInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


