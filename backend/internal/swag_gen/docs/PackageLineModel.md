# PackageLineModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique Identifier | [optional] [readonly] 
**ActivityId** | Pointer to **int64** | Unique Identifier for activity where this packageline is related to | [optional] 
**ActivityIds** | Pointer to **[]int32** | Unique Identifier for activities where this packageline is related to | [optional] 
**Nr** | Pointer to **string** | Number of this PackageLine | [optional] 
**StatusId** | Pointer to **int32** | StatusId of this PackageLine, 31: package_line_cancelled, 23: package_line_incomplete, 24: package_line_new, 42: package_line_awaiting, 25: package_line_accepted, 10: package_line_planned, 11: package_line_in_progress, 12: package_line_executed | [optional] 
**StatusName** | Pointer to **string** | PackageLine Status | [optional] 
**NrOfPackages** | Pointer to **float32** | Number of packages in package line | [optional] 
**PackageTypeName** | Pointer to **string** | Type of the Packages in the package line | [optional] 
**PackageTypeId** | Pointer to **int32** | ID of the package type for the packages in this PackageLine | [optional] 
**Barcode** | Pointer to **string** | Barcode of this packageline | [optional] 
**Barcodes** | Pointer to **[]string** | For every barcode in this array, a seperate packageline will be created | [optional] 
**ActionTypeId** | Pointer to **string** | Action type name, 1:inbound, 2:outbound, 3:assess | [optional] 
**ActionTypeName** | Pointer to **string** | Action type name | [optional] 
**CheckedByDriver** | Pointer to **bool** | boolean for whether or not the packages have been checked by the driver | [optional] 
**ADR** | Pointer to **bool** | boolean for whether or not the packages in this package line should be considered as ADR | [optional] 
**ADRClass** | Pointer to **int64** | ADR class of packages in package line | [optional] 
**ADRUNNr** | Pointer to **int64** | ADR UN Nr of packages in package line | [optional] 
**Temp** | Pointer to **bool** | boolean for whether or not the packages in this package line should be considered as temperature dependent | [optional] 
**TempMin** | Pointer to **float32** | minimum temperature for packages in package line | [optional] 
**TempMax** | Pointer to **float32** | maximum temperature for packages in package line | [optional] 
**HsCode** | Pointer to **string** | Harmonized System code for packages in this package line | [optional] 
**Description** | Pointer to **string** | description of this package_line | [optional] 
**Compartments** | Pointer to [**[]CompartmentModel**](CompartmentModel.md) |  | [optional] 
**Active** | Pointer to **bool** | if active&#x3D;0: package line has been removed and is no longer visible in any bumbal interface | [optional] 
**AppliedCapacities** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**Capacities** | Pointer to [**[]CapacityModel**](CapacityModel.md) |  | [optional] 
**ActivityLinks** | Pointer to [**[]LinkModel**](LinkModel.md) | links to activities connected to this package_line | [optional] 
**Links** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**MetaData** | Pointer to [**[]MetaDataModel**](MetaDataModel.md) |  | [optional] 
**Notes** | Pointer to [**[]NoteModel**](NoteModel.md) |  | [optional] 
**Files** | Pointer to [**[]FileModel**](FileModel.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | created_at date time | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | updated_at date time | [optional] [readonly] 
**PackageLineCreatedAt** | Pointer to **time.Time** | created_at date time | [optional] [readonly] 
**PackageLineUpdatedAt** | Pointer to **time.Time** | updated_at date time | [optional] [readonly] 
**PackageLineCreatedBy** | Pointer to **int32** | created_by user id | [optional] [readonly] 
**PackageLineUpdatedBy** | Pointer to **int32** | updated_by user id | [optional] [readonly] 

## Methods

### NewPackageLineModel

`func NewPackageLineModel() *PackageLineModel`

NewPackageLineModel instantiates a new PackageLineModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageLineModelWithDefaults

`func NewPackageLineModelWithDefaults() *PackageLineModel`

NewPackageLineModelWithDefaults instantiates a new PackageLineModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PackageLineModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PackageLineModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PackageLineModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PackageLineModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActivityId

`func (o *PackageLineModel) GetActivityId() int64`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *PackageLineModel) GetActivityIdOk() (*int64, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *PackageLineModel) SetActivityId(v int64)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *PackageLineModel) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### GetActivityIds

`func (o *PackageLineModel) GetActivityIds() []int32`

GetActivityIds returns the ActivityIds field if non-nil, zero value otherwise.

### GetActivityIdsOk

