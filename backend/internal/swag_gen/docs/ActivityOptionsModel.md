# ActivityOptionsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeActivityStatus** | Pointer to **bool** | Deprecated! Use include_status_name instead | [optional] 
**IncludeActivityTypeName** | Pointer to **bool** |  | [optional] 
**IncludeAddress** | Pointer to **bool** | Include the address in the addressbook | [optional] 
**IncludeAddressApplied** | Pointer to **bool** | Include the addres applied to the activity (not the original addres from the address book) | [optional] 
**IncludeAllowedDrivers** | Pointer to **bool** |  | [optional] 
**IncludeAllowedDriverIds** | Pointer to **bool** |  | [optional] 
**IncludeAllowedDriversLinks** | Pointer to **bool** |  | [optional] 
**IncludeAssignment** | Pointer to **bool** |  | [optional] 
**IncludeAssignmentNr** | Pointer to **bool** |  | [optional] 
**IncludeBrand** | Pointer to **bool** |  | [optional] 
**IncludeCoDrivers** | Pointer to **bool** |  | [optional] 
**IncludeCoDriverIds** | Pointer to **bool** |  | [optional] 
**IncludeCompartmentIds** | Pointer to **bool** |  | [optional] 
**IncludeCompartments** | Pointer to **bool** |  | [optional] 
**IncludeDriver** | Pointer to **bool** |  | [optional] 
**IncludeDriverInfo** | Pointer to **bool** |  | [optional] 
**IncludeDriverLinks** | Pointer to **bool** |  | [optional] 
**IncludeEmails** | Pointer to **bool** |  | [optional] 
**IncludeMetaData** | Pointer to **bool** |  | [optional] 
**IncludeNotes** | Pointer to **bool** |  | [optional] 
**IncludePackageLines** | Pointer to **bool** |  | [optional] 
**IncludePackageLineLinks** | Pointer to **bool** | Deprecated! Use include_links instead | [optional] 
**IncludePackageLinesInfo** | Pointer to **bool** |  | [optional] 
**IncludePhoneNrs** | Pointer to **bool** |  | [optional] 
**IncludeRoute** | Pointer to **bool** |  | [optional] 
**IncludeRouteInfo** | Pointer to **bool** |  | [optional] 
**IncludeStatusName** | Pointer to **bool** |  | [optional] 
**IncludeTimeSlots** | Pointer to **bool** |  | [optional] 
**IncludeTrailer** | Pointer to **bool** |  | [optional] 
**IncludeVehicle** | Pointer to **bool** |  | [optional] 
**IncludeCommunication** | Pointer to **bool** |  | [optional] 
**IncludeCommunicationObject** | Pointer to **bool** |  | [optional] 
**IncludeLinks** | Pointer to **bool** |  | [optional] 
**IncludeFiles** | Pointer to **bool** |  | [optional] 
**IncludeTags** | Pointer to **bool** |  | [optional] 
**IncludeTagTypeName** | Pointer to **bool** |  | [optional] 
**IncludeRecordInfo** | Pointer to **bool** |  | [optional] 
**IncludeRecordObject** | Pointer to **bool** |  | [optional] 
**IncludeTagNames** | Pointer to **bool** |  | [optional] 
**IncludeActivityNoteTags** | Pointer to **bool** | Deprecated! Use include_notes instead | [optional] 
**IncludeDepotActivity** | Pointer to **bool** |  | [optional] 
**IncludeDepotAddress** | Pointer to **bool** |  | [optional] 
**IncludeDepotAddressObject** | Pointer to **bool** | Deprecated! Use include_depot_address instead | [optional] 
**IncludeCapacityObject** | Pointer to **bool** |  | [optional] 
**IncludeCapacities** | Pointer to **bool** |  | [optional] 
**IncludeFilledCapacities** | Pointer to **bool** |  | [optional] 
**IncludeAppliedCapacities** | Pointer to **bool** |  | [optional] 
**IncludeAppliedCapacitiesV2** | Pointer to **bool** |  | [optional] 
**IncludeZones** | Pointer to **bool** |  | [optional] 
**IncludeBrandName** | Pointer to **bool** |  | [optional] 
**IncludeBrandColours** | Pointer to **bool** |  | [optional] 
**ApplyAddressBundling** | Pointer to **bool** |  | [optional] 
**IncludeBundledActivityIds** | Pointer to **bool** |  | [optional] 
**IncludePartyName** | Pointer to **bool** |  | [optional] 
**IncludeShipmentActivityNr** | Pointer to **bool** |  | [optional] 
**IncludePayments** | Pointer to **bool** |  | [optional] 
**IncludeTransactions** | Pointer to **bool** |  | [optional] 
**IncludeRelations** | Pointer to **bool** |  | [optional] 
**IncludeStats** | Pointer to **bool** |  | [optional] 
**IncludePackagelinesForWholeRouteOnStartActivity** | Pointer to **bool** |  | [optional] 
**IncludeActivityMetaData** | Pointer to **bool** | Deprecated! Use include_meta_data instead | [optional] 
**IncludeAddressObject** | Pointer to **bool** | Deprecated! Use include_adderess or include_address_applied instead | [optional] 
**IncludeCar** | Pointer to **bool** | Deprecated! Use include_vehicle instead | [optional] 
**IncludeActivityLinks** | Pointer to **bool** | Deprecated! Use include_links instead | [optional] 
**IncludeActivityFiles** | Pointer to **bool** | Deprecated! Use include_files instead | [optional] 
**IncludeBrandFiles** | Pointer to **bool** | Deprecated! Use include_files instead | [optional] 
**IncludeActivityFilesMetaData** | Pointer to **bool** | Deprecated! Use include_meta_data instead | [optional] 
**IncludeActivityRecordInfo** | Pointer to **bool** | Deprecated! Use include_record_info instead | [optional] 
**IncludeActivityRecordObject** | Pointer to **bool** | Deprecated! Use include_record_object instead | [optional] 
**IncludeActivityNotes** | Pointer to **bool** | Deprecated! Use include_notes instead | [optional] 
**IncludeActivityTags** | Pointer to **bool** | Deprecated! Use include_tags instead | [optional] 
**IncludeRecurrence** | Pointer to **bool** | Inlcude recurrence object | [optional] 
**IncludeMatchingRouteIds** | Pointer to **bool** | Include matching reoute ids | [optional] 

## Methods

### NewActivityOptionsModel

`func NewActivityOptionsModel() *ActivityOptionsModel`

NewActivityOptionsModel instantiates a new ActivityOptionsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityOptionsModelWithDefaults

`func NewActivityOptionsModelWithDefaults() *ActivityOptionsModel`

NewActivityOptionsModelWithDefaults instantiates a new ActivityOptionsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeActivityStatus

`func (o *ActivityOptionsModel) GetIncludeActivityStatus() bool`

GetIncludeActivityStatus returns the IncludeActivityStatus field if non-nil, zero value otherwise.

### GetIncludeActivityStatusOk

`func (o *ActivityOptionsModel) GetIncludeActivityStatusOk() (*bool, bool)`

GetIncludeActivityStatusOk returns a tuple with the IncludeActivityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeActivityStatus

`func (o *ActivityOptionsModel) SetIncludeActivityStatus(v bool)`

SetIncludeActivityStatus sets IncludeActivityStatus field to given value.

### HasIncludeActivityStatus

`func (o *ActivityOptionsModel) HasIncludeActivityStatus() bool`

HasIncludeActivityStatus returns a boolean if a field has been set.

### GetIncludeActivityTypeName

`func (o *ActivityOptionsModel) GetIncludeActivityTypeName() bool`

GetIncludeActivityTypeName returns the IncludeActivityTypeName field if non-nil, zero value otherwise.

### GetIncludeActivityTypeNameOk

`func (o *ActivityOptionsModel) GetIncludeActivityTypeNameOk() (*bool, bool)`

GetIncludeActivityTypeNameOk returns a tuple with the IncludeActivityTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeActivityTypeName

`func (o *ActivityOptionsModel) SetIncludeActivityTypeName(v bool)`

SetIncludeActivityTypeName sets IncludeActivityTypeName field to given value.

