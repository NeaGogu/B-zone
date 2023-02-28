# VehicleOptionsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeTags** | Pointer to **bool** |  | [optional] 
**IncludeLinks** | Pointer to **bool** |  | [optional] 
**IncludeMetaData** | Pointer to **bool** |  | [optional] 
**IncludeUpdatedByName** | Pointer to **bool** |  | [optional] 
**IncludeVehicleMetaData** | Pointer to **bool** | Deprecated! use include_meta_data | [optional] 
**IncludeVehicleLinks** | Pointer to **bool** | Deprecated! use include_links | [optional] 
**IncludeVehicleTags** | Pointer to **bool** | Deprecated! use include_tags | [optional] 
**IncludeTagTypeName** | Pointer to **bool** | Deprecated! use include_tags | [optional] 

## Methods

### NewVehicleOptionsModel

`func NewVehicleOptionsModel() *VehicleOptionsModel`

NewVehicleOptionsModel instantiates a new VehicleOptionsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleOptionsModelWithDefaults

`func NewVehicleOptionsModelWithDefaults() *VehicleOptionsModel`

NewVehicleOptionsModelWithDefaults instantiates a new VehicleOptionsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeTags

`func (o *VehicleOptionsModel) GetIncludeTags() bool`

GetIncludeTags returns the IncludeTags field if non-nil, zero value otherwise.

### GetIncludeTagsOk

`func (o *VehicleOptionsModel) GetIncludeTagsOk() (*bool, bool)`

GetIncludeTagsOk returns a tuple with the IncludeTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTags

`func (o *VehicleOptionsModel) SetIncludeTags(v bool)`

SetIncludeTags sets IncludeTags field to given value.

### HasIncludeTags

`func (o *VehicleOptionsModel) HasIncludeTags() bool`

HasIncludeTags returns a boolean if a field has been set.

### GetIncludeLinks

`func (o *VehicleOptionsModel) GetIncludeLinks() bool`

GetIncludeLinks returns the IncludeLinks field if non-nil, zero value otherwise.

### GetIncludeLinksOk

`func (o *VehicleOptionsModel) GetIncludeLinksOk() (*bool, bool)`

GetIncludeLinksOk returns a tuple with the IncludeLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLinks

`func (o *VehicleOptionsModel) SetIncludeLinks(v bool)`

SetIncludeLinks sets IncludeLinks field to given value.

### HasIncludeLinks

`func (o *VehicleOptionsModel) HasIncludeLinks() bool`

HasIncludeLinks returns a boolean if a field has been set.

### GetIncludeMetaData

`func (o *VehicleOptionsModel) GetIncludeMetaData() bool`

GetIncludeMetaData returns the IncludeMetaData field if non-nil, zero value otherwise.

### GetIncludeMetaDataOk

`func (o *VehicleOptionsModel) GetIncludeMetaDataOk() (*bool, bool)`

GetIncludeMetaDataOk returns a tuple with the IncludeMetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMetaData

`func (o *VehicleOptionsModel) SetIncludeMetaData(v bool)`

SetIncludeMetaData sets IncludeMetaData field to given value.

### HasIncludeMetaData

`func (o *VehicleOptionsModel) HasIncludeMetaData() bool`

HasIncludeMetaData returns a boolean if a field has been set.

### GetIncludeUpdatedByName

`func (o *VehicleOptionsModel) GetIncludeUpdatedByName() bool`

GetIncludeUpdatedByName returns the IncludeUpdatedByName field if non-nil, zero value otherwise.

### GetIncludeUpdatedByNameOk

`func (o *VehicleOptionsModel) GetIncludeUpdatedByNameOk() (*bool, bool)`

GetIncludeUpdatedByNameOk returns a tuple with the IncludeUpdatedByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeUpdatedByName

`func (o *VehicleOptionsModel) SetIncludeUpdatedByName(v bool)`

SetIncludeUpdatedByName sets IncludeUpdatedByName field to given value.

### HasIncludeUpdatedByName

`func (o *VehicleOptionsModel) HasIncludeUpdatedByName() bool`

HasIncludeUpdatedByName returns a boolean if a field has been set.

### GetIncludeVehicleMetaData

`func (o *VehicleOptionsModel) GetIncludeVehicleMetaData() bool`

GetIncludeVehicleMetaData returns the IncludeVehicleMetaData field if non-nil, zero value otherwise.

### GetIncludeVehicleMetaDataOk

`func (o *VehicleOptionsModel) GetIncludeVehicleMetaDataOk() (*bool, bool)`

GetIncludeVehicleMetaDataOk returns a tuple with the IncludeVehicleMetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeVehicleMetaData

`func (o *VehicleOptionsModel) SetIncludeVehicleMetaData(v bool)`

SetIncludeVehicleMetaData sets IncludeVehicleMetaData field to given value.

### HasIncludeVehicleMetaData

`func (o *VehicleOptionsModel) HasIncludeVehicleMetaData() bool`

HasIncludeVehicleMetaData returns a boolean if a field has been set.

### GetIncludeVehicleLinks

`func (o *VehicleOptionsModel) GetIncludeVehicleLinks() bool`

GetIncludeVehicleLinks returns the IncludeVehicleLinks field if non-nil, zero value otherwise.

### GetIncludeVehicleLinksOk

`func (o *VehicleOptionsModel) GetIncludeVehicleLinksOk() (*bool, bool)`

GetIncludeVehicleLinksOk returns a tuple with the IncludeVehicleLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeVehicleLinks

`func (o *VehicleOptionsModel) SetIncludeVehicleLinks(v bool)`

SetIncludeVehicleLinks sets IncludeVehicleLinks field to given value.

### HasIncludeVehicleLinks

`func (o *VehicleOptionsModel) HasIncludeVehicleLinks() bool`

HasIncludeVehicleLinks returns a boolean if a field has been set.

### GetIncludeVehicleTags

`func (o *VehicleOptionsModel) GetIncludeVehicleTags() bool`

GetIncludeVehicleTags returns the IncludeVehicleTags field if non-nil, zero value otherwise.

### GetIncludeVehicleTagsOk

`func (o *VehicleOptionsModel) GetIncludeVehicleTagsOk() (*bool, bool)`

GetIncludeVehicleTagsOk returns a tuple with the IncludeVehicleTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeVehicleTags

`func (o *VehicleOptionsModel) SetIncludeVehicleTags(v bool)`

SetIncludeVehicleTags sets IncludeVehicleTags field to given value.

### HasIncludeVehicleTags

`func (o *VehicleOptionsModel) HasIncludeVehicleTags() bool`

HasIncludeVehicleTags returns a boolean if a field has been set.

### GetIncludeTagTypeName

`func (o *VehicleOptionsModel) GetIncludeTagTypeName() bool`

GetIncludeTagTypeName returns the IncludeTagTypeName field if non-nil, zero value otherwise.

### GetIncludeTagTypeNameOk

`func (o *VehicleOptionsModel) GetIncludeTagTypeNameOk() (*bool, bool)`

GetIncludeTagTypeNameOk returns a tuple with the IncludeTagTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTagTypeName

`func (o *VehicleOptionsModel) SetIncludeTagTypeName(v bool)`

SetIncludeTagTypeName sets IncludeTagTypeName field to given value.

### HasIncludeTagTypeName

`func (o *VehicleOptionsModel) HasIncludeTagTypeName() bool`

HasIncludeTagTypeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


