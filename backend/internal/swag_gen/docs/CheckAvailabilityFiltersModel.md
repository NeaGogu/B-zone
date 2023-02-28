# CheckAvailabilityFiltersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RouteId** | Pointer to **[]int32** | Route id | [optional] 
**DateFrom** | Pointer to **string** |  | [optional] 
**DateTo** | Pointer to **string** |  | [optional] 
**MaxNrOfDaysWithAvailability** | Pointer to **int32** | Availability check will continue to analyze days until there is no availability in the system anymore or the number of days with available time slots has reached the max_nr_of_days_with_availability | [optional] 

## Methods

### NewCheckAvailabilityFiltersModel

`func NewCheckAvailabilityFiltersModel() *CheckAvailabilityFiltersModel`

NewCheckAvailabilityFiltersModel instantiates a new CheckAvailabilityFiltersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckAvailabilityFiltersModelWithDefaults

`func NewCheckAvailabilityFiltersModelWithDefaults() *CheckAvailabilityFiltersModel`

NewCheckAvailabilityFiltersModelWithDefaults instantiates a new CheckAvailabilityFiltersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRouteId

`func (o *CheckAvailabilityFiltersModel) GetRouteId() []int32`

GetRouteId returns the RouteId field if non-nil, zero value otherwise.

### GetRouteIdOk

`func (o *CheckAvailabilityFiltersModel) GetRouteIdOk() (*[]int32, bool)`

GetRouteIdOk returns a tuple with the RouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteId

`func (o *CheckAvailabilityFiltersModel) SetRouteId(v []int32)`

SetRouteId sets RouteId field to given value.

### HasRouteId

`func (o *CheckAvailabilityFiltersModel) HasRouteId() bool`

HasRouteId returns a boolean if a field has been set.

### GetDateFrom

`func (o *CheckAvailabilityFiltersModel) GetDateFrom() string`

GetDateFrom returns the DateFrom field if non-nil, zero value otherwise.

### GetDateFromOk

`func (o *CheckAvailabilityFiltersModel) GetDateFromOk() (*string, bool)`

GetDateFromOk returns a tuple with the DateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFrom

`func (o *CheckAvailabilityFiltersModel) SetDateFrom(v string)`

SetDateFrom sets DateFrom field to given value.

### HasDateFrom

`func (o *CheckAvailabilityFiltersModel) HasDateFrom() bool`

HasDateFrom returns a boolean if a field has been set.

### GetDateTo

`func (o *CheckAvailabilityFiltersModel) GetDateTo() string`

GetDateTo returns the DateTo field if non-nil, zero value otherwise.

### GetDateToOk

`func (o *CheckAvailabilityFiltersModel) GetDateToOk() (*string, bool)`

GetDateToOk returns a tuple with the DateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTo

`func (o *CheckAvailabilityFiltersModel) SetDateTo(v string)`

SetDateTo sets DateTo field to given value.

### HasDateTo

`func (o *CheckAvailabilityFiltersModel) HasDateTo() bool`

HasDateTo returns a boolean if a field has been set.

### GetMaxNrOfDaysWithAvailability

`func (o *CheckAvailabilityFiltersModel) GetMaxNrOfDaysWithAvailability() int32`

GetMaxNrOfDaysWithAvailability returns the MaxNrOfDaysWithAvailability field if non-nil, zero value otherwise.

### GetMaxNrOfDaysWithAvailabilityOk

`func (o *CheckAvailabilityFiltersModel) GetMaxNrOfDaysWithAvailabilityOk() (*int32, bool)`

GetMaxNrOfDaysWithAvailabilityOk returns a tuple with the MaxNrOfDaysWithAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNrOfDaysWithAvailability

`func (o *CheckAvailabilityFiltersModel) SetMaxNrOfDaysWithAvailability(v int32)`

SetMaxNrOfDaysWithAvailability sets MaxNrOfDaysWithAvailability field to given value.

### HasMaxNrOfDaysWithAvailability

`func (o *CheckAvailabilityFiltersModel) HasMaxNrOfDaysWithAvailability() bool`

HasMaxNrOfDaysWithAvailability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