### HasIncludeActivityTypeName

`func (o *ActivityOptionsModel) HasIncludeActivityTypeName() bool`

HasIncludeActivityTypeName returns a boolean if a field has been set.

### GetIncludeAddress

`func (o *ActivityOptionsModel) GetIncludeAddress() bool`

GetIncludeAddress returns the IncludeAddress field if non-nil, zero value otherwise.

### GetIncludeAddressOk

`func (o *ActivityOptionsModel) GetIncludeAddressOk() (*bool, bool)`

GetIncludeAddressOk returns a tuple with the IncludeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAddress

`func (o *ActivityOptionsModel) SetIncludeAddress(v bool)`

SetIncludeAddress sets IncludeAddress field to given value.

### HasIncludeAddress

`func (o *ActivityOptionsModel) HasIncludeAddress() bool`

HasIncludeAddress returns a boolean if a field has been set.

### GetIncludeAddressApplied

`func (o *ActivityOptionsModel) GetIncludeAddressApplied() bool`

GetIncludeAddressApplied returns the IncludeAddressApplied field if non-nil, zero value otherwise.

### GetIncludeAddressAppliedOk

`func (o *ActivityOptionsModel) GetIncludeAddressAppliedOk() (*bool, bool)`

GetIncludeAddressAppliedOk returns a tuple with the IncludeAddressApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAddressApplied

`func (o *ActivityOptionsModel) SetIncludeAddressApplied(v bool)`

SetIncludeAddressApplied sets IncludeAddressApplied field to given value.

### HasIncludeAddressApplied

`func (o *ActivityOptionsModel) HasIncludeAddressApplied() bool`

HasIncludeAddressApplied returns a boolean if a field has been set.

### GetIncludeAllowedDrivers

`func (o *ActivityOptionsModel) GetIncludeAllowedDrivers() bool`

GetIncludeAllowedDrivers returns the IncludeAllowedDrivers field if non-nil, zero value otherwise.

### GetIncludeAllowedDriversOk

`func (o *ActivityOptionsModel) GetIncludeAllowedDriversOk() (*bool, bool)`

GetIncludeAllowedDriversOk returns a tuple with the IncludeAllowedDrivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllowedDrivers

`func (o *ActivityOptionsModel) SetIncludeAllowedDrivers(v bool)`

SetIncludeAllowedDrivers sets IncludeAllowedDrivers field to given value.

### HasIncludeAllowedDrivers

`func (o *ActivityOptionsModel) HasIncludeAllowedDrivers() bool`

HasIncludeAllowedDrivers returns a boolean if a field has been set.

### GetIncludeAllowedDriverIds

`func (o *ActivityOptionsModel) GetIncludeAllowedDriverIds() bool`

GetIncludeAllowedDriverIds returns the IncludeAllowedDriverIds field if non-nil, zero value otherwise.

### GetIncludeAllowedDriverIdsOk

`func (o *ActivityOptionsModel) GetIncludeAllowedDriverIdsOk() (*bool, bool)`

GetIncludeAllowedDriverIdsOk returns a tuple with the IncludeAllowedDriverIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllowedDriverIds

`func (o *ActivityOptionsModel) SetIncludeAllowedDriverIds(v bool)`

SetIncludeAllowedDriverIds sets IncludeAllowedDriverIds field to given value.

### HasIncludeAllowedDriverIds

`func (o *ActivityOptionsModel) HasIncludeAllowedDriverIds() bool`

HasIncludeAllowedDriverIds returns a boolean if a field has been set.

### GetIncludeAllowedDriversLinks

`func (o *ActivityOptionsModel) GetIncludeAllowedDriversLinks() bool`

GetIncludeAllowedDriversLinks returns the IncludeAllowedDriversLinks field if non-nil, zero value otherwise.

### GetIncludeAllowedDriversLinksOk

`func (o *ActivityOptionsModel) GetIncludeAllowedDriversLinksOk() (*bool, bool)`

GetIncludeAllowedDriversLinksOk returns a tuple with the IncludeAllowedDriversLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllowedDriversLinks

`func (o *ActivityOptionsModel) SetIncludeAllowedDriversLinks(v bool)`

SetIncludeAllowedDriversLinks sets IncludeAllowedDriversLinks field to given value.

### HasIncludeAllowedDriversLinks

`func (o *ActivityOptionsModel) HasIncludeAllowedDriversLinks() bool`

HasIncludeAllowedDriversLinks returns a boolean if a field has been set.

### GetIncludeAssignment

`func (o *ActivityOptionsModel) GetIncludeAssignment() bool`

GetIncludeAssignment returns the IncludeAssignment field if non-nil, zero value otherwise.

### GetIncludeAssignmentOk

`func (o *ActivityOptionsModel) GetIncludeAssignmentOk() (*bool, bool)`

GetIncludeAssignmentOk returns a tuple with the IncludeAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAssignment

`func (o *ActivityOptionsModel) SetIncludeAssignment(v bool)`

SetIncludeAssignment sets IncludeAssignment field to given value.

### HasIncludeAssignment

`func (o *ActivityOptionsModel) HasIncludeAssignment() bool`

HasIncludeAssignment returns a boolean if a field has been set.

### GetIncludeAssignmentNr

`func (o *ActivityOptionsModel) GetIncludeAssignmentNr() bool`

GetIncludeAssignmentNr returns the IncludeAssignmentNr field if non-nil, zero value otherwise.

### GetIncludeAssignmentNrOk

`func (o *ActivityOptionsModel) GetIncludeAssignmentNrOk() (*bool, bool)`

GetIncludeAssignmentNrOk returns a tuple with the IncludeAssignmentNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAssignmentNr

`func (o *ActivityOptionsModel) SetIncludeAssignmentNr(v bool)`

SetIncludeAssignmentNr sets IncludeAssignmentNr field to given value.

### HasIncludeAssignmentNr

`func (o *ActivityOptionsModel) HasIncludeAssignmentNr() bool`

HasIncludeAssignmentNr returns a boolean if a field has been set.

### GetIncludeBrand

`func (o *ActivityOptionsModel) GetIncludeBrand() bool`

GetIncludeBrand returns the IncludeBrand field if non-nil, zero value otherwise.

### GetIncludeBrandOk

`func (o *ActivityOptionsModel) GetIncludeBrandOk() (*bool, bool)`

GetIncludeBrandOk returns a tuple with the IncludeBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBrand

`func (o *ActivityOptionsModel) SetIncludeBrand(v bool)`

SetIncludeBrand sets IncludeBrand field to given value.

### HasIncludeBrand

`func (o *ActivityOptionsModel) HasIncludeBrand() bool`

HasIncludeBrand returns a boolean if a field has been set.

### GetIncludeCoDrivers

`func (o *ActivityOptionsModel) GetIncludeCoDrivers() bool`

GetIncludeCoDrivers returns the IncludeCoDrivers field if non-nil, zero value otherwise.

### GetIncludeCoDriversOk

`func (o *ActivityOptionsModel) GetIncludeCoDriversOk() (*bool, bool)`

GetIncludeCoDriversOk returns a tuple with the IncludeCoDrivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCoDrivers

`func (o *ActivityOptionsModel) SetIncludeCoDrivers(v bool)`

SetIncludeCoDrivers sets IncludeCoDrivers field to given value.

### HasIncludeCoDrivers

`func (o *ActivityOptionsModel) HasIncludeCoDrivers() bool`

HasIncludeCoDrivers returns a boolean if a field has been set.

### GetIncludeCoDriverIds

`func (o *ActivityOptionsModel) GetIncludeCoDriverIds() bool`

GetIncludeCoDriverIds returns the IncludeCoDriverIds field if non-nil, zero value otherwise.

### GetIncludeCoDriverIdsOk

`func (o *ActivityOptionsModel) GetIncludeCoDriverIdsOk() (*bool, bool)`

GetIncludeCoDriverIdsOk returns a tuple with the IncludeCoDriverIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCoDriverIds

`func (o *ActivityOptionsModel) SetIncludeCoDriverIds(v bool)`

SetIncludeCoDriverIds sets IncludeCoDriverIds field to given value.

