# LogFiltersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to **string** | Graylog query, see Graylog docs for syntax (https://docs.graylog.org/en/3.1/pages/queries.html#time-frame-selector) | [optional] [readonly] 
**DateTimeFrom** | Pointer to **string** | Datetime from. When left empty, default value will be now -10 minutes | [optional] [readonly] 
**DateTimeTo** | Pointer to **string** | Datetime to. When left empty, default value will be now | [optional] [readonly] 

## Methods

### NewLogFiltersModel

`func NewLogFiltersModel() *LogFiltersModel`

NewLogFiltersModel instantiates a new LogFiltersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogFiltersModelWithDefaults

`func NewLogFiltersModelWithDefaults() *LogFiltersModel`

NewLogFiltersModelWithDefaults instantiates a new LogFiltersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *LogFiltersModel) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *LogFiltersModel) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *LogFiltersModel) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *LogFiltersModel) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetDateTimeFrom

`func (o *LogFiltersModel) GetDateTimeFrom() string`

GetDateTimeFrom returns the DateTimeFrom field if non-nil, zero value otherwise.

### GetDateTimeFromOk

`func (o *LogFiltersModel) GetDateTimeFromOk() (*string, bool)`

GetDateTimeFromOk returns a tuple with the DateTimeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFrom

`func (o *LogFiltersModel) SetDateTimeFrom(v string)`

SetDateTimeFrom sets DateTimeFrom field to given value.

### HasDateTimeFrom

`func (o *LogFiltersModel) HasDateTimeFrom() bool`

HasDateTimeFrom returns a boolean if a field has been set.

### GetDateTimeTo

`func (o *LogFiltersModel) GetDateTimeTo() string`

GetDateTimeTo returns the DateTimeTo field if non-nil, zero value otherwise.

### GetDateTimeToOk

`func (o *LogFiltersModel) GetDateTimeToOk() (*string, bool)`

GetDateTimeToOk returns a tuple with the DateTimeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeTo

`func (o *LogFiltersModel) SetDateTimeTo(v string)`

SetDateTimeTo sets DateTimeTo field to given value.

### HasDateTimeTo

`func (o *LogFiltersModel) HasDateTimeTo() bool`

HasDateTimeTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


