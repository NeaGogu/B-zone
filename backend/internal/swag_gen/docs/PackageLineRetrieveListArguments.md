# PackageLineRetrieveListArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | Pointer to [**PackageLineOptionsModel**](PackageLineOptionsModel.md) |  | [optional] 
**Filters** | Pointer to [**PackageLineFiltersModel**](PackageLineFiltersModel.md) |  | [optional] 
**Limit** | Pointer to **int64** |  | [optional] 
**Offset** | Pointer to **int64** |  | [optional] 
**SortingColumn** | Pointer to **string** | Sorting Column | [optional] 
**SortingDirection** | Pointer to **string** | Sorting Direction | [optional] 
**SearchText** | Pointer to **string** |  | [optional] 
**AsList** | Pointer to **bool** |  | [optional] 
**CountOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewPackageLineRetrieveListArguments

`func NewPackageLineRetrieveListArguments() *PackageLineRetrieveListArguments`

NewPackageLineRetrieveListArguments instantiates a new PackageLineRetrieveListArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageLineRetrieveListArgumentsWithDefaults

`func NewPackageLineRetrieveListArgumentsWithDefaults() *PackageLineRetrieveListArguments`

NewPackageLineRetrieveListArgumentsWithDefaults instantiates a new PackageLineRetrieveListArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *PackageLineRetrieveListArguments) GetOptions() PackageLineOptionsModel`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PackageLineRetrieveListArguments) GetOptionsOk() (*PackageLineOptionsModel, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PackageLineRetrieveListArguments) SetOptions(v PackageLineOptionsModel)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *PackageLineRetrieveListArguments) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetFilters

`func (o *PackageLineRetrieveListArguments) GetFilters() PackageLineFiltersModel`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *PackageLineRetrieveListArguments) GetFiltersOk() (*PackageLineFiltersModel, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *PackageLineRetrieveListArguments) SetFilters(v PackageLineFiltersModel)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *PackageLineRetrieveListArguments) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetLimit

`func (o *PackageLineRetrieveListArguments) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PackageLineRetrieveListArguments) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PackageLineRetrieveListArguments) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PackageLineRetrieveListArguments) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *PackageLineRetrieveListArguments) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PackageLineRetrieveListArguments) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PackageLineRetrieveListArguments) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *PackageLineRetrieveListArguments) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetSortingColumn

`func (o *PackageLineRetrieveListArguments) GetSortingColumn() string`

GetSortingColumn returns the SortingColumn field if non-nil, zero value otherwise.

### GetSortingColumnOk

`func (o *PackageLineRetrieveListArguments) GetSortingColumnOk() (*string, bool)`

GetSortingColumnOk returns a tuple with the SortingColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortingColumn

`func (o *PackageLineRetrieveListArguments) SetSortingColumn(v string)`

SetSortingColumn sets SortingColumn field to given value.

### HasSortingColumn

`func (o *PackageLineRetrieveListArguments) HasSortingColumn() bool`

HasSortingColumn returns a boolean if a field has been set.

### GetSortingDirection

`func (o *PackageLineRetrieveListArguments) GetSortingDirection() string`

GetSortingDirection returns the SortingDirection field if non-nil, zero value otherwise.

### GetSortingDirectionOk

`func (o *PackageLineRetrieveListArguments) GetSortingDirectionOk() (*string, bool)`

GetSortingDirectionOk returns a tuple with the SortingDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortingDirection

`func (o *PackageLineRetrieveListArguments) SetSortingDirection(v string)`

SetSortingDirection sets SortingDirection field to given value.

### HasSortingDirection

`func (o *PackageLineRetrieveListArguments) HasSortingDirection() bool`

HasSortingDirection returns a boolean if a field has been set.

### GetSearchText

`func (o *PackageLineRetrieveListArguments) GetSearchText() string`

GetSearchText returns the SearchText field if non-nil, zero value otherwise.

### GetSearchTextOk

`func (o *PackageLineRetrieveListArguments) GetSearchTextOk() (*string, bool)`

GetSearchTextOk returns a tuple with the SearchText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchText

`func (o *PackageLineRetrieveListArguments) SetSearchText(v string)`

SetSearchText sets SearchText field to given value.

### HasSearchText

`func (o *PackageLineRetrieveListArguments) HasSearchText() bool`

HasSearchText returns a boolean if a field has been set.

### GetAsList

`func (o *PackageLineRetrieveListArguments) GetAsList() bool`

GetAsList returns the AsList field if non-nil, zero value otherwise.

### GetAsListOk

`func (o *PackageLineRetrieveListArguments) GetAsListOk() (*bool, bool)`

GetAsListOk returns a tuple with the AsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsList

`func (o *PackageLineRetrieveListArguments) SetAsList(v bool)`

SetAsList sets AsList field to given value.

### HasAsList

`func (o *PackageLineRetrieveListArguments) HasAsList() bool`

HasAsList returns a boolean if a field has been set.

### GetCountOnly

`func (o *PackageLineRetrieveListArguments) GetCountOnly() bool`

GetCountOnly returns the CountOnly field if non-nil, zero value otherwise.

### GetCountOnlyOk

`func (o *PackageLineRetrieveListArguments) GetCountOnlyOk() (*bool, bool)`

GetCountOnlyOk returns a tuple with the CountOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountOnly

`func (o *PackageLineRetrieveListArguments) SetCountOnly(v bool)`

SetCountOnly sets CountOnly field to given value.

### HasCountOnly

`func (o *PackageLineRetrieveListArguments) HasCountOnly() bool`

HasCountOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


