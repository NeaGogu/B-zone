# AssignmentFiltersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **[]int32** | Unique Identifier | [optional] 
**Nr** | Pointer to **[]string** | Number of this assignment | [optional] 
**DateTimeFrom** | Pointer to **time.Time** | DateTime From | [optional] 
**DateTimeTo** | Pointer to **time.Time** | DateTime To | [optional] 
**DateTimeFromSince** | Pointer to **time.Time** | filter assignments with a DateTime From since this input | [optional] 
**DateTimeFromTill** | Pointer to **time.Time** | filter assignments with a DateTime From till this input | [optional] 
**DateTimeToSince** | Pointer to **time.Time** | filter assignments with a DateTime To since this input | [optional] 
**DateTimeToTill** | Pointer to **time.Time** | filter assignments with a DateTime To till this input | [optional] 
**Links** | Pointer to **[]map[string]interface{}** | Activity Link ids | [optional] 
**UpdatedAtSince** | Pointer to **time.Time** | Show updated since | [optional] 
**UpdatedAtTill** | Pointer to **time.Time** | Show updated till | [optional] 

## Methods

### NewAssignmentFiltersModel

`func NewAssignmentFiltersModel() *AssignmentFiltersModel`

NewAssignmentFiltersModel instantiates a new AssignmentFiltersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignmentFiltersModelWithDefaults

`func NewAssignmentFiltersModelWithDefaults() *AssignmentFiltersModel`

