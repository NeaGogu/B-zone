# QuestionnaireAnswerFiltersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **[]int32** | QuestionnaireAnswer id&#39;s | [optional] 
**QuestionnaireId** | Pointer to **[]int32** | Questionnaire id&#39;s | [optional] 
**SearchText** | Pointer to **string** | free search through text and numeric type columns | [optional] 
**UpdatedAtSince** | Pointer to **time.Time** | Show updated since | [optional] 
**UpdatedAtTill** | Pointer to **time.Time** | Show updated till | [optional] 

## Methods

### NewQuestionnaireAnswerFiltersModel

`func NewQuestionnaireAnswerFiltersModel() *QuestionnaireAnswerFiltersModel`

NewQuestionnaireAnswerFiltersModel instantiates a new QuestionnaireAnswerFiltersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionnaireAnswerFiltersModelWithDefaults

`func NewQuestionnaireAnswerFiltersModelWithDefaults() *QuestionnaireAnswerFiltersModel`

NewQuestionnaireAnswerFiltersModelWithDefaults instantiates a new QuestionnaireAnswerFiltersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuestionnaireAnswerFiltersModel) GetId() []int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuestionnaireAnswerFiltersModel) GetIdOk() (*[]int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuestionnaireAnswerFiltersModel) SetId(v []int32)`

SetId sets Id field to given value.

### HasId

`func (o *QuestionnaireAnswerFiltersModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQuestionnaireId

`func (o *QuestionnaireAnswerFiltersModel) GetQuestionnaireId() []int32`

GetQuestionnaireId returns the QuestionnaireId field if non-nil, zero value otherwise.

### GetQuestionnaireIdOk

`func (o *QuestionnaireAnswerFiltersModel) GetQuestionnaireIdOk() (*[]int32, bool)`

GetQuestionnaireIdOk returns a tuple with the QuestionnaireId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionnaireId

`func (o *QuestionnaireAnswerFiltersModel) SetQuestionnaireId(v []int32)`

SetQuestionnaireId sets QuestionnaireId field to given value.

### HasQuestionnaireId

`func (o *QuestionnaireAnswerFiltersModel) HasQuestionnaireId() bool`

HasQuestionnaireId returns a boolean if a field has been set.

### GetSearchText

`func (o *QuestionnaireAnswerFiltersModel) GetSearchText() string`

GetSearchText returns the SearchText field if non-nil, zero value otherwise.

### GetSearchTextOk

`func (o *QuestionnaireAnswerFiltersModel) GetSearchTextOk() (*string, bool)`

GetSearchTextOk returns a tuple with the SearchText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchText

`func (o *QuestionnaireAnswerFiltersModel) SetSearchText(v string)`

SetSearchText sets SearchText field to given value.

### HasSearchText

`func (o *QuestionnaireAnswerFiltersModel) HasSearchText() bool`

HasSearchText returns a boolean if a field has been set.

### GetUpdatedAtSince

`func (o *QuestionnaireAnswerFiltersModel) GetUpdatedAtSince() time.Time`

GetUpdatedAtSince returns the UpdatedAtSince field if non-nil, zero value otherwise.

### GetUpdatedAtSinceOk

`func (o *QuestionnaireAnswerFiltersModel) GetUpdatedAtSinceOk() (*time.Time, bool)`

GetUpdatedAtSinceOk returns a tuple with the UpdatedAtSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtSince

`func (o *QuestionnaireAnswerFiltersModel) SetUpdatedAtSince(v time.Time)`

SetUpdatedAtSince sets UpdatedAtSince field to given value.

### HasUpdatedAtSince

`func (o *QuestionnaireAnswerFiltersModel) HasUpdatedAtSince() bool`

HasUpdatedAtSince returns a boolean if a field has been set.

### GetUpdatedAtTill

`func (o *QuestionnaireAnswerFiltersModel) GetUpdatedAtTill() time.Time`

GetUpdatedAtTill returns the UpdatedAtTill field if non-nil, zero value otherwise.

### GetUpdatedAtTillOk

`func (o *QuestionnaireAnswerFiltersModel) GetUpdatedAtTillOk() (*time.Time, bool)`

GetUpdatedAtTillOk returns a tuple with the UpdatedAtTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtTill

`func (o *QuestionnaireAnswerFiltersModel) SetUpdatedAtTill(v time.Time)`

SetUpdatedAtTill sets UpdatedAtTill field to given value.

### HasUpdatedAtTill

`func (o *QuestionnaireAnswerFiltersModel) HasUpdatedAtTill() bool`

HasUpdatedAtTill returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

