# ZoneModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique ID | [optional] 
**Name** | Pointer to **string** | Zone Name | [optional] 
**FilterTagNames** | Pointer to **[]string** | Tag names | [optional] 
**FilterTags** | Pointer to **map[string]interface{}** |  | [optional] 
**Links** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**MetaData** | Pointer to [**[]MetaDataModel**](MetaDataModel.md) |  | [optional] 
**ZoneRanges** | Pointer to [**[]ZoneRangeModel**](ZoneRangeModel.md) |  | [optional] 
**Brands** | Pointer to [**[]BrandModel**](BrandModel.md) |  | [optional] 
**BrandIds** | Pointer to **[]int32** | Brand ID&#39;s | [optional] 
**ZoneCreatedAt** | Pointer to **time.Time** | created_at date time | [optional] [readonly] 
**ZoneUpdatedAt** | Pointer to **time.Time** | updated_at date time | [optional] [readonly] 
**ZoneCreatedBy** | Pointer to **int32** | created_by user id | [optional] [readonly] 
**ZoneUpdatedBy** | Pointer to **int32** | created_by user id | [optional] [readonly] 

## Methods

### NewZoneModel

`func NewZoneModel() *ZoneModel`

NewZoneModel instantiates a new ZoneModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneModelWithDefaults

`func NewZoneModelWithDefaults() *ZoneModel`

NewZoneModelWithDefaults instantiates a new ZoneModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ZoneModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZoneModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZoneModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ZoneModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ZoneModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZoneModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZoneModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ZoneModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFilterTagNames

`func (o *ZoneModel) GetFilterTagNames() []string`

GetFilterTagNames returns the FilterTagNames field if non-nil, zero value otherwise.

### GetFilterTagNamesOk

`func (o *ZoneModel) GetFilterTagNamesOk() (*[]string, bool)`

GetFilterTagNamesOk returns a tuple with the FilterTagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTagNames

`func (o *ZoneModel) SetFilterTagNames(v []string)`

SetFilterTagNames sets FilterTagNames field to given value.

### HasFilterTagNames

`func (o *ZoneModel) HasFilterTagNames() bool`

HasFilterTagNames returns a boolean if a field has been set.

### GetFilterTags

`func (o *ZoneModel) GetFilterTags() map[string]interface{}`

GetFilterTags returns the FilterTags field if non-nil, zero value otherwise.

### GetFilterTagsOk

`func (o *ZoneModel) GetFilterTagsOk() (*map[string]interface{}, bool)`

GetFilterTagsOk returns a tuple with the FilterTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterTags

`func (o *ZoneModel) SetFilterTags(v map[string]interface{})`

SetFilterTags sets FilterTags field to given value.

### HasFilterTags

`func (o *ZoneModel) HasFilterTags() bool`

HasFilterTags returns a boolean if a field has been set.

### GetLinks

