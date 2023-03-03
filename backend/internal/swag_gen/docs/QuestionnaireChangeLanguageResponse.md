# QuestionnaireChangeLanguageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **[]int32** | List with answer id&#39;s succesful updated | [optional] 
**Failed** | Pointer to **[]int32** | List with answer id&#39;s failed to update | [optional] 
**NewLangCode** | Pointer to **string** | New Language Code | [optional] 
**OldLangCode** | Pointer to **string** | Old Language code | [optional] 

## Methods

### NewQuestionnaireChangeLanguageResponse

`func NewQuestionnaireChangeLanguageResponse() *QuestionnaireChangeLanguageResponse`

NewQuestionnaireChangeLanguageResponse instantiates a new QuestionnaireChangeLanguageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuestionnaireChangeLanguageResponseWithDefaults

`func NewQuestionnaireChangeLanguageResponseWithDefaults() *QuestionnaireChangeLanguageResponse`

NewQuestionnaireChangeLanguageResponseWithDefaults instantiates a new QuestionnaireChangeLanguageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *QuestionnaireChangeLanguageResponse) GetSuccess() []int32`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *QuestionnaireChangeLanguageResponse) GetSuccessOk() (*[]int32, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *QuestionnaireChangeLanguageResponse) SetSuccess(v []int32)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *QuestionnaireChangeLanguageResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetFailed

`func (o *QuestionnaireChangeLanguageResponse) GetFailed() []int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *QuestionnaireChangeLanguageResponse) GetFailedOk() (*[]int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *QuestionnaireChangeLanguageResponse) SetFailed(v []int32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *QuestionnaireChangeLanguageResponse) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetNewLangCode

`func (o *QuestionnaireChangeLanguageResponse) GetNewLangCode() string`

GetNewLangCode returns the NewLangCode field if non-nil, zero value otherwise.

### GetNewLangCodeOk

`func (o *QuestionnaireChangeLanguageResponse) GetNewLangCodeOk() (*string, bool)`

GetNewLangCodeOk returns a tuple with the NewLangCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewLangCode

`func (o *QuestionnaireChangeLanguageResponse) SetNewLangCode(v string)`

SetNewLangCode sets NewLangCode field to given value.

### HasNewLangCode

`func (o *QuestionnaireChangeLanguageResponse) HasNewLangCode() bool`

HasNewLangCode returns a boolean if a field has been set.

### GetOldLangCode

`func (o *QuestionnaireChangeLanguageResponse) GetOldLangCode() string`

GetOldLangCode returns the OldLangCode field if non-nil, zero value otherwise.

### GetOldLangCodeOk

`func (o *QuestionnaireChangeLanguageResponse) GetOldLangCodeOk() (*string, bool)`

GetOldLangCodeOk returns a tuple with the OldLangCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldLangCode

`func (o *QuestionnaireChangeLanguageResponse) SetOldLangCode(v string)`

SetOldLangCode sets OldLangCode field to given value.

### HasOldLangCode

`func (o *QuestionnaireChangeLanguageResponse) HasOldLangCode() bool`

HasOldLangCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


