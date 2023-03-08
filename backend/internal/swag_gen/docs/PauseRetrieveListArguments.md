# PauseRetrieveListArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | Pointer to [**PauseOptionsModel**](PauseOptionsModel.md) |  | [optional] 
**Filters** | Pointer to [**PauseFiltersModel**](PauseFiltersModel.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**SortingColumn** | Pointer to **string** | Sorting Column | [optional] 
**SortingDirection** | Pointer to **string** | Sorting Direction | [optional] 
**SearchText** | Pointer to **string** |  | [optional] 
**AsList** | Pointer to **bool** |  | [optional] 

## Methods

### NewPauseRetrieveListArguments

`func NewPauseRetrieveListArguments() *PauseRetrieveListArguments`

NewPauseRetrieveListArguments instantiates a new PauseRetrieveListArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPauseRetrieveListArgumentsWithDefaults

`func NewPauseRetrieveListArgumentsWithDefaults() *PauseRetrieveListArguments`

NewPauseRetrieveListArgumentsWithDefaults instantiates a new PauseRetrieveListArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *PauseRetrieveListArguments) GetOptions() PauseOptionsModel`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PauseRetrieveListArguments) GetOptionsOk() (*PauseOptionsModel, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PauseRetrieveListArguments) SetOptions(v PauseOptionsModel)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *PauseRetrieveListArguments) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetFilters

`func (o *PauseRetrieveListArguments) GetFilters() PauseFiltersModel`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PauseRetrieveListArguments) GetFiltersOk() (*PauseFiltersModel, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PauseRetrieveListArguments) SetFilters(v PauseFiltersModel)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *PauseRetrieveListArguments) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetLimit

`func (o *PauseRetrieveListArguments) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PauseRetrieveListArguments) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PauseRetrieveListArguments) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PauseRetrieveListArguments) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *PauseRetrieveListArguments) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PauseRetrieveListArguments) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PauseRetrieveListArguments) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *PauseRetrieveListArguments) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetSortingColumn

`func (o *PauseRetrieveListArguments) GetSortingColumn() string`

GetSortingColumn returns the SortingColumn field if non-nil, zero value otherwise.

### GetSortingColumnOk

`func (o *PauseRetrieveListArguments) GetSortingColumnOk() (*string, bool)`

GetSortingColumnOk returns a tuple with the SortingColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortingColumn

`func (o *PauseRetrieveListArguments) SetSortingColumn(v string)`

SetSortingColumn sets SortingColumn field to given value.

### HasSortingColumn

`func (o *PauseRetrieveListArguments) HasSortingColumn() bool`

HasSortingColumn returns a boolean if a field has been set.

### GetSortingDirection

`func (o *PauseRetrieveListArguments) GetSortingDirection() string`

GetSortingDirection returns the SortingDirection field if non-nil, zero value otherwise.

### GetSortingDirectionOk

`func (o *PauseRetrieveListArguments) GetSortingDirectionOk() (*string, bool)`

GetSortingDirectionOk returns a tuple with the SortingDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortingDirection

`func (o *PauseRetrieveListArguments) SetSortingDirection(v string)`

SetSortingDirection sets SortingDirection field to given value.

### HasSortingDirection

`func (o *PauseRetrieveListArguments) HasSortingDirection() bool`

HasSortingDirection returns a boolean if a field has been set.

### GetSearchText

`func (o *PauseRetrieveListArguments) GetSearchText() string`

GetSearchText returns the SearchText field if non-nil, zero value otherwise.

### GetSearchTextOk

`func (o *PauseRetrieveListArguments) GetSearchTextOk() (*string, bool)`

GetSearchTextOk returns a tuple with the SearchText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchText

`func (o *PauseRetrieveListArguments) SetSearchText(v string)`

SetSearchText sets SearchText field to given value.

### HasSearchText

`func (o *PauseRetrieveListArguments) HasSearchText() bool`

HasSearchText returns a boolean if a field has been set.

### GetAsList

`func (o *PauseRetrieveListArguments) GetAsList() bool`

GetAsList returns the AsList field if non-nil, zero value otherwise.

### GetAsListOk

`func (o *PauseRetrieveListArguments) GetAsListOk() (*bool, bool)`

GetAsListOk returns a tuple with the AsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsList

`func (o *PauseRetrieveListArguments) SetAsList(v bool)`

SetAsList sets AsList field to given value.

### HasAsList

`func (o *PauseRetrieveListArguments) HasAsList() bool`

HasAsList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


