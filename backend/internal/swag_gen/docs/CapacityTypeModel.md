# CapacityTypeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique ID | [optional] 
**Name** | Pointer to **string** | name of capacity type | [optional] 
**Properties** | Pointer to **[]string** | extra properties of capacity type | [optional] 
**UomId** | Pointer to **string** | Unit of Measurement ID | [optional] 
**UomName** | Pointer to **string** | Unit of Measurement Name | [optional] 
**Uom** | Pointer to [**UomModel**](UomModel.md) |  | [optional] 
**ValuesUom** | Pointer to [**UomModel**](UomModel.md) |  | [optional] 
**Sys** | Pointer to **bool** | created by system boolean | [optional] 

## Methods

### NewCapacityTypeModel

`func NewCapacityTypeModel() *CapacityTypeModel`

NewCapacityTypeModel instantiates a new CapacityTypeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapacityTypeModelWithDefaults

`func NewCapacityTypeModelWithDefaults() *CapacityTypeModel`

NewCapacityTypeModelWithDefaults instantiates a new CapacityTypeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CapacityTypeModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CapacityTypeModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CapacityTypeModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CapacityTypeModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CapacityTypeModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CapacityTypeModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CapacityTypeModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CapacityTypeModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProperties

`func (o *CapacityTypeModel) GetProperties() []string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CapacityTypeModel) GetPropertiesOk() (*[]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CapacityTypeModel) SetProperties(v []string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CapacityTypeModel) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetUomId

`func (o *CapacityTypeModel) GetUomId() string`

GetUomId returns the UomId field if non-nil, zero value otherwise.

### GetUomIdOk

`func (o *CapacityTypeModel) GetUomIdOk() (*string, bool)`

GetUomIdOk returns a tuple with the UomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUomId

`func (o *CapacityTypeModel) SetUomId(v string)`

SetUomId sets UomId field to given value.

### HasUomId

`func (o *CapacityTypeModel) HasUomId() bool`

HasUomId returns a boolean if a field has been set.

### GetUomName

`func (o *CapacityTypeModel) GetUomName() string`

GetUomName returns the UomName field if non-nil, zero value otherwise.

### GetUomNameOk

`func (o *CapacityTypeModel) GetUomNameOk() (*string, bool)`

GetUomNameOk returns a tuple with the UomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUomName

`func (o *CapacityTypeModel) SetUomName(v string)`

SetUomName sets UomName field to given value.

### HasUomName

`func (o *CapacityTypeModel) HasUomName() bool`

HasUomName returns a boolean if a field has been set.

### GetUom

`func (o *CapacityTypeModel) GetUom() UomModel`

GetUom returns the Uom field if non-nil, zero value otherwise.

### GetUomOk

`func (o *CapacityTypeModel) GetUomOk() (*UomModel, bool)`

GetUomOk returns a tuple with the Uom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUom

`func (o *CapacityTypeModel) SetUom(v UomModel)`

SetUom sets Uom field to given value.

### HasUom

`func (o *CapacityTypeModel) HasUom() bool`

HasUom returns a boolean if a field has been set.

### GetValuesUom

`func (o *CapacityTypeModel) GetValuesUom() UomModel`

GetValuesUom returns the ValuesUom field if non-nil, zero value otherwise.

### GetValuesUomOk

`func (o *CapacityTypeModel) GetValuesUomOk() (*UomModel, bool)`

GetValuesUomOk returns a tuple with the ValuesUom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesUom

`func (o *CapacityTypeModel) SetValuesUom(v UomModel)`

SetValuesUom sets ValuesUom field to given value.

### HasValuesUom

`func (o *CapacityTypeModel) HasValuesUom() bool`

HasValuesUom returns a boolean if a field has been set.

### GetSys

`func (o *CapacityTypeModel) GetSys() bool`

GetSys returns the Sys field if non-nil, zero value otherwise.

### GetSysOk

`func (o *CapacityTypeModel) GetSysOk() (*bool, bool)`

GetSysOk returns a tuple with the Sys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSys

`func (o *CapacityTypeModel) SetSys(v bool)`

SetSys sets Sys field to given value.

### HasSys

`func (o *CapacityTypeModel) HasSys() bool`

HasSys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