`func (o *PackageLineModel) GetActivityIdsOk() (*[]int32, bool)`

GetActivityIdsOk returns a tuple with the ActivityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityIds

`func (o *PackageLineModel) SetActivityIds(v []int32)`

SetActivityIds sets ActivityIds field to given value.

### HasActivityIds

`func (o *PackageLineModel) HasActivityIds() bool`

HasActivityIds returns a boolean if a field has been set.

### GetNr

`func (o *PackageLineModel) GetNr() string`

GetNr returns the Nr field if non-nil, zero value otherwise.

### GetNrOk

`func (o *PackageLineModel) GetNrOk() (*string, bool)`

GetNrOk returns a tuple with the Nr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNr

`func (o *PackageLineModel) SetNr(v string)`

SetNr sets Nr field to given value.

### HasNr

`func (o *PackageLineModel) HasNr() bool`

HasNr returns a boolean if a field has been set.

### GetStatusId

`func (o *PackageLineModel) GetStatusId() int32`

GetStatusId returns the StatusId field if non-nil, zero value otherwise.

### GetStatusIdOk

`func (o *PackageLineModel) GetStatusIdOk() (*int32, bool)`

GetStatusIdOk returns a tuple with the StatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusId

`func (o *PackageLineModel) SetStatusId(v int32)`

SetStatusId sets StatusId field to given value.

### HasStatusId

`func (o *PackageLineModel) HasStatusId() bool`

HasStatusId returns a boolean if a field has been set.

### GetStatusName

`func (o *PackageLineModel) GetStatusName() string`

GetStatusName returns the StatusName field if non-nil, zero value otherwise.

### GetStatusNameOk

`func (o *PackageLineModel) GetStatusNameOk() (*string, bool)`

GetStatusNameOk returns a tuple with the StatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusName

`func (o *PackageLineModel) SetStatusName(v string)`

SetStatusName sets StatusName field to given value.

### HasStatusName

`func (o *PackageLineModel) HasStatusName() bool`

HasStatusName returns a boolean if a field has been set.

### GetNrOfPackages

`func (o *PackageLineModel) GetNrOfPackages() float32`

GetNrOfPackages returns the NrOfPackages field if non-nil, zero value otherwise.

### GetNrOfPackagesOk

`func (o *PackageLineModel) GetNrOfPackagesOk() (*float32, bool)`

GetNrOfPackagesOk returns a tuple with the NrOfPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrOfPackages

`func (o *PackageLineModel) SetNrOfPackages(v float32)`

SetNrOfPackages sets NrOfPackages field to given value.

### HasNrOfPackages

`func (o *PackageLineModel) HasNrOfPackages() bool`

HasNrOfPackages returns a boolean if a field has been set.

### GetPackageTypeName

`func (o *PackageLineModel) GetPackageTypeName() string`

GetPackageTypeName returns the PackageTypeName field if non-nil, zero value otherwise.

### GetPackageTypeNameOk

`func (o *PackageLineModel) GetPackageTypeNameOk() (*string, bool)`

GetPackageTypeNameOk returns a tuple with the PackageTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageTypeName

`func (o *PackageLineModel) SetPackageTypeName(v string)`

SetPackageTypeName sets PackageTypeName field to given value.

### HasPackageTypeName

`func (o *PackageLineModel) HasPackageTypeName() bool`

HasPackageTypeName returns a boolean if a field has been set.

### GetPackageTypeId

`func (o *PackageLineModel) GetPackageTypeId() int32`

GetPackageTypeId returns the PackageTypeId field if non-nil, zero value otherwise.

### GetPackageTypeIdOk

`func (o *PackageLineModel) GetPackageTypeIdOk() (*int32, bool)`

GetPackageTypeIdOk returns a tuple with the PackageTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageTypeId

`func (o *PackageLineModel) SetPackageTypeId(v int32)`

SetPackageTypeId sets PackageTypeId field to given value.

### HasPackageTypeId

`func (o *PackageLineModel) HasPackageTypeId() bool`

HasPackageTypeId returns a boolean if a field has been set.

### GetBarcode

