# ServiceWindowsSchemeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique ID | [optional] 
**Name** | Pointer to **string** | Service windows scheme name | [optional] 
**EvenWeeks** | Pointer to **bool** | even weeks | [optional] 
**OddWeeks** | Pointer to **bool** | odd weeks | [optional] 
**Tags** | Pointer to [**[]TagModel**](TagModel.md) |  | [optional] 
**Zones** | Pointer to [**[]ZoneModel**](ZoneModel.md) |  | [optional] 
**Brands** | Pointer to [**[]BrandModel**](BrandModel.md) |  | [optional] 
**NoTags** | Pointer to **bool** | No tags boolean value | [optional] 
**NoZones** | Pointer to **bool** | No zones boolean value | [optional] 
**ActivityTypeIds** | Pointer to **[]int32** | Activity type ids | [optional] 
**Monday** | Pointer to **map[string]interface{}** | ServiceWindowDayModel containing the cut off information | [optional] 
**Tuesday** | Pointer to **map[string]interface{}** | ServiceWindowDayModel containing the cut off information | [optional] 
**Wednesday** | Pointer to **map[string]interface{}** | ServiceWindowDayModel containing the cut off information | [optional] 
**Thursday** | Pointer to **map[string]interface{}** | ServiceWindowDayModel containing the cut off information | [optional] 
**Friday** | Pointer to **map[string]interface{}** | ServiceWindowDayModel containing the cut off information | [optional] 
**Saturday** | Pointer to **map[string]interface{}** | ServiceWindowDayModel containing the cut off information | [optional] 
**Sunday** | Pointer to **map[string]interface{}** | ServiceWindowDayModel containing the cut off information | [optional] 
**MinDaysAhead** | Pointer to **int32** | Number of min. days ahead | [optional] 
**MaxDaysAhead** | Pointer to **int32** | Number of max. days ahead | [optional] 
**ServiceWindows** | Pointer to [**[]ServiceWindowModel**](ServiceWindowModel.md) |  | [optional] 

## Methods

### NewServiceWindowsSchemeModel

`func NewServiceWindowsSchemeModel() *ServiceWindowsSchemeModel`

NewServiceWindowsSchemeModel instantiates a new ServiceWindowsSchemeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWindowsSchemeModelWithDefaults

`func NewServiceWindowsSchemeModelWithDefaults() *ServiceWindowsSchemeModel`

NewServiceWindowsSchemeModelWithDefaults instantiates a new ServiceWindowsSchemeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceWindowsSchemeModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceWindowsSchemeModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceWindowsSchemeModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceWindowsSchemeModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ServiceWindowsSchemeModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceWindowsSchemeModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceWindowsSchemeModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceWindowsSchemeModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEvenWeeks

`func (o *ServiceWindowsSchemeModel) GetEvenWeeks() bool`

GetEvenWeeks returns the EvenWeeks field if non-nil, zero value otherwise.

### GetEvenWeeksOk

`func (o *ServiceWindowsSchemeModel) GetEvenWeeksOk() (*bool, bool)`

GetEvenWeeksOk returns a tuple with the EvenWeeks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvenWeeks

`func (o *ServiceWindowsSchemeModel) SetEvenWeeks(v bool)`

SetEvenWeeks sets EvenWeeks field to given value.

### HasEvenWeeks

`func (o *ServiceWindowsSchemeModel) HasEvenWeeks() bool`

HasEvenWeeks returns a boolean if a field has been set.

### GetOddWeeks

`func (o *ServiceWindowsSchemeModel) GetOddWeeks() bool`

GetOddWeeks returns the OddWeeks field if non-nil, zero value otherwise.

### GetOddWeeksOk

`func (o *ServiceWindowsSchemeModel) GetOddWeeksOk() (*bool, bool)`

GetOddWeeksOk returns a tuple with the OddWeeks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOddWeeks

`func (o *ServiceWindowsSchemeModel) SetOddWeeks(v bool)`

SetOddWeeks sets OddWeeks field to given value.

### HasOddWeeks

`func (o *ServiceWindowsSchemeModel) HasOddWeeks() bool`

HasOddWeeks returns a boolean if a field has been set.

### GetTags

`func (o *ServiceWindowsSchemeModel) GetTags() []TagModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServiceWindowsSchemeModel) GetTagsOk() (*[]TagModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServiceWindowsSchemeModel) SetTags(v []TagModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServiceWindowsSchemeModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetZones

`func (o *ServiceWindowsSchemeModel) GetZones() []ZoneModel`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *ServiceWindowsSchemeModel) GetZonesOk() (*[]ZoneModel, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *ServiceWindowsSchemeModel) SetZones(v []ZoneModel)`

SetZones sets Zones field to given value.

### HasZones

