# QuestionnaireTemplateQuestionFiltersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **[]int32** | QuestionnaireTemplateQuestion id&#39;s | [optional] 
**SearchText** | Pointer to **string** | free search through text and numeric type columns | [optional] 
**UpdatedAtSince** | Pointer to **time.Time** | Show updated since | [optional] 
**UpdatedAtTill** | Pointer to **time.Time** | Show updated till | [optional] 
**QuestionType** | Pointer to **string** | Question Type | [optional] 

## Methods

### NewQuestionnaireTemplateQuestionFiltersModel

`func NewQuestionnaireTemplateQuestionFiltersModel() *QuestionnaireTemplateQuestionFiltersModel`

NewQuestionnaireTemplateQuestionFiltersModel instantiates a new QuestionnaireTemplateQuestionFiltersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionnaireTemplateQuestionFiltersModelWithDefaults

`func NewQuestionnaireTemplateQuestionFiltersModelWithDefaults() *QuestionnaireTemplateQuestionFiltersModel`

NewQuestionnaireTemplateQuestionFiltersModelWithDefaults instantiates a new QuestionnaireTemplateQuestionFiltersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuestionnaireTemplateQuestionFiltersModel) GetId() []int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuestionnaireTemplateQuestionFiltersModel) GetIdOk() (*[]int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuestionnaireTemplateQuestionFiltersModel) SetId(v []int32)`

SetId sets Id field to given value.

### HasId

`func (o *QuestionnaireTemplateQuestionFiltersModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSearchText

`func (o *QuestionnaireTemplateQuestionFiltersModel) GetSearchText() string`

GetSearchText returns the SearchText field if non-nil, zero value otherwise.

### GetSearchTextOk

`func (o *QuestionnaireTemplateQuestionFiltersModel) GetSearchTextOk() (*string, bool)`

GetSearchTextOk returns a tuple with the SearchText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchText

`func (o *QuestionnaireTemplateQuestionFiltersModel) SetSearchText(v string)`

SetSearchText sets SearchText field to given value.

### HasSearchText

`func (o *QuestionnaireTemplateQuestionFiltersModel) HasSearchText() bool`

HasSearchText returns a boolean if a field has been set.

### GetUpdatedAtSince

`func (o *QuestionnaireTemplateQuestionFiltersModel) GetUpdatedAtSince() time.Time`

GetUpdatedAtSince returns the UpdatedAtSince field if non-nil, zero value otherwise.

### GetUpdatedAtSinceOk

`func (o *QuestionnaireTemplateQuestionFiltersModel) GetUpdatedAtSinceOk() (*time.Time, bool)`

GetUpdatedAtSinceOk returns a tuple with the UpdatedAtSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtSince

`func (o *QuestionnaireTemplateQuestionFiltersModel) SetUpdatedAtSince(v time.Time)`

SetUpdatedAtSince sets UpdatedAtSince field to given value.

### HasUpdatedAtSince

`func (o *QuestionnaireTemplateQuestionFiltersModel) HasUpdatedAtSince() bool`

HasUpdatedAtSince returns a boolean if a field has been set.

### GetUpdatedAtTill

`func (o *QuestionnaireTemplateQuestionFiltersModel) GetUpdatedAtTill() time.Time`

GetUpdatedAtTill returns the UpdatedAtTill field if non-nil, zero value otherwise.

### GetUpdatedAtTillOk

`func (o *QuestionnaireTemplateQuestionFiltersModel) GetUpdatedAtTillOk() (*time.Time, bool)`

GetUpdatedAtTillOk returns a tuple with the UpdatedAtTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtTill

`func (o *QuestionnaireTemplateQuestionFiltersModel) SetUpdatedAtTill(v time.Time)`

SetUpdatedAtTill sets UpdatedAtTill field to given value.

### HasUpdatedAtTill

`func (o *QuestionnaireTemplateQuestionFiltersModel) HasUpdatedAtTill() bool`

HasUpdatedAtTill returns a boolean if a field has been set.

### GetQuestionType

`func (o *QuestionnaireTemplateQuestionFiltersModel) GetQuestionType() string`

GetQuestionType returns the QuestionType field if non-nil, zero value otherwise.

### GetQuestionTypeOk

`func (o *QuestionnaireTemplateQuestionFiltersModel) GetQuestionTypeOk() (*string, bool)`

GetQuestionTypeOk returns a tuple with the QuestionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionType

`func (o *QuestionnaireTemplateQuestionFiltersModel) SetQuestionType(v string)`

SetQuestionType sets QuestionType field to given value.

### HasQuestionType

`func (o *QuestionnaireTemplateQuestionFiltersModel) HasQuestionType() bool`

HasQuestionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


