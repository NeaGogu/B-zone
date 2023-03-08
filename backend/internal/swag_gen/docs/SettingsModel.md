# SettingsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**SettingsGroupId** | Pointer to **int32** | SettingsGroup id of this setting. Possible values: 1: general, 2: address, 3: package, 4: activity, 5: equipment, 6: note, 7: optimisation, 8: filters | [optional] 
**SettingsGroupName** | Pointer to **string** | SettingsGroup name of this setting | [optional] 
**Key** | Pointer to **string** | Unique string key for setting identification | [optional] 
**Value** | Pointer to **string** | Set value for setting | [optional] 
**ValueOptions** | Pointer to [**[]ValueOptionModel**](ValueOptionModel.md) |  | [optional] 
**Obscured** | Pointer to **bool** | If the return value is obscured because it is sensitive data | [optional] [readonly] 
**SettingUpdatedAt** | Pointer to **time.Time** | updated_at date time | [optional] [readonly] 
**SettingUpdatedBy** | Pointer to **int32** | updated_by user id | [optional] [readonly] 
**SettingUpdatedByUser** | Pointer to [**UsersModel**](UsersModel.md) |  | [optional] 

## Methods

### NewSettingsModel

`func NewSettingsModel() *SettingsModel`

NewSettingsModel instantiates a new SettingsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsModelWithDefaults

`func NewSettingsModelWithDefaults() *SettingsModel`

NewSettingsModelWithDefaults instantiates a new SettingsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SettingsModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SettingsModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SettingsModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SettingsModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSettingsGroupId

`func (o *SettingsModel) GetSettingsGroupId() int32`

GetSettingsGroupId returns the SettingsGroupId field if non-nil, zero value otherwise.

### GetSettingsGroupIdOk

`func (o *SettingsModel) GetSettingsGroupIdOk() (*int32, bool)`

GetSettingsGroupIdOk returns a tuple with the SettingsGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingsGroupId

`func (o *SettingsModel) SetSettingsGroupId(v int32)`

SetSettingsGroupId sets SettingsGroupId field to given value.

### HasSettingsGroupId

`func (o *SettingsModel) HasSettingsGroupId() bool`

HasSettingsGroupId returns a boolean if a field has been set.

### GetSettingsGroupName

`func (o *SettingsModel) GetSettingsGroupName() string`

GetSettingsGroupName returns the SettingsGroupName field if non-nil, zero value otherwise.

### GetSettingsGroupNameOk

`func (o *SettingsModel) GetSettingsGroupNameOk() (*string, bool)`

GetSettingsGroupNameOk returns a tuple with the SettingsGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingsGroupName

`func (o *SettingsModel) SetSettingsGroupName(v string)`

SetSettingsGroupName sets SettingsGroupName field to given value.

### HasSettingsGroupName

`func (o *SettingsModel) HasSettingsGroupName() bool`

HasSettingsGroupName returns a boolean if a field has been set.

### GetKey

`func (o *SettingsModel) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SettingsModel) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SettingsModel) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SettingsModel) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *SettingsModel) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SettingsModel) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SettingsModel) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SettingsModel) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueOptions

`func (o *SettingsModel) GetValueOptions() []ValueOptionModel`

GetValueOptions returns the ValueOptions field if non-nil, zero value otherwise.

### GetValueOptionsOk

`func (o *SettingsModel) GetValueOptionsOk() (*[]ValueOptionModel, bool)`

GetValueOptionsOk returns a tuple with the ValueOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueOptions

`func (o *SettingsModel) SetValueOptions(v []ValueOptionModel)`

SetValueOptions sets ValueOptions field to given value.

### HasValueOptions

`func (o *SettingsModel) HasValueOptions() bool`

HasValueOptions returns a boolean if a field has been set.

### GetObscured

`func (o *SettingsModel) GetObscured() bool`

GetObscured returns the Obscured field if non-nil, zero value otherwise.

### GetObscuredOk

`func (o *SettingsModel) GetObscuredOk() (*bool, bool)`

GetObscuredOk returns a tuple with the Obscured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObscured

`func (o *SettingsModel) SetObscured(v bool)`

SetObscured sets Obscured field to given value.

### HasObscured

`func (o *SettingsModel) HasObscured() bool`

HasObscured returns a boolean if a field has been set.

### GetSettingUpdatedAt

`func (o *SettingsModel) GetSettingUpdatedAt() time.Time`

GetSettingUpdatedAt returns the SettingUpdatedAt field if non-nil, zero value otherwise.

### GetSettingUpdatedAtOk

`func (o *SettingsModel) GetSettingUpdatedAtOk() (*time.Time, bool)`

GetSettingUpdatedAtOk returns a tuple with the SettingUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingUpdatedAt

`func (o *SettingsModel) SetSettingUpdatedAt(v time.Time)`

SetSettingUpdatedAt sets SettingUpdatedAt field to given value.

### HasSettingUpdatedAt

`func (o *SettingsModel) HasSettingUpdatedAt() bool`

HasSettingUpdatedAt returns a boolean if a field has been set.

### GetSettingUpdatedBy

`func (o *SettingsModel) GetSettingUpdatedBy() int32`

GetSettingUpdatedBy returns the SettingUpdatedBy field if non-nil, zero value otherwise.

### GetSettingUpdatedByOk

`func (o *SettingsModel) GetSettingUpdatedByOk() (*int32, bool)`

GetSettingUpdatedByOk returns a tuple with the SettingUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingUpdatedBy

`func (o *SettingsModel) SetSettingUpdatedBy(v int32)`

SetSettingUpdatedBy sets SettingUpdatedBy field to given value.

### HasSettingUpdatedBy

`func (o *SettingsModel) HasSettingUpdatedBy() bool`

HasSettingUpdatedBy returns a boolean if a field has been set.

### GetSettingUpdatedByUser

`func (o *SettingsModel) GetSettingUpdatedByUser() UsersModel`

GetSettingUpdatedByUser returns the SettingUpdatedByUser field if non-nil, zero value otherwise.

### GetSettingUpdatedByUserOk

`func (o *SettingsModel) GetSettingUpdatedByUserOk() (*UsersModel, bool)`

GetSettingUpdatedByUserOk returns a tuple with the SettingUpdatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingUpdatedByUser

`func (o *SettingsModel) SetSettingUpdatedByUser(v UsersModel)`

SetSettingUpdatedByUser sets SettingUpdatedByUser field to given value.

### HasSettingUpdatedByUser

`func (o *SettingsModel) HasSettingUpdatedByUser() bool`

HasSettingUpdatedByUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


