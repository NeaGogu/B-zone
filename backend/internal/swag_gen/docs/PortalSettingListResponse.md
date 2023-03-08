# PortalSettingListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]PortalSettingModel**](PortalSettingModel.md) |  | [optional] 
**CountFiltered** | Pointer to **int32** | Count of total items with filters in place | [optional] 
**CountUnfiltered** | Pointer to **int32** | Count of total items without filters in place | [optional] 
**CountLimited** | Pointer to **int32** | Count of items with limit in place | [optional] 

## Methods

### NewPortalSettingListResponse

`func NewPortalSettingListResponse() *PortalSettingListResponse`

NewPortalSettingListResponse instantiates a new PortalSettingListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalSettingListResponseWithDefaults

`func NewPortalSettingListResponseWithDefaults() *PortalSettingListResponse`

NewPortalSettingListResponseWithDefaults instantiates a new PortalSettingListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *PortalSettingListResponse) GetItems() []PortalSettingModel`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PortalSettingListResponse) GetItemsOk() (*[]PortalSettingModel, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PortalSettingListResponse) SetItems(v []PortalSettingModel)`

SetItems sets Items field to given value.

### HasItems

`func (o *PortalSettingListResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetCountFiltered

`func (o *PortalSettingListResponse) GetCountFiltered() int32`

GetCountFiltered returns the CountFiltered field if non-nil, zero value otherwise.

### GetCountFilteredOk

`func (o *PortalSettingListResponse) GetCountFilteredOk() (*int32, bool)`

GetCountFilteredOk returns a tuple with the CountFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountFiltered

`func (o *PortalSettingListResponse) SetCountFiltered(v int32)`

SetCountFiltered sets CountFiltered field to given value.

### HasCountFiltered

`func (o *PortalSettingListResponse) HasCountFiltered() bool`

HasCountFiltered returns a boolean if a field has been set.

### GetCountUnfiltered

`func (o *PortalSettingListResponse) GetCountUnfiltered() int32`

GetCountUnfiltered returns the CountUnfiltered field if non-nil, zero value otherwise.

### GetCountUnfilteredOk

`func (o *PortalSettingListResponse) GetCountUnfilteredOk() (*int32, bool)`

GetCountUnfilteredOk returns a tuple with the CountUnfiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountUnfiltered

`func (o *PortalSettingListResponse) SetCountUnfiltered(v int32)`

SetCountUnfiltered sets CountUnfiltered field to given value.

### HasCountUnfiltered

`func (o *PortalSettingListResponse) HasCountUnfiltered() bool`

HasCountUnfiltered returns a boolean if a field has been set.

### GetCountLimited

`func (o *PortalSettingListResponse) GetCountLimited() int32`

GetCountLimited returns the CountLimited field if non-nil, zero value otherwise.

### GetCountLimitedOk

`func (o *PortalSettingListResponse) GetCountLimitedOk() (*int32, bool)`

GetCountLimitedOk returns a tuple with the CountLimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountLimited

`func (o *PortalSettingListResponse) SetCountLimited(v int32)`

SetCountLimited sets CountLimited field to given value.

### HasCountLimited

`func (o *PortalSettingListResponse) HasCountLimited() bool`

HasCountLimited returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