NewAssignmentFiltersModelWithDefaults instantiates a new AssignmentFiltersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AssignmentFiltersModel) GetId() []int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssignmentFiltersModel) GetIdOk() (*[]int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssignmentFiltersModel) SetId(v []int32)`

SetId sets Id field to given value.

### HasId

`func (o *AssignmentFiltersModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNr

`func (o *AssignmentFiltersModel) GetNr() []string`

GetNr returns the Nr field if non-nil, zero value otherwise.

### GetNrOk

`func (o *AssignmentFiltersModel) GetNrOk() (*[]string, bool)`

GetNrOk returns a tuple with the Nr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNr

`func (o *AssignmentFiltersModel) SetNr(v []string)`

SetNr sets Nr field to given value.

### HasNr

`func (o *AssignmentFiltersModel) HasNr() bool`

HasNr returns a boolean if a field has been set.

### GetDateTimeFrom

`func (o *AssignmentFiltersModel) GetDateTimeFrom() time.Time`

GetDateTimeFrom returns the DateTimeFrom field if non-nil, zero value otherwise.

### GetDateTimeFromOk

`func (o *AssignmentFiltersModel) GetDateTimeFromOk() (*time.Time, bool)`

GetDateTimeFromOk returns a tuple with the DateTimeFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFrom

`func (o *AssignmentFiltersModel) SetDateTimeFrom(v time.Time)`

SetDateTimeFrom sets DateTimeFrom field to given value.

### HasDateTimeFrom

`func (o *AssignmentFiltersModel) HasDateTimeFrom() bool`

HasDateTimeFrom returns a boolean if a field has been set.

### GetDateTimeTo

`func (o *AssignmentFiltersModel) GetDateTimeTo() time.Time`

GetDateTimeTo returns the DateTimeTo field if non-nil, zero value otherwise.

### GetDateTimeToOk

`func (o *AssignmentFiltersModel) GetDateTimeToOk() (*time.Time, bool)`

GetDateTimeToOk returns a tuple with the DateTimeTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeTo

`func (o *AssignmentFiltersModel) SetDateTimeTo(v time.Time)`

SetDateTimeTo sets DateTimeTo field to given value.

### HasDateTimeTo

`func (o *AssignmentFiltersModel) HasDateTimeTo() bool`

HasDateTimeTo returns a boolean if a field has been set.

### GetDateTimeFromSince

`func (o *AssignmentFiltersModel) GetDateTimeFromSince() time.Time`

GetDateTimeFromSince returns the DateTimeFromSince field if non-nil, zero value otherwise.

### GetDateTimeFromSinceOk

`func (o *AssignmentFiltersModel) GetDateTimeFromSinceOk() (*time.Time, bool)`

GetDateTimeFromSinceOk returns a tuple with the DateTimeFromSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFromSince

`func (o *AssignmentFiltersModel) SetDateTimeFromSince(v time.Time)`

SetDateTimeFromSince sets DateTimeFromSince field to given value.

### HasDateTimeFromSince

`func (o *AssignmentFiltersModel) HasDateTimeFromSince() bool`

HasDateTimeFromSince returns a boolean if a field has been set.

### GetDateTimeFromTill

`func (o *AssignmentFiltersModel) GetDateTimeFromTill() time.Time`

GetDateTimeFromTill returns the DateTimeFromTill field if non-nil, zero value otherwise.

### GetDateTimeFromTillOk

`func (o *AssignmentFiltersModel) GetDateTimeFromTillOk() (*time.Time, bool)`

GetDateTimeFromTillOk returns a tuple with the DateTimeFromTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeFromTill

`func (o *AssignmentFiltersModel) SetDateTimeFromTill(v time.Time)`

SetDateTimeFromTill sets DateTimeFromTill field to given value.

### HasDateTimeFromTill

`func (o *AssignmentFiltersModel) HasDateTimeFromTill() bool`

HasDateTimeFromTill returns a boolean if a field has been set.

### GetDateTimeToSince

`func (o *AssignmentFiltersModel) GetDateTimeToSince() time.Time`

GetDateTimeToSince returns the DateTimeToSince field if non-nil, zero value otherwise.

### GetDateTimeToSinceOk

`func (o *AssignmentFiltersModel) GetDateTimeToSinceOk() (*time.Time, bool)`

GetDateTimeToSinceOk returns a tuple with the DateTimeToSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeToSince

`func (o *AssignmentFiltersModel) SetDateTimeToSince(v time.Time)`

SetDateTimeToSince sets DateTimeToSince field to given value.

### HasDateTimeToSince

`func (o *AssignmentFiltersModel) HasDateTimeToSince() bool`

HasDateTimeToSince returns a boolean if a field has been set.

### GetDateTimeToTill

`func (o *AssignmentFiltersModel) GetDateTimeToTill() time.Time`

GetDateTimeToTill returns the DateTimeToTill field if non-nil, zero value otherwise.

### GetDateTimeToTillOk

`func (o *AssignmentFiltersModel) GetDateTimeToTillOk() (*time.Time, bool)`

GetDateTimeToTillOk returns a tuple with the DateTimeToTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTimeToTill

`func (o *AssignmentFiltersModel) SetDateTimeToTill(v time.Time)`

SetDateTimeToTill sets DateTimeToTill field to given value.

### HasDateTimeToTill

`func (o *AssignmentFiltersModel) HasDateTimeToTill() bool`

HasDateTimeToTill returns a boolean if a field has been set.

### GetLinks

`func (o *AssignmentFiltersModel) GetLinks() []map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AssignmentFiltersModel) GetLinksOk() (*[]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AssignmentFiltersModel) SetLinks(v []map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AssignmentFiltersModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetUpdatedAtSince

`func (o *AssignmentFiltersModel) GetUpdatedAtSince() time.Time`

GetUpdatedAtSince returns the UpdatedAtSince field if non-nil, zero value otherwise.

### GetUpdatedAtSinceOk

`func (o *AssignmentFiltersModel) GetUpdatedAtSinceOk() (*time.Time, bool)`

GetUpdatedAtSinceOk returns a tuple with the UpdatedAtSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtSince

`func (o *AssignmentFiltersModel) SetUpdatedAtSince(v time.Time)`

SetUpdatedAtSince sets UpdatedAtSince field to given value.

### HasUpdatedAtSince

`func (o *AssignmentFiltersModel) HasUpdatedAtSince() bool`

HasUpdatedAtSince returns a boolean if a field has been set.

### GetUpdatedAtTill

`func (o *AssignmentFiltersModel) GetUpdatedAtTill() time.Time`

GetUpdatedAtTill returns the UpdatedAtTill field if non-nil, zero value otherwise.

### GetUpdatedAtTillOk

`func (o *AssignmentFiltersModel) GetUpdatedAtTillOk() (*time.Time, bool)`

GetUpdatedAtTillOk returns a tuple with the UpdatedAtTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtTill

`func (o *AssignmentFiltersModel) SetUpdatedAtTill(v time.Time)`

SetUpdatedAtTill sets UpdatedAtTill field to given value.

### HasUpdatedAtTill

`func (o *AssignmentFiltersModel) HasUpdatedAtTill() bool`

HasUpdatedAtTill returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


