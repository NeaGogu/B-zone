# CommunicationMessageModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Identifier | [optional] 
**ActivityId** | Pointer to **int32** | Activity ID | [optional] 
**CommunicationMappingId** | Pointer to **int32** | Activity ID | [optional] 
**CommunicationMapping** | Pointer to [**CommunicationTemplateModel**](CommunicationTemplateModel.md) |  | [optional] 
**Recipient** | Pointer to **string** | recipient | [optional] 
**Subject** | Pointer to **string** | subject | [optional] 
**HtmlContent** | Pointer to **string** | HTML content | [optional] 
**TextContent** | Pointer to **string** | text content | [optional] 
**CommunicationMessageHistory** | Pointer to [**[]CommunicationMessageHistoryModel**](CommunicationMessageHistoryModel.md) |  | [optional] 

## Methods

### NewCommunicationMessageModel

`func NewCommunicationMessageModel() *CommunicationMessageModel`

NewCommunicationMessageModel instantiates a new CommunicationMessageModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunicationMessageModelWithDefaults

`func NewCommunicationMessageModelWithDefaults() *CommunicationMessageModel`

NewCommunicationMessageModelWithDefaults instantiates a new CommunicationMessageModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommunicationMessageModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommunicationMessageModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommunicationMessageModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CommunicationMessageModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActivityId

`func (o *CommunicationMessageModel) GetActivityId() int32`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *CommunicationMessageModel) GetActivityIdOk() (*int32, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *CommunicationMessageModel) SetActivityId(v int32)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *CommunicationMessageModel) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### GetCommunicationMappingId

`func (o *CommunicationMessageModel) GetCommunicationMappingId() int32`

GetCommunicationMappingId returns the CommunicationMappingId field if non-nil, zero value otherwise.

### GetCommunicationMappingIdOk

`func (o *CommunicationMessageModel) GetCommunicationMappingIdOk() (*int32, bool)`

GetCommunicationMappingIdOk returns a tuple with the CommunicationMappingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationMappingId

`func (o *CommunicationMessageModel) SetCommunicationMappingId(v int32)`

SetCommunicationMappingId sets CommunicationMappingId field to given value.

### HasCommunicationMappingId

`func (o *CommunicationMessageModel) HasCommunicationMappingId() bool`

HasCommunicationMappingId returns a boolean if a field has been set.

### GetCommunicationMapping

`func (o *CommunicationMessageModel) GetCommunicationMapping() CommunicationTemplateModel`

GetCommunicationMapping returns the CommunicationMapping field if non-nil, zero value otherwise.

### GetCommunicationMappingOk

`func (o *CommunicationMessageModel) GetCommunicationMappingOk() (*CommunicationTemplateModel, bool)`

GetCommunicationMappingOk returns a tuple with the CommunicationMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationMapping

`func (o *CommunicationMessageModel) SetCommunicationMapping(v CommunicationTemplateModel)`

SetCommunicationMapping sets CommunicationMapping field to given value.

### HasCommunicationMapping

`func (o *CommunicationMessageModel) HasCommunicationMapping() bool`

HasCommunicationMapping returns a boolean if a field has been set.

### GetRecipient

`func (o *CommunicationMessageModel) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *CommunicationMessageModel) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *CommunicationMessageModel) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *CommunicationMessageModel) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetSubject

`func (o *CommunicationMessageModel) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CommunicationMessageModel) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CommunicationMessageModel) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CommunicationMessageModel) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetHtmlContent

`func (o *CommunicationMessageModel) GetHtmlContent() string`

GetHtmlContent returns the HtmlContent field if non-nil, zero value otherwise.

### GetHtmlContentOk

`func (o *CommunicationMessageModel) GetHtmlContentOk() (*string, bool)`

GetHtmlContentOk returns a tuple with the HtmlContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlContent

`func (o *CommunicationMessageModel) SetHtmlContent(v string)`

SetHtmlContent sets HtmlContent field to given value.

### HasHtmlContent

`func (o *CommunicationMessageModel) HasHtmlContent() bool`

HasHtmlContent returns a boolean if a field has been set.

### GetTextContent

`func (o *CommunicationMessageModel) GetTextContent() string`

GetTextContent returns the TextContent field if non-nil, zero value otherwise.

### GetTextContentOk

`func (o *CommunicationMessageModel) GetTextContentOk() (*string, bool)`

GetTextContentOk returns a tuple with the TextContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextContent

`func (o *CommunicationMessageModel) SetTextContent(v string)`

SetTextContent sets TextContent field to given value.

### HasTextContent

`func (o *CommunicationMessageModel) HasTextContent() bool`

HasTextContent returns a boolean if a field has been set.

### GetCommunicationMessageHistory

`func (o *CommunicationMessageModel) GetCommunicationMessageHistory() []CommunicationMessageHistoryModel`

GetCommunicationMessageHistory returns the CommunicationMessageHistory field if non-nil, zero value otherwise.

### GetCommunicationMessageHistoryOk

`func (o *CommunicationMessageModel) GetCommunicationMessageHistoryOk() (*[]CommunicationMessageHistoryModel, bool)`

GetCommunicationMessageHistoryOk returns a tuple with the CommunicationMessageHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationMessageHistory

`func (o *CommunicationMessageModel) SetCommunicationMessageHistory(v []CommunicationMessageHistoryModel)`

SetCommunicationMessageHistory sets CommunicationMessageHistory field to given value.

### HasCommunicationMessageHistory

`func (o *CommunicationMessageModel) HasCommunicationMessageHistory() bool`

HasCommunicationMessageHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


