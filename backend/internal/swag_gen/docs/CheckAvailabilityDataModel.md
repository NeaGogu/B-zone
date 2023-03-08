# CheckAvailabilityDataModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activity** | [**ActivityModel**](ActivityModel.md) |  | 
**RouteFilters** | Pointer to **map[string]interface{}** |  | [optional] 
**Token** | Pointer to **string** | unique per api request | [optional] 

## Methods

### NewCheckAvailabilityDataModel

`func NewCheckAvailabilityDataModel(activity ActivityModel, ) *CheckAvailabilityDataModel`

NewCheckAvailabilityDataModel instantiates a new CheckAvailabilityDataModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckAvailabilityDataModelWithDefaults

`func NewCheckAvailabilityDataModelWithDefaults() *CheckAvailabilityDataModel`

NewCheckAvailabilityDataModelWithDefaults instantiates a new CheckAvailabilityDataModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivity

`func (o *CheckAvailabilityDataModel) GetActivity() ActivityModel`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *CheckAvailabilityDataModel) GetActivityOk() (*ActivityModel, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *CheckAvailabilityDataModel) SetActivity(v ActivityModel)`

SetActivity sets Activity field to given value.


### GetRouteFilters

`func (o *CheckAvailabilityDataModel) GetRouteFilters() map[string]interface{}`

GetRouteFilters returns the RouteFilters field if non-nil, zero value otherwise.

### GetRouteFiltersOk

`func (o *CheckAvailabilityDataModel) GetRouteFiltersOk() (*map[string]interface{}, bool)`

GetRouteFiltersOk returns a tuple with the RouteFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteFilters

`func (o *CheckAvailabilityDataModel) SetRouteFilters(v map[string]interface{})`

SetRouteFilters sets RouteFilters field to given value.

### HasRouteFilters

`func (o *CheckAvailabilityDataModel) HasRouteFilters() bool`

HasRouteFilters returns a boolean if a field has been set.

### GetToken

`func (o *CheckAvailabilityDataModel) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CheckAvailabilityDataModel) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CheckAvailabilityDataModel) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CheckAvailabilityDataModel) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


