# QuestionnaireQuestionModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuestionnaireId** | Pointer to **int32** | ID of the questionnaire id | [optional] 
**QuestionnaireQuestionId** | Pointer to **int32** | ID of the questionnaire question id | [optional] 
**Order** | Pointer to **int32** | order of the question inside the questionnaire | [optional] 
**Question** | Pointer to **string** | Question | [optional] 
**AnswerOptions** | Pointer to [**[]QuestionnaireQuestionOptionModel**](QuestionnaireQuestionOptionModel.md) |  | [optional] 
**Answer** | Pointer to **string** | Answer | [optional] 
**ChosenOptions** | Pointer to **[]int32** | Chosen option as answer | [optional] 
**QuestionnaireQuestionTypeName** | Pointer to **string** | Question type name | [optional] 

## Methods

### NewQuestionnaireQuestionModel

`func NewQuestionnaireQuestionModel() *QuestionnaireQuestionModel`

NewQuestionnaireQuestionModel instantiates a new QuestionnaireQuestionModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionnaireQuestionModelWithDefaults

`func NewQuestionnaireQuestionModelWithDefaults() *QuestionnaireQuestionModel`

NewQuestionnaireQuestionModelWithDefaults instantiates a new QuestionnaireQuestionModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestionnaireId

`func (o *QuestionnaireQuestionModel) GetQuestionnaireId() int32`

GetQuestionnaireId returns the QuestionnaireId field if non-nil, zero value otherwise.

### GetQuestionnaireIdOk

`func (o *QuestionnaireQuestionModel) GetQuestionnaireIdOk() (*int32, bool)`

GetQuestionnaireIdOk returns a tuple with the QuestionnaireId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionnaireId

`func (o *QuestionnaireQuestionModel) SetQuestionnaireId(v int32)`

SetQuestionnaireId sets QuestionnaireId field to given value.

### HasQuestionnaireId

`func (o *QuestionnaireQuestionModel) HasQuestionnaireId() bool`

HasQuestionnaireId returns a boolean if a field has been set.

### GetQuestionnaireQuestionId

`func (o *QuestionnaireQuestionModel) GetQuestionnaireQuestionId() int32`

GetQuestionnaireQuestionId returns the QuestionnaireQuestionId field if non-nil, zero value otherwise.

### GetQuestionnaireQuestionIdOk

`func (o *QuestionnaireQuestionModel) GetQuestionnaireQuestionIdOk() (*int32, bool)`

GetQuestionnaireQuestionIdOk returns a tuple with the QuestionnaireQuestionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionnaireQuestionId

`func (o *QuestionnaireQuestionModel) SetQuestionnaireQuestionId(v int32)`

SetQuestionnaireQuestionId sets QuestionnaireQuestionId field to given value.

### HasQuestionnaireQuestionId

`func (o *QuestionnaireQuestionModel) HasQuestionnaireQuestionId() bool`

HasQuestionnaireQuestionId returns a boolean if a field has been set.

### GetOrder

`func (o *QuestionnaireQuestionModel) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *QuestionnaireQuestionModel) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *QuestionnaireQuestionModel) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *QuestionnaireQuestionModel) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetQuestion

`func (o *QuestionnaireQuestionModel) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *QuestionnaireQuestionModel) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *QuestionnaireQuestionModel) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *QuestionnaireQuestionModel) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetAnswerOptions

`func (o *QuestionnaireQuestionModel) GetAnswerOptions() []QuestionnaireQuestionOptionModel`

GetAnswerOptions returns the AnswerOptions field if non-nil, zero value otherwise.

### GetAnswerOptionsOk

`func (o *QuestionnaireQuestionModel) GetAnswerOptionsOk() (*[]QuestionnaireQuestionOptionModel, bool)`

GetAnswerOptionsOk returns a tuple with the AnswerOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerOptions

`func (o *QuestionnaireQuestionModel) SetAnswerOptions(v []QuestionnaireQuestionOptionModel)`

SetAnswerOptions sets AnswerOptions field to given value.

### HasAnswerOptions

`func (o *QuestionnaireQuestionModel) HasAnswerOptions() bool`

HasAnswerOptions returns a boolean if a field has been set.

### GetAnswer

`func (o *QuestionnaireQuestionModel) GetAnswer() string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *QuestionnaireQuestionModel) GetAnswerOk() (*string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *QuestionnaireQuestionModel) SetAnswer(v string)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *QuestionnaireQuestionModel) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### GetChosenOptions

`func (o *QuestionnaireQuestionModel) GetChosenOptions() []int32`

GetChosenOptions returns the ChosenOptions field if non-nil, zero value otherwise.

### GetChosenOptionsOk

`func (o *QuestionnaireQuestionModel) GetChosenOptionsOk() (*[]int32, bool)`

GetChosenOptionsOk returns a tuple with the ChosenOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChosenOptions

`func (o *QuestionnaireQuestionModel) SetChosenOptions(v []int32)`

SetChosenOptions sets ChosenOptions field to given value.

### HasChosenOptions

`func (o *QuestionnaireQuestionModel) HasChosenOptions() bool`

HasChosenOptions returns a boolean if a field has been set.

### GetQuestionnaireQuestionTypeName

`func (o *QuestionnaireQuestionModel) GetQuestionnaireQuestionTypeName() string`

GetQuestionnaireQuestionTypeName returns the QuestionnaireQuestionTypeName field if non-nil, zero value otherwise.

### GetQuestionnaireQuestionTypeNameOk

`func (o *QuestionnaireQuestionModel) GetQuestionnaireQuestionTypeNameOk() (*string, bool)`

GetQuestionnaireQuestionTypeNameOk returns a tuple with the QuestionnaireQuestionTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionnaireQuestionTypeName

`func (o *QuestionnaireQuestionModel) SetQuestionnaireQuestionTypeName(v string)`

SetQuestionnaireQuestionTypeName sets QuestionnaireQuestionTypeName field to given value.

### HasQuestionnaireQuestionTypeName

`func (o *QuestionnaireQuestionModel) HasQuestionnaireQuestionTypeName() bool`

HasQuestionnaireQuestionTypeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


