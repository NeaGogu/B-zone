# RecurrenceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique Identifier | [optional] [readonly] 
**Name** | Pointer to **string** | Recurrence name | [optional] 
**TypeId** | Pointer to **int32** | recurrence type_id, 11:activity, 24:route | [optional] [readonly] 
**TypeName** | Pointer to **string** | recurrence type_name, activity, route | [optional] 
**StartDate** | Pointer to **time.Time** | Start date | [optional] [readonly] 
**EndDate** | Pointer to **time.Time** | End date | [optional] 
**EndOption** | Pointer to **string** | End option | [optional] 
**PeriodName** | Pointer to **string** | recurrence period_name, 1:day, 2:week, 3:month | [optional] 
**PeriodId** | Pointer to **int32** | recurrence period_id, 1:day, 2:week, 3:month | [optional] [readonly] 
**Frequency** | Pointer to **int64** | period frequency of recurrence (2 &#x3D; repeat each 2 days/weeks/months) | [optional] 
**Count** | Pointer to **int64** | nr of last recurrence to be created | [optional] 
**ShowAhead** | Pointer to **int64** | nr of recurrences to show ahead in system | [optional] 
**Current** | Pointer to **int64** | last created recurrence nr | [optional] [readonly] 
**Summary** | Pointer to **string** | summary of recurrence | [optional] [readonly] 
**NextRun** | Pointer to **time.Time** | Next date on which a new recurrence will be added | [optional] [readonly] 
**LastRun** | Pointer to **time.Time** | Last date on which a new recurrence was added | [optional] [readonly] 
**Active** | Pointer to **bool** | if active&#x3D;0: recurrence has been removed and is no longer visible in any bumbal interface | [optional] [readonly] 
**HasUncreatedObjects** | Pointer to **bool** | If the recurrence has uncreated objects | [optional] [readonly] 
**MetaData** | Pointer to [**[]MetaDataModel**](MetaDataModel.md) |  | [optional] 
**Base** | Pointer to **string** | Recurrence base | [optional] [readonly] 
**BaseDate** | Pointer to **string** | Recurrence base date | [optional] [readonly] 
**Monday** | Pointer to **bool** | Monday | [optional] [readonly] 
**Tuesday** | Pointer to **bool** | Tuesday | [optional] [readonly] 
**Wednesday** | Pointer to **bool** | Wednesday | [optional] [readonly] 
**Thursday** | Pointer to **bool** | Thursday | [optional] [readonly] 
**Friday** | Pointer to **bool** | Friday | [optional] [readonly] 
**Saturday** | Pointer to **bool** | Saturday | [optional] [readonly] 
**Sunday** | Pointer to **bool** | Sunday | [optional] [readonly] 
**MonthDay** | Pointer to **bool** | month day | [optional] [readonly] 
**MonthlyOption** | Pointer to **string** | Recurrence occurs every beginning or ending of the month | [optional] 
**ObjectId** | **int32** | ID of the object | 

## Methods

### NewRecurrenceModel

`func NewRecurrenceModel(objectId int32, ) *RecurrenceModel`

NewRecurrenceModel instantiates a new RecurrenceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurrenceModelWithDefaults

`func NewRecurrenceModelWithDefaults() *RecurrenceModel`

NewRecurrenceModelWithDefaults instantiates a new RecurrenceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RecurrenceModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RecurrenceModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RecurrenceModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RecurrenceModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RecurrenceModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecurrenceModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecurrenceModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecurrenceModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTypeId

`func (o *RecurrenceModel) GetTypeId() int32`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *RecurrenceModel) GetTypeIdOk() (*int32, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *RecurrenceModel) SetTypeId(v int32)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *RecurrenceModel) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetTypeName

