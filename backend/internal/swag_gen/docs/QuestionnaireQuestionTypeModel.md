# QuestionnaireQuestionTypeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique Identifier | [optional] 
**Active** | Pointer to **bool** | if active&#x3D;0: note has been removed and is no longer visible in any bumbal interface | [optional] 
**Name** | Pointer to **string** | Name of the question type | [optional] 

## Methods

### NewQuestionnaireQuestionTypeModel

`func NewQuestionnaireQuestionTypeModel() *QuestionnaireQuestionTypeModel`

NewQuestionnaireQuestionTypeModel instantiates a new QuestionnaireQuestionTypeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionnaireQuestionTypeModelWithDefaults

`func NewQuestionnaireQuestionTypeModelWithDefaults() *QuestionnaireQuestionTypeModel`

NewQuestionnaireQuestionTypeModelWithDefaults instantiates a new QuestionnaireQuestionTypeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuestionnaireQuestionTypeModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuestionnaireQuestionTypeModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuestionnaireQuestionTypeModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *QuestionnaireQuestionTypeModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActive

`func (o *QuestionnaireQuestionTypeModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *QuestionnaireQuestionTypeModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *QuestionnaireQuestionTypeModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *QuestionnaireQuestionTypeModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetName

`func (o *QuestionnaireQuestionTypeModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuestionnaireQuestionTypeModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuestionnaireQuestionTypeModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QuestionnaireQuestionTypeModel) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


