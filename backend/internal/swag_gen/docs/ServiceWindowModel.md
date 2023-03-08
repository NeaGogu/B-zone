# ServiceWindowModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique ID | [optional] 
**ServiceWindowsSchemeId** | Pointer to **int64** | Service windows scheme ID | [optional] 
**WeekDay** | Pointer to **int64** | Week day | [optional] 
**TimeFrom** | Pointer to **string** | time from | [optional] 
**TimeTo** | Pointer to **string** | time to | [optional] 

## Methods

### NewServiceWindowModel

`func NewServiceWindowModel() *ServiceWindowModel`

NewServiceWindowModel instantiates a new ServiceWindowModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceWindowModelWithDefaults

`func NewServiceWindowModelWithDefaults() *ServiceWindowModel`

NewServiceWindowModelWithDefaults instantiates a new ServiceWindowModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceWindowModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceWindowModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceWindowModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceWindowModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetServiceWindowsSchemeId

`func (o *ServiceWindowModel) GetServiceWindowsSchemeId() int64`

GetServiceWindowsSchemeId returns the ServiceWindowsSchemeId field if non-nil, zero value otherwise.

### GetServiceWindowsSchemeIdOk

`func (o *ServiceWindowModel) GetServiceWindowsSchemeIdOk() (*int64, bool)`

GetServiceWindowsSchemeIdOk returns a tuple with the ServiceWindowsSchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceWindowsSchemeId

`func (o *ServiceWindowModel) SetServiceWindowsSchemeId(v int64)`

SetServiceWindowsSchemeId sets ServiceWindowsSchemeId field to given value.

### HasServiceWindowsSchemeId

`func (o *ServiceWindowModel) HasServiceWindowsSchemeId() bool`

HasServiceWindowsSchemeId returns a boolean if a field has been set.

### GetWeekDay

`func (o *ServiceWindowModel) GetWeekDay() int64`

GetWeekDay returns the WeekDay field if non-nil, zero value otherwise.

### GetWeekDayOk

`func (o *ServiceWindowModel) GetWeekDayOk() (*int64, bool)`

GetWeekDayOk returns a tuple with the WeekDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekDay

`func (o *ServiceWindowModel) SetWeekDay(v int64)`

SetWeekDay sets WeekDay field to given value.

### HasWeekDay

`func (o *ServiceWindowModel) HasWeekDay() bool`

HasWeekDay returns a boolean if a field has been set.

### GetTimeFrom

`func (o *ServiceWindowModel) GetTimeFrom() string`

GetTimeFrom returns the TimeFrom field if non-nil, zero value otherwise.

### GetTimeFromOk

`func (o *ServiceWindowModel) GetTimeFromOk() (*string, bool)`

GetTimeFromOk returns a tuple with the TimeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrom

`func (o *ServiceWindowModel) SetTimeFrom(v string)`

SetTimeFrom sets TimeFrom field to given value.

### HasTimeFrom

`func (o *ServiceWindowModel) HasTimeFrom() bool`

HasTimeFrom returns a boolean if a field has been set.

### GetTimeTo

`func (o *ServiceWindowModel) GetTimeTo() string`

GetTimeTo returns the TimeTo field if non-nil, zero value otherwise.

### GetTimeToOk

`func (o *ServiceWindowModel) GetTimeToOk() (*string, bool)`

GetTimeToOk returns a tuple with the TimeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeTo

`func (o *ServiceWindowModel) SetTimeTo(v string)`

SetTimeTo sets TimeTo field to given value.

### HasTimeTo

`func (o *ServiceWindowModel) HasTimeTo() bool`

HasTimeTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


