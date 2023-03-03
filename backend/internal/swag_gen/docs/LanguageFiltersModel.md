# LanguageFiltersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **[]int32** | Language id&#39;s | [optional] 
**SearchText** | Pointer to **string** | free search through text and numeric type columns | [optional] 
**LangCode** | Pointer to **string** | Find given lang_code | [optional] 

## Methods

### NewLanguageFiltersModel

`func NewLanguageFiltersModel() *LanguageFiltersModel`

NewLanguageFiltersModel instantiates a new LanguageFiltersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanguageFiltersModelWithDefaults

`func NewLanguageFiltersModelWithDefaults() *LanguageFiltersModel`

NewLanguageFiltersModelWithDefaults instantiates a new LanguageFiltersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LanguageFiltersModel) GetId() []int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LanguageFiltersModel) GetIdOk() (*[]int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LanguageFiltersModel) SetId(v []int32)`

SetId sets Id field to given value.

### HasId

`func (o *LanguageFiltersModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSearchText

`func (o *LanguageFiltersModel) GetSearchText() string`

GetSearchText returns the SearchText field if non-nil, zero value otherwise.

### GetSearchTextOk

`func (o *LanguageFiltersModel) GetSearchTextOk() (*string, bool)`

GetSearchTextOk returns a tuple with the SearchText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchText

`func (o *LanguageFiltersModel) SetSearchText(v string)`

SetSearchText sets SearchText field to given value.

### HasSearchText

`func (o *LanguageFiltersModel) HasSearchText() bool`

HasSearchText returns a boolean if a field has been set.

### GetLangCode

`func (o *LanguageFiltersModel) GetLangCode() string`

GetLangCode returns the LangCode field if non-nil, zero value otherwise.

### GetLangCodeOk

`func (o *LanguageFiltersModel) GetLangCodeOk() (*string, bool)`

GetLangCodeOk returns a tuple with the LangCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLangCode

`func (o *LanguageFiltersModel) SetLangCode(v string)`

SetLangCode sets LangCode field to given value.

### HasLangCode

`func (o *LanguageFiltersModel) HasLangCode() bool`

HasLangCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


