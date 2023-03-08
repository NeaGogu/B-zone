# QuestionnaireTemplateQuestionOptionRetrieveListArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | Pointer to [**QuestionnaireTemplateQuestionOptionOptionsModel**](QuestionnaireTemplateQuestionOptionOptionsModel.md) |  | [optional] 
**Filters** | Pointer to [**QuestionnaireTemplateQuestionOptionFiltersModel**](QuestionnaireTemplateQuestionOptionFiltersModel.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**SortingColumn** | Pointer to **string** | Sorting Column | [optional] 
**SortingDirection** | Pointer to **string** | Sorting Direction | [optional] 
**SearchText** | Pointer to **string** |  | [optional] 

## Methods

### NewQuestionnaireTemplateQuestionOptionRetrieveListArguments

`func NewQuestionnaireTemplateQuestionOptionRetrieveListArguments() *QuestionnaireTemplateQuestionOptionRetrieveListArguments`

NewQuestionnaireTemplateQuestionOptionRetrieveListArguments instantiates a new QuestionnaireTemplateQuestionOptionRetrieveListArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionnaireTemplateQuestionOptionRetrieveListArgumentsWithDefaults

`func NewQuestionnaireTemplateQuestionOptionRetrieveListArgumentsWithDefaults() *QuestionnaireTemplateQuestionOptionRetrieveListArguments`

NewQuestionnaireTemplateQuestionOptionRetrieveListArgumentsWithDefaults instantiates a new QuestionnaireTemplateQuestionOptionRetrieveListArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *QuestionnaireTemplateQuestionOptionRetrieveListArguments) GetOptions() QuestionnaireTemplateQuestionOptionOptionsModel`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *QuestionnaireTemplateQuestionOptionRetrieveListArguments) GetOptionsOk() (*QuestionnaireTemplateQuestionOptionOptionsModel, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *QuestionnaireTemplateQuestionOptionRetrieveListArguments) SetOptions(v QuestionnaireTemplateQuestionOptionOptionsModel)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *QuestionnaireTemplateQuestionOptionRetrieveListArguments) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetFilters

`func (o *QuestionnaireTemplateQuestionOptionRetrieveListArguments) GetFilters() QuestionnaireTemplateQuestionOptionFiltersModel`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *QuestionnaireTemplateQuestionOptionRetrieveListArguments) GetFiltersOk() (*QuestionnaireTemplateQuestionOptionFiltersModel, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *QuestionnaireTemplateQuestionOptionRetrieveListArguments) SetFilters(v QuestionnaireTemplateQuestionOptionFiltersModel)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *QuestionnaireTemplateQuestionOptionRetrieveListArguments) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetLimit

`func (o *QuestionnaireTemplateQuestionOptionRetrieveListArguments) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *QuestionnaireTemplateQuestionOptionRetrieveListArguments) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *QuestionnaireTemplateQuestionOptionRetrieveListArguments) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *QuestionnaireTemplateQuestionOptionRetrieveListArguments) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *QuestionnaireTemplateQuestionOptionRetrieveListArguments) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *QuestionnaireTemplateQuestionOptionRetrieveListArguments) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *QuestionnaireTemplateQuestionOptionRetrieveListArguments) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *QuestionnaireTemplateQuestionOptionRetrieveListArguments) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetSortingColumn

`func (o *QuestionnaireTemplateQuestionOptionRetrieveListArguments) GetSortingColumn() string`

GetSortingColumn returns the SortingColumn field if non-nil, zero value otherwise.

### GetSortingColumnOk

`func (o *QuestionnaireTemplateQuestionOptionRetrieveListArguments) GetSortingColumnOk() (*string, bool)`

GetSortingColumnOk returns a tuple with the SortingColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortingColumn

`func (o *QuestionnaireTemplateQuestionOptionRetrieveListArguments) SetSortingColumn(v string)`

SetSortingColumn sets SortingColumn field to given value.

### HasSortingColumn

`func (o *QuestionnaireTemplateQuestionOptionRetrieveListArguments) HasSortingColumn() bool`

HasSortingColumn returns a boolean if a field has been set.

### GetSortingDirection

`func (o *QuestionnaireTemplateQuestionOptionRetrieveListArguments) GetSortingDirection() string`

GetSortingDirection returns the SortingDirection field if non-nil, zero value otherwise.

### GetSortingDirectionOk

`func (o *QuestionnaireTemplateQuestionOptionRetrieveListArguments) GetSortingDirectionOk() (*string, bool)`

GetSortingDirectionOk returns a tuple with the SortingDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortingDirection

`func (o *QuestionnaireTemplateQuestionOptionRetrieveListArguments) SetSortingDirection(v string)`

SetSortingDirection sets SortingDirection field to given value.

### HasSortingDirection

`func (o *QuestionnaireTemplateQuestionOptionRetrieveListArguments) HasSortingDirection() bool`

HasSortingDirection returns a boolean if a field has been set.

### GetSearchText

`func (o *QuestionnaireTemplateQuestionOptionRetrieveListArguments) GetSearchText() string`

GetSearchText returns the SearchText field if non-nil, zero value otherwise.

### GetSearchTextOk

`func (o *QuestionnaireTemplateQuestionOptionRetrieveListArguments) GetSearchTextOk() (*string, bool)`

GetSearchTextOk returns a tuple with the SearchText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchText

`func (o *QuestionnaireTemplateQuestionOptionRetrieveListArguments) SetSearchText(v string)`

SetSearchText sets SearchText field to given value.

### HasSearchText

`func (o *QuestionnaireTemplateQuestionOptionRetrieveListArguments) HasSearchText() bool`

HasSearchText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


