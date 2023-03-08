# QuestionnaireTemplateQuestionTextModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique Identifier | [optional] 
**Active** | Pointer to **bool** | if active&#x3D;0: note has been removed and is no longer visible in any bumbal interface | [optional] 
**QuestionnaireTemplateQuestionId** | Pointer to **int64** | Questionnaire Template Question ID | [optional] 
**LangCode** | Pointer to **string** | ISO lang code | [optional] 
**Question** | Pointer to **string** | Text representation of the question | [optional] 

## Methods

### NewQuestionnaireTemplateQuestionTextModel

`func NewQuestionnaireTemplateQuestionTextModel() *QuestionnaireTemplateQuestionTextModel`

NewQuestionnaireTemplateQuestionTextModel instantiates a new QuestionnaireTemplateQuestionTextModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionnaireTemplateQuestionTextModelWithDefaults

`func NewQuestionnaireTemplateQuestionTextModelWithDefaults() *QuestionnaireTemplateQuestionTextModel`

NewQuestionnaireTemplateQuestionTextModelWithDefaults instantiates a new QuestionnaireTemplateQuestionTextModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuestionnaireTemplateQuestionTextModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuestionnaireTemplateQuestionTextModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuestionnaireTemplateQuestionTextModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *QuestionnaireTemplateQuestionTextModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActive

`func (o *QuestionnaireTemplateQuestionTextModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *QuestionnaireTemplateQuestionTextModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *QuestionnaireTemplateQuestionTextModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *QuestionnaireTemplateQuestionTextModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetQuestionnaireTemplateQuestionId

`func (o *QuestionnaireTemplateQuestionTextModel) GetQuestionnaireTemplateQuestionId() int64`

GetQuestionnaireTemplateQuestionId returns the QuestionnaireTemplateQuestionId field if non-nil, zero value otherwise.

### GetQuestionnaireTemplateQuestionIdOk

`func (o *QuestionnaireTemplateQuestionTextModel) GetQuestionnaireTemplateQuestionIdOk() (*int64, bool)`

GetQuestionnaireTemplateQuestionIdOk returns a tuple with the QuestionnaireTemplateQuestionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionnaireTemplateQuestionId

`func (o *QuestionnaireTemplateQuestionTextModel) SetQuestionnaireTemplateQuestionId(v int64)`

SetQuestionnaireTemplateQuestionId sets QuestionnaireTemplateQuestionId field to given value.

### HasQuestionnaireTemplateQuestionId

`func (o *QuestionnaireTemplateQuestionTextModel) HasQuestionnaireTemplateQuestionId() bool`

HasQuestionnaireTemplateQuestionId returns a boolean if a field has been set.

### GetLangCode

`func (o *QuestionnaireTemplateQuestionTextModel) GetLangCode() string`

GetLangCode returns the LangCode field if non-nil, zero value otherwise.

### GetLangCodeOk

`func (o *QuestionnaireTemplateQuestionTextModel) GetLangCodeOk() (*string, bool)`

GetLangCodeOk returns a tuple with the LangCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLangCode

`func (o *QuestionnaireTemplateQuestionTextModel) SetLangCode(v string)`

SetLangCode sets LangCode field to given value.

### HasLangCode

`func (o *QuestionnaireTemplateQuestionTextModel) HasLangCode() bool`

HasLangCode returns a boolean if a field has been set.

### GetQuestion

`func (o *QuestionnaireTemplateQuestionTextModel) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *QuestionnaireTemplateQuestionTextModel) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *QuestionnaireTemplateQuestionTextModel) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *QuestionnaireTemplateQuestionTextModel) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


