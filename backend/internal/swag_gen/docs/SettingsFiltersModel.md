# SettingsFiltersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **[]int32** |  | [optional] 
**SettingsGroupId** | Pointer to **[]int32** | SettingsGroup id of this setting. Possible values: 1: general, 2: address, 3: package, 4: activity, 5: equipment, 6: note, 7: optimisation, 8: filters | [optional] 
**SettingsGroupName** | Pointer to **[]string** | SettingsGroup name of this setting | [optional] 
**Key** | Pointer to **[]string** | Unique string key for setting identification | [optional] 
**SearchText** | Pointer to **string** | free search through text and numeric type columns | [optional] 
**UpdatedAtSince** | Pointer to **time.Time** | Show updated since | [optional] 
**UpdatedAtTill** | Pointer to **time.Time** | Show updated till | [optional] 

## Methods

### NewSettingsFiltersModel

`func NewSettingsFiltersModel() *SettingsFiltersModel`

NewSettingsFiltersModel instantiates a new SettingsFiltersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsFiltersModelWithDefaults

`func NewSettingsFiltersModelWithDefaults() *SettingsFiltersModel`

NewSettingsFiltersModelWithDefaults instantiates a new SettingsFiltersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SettingsFiltersModel) GetId() []int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SettingsFiltersModel) GetIdOk() (*[]int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SettingsFiltersModel) SetId(v []int32)`

SetId sets Id field to given value.

### HasId

`func (o *SettingsFiltersModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSettingsGroupId

`func (o *SettingsFiltersModel) GetSettingsGroupId() []int32`

GetSettingsGroupId returns the SettingsGroupId field if non-nil, zero value otherwise.

### GetSettingsGroupIdOk

`func (o *SettingsFiltersModel) GetSettingsGroupIdOk() (*[]int32, bool)`

GetSettingsGroupIdOk returns a tuple with the SettingsGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingsGroupId

`func (o *SettingsFiltersModel) SetSettingsGroupId(v []int32)`

SetSettingsGroupId sets SettingsGroupId field to given value.

### HasSettingsGroupId

`func (o *SettingsFiltersModel) HasSettingsGroupId() bool`

HasSettingsGroupId returns a boolean if a field has been set.

### GetSettingsGroupName

`func (o *SettingsFiltersModel) GetSettingsGroupName() []string`

GetSettingsGroupName returns the SettingsGroupName field if non-nil, zero value otherwise.

### GetSettingsGroupNameOk

`func (o *SettingsFiltersModel) GetSettingsGroupNameOk() (*[]string, bool)`

GetSettingsGroupNameOk returns a tuple with the SettingsGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingsGroupName

`func (o *SettingsFiltersModel) SetSettingsGroupName(v []string)`

SetSettingsGroupName sets SettingsGroupName field to given value.

### HasSettingsGroupName

`func (o *SettingsFiltersModel) HasSettingsGroupName() bool`

HasSettingsGroupName returns a boolean if a field has been set.

### GetKey

`func (o *SettingsFiltersModel) GetKey() []string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SettingsFiltersModel) GetKeyOk() (*[]string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SettingsFiltersModel) SetKey(v []string)`

SetKey sets Key field to given value.

### HasKey

`func (o *SettingsFiltersModel) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSearchText

`func (o *SettingsFiltersModel) GetSearchText() string`

GetSearchText returns the SearchText field if non-nil, zero value otherwise.

### GetSearchTextOk

`func (o *SettingsFiltersModel) GetSearchTextOk() (*string, bool)`

GetSearchTextOk returns a tuple with the SearchText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchText

`func (o *SettingsFiltersModel) SetSearchText(v string)`

SetSearchText sets SearchText field to given value.

### HasSearchText

`func (o *SettingsFiltersModel) HasSearchText() bool`

HasSearchText returns a boolean if a field has been set.

### GetUpdatedAtSince

`func (o *SettingsFiltersModel) GetUpdatedAtSince() time.Time`

GetUpdatedAtSince returns the UpdatedAtSince field if non-nil, zero value otherwise.

### GetUpdatedAtSinceOk

`func (o *SettingsFiltersModel) GetUpdatedAtSinceOk() (*time.Time, bool)`

GetUpdatedAtSinceOk returns a tuple with the UpdatedAtSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtSince

`func (o *SettingsFiltersModel) SetUpdatedAtSince(v time.Time)`

SetUpdatedAtSince sets UpdatedAtSince field to given value.

### HasUpdatedAtSince

`func (o *SettingsFiltersModel) HasUpdatedAtSince() bool`

HasUpdatedAtSince returns a boolean if a field has been set.

### GetUpdatedAtTill

`func (o *SettingsFiltersModel) GetUpdatedAtTill() time.Time`

GetUpdatedAtTill returns the UpdatedAtTill field if non-nil, zero value otherwise.

### GetUpdatedAtTillOk

`func (o *SettingsFiltersModel) GetUpdatedAtTillOk() (*time.Time, bool)`

GetUpdatedAtTillOk returns a tuple with the UpdatedAtTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtTill

`func (o *SettingsFiltersModel) SetUpdatedAtTill(v time.Time)`

SetUpdatedAtTill sets UpdatedAtTill field to given value.

### HasUpdatedAtTill

`func (o *SettingsFiltersModel) HasUpdatedAtTill() bool`

HasUpdatedAtTill returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


