# AddActivitiesToRouteArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RouteId** | Pointer to **int64** | Unique ID of Route | [optional] 
**Activities** | Pointer to [**[]ActivityForToRouteModel**](ActivityForToRouteModel.md) |  | [optional] 
**Options** | Pointer to **map[string]bool** |  | [optional] 
**Filters** | Pointer to [**AddActivitiesToRouteFiltersModel**](AddActivitiesToRouteFiltersModel.md) |  | [optional] 

## Methods

### NewAddActivitiesToRouteArguments

`func NewAddActivitiesToRouteArguments() *AddActivitiesToRouteArguments`

NewAddActivitiesToRouteArguments instantiates a new AddActivitiesToRouteArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddActivitiesToRouteArgumentsWithDefaults

`func NewAddActivitiesToRouteArgumentsWithDefaults() *AddActivitiesToRouteArguments`

NewAddActivitiesToRouteArgumentsWithDefaults instantiates a new AddActivitiesToRouteArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRouteId

`func (o *AddActivitiesToRouteArguments) GetRouteId() int64`

GetRouteId returns the RouteId field if non-nil, zero value otherwise.

### GetRouteIdOk

`func (o *AddActivitiesToRouteArguments) GetRouteIdOk() (*int64, bool)`

GetRouteIdOk returns a tuple with the RouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteId

`func (o *AddActivitiesToRouteArguments) SetRouteId(v int64)`

SetRouteId sets RouteId field to given value.

### HasRouteId

`func (o *AddActivitiesToRouteArguments) HasRouteId() bool`

HasRouteId returns a boolean if a field has been set.

### GetActivities

`func (o *AddActivitiesToRouteArguments) GetActivities() []ActivityForToRouteModel`

GetActivities returns the Activities field if non-nil, zero value otherwise.

### GetActivitiesOk

`func (o *AddActivitiesToRouteArguments) GetActivitiesOk() (*[]ActivityForToRouteModel, bool)`

GetActivitiesOk returns a tuple with the Activities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivities

`func (o *AddActivitiesToRouteArguments) SetActivities(v []ActivityForToRouteModel)`

SetActivities sets Activities field to given value.

### HasActivities

`func (o *AddActivitiesToRouteArguments) HasActivities() bool`

HasActivities returns a boolean if a field has been set.

### GetOptions

`func (o *AddActivitiesToRouteArguments) GetOptions() map[string]bool`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AddActivitiesToRouteArguments) GetOptionsOk() (*map[string]bool, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AddActivitiesToRouteArguments) SetOptions(v map[string]bool)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AddActivitiesToRouteArguments) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetFilters

`func (o *AddActivitiesToRouteArguments) GetFilters() AddActivitiesToRouteFiltersModel`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *AddActivitiesToRouteArguments) GetFiltersOk() (*AddActivitiesToRouteFiltersModel, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *AddActivitiesToRouteArguments) SetFilters(v AddActivitiesToRouteFiltersModel)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *AddActivitiesToRouteArguments) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


