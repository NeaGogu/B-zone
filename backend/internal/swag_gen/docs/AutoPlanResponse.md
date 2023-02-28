# AutoPlanResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** | token for the auto plan job | [optional] 
**Status** | Pointer to **string** | current status for request | [optional] 
**AffectedActivities** | Pointer to [**[]ActivityModel**](ActivityModel.md) |  | [optional] 
**UnavailableTimewindows** | Pointer to [**[]TimeSlotModel**](TimeSlotModel.md) |  | [optional] 
**LatestAnalyzedDate** | Pointer to **string** |  | [optional] 

## Methods

### NewAutoPlanResponse

`func NewAutoPlanResponse() *AutoPlanResponse`

NewAutoPlanResponse instantiates a new AutoPlanResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoPlanResponseWithDefaults

`func NewAutoPlanResponseWithDefaults() *AutoPlanResponse`

NewAutoPlanResponseWithDefaults instantiates a new AutoPlanResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *AutoPlanResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AutoPlanResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AutoPlanResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AutoPlanResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetStatus

`func (o *AutoPlanResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AutoPlanResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AutoPlanResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AutoPlanResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAffectedActivities

`func (o *AutoPlanResponse) GetAffectedActivities() []ActivityModel`

GetAffectedActivities returns the AffectedActivities field if non-nil, zero value otherwise.

### GetAffectedActivitiesOk

`func (o *AutoPlanResponse) GetAffectedActivitiesOk() (*[]ActivityModel, bool)`

GetAffectedActivitiesOk returns a tuple with the AffectedActivities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedActivities

`func (o *AutoPlanResponse) SetAffectedActivities(v []ActivityModel)`

SetAffectedActivities sets AffectedActivities field to given value.

### HasAffectedActivities

`func (o *AutoPlanResponse) HasAffectedActivities() bool`

HasAffectedActivities returns a boolean if a field has been set.

### GetUnavailableTimewindows

`func (o *AutoPlanResponse) GetUnavailableTimewindows() []TimeSlotModel`

GetUnavailableTimewindows returns the UnavailableTimewindows field if non-nil, zero value otherwise.

### GetUnavailableTimewindowsOk

`func (o *AutoPlanResponse) GetUnavailableTimewindowsOk() (*[]TimeSlotModel, bool)`

GetUnavailableTimewindowsOk returns a tuple with the UnavailableTimewindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailableTimewindows

`func (o *AutoPlanResponse) SetUnavailableTimewindows(v []TimeSlotModel)`

SetUnavailableTimewindows sets UnavailableTimewindows field to given value.

### HasUnavailableTimewindows

`func (o *AutoPlanResponse) HasUnavailableTimewindows() bool`

HasUnavailableTimewindows returns a boolean if a field has been set.

### GetLatestAnalyzedDate

`func (o *AutoPlanResponse) GetLatestAnalyzedDate() string`

GetLatestAnalyzedDate returns the LatestAnalyzedDate field if non-nil, zero value otherwise.

### GetLatestAnalyzedDateOk

`func (o *AutoPlanResponse) GetLatestAnalyzedDateOk() (*string, bool)`

GetLatestAnalyzedDateOk returns a tuple with the LatestAnalyzedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestAnalyzedDate

`func (o *AutoPlanResponse) SetLatestAnalyzedDate(v string)`

SetLatestAnalyzedDate sets LatestAnalyzedDate field to given value.

### HasLatestAnalyzedDate

`func (o *AutoPlanResponse) HasLatestAnalyzedDate() bool`

HasLatestAnalyzedDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


