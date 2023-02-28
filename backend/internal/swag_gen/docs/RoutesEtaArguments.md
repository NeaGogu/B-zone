# RoutesEtaArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | Pointer to [**RoutesEtaOptionsModel**](RoutesEtaOptionsModel.md) |  | [optional] 
**Filters** | Pointer to [**RoutesEtaFiltersModel**](RoutesEtaFiltersModel.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**SearchText** | Pointer to **string** |  | [optional] 
**SortingColumn** | Pointer to **string** | Sorting Column | [optional] 
**SortingDirection** | Pointer to **string** | Sorting Direction | [optional] 
**AsList** | Pointer to **bool** |  | [optional] 

## Methods

### NewRoutesEtaArguments

`func NewRoutesEtaArguments() *RoutesEtaArguments`

NewRoutesEtaArguments instantiates a new RoutesEtaArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutesEtaArgumentsWithDefaults

`func NewRoutesEtaArgumentsWithDefaults() *RoutesEtaArguments`

NewRoutesEtaArgumentsWithDefaults instantiates a new RoutesEtaArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *RoutesEtaArguments) GetOptions() RoutesEtaOptionsModel`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *RoutesEtaArguments) GetOptionsOk() (*RoutesEtaOptionsModel, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *RoutesEtaArguments) SetOptions(v RoutesEtaOptionsModel)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *RoutesEtaArguments) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetFilters

`func (o *RoutesEtaArguments) GetFilters() RoutesEtaFiltersModel`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *RoutesEtaArguments) GetFiltersOk() (*RoutesEtaFiltersModel, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *RoutesEtaArguments) SetFilters(v RoutesEtaFiltersModel)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *RoutesEtaArguments) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetLimit

`func (o *RoutesEtaArguments) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *RoutesEtaArguments) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *RoutesEtaArguments) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *RoutesEtaArguments) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *RoutesEtaArguments) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *RoutesEtaArguments) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *RoutesEtaArguments) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *RoutesEtaArguments) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetSearchText

`func (o *RoutesEtaArguments) GetSearchText() string`

GetSearchText returns the SearchText field if non-nil, zero value otherwise.

### GetSearchTextOk

`func (o *RoutesEtaArguments) GetSearchTextOk() (*string, bool)`

GetSearchTextOk returns a tuple with the SearchText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchText

`func (o *RoutesEtaArguments) SetSearchText(v string)`

SetSearchText sets SearchText field to given value.

### HasSearchText

`func (o *RoutesEtaArguments) HasSearchText() bool`

HasSearchText returns a boolean if a field has been set.

### GetSortingColumn

`func (o *RoutesEtaArguments) GetSortingColumn() string`

GetSortingColumn returns the SortingColumn field if non-nil, zero value otherwise.

### GetSortingColumnOk

`func (o *RoutesEtaArguments) GetSortingColumnOk() (*string, bool)`

GetSortingColumnOk returns a tuple with the SortingColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortingColumn

`func (o *RoutesEtaArguments) SetSortingColumn(v string)`

SetSortingColumn sets SortingColumn field to given value.

### HasSortingColumn

`func (o *RoutesEtaArguments) HasSortingColumn() bool`

HasSortingColumn returns a boolean if a field has been set.

### GetSortingDirection

`func (o *RoutesEtaArguments) GetSortingDirection() string`

GetSortingDirection returns the SortingDirection field if non-nil, zero value otherwise.

### GetSortingDirectionOk

`func (o *RoutesEtaArguments) GetSortingDirectionOk() (*string, bool)`

GetSortingDirectionOk returns a tuple with the SortingDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortingDirection

`func (o *RoutesEtaArguments) SetSortingDirection(v string)`

SetSortingDirection sets SortingDirection field to given value.

### HasSortingDirection

`func (o *RoutesEtaArguments) HasSortingDirection() bool`

HasSortingDirection returns a boolean if a field has been set.

### GetAsList

`func (o *RoutesEtaArguments) GetAsList() bool`

GetAsList returns the AsList field if non-nil, zero value otherwise.

### GetAsListOk

`func (o *RoutesEtaArguments) GetAsListOk() (*bool, bool)`

GetAsListOk returns a tuple with the AsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsList

`func (o *RoutesEtaArguments) SetAsList(v bool)`

SetAsList sets AsList field to given value.

### HasAsList

`func (o *RoutesEtaArguments) HasAsList() bool`

HasAsList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


