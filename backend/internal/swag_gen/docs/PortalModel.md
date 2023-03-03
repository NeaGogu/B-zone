# PortalModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique Identifier | [optional] [readonly] 
**Name** | Pointer to **string** | Name of brand | [optional] [readonly] 
**PortalSettings** | Pointer to [**[]PortalSettingModel**](PortalSettingModel.md) |  | [optional] 

## Methods

### NewPortalModel

`func NewPortalModel() *PortalModel`

NewPortalModel instantiates a new PortalModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalModelWithDefaults

`func NewPortalModelWithDefaults() *PortalModel`

NewPortalModelWithDefaults instantiates a new PortalModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortalModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortalModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortalModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PortalModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PortalModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PortalModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PortalModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PortalModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPortalSettings

`func (o *PortalModel) GetPortalSettings() []PortalSettingModel`

GetPortalSettings returns the PortalSettings field if non-nil, zero value otherwise.

### GetPortalSettingsOk

`func (o *PortalModel) GetPortalSettingsOk() (*[]PortalSettingModel, bool)`

GetPortalSettingsOk returns a tuple with the PortalSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalSettings

`func (o *PortalModel) SetPortalSettings(v []PortalSettingModel)`

SetPortalSettings sets PortalSettings field to given value.

### HasPortalSettings

`func (o *PortalModel) HasPortalSettings() bool`

HasPortalSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


