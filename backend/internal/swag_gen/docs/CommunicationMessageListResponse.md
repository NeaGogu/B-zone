# CommunicationMessageListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]CommunicationMessageModel**](CommunicationMessageModel.md) |  | [optional] 
**CountFiltered** | Pointer to **int32** | Count of total items with filters in place | [optional] 
**CountUnfiltered** | Pointer to **int32** | Count of total items without filters in place | [optional] 
**CountLimited** | Pointer to **int32** | Count of items with limit in place | [optional] 

## Methods

### NewCommunicationMessageListResponse

`func NewCommunicationMessageListResponse() *CommunicationMessageListResponse`

NewCommunicationMessageListResponse instantiates a new CommunicationMessageListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunicationMessageListResponseWithDefaults

`func NewCommunicationMessageListResponseWithDefaults() *CommunicationMessageListResponse`

NewCommunicationMessageListResponseWithDefaults instantiates a new CommunicationMessageListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CommunicationMessageListResponse) GetItems() []CommunicationMessageModel`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CommunicationMessageListResponse) GetItemsOk() (*[]CommunicationMessageModel, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CommunicationMessageListResponse) SetItems(v []CommunicationMessageModel)`

SetItems sets Items field to given value.

### HasItems

`func (o *CommunicationMessageListResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetCountFiltered

`func (o *CommunicationMessageListResponse) GetCountFiltered() int32`

GetCountFiltered returns the CountFiltered field if non-nil, zero value otherwise.

### GetCountFilteredOk

`func (o *CommunicationMessageListResponse) GetCountFilteredOk() (*int32, bool)`

GetCountFilteredOk returns a tuple with the CountFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountFiltered

`func (o *CommunicationMessageListResponse) SetCountFiltered(v int32)`

SetCountFiltered sets CountFiltered field to given value.

### HasCountFiltered

`func (o *CommunicationMessageListResponse) HasCountFiltered() bool`

HasCountFiltered returns a boolean if a field has been set.

### GetCountUnfiltered

`func (o *CommunicationMessageListResponse) GetCountUnfiltered() int32`

GetCountUnfiltered returns the CountUnfiltered field if non-nil, zero value otherwise.

### GetCountUnfilteredOk

`func (o *CommunicationMessageListResponse) GetCountUnfilteredOk() (*int32, bool)`

GetCountUnfilteredOk returns a tuple with the CountUnfiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountUnfiltered

`func (o *CommunicationMessageListResponse) SetCountUnfiltered(v int32)`

SetCountUnfiltered sets CountUnfiltered field to given value.

### HasCountUnfiltered

`func (o *CommunicationMessageListResponse) HasCountUnfiltered() bool`

HasCountUnfiltered returns a boolean if a field has been set.

### GetCountLimited

`func (o *CommunicationMessageListResponse) GetCountLimited() int32`

GetCountLimited returns the CountLimited field if non-nil, zero value otherwise.

### GetCountLimitedOk

`func (o *CommunicationMessageListResponse) GetCountLimitedOk() (*int32, bool)`

GetCountLimitedOk returns a tuple with the CountLimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountLimited

`func (o *CommunicationMessageListResponse) SetCountLimited(v int32)`

SetCountLimited sets CountLimited field to given value.

### HasCountLimited

`func (o *CommunicationMessageListResponse) HasCountLimited() bool`

HasCountLimited returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


