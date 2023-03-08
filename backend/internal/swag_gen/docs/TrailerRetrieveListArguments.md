# TrailerRetrieveListArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | Pointer to [**TrailerOptionsModel**](TrailerOptionsModel.md) |  | [optional] 
**Filters** | Pointer to [**TrailerFiltersModel**](TrailerFiltersModel.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**SortingColumn** | Pointer to **string** | Sorting Column | [optional] 
**SortingDirection** | Pointer to **string** | Sorting Direction | [optional] 
**SearchText** | Pointer to **string** |  | [optional] 
**AsList** | Pointer to **bool** |  | [optional] 
**CountOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewTrailerRetrieveListArguments

`func NewTrailerRetrieveListArguments() *TrailerRetrieveListArguments`

NewTrailerRetrieveListArguments instantiates a new TrailerRetrieveListArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrailerRetrieveListArgumentsWithDefaults

`func NewTrailerRetrieveListArgumentsWithDefaults() *TrailerRetrieveListArguments`

NewTrailerRetrieveListArgumentsWithDefaults instantiates a new TrailerRetrieveListArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *TrailerRetrieveListArguments) GetOptions() TrailerOptionsModel`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *TrailerRetrieveListArguments) GetOptionsOk() (*TrailerOptionsModel, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *TrailerRetrieveListArguments) SetOptions(v TrailerOptionsModel)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *TrailerRetrieveListArguments) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetFilters

`func (o *TrailerRetrieveListArguments) GetFilters() TrailerFiltersModel`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *TrailerRetrieveListArguments) GetFiltersOk() (*TrailerFiltersModel, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *TrailerRetrieveListArguments) SetFilters(v TrailerFiltersModel)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *TrailerRetrieveListArguments) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetLimit

`func (o *TrailerRetrieveListArguments) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TrailerRetrieveListArguments) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TrailerRetrieveListArguments) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *TrailerRetrieveListArguments) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *TrailerRetrieveListArguments) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TrailerRetrieveListArguments) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TrailerRetrieveListArguments) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *TrailerRetrieveListArguments) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetSortingColumn

`func (o *TrailerRetrieveListArguments) GetSortingColumn() string`

GetSortingColumn returns the SortingColumn field if non-nil, zero value otherwise.

### GetSortingColumnOk

`func (o *TrailerRetrieveListArguments) GetSortingColumnOk() (*string, bool)`

GetSortingColumnOk returns a tuple with the SortingColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortingColumn

`func (o *TrailerRetrieveListArguments) SetSortingColumn(v string)`

SetSortingColumn sets SortingColumn field to given value.

### HasSortingColumn

`func (o *TrailerRetrieveListArguments) HasSortingColumn() bool`

HasSortingColumn returns a boolean if a field has been set.

### GetSortingDirection

`func (o *TrailerRetrieveListArguments) GetSortingDirection() string`

GetSortingDirection returns the SortingDirection field if non-nil, zero value otherwise.

### GetSortingDirectionOk

`func (o *TrailerRetrieveListArguments) GetSortingDirectionOk() (*string, bool)`

GetSortingDirectionOk returns a tuple with the SortingDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortingDirection

`func (o *TrailerRetrieveListArguments) SetSortingDirection(v string)`

SetSortingDirection sets SortingDirection field to given value.

### HasSortingDirection

`func (o *TrailerRetrieveListArguments) HasSortingDirection() bool`

HasSortingDirection returns a boolean if a field has been set.

### GetSearchText

`func (o *TrailerRetrieveListArguments) GetSearchText() string`

GetSearchText returns the SearchText field if non-nil, zero value otherwise.

### GetSearchTextOk

`func (o *TrailerRetrieveListArguments) GetSearchTextOk() (*string, bool)`

GetSearchTextOk returns a tuple with the SearchText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchText

`func (o *TrailerRetrieveListArguments) SetSearchText(v string)`

SetSearchText sets SearchText field to given value.

### HasSearchText

`func (o *TrailerRetrieveListArguments) HasSearchText() bool`

HasSearchText returns a boolean if a field has been set.

### GetAsList

`func (o *TrailerRetrieveListArguments) GetAsList() bool`

GetAsList returns the AsList field if non-nil, zero value otherwise.

### GetAsListOk

`func (o *TrailerRetrieveListArguments) GetAsListOk() (*bool, bool)`

GetAsListOk returns a tuple with the AsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsList

`func (o *TrailerRetrieveListArguments) SetAsList(v bool)`

SetAsList sets AsList field to given value.

### HasAsList

`func (o *TrailerRetrieveListArguments) HasAsList() bool`

HasAsList returns a boolean if a field has been set.

### GetCountOnly

`func (o *TrailerRetrieveListArguments) GetCountOnly() bool`

GetCountOnly returns the CountOnly field if non-nil, zero value otherwise.

### GetCountOnlyOk

`func (o *TrailerRetrieveListArguments) GetCountOnlyOk() (*bool, bool)`

GetCountOnlyOk returns a tuple with the CountOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountOnly

`func (o *TrailerRetrieveListArguments) SetCountOnly(v bool)`

SetCountOnly sets CountOnly field to given value.

### HasCountOnly

`func (o *TrailerRetrieveListArguments) HasCountOnly() bool`

HasCountOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


