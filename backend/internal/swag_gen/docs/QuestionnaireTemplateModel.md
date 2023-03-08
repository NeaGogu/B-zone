# QuestionnaireTemplateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique Identifier | [optional] 
**Active** | Pointer to **bool** | if active&#x3D;0: QuestionnaireTemplate has been removed and is no longer visible in any bumbal interface | [optional] 
**Brands** | Pointer to [**[]BrandModel**](BrandModel.md) |  | [optional] 
**BrandIds** | Pointer to **[]int32** | Brand ids | [optional] 
**BrandNames** | Pointer to **[]string** | Brand names | [optional] 
**Zones** | Pointer to [**[]ZoneModel**](ZoneModel.md) |  | [optional] 
**ZoneIds** | Pointer to **[]int32** | Zone ids | [optional] 
**ZoneNames** | Pointer to **[]string** | Zone names | [optional] 
**Tags** | Pointer to [**[]TagModel**](TagModel.md) |  | [optional] 
**TagIds** | Pointer to **[]int32** | Tag ids | [optional] 
**TagNames** | Pointer to **[]string** | Tag names | [optional] 
**Name** | Pointer to **string** | Name of the questionnaire template | [optional] 
**QuestionnaireTypeId** | Pointer to **int64** | Questionnaire type id | [optional] 
**QuestionnaireTypeName** | Pointer to **string** | Name of the questionnaire type | [optional] 
**NoTags** | Pointer to **bool** | if no_tags&#x3D;1: no tags are used for matching | [optional] 
**NoZones** | Pointer to **bool** | if no_zones&#x3D;1: no zones are used for matching | [optional] 
**Questions** | Pointer to [**[]QuestionnaireTemplateQuestionModel**](QuestionnaireTemplateQuestionModel.md) |  | [optional] 

## Methods

### NewQuestionnaireTemplateModel

`func NewQuestionnaireTemplateModel() *QuestionnaireTemplateModel`

NewQuestionnaireTemplateModel instantiates a new QuestionnaireTemplateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionnaireTemplateModelWithDefaults

`func NewQuestionnaireTemplateModelWithDefaults() *QuestionnaireTemplateModel`

NewQuestionnaireTemplateModelWithDefaults instantiates a new QuestionnaireTemplateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuestionnaireTemplateModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuestionnaireTemplateModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuestionnaireTemplateModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *QuestionnaireTemplateModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActive

`func (o *QuestionnaireTemplateModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *QuestionnaireTemplateModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *QuestionnaireTemplateModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *QuestionnaireTemplateModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetBrands

`func (o *QuestionnaireTemplateModel) GetBrands() []BrandModel`

GetBrands returns the Brands field if non-nil, zero value otherwise.

### GetBrandsOk

`func (o *QuestionnaireTemplateModel) GetBrandsOk() (*[]BrandModel, bool)`

GetBrandsOk returns a tuple with the Brands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrands

`func (o *QuestionnaireTemplateModel) SetBrands(v []BrandModel)`

SetBrands sets Brands field to given value.

### HasBrands

`func (o *QuestionnaireTemplateModel) HasBrands() bool`

HasBrands returns a boolean if a field has been set.

### GetBrandIds

`func (o *QuestionnaireTemplateModel) GetBrandIds() []int32`

GetBrandIds returns the BrandIds field if non-nil, zero value otherwise.

### GetBrandIdsOk

`func (o *QuestionnaireTemplateModel) GetBrandIdsOk() (*[]int32, bool)`

GetBrandIdsOk returns a tuple with the BrandIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandIds

`func (o *QuestionnaireTemplateModel) SetBrandIds(v []int32)`

SetBrandIds sets BrandIds field to given value.

### HasBrandIds

`func (o *QuestionnaireTemplateModel) HasBrandIds() bool`

HasBrandIds returns a boolean if a field has been set.

### GetBrandNames

`func (o *QuestionnaireTemplateModel) GetBrandNames() []string`

GetBrandNames returns the BrandNames field if non-nil, zero value otherwise.

### GetBrandNamesOk

`func (o *QuestionnaireTemplateModel) GetBrandNamesOk() (*[]string, bool)`

GetBrandNamesOk returns a tuple with the BrandNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandNames

`func (o *QuestionnaireTemplateModel) SetBrandNames(v []string)`

SetBrandNames sets BrandNames field to given value.

### HasBrandNames

`func (o *QuestionnaireTemplateModel) HasBrandNames() bool`

HasBrandNames returns a boolean if a field has been set.

### GetZones

`func (o *QuestionnaireTemplateModel) GetZones() []ZoneModel`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *QuestionnaireTemplateModel) GetZonesOk() (*[]ZoneModel, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *QuestionnaireTemplateModel) SetZones(v []ZoneModel)`

SetZones sets Zones field to given value.

### HasZones

