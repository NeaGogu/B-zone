# QuestionnaireTemplateQuestionRetrieveListArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | Pointer to [**QuestionnaireTemplateQuestionOptionsModel**](QuestionnaireTemplateQuestionOptionsModel.md) |  | [optional] 
**Filters** | Pointer to [**QuestionnaireTemplateQuestionFiltersModel**](QuestionnaireTemplateQuestionFiltersModel.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**SortingColumn** | Pointer to **string** | Sorting Column | [optional] [default to "order"]
**SortingDirection** | Pointer to **string** | Sorting Direction | [optional] 
**SearchText** | Pointer to **string** |  | [optional] 

## Methods

### NewQuestionnaireTemplateQuestionRetrieveListArguments

`func NewQuestionnaireTemplateQuestionRetrieveListArguments() *QuestionnaireTemplateQuestionRetrieveListArguments`

NewQuestionnaireTemplateQuestionRetrieveListArguments instantiates a new QuestionnaireTemplateQuestionRetrieveListArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionnaireTemplateQuestionRetrieveListArgumentsWithDefaults

`func NewQuestionnaireTemplateQuestionRetrieveListArgumentsWithDefaults() *QuestionnaireTemplateQuestionRetrieveListArguments`

NewQuestionnaireTemplateQuestionRetrieveListArgumentsWithDefaults instantiates a new QuestionnaireTemplateQuestionRetrieveListArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *QuestionnaireTemplateQuestionRetrieveListArguments) GetOptions() QuestionnaireTemplateQuestionOptionsModel`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *QuestionnaireTemplateQuestionRetrieveListArguments) GetOptionsOk() (*QuestionnaireTemplateQuestionOptionsModel, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *QuestionnaireTemplateQuestionRetrieveListArguments) SetOptions(v QuestionnaireTemplateQuestionOptionsModel)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *QuestionnaireTemplateQuestionRetrieveListArguments) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetFilters

`func (o *QuestionnaireTemplateQuestionRetrieveListArguments) GetFilters() QuestionnaireTemplateQuestionFiltersModel`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *QuestionnaireTemplateQuestionRetrieveListArguments) GetFiltersOk() (*QuestionnaireTemplateQuestionFiltersModel, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *QuestionnaireTemplateQuestionRetrieveListArguments) SetFilters(v QuestionnaireTemplateQuestionFiltersModel)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *QuestionnaireTemplateQuestionRetrieveListArguments) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetLimit

`func (o *QuestionnaireTemplateQuestionRetrieveListArguments) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *QuestionnaireTemplateQuestionRetrieveListArguments) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *QuestionnaireTemplateQuestionRetrieveListArguments) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *QuestionnaireTemplateQuestionRetrieveListArguments) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *QuestionnaireTemplateQuestionRetrieveListArguments) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *QuestionnaireTemplateQuestionRetrieveListArguments) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *QuestionnaireTemplateQuestionRetrieveListArguments) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *QuestionnaireTemplateQuestionRetrieveListArguments) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetSortingColumn

`func (o *QuestionnaireTemplateQuestionRetrieveListArguments) GetSortingColumn() string`

GetSortingColumn returns the SortingColumn field if non-nil, zero value otherwise.

### GetSortingColumnOk

`func (o *QuestionnaireTemplateQuestionRetrieveListArguments) GetSortingColumnOk() (*string, bool)`

GetSortingColumnOk returns a tuple with the SortingColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortingColumn

`func (o *QuestionnaireTemplateQuestionRetrieveListArguments) SetSortingColumn(v string)`

SetSortingColumn sets SortingColumn field to given value.

### HasSortingColumn

`func (o *QuestionnaireTemplateQuestionRetrieveListArguments) HasSortingColumn() bool`

HasSortingColumn returns a boolean if a field has been set.

### GetSortingDirection

`func (o *QuestionnaireTemplateQuestionRetrieveListArguments) GetSortingDirection() string`

GetSortingDirection returns the SortingDirection field if non-nil, zero value otherwise.

### GetSortingDirectionOk

`func (o *QuestionnaireTemplateQuestionRetrieveListArguments) GetSortingDirectionOk() (*string, bool)`

GetSortingDirectionOk returns a tuple with the SortingDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortingDirection

`func (o *QuestionnaireTemplateQuestionRetrieveListArguments) SetSortingDirection(v string)`

SetSortingDirection sets SortingDirection field to given value.

### HasSortingDirection

`func (o *QuestionnaireTemplateQuestionRetrieveListArguments) HasSortingDirection() bool`

HasSortingDirection returns a boolean if a field has been set.

### GetSearchText

`func (o *QuestionnaireTemplateQuestionRetrieveListArguments) GetSearchText() string`

GetSearchText returns the SearchText field if non-nil, zero value otherwise.

### GetSearchTextOk

`func (o *QuestionnaireTemplateQuestionRetrieveListArguments) GetSearchTextOk() (*string, bool)`

GetSearchTextOk returns a tuple with the SearchText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchText

`func (o *QuestionnaireTemplateQuestionRetrieveListArguments) SetSearchText(v string)`

SetSearchText sets SearchText field to given value.

### HasSearchText

`func (o *QuestionnaireTemplateQuestionRetrieveListArguments) HasSearchText() bool`

HasSearchText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


