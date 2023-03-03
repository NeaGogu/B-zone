# CommunicationTriggerMessageModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityId** | Pointer to **int64** | Unique Identifier | [optional] [readonly] 
**MessageType** | Pointer to **string** | Type of message | [optional] 
**CheckPreference** | Pointer to **bool** | check preference | [optional] 

## Methods

### NewCommunicationTriggerMessageModel

`func NewCommunicationTriggerMessageModel() *CommunicationTriggerMessageModel`

NewCommunicationTriggerMessageModel instantiates a new CommunicationTriggerMessageModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunicationTriggerMessageModelWithDefaults

`func NewCommunicationTriggerMessageModelWithDefaults() *CommunicationTriggerMessageModel`

NewCommunicationTriggerMessageModelWithDefaults instantiates a new CommunicationTriggerMessageModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityId

`func (o *CommunicationTriggerMessageModel) GetActivityId() int64`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *CommunicationTriggerMessageModel) GetActivityIdOk() (*int64, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *CommunicationTriggerMessageModel) SetActivityId(v int64)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *CommunicationTriggerMessageModel) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### GetMessageType

`func (o *CommunicationTriggerMessageModel) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *CommunicationTriggerMessageModel) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *CommunicationTriggerMessageModel) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *CommunicationTriggerMessageModel) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetCheckPreference

`func (o *CommunicationTriggerMessageModel) GetCheckPreference() bool`

GetCheckPreference returns the CheckPreference field if non-nil, zero value otherwise.

### GetCheckPreferenceOk

`func (o *CommunicationTriggerMessageModel) GetCheckPreferenceOk() (*bool, bool)`

GetCheckPreferenceOk returns a tuple with the CheckPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPreference

`func (o *CommunicationTriggerMessageModel) SetCheckPreference(v bool)`

SetCheckPreference sets CheckPreference field to given value.

### HasCheckPreference

`func (o *CommunicationTriggerMessageModel) HasCheckPreference() bool`

HasCheckPreference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


