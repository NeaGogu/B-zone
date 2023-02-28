# ManualWebHookModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectId** | Pointer to **int64** | Unique Identifier | [optional] 
**WebHookName** | Pointer to **string** | Name of this Web Hook | [optional] 
**ExtraPayload** | Pointer to [**[]PayloadItem**](PayloadItem.md) | extra payload to be sent when the webhook is triggered | [optional] 

## Methods

### NewManualWebHookModel

`func NewManualWebHookModel() *ManualWebHookModel`

NewManualWebHookModel instantiates a new ManualWebHookModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManualWebHookModelWithDefaults

`func NewManualWebHookModelWithDefaults() *ManualWebHookModel`

NewManualWebHookModelWithDefaults instantiates a new ManualWebHookModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectId

`func (o *ManualWebHookModel) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *ManualWebHookModel) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *ManualWebHookModel) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *ManualWebHookModel) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetWebHookName

`func (o *ManualWebHookModel) GetWebHookName() string`

GetWebHookName returns the WebHookName field if non-nil, zero value otherwise.

### GetWebHookNameOk

`func (o *ManualWebHookModel) GetWebHookNameOk() (*string, bool)`

GetWebHookNameOk returns a tuple with the WebHookName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebHookName

`func (o *ManualWebHookModel) SetWebHookName(v string)`

SetWebHookName sets WebHookName field to given value.

### HasWebHookName

`func (o *ManualWebHookModel) HasWebHookName() bool`

HasWebHookName returns a boolean if a field has been set.

### GetExtraPayload

`func (o *ManualWebHookModel) GetExtraPayload() []PayloadItem`

GetExtraPayload returns the ExtraPayload field if non-nil, zero value otherwise.

### GetExtraPayloadOk

`func (o *ManualWebHookModel) GetExtraPayloadOk() (*[]PayloadItem, bool)`

GetExtraPayloadOk returns a tuple with the ExtraPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraPayload

`func (o *ManualWebHookModel) SetExtraPayload(v []PayloadItem)`

SetExtraPayload sets ExtraPayload field to given value.

### HasExtraPayload

`func (o *ManualWebHookModel) HasExtraPayload() bool`

HasExtraPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


