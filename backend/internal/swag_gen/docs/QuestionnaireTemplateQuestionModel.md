# QuestionnaireTemplateQuestionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique Identifier | [optional] 
**Active** | Pointer to **bool** | if active&#x3D;0: note has been removed and is no longer visible in any bumbal interface | [optional] 
**QuestionnaireTemplateId** | Pointer to **int64** | Questionnaire template id | [optional] 
**QuestionTypeId** | Pointer to **int64** | Question type id | [optional] 
**QuestionTypeName** | Pointer to **string** | Question type name | [optional] 
**Order** | Pointer to **int64** | Order | [optional] 
**Options** | Pointer to [**[]QuestionnaireTemplateQuestionOptionModel**](QuestionnaireTemplateQuestionOptionModel.md) |  | [optional] 
**Texts** | Pointer to [**[]QuestionnaireTemplateQuestionTextModel**](QuestionnaireTemplateQuestionTextModel.md) |  | [optional] 

## Methods

### NewQuestionnaireTemplateQuestionModel

`func NewQuestionnaireTemplateQuestionModel() *QuestionnaireTemplateQuestionModel`

NewQuestionnaireTemplateQuestionModel instantiates a new QuestionnaireTemplateQuestionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionnaireTemplateQuestionModelWithDefaults

`func NewQuestionnaireTemplateQuestionModelWithDefaults() *QuestionnaireTemplateQuestionModel`

NewQuestionnaireTemplateQuestionModelWithDefaults instantiates a new QuestionnaireTemplateQuestionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuestionnaireTemplateQuestionModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuestionnaireTemplateQuestionModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuestionnaireTemplateQuestionModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *QuestionnaireTemplateQuestionModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActive

`func (o *QuestionnaireTemplateQuestionModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *QuestionnaireTemplateQuestionModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *QuestionnaireTemplateQuestionModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *QuestionnaireTemplateQuestionModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetQuestionnaireTemplateId

`func (o *QuestionnaireTemplateQuestionModel) GetQuestionnaireTemplateId() int64`

GetQuestionnaireTemplateId returns the QuestionnaireTemplateId field if non-nil, zero value otherwise.

### GetQuestionnaireTemplateIdOk

`func (o *QuestionnaireTemplateQuestionModel) GetQuestionnaireTemplateIdOk() (*int64, bool)`

GetQuestionnaireTemplateIdOk returns a tuple with the QuestionnaireTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionnaireTemplateId

`func (o *QuestionnaireTemplateQuestionModel) SetQuestionnaireTemplateId(v int64)`

SetQuestionnaireTemplateId sets QuestionnaireTemplateId field to given value.

### HasQuestionnaireTemplateId

`func (o *QuestionnaireTemplateQuestionModel) HasQuestionnaireTemplateId() bool`

HasQuestionnaireTemplateId returns a boolean if a field has been set.

### GetQuestionTypeId

`func (o *QuestionnaireTemplateQuestionModel) GetQuestionTypeId() int64`

GetQuestionTypeId returns the QuestionTypeId field if non-nil, zero value otherwise.

### GetQuestionTypeIdOk

`func (o *QuestionnaireTemplateQuestionModel) GetQuestionTypeIdOk() (*int64, bool)`

GetQuestionTypeIdOk returns a tuple with the QuestionTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionTypeId

`func (o *QuestionnaireTemplateQuestionModel) SetQuestionTypeId(v int64)`

SetQuestionTypeId sets QuestionTypeId field to given value.

### HasQuestionTypeId

`func (o *QuestionnaireTemplateQuestionModel) HasQuestionTypeId() bool`

HasQuestionTypeId returns a boolean if a field has been set.

### GetQuestionTypeName

`func (o *QuestionnaireTemplateQuestionModel) GetQuestionTypeName() string`

GetQuestionTypeName returns the QuestionTypeName field if non-nil, zero value otherwise.

### GetQuestionTypeNameOk

`func (o *QuestionnaireTemplateQuestionModel) GetQuestionTypeNameOk() (*string, bool)`

GetQuestionTypeNameOk returns a tuple with the QuestionTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionTypeName

`func (o *QuestionnaireTemplateQuestionModel) SetQuestionTypeName(v string)`

SetQuestionTypeName sets QuestionTypeName field to given value.

### HasQuestionTypeName

`func (o *QuestionnaireTemplateQuestionModel) HasQuestionTypeName() bool`

HasQuestionTypeName returns a boolean if a field has been set.

### GetOrder

`func (o *QuestionnaireTemplateQuestionModel) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *QuestionnaireTemplateQuestionModel) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *QuestionnaireTemplateQuestionModel) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *QuestionnaireTemplateQuestionModel) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetOptions

`func (o *QuestionnaireTemplateQuestionModel) GetOptions() []QuestionnaireTemplateQuestionOptionModel`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *QuestionnaireTemplateQuestionModel) GetOptionsOk() (*[]QuestionnaireTemplateQuestionOptionModel, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *QuestionnaireTemplateQuestionModel) SetOptions(v []QuestionnaireTemplateQuestionOptionModel)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *QuestionnaireTemplateQuestionModel) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetTexts

`func (o *QuestionnaireTemplateQuestionModel) GetTexts() []QuestionnaireTemplateQuestionTextModel`

GetTexts returns the Texts field if non-nil, zero value otherwise.

### GetTextsOk

`func (o *QuestionnaireTemplateQuestionModel) GetTextsOk() (*[]QuestionnaireTemplateQuestionTextModel, bool)`

GetTextsOk returns a tuple with the Texts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTexts

`func (o *QuestionnaireTemplateQuestionModel) SetTexts(v []QuestionnaireTemplateQuestionTextModel)`

SetTexts sets Texts field to given value.

### HasTexts

`func (o *QuestionnaireTemplateQuestionModel) HasTexts() bool`

HasTexts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


