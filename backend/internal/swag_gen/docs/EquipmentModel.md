# EquipmentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] [readonly] 
**EquipmentTypeId** | Pointer to **int32** | Equipment Type ID | [optional] 
**EquipmentTypeName** | Pointer to **string** | Equipment Type Name | [optional] 
**Info** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RegistrationNr** | Pointer to **string** |  | [optional] 
**MaxSpeed** | Pointer to **int32** |  | [optional] 
**Tags** | Pointer to [**[]TagModel**](TagModel.md) |  | [optional] 
**AppliedCapacities** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**Capacities** | Pointer to [**[]CapacityModel**](CapacityModel.md) |  | [optional] 
**Links** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**MetaData** | Pointer to [**[]MetaDataModel**](MetaDataModel.md) |  | [optional] 

## Methods

### NewEquipmentModel

`func NewEquipmentModel() *EquipmentModel`

NewEquipmentModel instantiates a new EquipmentModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEquipmentModelWithDefaults

`func NewEquipmentModelWithDefaults() *EquipmentModel`

NewEquipmentModelWithDefaults instantiates a new EquipmentModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EquipmentModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EquipmentModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EquipmentModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *EquipmentModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEquipmentTypeId

`func (o *EquipmentModel) GetEquipmentTypeId() int32`

GetEquipmentTypeId returns the EquipmentTypeId field if non-nil, zero value otherwise.

### GetEquipmentTypeIdOk

`func (o *EquipmentModel) GetEquipmentTypeIdOk() (*int32, bool)`

GetEquipmentTypeIdOk returns a tuple with the EquipmentTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentTypeId

`func (o *EquipmentModel) SetEquipmentTypeId(v int32)`

SetEquipmentTypeId sets EquipmentTypeId field to given value.

### HasEquipmentTypeId

`func (o *EquipmentModel) HasEquipmentTypeId() bool`

HasEquipmentTypeId returns a boolean if a field has been set.

### GetEquipmentTypeName

`func (o *EquipmentModel) GetEquipmentTypeName() string`

GetEquipmentTypeName returns the EquipmentTypeName field if non-nil, zero value otherwise.

### GetEquipmentTypeNameOk

`func (o *EquipmentModel) GetEquipmentTypeNameOk() (*string, bool)`

GetEquipmentTypeNameOk returns a tuple with the EquipmentTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentTypeName

`func (o *EquipmentModel) SetEquipmentTypeName(v string)`

SetEquipmentTypeName sets EquipmentTypeName field to given value.

### HasEquipmentTypeName

`func (o *EquipmentModel) HasEquipmentTypeName() bool`

HasEquipmentTypeName returns a boolean if a field has been set.

### GetInfo

`func (o *EquipmentModel) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *EquipmentModel) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *EquipmentModel) SetInfo(v string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *EquipmentModel) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetName

`func (o *EquipmentModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EquipmentModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EquipmentModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EquipmentModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegistrationNr

`func (o *EquipmentModel) GetRegistrationNr() string`

GetRegistrationNr returns the RegistrationNr field if non-nil, zero value otherwise.

### GetRegistrationNrOk

`func (o *EquipmentModel) GetRegistrationNrOk() (*string, bool)`

GetRegistrationNrOk returns a tuple with the RegistrationNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationNr

`func (o *EquipmentModel) SetRegistrationNr(v string)`

SetRegistrationNr sets RegistrationNr field to given value.

### HasRegistrationNr

`func (o *EquipmentModel) HasRegistrationNr() bool`

HasRegistrationNr returns a boolean if a field has been set.

### GetMaxSpeed

`func (o *EquipmentModel) GetMaxSpeed() int32`

GetMaxSpeed returns the MaxSpeed field if non-nil, zero value otherwise.

### GetMaxSpeedOk

`func (o *EquipmentModel) GetMaxSpeedOk() (*int32, bool)`

GetMaxSpeedOk returns a tuple with the MaxSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpeed

`func (o *EquipmentModel) SetMaxSpeed(v int32)`

SetMaxSpeed sets MaxSpeed field to given value.

### HasMaxSpeed

`func (o *EquipmentModel) HasMaxSpeed() bool`

HasMaxSpeed returns a boolean if a field has been set.

### GetTags

`func (o *EquipmentModel) GetTags() []TagModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EquipmentModel) GetTagsOk() (*[]TagModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EquipmentModel) SetTags(v []TagModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EquipmentModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAppliedCapacities

`func (o *EquipmentModel) GetAppliedCapacities() map[string]interface{}`

GetAppliedCapacities returns the AppliedCapacities field if non-nil, zero value otherwise.

### GetAppliedCapacitiesOk

`func (o *EquipmentModel) GetAppliedCapacitiesOk() (*map[string]interface{}, bool)`

GetAppliedCapacitiesOk returns a tuple with the AppliedCapacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedCapacities

`func (o *EquipmentModel) SetAppliedCapacities(v map[string]interface{})`

SetAppliedCapacities sets AppliedCapacities field to given value.

### HasAppliedCapacities

`func (o *EquipmentModel) HasAppliedCapacities() bool`

HasAppliedCapacities returns a boolean if a field has been set.

### GetCapacities

`func (o *EquipmentModel) GetCapacities() []CapacityModel`

GetCapacities returns the Capacities field if non-nil, zero value otherwise.

### GetCapacitiesOk

`func (o *EquipmentModel) GetCapacitiesOk() (*[]CapacityModel, bool)`

GetCapacitiesOk returns a tuple with the Capacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacities

`func (o *EquipmentModel) SetCapacities(v []CapacityModel)`

SetCapacities sets Capacities field to given value.

### HasCapacities

`func (o *EquipmentModel) HasCapacities() bool`

HasCapacities returns a boolean if a field has been set.

### GetLinks

`func (o *EquipmentModel) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EquipmentModel) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EquipmentModel) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EquipmentModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetaData

`func (o *EquipmentModel) GetMetaData() []MetaDataModel`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *EquipmentModel) GetMetaDataOk() (*[]MetaDataModel, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *EquipmentModel) SetMetaData(v []MetaDataModel)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *EquipmentModel) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


