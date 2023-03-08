# UomModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique ID | [optional] 
**Name** | Pointer to **string** | name of unit | [optional] 
**UomGroup** | Pointer to [**UomGroupModel**](UomGroupModel.md) |  | [optional] 

## Methods

### NewUomModel

`func NewUomModel() *UomModel`

NewUomModel instantiates a new UomModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUomModelWithDefaults

`func NewUomModelWithDefaults() *UomModel`

NewUomModelWithDefaults instantiates a new UomModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UomModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UomModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UomModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UomModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UomModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UomModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UomModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UomModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUomGroup

`func (o *UomModel) GetUomGroup() UomGroupModel`

GetUomGroup returns the UomGroup field if non-nil, zero value otherwise.

### GetUomGroupOk

`func (o *UomModel) GetUomGroupOk() (*UomGroupModel, bool)`

GetUomGroupOk returns a tuple with the UomGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUomGroup

`func (o *UomModel) SetUomGroup(v UomGroupModel)`

SetUomGroup sets UomGroup field to given value.

### HasUomGroup

`func (o *UomModel) HasUomGroup() bool`

HasUomGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


