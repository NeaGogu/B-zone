# BlockedDateModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique Identifier | [optional] [readonly] 
**Description** | Pointer to **string** | Description of blocked date | [optional] [readonly] 
**Date** | Pointer to **time.Time** | updated_at date time | [optional] [readonly] 

## Methods

### NewBlockedDateModel

`func NewBlockedDateModel() *BlockedDateModel`

NewBlockedDateModel instantiates a new BlockedDateModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockedDateModelWithDefaults

`func NewBlockedDateModelWithDefaults() *BlockedDateModel`

NewBlockedDateModelWithDefaults instantiates a new BlockedDateModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BlockedDateModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlockedDateModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlockedDateModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *BlockedDateModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *BlockedDateModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BlockedDateModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BlockedDateModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BlockedDateModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDate

`func (o *BlockedDateModel) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *BlockedDateModel) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *BlockedDateModel) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *BlockedDateModel) HasDate() bool`

HasDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