### HasIncludeCoDriverIds

`func (o *ActivityOptionsModel) HasIncludeCoDriverIds() bool`

HasIncludeCoDriverIds returns a boolean if a field has been set.

### GetIncludeCompartmentIds

`func (o *ActivityOptionsModel) GetIncludeCompartmentIds() bool`

GetIncludeCompartmentIds returns the IncludeCompartmentIds field if non-nil, zero value otherwise.

### GetIncludeCompartmentIdsOk

`func (o *ActivityOptionsModel) GetIncludeCompartmentIdsOk() (*bool, bool)`

GetIncludeCompartmentIdsOk returns a tuple with the IncludeCompartmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCompartmentIds

`func (o *ActivityOptionsModel) SetIncludeCompartmentIds(v bool)`

SetIncludeCompartmentIds sets IncludeCompartmentIds field to given value.

### HasIncludeCompartmentIds

`func (o *ActivityOptionsModel) HasIncludeCompartmentIds() bool`

HasIncludeCompartmentIds returns a boolean if a field has been set.

### GetIncludeCompartments

`func (o *ActivityOptionsModel) GetIncludeCompartments() bool`

GetIncludeCompartments returns the IncludeCompartments field if non-nil, zero value otherwise.

### GetIncludeCompartmentsOk

`func (o *ActivityOptionsModel) GetIncludeCompartmentsOk() (*bool, bool)`

GetIncludeCompartmentsOk returns a tuple with the IncludeCompartments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCompartments

`func (o *ActivityOptionsModel) SetIncludeCompartments(v bool)`

SetIncludeCompartments sets IncludeCompartments field to given value.

### HasIncludeCompartments

`func (o *ActivityOptionsModel) HasIncludeCompartments() bool`

HasIncludeCompartments returns a boolean if a field has been set.

### GetIncludeDriver

`func (o *ActivityOptionsModel) GetIncludeDriver() bool`

GetIncludeDriver returns the IncludeDriver field if non-nil, zero value otherwise.

### GetIncludeDriverOk

`func (o *ActivityOptionsModel) GetIncludeDriverOk() (*bool, bool)`

GetIncludeDriverOk returns a tuple with the IncludeDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDriver

`func (o *ActivityOptionsModel) SetIncludeDriver(v bool)`

SetIncludeDriver sets IncludeDriver field to given value.

### HasIncludeDriver

`func (o *ActivityOptionsModel) HasIncludeDriver() bool`

HasIncludeDriver returns a boolean if a field has been set.

### GetIncludeDriverInfo

`func (o *ActivityOptionsModel) GetIncludeDriverInfo() bool`

GetIncludeDriverInfo returns the IncludeDriverInfo field if non-nil, zero value otherwise.

### GetIncludeDriverInfoOk

`func (o *ActivityOptionsModel) GetIncludeDriverInfoOk() (*bool, bool)`

GetIncludeDriverInfoOk returns a tuple with the IncludeDriverInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDriverInfo

`func (o *ActivityOptionsModel) SetIncludeDriverInfo(v bool)`

SetIncludeDriverInfo sets IncludeDriverInfo field to given value.

### HasIncludeDriverInfo

`func (o *ActivityOptionsModel) HasIncludeDriverInfo() bool`

HasIncludeDriverInfo returns a boolean if a field has been set.

### GetIncludeDriverLinks

`func (o *ActivityOptionsModel) GetIncludeDriverLinks() bool`

GetIncludeDriverLinks returns the IncludeDriverLinks field if non-nil, zero value otherwise.

### GetIncludeDriverLinksOk

`func (o *ActivityOptionsModel) GetIncludeDriverLinksOk() (*bool, bool)`

GetIncludeDriverLinksOk returns a tuple with the IncludeDriverLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDriverLinks

`func (o *ActivityOptionsModel) SetIncludeDriverLinks(v bool)`

SetIncludeDriverLinks sets IncludeDriverLinks field to given value.

### HasIncludeDriverLinks

`func (o *ActivityOptionsModel) HasIncludeDriverLinks() bool`

HasIncludeDriverLinks returns a boolean if a field has been set.

### GetIncludeEmails

`func (o *ActivityOptionsModel) GetIncludeEmails() bool`

GetIncludeEmails returns the IncludeEmails field if non-nil, zero value otherwise.

### GetIncludeEmailsOk

`func (o *ActivityOptionsModel) GetIncludeEmailsOk() (*bool, bool)`

GetIncludeEmailsOk returns a tuple with the IncludeEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeEmails

`func (o *ActivityOptionsModel) SetIncludeEmails(v bool)`

SetIncludeEmails sets IncludeEmails field to given value.

### HasIncludeEmails

`func (o *ActivityOptionsModel) HasIncludeEmails() bool`

HasIncludeEmails returns a boolean if a field has been set.

### GetIncludeMetaData

`func (o *ActivityOptionsModel) GetIncludeMetaData() bool`

GetIncludeMetaData returns the IncludeMetaData field if non-nil, zero value otherwise.

### GetIncludeMetaDataOk

`func (o *ActivityOptionsModel) GetIncludeMetaDataOk() (*bool, bool)`

GetIncludeMetaDataOk returns a tuple with the IncludeMetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMetaData

`func (o *ActivityOptionsModel) SetIncludeMetaData(v bool)`

SetIncludeMetaData sets IncludeMetaData field to given value.

### HasIncludeMetaData

`func (o *ActivityOptionsModel) HasIncludeMetaData() bool`

HasIncludeMetaData returns a boolean if a field has been set.

### GetIncludeNotes

`func (o *ActivityOptionsModel) GetIncludeNotes() bool`

GetIncludeNotes returns the IncludeNotes field if non-nil, zero value otherwise.

### GetIncludeNotesOk

`func (o *ActivityOptionsModel) GetIncludeNotesOk() (*bool, bool)`

GetIncludeNotesOk returns a tuple with the IncludeNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeNotes

`func (o *ActivityOptionsModel) SetIncludeNotes(v bool)`

SetIncludeNotes sets IncludeNotes field to given value.

### HasIncludeNotes

`func (o *ActivityOptionsModel) HasIncludeNotes() bool`

HasIncludeNotes returns a boolean if a field has been set.

### GetIncludePackageLines

`func (o *ActivityOptionsModel) GetIncludePackageLines() bool`

GetIncludePackageLines returns the IncludePackageLines field if non-nil, zero value otherwise.

### GetIncludePackageLinesOk

`func (o *ActivityOptionsModel) GetIncludePackageLinesOk() (*bool, bool)`

GetIncludePackageLinesOk returns a tuple with the IncludePackageLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePackageLines

`func (o *ActivityOptionsModel) SetIncludePackageLines(v bool)`

SetIncludePackageLines sets IncludePackageLines field to given value.

### HasIncludePackageLines

`func (o *ActivityOptionsModel) HasIncludePackageLines() bool`

HasIncludePackageLines returns a boolean if a field has been set.

### GetIncludePackageLineLinks

`func (o *ActivityOptionsModel) GetIncludePackageLineLinks() bool`

GetIncludePackageLineLinks returns the IncludePackageLineLinks field if non-nil, zero value otherwise.

### GetIncludePackageLineLinksOk

`func (o *ActivityOptionsModel) GetIncludePackageLineLinksOk() (*bool, bool)`

GetIncludePackageLineLinksOk returns a tuple with the IncludePackageLineLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePackageLineLinks

`func (o *ActivityOptionsModel) SetIncludePackageLineLinks(v bool)`

SetIncludePackageLineLinks sets IncludePackageLineLinks field to given value.

### HasIncludePackageLineLinks

`func (o *ActivityOptionsModel) HasIncludePackageLineLinks() bool`

HasIncludePackageLineLinks returns a boolean if a field has been set.

### GetIncludePackageLinesInfo

`func (o *ActivityOptionsModel) GetIncludePackageLinesInfo() bool`

GetIncludePackageLinesInfo returns the IncludePackageLinesInfo field if non-nil, zero value otherwise.

### GetIncludePackageLinesInfoOk

