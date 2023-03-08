# QuestionnaireTemplateQuestionTextRetrieveListArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | Pointer to **map[string]interface{}** |  | [optional] 
**Filters** | Pointer to [**QuestionnaireTemplateQuestionTextFiltersModel**](QuestionnaireTemplateQuestionTextFiltersModel.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**SortingColumn** | Pointer to **string** | Sorting Column | [optional] 
**SortingDirection** | Pointer to **string** | Sorting Direction | [optional] 
**SearchText** | Pointer to **string** |  | [optional] 

## Methods

### NewQuestionnaireTemplateQuestionTextRetrieveListArguments

`func NewQuestionnaireTemplateQuestionTextRetrieveListArguments() *QuestionnaireTemplateQuestionTextRetrieveListArguments`

NewQuestionnaireTemplateQuestionTextRetrieveListArguments instantiates a new QuestionnaireTemplateQuestionTextRetrieveListArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionnaireTemplateQuestionTextRetrieveListArgumentsWithDefaults

`func NewQuestionnaireTemplateQuestionTextRetrieveListArgumentsWithDefaults() *QuestionnaireTemplateQuestionTextRetrieveListArguments`

NewQuestionnaireTemplateQuestionTextRetrieveListArgumentsWithDefaults instantiates a new QuestionnaireTemplateQuestionTextRetrieveListArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *QuestionnaireTemplateQuestionTextRetrieveListArguments) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *QuestionnaireTemplateQuestionTextRetrieveListArguments) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *QuestionnaireTemplateQuestionTextRetrieveListArguments) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *QuestionnaireTemplateQuestionTextRetrieveListArguments) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetFilters

`func (o *QuestionnaireTemplateQuestionTextRetrieveListArguments) GetFilters() QuestionnaireTemplateQuestionTextFiltersModel`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *QuestionnaireTemplateQuestionTextRetrieveListArguments) GetFiltersOk() (*QuestionnaireTemplateQuestionTextFiltersModel, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *QuestionnaireTemplateQuestionTextRetrieveListArguments) SetFilters(v QuestionnaireTemplateQuestionTextFiltersModel)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *QuestionnaireTemplateQuestionTextRetrieveListArguments) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetLimit

`func (o *QuestionnaireTemplateQuestionTextRetrieveListArguments) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *QuestionnaireTemplateQuestionTextRetrieveListArguments) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *QuestionnaireTemplateQuestionTextRetrieveListArguments) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *QuestionnaireTemplateQuestionTextRetrieveListArguments) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *QuestionnaireTemplateQuestionTextRetrieveListArguments) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *QuestionnaireTemplateQuestionTextRetrieveListArguments) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *QuestionnaireTemplateQuestionTextRetrieveListArguments) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *QuestionnaireTemplateQuestionTextRetrieveListArguments) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetSortingColumn

`func (o *QuestionnaireTemplateQuestionTextRetrieveListArguments) GetSortingColumn() string`

GetSortingColumn returns the SortingColumn field if non-nil, zero value otherwise.

### GetSortingColumnOk

`func (o *QuestionnaireTemplateQuestionTextRetrieveListArguments) GetSortingColumnOk() (*string, bool)`

GetSortingColumnOk returns a tuple with the SortingColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortingColumn

`func (o *QuestionnaireTemplateQuestionTextRetrieveListArguments) SetSortingColumn(v string)`

SetSortingColumn sets SortingColumn field to given value.

### HasSortingColumn

`func (o *QuestionnaireTemplateQuestionTextRetrieveListArguments) HasSortingColumn() bool`

HasSortingColumn returns a boolean if a field has been set.

### GetSortingDirection

`func (o *QuestionnaireTemplateQuestionTextRetrieveListArguments) GetSortingDirection() string`

GetSortingDirection returns the SortingDirection field if non-nil, zero value otherwise.

### GetSortingDirectionOk

`func (o *QuestionnaireTemplateQuestionTextRetrieveListArguments) GetSortingDirectionOk() (*string, bool)`

GetSortingDirectionOk returns a tuple with the SortingDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortingDirection

`func (o *QuestionnaireTemplateQuestionTextRetrieveListArguments) SetSortingDirection(v string)`

SetSortingDirection sets SortingDirection field to given value.

### HasSortingDirection

`func (o *QuestionnaireTemplateQuestionTextRetrieveListArguments) HasSortingDirection() bool`

HasSortingDirection returns a boolean if a field has been set.

### GetSearchText

`func (o *QuestionnaireTemplateQuestionTextRetrieveListArguments) GetSearchText() string`

GetSearchText returns the SearchText field if non-nil, zero value otherwise.

### GetSearchTextOk

`func (o *QuestionnaireTemplateQuestionTextRetrieveListArguments) GetSearchTextOk() (*string, bool)`

GetSearchTextOk returns a tuple with the SearchText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchText

`func (o *QuestionnaireTemplateQuestionTextRetrieveListArguments) SetSearchText(v string)`

SetSearchText sets SearchText field to given value.

### HasSearchText

`func (o *QuestionnaireTemplateQuestionTextRetrieveListArguments) HasSearchText() bool`

HasSearchText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


