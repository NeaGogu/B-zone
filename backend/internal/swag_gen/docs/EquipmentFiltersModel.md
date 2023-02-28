# EquipmentFiltersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Equipment ID | [optional] 
**EquipmentTypeId** | Pointer to **[]int32** | Equipment Type ID&#39;s | [optional] 

## Methods

### NewEquipmentFiltersModel

`func NewEquipmentFiltersModel() *EquipmentFiltersModel`

NewEquipmentFiltersModel instantiates a new EquipmentFiltersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentFiltersModelWithDefaults

`func NewEquipmentFiltersModelWithDefaults() *EquipmentFiltersModel`

NewEquipmentFiltersModelWithDefaults instantiates a new EquipmentFiltersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EquipmentFiltersModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EquipmentFiltersModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EquipmentFiltersModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EquipmentFiltersModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEquipmentTypeId

`func (o *EquipmentFiltersModel) GetEquipmentTypeId() []int32`

GetEquipmentTypeId returns the EquipmentTypeId field if non-nil, zero value otherwise.

### GetEquipmentTypeIdOk

`func (o *EquipmentFiltersModel) GetEquipmentTypeIdOk() (*[]int32, bool)`

GetEquipmentTypeIdOk returns a tuple with the EquipmentTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentTypeId

`func (o *EquipmentFiltersModel) SetEquipmentTypeId(v []int32)`

SetEquipmentTypeId sets EquipmentTypeId field to given value.

### HasEquipmentTypeId

`func (o *EquipmentFiltersModel) HasEquipmentTypeId() bool`

HasEquipmentTypeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


