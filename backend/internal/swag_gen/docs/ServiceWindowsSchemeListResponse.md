# ServiceWindowsSchemeListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]ServiceWindowsSchemeModel**](ServiceWindowsSchemeModel.md) |  | [optional] 
**CountFiltered** | Pointer to **int32** | Count of total items with filters in place | [optional] 
**CountUnfiltered** | Pointer to **int32** | Count of total items without filters in place | [optional] 
**CountLimited** | Pointer to **int32** | Count of items with limit in place | [optional] 

## Methods

### NewServiceWindowsSchemeListResponse

`func NewServiceWindowsSchemeListResponse() *ServiceWindowsSchemeListResponse`

NewServiceWindowsSchemeListResponse instantiates a new ServiceWindowsSchemeListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWindowsSchemeListResponseWithDefaults

`func NewServiceWindowsSchemeListResponseWithDefaults() *ServiceWindowsSchemeListResponse`

NewServiceWindowsSchemeListResponseWithDefaults instantiates a new ServiceWindowsSchemeListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ServiceWindowsSchemeListResponse) GetItems() []ServiceWindowsSchemeModel`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ServiceWindowsSchemeListResponse) GetItemsOk() (*[]ServiceWindowsSchemeModel, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ServiceWindowsSchemeListResponse) SetItems(v []ServiceWindowsSchemeModel)`

SetItems sets Items field to given value.

### HasItems

`func (o *ServiceWindowsSchemeListResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetCountFiltered

`func (o *ServiceWindowsSchemeListResponse) GetCountFiltered() int32`

GetCountFiltered returns the CountFiltered field if non-nil, zero value otherwise.

### GetCountFilteredOk

`func (o *ServiceWindowsSchemeListResponse) GetCountFilteredOk() (*int32, bool)`

GetCountFilteredOk returns a tuple with the CountFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountFiltered

`func (o *ServiceWindowsSchemeListResponse) SetCountFiltered(v int32)`

SetCountFiltered sets CountFiltered field to given value.

### HasCountFiltered

`func (o *ServiceWindowsSchemeListResponse) HasCountFiltered() bool`

HasCountFiltered returns a boolean if a field has been set.

### GetCountUnfiltered

`func (o *ServiceWindowsSchemeListResponse) GetCountUnfiltered() int32`

GetCountUnfiltered returns the CountUnfiltered field if non-nil, zero value otherwise.

### GetCountUnfilteredOk

`func (o *ServiceWindowsSchemeListResponse) GetCountUnfilteredOk() (*int32, bool)`

GetCountUnfilteredOk returns a tuple with the CountUnfiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountUnfiltered

`func (o *ServiceWindowsSchemeListResponse) SetCountUnfiltered(v int32)`

SetCountUnfiltered sets CountUnfiltered field to given value.

### HasCountUnfiltered

`func (o *ServiceWindowsSchemeListResponse) HasCountUnfiltered() bool`

HasCountUnfiltered returns a boolean if a field has been set.

### GetCountLimited

`func (o *ServiceWindowsSchemeListResponse) GetCountLimited() int32`

GetCountLimited returns the CountLimited field if non-nil, zero value otherwise.

### GetCountLimitedOk

`func (o *ServiceWindowsSchemeListResponse) GetCountLimitedOk() (*int32, bool)`

GetCountLimitedOk returns a tuple with the CountLimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountLimited

`func (o *ServiceWindowsSchemeListResponse) SetCountLimited(v int32)`

SetCountLimited sets CountLimited field to given value.

### HasCountLimited

`func (o *ServiceWindowsSchemeListResponse) HasCountLimited() bool`

HasCountLimited returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

