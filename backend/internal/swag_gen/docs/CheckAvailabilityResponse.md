# CheckAvailabilityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | token for the check availability job | [optional] 
**Status** | Pointer to **string** | current status for request | [optional] 
**AvailableTimewindows** | Pointer to [**[]AvailabilityTimeSlotModel**](AvailabilityTimeSlotModel.md) |  | [optional] 
**UnavailableTimewindows** | Pointer to [**[]AvailabilityTimeSlotModel**](AvailabilityTimeSlotModel.md) |  | [optional] 
**LatestAnalyzedDate** | Pointer to **string** |  | [optional] 

## Methods

### NewCheckAvailabilityResponse

`func NewCheckAvailabilityResponse() *CheckAvailabilityResponse`

NewCheckAvailabilityResponse instantiates a new CheckAvailabilityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckAvailabilityResponseWithDefaults

`func NewCheckAvailabilityResponseWithDefaults() *CheckAvailabilityResponse`

NewCheckAvailabilityResponseWithDefaults instantiates a new CheckAvailabilityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *CheckAvailabilityResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CheckAvailabilityResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CheckAvailabilityResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CheckAvailabilityResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetStatus

`func (o *CheckAvailabilityResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CheckAvailabilityResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CheckAvailabilityResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CheckAvailabilityResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAvailableTimewindows

`func (o *CheckAvailabilityResponse) GetAvailableTimewindows() []AvailabilityTimeSlotModel`

GetAvailableTimewindows returns the AvailableTimewindows field if non-nil, zero value otherwise.

### GetAvailableTimewindowsOk

`func (o *CheckAvailabilityResponse) GetAvailableTimewindowsOk() (*[]AvailabilityTimeSlotModel, bool)`

GetAvailableTimewindowsOk returns a tuple with the AvailableTimewindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableTimewindows

`func (o *CheckAvailabilityResponse) SetAvailableTimewindows(v []AvailabilityTimeSlotModel)`

SetAvailableTimewindows sets AvailableTimewindows field to given value.

### HasAvailableTimewindows

`func (o *CheckAvailabilityResponse) HasAvailableTimewindows() bool`

HasAvailableTimewindows returns a boolean if a field has been set.

### GetUnavailableTimewindows

`func (o *CheckAvailabilityResponse) GetUnavailableTimewindows() []AvailabilityTimeSlotModel`

GetUnavailableTimewindows returns the UnavailableTimewindows field if non-nil, zero value otherwise.

### GetUnavailableTimewindowsOk

`func (o *CheckAvailabilityResponse) GetUnavailableTimewindowsOk() (*[]AvailabilityTimeSlotModel, bool)`

GetUnavailableTimewindowsOk returns a tuple with the UnavailableTimewindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailableTimewindows

`func (o *CheckAvailabilityResponse) SetUnavailableTimewindows(v []AvailabilityTimeSlotModel)`

SetUnavailableTimewindows sets UnavailableTimewindows field to given value.

### HasUnavailableTimewindows

`func (o *CheckAvailabilityResponse) HasUnavailableTimewindows() bool`

HasUnavailableTimewindows returns a boolean if a field has been set.

### GetLatestAnalyzedDate

`func (o *CheckAvailabilityResponse) GetLatestAnalyzedDate() string`

GetLatestAnalyzedDate returns the LatestAnalyzedDate field if non-nil, zero value otherwise.

### GetLatestAnalyzedDateOk

`func (o *CheckAvailabilityResponse) GetLatestAnalyzedDateOk() (*string, bool)`

GetLatestAnalyzedDateOk returns a tuple with the LatestAnalyzedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestAnalyzedDate

`func (o *CheckAvailabilityResponse) SetLatestAnalyzedDate(v string)`

SetLatestAnalyzedDate sets LatestAnalyzedDate field to given value.

### HasLatestAnalyzedDate

`func (o *CheckAvailabilityResponse) HasLatestAnalyzedDate() bool`

HasLatestAnalyzedDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


