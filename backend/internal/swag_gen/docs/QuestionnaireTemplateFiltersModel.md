# QuestionnaireTemplateFiltersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **[]int32** | QuestionnaireTemplate id&#39;s | [optional] 
**TagNames** | Pointer to **[]string** | Tag names | [optional] 
**ZoneNames** | Pointer to **[]string** | Zone names | [optional] 
**Link** | Pointer to **[]map[string]interface{}** | Activity Link ids | [optional] 
**Links** | Pointer to **[]map[string]interface{}** | Activity Link ids | [optional] 
**SearchText** | Pointer to **string** | free search through text and numeric type columns | [optional] 
**ZoneId** | Pointer to **int32** | Zone ID | [optional] 
**BrandId** | Pointer to **int32** | Brand ID | [optional] 
**UpdatedAtSince** | Pointer to **time.Time** | Show updated since | [optional] 
**UpdatedAtTill** | Pointer to **time.Time** | Show updated till | [optional] 
**QuestionnaireTypeName** | Pointer to **map[string]interface{}** | Questionnaire type name | [optional] [readonly] 

## Methods

### NewQuestionnaireTemplateFiltersModel

`func NewQuestionnaireTemplateFiltersModel() *QuestionnaireTemplateFiltersModel`

NewQuestionnaireTemplateFiltersModel instantiates a new QuestionnaireTemplateFiltersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionnaireTemplateFiltersModelWithDefaults

`func NewQuestionnaireTemplateFiltersModelWithDefaults() *QuestionnaireTemplateFiltersModel`

NewQuestionnaireTemplateFiltersModelWithDefaults instantiates a new QuestionnaireTemplateFiltersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuestionnaireTemplateFiltersModel) GetId() []int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuestionnaireTemplateFiltersModel) GetIdOk() (*[]int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuestionnaireTemplateFiltersModel) SetId(v []int32)`

SetId sets Id field to given value.

### HasId

`func (o *QuestionnaireTemplateFiltersModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTagNames

`func (o *QuestionnaireTemplateFiltersModel) GetTagNames() []string`

GetTagNames returns the TagNames field if non-nil, zero value otherwise.

### GetTagNamesOk

`func (o *QuestionnaireTemplateFiltersModel) GetTagNamesOk() (*[]string, bool)`

GetTagNamesOk returns a tuple with the TagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNames

`func (o *QuestionnaireTemplateFiltersModel) SetTagNames(v []string)`

SetTagNames sets TagNames field to given value.

### HasTagNames

`func (o *QuestionnaireTemplateFiltersModel) HasTagNames() bool`

HasTagNames returns a boolean if a field has been set.

### GetZoneNames

`func (o *QuestionnaireTemplateFiltersModel) GetZoneNames() []string`

GetZoneNames returns the ZoneNames field if non-nil, zero value otherwise.

### GetZoneNamesOk

`func (o *QuestionnaireTemplateFiltersModel) GetZoneNamesOk() (*[]string, bool)`

GetZoneNamesOk returns a tuple with the ZoneNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneNames

`func (o *QuestionnaireTemplateFiltersModel) SetZoneNames(v []string)`

SetZoneNames sets ZoneNames field to given value.

### HasZoneNames

`func (o *QuestionnaireTemplateFiltersModel) HasZoneNames() bool`

HasZoneNames returns a boolean if a field has been set.

### GetLink

`func (o *QuestionnaireTemplateFiltersModel) GetLink() []map[string]interface{}`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *QuestionnaireTemplateFiltersModel) GetLinkOk() (*[]map[string]interface{}, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *QuestionnaireTemplateFiltersModel) SetLink(v []map[string]interface{})`

SetLink sets Link field to given value.

### HasLink

`func (o *QuestionnaireTemplateFiltersModel) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetLinks

`func (o *QuestionnaireTemplateFiltersModel) GetLinks() []map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *QuestionnaireTemplateFiltersModel) GetLinksOk() (*[]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *QuestionnaireTemplateFiltersModel) SetLinks(v []map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *QuestionnaireTemplateFiltersModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetSearchText

`func (o *QuestionnaireTemplateFiltersModel) GetSearchText() string`

GetSearchText returns the SearchText field if non-nil, zero value otherwise.

### GetSearchTextOk

`func (o *QuestionnaireTemplateFiltersModel) GetSearchTextOk() (*string, bool)`

GetSearchTextOk returns a tuple with the SearchText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchText

`func (o *QuestionnaireTemplateFiltersModel) SetSearchText(v string)`

SetSearchText sets SearchText field to given value.

### HasSearchText

`func (o *QuestionnaireTemplateFiltersModel) HasSearchText() bool`

HasSearchText returns a boolean if a field has been set.

### GetZoneId

`func (o *QuestionnaireTemplateFiltersModel) GetZoneId() int32`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *QuestionnaireTemplateFiltersModel) GetZoneIdOk() (*int32, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *QuestionnaireTemplateFiltersModel) SetZoneId(v int32)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *QuestionnaireTemplateFiltersModel) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetBrandId

`func (o *QuestionnaireTemplateFiltersModel) GetBrandId() int32`

GetBrandId returns the BrandId field if non-nil, zero value otherwise.

### GetBrandIdOk

`func (o *QuestionnaireTemplateFiltersModel) GetBrandIdOk() (*int32, bool)`

GetBrandIdOk returns a tuple with the BrandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandId

`func (o *QuestionnaireTemplateFiltersModel) SetBrandId(v int32)`

SetBrandId sets BrandId field to given value.

### HasBrandId

`func (o *QuestionnaireTemplateFiltersModel) HasBrandId() bool`

HasBrandId returns a boolean if a field has been set.

### GetUpdatedAtSince

`func (o *QuestionnaireTemplateFiltersModel) GetUpdatedAtSince() time.Time`

GetUpdatedAtSince returns the UpdatedAtSince field if non-nil, zero value otherwise.

### GetUpdatedAtSinceOk

`func (o *QuestionnaireTemplateFiltersModel) GetUpdatedAtSinceOk() (*time.Time, bool)`

GetUpdatedAtSinceOk returns a tuple with the UpdatedAtSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtSince

`func (o *QuestionnaireTemplateFiltersModel) SetUpdatedAtSince(v time.Time)`

SetUpdatedAtSince sets UpdatedAtSince field to given value.

### HasUpdatedAtSince

`func (o *QuestionnaireTemplateFiltersModel) HasUpdatedAtSince() bool`

HasUpdatedAtSince returns a boolean if a field has been set.

### GetUpdatedAtTill

`func (o *QuestionnaireTemplateFiltersModel) GetUpdatedAtTill() time.Time`

GetUpdatedAtTill returns the UpdatedAtTill field if non-nil, zero value otherwise.

### GetUpdatedAtTillOk

`func (o *QuestionnaireTemplateFiltersModel) GetUpdatedAtTillOk() (*time.Time, bool)`

GetUpdatedAtTillOk returns a tuple with the UpdatedAtTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtTill

`func (o *QuestionnaireTemplateFiltersModel) SetUpdatedAtTill(v time.Time)`

SetUpdatedAtTill sets UpdatedAtTill field to given value.

### HasUpdatedAtTill

`func (o *QuestionnaireTemplateFiltersModel) HasUpdatedAtTill() bool`

HasUpdatedAtTill returns a boolean if a field has been set.

### GetQuestionnaireTypeName

`func (o *QuestionnaireTemplateFiltersModel) GetQuestionnaireTypeName() map[string]interface{}`

GetQuestionnaireTypeName returns the QuestionnaireTypeName field if non-nil, zero value otherwise.

### GetQuestionnaireTypeNameOk

`func (o *QuestionnaireTemplateFiltersModel) GetQuestionnaireTypeNameOk() (*map[string]interface{}, bool)`

GetQuestionnaireTypeNameOk returns a tuple with the QuestionnaireTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionnaireTypeName

`func (o *QuestionnaireTemplateFiltersModel) SetQuestionnaireTypeName(v map[string]interface{})`

SetQuestionnaireTypeName sets QuestionnaireTypeName field to given value.

### HasQuestionnaireTypeName

`func (o *QuestionnaireTemplateFiltersModel) HasQuestionnaireTypeName() bool`

HasQuestionnaireTypeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