`func (o *RecurrenceModel) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *RecurrenceModel) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *RecurrenceModel) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.

### HasTypeName

`func (o *RecurrenceModel) HasTypeName() bool`

HasTypeName returns a boolean if a field has been set.

### GetStartDate

`func (o *RecurrenceModel) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *RecurrenceModel) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *RecurrenceModel) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *RecurrenceModel) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *RecurrenceModel) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *RecurrenceModel) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *RecurrenceModel) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *RecurrenceModel) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetEndOption

`func (o *RecurrenceModel) GetEndOption() string`

GetEndOption returns the EndOption field if non-nil, zero value otherwise.

### GetEndOptionOk

`func (o *RecurrenceModel) GetEndOptionOk() (*string, bool)`

GetEndOptionOk returns a tuple with the EndOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOption

`func (o *RecurrenceModel) SetEndOption(v string)`

SetEndOption sets EndOption field to given value.

### HasEndOption

`func (o *RecurrenceModel) HasEndOption() bool`

HasEndOption returns a boolean if a field has been set.

### GetPeriodName

`func (o *RecurrenceModel) GetPeriodName() string`

GetPeriodName returns the PeriodName field if non-nil, zero value otherwise.

### GetPeriodNameOk

`func (o *RecurrenceModel) GetPeriodNameOk() (*string, bool)`

GetPeriodNameOk returns a tuple with the PeriodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodName

`func (o *RecurrenceModel) SetPeriodName(v string)`

SetPeriodName sets PeriodName field to given value.

### HasPeriodName

`func (o *RecurrenceModel) HasPeriodName() bool`

HasPeriodName returns a boolean if a field has been set.

### GetPeriodId

`func (o *RecurrenceModel) GetPeriodId() int32`

GetPeriodId returns the PeriodId field if non-nil, zero value otherwise.

### GetPeriodIdOk

`func (o *RecurrenceModel) GetPeriodIdOk() (*int32, bool)`

GetPeriodIdOk returns a tuple with the PeriodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodId

`func (o *RecurrenceModel) SetPeriodId(v int32)`

SetPeriodId sets PeriodId field to given value.

### HasPeriodId

`func (o *RecurrenceModel) HasPeriodId() bool`

HasPeriodId returns a boolean if a field has been set.

### GetFrequency

`func (o *RecurrenceModel) GetFrequency() int64`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *RecurrenceModel) GetFrequencyOk() (*int64, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *RecurrenceModel) SetFrequency(v int64)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *RecurrenceModel) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetCount

`func (o *RecurrenceModel) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *RecurrenceModel) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *RecurrenceModel) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *RecurrenceModel) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetShowAhead

`func (o *RecurrenceModel) GetShowAhead() int64`

GetShowAhead returns the ShowAhead field if non-nil, zero value otherwise.

### GetShowAheadOk

`func (o *RecurrenceModel) GetShowAheadOk() (*int64, bool)`

GetShowAheadOk returns a tuple with the ShowAhead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAhead

`func (o *RecurrenceModel) SetShowAhead(v int64)`

SetShowAhead sets ShowAhead field to given value.

### HasShowAhead

`func (o *RecurrenceModel) HasShowAhead() bool`

HasShowAhead returns a boolean if a field has been set.

### GetCurrent

`func (o *RecurrenceModel) GetCurrent() int64`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *RecurrenceModel) GetCurrentOk() (*int64, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *RecurrenceModel) SetCurrent(v int64)`

SetCurrent sets Current field to given value.

### HasCurrent

`func (o *RecurrenceModel) HasCurrent() bool`

HasCurrent returns a boolean if a field has been set.

### GetSummary

