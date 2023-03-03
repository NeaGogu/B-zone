# CapacityModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique ID | [optional] 
**CapacityTypeId** | Pointer to **int32** | id for capacity type | [optional] 
**CapacityTypeName** | Pointer to **string** | name of capacity type | [optional] 
**CapacityType** | Pointer to [**CapacityTypeModel**](CapacityTypeModel.md) |  | [optional] 
**CapacityValue** | Pointer to **float64** | Capacity value | [optional] 
**UnitValues** | Pointer to [**[]UnitValueModel**](UnitValueModel.md) |  | [optional] 
**CapacityValueUomName** | Pointer to **string** | Name of used unit of measurement for values provided in capacity_value | [optional] 
**UnitValuesUomName** | Pointer to **string** | Name of used unit of measurement for values provided in unit_values | [optional] 

## Methods

### NewCapacityModel

`func NewCapacityModel() *CapacityModel`

NewCapacityModel instantiates a new CapacityModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapacityModelWithDefaults

`func NewCapacityModelWithDefaults() *CapacityModel`

NewCapacityModelWithDefaults instantiates a new CapacityModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CapacityModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CapacityModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CapacityModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CapacityModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCapacityTypeId

`func (o *CapacityModel) GetCapacityTypeId() int32`

GetCapacityTypeId returns the CapacityTypeId field if non-nil, zero value otherwise.

### GetCapacityTypeIdOk

`func (o *CapacityModel) GetCapacityTypeIdOk() (*int32, bool)`

GetCapacityTypeIdOk returns a tuple with the CapacityTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityTypeId

`func (o *CapacityModel) SetCapacityTypeId(v int32)`

SetCapacityTypeId sets CapacityTypeId field to given value.

### HasCapacityTypeId

`func (o *CapacityModel) HasCapacityTypeId() bool`

HasCapacityTypeId returns a boolean if a field has been set.

### GetCapacityTypeName

`func (o *CapacityModel) GetCapacityTypeName() string`

GetCapacityTypeName returns the CapacityTypeName field if non-nil, zero value otherwise.

### GetCapacityTypeNameOk

`func (o *CapacityModel) GetCapacityTypeNameOk() (*string, bool)`

GetCapacityTypeNameOk returns a tuple with the CapacityTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityTypeName

`func (o *CapacityModel) SetCapacityTypeName(v string)`

SetCapacityTypeName sets CapacityTypeName field to given value.

### HasCapacityTypeName

`func (o *CapacityModel) HasCapacityTypeName() bool`

HasCapacityTypeName returns a boolean if a field has been set.

### GetCapacityType

`func (o *CapacityModel) GetCapacityType() CapacityTypeModel`

GetCapacityType returns the CapacityType field if non-nil, zero value otherwise.

### GetCapacityTypeOk

`func (o *CapacityModel) GetCapacityTypeOk() (*CapacityTypeModel, bool)`

GetCapacityTypeOk returns a tuple with the CapacityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityType

`func (o *CapacityModel) SetCapacityType(v CapacityTypeModel)`

SetCapacityType sets CapacityType field to given value.

### HasCapacityType

`func (o *CapacityModel) HasCapacityType() bool`

HasCapacityType returns a boolean if a field has been set.

### GetCapacityValue

`func (o *CapacityModel) GetCapacityValue() float64`

GetCapacityValue returns the CapacityValue field if non-nil, zero value otherwise.

### GetCapacityValueOk

`func (o *CapacityModel) GetCapacityValueOk() (*float64, bool)`

GetCapacityValueOk returns a tuple with the CapacityValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityValue

`func (o *CapacityModel) SetCapacityValue(v float64)`

SetCapacityValue sets CapacityValue field to given value.

### HasCapacityValue

`func (o *CapacityModel) HasCapacityValue() bool`

HasCapacityValue returns a boolean if a field has been set.

### GetUnitValues

`func (o *CapacityModel) GetUnitValues() []UnitValueModel`

GetUnitValues returns the UnitValues field if non-nil, zero value otherwise.

### GetUnitValuesOk

`func (o *CapacityModel) GetUnitValuesOk() (*[]UnitValueModel, bool)`

GetUnitValuesOk returns a tuple with the UnitValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitValues

`func (o *CapacityModel) SetUnitValues(v []UnitValueModel)`

SetUnitValues sets UnitValues field to given value.

### HasUnitValues

`func (o *CapacityModel) HasUnitValues() bool`

HasUnitValues returns a boolean if a field has been set.

### GetCapacityValueUomName

`func (o *CapacityModel) GetCapacityValueUomName() string`

GetCapacityValueUomName returns the CapacityValueUomName field if non-nil, zero value otherwise.

### GetCapacityValueUomNameOk

`func (o *CapacityModel) GetCapacityValueUomNameOk() (*string, bool)`

GetCapacityValueUomNameOk returns a tuple with the CapacityValueUomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityValueUomName

`func (o *CapacityModel) SetCapacityValueUomName(v string)`

SetCapacityValueUomName sets CapacityValueUomName field to given value.

### HasCapacityValueUomName

`func (o *CapacityModel) HasCapacityValueUomName() bool`

HasCapacityValueUomName returns a boolean if a field has been set.

### GetUnitValuesUomName

`func (o *CapacityModel) GetUnitValuesUomName() string`

GetUnitValuesUomName returns the UnitValuesUomName field if non-nil, zero value otherwise.

### GetUnitValuesUomNameOk

`func (o *CapacityModel) GetUnitValuesUomNameOk() (*string, bool)`

GetUnitValuesUomNameOk returns a tuple with the UnitValuesUomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitValuesUomName

`func (o *CapacityModel) SetUnitValuesUomName(v string)`

SetUnitValuesUomName sets UnitValuesUomName field to given value.

### HasUnitValuesUomName

`func (o *CapacityModel) HasUnitValuesUomName() bool`

HasUnitValuesUomName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


