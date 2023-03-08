# VehicleFiltersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **[]int32** | Vehicle id&#39;s | [optional] 
**VehicleTypeId** | Pointer to **[]int32** | Vehicle type id&#39;s | [optional] 
**Links** | Pointer to **[]map[string]interface{}** | Vehicle Link ids | [optional] 
**UpdatedAtSince** | Pointer to **time.Time** | Show updated since | [optional] 
**UpdatedAtTill** | Pointer to **time.Time** | Show updated till | [optional] 

## Methods

### NewVehicleFiltersModel

`func NewVehicleFiltersModel() *VehicleFiltersModel`

NewVehicleFiltersModel instantiates a new VehicleFiltersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleFiltersModelWithDefaults

`func NewVehicleFiltersModelWithDefaults() *VehicleFiltersModel`

NewVehicleFiltersModelWithDefaults instantiates a new VehicleFiltersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VehicleFiltersModel) GetId() []int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VehicleFiltersModel) GetIdOk() (*[]int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VehicleFiltersModel) SetId(v []int32)`

SetId sets Id field to given value.

### HasId

`func (o *VehicleFiltersModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVehicleTypeId

`func (o *VehicleFiltersModel) GetVehicleTypeId() []int32`

GetVehicleTypeId returns the VehicleTypeId field if non-nil, zero value otherwise.

### GetVehicleTypeIdOk

`func (o *VehicleFiltersModel) GetVehicleTypeIdOk() (*[]int32, bool)`

GetVehicleTypeIdOk returns a tuple with the VehicleTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleTypeId

`func (o *VehicleFiltersModel) SetVehicleTypeId(v []int32)`

SetVehicleTypeId sets VehicleTypeId field to given value.

### HasVehicleTypeId

`func (o *VehicleFiltersModel) HasVehicleTypeId() bool`

HasVehicleTypeId returns a boolean if a field has been set.

### GetLinks

`func (o *VehicleFiltersModel) GetLinks() []map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VehicleFiltersModel) GetLinksOk() (*[]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VehicleFiltersModel) SetLinks(v []map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *VehicleFiltersModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetUpdatedAtSince

`func (o *VehicleFiltersModel) GetUpdatedAtSince() time.Time`

GetUpdatedAtSince returns the UpdatedAtSince field if non-nil, zero value otherwise.

### GetUpdatedAtSinceOk

`func (o *VehicleFiltersModel) GetUpdatedAtSinceOk() (*time.Time, bool)`

GetUpdatedAtSinceOk returns a tuple with the UpdatedAtSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtSince

`func (o *VehicleFiltersModel) SetUpdatedAtSince(v time.Time)`

SetUpdatedAtSince sets UpdatedAtSince field to given value.

### HasUpdatedAtSince

`func (o *VehicleFiltersModel) HasUpdatedAtSince() bool`

HasUpdatedAtSince returns a boolean if a field has been set.

### GetUpdatedAtTill

`func (o *VehicleFiltersModel) GetUpdatedAtTill() time.Time`

GetUpdatedAtTill returns the UpdatedAtTill field if non-nil, zero value otherwise.

### GetUpdatedAtTillOk

`func (o *VehicleFiltersModel) GetUpdatedAtTillOk() (*time.Time, bool)`

GetUpdatedAtTillOk returns a tuple with the UpdatedAtTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtTill

`func (o *VehicleFiltersModel) SetUpdatedAtTill(v time.Time)`

SetUpdatedAtTill sets UpdatedAtTill field to given value.

### HasUpdatedAtTill

`func (o *VehicleFiltersModel) HasUpdatedAtTill() bool`

HasUpdatedAtTill returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


