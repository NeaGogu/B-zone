# PauseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique Identifier | [optional] [readonly] 
**DriveTimeMode** | Pointer to **bool** | Determines if pause is a drivetime pause or a timewindow pause | [optional] 
**Name** | Pointer to **string** | Name of pause | [optional] [readonly] 
**InitialDrivingDuration** | Pointer to **int64** | initial driving time in minutes until first possibility of starting pause. (Only used in servicewindow pause) | [optional] 
**MaxDrivingDuration** | Pointer to **int64** | max driving time in minutes before a pause must be started | [optional] 
**PauseDuration** | Pointer to **int64** | (total) duration for pause(s) in minutes | [optional] 
**PossibleSplit** | Pointer to **[]int32** | A list of durations in minutes by which the total duration of the pause may be split. (Only used in servicewindow pause) | [optional] 
**EarliestTime** | Pointer to **string** | Earliest time. (Only used in drivetime pause) | [optional] 
**LatestTime** | Pointer to **string** | Latest time. (Only used in drivetime pause) | [optional] 
**Links** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**MetaData** | Pointer to [**[]MetaDataModel**](MetaDataModel.md) |  | [optional] 

## Methods

### NewPauseModel

`func NewPauseModel() *PauseModel`

NewPauseModel instantiates a new PauseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPauseModelWithDefaults

`func NewPauseModelWithDefaults() *PauseModel`

NewPauseModelWithDefaults instantiates a new PauseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PauseModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PauseModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PauseModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PauseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDriveTimeMode

`func (o *PauseModel) GetDriveTimeMode() bool`

GetDriveTimeMode returns the DriveTimeMode field if non-nil, zero value otherwise.

### GetDriveTimeModeOk

`func (o *PauseModel) GetDriveTimeModeOk() (*bool, bool)`

GetDriveTimeModeOk returns a tuple with the DriveTimeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveTimeMode

`func (o *PauseModel) SetDriveTimeMode(v bool)`

SetDriveTimeMode sets DriveTimeMode field to given value.

### HasDriveTimeMode

`func (o *PauseModel) HasDriveTimeMode() bool`

HasDriveTimeMode returns a boolean if a field has been set.

### GetName

`func (o *PauseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PauseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PauseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PauseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInitialDrivingDuration

`func (o *PauseModel) GetInitialDrivingDuration() int64`

GetInitialDrivingDuration returns the InitialDrivingDuration field if non-nil, zero value otherwise.

### GetInitialDrivingDurationOk

`func (o *PauseModel) GetInitialDrivingDurationOk() (*int64, bool)`

GetInitialDrivingDurationOk returns a tuple with the InitialDrivingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialDrivingDuration

`func (o *PauseModel) SetInitialDrivingDuration(v int64)`

SetInitialDrivingDuration sets InitialDrivingDuration field to given value.

### HasInitialDrivingDuration

`func (o *PauseModel) HasInitialDrivingDuration() bool`

HasInitialDrivingDuration returns a boolean if a field has been set.

### GetMaxDrivingDuration

`func (o *PauseModel) GetMaxDrivingDuration() int64`

GetMaxDrivingDuration returns the MaxDrivingDuration field if non-nil, zero value otherwise.

### GetMaxDrivingDurationOk

`func (o *PauseModel) GetMaxDrivingDurationOk() (*int64, bool)`

GetMaxDrivingDurationOk returns a tuple with the MaxDrivingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDrivingDuration

`func (o *PauseModel) SetMaxDrivingDuration(v int64)`

SetMaxDrivingDuration sets MaxDrivingDuration field to given value.

### HasMaxDrivingDuration

`func (o *PauseModel) HasMaxDrivingDuration() bool`

HasMaxDrivingDuration returns a boolean if a field has been set.

### GetPauseDuration

`func (o *PauseModel) GetPauseDuration() int64`

GetPauseDuration returns the PauseDuration field if non-nil, zero value otherwise.

### GetPauseDurationOk

`func (o *PauseModel) GetPauseDurationOk() (*int64, bool)`

GetPauseDurationOk returns a tuple with the PauseDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseDuration

`func (o *PauseModel) SetPauseDuration(v int64)`

SetPauseDuration sets PauseDuration field to given value.

### HasPauseDuration

`func (o *PauseModel) HasPauseDuration() bool`

HasPauseDuration returns a boolean if a field has been set.

### GetPossibleSplit

`func (o *PauseModel) GetPossibleSplit() []int32`

GetPossibleSplit returns the PossibleSplit field if non-nil, zero value otherwise.

### GetPossibleSplitOk

`func (o *PauseModel) GetPossibleSplitOk() (*[]int32, bool)`

GetPossibleSplitOk returns a tuple with the PossibleSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossibleSplit

`func (o *PauseModel) SetPossibleSplit(v []int32)`

SetPossibleSplit sets PossibleSplit field to given value.

### HasPossibleSplit

`func (o *PauseModel) HasPossibleSplit() bool`

HasPossibleSplit returns a boolean if a field has been set.

### GetEarliestTime

`func (o *PauseModel) GetEarliestTime() string`

GetEarliestTime returns the EarliestTime field if non-nil, zero value otherwise.

### GetEarliestTimeOk

`func (o *PauseModel) GetEarliestTimeOk() (*string, bool)`

GetEarliestTimeOk returns a tuple with the EarliestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestTime

`func (o *PauseModel) SetEarliestTime(v string)`

SetEarliestTime sets EarliestTime field to given value.

### HasEarliestTime

`func (o *PauseModel) HasEarliestTime() bool`

HasEarliestTime returns a boolean if a field has been set.

### GetLatestTime

`func (o *PauseModel) GetLatestTime() string`

GetLatestTime returns the LatestTime field if non-nil, zero value otherwise.

### GetLatestTimeOk

`func (o *PauseModel) GetLatestTimeOk() (*string, bool)`

GetLatestTimeOk returns a tuple with the LatestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestTime

`func (o *PauseModel) SetLatestTime(v string)`

SetLatestTime sets LatestTime field to given value.

### HasLatestTime

`func (o *PauseModel) HasLatestTime() bool`

HasLatestTime returns a boolean if a field has been set.

### GetLinks

`func (o *PauseModel) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PauseModel) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PauseModel) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PauseModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetaData

`func (o *PauseModel) GetMetaData() []MetaDataModel`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *PauseModel) GetMetaDataOk() (*[]MetaDataModel, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *PauseModel) SetMetaData(v []MetaDataModel)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *PauseModel) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