`func (o *RecurrenceModel) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *RecurrenceModel) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *RecurrenceModel) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *RecurrenceModel) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetNextRun

`func (o *RecurrenceModel) GetNextRun() time.Time`

GetNextRun returns the NextRun field if non-nil, zero value otherwise.

### GetNextRunOk

`func (o *RecurrenceModel) GetNextRunOk() (*time.Time, bool)`

GetNextRunOk returns a tuple with the NextRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRun

`func (o *RecurrenceModel) SetNextRun(v time.Time)`

SetNextRun sets NextRun field to given value.

### HasNextRun

`func (o *RecurrenceModel) HasNextRun() bool`

HasNextRun returns a boolean if a field has been set.

### GetLastRun

`func (o *RecurrenceModel) GetLastRun() time.Time`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *RecurrenceModel) GetLastRunOk() (*time.Time, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *RecurrenceModel) SetLastRun(v time.Time)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *RecurrenceModel) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetActive

`func (o *RecurrenceModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *RecurrenceModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *RecurrenceModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *RecurrenceModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetHasUncreatedObjects

`func (o *RecurrenceModel) GetHasUncreatedObjects() bool`

GetHasUncreatedObjects returns the HasUncreatedObjects field if non-nil, zero value otherwise.

### GetHasUncreatedObjectsOk

`func (o *RecurrenceModel) GetHasUncreatedObjectsOk() (*bool, bool)`

GetHasUncreatedObjectsOk returns a tuple with the HasUncreatedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUncreatedObjects

`func (o *RecurrenceModel) SetHasUncreatedObjects(v bool)`

SetHasUncreatedObjects sets HasUncreatedObjects field to given value.

### HasHasUncreatedObjects

`func (o *RecurrenceModel) HasHasUncreatedObjects() bool`

HasHasUncreatedObjects returns a boolean if a field has been set.

### GetMetaData

`func (o *RecurrenceModel) GetMetaData() []MetaDataModel`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *RecurrenceModel) GetMetaDataOk() (*[]MetaDataModel, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *RecurrenceModel) SetMetaData(v []MetaDataModel)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *RecurrenceModel) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetBase

`func (o *RecurrenceModel) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *RecurrenceModel) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *RecurrenceModel) SetBase(v string)`

SetBase sets Base field to given value.

### HasBase

`func (o *RecurrenceModel) HasBase() bool`

HasBase returns a boolean if a field has been set.

### GetBaseDate

`func (o *RecurrenceModel) GetBaseDate() string`

GetBaseDate returns the BaseDate field if non-nil, zero value otherwise.

### GetBaseDateOk

`func (o *RecurrenceModel) GetBaseDateOk() (*string, bool)`

GetBaseDateOk returns a tuple with the BaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseDate

`func (o *RecurrenceModel) SetBaseDate(v string)`

SetBaseDate sets BaseDate field to given value.

### HasBaseDate

`func (o *RecurrenceModel) HasBaseDate() bool`

HasBaseDate returns a boolean if a field has been set.

### GetMonday

`func (o *RecurrenceModel) GetMonday() bool`

GetMonday returns the Monday field if non-nil, zero value otherwise.

### GetMondayOk

`func (o *RecurrenceModel) GetMondayOk() (*bool, bool)`

GetMondayOk returns a tuple with the Monday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonday

`func (o *RecurrenceModel) SetMonday(v bool)`

SetMonday sets Monday field to given value.

### HasMonday

`func (o *RecurrenceModel) HasMonday() bool`

HasMonday returns a boolean if a field has been set.

### GetTuesday

`func (o *RecurrenceModel) GetTuesday() bool`

GetTuesday returns the Tuesday field if non-nil, zero value otherwise.

### GetTuesdayOk

`func (o *RecurrenceModel) GetTuesdayOk() (*bool, bool)`

GetTuesdayOk returns a tuple with the Tuesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuesday

`func (o *RecurrenceModel) SetTuesday(v bool)`

SetTuesday sets Tuesday field to given value.

### HasTuesday

`func (o *RecurrenceModel) HasTuesday() bool`

HasTuesday returns a boolean if a field has been set.

### GetWednesday

`func (o *RecurrenceModel) GetWednesday() bool`

GetWednesday returns the Wednesday field if non-nil, zero value otherwise.

### GetWednesdayOk

`func (o *RecurrenceModel) GetWednesdayOk() (*bool, bool)`

GetWednesdayOk returns a tuple with the Wednesday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWednesday

`func (o *RecurrenceModel) SetWednesday(v bool)`

