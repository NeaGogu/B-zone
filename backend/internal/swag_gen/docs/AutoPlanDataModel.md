# AutoPlanDataModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **int64** | unique per api request | [optional] [readonly] 
**AvailabilityKey** | Pointer to **string** | unique key from availability result, used to reuse fromer result | [optional] 
**Activity** | Pointer to [**[]ActivityModel**](ActivityModel.md) |  | [optional] 

## Methods

### NewAutoPlanDataModel

`func NewAutoPlanDataModel() *AutoPlanDataModel`

NewAutoPlanDataModel instantiates a new AutoPlanDataModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoPlanDataModelWithDefaults

`func NewAutoPlanDataModelWithDefaults() *AutoPlanDataModel`

NewAutoPlanDataModelWithDefaults instantiates a new AutoPlanDataModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *AutoPlanDataModel) GetToken() int64`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AutoPlanDataModel) GetTokenOk() (*int64, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AutoPlanDataModel) SetToken(v int64)`

SetToken sets Token field to given value.

### HasToken

`func (o *AutoPlanDataModel) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetAvailabilityKey

`func (o *AutoPlanDataModel) GetAvailabilityKey() string`

GetAvailabilityKey returns the AvailabilityKey field if non-nil, zero value otherwise.

### GetAvailabilityKeyOk

`func (o *AutoPlanDataModel) GetAvailabilityKeyOk() (*string, bool)`

GetAvailabilityKeyOk returns a tuple with the AvailabilityKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityKey

`func (o *AutoPlanDataModel) SetAvailabilityKey(v string)`

SetAvailabilityKey sets AvailabilityKey field to given value.

### HasAvailabilityKey

`func (o *AutoPlanDataModel) HasAvailabilityKey() bool`

HasAvailabilityKey returns a boolean if a field has been set.

### GetActivity

`func (o *AutoPlanDataModel) GetActivity() []ActivityModel`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *AutoPlanDataModel) GetActivityOk() (*[]ActivityModel, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *AutoPlanDataModel) SetActivity(v []ActivityModel)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *AutoPlanDataModel) HasActivity() bool`

HasActivity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