`func (o *PackageLineModel) GetBarcode() string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *PackageLineModel) GetBarcodeOk() (*string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *PackageLineModel) SetBarcode(v string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *PackageLineModel) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetBarcodes

`func (o *PackageLineModel) GetBarcodes() []string`

GetBarcodes returns the Barcodes field if non-nil, zero value otherwise.

### GetBarcodesOk

`func (o *PackageLineModel) GetBarcodesOk() (*[]string, bool)`

GetBarcodesOk returns a tuple with the Barcodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcodes

`func (o *PackageLineModel) SetBarcodes(v []string)`

SetBarcodes sets Barcodes field to given value.

### HasBarcodes

`func (o *PackageLineModel) HasBarcodes() bool`

HasBarcodes returns a boolean if a field has been set.

### GetActionTypeId

`func (o *PackageLineModel) GetActionTypeId() string`

GetActionTypeId returns the ActionTypeId field if non-nil, zero value otherwise.

### GetActionTypeIdOk

`func (o *PackageLineModel) GetActionTypeIdOk() (*string, bool)`

GetActionTypeIdOk returns a tuple with the ActionTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTypeId

`func (o *PackageLineModel) SetActionTypeId(v string)`

SetActionTypeId sets ActionTypeId field to given value.

### HasActionTypeId

`func (o *PackageLineModel) HasActionTypeId() bool`

HasActionTypeId returns a boolean if a field has been set.

### GetActionTypeName

`func (o *PackageLineModel) GetActionTypeName() string`

GetActionTypeName returns the ActionTypeName field if non-nil, zero value otherwise.

### GetActionTypeNameOk

`func (o *PackageLineModel) GetActionTypeNameOk() (*string, bool)`

GetActionTypeNameOk returns a tuple with the ActionTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTypeName

`func (o *PackageLineModel) SetActionTypeName(v string)`

SetActionTypeName sets ActionTypeName field to given value.

### HasActionTypeName

`func (o *PackageLineModel) HasActionTypeName() bool`

HasActionTypeName returns a boolean if a field has been set.

### GetCheckedByDriver

`func (o *PackageLineModel) GetCheckedByDriver() bool`

GetCheckedByDriver returns the CheckedByDriver field if non-nil, zero value otherwise.

### GetCheckedByDriverOk

`func (o *PackageLineModel) GetCheckedByDriverOk() (*bool, bool)`

GetCheckedByDriverOk returns a tuple with the CheckedByDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckedByDriver

`func (o *PackageLineModel) SetCheckedByDriver(v bool)`

SetCheckedByDriver sets CheckedByDriver field to given value.

### HasCheckedByDriver

`func (o *PackageLineModel) HasCheckedByDriver() bool`

HasCheckedByDriver returns a boolean if a field has been set.

### GetADR

`func (o *PackageLineModel) GetADR() bool`

GetADR returns the ADR field if non-nil, zero value otherwise.

### GetADROk

`func (o *PackageLineModel) GetADROk() (*bool, bool)`

GetADROk returns a tuple with the ADR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADR

`func (o *PackageLineModel) SetADR(v bool)`

SetADR sets ADR field to given value.

### HasADR

`func (o *PackageLineModel) HasADR() bool`

HasADR returns a boolean if a field has been set.

### GetADRClass

`func (o *PackageLineModel) GetADRClass() int64`

GetADRClass returns the ADRClass field if non-nil, zero value otherwise.

### GetADRClassOk

`func (o *PackageLineModel) GetADRClassOk() (*int64, bool)`

GetADRClassOk returns a tuple with the ADRClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADRClass

`func (o *PackageLineModel) SetADRClass(v int64)`

SetADRClass sets ADRClass field to given value.

### HasADRClass

`func (o *PackageLineModel) HasADRClass() bool`

HasADRClass returns a boolean if a field has been set.

### GetADRUNNr

`func (o *PackageLineModel) GetADRUNNr() int64`

GetADRUNNr returns the ADRUNNr field if non-nil, zero value otherwise.

### GetADRUNNrOk

`func (o *PackageLineModel) GetADRUNNrOk() (*int64, bool)`

GetADRUNNrOk returns a tuple with the ADRUNNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADRUNNr

`func (o *PackageLineModel) SetADRUNNr(v int64)`

SetADRUNNr sets ADRUNNr field to given value.

### HasADRUNNr

`func (o *PackageLineModel) HasADRUNNr() bool`

HasADRUNNr returns a boolean if a field has been set.

### GetTemp

`func (o *PackageLineModel) GetTemp() bool`

GetTemp returns the Temp field if non-nil, zero value otherwise.

### GetTempOk

`func (o *PackageLineModel) GetTempOk() (*bool, bool)`

GetTempOk returns a tuple with the Temp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemp

`func (o *PackageLineModel) SetTemp(v bool)`

SetTemp sets Temp field to given value.

### HasTemp

`func (o *PackageLineModel) HasTemp() bool`

HasTemp returns a boolean if a field has been set.

### GetTempMin

`func (o *PackageLineModel) GetTempMin() float32`

GetTempMin returns the TempMin field if non-nil, zero value otherwise.

### GetTempMinOk

`func (o *PackageLineModel) GetTempMinOk() (*float32, bool)`

GetTempMinOk returns a tuple with the TempMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempMin

`func (o *PackageLineModel) SetTempMin(v float32)`

SetTempMin sets TempMin field to given value.

### HasTempMin

`func (o *PackageLineModel) HasTempMin() bool`

HasTempMin returns a boolean if a field has been set.

### GetTempMax

`func (o *PackageLineModel) GetTempMax() float32`

GetTempMax returns the TempMax field if non-nil, zero value otherwise.

### GetTempMaxOk

`func (o *PackageLineModel) GetTempMaxOk() (*float32, bool)`

GetTempMaxOk returns a tuple with the TempMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempMax

`func (o *PackageLineModel) SetTempMax(v float32)`

SetTempMax sets TempMax field to given value.

### HasTempMax

`func (o *PackageLineModel) HasTempMax() bool`

HasTempMax returns a boolean if a field has been set.

### GetHsCode

`func (o *PackageLineModel) GetHsCode() string`

GetHsCode returns the HsCode field if non-nil, zero value otherwise.

### GetHsCodeOk

`func (o *PackageLineModel) GetHsCodeOk() (*string, bool)`

GetHsCodeOk returns a tuple with the HsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsCode

`func (o *PackageLineModel) SetHsCode(v string)`

SetHsCode sets HsCode field to given value.

### HasHsCode

`func (o *PackageLineModel) HasHsCode() bool`

HasHsCode returns a boolean if a field has been set.

### GetDescription

`func (o *PackageLineModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PackageLineModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PackageLineModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PackageLineModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCompartments

`func (o *PackageLineModel) GetCompartments() []CompartmentModel`

GetCompartments returns the Compartments field if non-nil, zero value otherwise.

### GetCompartmentsOk

`func (o *PackageLineModel) GetCompartmentsOk() (*[]CompartmentModel, bool)`

GetCompartmentsOk returns a tuple with the Compartments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompartments

`func (o *PackageLineModel) SetCompartments(v []CompartmentModel)`

SetCompartments sets Compartments field to given value.

### HasCompartments

`func (o *PackageLineModel) HasCompartments() bool`

HasCompartments returns a boolean if a field has been set.

### GetActive

`func (o *PackageLineModel) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PackageLineModel) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PackageLineModel) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *PackageLineModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAppliedCapacities

`func (o *PackageLineModel) GetAppliedCapacities() map[string]interface{}`

GetAppliedCapacities returns the AppliedCapacities field if non-nil, zero value otherwise.

### GetAppliedCapacitiesOk

`func (o *PackageLineModel) GetAppliedCapacitiesOk() (*map[string]interface{}, bool)`

GetAppliedCapacitiesOk returns a tuple with the AppliedCapacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedCapacities

`func (o *PackageLineModel) SetAppliedCapacities(v map[string]interface{})`

SetAppliedCapacities sets AppliedCapacities field to given value.

### HasAppliedCapacities

`func (o *PackageLineModel) HasAppliedCapacities() bool`

HasAppliedCapacities returns a boolean if a field has been set.

### GetCapacities

`func (o *PackageLineModel) GetCapacities() []CapacityModel`

GetCapacities returns the Capacities field if non-nil, zero value otherwise.

### GetCapacitiesOk

`func (o *PackageLineModel) GetCapacitiesOk() (*[]CapacityModel, bool)`

GetCapacitiesOk returns a tuple with the Capacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacities

`func (o *PackageLineModel) SetCapacities(v []CapacityModel)`

SetCapacities sets Capacities field to given value.

### HasCapacities

`func (o *PackageLineModel) HasCapacities() bool`

HasCapacities returns a boolean if a field has been set.

### GetActivityLinks

`func (o *PackageLineModel) GetActivityLinks() []LinkModel`

GetActivityLinks returns the ActivityLinks field if non-nil, zero value otherwise.

### GetActivityLinksOk

`func (o *PackageLineModel) GetActivityLinksOk() (*[]LinkModel, bool)`

GetActivityLinksOk returns a tuple with the ActivityLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityLinks

`func (o *PackageLineModel) SetActivityLinks(v []LinkModel)`

SetActivityLinks sets ActivityLinks field to given value.

### HasActivityLinks

`func (o *PackageLineModel) HasActivityLinks() bool`

HasActivityLinks returns a boolean if a field has been set.

### GetLinks

`func (o *PackageLineModel) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PackageLineModel) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PackageLineModel) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PackageLineModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMetaData

`func (o *PackageLineModel) GetMetaData() []MetaDataModel`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *PackageLineModel) GetMetaDataOk() (*[]MetaDataModel, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *PackageLineModel) SetMetaData(v []MetaDataModel)`

SetMetaData sets MetaData field to given value.

### HasMetaData

`func (o *PackageLineModel) HasMetaData() bool`

HasMetaData returns a boolean if a field has been set.

### GetNotes

`func (o *PackageLineModel) GetNotes() []NoteModel`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *PackageLineModel) GetNotesOk() (*[]NoteModel, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *PackageLineModel) SetNotes(v []NoteModel)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *PackageLineModel) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetFiles

`func (o *PackageLineModel) GetFiles() []FileModel`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *PackageLineModel) GetFilesOk() (*[]FileModel, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *PackageLineModel) SetFiles(v []FileModel)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *PackageLineModel) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PackageLineModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PackageLineModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PackageLineModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PackageLineModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PackageLineModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PackageLineModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PackageLineModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PackageLineModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetPackageLineCreatedAt

`func (o *PackageLineModel) GetPackageLineCreatedAt() time.Time`

GetPackageLineCreatedAt returns the PackageLineCreatedAt field if non-nil, zero value otherwise.

### GetPackageLineCreatedAtOk

`func (o *PackageLineModel) GetPackageLineCreatedAtOk() (*time.Time, bool)`

GetPackageLineCreatedAtOk returns a tuple with the PackageLineCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLineCreatedAt

`func (o *PackageLineModel) SetPackageLineCreatedAt(v time.Time)`

SetPackageLineCreatedAt sets PackageLineCreatedAt field to given value.

### HasPackageLineCreatedAt

`func (o *PackageLineModel) HasPackageLineCreatedAt() bool`

HasPackageLineCreatedAt returns a boolean if a field has been set.

### GetPackageLineUpdatedAt

`func (o *PackageLineModel) GetPackageLineUpdatedAt() time.Time`

GetPackageLineUpdatedAt returns the PackageLineUpdatedAt field if non-nil, zero value otherwise.

### GetPackageLineUpdatedAtOk

`func (o *PackageLineModel) GetPackageLineUpdatedAtOk() (*time.Time, bool)`

GetPackageLineUpdatedAtOk returns a tuple with the PackageLineUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLineUpdatedAt

`func (o *PackageLineModel) SetPackageLineUpdatedAt(v time.Time)`

SetPackageLineUpdatedAt sets PackageLineUpdatedAt field to given value.

### HasPackageLineUpdatedAt

`func (o *PackageLineModel) HasPackageLineUpdatedAt() bool`

HasPackageLineUpdatedAt returns a boolean if a field has been set.

### GetPackageLineCreatedBy

`func (o *PackageLineModel) GetPackageLineCreatedBy() int32`

GetPackageLineCreatedBy returns the PackageLineCreatedBy field if non-nil, zero value otherwise.

### GetPackageLineCreatedByOk

`func (o *PackageLineModel) GetPackageLineCreatedByOk() (*int32, bool)`

GetPackageLineCreatedByOk returns a tuple with the PackageLineCreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLineCreatedBy

`func (o *PackageLineModel) SetPackageLineCreatedBy(v int32)`

SetPackageLineCreatedBy sets PackageLineCreatedBy field to given value.

### HasPackageLineCreatedBy

`func (o *PackageLineModel) HasPackageLineCreatedBy() bool`

HasPackageLineCreatedBy returns a boolean if a field has been set.

### GetPackageLineUpdatedBy

`func (o *PackageLineModel) GetPackageLineUpdatedBy() int32`

GetPackageLineUpdatedBy returns the PackageLineUpdatedBy field if non-nil, zero value otherwise.

### GetPackageLineUpdatedByOk

`func (o *PackageLineModel) GetPackageLineUpdatedByOk() (*int32, bool)`

GetPackageLineUpdatedByOk returns a tuple with the PackageLineUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLineUpdatedBy

`func (o *PackageLineModel) SetPackageLineUpdatedBy(v int32)`

SetPackageLineUpdatedBy sets PackageLineUpdatedBy field to given value.

### HasPackageLineUpdatedBy

`func (o *PackageLineModel) HasPackageLineUpdatedBy() bool`

HasPackageLineUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


