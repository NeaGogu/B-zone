# ServiceWindowsSchemeFiltersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **[]int32** | Service window scheme id&#39;s to filter by | [optional] 
**Name** | Pointer to **[]string** | Name of the service windows scheme(s) to filter by | [optional] 
**TagNames** | Pointer to **[]string** | tags to filter by | [optional] 
**TagIds** | Pointer to **[]int32** | tag ID&#39;s to filter by | [optional] 
**ZoneNames** | Pointer to **[]string** | zones to filter by | [optional] 
**ZoneIds** | Pointer to **[]int32** | zone ID&#39;s to filter by | [optional] 
**BrandNames** | Pointer to **[]string** | brands to filter by | [optional] 
**BrandIds** | Pointer to **[]int32** | brand ID&#39;s to filter by | [optional] 

## Methods

### NewServiceWindowsSchemeFiltersModel

`func NewServiceWindowsSchemeFiltersModel() *ServiceWindowsSchemeFiltersModel`

NewServiceWindowsSchemeFiltersModel instantiates a new ServiceWindowsSchemeFiltersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWindowsSchemeFiltersModelWithDefaults

`func NewServiceWindowsSchemeFiltersModelWithDefaults() *ServiceWindowsSchemeFiltersModel`

NewServiceWindowsSchemeFiltersModelWithDefaults instantiates a new ServiceWindowsSchemeFiltersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceWindowsSchemeFiltersModel) GetId() []int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceWindowsSchemeFiltersModel) GetIdOk() (*[]int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceWindowsSchemeFiltersModel) SetId(v []int32)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceWindowsSchemeFiltersModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ServiceWindowsSchemeFiltersModel) GetName() []string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceWindowsSchemeFiltersModel) GetNameOk() (*[]string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceWindowsSchemeFiltersModel) SetName(v []string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceWindowsSchemeFiltersModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTagNames

`func (o *ServiceWindowsSchemeFiltersModel) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *ServiceWindowsSchemeFiltersModel) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *ServiceWindowsSchemeFiltersModel) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.

### HasTagNames

`func (o *ServiceWindowsSchemeFiltersModel) HasTagNames() bool`

HasTagNames returns a boolean if a field has been set.

### GetTagIds

`func (o *ServiceWindowsSchemeFiltersModel) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *ServiceWindowsSchemeFiltersModel) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *ServiceWindowsSchemeFiltersModel) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *ServiceWindowsSchemeFiltersModel) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetZoneNames

`func (o *ServiceWindowsSchemeFiltersModel) GetZoneNames() []string`

GetZoneNames returns the ZoneNames field if non-nil, zero value otherwise.

### GetZoneNamesOk

`func (o *ServiceWindowsSchemeFiltersModel) GetZoneNamesOk() (*[]string, bool)`

GetZoneNamesOk returns a tuple with the ZoneNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneNames

`func (o *ServiceWindowsSchemeFiltersModel) SetZoneNames(v []string)`

SetZoneNames sets ZoneNames field to given value.

### HasZoneNames

`func (o *ServiceWindowsSchemeFiltersModel) HasZoneNames() bool`

HasZoneNames returns a boolean if a field has been set.

### GetZoneIds

`func (o *ServiceWindowsSchemeFiltersModel) GetZoneIds() []int32`

GetZoneIds returns the ZoneIds field if non-nil, zero value otherwise.

### GetZoneIdsOk

`func (o *ServiceWindowsSchemeFiltersModel) GetZoneIdsOk() (*[]int32, bool)`

GetZoneIdsOk returns a tuple with the ZoneIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneIds

`func (o *ServiceWindowsSchemeFiltersModel) SetZoneIds(v []int32)`

SetZoneIds sets ZoneIds field to given value.

### HasZoneIds

`func (o *ServiceWindowsSchemeFiltersModel) HasZoneIds() bool`

HasZoneIds returns a boolean if a field has been set.

### GetBrandNames

`func (o *ServiceWindowsSchemeFiltersModel) GetBrandNames() []string`

GetBrandNames returns the BrandNames field if non-nil, zero value otherwise.

### GetBrandNamesOk

`func (o *ServiceWindowsSchemeFiltersModel) GetBrandNamesOk() (*[]string, bool)`

GetBrandNamesOk returns a tuple with the BrandNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandNames

`func (o *ServiceWindowsSchemeFiltersModel) SetBrandNames(v []string)`

SetBrandNames sets BrandNames field to given value.

### HasBrandNames

`func (o *ServiceWindowsSchemeFiltersModel) HasBrandNames() bool`

HasBrandNames returns a boolean if a field has been set.

### GetBrandIds

`func (o *ServiceWindowsSchemeFiltersModel) GetBrandIds() []int32`

GetBrandIds returns the BrandIds field if non-nil, zero value otherwise.

### GetBrandIdsOk

`func (o *ServiceWindowsSchemeFiltersModel) GetBrandIdsOk() (*[]int32, bool)`

GetBrandIdsOk returns a tuple with the BrandIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandIds

`func (o *ServiceWindowsSchemeFiltersModel) SetBrandIds(v []int32)`

SetBrandIds sets BrandIds field to given value.

### HasBrandIds

`func (o *ServiceWindowsSchemeFiltersModel) HasBrandIds() bool`

HasBrandIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


