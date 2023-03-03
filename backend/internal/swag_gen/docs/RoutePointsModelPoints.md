# RoutePointsModelPoints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Type as defined by GeoJSON standard | [optional] 
**Coordinates** | Pointer to **[][]float32** | Array of longitude/latitude pairs | [optional] 

## Methods

### NewRoutePointsModelPoints

`func NewRoutePointsModelPoints() *RoutePointsModelPoints`

NewRoutePointsModelPoints instantiates a new RoutePointsModelPoints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutePointsModelPointsWithDefaults

`func NewRoutePointsModelPointsWithDefaults() *RoutePointsModelPoints`

NewRoutePointsModelPointsWithDefaults instantiates a new RoutePointsModelPoints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RoutePointsModelPoints) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoutePointsModelPoints) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoutePointsModelPoints) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RoutePointsModelPoints) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCoordinates

`func (o *RoutePointsModelPoints) GetCoordinates() [][]float32`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *RoutePointsModelPoints) GetCoordinatesOk() (*[][]float32, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *RoutePointsModelPoints) SetCoordinates(v [][]float32)`

SetCoordinates sets Coordinates field to given value.

### HasCoordinates

`func (o *RoutePointsModelPoints) HasCoordinates() bool`

HasCoordinates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