`func (o *ActivityOptionsModel) GetIncludePackageLinesInfoOk() (*bool, bool)`

GetIncludePackageLinesInfoOk returns a tuple with the IncludePackageLinesInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePackageLinesInfo

`func (o *ActivityOptionsModel) SetIncludePackageLinesInfo(v bool)`

SetIncludePackageLinesInfo sets IncludePackageLinesInfo field to given value.

### HasIncludePackageLinesInfo

`func (o *ActivityOptionsModel) HasIncludePackageLinesInfo() bool`

HasIncludePackageLinesInfo returns a boolean if a field has been set.

### GetIncludePhoneNrs

`func (o *ActivityOptionsModel) GetIncludePhoneNrs() bool`

GetIncludePhoneNrs returns the IncludePhoneNrs field if non-nil, zero value otherwise.

### GetIncludePhoneNrsOk

`func (o *ActivityOptionsModel) GetIncludePhoneNrsOk() (*bool, bool)`

GetIncludePhoneNrsOk returns a tuple with the IncludePhoneNrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePhoneNrs

`func (o *ActivityOptionsModel) SetIncludePhoneNrs(v bool)`

SetIncludePhoneNrs sets IncludePhoneNrs field to given value.

### HasIncludePhoneNrs

`func (o *ActivityOptionsModel) HasIncludePhoneNrs() bool`

HasIncludePhoneNrs returns a boolean if a field has been set.

### GetIncludeRoute

`func (o *ActivityOptionsModel) GetIncludeRoute() bool`

GetIncludeRoute returns the IncludeRoute field if non-nil, zero value otherwise.

### GetIncludeRouteOk

`func (o *ActivityOptionsModel) GetIncludeRouteOk() (*bool, bool)`

GetIncludeRouteOk returns a tuple with the IncludeRoute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRoute

`func (o *ActivityOptionsModel) SetIncludeRoute(v bool)`

SetIncludeRoute sets IncludeRoute field to given value.

### HasIncludeRoute

`func (o *ActivityOptionsModel) HasIncludeRoute() bool`

HasIncludeRoute returns a boolean if a field has been set.

### GetIncludeRouteInfo

`func (o *ActivityOptionsModel) GetIncludeRouteInfo() bool`

GetIncludeRouteInfo returns the IncludeRouteInfo field if non-nil, zero value otherwise.

### GetIncludeRouteInfoOk

`func (o *ActivityOptionsModel) GetIncludeRouteInfoOk() (*bool, bool)`

GetIncludeRouteInfoOk returns a tuple with the IncludeRouteInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRouteInfo

`func (o *ActivityOptionsModel) SetIncludeRouteInfo(v bool)`

SetIncludeRouteInfo sets IncludeRouteInfo field to given value.

### HasIncludeRouteInfo

`func (o *ActivityOptionsModel) HasIncludeRouteInfo() bool`

HasIncludeRouteInfo returns a boolean if a field has been set.

### GetIncludeStatusName

`func (o *ActivityOptionsModel) GetIncludeStatusName() bool`

GetIncludeStatusName returns the IncludeStatusName field if non-nil, zero value otherwise.

### GetIncludeStatusNameOk

`func (o *ActivityOptionsModel) GetIncludeStatusNameOk() (*bool, bool)`

GetIncludeStatusNameOk returns a tuple with the IncludeStatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeStatusName

`func (o *ActivityOptionsModel) SetIncludeStatusName(v bool)`

SetIncludeStatusName sets IncludeStatusName field to given value.

### HasIncludeStatusName

`func (o *ActivityOptionsModel) HasIncludeStatusName() bool`

HasIncludeStatusName returns a boolean if a field has been set.

### GetIncludeTimeSlots

`func (o *ActivityOptionsModel) GetIncludeTimeSlots() bool`

GetIncludeTimeSlots returns the IncludeTimeSlots field if non-nil, zero value otherwise.

### GetIncludeTimeSlotsOk

`func (o *ActivityOptionsModel) GetIncludeTimeSlotsOk() (*bool, bool)`

GetIncludeTimeSlotsOk returns a tuple with the IncludeTimeSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTimeSlots

`func (o *ActivityOptionsModel) SetIncludeTimeSlots(v bool)`

SetIncludeTimeSlots sets IncludeTimeSlots field to given value.

### HasIncludeTimeSlots

`func (o *ActivityOptionsModel) HasIncludeTimeSlots() bool`

HasIncludeTimeSlots returns a boolean if a field has been set.

### GetIncludeTrailer

`func (o *ActivityOptionsModel) GetIncludeTrailer() bool`

GetIncludeTrailer returns the IncludeTrailer field if non-nil, zero value otherwise.

### GetIncludeTrailerOk

`func (o *ActivityOptionsModel) GetIncludeTrailerOk() (*bool, bool)`

GetIncludeTrailerOk returns a tuple with the IncludeTrailer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTrailer

`func (o *ActivityOptionsModel) SetIncludeTrailer(v bool)`

SetIncludeTrailer sets IncludeTrailer field to given value.

### HasIncludeTrailer

`func (o *ActivityOptionsModel) HasIncludeTrailer() bool`

HasIncludeTrailer returns a boolean if a field has been set.

### GetIncludeVehicle

`func (o *ActivityOptionsModel) GetIncludeVehicle() bool`

GetIncludeVehicle returns the IncludeVehicle field if non-nil, zero value otherwise.

### GetIncludeVehicleOk

`func (o *ActivityOptionsModel) GetIncludeVehicleOk() (*bool, bool)`

GetIncludeVehicleOk returns a tuple with the IncludeVehicle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeVehicle

`func (o *ActivityOptionsModel) SetIncludeVehicle(v bool)`

SetIncludeVehicle sets IncludeVehicle field to given value.

### HasIncludeVehicle

`func (o *ActivityOptionsModel) HasIncludeVehicle() bool`

HasIncludeVehicle returns a boolean if a field has been set.

### GetIncludeCommunication

`func (o *ActivityOptionsModel) GetIncludeCommunication() bool`

GetIncludeCommunication returns the IncludeCommunication field if non-nil, zero value otherwise.

### GetIncludeCommunicationOk

`func (o *ActivityOptionsModel) GetIncludeCommunicationOk() (*bool, bool)`

GetIncludeCommunicationOk returns a tuple with the IncludeCommunication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCommunication

`func (o *ActivityOptionsModel) SetIncludeCommunication(v bool)`

SetIncludeCommunication sets IncludeCommunication field to given value.

### HasIncludeCommunication

`func (o *ActivityOptionsModel) HasIncludeCommunication() bool`

HasIncludeCommunication returns a boolean if a field has been set.

### GetIncludeCommunicationObject

`func (o *ActivityOptionsModel) GetIncludeCommunicationObject() bool`

GetIncludeCommunicationObject returns the IncludeCommunicationObject field if non-nil, zero value otherwise.

### GetIncludeCommunicationObjectOk

`func (o *ActivityOptionsModel) GetIncludeCommunicationObjectOk() (*bool, bool)`

GetIncludeCommunicationObjectOk returns a tuple with the IncludeCommunicationObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCommunicationObject

`func (o *ActivityOptionsModel) SetIncludeCommunicationObject(v bool)`

SetIncludeCommunicationObject sets IncludeCommunicationObject field to given value.

### HasIncludeCommunicationObject

`func (o *ActivityOptionsModel) HasIncludeCommunicationObject() bool`

HasIncludeCommunicationObject returns a boolean if a field has been set.

### GetIncludeLinks

`func (o *ActivityOptionsModel) GetIncludeLinks() bool`

GetIncludeLinks returns the IncludeLinks field if non-nil, zero value otherwise.

### GetIncludeLinksOk

`func (o *ActivityOptionsModel) GetIncludeLinksOk() (*bool, bool)`

GetIncludeLinksOk returns a tuple with the IncludeLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLinks

`func (o *ActivityOptionsModel) SetIncludeLinks(v bool)`

SetIncludeLinks sets IncludeLinks field to given value.

### HasIncludeLinks

`func (o *ActivityOptionsModel) HasIncludeLinks() bool`

