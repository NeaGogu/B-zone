# CommunicationMappingModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Identifier | [optional] [readonly] 
**CommunicationTemplateId** | Pointer to **int32** | Template ID | [optional] 
**CommunicationTemplateName** | Pointer to **string** | Template Name | [optional] 
**CommunicationTemplate** | Pointer to [**CommunicationTemplateModel**](CommunicationTemplateModel.md) |  | [optional] 
**CommunicationDeliveryMethodId** | Pointer to **int32** | Delivery Method ID | [optional] 
**CommunicationDeliveryMethodName** | Pointer to **string** | Delivery Method Name | [optional] 
**CommunicationMessageTypeId** | Pointer to **int32** | MessageType ID | [optional] 
**CommunicationMessageTypeName** | Pointer to **string** | MessageType Name | [optional] 
**NoTags** | Pointer to **bool** | No Tag Flag | [optional] 
**TagNames** | Pointer to **[]string** | Tag names | [optional] 
**Tags** | Pointer to [**[]TagModel**](TagModel.md) |  | [optional] 
**NoZones** | Pointer to **bool** | No zones Flag | [optional] 
**Zones** | Pointer to [**[]ZoneModel**](ZoneModel.md) |  | [optional] 
**ZoneNames** | Pointer to **[]string** | Zone names | [optional] 

## Methods

### NewCommunicationMappingModel

`func NewCommunicationMappingModel() *CommunicationMappingModel`

NewCommunicationMappingModel instantiates a new CommunicationMappingModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommunicationMappingModelWithDefaults

`func NewCommunicationMappingModelWithDefaults() *CommunicationMappingModel`

NewCommunicationMappingModelWithDefaults instantiates a new CommunicationMappingModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommunicationMappingModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommunicationMappingModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommunicationMappingModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CommunicationMappingModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCommunicationTemplateId

`func (o *CommunicationMappingModel) GetCommunicationTemplateId() int32`

GetCommunicationTemplateId returns the CommunicationTemplateId field if non-nil, zero value otherwise.

### GetCommunicationTemplateIdOk

`func (o *CommunicationMappingModel) GetCommunicationTemplateIdOk() (*int32, bool)`

GetCommunicationTemplateIdOk returns a tuple with the CommunicationTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationTemplateId

`func (o *CommunicationMappingModel) SetCommunicationTemplateId(v int32)`

SetCommunicationTemplateId sets CommunicationTemplateId field to given value.

### HasCommunicationTemplateId

`func (o *CommunicationMappingModel) HasCommunicationTemplateId() bool`

HasCommunicationTemplateId returns a boolean if a field has been set.

### GetCommunicationTemplateName

`func (o *CommunicationMappingModel) GetCommunicationTemplateName() string`

GetCommunicationTemplateName returns the CommunicationTemplateName field if non-nil, zero value otherwise.

### GetCommunicationTemplateNameOk

`func (o *CommunicationMappingModel) GetCommunicationTemplateNameOk() (*string, bool)`

GetCommunicationTemplateNameOk returns a tuple with the CommunicationTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationTemplateName

`func (o *CommunicationMappingModel) SetCommunicationTemplateName(v string)`

SetCommunicationTemplateName sets CommunicationTemplateName field to given value.

### HasCommunicationTemplateName

`func (o *CommunicationMappingModel) HasCommunicationTemplateName() bool`

HasCommunicationTemplateName returns a boolean if a field has been set.

### GetCommunicationTemplate

`func (o *CommunicationMappingModel) GetCommunicationTemplate() CommunicationTemplateModel`

GetCommunicationTemplate returns the CommunicationTemplate field if non-nil, zero value otherwise.

### GetCommunicationTemplateOk

`func (o *CommunicationMappingModel) GetCommunicationTemplateOk() (*CommunicationTemplateModel, bool)`

GetCommunicationTemplateOk returns a tuple with the CommunicationTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationTemplate

`func (o *CommunicationMappingModel) SetCommunicationTemplate(v CommunicationTemplateModel)`

SetCommunicationTemplate sets CommunicationTemplate field to given value.

### HasCommunicationTemplate

`func (o *CommunicationMappingModel) HasCommunicationTemplate() bool`

HasCommunicationTemplate returns a boolean if a field has been set.

### GetCommunicationDeliveryMethodId

`func (o *CommunicationMappingModel) GetCommunicationDeliveryMethodId() int32`

GetCommunicationDeliveryMethodId returns the CommunicationDeliveryMethodId field if non-nil, zero value otherwise.

### GetCommunicationDeliveryMethodIdOk

`func (o *CommunicationMappingModel) GetCommunicationDeliveryMethodIdOk() (*int32, bool)`

GetCommunicationDeliveryMethodIdOk returns a tuple with the CommunicationDeliveryMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationDeliveryMethodId

`func (o *CommunicationMappingModel) SetCommunicationDeliveryMethodId(v int32)`

SetCommunicationDeliveryMethodId sets CommunicationDeliveryMethodId field to given value.

### HasCommunicationDeliveryMethodId

`func (o *CommunicationMappingModel) HasCommunicationDeliveryMethodId() bool`

HasCommunicationDeliveryMethodId returns a boolean if a field has been set.

### GetCommunicationDeliveryMethodName

