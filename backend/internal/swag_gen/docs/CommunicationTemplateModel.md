# CommunicationTemplateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Identifier | [optional] [readonly] 
**Name** | Pointer to **string** | Name of this Template | [optional] 
**Description** | Pointer to **string** | Description of this Template | [optional] 
**Subject** | Pointer to **string** | Subject used in email | [optional] 
**ContentHtml** | Pointer to **string** | Mustache based HTML content template | [optional] 
**ContentText** | Pointer to **string** | Mustache based Text content template | [optional] 
**IncludeAttachments** | Pointer to **bool** | Include the files associated with this activity as attachment, mail only | [optional] 
**Mappings** | Pointer to [**[]CommunicationMappingModel**](CommunicationMappingModel.md) |  | [optional] 

## Methods

### NewCommunicationTemplateModel

`func NewCommunicationTemplateModel() *CommunicationTemplateModel`

NewCommunicationTemplateModel instantiates a new CommunicationTemplateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunicationTemplateModelWithDefaults

`func NewCommunicationTemplateModelWithDefaults() *CommunicationTemplateModel`

NewCommunicationTemplateModelWithDefaults instantiates a new CommunicationTemplateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommunicationTemplateModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommunicationTemplateModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommunicationTemplateModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CommunicationTemplateModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CommunicationTemplateModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommunicationTemplateModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommunicationTemplateModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CommunicationTemplateModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CommunicationTemplateModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CommunicationTemplateModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CommunicationTemplateModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CommunicationTemplateModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSubject

`func (o *CommunicationTemplateModel) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CommunicationTemplateModel) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CommunicationTemplateModel) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CommunicationTemplateModel) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetContentHtml

`func (o *CommunicationTemplateModel) GetContentHtml() string`

GetContentHtml returns the ContentHtml field if non-nil, zero value otherwise.

### GetContentHtmlOk

`func (o *CommunicationTemplateModel) GetContentHtmlOk() (*string, bool)`

GetContentHtmlOk returns a tuple with the ContentHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentHtml

`func (o *CommunicationTemplateModel) SetContentHtml(v string)`

SetContentHtml sets ContentHtml field to given value.

### HasContentHtml

`func (o *CommunicationTemplateModel) HasContentHtml() bool`

HasContentHtml returns a boolean if a field has been set.

### GetContentText

`func (o *CommunicationTemplateModel) GetContentText() string`

GetContentText returns the ContentText field if non-nil, zero value otherwise.

### GetContentTextOk

`func (o *CommunicationTemplateModel) GetContentTextOk() (*string, bool)`

GetContentTextOk returns a tuple with the ContentText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentText

`func (o *CommunicationTemplateModel) SetContentText(v string)`

SetContentText sets ContentText field to given value.

### HasContentText

`func (o *CommunicationTemplateModel) HasContentText() bool`

HasContentText returns a boolean if a field has been set.

### GetIncludeAttachments

`func (o *CommunicationTemplateModel) GetIncludeAttachments() bool`

GetIncludeAttachments returns the IncludeAttachments field if non-nil, zero value otherwise.

### GetIncludeAttachmentsOk

`func (o *CommunicationTemplateModel) GetIncludeAttachmentsOk() (*bool, bool)`

GetIncludeAttachmentsOk returns a tuple with the IncludeAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAttachments

`func (o *CommunicationTemplateModel) SetIncludeAttachments(v bool)`

SetIncludeAttachments sets IncludeAttachments field to given value.

### HasIncludeAttachments

`func (o *CommunicationTemplateModel) HasIncludeAttachments() bool`

HasIncludeAttachments returns a boolean if a field has been set.

### GetMappings

`func (o *CommunicationTemplateModel) GetMappings() []CommunicationMappingModel`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *CommunicationTemplateModel) GetMappingsOk() (*[]CommunicationMappingModel, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *CommunicationTemplateModel) SetMappings(v []CommunicationMappingModel)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *CommunicationTemplateModel) HasMappings() bool`

HasMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