`func (o *ZoneModel) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ZoneModel) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ZoneModel) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ZoneModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetaData

`func (o *ZoneModel) GetMetaData() []MetaDataModel`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *ZoneModel) GetMetaDataOk() (*[]MetaDataModel, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *ZoneModel) SetMetaData(v []MetaDataModel)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *ZoneModel) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetZoneRanges

`func (o *ZoneModel) GetZoneRanges() []ZoneRangeModel`

GetZoneRanges returns the ZoneRanges field if non-nil, zero value otherwise.

### GetZoneRangesOk

`func (o *ZoneModel) GetZoneRangesOk() (*[]ZoneRangeModel, bool)`

GetZoneRangesOk returns a tuple with the ZoneRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneRanges

`func (o *ZoneModel) SetZoneRanges(v []ZoneRangeModel)`

SetZoneRanges sets ZoneRanges field to given value.

### HasZoneRanges

`func (o *ZoneModel) HasZoneRanges() bool`

HasZoneRanges returns a boolean if a field has been set.

### GetBrands

`func (o *ZoneModel) GetBrands() []BrandModel`

GetBrands returns the Brands field if non-nil, zero value otherwise.

### GetBrandsOk

`func (o *ZoneModel) GetBrandsOk() (*[]BrandModel, bool)`

GetBrandsOk returns a tuple with the Brands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrands

`func (o *ZoneModel) SetBrands(v []BrandModel)`

SetBrands sets Brands field to given value.

### HasBrands

`func (o *ZoneModel) HasBrands() bool`

HasBrands returns a boolean if a field has been set.

### GetBrandIds

`func (o *ZoneModel) GetBrandIds() []int32`

GetBrandIds returns the BrandIds field if non-nil, zero value otherwise.

### GetBrandIdsOk

`func (o *ZoneModel) GetBrandIdsOk() (*[]int32, bool)`

GetBrandIdsOk returns a tuple with the BrandIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandIds

`func (o *ZoneModel) SetBrandIds(v []int32)`

SetBrandIds sets BrandIds field to given value.

### HasBrandIds

`func (o *ZoneModel) HasBrandIds() bool`

HasBrandIds returns a boolean if a field has been set.

### GetZoneCreatedAt

`func (o *ZoneModel) GetZoneCreatedAt() time.Time`

GetZoneCreatedAt returns the ZoneCreatedAt field if non-nil, zero value otherwise.

### GetZoneCreatedAtOk

`func (o *ZoneModel) GetZoneCreatedAtOk() (*time.Time, bool)`

GetZoneCreatedAtOk returns a tuple with the ZoneCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneCreatedAt

`func (o *ZoneModel) SetZoneCreatedAt(v time.Time)`

SetZoneCreatedAt sets ZoneCreatedAt field to given value.

### HasZoneCreatedAt

`func (o *ZoneModel) HasZoneCreatedAt() bool`

HasZoneCreatedAt returns a boolean if a field has been set.

### GetZoneUpdatedAt

`func (o *ZoneModel) GetZoneUpdatedAt() time.Time`

GetZoneUpdatedAt returns the ZoneUpdatedAt field if non-nil, zero value otherwise.

### GetZoneUpdatedAtOk

`func (o *ZoneModel) GetZoneUpdatedAtOk() (*time.Time, bool)`

GetZoneUpdatedAtOk returns a tuple with the ZoneUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneUpdatedAt

`func (o *ZoneModel) SetZoneUpdatedAt(v time.Time)`

SetZoneUpdatedAt sets ZoneUpdatedAt field to given value.

### HasZoneUpdatedAt

`func (o *ZoneModel) HasZoneUpdatedAt() bool`

HasZoneUpdatedAt returns a boolean if a field has been set.

### GetZoneCreatedBy

`func (o *ZoneModel) GetZoneCreatedBy() int32`

GetZoneCreatedBy returns the ZoneCreatedBy field if non-nil, zero value otherwise.

### GetZoneCreatedByOk

`func (o *ZoneModel) GetZoneCreatedByOk() (*int32, bool)`

GetZoneCreatedByOk returns a tuple with the ZoneCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneCreatedBy

`func (o *ZoneModel) SetZoneCreatedBy(v int32)`

SetZoneCreatedBy sets ZoneCreatedBy field to given value.

### HasZoneCreatedBy

`func (o *ZoneModel) HasZoneCreatedBy() bool`

HasZoneCreatedBy returns a boolean if a field has been set.

### GetZoneUpdatedBy

`func (o *ZoneModel) GetZoneUpdatedBy() int32`

GetZoneUpdatedBy returns the ZoneUpdatedBy field if non-nil, zero value otherwise.

### GetZoneUpdatedByOk

`func (o *ZoneModel) GetZoneUpdatedByOk() (*int32, bool)`

GetZoneUpdatedByOk returns a tuple with the ZoneUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneUpdatedBy

`func (o *ZoneModel) SetZoneUpdatedBy(v int32)`

SetZoneUpdatedBy sets ZoneUpdatedBy field to given value.

### HasZoneUpdatedBy

`func (o *ZoneModel) HasZoneUpdatedBy() bool`

HasZoneUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