SetWednesday sets Wednesday field to given value.

### HasWednesday

`func (o *RecurrenceModel) HasWednesday() bool`

HasWednesday returns a boolean if a field has been set.

### GetThursday

`func (o *RecurrenceModel) GetThursday() bool`

GetThursday returns the Thursday field if non-nil, zero value otherwise.

### GetThursdayOk

`func (o *RecurrenceModel) GetThursdayOk() (*bool, bool)`

GetThursdayOk returns a tuple with the Thursday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThursday

`func (o *RecurrenceModel) SetThursday(v bool)`

SetThursday sets Thursday field to given value.

### HasThursday

`func (o *RecurrenceModel) HasThursday() bool`

HasThursday returns a boolean if a field has been set.

### GetFriday

`func (o *RecurrenceModel) GetFriday() bool`

GetFriday returns the Friday field if non-nil, zero value otherwise.

### GetFridayOk

`func (o *RecurrenceModel) GetFridayOk() (*bool, bool)`

GetFridayOk returns a tuple with the Friday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriday

`func (o *RecurrenceModel) SetFriday(v bool)`

SetFriday sets Friday field to given value.

### HasFriday

`func (o *RecurrenceModel) HasFriday() bool`

HasFriday returns a boolean if a field has been set.

### GetSaturday

`func (o *RecurrenceModel) GetSaturday() bool`

GetSaturday returns the Saturday field if non-nil, zero value otherwise.

### GetSaturdayOk

`func (o *RecurrenceModel) GetSaturdayOk() (*bool, bool)`

GetSaturdayOk returns a tuple with the Saturday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturday

`func (o *RecurrenceModel) SetSaturday(v bool)`

SetSaturday sets Saturday field to given value.

### HasSaturday

`func (o *RecurrenceModel) HasSaturday() bool`

HasSaturday returns a boolean if a field has been set.

### GetSunday

`func (o *RecurrenceModel) GetSunday() bool`

GetSunday returns the Sunday field if non-nil, zero value otherwise.

### GetSundayOk

`func (o *RecurrenceModel) GetSundayOk() (*bool, bool)`

GetSundayOk returns a tuple with the Sunday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSunday

`func (o *RecurrenceModel) SetSunday(v bool)`

SetSunday sets Sunday field to given value.

### HasSunday

`func (o *RecurrenceModel) HasSunday() bool`

HasSunday returns a boolean if a field has been set.

### GetMonthDay

`func (o *RecurrenceModel) GetMonthDay() bool`

GetMonthDay returns the MonthDay field if non-nil, zero value otherwise.

### GetMonthDayOk

`func (o *RecurrenceModel) GetMonthDayOk() (*bool, bool)`

GetMonthDayOk returns a tuple with the MonthDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthDay

`func (o *RecurrenceModel) SetMonthDay(v bool)`

SetMonthDay sets MonthDay field to given value.

### HasMonthDay

`func (o *RecurrenceModel) HasMonthDay() bool`

HasMonthDay returns a boolean if a field has been set.

### GetMonthlyOption

`func (o *RecurrenceModel) GetMonthlyOption() string`

GetMonthlyOption returns the MonthlyOption field if non-nil, zero value otherwise.

### GetMonthlyOptionOk

`func (o *RecurrenceModel) GetMonthlyOptionOk() (*string, bool)`

GetMonthlyOptionOk returns a tuple with the MonthlyOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyOption

`func (o *RecurrenceModel) SetMonthlyOption(v string)`

SetMonthlyOption sets MonthlyOption field to given value.

### HasMonthlyOption

`func (o *RecurrenceModel) HasMonthlyOption() bool`

HasMonthlyOption returns a boolean if a field has been set.

### GetObjectId

`func (o *RecurrenceModel) GetObjectId() int32`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *RecurrenceModel) GetObjectIdOk() (*int32, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *RecurrenceModel) SetObjectId(v int32)`

SetObjectId sets ObjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


