# WebhookModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectId** | **int32** | objectId | 
**WebHookName** | **string** | Name of this Web Hook | 
**ExtraPayload** | Pointer to [**[]PayloadItem**](PayloadItem.md) | extra payload to be sent when the webhook is triggered | [optional] 

## Methods

### NewWebhookModel

`func NewWebhookModel(objectId int32, webHookName string, ) *WebhookModel`

NewWebhookModel instantiates a new WebhookModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookModelWithDefaults

`func NewWebhookModelWithDefaults() *WebhookModel`

NewWebhookModelWithDefaults instantiates a new WebhookModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectId

`func (o *WebhookModel) GetObjectId() int32`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *WebhookModel) GetObjectIdOk() (*int32, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *WebhookModel) SetObjectId(v int32)`

SetObjectId sets ObjectId field to given value.


### GetWebHookName

`func (o *WebhookModel) GetWebHookName() string`

GetWebHookName returns the WebHookName field if non-nil, zero value otherwise.

### GetWebHookNameOk

`func (o *WebhookModel) GetWebHookNameOk() (*string, bool)`

GetWebHookNameOk returns a tuple with the WebHookName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebHookName

`func (o *WebhookModel) SetWebHookName(v string)`

SetWebHookName sets WebHookName field to given value.


### GetExtraPayload

`func (o *WebhookModel) GetExtraPayload() []PayloadItem`

GetExtraPayload returns the ExtraPayload field if non-nil, zero value otherwise.

### GetExtraPayloadOk

`func (o *WebhookModel) GetExtraPayloadOk() (*[]PayloadItem, bool)`

GetExtraPayloadOk returns a tuple with the ExtraPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraPayload

`func (o *WebhookModel) SetExtraPayload(v []PayloadItem)`

SetExtraPayload sets ExtraPayload field to given value.

### HasExtraPayload

`func (o *WebhookModel) HasExtraPayload() bool`

HasExtraPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