`func (o *ServiceWindowsSchemeModel) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetBrands

`func (o *ServiceWindowsSchemeModel) GetBrands() []BrandModel`

GetBrands returns the Brands field if non-nil, zero value otherwise.

### GetBrandsOk

`func (o *ServiceWindowsSchemeModel) GetBrandsOk() (*[]BrandModel, bool)`

GetBrandsOk returns a tuple with the Brands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrands

`func (o *ServiceWindowsSchemeModel) SetBrands(v []BrandModel)`

SetBrands sets Brands field to given value.

### HasBrands

`func (o *ServiceWindowsSchemeModel) HasBrands() bool`

HasBrands returns a boolean if a field has been set.

### GetNoTags

`func (o *ServiceWindowsSchemeModel) GetNoTags() bool`

GetNoTags returns the NoTags field if non-nil, zero value otherwise.

### GetNoTagsOk

`func (o *ServiceWindowsSchemeModel) GetNoTagsOk() (*bool, bool)`

GetNoTagsOk returns a tuple with the NoTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoTags

`func (o *ServiceWindowsSchemeModel) SetNoTags(v bool)`

SetNoTags sets NoTags field to given value.

### HasNoTags

`func (o *ServiceWindowsSchemeModel) HasNoTags() bool`

HasNoTags returns a boolean if a field has been set.

### GetNoZones

`func (o *ServiceWindowsSchemeModel) GetNoZones() bool`

GetNoZones returns the NoZones field if non-nil, zero value otherwise.

### GetNoZonesOk

`func (o *ServiceWindowsSchemeModel) GetNoZonesOk() (*bool, bool)`

GetNoZonesOk returns a tuple with the NoZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoZones

`func (o *ServiceWindowsSchemeModel) SetNoZones(v bool)`

SetNoZones sets NoZones field to given value.

### HasNoZones

`func (o *ServiceWindowsSchemeModel) HasNoZones() bool`

HasNoZones returns a boolean if a field has been set.

### GetActivityTypeIds

`func (o *ServiceWindowsSchemeModel) GetActivityTypeIds() []int32`

GetActivityTypeIds returns the ActivityTypeIds field if non-nil, zero value otherwise.

### GetActivityTypeIdsOk

`func (o *ServiceWindowsSchemeModel) GetActivityTypeIdsOk() (*[]int32, bool)`

GetActivityTypeIdsOk returns a tuple with the ActivityTypeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityTypeIds

`func (o *ServiceWindowsSchemeModel) SetActivityTypeIds(v []int32)`

SetActivityTypeIds sets ActivityTypeIds field to given value.

### HasActivityTypeIds

`func (o *ServiceWindowsSchemeModel) HasActivityTypeIds() bool`

HasActivityTypeIds returns a boolean if a field has been set.

### GetMonday

`func (o *ServiceWindowsSchemeModel) GetMonday() map[string]interface{}`

GetMonday returns the Monday field if non-nil, zero value otherwise.

### GetMondayOk

`func (o *ServiceWindowsSchemeModel) GetMondayOk() (*map[string]interface{}, bool)`

GetMondayOk returns a tuple with the Monday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonday

`func (o *ServiceWindowsSchemeModel) SetMonday(v map[string]interface{})`

SetMonday sets Monday field to given value.

### HasMonday

`func (o *ServiceWindowsSchemeModel) HasMonday() bool`

HasMonday returns a boolean if a field has been set.

### GetTuesday

`func (o *ServiceWindowsSchemeModel) GetTuesday() map[string]interface{}`

GetTuesday returns the Tuesday field if non-nil, zero value otherwise.

### GetTuesdayOk

`func (o *ServiceWindowsSchemeModel) GetTuesdayOk() (*map[string]interface{}, bool)`

GetTuesdayOk returns a tuple with the Tuesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuesday

`func (o *ServiceWindowsSchemeModel) SetTuesday(v map[string]interface{})`

SetTuesday sets Tuesday field to given value.

### HasTuesday

`func (o *ServiceWindowsSchemeModel) HasTuesday() bool`

HasTuesday returns a boolean if a field has been set.

### GetWednesday

`func (o *ServiceWindowsSchemeModel) GetWednesday() map[string]interface{}`

GetWednesday returns the Wednesday field if non-nil, zero value otherwise.

### GetWednesdayOk

`func (o *ServiceWindowsSchemeModel) GetWednesdayOk() (*map[string]interface{}, bool)`

GetWednesdayOk returns a tuple with the Wednesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWednesday

`func (o *ServiceWindowsSchemeModel) SetWednesday(v map[string]interface{})`

SetWednesday sets Wednesday field to given value.

### HasWednesday

`func (o *ServiceWindowsSchemeModel) HasWednesday() bool`

HasWednesday returns a boolean if a field has been set.

### GetThursday

`func (o *ServiceWindowsSchemeModel) GetThursday() map[string]interface{}`

GetThursday returns the Thursday field if non-nil, zero value otherwise.

### GetThursdayOk

`func (o *ServiceWindowsSchemeModel) GetThursdayOk() (*map[string]interface{}, bool)`

GetThursdayOk returns a tuple with the Thursday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThursday

`func (o *ServiceWindowsSchemeModel) SetThursday(v map[string]interface{})`

SetThursday sets Thursday field to given value.

### HasThursday

`func (o *ServiceWindowsSchemeModel) HasThursday() bool`

HasThursday returns a boolean if a field has been set.

### GetFriday

`func (o *ServiceWindowsSchemeModel) GetFriday() map[string]interface{}`

GetFriday returns the Friday field if non-nil, zero value otherwise.

### GetFridayOk

`func (o *ServiceWindowsSchemeModel) GetFridayOk() (*map[string]interface{}, bool)`

GetFridayOk returns a tuple with the Friday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriday

`func (o *ServiceWindowsSchemeModel) SetFriday(v map[string]interface{})`

SetFriday sets Friday field to given value.

### HasFriday

`func (o *ServiceWindowsSchemeModel) HasFriday() bool`

HasFriday returns a boolean if a field has been set.

### GetSaturday

`func (o *ServiceWindowsSchemeModel) GetSaturday() map[string]interface{}`

GetSaturday returns the Saturday field if non-nil, zero value otherwise.

### GetSaturdayOk

`func (o *ServiceWindowsSchemeModel) GetSaturdayOk() (*map[string]interface{}, bool)`

GetSaturdayOk returns a tuple with the Saturday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturday

`func (o *ServiceWindowsSchemeModel) SetSaturday(v map[string]interface{})`

SetSaturday sets Saturday field to given value.

### HasSaturday

`func (o *ServiceWindowsSchemeModel) HasSaturday() bool`

HasSaturday returns a boolean if a field has been set.

### GetSunday

`func (o *ServiceWindowsSchemeModel) GetSunday() map[string]interface{}`

GetSunday returns the Sunday field if non-nil, zero value otherwise.

### GetSundayOk

`func (o *ServiceWindowsSchemeModel) GetSundayOk() (*map[string]interface{}, bool)`

GetSundayOk returns a tuple with the Sunday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSunday

`func (o *ServiceWindowsSchemeModel) SetSunday(v map[string]interface{})`

SetSunday sets Sunday field to given value.

### HasSunday

`func (o *ServiceWindowsSchemeModel) HasSunday() bool`

HasSunday returns a boolean if a field has been set.

### GetMinDaysAhead

`func (o *ServiceWindowsSchemeModel) GetMinDaysAhead() int32`

GetMinDaysAhead returns the MinDaysAhead field if non-nil, zero value otherwise.

### GetMinDaysAheadOk

`func (o *ServiceWindowsSchemeModel) GetMinDaysAheadOk() (*int32, bool)`

GetMinDaysAheadOk returns a tuple with the MinDaysAhead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDaysAhead

`func (o *ServiceWindowsSchemeModel) SetMinDaysAhead(v int32)`

SetMinDaysAhead sets MinDaysAhead field to given value.

### HasMinDaysAhead

`func (o *ServiceWindowsSchemeModel) HasMinDaysAhead() bool`

HasMinDaysAhead returns a boolean if a field has been set.

### GetMaxDaysAhead

`func (o *ServiceWindowsSchemeModel) GetMaxDaysAhead() int32`

GetMaxDaysAhead returns the MaxDaysAhead field if non-nil, zero value otherwise.

### GetMaxDaysAheadOk

`func (o *ServiceWindowsSchemeModel) GetMaxDaysAheadOk() (*int32, bool)`

GetMaxDaysAheadOk returns a tuple with the MaxDaysAhead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDaysAhead

`func (o *ServiceWindowsSchemeModel) SetMaxDaysAhead(v int32)`

SetMaxDaysAhead sets MaxDaysAhead field to given value.

### HasMaxDaysAhead

`func (o *ServiceWindowsSchemeModel) HasMaxDaysAhead() bool`

HasMaxDaysAhead returns a boolean if a field has been set.

### GetServiceWindows

`func (o *ServiceWindowsSchemeModel) GetServiceWindows() []ServiceWindowModel`

GetServiceWindows returns the ServiceWindows field if non-nil, zero value otherwise.

### GetServiceWindowsOk

`func (o *ServiceWindowsSchemeModel) GetServiceWindowsOk() (*[]ServiceWindowModel, bool)`

GetServiceWindowsOk returns a tuple with the ServiceWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceWindows

`func (o *ServiceWindowsSchemeModel) SetServiceWindows(v []ServiceWindowModel)`

SetServiceWindows sets ServiceWindows field to given value.

### HasServiceWindows

`func (o *ServiceWindowsSchemeModel) HasServiceWindows() bool`

HasServiceWindows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