HasIncludeLinks returns a boolean if a field has been set.

### GetIncludeFiles

`func (o *ActivityOptionsModel) GetIncludeFiles() bool`

GetIncludeFiles returns the IncludeFiles field if non-nil, zero value otherwise.

### GetIncludeFilesOk

`func (o *ActivityOptionsModel) GetIncludeFilesOk() (*bool, bool)`

GetIncludeFilesOk returns a tuple with the IncludeFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFiles

`func (o *ActivityOptionsModel) SetIncludeFiles(v bool)`

SetIncludeFiles sets IncludeFiles field to given value.

### HasIncludeFiles

`func (o *ActivityOptionsModel) HasIncludeFiles() bool`

HasIncludeFiles returns a boolean if a field has been set.

### GetIncludeTags

`func (o *ActivityOptionsModel) GetIncludeTags() bool`

GetIncludeTags returns the IncludeTags field if non-nil, zero value otherwise.

### GetIncludeTagsOk

`func (o *ActivityOptionsModel) GetIncludeTagsOk() (*bool, bool)`

GetIncludeTagsOk returns a tuple with the IncludeTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTags

`func (o *ActivityOptionsModel) SetIncludeTags(v bool)`

SetIncludeTags sets IncludeTags field to given value.

### HasIncludeTags

`func (o *ActivityOptionsModel) HasIncludeTags() bool`

HasIncludeTags returns a boolean if a field has been set.

### GetIncludeTagTypeName

`func (o *ActivityOptionsModel) GetIncludeTagTypeName() bool`

GetIncludeTagTypeName returns the IncludeTagTypeName field if non-nil, zero value otherwise.

### GetIncludeTagTypeNameOk

`func (o *ActivityOptionsModel) GetIncludeTagTypeNameOk() (*bool, bool)`

GetIncludeTagTypeNameOk returns a tuple with the IncludeTagTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTagTypeName

`func (o *ActivityOptionsModel) SetIncludeTagTypeName(v bool)`

SetIncludeTagTypeName sets IncludeTagTypeName field to given value.

### HasIncludeTagTypeName

`func (o *ActivityOptionsModel) HasIncludeTagTypeName() bool`

HasIncludeTagTypeName returns a boolean if a field has been set.

### GetIncludeRecordInfo

`func (o *ActivityOptionsModel) GetIncludeRecordInfo() bool`

GetIncludeRecordInfo returns the IncludeRecordInfo field if non-nil, zero value otherwise.

### GetIncludeRecordInfoOk

`func (o *ActivityOptionsModel) GetIncludeRecordInfoOk() (*bool, bool)`

GetIncludeRecordInfoOk returns a tuple with the IncludeRecordInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRecordInfo

`func (o *ActivityOptionsModel) SetIncludeRecordInfo(v bool)`

SetIncludeRecordInfo sets IncludeRecordInfo field to given value.

### HasIncludeRecordInfo

`func (o *ActivityOptionsModel) HasIncludeRecordInfo() bool`

HasIncludeRecordInfo returns a boolean if a field has been set.

### GetIncludeRecordObject

`func (o *ActivityOptionsModel) GetIncludeRecordObject() bool`

GetIncludeRecordObject returns the IncludeRecordObject field if non-nil, zero value otherwise.

### GetIncludeRecordObjectOk

`func (o *ActivityOptionsModel) GetIncludeRecordObjectOk() (*bool, bool)`

GetIncludeRecordObjectOk returns a tuple with the IncludeRecordObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRecordObject

`func (o *ActivityOptionsModel) SetIncludeRecordObject(v bool)`

SetIncludeRecordObject sets IncludeRecordObject field to given value.

### HasIncludeRecordObject

`func (o *ActivityOptionsModel) HasIncludeRecordObject() bool`

HasIncludeRecordObject returns a boolean if a field has been set.

### GetIncludeTagNames

`func (o *ActivityOptionsModel) GetIncludeTagNames() bool`

GetIncludeTagNames returns the IncludeTagNames field if non-nil, zero value otherwise.

### GetIncludeTagNamesOk

`func (o *ActivityOptionsModel) GetIncludeTagNamesOk() (*bool, bool)`

GetIncludeTagNamesOk returns a tuple with the IncludeTagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTagNames

`func (o *ActivityOptionsModel) SetIncludeTagNames(v bool)`

SetIncludeTagNames sets IncludeTagNames field to given value.

### HasIncludeTagNames

`func (o *ActivityOptionsModel) HasIncludeTagNames() bool`

HasIncludeTagNames returns a boolean if a field has been set.

### GetIncludeActivityNoteTags

`func (o *ActivityOptionsModel) GetIncludeActivityNoteTags() bool`

GetIncludeActivityNoteTags returns the IncludeActivityNoteTags field if non-nil, zero value otherwise.

### GetIncludeActivityNoteTagsOk

`func (o *ActivityOptionsModel) GetIncludeActivityNoteTagsOk() (*bool, bool)`

GetIncludeActivityNoteTagsOk returns a tuple with the IncludeActivityNoteTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeActivityNoteTags

`func (o *ActivityOptionsModel) SetIncludeActivityNoteTags(v bool)`

SetIncludeActivityNoteTags sets IncludeActivityNoteTags field to given value.

### HasIncludeActivityNoteTags

`func (o *ActivityOptionsModel) HasIncludeActivityNoteTags() bool`

HasIncludeActivityNoteTags returns a boolean if a field has been set.

### GetIncludeDepotActivity

`func (o *ActivityOptionsModel) GetIncludeDepotActivity() bool`

GetIncludeDepotActivity returns the IncludeDepotActivity field if non-nil, zero value otherwise.

### GetIncludeDepotActivityOk

`func (o *ActivityOptionsModel) GetIncludeDepotActivityOk() (*bool, bool)`

GetIncludeDepotActivityOk returns a tuple with the IncludeDepotActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDepotActivity

`func (o *ActivityOptionsModel) SetIncludeDepotActivity(v bool)`

SetIncludeDepotActivity sets IncludeDepotActivity field to given value.

### HasIncludeDepotActivity

`func (o *ActivityOptionsModel) HasIncludeDepotActivity() bool`

HasIncludeDepotActivity returns a boolean if a field has been set.

### GetIncludeDepotAddress

`func (o *ActivityOptionsModel) GetIncludeDepotAddress() bool`

GetIncludeDepotAddress returns the IncludeDepotAddress field if non-nil, zero value otherwise.

### GetIncludeDepotAddressOk

`func (o *ActivityOptionsModel) GetIncludeDepotAddressOk() (*bool, bool)`

GetIncludeDepotAddressOk returns a tuple with the IncludeDepotAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDepotAddress

`func (o *ActivityOptionsModel) SetIncludeDepotAddress(v bool)`

SetIncludeDepotAddress sets IncludeDepotAddress field to given value.

### HasIncludeDepotAddress

`func (o *ActivityOptionsModel) HasIncludeDepotAddress() bool`

HasIncludeDepotAddress returns a boolean if a field has been set.

### GetIncludeDepotAddressObject

`func (o *ActivityOptionsModel) GetIncludeDepotAddressObject() bool`

GetIncludeDepotAddressObject returns the IncludeDepotAddressObject field if non-nil, zero value otherwise.

### GetIncludeDepotAddressObjectOk

`func (o *ActivityOptionsModel) GetIncludeDepotAddressObjectOk() (*bool, bool)`

GetIncludeDepotAddressObjectOk returns a tuple with the IncludeDepotAddressObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDepotAddressObject

`func (o *ActivityOptionsModel) SetIncludeDepotAddressObject(v bool)`

SetIncludeDepotAddressObject sets IncludeDepotAddressObject field to given value.

### HasIncludeDepotAddressObject

`func (o *ActivityOptionsModel) HasIncludeDepotAddressObject() bool`

HasIncludeDepotAddressObject returns a boolean if a field has been set.

### GetIncludeCapacityObject

`func (o *ActivityOptionsModel) GetIncludeCapacityObject() bool`

GetIncludeCapacityObject returns the IncludeCapacityObject field if non-nil, zero value otherwise.

