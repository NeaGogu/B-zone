# OpeningHoursRuleModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique ID | [optional] 
**AddressId** | Pointer to **int64** | Address ID associated with this OpeningHoursRule | [optional] 
**TimeFrom** | Pointer to **string** | We define time_from &#x3D; &#39;00:00&#39; and time_to &#x3D; &#39;00:00&#39; to mean no opening times, i.e. closed. | [optional] [default to "08:00"]
**TimeTo** | Pointer to **string** |  | [optional] [default to "18:00"]
**ValidDateFrom** | Pointer to **string** |  | [optional] [default to "1970-01-01"]
**ValidDateTo** | Pointer to **string** |  | [optional] [default to "2200-01-01"]
**Weekday** | Pointer to **int32** | value 7 represents all weekdays. 0 &#x3D; monday, etc. | [optional] 
**Precedence** | Pointer to **int32** | lowest precedence is 0. Accepted values are 0 and 1 right now. If more rules apply to a date, only rules with highest precendence will apply. | [optional] 

## Methods

### NewOpeningHoursRuleModel

`func NewOpeningHoursRuleModel() *OpeningHoursRuleModel`

NewOpeningHoursRuleModel instantiates a new OpeningHoursRuleModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpeningHoursRuleModelWithDefaults

`func NewOpeningHoursRuleModelWithDefaults() *OpeningHoursRuleModel`

NewOpeningHoursRuleModelWithDefaults instantiates a new OpeningHoursRuleModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpeningHoursRuleModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpeningHoursRuleModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpeningHoursRuleModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OpeningHoursRuleModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAddressId

`func (o *OpeningHoursRuleModel) GetAddressId() int64`

GetAddressId returns the AddressId field if non-nil, zero value otherwise.

### GetAddressIdOk

`func (o *OpeningHoursRuleModel) GetAddressIdOk() (*int64, bool)`

GetAddressIdOk returns a tuple with the AddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressId

`func (o *OpeningHoursRuleModel) SetAddressId(v int64)`

SetAddressId sets AddressId field to given value.

### HasAddressId

`func (o *OpeningHoursRuleModel) HasAddressId() bool`

HasAddressId returns a boolean if a field has been set.

### GetTimeFrom

`func (o *OpeningHoursRuleModel) GetTimeFrom() string`

GetTimeFrom returns the TimeFrom field if non-nil, zero value otherwise.

### GetTimeFromOk

`func (o *OpeningHoursRuleModel) GetTimeFromOk() (*string, bool)`

GetTimeFromOk returns a tuple with the TimeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrom

`func (o *OpeningHoursRuleModel) SetTimeFrom(v string)`

SetTimeFrom sets TimeFrom field to given value.

### HasTimeFrom

`func (o *OpeningHoursRuleModel) HasTimeFrom() bool`

HasTimeFrom returns a boolean if a field has been set.

### GetTimeTo

`func (o *OpeningHoursRuleModel) GetTimeTo() string`

GetTimeTo returns the TimeTo field if non-nil, zero value otherwise.

### GetTimeToOk

`func (o *OpeningHoursRuleModel) GetTimeToOk() (*string, bool)`

GetTimeToOk returns a tuple with the TimeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeTo

`func (o *OpeningHoursRuleModel) SetTimeTo(v string)`

SetTimeTo sets TimeTo field to given value.

### HasTimeTo

`func (o *OpeningHoursRuleModel) HasTimeTo() bool`

HasTimeTo returns a boolean if a field has been set.

### GetValidDateFrom

`func (o *OpeningHoursRuleModel) GetValidDateFrom() string`

GetValidDateFrom returns the ValidDateFrom field if non-nil, zero value otherwise.

### GetValidDateFromOk

`func (o *OpeningHoursRuleModel) GetValidDateFromOk() (*string, bool)`

GetValidDateFromOk returns a tuple with the ValidDateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidDateFrom

`func (o *OpeningHoursRuleModel) SetValidDateFrom(v string)`

SetValidDateFrom sets ValidDateFrom field to given value.

### HasValidDateFrom

`func (o *OpeningHoursRuleModel) HasValidDateFrom() bool`

HasValidDateFrom returns a boolean if a field has been set.

### GetValidDateTo

`func (o *OpeningHoursRuleModel) GetValidDateTo() string`

GetValidDateTo returns the ValidDateTo field if non-nil, zero value otherwise.

### GetValidDateToOk

`func (o *OpeningHoursRuleModel) GetValidDateToOk() (*string, bool)`

GetValidDateToOk returns a tuple with the ValidDateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidDateTo

`func (o *OpeningHoursRuleModel) SetValidDateTo(v string)`

SetValidDateTo sets ValidDateTo field to given value.

### HasValidDateTo

`func (o *OpeningHoursRuleModel) HasValidDateTo() bool`

HasValidDateTo returns a boolean if a field has been set.

### GetWeekday

`func (o *OpeningHoursRuleModel) GetWeekday() int32`

GetWeekday returns the Weekday field if non-nil, zero value otherwise.

### GetWeekdayOk

`func (o *OpeningHoursRuleModel) GetWeekdayOk() (*int32, bool)`

GetWeekdayOk returns a tuple with the Weekday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekday

`func (o *OpeningHoursRuleModel) SetWeekday(v int32)`

SetWeekday sets Weekday field to given value.

### HasWeekday

`func (o *OpeningHoursRuleModel) HasWeekday() bool`

HasWeekday returns a boolean if a field has been set.

### GetPrecedence

`func (o *OpeningHoursRuleModel) GetPrecedence() int32`

GetPrecedence returns the Precedence field if non-nil, zero value otherwise.

### GetPrecedenceOk

`func (o *OpeningHoursRuleModel) GetPrecedenceOk() (*int32, bool)`

GetPrecedenceOk returns a tuple with the Precedence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecedence

`func (o *OpeningHoursRuleModel) SetPrecedence(v int32)`

SetPrecedence sets Precedence field to given value.

### HasPrecedence

`func (o *OpeningHoursRuleModel) HasPrecedence() bool`

HasPrecedence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


