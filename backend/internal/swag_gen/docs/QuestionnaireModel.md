# QuestionnaireModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique Identifier | [optional] 
**Active** | Pointer to **bool** | if active&#x3D;0: Questionnaire has been removed and is no longer visible in any bumbal interface | [optional] 
**QuestionnaireTemplateId** | Pointer to **int64** | ID of the parent questionnaire template, just for reference | [optional] 
**Name** | Pointer to **string** | Questionnaire name | [optional] 
**QuestionnaireTypeId** | Pointer to **int64** | ID of the questionnaire type, just for reference | [optional] 
**QuestionnaireTypeName** | Pointer to **string** | name of the questionnaire type | [optional] 
**LangCode** | Pointer to **string** | ISO lang code | [optional] 
**Answers** | Pointer to [**[]QuestionnaireAnswerModel**](QuestionnaireAnswerModel.md) |  | [optional] 
**ObjectType** | Pointer to **map[string]interface{}** | Object type IDs available for this questionnaire | [optional] 
**ObjectTypeName** | Pointer to **string** | Object type name for the bound object to this questionnaire | [optional] 
**ObjectId** | Pointer to **int32** | Object ID | [optional] 

## Methods

### NewQuestionnaireModel

`func NewQuestionnaireModel() *QuestionnaireModel`

NewQuestionnaireModel instantiates a new QuestionnaireModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionnaireModelWithDefaults

`func NewQuestionnaireModelWithDefaults() *QuestionnaireModel`

NewQuestionnaireModelWithDefaults instantiates a new QuestionnaireModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuestionnaireModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuestionnaireModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuestionnaireModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *QuestionnaireModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActive

`func (o *QuestionnaireModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *QuestionnaireModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *QuestionnaireModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *QuestionnaireModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetQuestionnaireTemplateId

`func (o *QuestionnaireModel) GetQuestionnaireTemplateId() int64`

GetQuestionnaireTemplateId returns the QuestionnaireTemplateId field if non-nil, zero value otherwise.

### GetQuestionnaireTemplateIdOk

`func (o *QuestionnaireModel) GetQuestionnaireTemplateIdOk() (*int64, bool)`

GetQuestionnaireTemplateIdOk returns a tuple with the QuestionnaireTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionnaireTemplateId

`func (o *QuestionnaireModel) SetQuestionnaireTemplateId(v int64)`

SetQuestionnaireTemplateId sets QuestionnaireTemplateId field to given value.

### HasQuestionnaireTemplateId

`func (o *QuestionnaireModel) HasQuestionnaireTemplateId() bool`

HasQuestionnaireTemplateId returns a boolean if a field has been set.

### GetName

`func (o *QuestionnaireModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuestionnaireModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuestionnaireModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QuestionnaireModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuestionnaireTypeId

`func (o *QuestionnaireModel) GetQuestionnaireTypeId() int64`

GetQuestionnaireTypeId returns the QuestionnaireTypeId field if non-nil, zero value otherwise.

### GetQuestionnaireTypeIdOk

`func (o *QuestionnaireModel) GetQuestionnaireTypeIdOk() (*int64, bool)`

GetQuestionnaireTypeIdOk returns a tuple with the QuestionnaireTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionnaireTypeId

`func (o *QuestionnaireModel) SetQuestionnaireTypeId(v int64)`

SetQuestionnaireTypeId sets QuestionnaireTypeId field to given value.

### HasQuestionnaireTypeId

`func (o *QuestionnaireModel) HasQuestionnaireTypeId() bool`

HasQuestionnaireTypeId returns a boolean if a field has been set.

### GetQuestionnaireTypeName

`func (o *QuestionnaireModel) GetQuestionnaireTypeName() string`

GetQuestionnaireTypeName returns the QuestionnaireTypeName field if non-nil, zero value otherwise.

### GetQuestionnaireTypeNameOk

`func (o *QuestionnaireModel) GetQuestionnaireTypeNameOk() (*string, bool)`

GetQuestionnaireTypeNameOk returns a tuple with the QuestionnaireTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionnaireTypeName

`func (o *QuestionnaireModel) SetQuestionnaireTypeName(v string)`

SetQuestionnaireTypeName sets QuestionnaireTypeName field to given value.

### HasQuestionnaireTypeName

`func (o *QuestionnaireModel) HasQuestionnaireTypeName() bool`

HasQuestionnaireTypeName returns a boolean if a field has been set.

### GetLangCode

`func (o *QuestionnaireModel) GetLangCode() string`

GetLangCode returns the LangCode field if non-nil, zero value otherwise.

### GetLangCodeOk

`func (o *QuestionnaireModel) GetLangCodeOk() (*string, bool)`

GetLangCodeOk returns a tuple with the LangCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLangCode

`func (o *QuestionnaireModel) SetLangCode(v string)`

SetLangCode sets LangCode field to given value.

### HasLangCode

`func (o *QuestionnaireModel) HasLangCode() bool`

HasLangCode returns a boolean if a field has been set.

### GetAnswers

`func (o *QuestionnaireModel) GetAnswers() []QuestionnaireAnswerModel`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *QuestionnaireModel) GetAnswersOk() (*[]QuestionnaireAnswerModel, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *QuestionnaireModel) SetAnswers(v []QuestionnaireAnswerModel)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *QuestionnaireModel) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.

### GetObjectType

`func (o *QuestionnaireModel) GetObjectType() map[string]interface{}`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QuestionnaireModel) GetObjectTypeOk() (*map[string]interface{}, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QuestionnaireModel) SetObjectType(v map[string]interface{})`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *QuestionnaireModel) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObjectTypeName

`func (o *QuestionnaireModel) GetObjectTypeName() string`

GetObjectTypeName returns the ObjectTypeName field if non-nil, zero value otherwise.

### GetObjectTypeNameOk

`func (o *QuestionnaireModel) GetObjectTypeNameOk() (*string, bool)`

GetObjectTypeNameOk returns a tuple with the ObjectTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeName

`func (o *QuestionnaireModel) SetObjectTypeName(v string)`

SetObjectTypeName sets ObjectTypeName field to given value.

### HasObjectTypeName

`func (o *QuestionnaireModel) HasObjectTypeName() bool`

HasObjectTypeName returns a boolean if a field has been set.

### GetObjectId

`func (o *QuestionnaireModel) GetObjectId() int32`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *QuestionnaireModel) GetObjectIdOk() (*int32, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *QuestionnaireModel) SetObjectId(v int32)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *QuestionnaireModel) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


