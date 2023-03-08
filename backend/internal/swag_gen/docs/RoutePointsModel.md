# RoutePointsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RouteId** | Pointer to **int64** | RouteId | [optional] [readonly] 
**Points** | Pointer to [**RoutePointsModelPoints**](RoutePointsModel_points.md) |  | [optional] 
**Failed** | Pointer to **string** | Reason why retrieving this route point failed | [optional] [readonly] 

## Methods

### NewRoutePointsModel

`func NewRoutePointsModel() *RoutePointsModel`

NewRoutePointsModel instantiates a new RoutePointsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutePointsModelWithDefaults

`func NewRoutePointsModelWithDefaults() *RoutePointsModel`

NewRoutePointsModelWithDefaults instantiates a new RoutePointsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRouteId

`func (o *RoutePointsModel) GetRouteId() int64`

GetRouteId returns the RouteId field if non-nil, zero value otherwise.

### GetRouteIdOk

`func (o *RoutePointsModel) GetRouteIdOk() (*int64, bool)`

GetRouteIdOk returns a tuple with the RouteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteId

`func (o *RoutePointsModel) SetRouteId(v int64)`

SetRouteId sets RouteId field to given value.

### HasRouteId

`func (o *RoutePointsModel) HasRouteId() bool`

HasRouteId returns a boolean if a field has been set.

### GetPoints

`func (o *RoutePointsModel) GetPoints() RoutePointsModelPoints`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *RoutePointsModel) GetPointsOk() (*RoutePointsModelPoints, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *RoutePointsModel) SetPoints(v RoutePointsModelPoints)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *RoutePointsModel) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetFailed

`func (o *RoutePointsModel) GetFailed() string`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *RoutePointsModel) GetFailedOk() (*string, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *RoutePointsModel) SetFailed(v string)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *RoutePointsModel) HasFailed() bool`

HasFailed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


