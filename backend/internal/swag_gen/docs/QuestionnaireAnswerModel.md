# QuestionnaireAnswerModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique Identifier | [optional] 
**Active** | Pointer to **bool** | if active&#x3D;0: QuestionnaireAnswer has been removed and is no longer visible in any bumbal interface | [optional] 
**QuestionnaireId** | Pointer to **int64** | Questionnaire ID | [optional] 
**QuestionnaireQuestionId** | Pointer to **int64** | Questionnaire Question ID | [optional] 
**Order** | Pointer to **int64** | Order | [optional] 
**Question** | Pointer to **string** | Textuale representation of the question | [optional] 
**Answer** | Pointer to **string** | Textuale representation of the answer | [optional] 
**ChosenOptions** | Pointer to **[]int32** | Chosen options id&#39;s | [optional] 
**AnswerOptions** | Pointer to [**[]QuestionnaireQuestionOptionModel**](QuestionnaireQuestionOptionModel.md) |  | [optional] [readonly] 
**QuestionnaireQuestionTypeName** | Pointer to **string** | Question type name | [optional] 

## Methods

### NewQuestionnaireAnswerModel

`func NewQuestionnaireAnswerModel() *QuestionnaireAnswerModel`

NewQuestionnaireAnswerModel instantiates a new QuestionnaireAnswerModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionnaireAnswerModelWithDefaults

`func NewQuestionnaireAnswerModelWithDefaults() *QuestionnaireAnswerModel`

NewQuestionnaireAnswerModelWithDefaults instantiates a new QuestionnaireAnswerModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuestionnaireAnswerModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuestionnaireAnswerModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuestionnaireAnswerModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *QuestionnaireAnswerModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActive

`func (o *QuestionnaireAnswerModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *QuestionnaireAnswerModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *QuestionnaireAnswerModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *QuestionnaireAnswerModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetQuestionnaireId

`func (o *QuestionnaireAnswerModel) GetQuestionnaireId() int64`

GetQuestionnaireId returns the QuestionnaireId field if non-nil, zero value otherwise.

### GetQuestionnaireIdOk

`func (o *QuestionnaireAnswerModel) GetQuestionnaireIdOk() (*int64, bool)`

GetQuestionnaireIdOk returns a tuple with the QuestionnaireId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionnaireId

`func (o *QuestionnaireAnswerModel) SetQuestionnaireId(v int64)`

SetQuestionnaireId sets QuestionnaireId field to given value.

### HasQuestionnaireId

`func (o *QuestionnaireAnswerModel) HasQuestionnaireId() bool`

HasQuestionnaireId returns a boolean if a field has been set.

### GetQuestionnaireQuestionId

`func (o *QuestionnaireAnswerModel) GetQuestionnaireQuestionId() int64`

GetQuestionnaireQuestionId returns the QuestionnaireQuestionId field if non-nil, zero value otherwise.

### GetQuestionnaireQuestionIdOk

`func (o *QuestionnaireAnswerModel) GetQuestionnaireQuestionIdOk() (*int64, bool)`

GetQuestionnaireQuestionIdOk returns a tuple with the QuestionnaireQuestionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionnaireQuestionId

`func (o *QuestionnaireAnswerModel) SetQuestionnaireQuestionId(v int64)`

SetQuestionnaireQuestionId sets QuestionnaireQuestionId field to given value.

### HasQuestionnaireQuestionId

`func (o *QuestionnaireAnswerModel) HasQuestionnaireQuestionId() bool`

HasQuestionnaireQuestionId returns a boolean if a field has been set.

### GetOrder

`func (o *QuestionnaireAnswerModel) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *QuestionnaireAnswerModel) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *QuestionnaireAnswerModel) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *QuestionnaireAnswerModel) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetQuestion

`func (o *QuestionnaireAnswerModel) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *QuestionnaireAnswerModel) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *QuestionnaireAnswerModel) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *QuestionnaireAnswerModel) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetAnswer

`func (o *QuestionnaireAnswerModel) GetAnswer() string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *QuestionnaireAnswerModel) GetAnswerOk() (*string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *QuestionnaireAnswerModel) SetAnswer(v string)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *QuestionnaireAnswerModel) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### GetChosenOptions

`func (o *QuestionnaireAnswerModel) GetChosenOptions() []int32`

GetChosenOptions returns the ChosenOptions field if non-nil, zero value otherwise.

### GetChosenOptionsOk

`func (o *QuestionnaireAnswerModel) GetChosenOptionsOk() (*[]int32, bool)`

GetChosenOptionsOk returns a tuple with the ChosenOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChosenOptions

`func (o *QuestionnaireAnswerModel) SetChosenOptions(v []int32)`

SetChosenOptions sets ChosenOptions field to given value.

### HasChosenOptions

`func (o *QuestionnaireAnswerModel) HasChosenOptions() bool`

HasChosenOptions returns a boolean if a field has been set.

### GetAnswerOptions

`func (o *QuestionnaireAnswerModel) GetAnswerOptions() []QuestionnaireQuestionOptionModel`

GetAnswerOptions returns the AnswerOptions field if non-nil, zero value otherwise.

### GetAnswerOptionsOk

`func (o *QuestionnaireAnswerModel) GetAnswerOptionsOk() (*[]QuestionnaireQuestionOptionModel, bool)`

GetAnswerOptionsOk returns a tuple with the AnswerOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerOptions

`func (o *QuestionnaireAnswerModel) SetAnswerOptions(v []QuestionnaireQuestionOptionModel)`

SetAnswerOptions sets AnswerOptions field to given value.

### HasAnswerOptions

`func (o *QuestionnaireAnswerModel) HasAnswerOptions() bool`

HasAnswerOptions returns a boolean if a field has been set.

### GetQuestionnaireQuestionTypeName

`func (o *QuestionnaireAnswerModel) GetQuestionnaireQuestionTypeName() string`

GetQuestionnaireQuestionTypeName returns the QuestionnaireQuestionTypeName field if non-nil, zero value otherwise.

### GetQuestionnaireQuestionTypeNameOk

`func (o *QuestionnaireAnswerModel) GetQuestionnaireQuestionTypeNameOk() (*string, bool)`

GetQuestionnaireQuestionTypeNameOk returns a tuple with the QuestionnaireQuestionTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionnaireQuestionTypeName

`func (o *QuestionnaireAnswerModel) SetQuestionnaireQuestionTypeName(v string)`

SetQuestionnaireQuestionTypeName sets QuestionnaireQuestionTypeName field to given value.

### HasQuestionnaireQuestionTypeName

`func (o *QuestionnaireAnswerModel) HasQuestionnaireQuestionTypeName() bool`

HasQuestionnaireQuestionTypeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


