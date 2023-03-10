# QuestionnaireQuestionTypeRetrieveListArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | Pointer to [**QuestionnaireQuestionTypeOptionsModel**](QuestionnaireQuestionTypeOptionsModel.md) |  | [optional] 
**Filters** | Pointer to [**QuestionnaireQuestionTypeFiltersModel**](QuestionnaireQuestionTypeFiltersModel.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**SortingColumn** | Pointer to **string** | Sorting Column | [optional] 
**SortingDirection** | Pointer to **string** | Sorting Direction | [optional] 
**SearchText** | Pointer to **string** |  | [optional] 

## Methods

### NewQuestionnaireQuestionTypeRetrieveListArguments

`func NewQuestionnaireQuestionTypeRetrieveListArguments() *QuestionnaireQuestionTypeRetrieveListArguments`

NewQuestionnaireQuestionTypeRetrieveListArguments instantiates a new QuestionnaireQuestionTypeRetrieveListArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionnaireQuestionTypeRetrieveListArgumentsWithDefaults

`func NewQuestionnaireQuestionTypeRetrieveListArgumentsWithDefaults() *QuestionnaireQuestionTypeRetrieveListArguments`

NewQuestionnaireQuestionTypeRetrieveListArgumentsWithDefaults instantiates a new QuestionnaireQuestionTypeRetrieveListArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *QuestionnaireQuestionTypeRetrieveListArguments) GetOptions() QuestionnaireQuestionTypeOptionsModel`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *QuestionnaireQuestionTypeRetrieveListArguments) GetOptionsOk() (*QuestionnaireQuestionTypeOptionsModel, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *QuestionnaireQuestionTypeRetrieveListArguments) SetOptions(v QuestionnaireQuestionTypeOptionsModel)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *QuestionnaireQuestionTypeRetrieveListArguments) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetFilters

`func (o *QuestionnaireQuestionTypeRetrieveListArguments) GetFilters() QuestionnaireQuestionTypeFiltersModel`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *QuestionnaireQuestionTypeRetrieveListArguments) GetFiltersOk() (*QuestionnaireQuestionTypeFiltersModel, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *QuestionnaireQuestionTypeRetrieveListArguments) SetFilters(v QuestionnaireQuestionTypeFiltersModel)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *QuestionnaireQuestionTypeRetrieveListArguments) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetLimit

`func (o *QuestionnaireQuestionTypeRetrieveListArguments) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *QuestionnaireQuestionTypeRetrieveListArguments) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *QuestionnaireQuestionTypeRetrieveListArguments) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *QuestionnaireQuestionTypeRetrieveListArguments) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *QuestionnaireQuestionTypeRetrieveListArguments) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *QuestionnaireQuestionTypeRetrieveListArguments) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *QuestionnaireQuestionTypeRetrieveListArguments) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *QuestionnaireQuestionTypeRetrieveListArguments) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetSortingColumn

`func (o *QuestionnaireQuestionTypeRetrieveListArguments) GetSortingColumn() string`

GetSortingColumn returns the SortingColumn field if non-nil, zero value otherwise.

### GetSortingColumnOk

`func (o *QuestionnaireQuestionTypeRetrieveListArguments) GetSortingColumnOk() (*string, bool)`

GetSortingColumnOk returns a tuple with the SortingColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortingColumn

`func (o *QuestionnaireQuestionTypeRetrieveListArguments) SetSortingColumn(v string)`

SetSortingColumn sets SortingColumn field to given value.

### HasSortingColumn

`func (o *QuestionnaireQuestionTypeRetrieveListArguments) HasSortingColumn() bool`

HasSortingColumn returns a boolean if a field has been set.

### GetSortingDirection

`func (o *QuestionnaireQuestionTypeRetrieveListArguments) GetSortingDirection() string`

GetSortingDirection returns the SortingDirection field if non-nil, zero value otherwise.

### GetSortingDirectionOk

`func (o *QuestionnaireQuestionTypeRetrieveListArguments) GetSortingDirectionOk() (*string, bool)`

GetSortingDirectionOk returns a tuple with the SortingDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortingDirection

`func (o *QuestionnaireQuestionTypeRetrieveListArguments) SetSortingDirection(v string)`

SetSortingDirection sets SortingDirection field to given value.

### HasSortingDirection

`func (o *QuestionnaireQuestionTypeRetrieveListArguments) HasSortingDirection() bool`

HasSortingDirection returns a boolean if a field has been set.

### GetSearchText

`func (o *QuestionnaireQuestionTypeRetrieveListArguments) GetSearchText() string`

GetSearchText returns the SearchText field if non-nil, zero value otherwise.

### GetSearchTextOk

`func (o *QuestionnaireQuestionTypeRetrieveListArguments) GetSearchTextOk() (*string, bool)`

GetSearchTextOk returns a tuple with the SearchText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchText

`func (o *QuestionnaireQuestionTypeRetrieveListArguments) SetSearchText(v string)`

SetSearchText sets SearchText field to given value.

### HasSearchText

`func (o *QuestionnaireQuestionTypeRetrieveListArguments) HasSearchText() bool`

HasSearchText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