`func (o *CommunicationMappingModel) GetCommunicationDeliveryMethodName() string`

GetCommunicationDeliveryMethodName returns the CommunicationDeliveryMethodName field if non-nil, zero value otherwise.

### GetCommunicationDeliveryMethodNameOk

`func (o *CommunicationMappingModel) GetCommunicationDeliveryMethodNameOk() (*string, bool)`

GetCommunicationDeliveryMethodNameOk returns a tuple with the CommunicationDeliveryMethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationDeliveryMethodName

`func (o *CommunicationMappingModel) SetCommunicationDeliveryMethodName(v string)`

SetCommunicationDeliveryMethodName sets CommunicationDeliveryMethodName field to given value.

### HasCommunicationDeliveryMethodName

`func (o *CommunicationMappingModel) HasCommunicationDeliveryMethodName() bool`

HasCommunicationDeliveryMethodName returns a boolean if a field has been set.

### GetCommunicationMessageTypeId

`func (o *CommunicationMappingModel) GetCommunicationMessageTypeId() int32`

GetCommunicationMessageTypeId returns the CommunicationMessageTypeId field if non-nil, zero value otherwise.

### GetCommunicationMessageTypeIdOk

`func (o *CommunicationMappingModel) GetCommunicationMessageTypeIdOk() (*int32, bool)`

GetCommunicationMessageTypeIdOk returns a tuple with the CommunicationMessageTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationMessageTypeId

`func (o *CommunicationMappingModel) SetCommunicationMessageTypeId(v int32)`

SetCommunicationMessageTypeId sets CommunicationMessageTypeId field to given value.

### HasCommunicationMessageTypeId

`func (o *CommunicationMappingModel) HasCommunicationMessageTypeId() bool`

HasCommunicationMessageTypeId returns a boolean if a field has been set.

### GetCommunicationMessageTypeName

`func (o *CommunicationMappingModel) GetCommunicationMessageTypeName() string`

GetCommunicationMessageTypeName returns the CommunicationMessageTypeName field if non-nil, zero value otherwise.

### GetCommunicationMessageTypeNameOk

`func (o *CommunicationMappingModel) GetCommunicationMessageTypeNameOk() (*string, bool)`

GetCommunicationMessageTypeNameOk returns a tuple with the CommunicationMessageTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationMessageTypeName

`func (o *CommunicationMappingModel) SetCommunicationMessageTypeName(v string)`

SetCommunicationMessageTypeName sets CommunicationMessageTypeName field to given value.

### HasCommunicationMessageTypeName

`func (o *CommunicationMappingModel) HasCommunicationMessageTypeName() bool`

HasCommunicationMessageTypeName returns a boolean if a field has been set.

### GetNoTags

`func (o *CommunicationMappingModel) GetNoTags() bool`

GetNoTags returns the NoTags field if non-nil, zero value otherwise.

### GetNoTagsOk

`func (o *CommunicationMappingModel) GetNoTagsOk() (*bool, bool)`

GetNoTagsOk returns a tuple with the NoTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoTags

`func (o *CommunicationMappingModel) SetNoTags(v bool)`

SetNoTags sets NoTags field to given value.

### HasNoTags

`func (o *CommunicationMappingModel) HasNoTags() bool`

HasNoTags returns a boolean if a field has been set.

### GetTagNames

`func (o *CommunicationMappingModel) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *CommunicationMappingModel) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *CommunicationMappingModel) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.

### HasTagNames

`func (o *CommunicationMappingModel) HasTagNames() bool`

HasTagNames returns a boolean if a field has been set.

### GetTags

`func (o *CommunicationMappingModel) GetTags() []TagModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CommunicationMappingModel) GetTagsOk() (*[]TagModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CommunicationMappingModel) SetTags(v []TagModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CommunicationMappingModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNoZones

`func (o *CommunicationMappingModel) GetNoZones() bool`

GetNoZones returns the NoZones field if non-nil, zero value otherwise.

### GetNoZonesOk

`func (o *CommunicationMappingModel) GetNoZonesOk() (*bool, bool)`

GetNoZonesOk returns a tuple with the NoZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoZones

`func (o *CommunicationMappingModel) SetNoZones(v bool)`

SetNoZones sets NoZones field to given value.

### HasNoZones

`func (o *CommunicationMappingModel) HasNoZones() bool`

HasNoZones returns a boolean if a field has been set.

### GetZones

`func (o *CommunicationMappingModel) GetZones() []ZoneModel`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *CommunicationMappingModel) GetZonesOk() (*[]ZoneModel, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *CommunicationMappingModel) SetZones(v []ZoneModel)`

SetZones sets Zones field to given value.

### HasZones

`func (o *CommunicationMappingModel) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetZoneNames

`func (o *CommunicationMappingModel) GetZoneNames() []string`

GetZoneNames returns the ZoneNames field if non-nil, zero value otherwise.

### GetZoneNamesOk

`func (o *CommunicationMappingModel) GetZoneNamesOk() (*[]string, bool)`

GetZoneNamesOk returns a tuple with the ZoneNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneNames

`func (o *CommunicationMappingModel) SetZoneNames(v []string)`

SetZoneNames sets ZoneNames field to given value.

### HasZoneNames

`func (o *CommunicationMappingModel) HasZoneNames() bool`

HasZoneNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