### GetIncludeCapacityObjectOk

`func (o *ActivityOptionsModel) GetIncludeCapacityObjectOk() (*bool, bool)`

GetIncludeCapacityObjectOk returns a tuple with the IncludeCapacityObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCapacityObject

`func (o *ActivityOptionsModel) SetIncludeCapacityObject(v bool)`

SetIncludeCapacityObject sets IncludeCapacityObject field to given value.

### HasIncludeCapacityObject

`func (o *ActivityOptionsModel) HasIncludeCapacityObject() bool`

HasIncludeCapacityObject returns a boolean if a field has been set.

### GetIncludeCapacities

`func (o *ActivityOptionsModel) GetIncludeCapacities() bool`

GetIncludeCapacities returns the IncludeCapacities field if non-nil, zero value otherwise.

### GetIncludeCapacitiesOk

`func (o *ActivityOptionsModel) GetIncludeCapacitiesOk() (*bool, bool)`

GetIncludeCapacitiesOk returns a tuple with the IncludeCapacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCapacities

`func (o *ActivityOptionsModel) SetIncludeCapacities(v bool)`

SetIncludeCapacities sets IncludeCapacities field to given value.

### HasIncludeCapacities

`func (o *ActivityOptionsModel) HasIncludeCapacities() bool`

HasIncludeCapacities returns a boolean if a field has been set.

### GetIncludeFilledCapacities

`func (o *ActivityOptionsModel) GetIncludeFilledCapacities() bool`

GetIncludeFilledCapacities returns the IncludeFilledCapacities field if non-nil, zero value otherwise.

### GetIncludeFilledCapacitiesOk

`func (o *ActivityOptionsModel) GetIncludeFilledCapacitiesOk() (*bool, bool)`

GetIncludeFilledCapacitiesOk returns a tuple with the IncludeFilledCapacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFilledCapacities

`func (o *ActivityOptionsModel) SetIncludeFilledCapacities(v bool)`

SetIncludeFilledCapacities sets IncludeFilledCapacities field to given value.

### HasIncludeFilledCapacities

`func (o *ActivityOptionsModel) HasIncludeFilledCapacities() bool`

HasIncludeFilledCapacities returns a boolean if a field has been set.

### GetIncludeAppliedCapacities

`func (o *ActivityOptionsModel) GetIncludeAppliedCapacities() bool`

GetIncludeAppliedCapacities returns the IncludeAppliedCapacities field if non-nil, zero value otherwise.

### GetIncludeAppliedCapacitiesOk

`func (o *ActivityOptionsModel) GetIncludeAppliedCapacitiesOk() (*bool, bool)`

GetIncludeAppliedCapacitiesOk returns a tuple with the IncludeAppliedCapacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAppliedCapacities

`func (o *ActivityOptionsModel) SetIncludeAppliedCapacities(v bool)`

SetIncludeAppliedCapacities sets IncludeAppliedCapacities field to given value.

### HasIncludeAppliedCapacities

`func (o *ActivityOptionsModel) HasIncludeAppliedCapacities() bool`

HasIncludeAppliedCapacities returns a boolean if a field has been set.

### GetIncludeAppliedCapacitiesV2

`func (o *ActivityOptionsModel) GetIncludeAppliedCapacitiesV2() bool`

GetIncludeAppliedCapacitiesV2 returns the IncludeAppliedCapacitiesV2 field if non-nil, zero value otherwise.

### GetIncludeAppliedCapacitiesV2Ok

`func (o *ActivityOptionsModel) GetIncludeAppliedCapacitiesV2Ok() (*bool, bool)`

GetIncludeAppliedCapacitiesV2Ok returns a tuple with the IncludeAppliedCapacitiesV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAppliedCapacitiesV2

`func (o *ActivityOptionsModel) SetIncludeAppliedCapacitiesV2(v bool)`

SetIncludeAppliedCapacitiesV2 sets IncludeAppliedCapacitiesV2 field to given value.

### HasIncludeAppliedCapacitiesV2

`func (o *ActivityOptionsModel) HasIncludeAppliedCapacitiesV2() bool`

HasIncludeAppliedCapacitiesV2 returns a boolean if a field has been set.

### GetIncludeZones

`func (o *ActivityOptionsModel) GetIncludeZones() bool`

GetIncludeZones returns the IncludeZones field if non-nil, zero value otherwise.

### GetIncludeZonesOk

`func (o *ActivityOptionsModel) GetIncludeZonesOk() (*bool, bool)`

GetIncludeZonesOk returns a tuple with the IncludeZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeZones

`func (o *ActivityOptionsModel) SetIncludeZones(v bool)`

SetIncludeZones sets IncludeZones field to given value.

### HasIncludeZones

`func (o *ActivityOptionsModel) HasIncludeZones() bool`

HasIncludeZones returns a boolean if a field has been set.

### GetIncludeBrandName

`func (o *ActivityOptionsModel) GetIncludeBrandName() bool`

GetIncludeBrandName returns the IncludeBrandName field if non-nil, zero value otherwise.

### GetIncludeBrandNameOk

`func (o *ActivityOptionsModel) GetIncludeBrandNameOk() (*bool, bool)`

GetIncludeBrandNameOk returns a tuple with the IncludeBrandName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBrandName

`func (o *ActivityOptionsModel) SetIncludeBrandName(v bool)`

SetIncludeBrandName sets IncludeBrandName field to given value.

### HasIncludeBrandName

`func (o *ActivityOptionsModel) HasIncludeBrandName() bool`

HasIncludeBrandName returns a boolean if a field has been set.

### GetIncludeBrandColours

`func (o *ActivityOptionsModel) GetIncludeBrandColours() bool`

GetIncludeBrandColours returns the IncludeBrandColours field if non-nil, zero value otherwise.

### GetIncludeBrandColoursOk

`func (o *ActivityOptionsModel) GetIncludeBrandColoursOk() (*bool, bool)`

GetIncludeBrandColoursOk returns a tuple with the IncludeBrandColours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBrandColours

`func (o *ActivityOptionsModel) SetIncludeBrandColours(v bool)`

SetIncludeBrandColours sets IncludeBrandColours field to given value.

### HasIncludeBrandColours

`func (o *ActivityOptionsModel) HasIncludeBrandColours() bool`

HasIncludeBrandColours returns a boolean if a field has been set.

### GetApplyAddressBundling

`func (o *ActivityOptionsModel) GetApplyAddressBundling() bool`

GetApplyAddressBundling returns the ApplyAddressBundling field if non-nil, zero value otherwise.

### GetApplyAddressBundlingOk

`func (o *ActivityOptionsModel) GetApplyAddressBundlingOk() (*bool, bool)`

GetApplyAddressBundlingOk returns a tuple with the ApplyAddressBundling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyAddressBundling

`func (o *ActivityOptionsModel) SetApplyAddressBundling(v bool)`

SetApplyAddressBundling sets ApplyAddressBundling field to given value.

### HasApplyAddressBundling

`func (o *ActivityOptionsModel) HasApplyAddressBundling() bool`

HasApplyAddressBundling returns a boolean if a field has been set.

### GetIncludeBundledActivityIds

`func (o *ActivityOptionsModel) GetIncludeBundledActivityIds() bool`

GetIncludeBundledActivityIds returns the IncludeBundledActivityIds field if non-nil, zero value otherwise.

### GetIncludeBundledActivityIdsOk

`func (o *ActivityOptionsModel) GetIncludeBundledActivityIdsOk() (*bool, bool)`

GetIncludeBundledActivityIdsOk returns a tuple with the IncludeBundledActivityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBundledActivityIds

`func (o *ActivityOptionsModel) SetIncludeBundledActivityIds(v bool)`

SetIncludeBundledActivityIds sets IncludeBundledActivityIds field to given value.

### HasIncludeBundledActivityIds

`func (o *ActivityOptionsModel) HasIncludeBundledActivityIds() bool`

HasIncludeBundledActivityIds returns a boolean if a field has been set.

### GetIncludePartyName

`func (o *ActivityOptionsModel) GetIncludePartyName() bool`

GetIncludePartyName returns the IncludePartyName field if non-nil, zero value otherwise.

