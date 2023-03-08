# LinkRetrieveListArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | Pointer to [**LinkOptionsModel**](LinkOptionsModel.md) |  | [optional] 
**Filters** | Pointer to [**LinkFiltersModel**](LinkFiltersModel.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**SortingColumn** | Pointer to **string** | Sorting Column | [optional] 
**SortingDirection** | Pointer to **string** | Sorting Direction | [optional] 
**SearchText** | Pointer to **string** |  | [optional] 
**CountOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewLinkRetrieveListArguments

`func NewLinkRetrieveListArguments() *LinkRetrieveListArguments`

NewLinkRetrieveListArguments instantiates a new LinkRetrieveListArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkRetrieveListArgumentsWithDefaults

`func NewLinkRetrieveListArgumentsWithDefaults() *LinkRetrieveListArguments`

NewLinkRetrieveListArgumentsWithDefaults instantiates a new LinkRetrieveListArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *LinkRetrieveListArguments) GetOptions() LinkOptionsModel`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *LinkRetrieveListArguments) GetOptionsOk() (*LinkOptionsModel, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *LinkRetrieveListArguments) SetOptions(v LinkOptionsModel)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *LinkRetrieveListArguments) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetFilters

`func (o *LinkRetrieveListArguments) GetFilters() LinkFiltersModel`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *LinkRetrieveListArguments) GetFiltersOk() (*LinkFiltersModel, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *LinkRetrieveListArguments) SetFilters(v LinkFiltersModel)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *LinkRetrieveListArguments) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetLimit

`func (o *LinkRetrieveListArguments) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *LinkRetrieveListArguments) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *LinkRetrieveListArguments) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *LinkRetrieveListArguments) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *LinkRetrieveListArguments) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *LinkRetrieveListArguments) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *LinkRetrieveListArguments) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *LinkRetrieveListArguments) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetSortingColumn

`func (o *LinkRetrieveListArguments) GetSortingColumn() string`

GetSortingColumn returns the SortingColumn field if non-nil, zero value otherwise.

### GetSortingColumnOk

`func (o *LinkRetrieveListArguments) GetSortingColumnOk() (*string, bool)`

GetSortingColumnOk returns a tuple with the SortingColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortingColumn

`func (o *LinkRetrieveListArguments) SetSortingColumn(v string)`

SetSortingColumn sets SortingColumn field to given value.

### HasSortingColumn

`func (o *LinkRetrieveListArguments) HasSortingColumn() bool`

HasSortingColumn returns a boolean if a field has been set.

### GetSortingDirection

`func (o *LinkRetrieveListArguments) GetSortingDirection() string`

GetSortingDirection returns the SortingDirection field if non-nil, zero value otherwise.

### GetSortingDirectionOk

`func (o *LinkRetrieveListArguments) GetSortingDirectionOk() (*string, bool)`

GetSortingDirectionOk returns a tuple with the SortingDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortingDirection

`func (o *LinkRetrieveListArguments) SetSortingDirection(v string)`

SetSortingDirection sets SortingDirection field to given value.

### HasSortingDirection

`func (o *LinkRetrieveListArguments) HasSortingDirection() bool`

HasSortingDirection returns a boolean if a field has been set.

### GetSearchText

`func (o *LinkRetrieveListArguments) GetSearchText() string`

GetSearchText returns the SearchText field if non-nil, zero value otherwise.

### GetSearchTextOk

`func (o *LinkRetrieveListArguments) GetSearchTextOk() (*string, bool)`

GetSearchTextOk returns a tuple with the SearchText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchText

`func (o *LinkRetrieveListArguments) SetSearchText(v string)`

SetSearchText sets SearchText field to given value.

### HasSearchText

`func (o *LinkRetrieveListArguments) HasSearchText() bool`

HasSearchText returns a boolean if a field has been set.

### GetCountOnly

`func (o *LinkRetrieveListArguments) GetCountOnly() bool`

GetCountOnly returns the CountOnly field if non-nil, zero value otherwise.

### GetCountOnlyOk

`func (o *LinkRetrieveListArguments) GetCountOnlyOk() (*bool, bool)`

GetCountOnlyOk returns a tuple with the CountOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountOnly

`func (o *LinkRetrieveListArguments) SetCountOnly(v bool)`

SetCountOnly sets CountOnly field to given value.

### HasCountOnly

`func (o *LinkRetrieveListArguments) HasCountOnly() bool`

HasCountOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


