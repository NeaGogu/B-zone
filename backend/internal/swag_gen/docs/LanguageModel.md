# LanguageModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique Identifier | [optional] 
**Name** | Pointer to **int32** | Name of the language in English | [optional] 
**LangCode** | Pointer to **int32** | Two character language code | [optional] 
**Active** | Pointer to **bool** | if active&#x3D;0: Language has been removed and is no longer visible in any bumbal interface | [optional] 

## Methods

### NewLanguageModel

`func NewLanguageModel() *LanguageModel`

NewLanguageModel instantiates a new LanguageModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanguageModelWithDefaults

`func NewLanguageModelWithDefaults() *LanguageModel`

NewLanguageModelWithDefaults instantiates a new LanguageModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LanguageModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LanguageModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LanguageModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *LanguageModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *LanguageModel) GetName() int32`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LanguageModel) GetNameOk() (*int32, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LanguageModel) SetName(v int32)`

SetName sets Name field to given value.

### HasName

`func (o *LanguageModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLangCode

`func (o *LanguageModel) GetLangCode() int32`

GetLangCode returns the LangCode field if non-nil, zero value otherwise.

### GetLangCodeOk

`func (o *LanguageModel) GetLangCodeOk() (*int32, bool)`

GetLangCodeOk returns a tuple with the LangCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLangCode

`func (o *LanguageModel) SetLangCode(v int32)`

SetLangCode sets LangCode field to given value.

### HasLangCode

`func (o *LanguageModel) HasLangCode() bool`

HasLangCode returns a boolean if a field has been set.

### GetActive

`func (o *LanguageModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *LanguageModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *LanguageModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *LanguageModel) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


