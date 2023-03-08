# UnsuccessfulModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityId** | Pointer to **int64** | Unsuccessful activity identifier | [optional] 
**ReasonId** | Pointer to **int64** | Reason why activity was unsuccessful identifier | [optional] 
**Reason** | Pointer to **string** | Reason why activity was unsuccessful content text | [optional] 

## Methods

### NewUnsuccessfulModel

`func NewUnsuccessfulModel() *UnsuccessfulModel`

NewUnsuccessfulModel instantiates a new UnsuccessfulModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnsuccessfulModelWithDefaults

`func NewUnsuccessfulModelWithDefaults() *UnsuccessfulModel`

NewUnsuccessfulModelWithDefaults instantiates a new UnsuccessfulModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityId

`func (o *UnsuccessfulModel) GetActivityId() int64`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *UnsuccessfulModel) GetActivityIdOk() (*int64, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *UnsuccessfulModel) SetActivityId(v int64)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *UnsuccessfulModel) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### GetReasonId

`func (o *UnsuccessfulModel) GetReasonId() int64`

GetReasonId returns the ReasonId field if non-nil, zero value otherwise.

### GetReasonIdOk

`func (o *UnsuccessfulModel) GetReasonIdOk() (*int64, bool)`

GetReasonIdOk returns a tuple with the ReasonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonId

`func (o *UnsuccessfulModel) SetReasonId(v int64)`

SetReasonId sets ReasonId field to given value.

### HasReasonId

`func (o *UnsuccessfulModel) HasReasonId() bool`

HasReasonId returns a boolean if a field has been set.

### GetReason

`func (o *UnsuccessfulModel) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UnsuccessfulModel) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UnsuccessfulModel) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *UnsuccessfulModel) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