`func (o *QuestionnaireTemplateModel) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetZoneIds

`func (o *QuestionnaireTemplateModel) GetZoneIds() []int32`

GetZoneIds returns the ZoneIds field if non-nil, zero value otherwise.

### GetZoneIdsOk

`func (o *QuestionnaireTemplateModel) GetZoneIdsOk() (*[]int32, bool)`

GetZoneIdsOk returns a tuple with the ZoneIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneIds

`func (o *QuestionnaireTemplateModel) SetZoneIds(v []int32)`

SetZoneIds sets ZoneIds field to given value.

### HasZoneIds

`func (o *QuestionnaireTemplateModel) HasZoneIds() bool`

HasZoneIds returns a boolean if a field has been set.

### GetZoneNames

`func (o *QuestionnaireTemplateModel) GetZoneNames() []string`

GetZoneNames returns the ZoneNames field if non-nil, zero value otherwise.

### GetZoneNamesOk

`func (o *QuestionnaireTemplateModel) GetZoneNamesOk() (*[]string, bool)`

GetZoneNamesOk returns a tuple with the ZoneNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneNames

`func (o *QuestionnaireTemplateModel) SetZoneNames(v []string)`

SetZoneNames sets ZoneNames field to given value.

### HasZoneNames

`func (o *QuestionnaireTemplateModel) HasZoneNames() bool`

HasZoneNames returns a boolean if a field has been set.

### GetTags

`func (o *QuestionnaireTemplateModel) GetTags() []TagModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *QuestionnaireTemplateModel) GetTagsOk() (*[]TagModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *QuestionnaireTemplateModel) SetTags(v []TagModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *QuestionnaireTemplateModel) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTagIds

`func (o *QuestionnaireTemplateModel) GetTagIds() []int32`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *QuestionnaireTemplateModel) GetTagIdsOk() (*[]int32, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *QuestionnaireTemplateModel) SetTagIds(v []int32)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *QuestionnaireTemplateModel) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTagNames

`func (o *QuestionnaireTemplateModel) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *QuestionnaireTemplateModel) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *QuestionnaireTemplateModel) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.

### HasTagNames

`func (o *QuestionnaireTemplateModel) HasTagNames() bool`

HasTagNames returns a boolean if a field has been set.

### GetName

`func (o *QuestionnaireTemplateModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuestionnaireTemplateModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuestionnaireTemplateModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QuestionnaireTemplateModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuestionnaireTypeId

`func (o *QuestionnaireTemplateModel) GetQuestionnaireTypeId() int64`

GetQuestionnaireTypeId returns the QuestionnaireTypeId field if non-nil, zero value otherwise.

### GetQuestionnaireTypeIdOk

`func (o *QuestionnaireTemplateModel) GetQuestionnaireTypeIdOk() (*int64, bool)`

GetQuestionnaireTypeIdOk returns a tuple with the QuestionnaireTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionnaireTypeId

`func (o *QuestionnaireTemplateModel) SetQuestionnaireTypeId(v int64)`

SetQuestionnaireTypeId sets QuestionnaireTypeId field to given value.

### HasQuestionnaireTypeId

`func (o *QuestionnaireTemplateModel) HasQuestionnaireTypeId() bool`

HasQuestionnaireTypeId returns a boolean if a field has been set.

### GetQuestionnaireTypeName

`func (o *QuestionnaireTemplateModel) GetQuestionnaireTypeName() string`

GetQuestionnaireTypeName returns the QuestionnaireTypeName field if non-nil, zero value otherwise.

### GetQuestionnaireTypeNameOk

`func (o *QuestionnaireTemplateModel) GetQuestionnaireTypeNameOk() (*string, bool)`

GetQuestionnaireTypeNameOk returns a tuple with the QuestionnaireTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionnaireTypeName

`func (o *QuestionnaireTemplateModel) SetQuestionnaireTypeName(v string)`

SetQuestionnaireTypeName sets QuestionnaireTypeName field to given value.

### HasQuestionnaireTypeName

`func (o *QuestionnaireTemplateModel) HasQuestionnaireTypeName() bool`

HasQuestionnaireTypeName returns a boolean if a field has been set.

### GetNoTags

`func (o *QuestionnaireTemplateModel) GetNoTags() bool`

GetNoTags returns the NoTags field if non-nil, zero value otherwise.

### GetNoTagsOk

`func (o *QuestionnaireTemplateModel) GetNoTagsOk() (*bool, bool)`

GetNoTagsOk returns a tuple with the NoTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoTags

`func (o *QuestionnaireTemplateModel) SetNoTags(v bool)`

SetNoTags sets NoTags field to given value.

### HasNoTags

`func (o *QuestionnaireTemplateModel) HasNoTags() bool`

HasNoTags returns a boolean if a field has been set.

### GetNoZones

`func (o *QuestionnaireTemplateModel) GetNoZones() bool`

GetNoZones returns the NoZones field if non-nil, zero value otherwise.

### GetNoZonesOk

`func (o *QuestionnaireTemplateModel) GetNoZonesOk() (*bool, bool)`

GetNoZonesOk returns a tuple with the NoZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoZones

`func (o *QuestionnaireTemplateModel) SetNoZones(v bool)`

SetNoZones sets NoZones field to given value.

### HasNoZones

`func (o *QuestionnaireTemplateModel) HasNoZones() bool`

HasNoZones returns a boolean if a field has been set.

### GetQuestions

`func (o *QuestionnaireTemplateModel) GetQuestions() []QuestionnaireTemplateQuestionModel`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *QuestionnaireTemplateModel) GetQuestionsOk() (*[]QuestionnaireTemplateQuestionModel, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *QuestionnaireTemplateModel) SetQuestions(v []QuestionnaireTemplateQuestionModel)`

SetQuestions sets Questions field to given value.

### HasQuestions

`func (o *QuestionnaireTemplateModel) HasQuestions() bool`

HasQuestions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