### GetIncludePartyNameOk

`func (o *ActivityOptionsModel) GetIncludePartyNameOk() (*bool, bool)`

GetIncludePartyNameOk returns a tuple with the IncludePartyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePartyName

`func (o *ActivityOptionsModel) SetIncludePartyName(v bool)`

SetIncludePartyName sets IncludePartyName field to given value.

### HasIncludePartyName

`func (o *ActivityOptionsModel) HasIncludePartyName() bool`

HasIncludePartyName returns a boolean if a field has been set.

### GetIncludeShipmentActivityNr

`func (o *ActivityOptionsModel) GetIncludeShipmentActivityNr() bool`

GetIncludeShipmentActivityNr returns the IncludeShipmentActivityNr field if non-nil, zero value otherwise.

### GetIncludeShipmentActivityNrOk

`func (o *ActivityOptionsModel) GetIncludeShipmentActivityNrOk() (*bool, bool)`

GetIncludeShipmentActivityNrOk returns a tuple with the IncludeShipmentActivityNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeShipmentActivityNr

`func (o *ActivityOptionsModel) SetIncludeShipmentActivityNr(v bool)`

SetIncludeShipmentActivityNr sets IncludeShipmentActivityNr field to given value.

### HasIncludeShipmentActivityNr

`func (o *ActivityOptionsModel) HasIncludeShipmentActivityNr() bool`

HasIncludeShipmentActivityNr returns a boolean if a field has been set.

### GetIncludePayments

`func (o *ActivityOptionsModel) GetIncludePayments() bool`

GetIncludePayments returns the IncludePayments field if non-nil, zero value otherwise.

### GetIncludePaymentsOk

`func (o *ActivityOptionsModel) GetIncludePaymentsOk() (*bool, bool)`

GetIncludePaymentsOk returns a tuple with the IncludePayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePayments

`func (o *ActivityOptionsModel) SetIncludePayments(v bool)`

SetIncludePayments sets IncludePayments field to given value.

### HasIncludePayments

`func (o *ActivityOptionsModel) HasIncludePayments() bool`

HasIncludePayments returns a boolean if a field has been set.

### GetIncludeTransactions

`func (o *ActivityOptionsModel) GetIncludeTransactions() bool`

GetIncludeTransactions returns the IncludeTransactions field if non-nil, zero value otherwise.

### GetIncludeTransactionsOk

`func (o *ActivityOptionsModel) GetIncludeTransactionsOk() (*bool, bool)`

GetIncludeTransactionsOk returns a tuple with the IncludeTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTransactions

`func (o *ActivityOptionsModel) SetIncludeTransactions(v bool)`

SetIncludeTransactions sets IncludeTransactions field to given value.

### HasIncludeTransactions

`func (o *ActivityOptionsModel) HasIncludeTransactions() bool`

HasIncludeTransactions returns a boolean if a field has been set.

### GetIncludeRelations

`func (o *ActivityOptionsModel) GetIncludeRelations() bool`

GetIncludeRelations returns the IncludeRelations field if non-nil, zero value otherwise.

### GetIncludeRelationsOk

`func (o *ActivityOptionsModel) GetIncludeRelationsOk() (*bool, bool)`

GetIncludeRelationsOk returns a tuple with the IncludeRelations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRelations

`func (o *ActivityOptionsModel) SetIncludeRelations(v bool)`

SetIncludeRelations sets IncludeRelations field to given value.

### HasIncludeRelations

`func (o *ActivityOptionsModel) HasIncludeRelations() bool`

HasIncludeRelations returns a boolean if a field has been set.

### GetIncludeStats

`func (o *ActivityOptionsModel) GetIncludeStats() bool`

GetIncludeStats returns the IncludeStats field if non-nil, zero value otherwise.

### GetIncludeStatsOk

`func (o *ActivityOptionsModel) GetIncludeStatsOk() (*bool, bool)`

GetIncludeStatsOk returns a tuple with the IncludeStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeStats

`func (o *ActivityOptionsModel) SetIncludeStats(v bool)`

SetIncludeStats sets IncludeStats field to given value.

### HasIncludeStats

`func (o *ActivityOptionsModel) HasIncludeStats() bool`

HasIncludeStats returns a boolean if a field has been set.

### GetIncludePackagelinesForWholeRouteOnStartActivity

`func (o *ActivityOptionsModel) GetIncludePackagelinesForWholeRouteOnStartActivity() bool`

GetIncludePackagelinesForWholeRouteOnStartActivity returns the IncludePackagelinesForWholeRouteOnStartActivity field if non-nil, zero value otherwise.

### GetIncludePackagelinesForWholeRouteOnStartActivityOk

`func (o *ActivityOptionsModel) GetIncludePackagelinesForWholeRouteOnStartActivityOk() (*bool, bool)`

GetIncludePackagelinesForWholeRouteOnStartActivityOk returns a tuple with the IncludePackagelinesForWholeRouteOnStartActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePackagelinesForWholeRouteOnStartActivity

`func (o *ActivityOptionsModel) SetIncludePackagelinesForWholeRouteOnStartActivity(v bool)`

SetIncludePackagelinesForWholeRouteOnStartActivity sets IncludePackagelinesForWholeRouteOnStartActivity field to given value.

### HasIncludePackagelinesForWholeRouteOnStartActivity

`func (o *ActivityOptionsModel) HasIncludePackagelinesForWholeRouteOnStartActivity() bool`

HasIncludePackagelinesForWholeRouteOnStartActivity returns a boolean if a field has been set.

### GetIncludeActivityMetaData

`func (o *ActivityOptionsModel) GetIncludeActivityMetaData() bool`

GetIncludeActivityMetaData returns the IncludeActivityMetaData field if non-nil, zero value otherwise.

### GetIncludeActivityMetaDataOk

`func (o *ActivityOptionsModel) GetIncludeActivityMetaDataOk() (*bool, bool)`

GetIncludeActivityMetaDataOk returns a tuple with the IncludeActivityMetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeActivityMetaData

`func (o *ActivityOptionsModel) SetIncludeActivityMetaData(v bool)`

SetIncludeActivityMetaData sets IncludeActivityMetaData field to given value.

### HasIncludeActivityMetaData

`func (o *ActivityOptionsModel) HasIncludeActivityMetaData() bool`

HasIncludeActivityMetaData returns a boolean if a field has been set.

### GetIncludeAddressObject

`func (o *ActivityOptionsModel) GetIncludeAddressObject() bool`

GetIncludeAddressObject returns the IncludeAddressObject field if non-nil, zero value otherwise.

### GetIncludeAddressObjectOk

`func (o *ActivityOptionsModel) GetIncludeAddressObjectOk() (*bool, bool)`

GetIncludeAddressObjectOk returns a tuple with the IncludeAddressObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAddressObject

`func (o *ActivityOptionsModel) SetIncludeAddressObject(v bool)`

SetIncludeAddressObject sets IncludeAddressObject field to given value.

### HasIncludeAddressObject

`func (o *ActivityOptionsModel) HasIncludeAddressObject() bool`

HasIncludeAddressObject returns a boolean if a field has been set.

### GetIncludeCar

`func (o *ActivityOptionsModel) GetIncludeCar() bool`

GetIncludeCar returns the IncludeCar field if non-nil, zero value otherwise.

### GetIncludeCarOk

`func (o *ActivityOptionsModel) GetIncludeCarOk() (*bool, bool)`

GetIncludeCarOk returns a tuple with the IncludeCar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCar

`func (o *ActivityOptionsModel) SetIncludeCar(v bool)`

SetIncludeCar sets IncludeCar field to given value.

### HasIncludeCar

`func (o *ActivityOptionsModel) HasIncludeCar() bool`

HasIncludeCar returns a boolean if a field has been set.

### GetIncludeActivityLinks

`func (o *ActivityOptionsModel) GetIncludeActivityLinks() bool`

GetIncludeActivityLinks returns the IncludeActivityLinks field if non-nil, zero value otherwise.

### GetIncludeActivityLinksOk

