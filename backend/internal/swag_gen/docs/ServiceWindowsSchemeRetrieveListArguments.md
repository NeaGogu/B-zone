# ServiceWindowsSchemeRetrieveListArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | Pointer to [**ServiceWindowsSchemeOptionsModel**](ServiceWindowsSchemeOptionsModel.md) |  | [optional] 
**Filters** | Pointer to [**ServiceWindowsSchemeFiltersModel**](ServiceWindowsSchemeFiltersModel.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**SearchText** | Pointer to **string** |  | [optional] 
**CountOnly** | Pointer to **bool** |  | [optional] 
**AsList** | Pointer to **bool** |  | [optional] 
**SortingColumn** | Pointer to **string** | Sorting Column | [optional] 
**SortingDirection** | Pointer to **string** | Sorting Direction | [optional] 

## Methods

### NewServiceWindowsSchemeRetrieveListArguments

`func NewServiceWindowsSchemeRetrieveListArguments() *ServiceWindowsSchemeRetrieveListArguments`

NewServiceWindowsSchemeRetrieveListArguments instantiates a new ServiceWindowsSchemeRetrieveListArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWindowsSchemeRetrieveListArgumentsWithDefaults

`func NewServiceWindowsSchemeRetrieveListArgumentsWithDefaults() *ServiceWindowsSchemeRetrieveListArguments`

NewServiceWindowsSchemeRetrieveListArgumentsWithDefaults instantiates a new ServiceWindowsSchemeRetrieveListArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *ServiceWindowsSchemeRetrieveListArguments) GetOptions() ServiceWindowsSchemeOptionsModel`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ServiceWindowsSchemeRetrieveListArguments) GetOptionsOk() (*ServiceWindowsSchemeOptionsModel, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ServiceWindowsSchemeRetrieveListArguments) SetOptions(v ServiceWindowsSchemeOptionsModel)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ServiceWindowsSchemeRetrieveListArguments) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetFilters

`func (o *ServiceWindowsSchemeRetrieveListArguments) GetFilters() ServiceWindowsSchemeFiltersModel`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ServiceWindowsSchemeRetrieveListArguments) GetFiltersOk() (*ServiceWindowsSchemeFiltersModel, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ServiceWindowsSchemeRetrieveListArguments) SetFilters(v ServiceWindowsSchemeFiltersModel)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ServiceWindowsSchemeRetrieveListArguments) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetLimit

`func (o *ServiceWindowsSchemeRetrieveListArguments) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ServiceWindowsSchemeRetrieveListArguments) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ServiceWindowsSchemeRetrieveListArguments) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ServiceWindowsSchemeRetrieveListArguments) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *ServiceWindowsSchemeRetrieveListArguments) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ServiceWindowsSchemeRetrieveListArguments) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ServiceWindowsSchemeRetrieveListArguments) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ServiceWindowsSchemeRetrieveListArguments) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetSearchText

`func (o *ServiceWindowsSchemeRetrieveListArguments) GetSearchText() string`

GetSearchText returns the SearchText field if non-nil, zero value otherwise.

### GetSearchTextOk

`func (o *ServiceWindowsSchemeRetrieveListArguments) GetSearchTextOk() (*string, bool)`

GetSearchTextOk returns a tuple with the SearchText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchText

`func (o *ServiceWindowsSchemeRetrieveListArguments) SetSearchText(v string)`

SetSearchText sets SearchText field to given value.

### HasSearchText

`func (o *ServiceWindowsSchemeRetrieveListArguments) HasSearchText() bool`

HasSearchText returns a boolean if a field has been set.

### GetCountOnly

`func (o *ServiceWindowsSchemeRetrieveListArguments) GetCountOnly() bool`

GetCountOnly returns the CountOnly field if non-nil, zero value otherwise.

### GetCountOnlyOk

`func (o *ServiceWindowsSchemeRetrieveListArguments) GetCountOnlyOk() (*bool, bool)`

GetCountOnlyOk returns a tuple with the CountOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountOnly

`func (o *ServiceWindowsSchemeRetrieveListArguments) SetCountOnly(v bool)`

SetCountOnly sets CountOnly field to given value.

### HasCountOnly

`func (o *ServiceWindowsSchemeRetrieveListArguments) HasCountOnly() bool`

HasCountOnly returns a boolean if a field has been set.

### GetAsList

`func (o *ServiceWindowsSchemeRetrieveListArguments) GetAsList() bool`

GetAsList returns the AsList field if non-nil, zero value otherwise.

### GetAsListOk

`func (o *ServiceWindowsSchemeRetrieveListArguments) GetAsListOk() (*bool, bool)`

GetAsListOk returns a tuple with the AsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsList

`func (o *ServiceWindowsSchemeRetrieveListArguments) SetAsList(v bool)`

SetAsList sets AsList field to given value.

### HasAsList

`func (o *ServiceWindowsSchemeRetrieveListArguments) HasAsList() bool`

HasAsList returns a boolean if a field has been set.

### GetSortingColumn

`func (o *ServiceWindowsSchemeRetrieveListArguments) GetSortingColumn() string`

GetSortingColumn returns the SortingColumn field if non-nil, zero value otherwise.

### GetSortingColumnOk

`func (o *ServiceWindowsSchemeRetrieveListArguments) GetSortingColumnOk() (*string, bool)`

GetSortingColumnOk returns a tuple with the SortingColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortingColumn

`func (o *ServiceWindowsSchemeRetrieveListArguments) SetSortingColumn(v string)`

SetSortingColumn sets SortingColumn field to given value.

### HasSortingColumn

`func (o *ServiceWindowsSchemeRetrieveListArguments) HasSortingColumn() bool`

HasSortingColumn returns a boolean if a field has been set.

### GetSortingDirection

`func (o *ServiceWindowsSchemeRetrieveListArguments) GetSortingDirection() string`

GetSortingDirection returns the SortingDirection field if non-nil, zero value otherwise.

### GetSortingDirectionOk

`func (o *ServiceWindowsSchemeRetrieveListArguments) GetSortingDirectionOk() (*string, bool)`

GetSortingDirectionOk returns a tuple with the SortingDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortingDirection

`func (o *ServiceWindowsSchemeRetrieveListArguments) SetSortingDirection(v string)`

SetSortingDirection sets SortingDirection field to given value.

### HasSortingDirection

`func (o *ServiceWindowsSchemeRetrieveListArguments) HasSortingDirection() bool`

HasSortingDirection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