`func (o *ActivityOptionsModel) GetIncludeActivityLinksOk() (*bool, bool)`

GetIncludeActivityLinksOk returns a tuple with the IncludeActivityLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeActivityLinks

`func (o *ActivityOptionsModel) SetIncludeActivityLinks(v bool)`

SetIncludeActivityLinks sets IncludeActivityLinks field to given value.

### HasIncludeActivityLinks

`func (o *ActivityOptionsModel) HasIncludeActivityLinks() bool`

HasIncludeActivityLinks returns a boolean if a field has been set.

### GetIncludeActivityFiles

`func (o *ActivityOptionsModel) GetIncludeActivityFiles() bool`

GetIncludeActivityFiles returns the IncludeActivityFiles field if non-nil, zero value otherwise.

### GetIncludeActivityFilesOk

`func (o *ActivityOptionsModel) GetIncludeActivityFilesOk() (*bool, bool)`

GetIncludeActivityFilesOk returns a tuple with the IncludeActivityFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeActivityFiles

`func (o *ActivityOptionsModel) SetIncludeActivityFiles(v bool)`

SetIncludeActivityFiles sets IncludeActivityFiles field to given value.

### HasIncludeActivityFiles

`func (o *ActivityOptionsModel) HasIncludeActivityFiles() bool`

HasIncludeActivityFiles returns a boolean if a field has been set.

### GetIncludeBrandFiles

`func (o *ActivityOptionsModel) GetIncludeBrandFiles() bool`

GetIncludeBrandFiles returns the IncludeBrandFiles field if non-nil, zero value otherwise.

### GetIncludeBrandFilesOk

`func (o *ActivityOptionsModel) GetIncludeBrandFilesOk() (*bool, bool)`

GetIncludeBrandFilesOk returns a tuple with the IncludeBrandFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBrandFiles

`func (o *ActivityOptionsModel) SetIncludeBrandFiles(v bool)`

SetIncludeBrandFiles sets IncludeBrandFiles field to given value.

### HasIncludeBrandFiles

`func (o *ActivityOptionsModel) HasIncludeBrandFiles() bool`

HasIncludeBrandFiles returns a boolean if a field has been set.

### GetIncludeActivityFilesMetaData

`func (o *ActivityOptionsModel) GetIncludeActivityFilesMetaData() bool`

GetIncludeActivityFilesMetaData returns the IncludeActivityFilesMetaData field if non-nil, zero value otherwise.

### GetIncludeActivityFilesMetaDataOk

`func (o *ActivityOptionsModel) GetIncludeActivityFilesMetaDataOk() (*bool, bool)`

GetIncludeActivityFilesMetaDataOk returns a tuple with the IncludeActivityFilesMetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeActivityFilesMetaData

`func (o *ActivityOptionsModel) SetIncludeActivityFilesMetaData(v bool)`

SetIncludeActivityFilesMetaData sets IncludeActivityFilesMetaData field to given value.

### HasIncludeActivityFilesMetaData

`func (o *ActivityOptionsModel) HasIncludeActivityFilesMetaData() bool`

HasIncludeActivityFilesMetaData returns a boolean if a field has been set.

### GetIncludeActivityRecordInfo

`func (o *ActivityOptionsModel) GetIncludeActivityRecordInfo() bool`

GetIncludeActivityRecordInfo returns the IncludeActivityRecordInfo field if non-nil, zero value otherwise.

### GetIncludeActivityRecordInfoOk

`func (o *ActivityOptionsModel) GetIncludeActivityRecordInfoOk() (*bool, bool)`

GetIncludeActivityRecordInfoOk returns a tuple with the IncludeActivityRecordInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeActivityRecordInfo

`func (o *ActivityOptionsModel) SetIncludeActivityRecordInfo(v bool)`

SetIncludeActivityRecordInfo sets IncludeActivityRecordInfo field to given value.

### HasIncludeActivityRecordInfo

`func (o *ActivityOptionsModel) HasIncludeActivityRecordInfo() bool`

HasIncludeActivityRecordInfo returns a boolean if a field has been set.

### GetIncludeActivityRecordObject

`func (o *ActivityOptionsModel) GetIncludeActivityRecordObject() bool`

GetIncludeActivityRecordObject returns the IncludeActivityRecordObject field if non-nil, zero value otherwise.

### GetIncludeActivityRecordObjectOk

`func (o *ActivityOptionsModel) GetIncludeActivityRecordObjectOk() (*bool, bool)`

GetIncludeActivityRecordObjectOk returns a tuple with the IncludeActivityRecordObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeActivityRecordObject

`func (o *ActivityOptionsModel) SetIncludeActivityRecordObject(v bool)`

SetIncludeActivityRecordObject sets IncludeActivityRecordObject field to given value.

### HasIncludeActivityRecordObject

`func (o *ActivityOptionsModel) HasIncludeActivityRecordObject() bool`

HasIncludeActivityRecordObject returns a boolean if a field has been set.

### GetIncludeActivityNotes

`func (o *ActivityOptionsModel) GetIncludeActivityNotes() bool`

GetIncludeActivityNotes returns the IncludeActivityNotes field if non-nil, zero value otherwise.

### GetIncludeActivityNotesOk

`func (o *ActivityOptionsModel) GetIncludeActivityNotesOk() (*bool, bool)`

GetIncludeActivityNotesOk returns a tuple with the IncludeActivityNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeActivityNotes

`func (o *ActivityOptionsModel) SetIncludeActivityNotes(v bool)`

SetIncludeActivityNotes sets IncludeActivityNotes field to given value.

### HasIncludeActivityNotes

`func (o *ActivityOptionsModel) HasIncludeActivityNotes() bool`

HasIncludeActivityNotes returns a boolean if a field has been set.

### GetIncludeActivityTags

`func (o *ActivityOptionsModel) GetIncludeActivityTags() bool`

GetIncludeActivityTags returns the IncludeActivityTags field if non-nil, zero value otherwise.

### GetIncludeActivityTagsOk

`func (o *ActivityOptionsModel) GetIncludeActivityTagsOk() (*bool, bool)`

GetIncludeActivityTagsOk returns a tuple with the IncludeActivityTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeActivityTags

`func (o *ActivityOptionsModel) SetIncludeActivityTags(v bool)`

SetIncludeActivityTags sets IncludeActivityTags field to given value.

### HasIncludeActivityTags

`func (o *ActivityOptionsModel) HasIncludeActivityTags() bool`

HasIncludeActivityTags returns a boolean if a field has been set.

### GetIncludeRecurrence

`func (o *ActivityOptionsModel) GetIncludeRecurrence() bool`

GetIncludeRecurrence returns the IncludeRecurrence field if non-nil, zero value otherwise.

### GetIncludeRecurrenceOk

`func (o *ActivityOptionsModel) GetIncludeRecurrenceOk() (*bool, bool)`

GetIncludeRecurrenceOk returns a tuple with the IncludeRecurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRecurrence

`func (o *ActivityOptionsModel) SetIncludeRecurrence(v bool)`

SetIncludeRecurrence sets IncludeRecurrence field to given value.

### HasIncludeRecurrence

`func (o *ActivityOptionsModel) HasIncludeRecurrence() bool`

HasIncludeRecurrence returns a boolean if a field has been set.

### GetIncludeMatchingRouteIds

`func (o *ActivityOptionsModel) GetIncludeMatchingRouteIds() bool`

GetIncludeMatchingRouteIds returns the IncludeMatchingRouteIds field if non-nil, zero value otherwise.

### GetIncludeMatchingRouteIdsOk

`func (o *ActivityOptionsModel) GetIncludeMatchingRouteIdsOk() (*bool, bool)`

GetIncludeMatchingRouteIdsOk returns a tuple with the IncludeMatchingRouteIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeMatchingRouteIds

`func (o *ActivityOptionsModel) SetIncludeMatchingRouteIds(v bool)`

SetIncludeMatchingRouteIds sets IncludeMatchingRouteIds field to given value.

### HasIncludeMatchingRouteIds

`func (o *ActivityOptionsModel) HasIncludeMatchingRouteIds() bool`

HasIncludeMatchingRouteIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


