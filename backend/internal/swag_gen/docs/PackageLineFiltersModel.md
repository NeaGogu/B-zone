# PackageLineFiltersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdatedAtSince** | Pointer to **time.Time** | Show updated since | [optional] 
**UpdatedAtTill** | Pointer to **time.Time** | Show updated till | [optional] 
**Id** | Pointer to **[]int32** | Bumbal package line id&#39;s | [optional] 
**Nr** | Pointer to **[]string** | PackageLine numbers | [optional] 
**StatusId** | Pointer to **[]int32** | StatusIds of PackageLine, 31: package_line_cancelled, 23: package_line_incomplete, 24: package_line_new, 42: package_line_awaiting, 25: package_line_accepted, 10: package_line_planned, 11: package_line_in_progress, 12: package_line_executed | [optional] 
**Active** | Pointer to **[]int32** | Active status of PackageLine, 0 values represent deleted PackageLines | [optional] 
**StatusName** | Pointer to **[]string** | PackageLine Status | [optional] 
**NrOfPackages** | Pointer to **float32** | Number of packages in package line | [optional] 
**PackageTypeName** | Pointer to **[]string** | Type of the Packages | [optional] 
**PackageTypeId** | Pointer to **[]int32** | ID of the package type for the packages | [optional] 
**AppliedCapacities** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**Capacities** | Pointer to [**[]CapacityModel**](CapacityModel.md) |  | [optional] 
**Barcode** | Pointer to **[]string** | barcode for packages | [optional] 
**ADR** | Pointer to **bool** | boolean for whether or not the packages in package line should be considered as ADR | [optional] 
**ADRClass** | Pointer to **[]int32** | ADR class of packages in package line | [optional] 
**ADRUNNr** | Pointer to **[]int32** | ADR UN Nr of packages in package line | [optional] 
**Temp** | Pointer to **bool** | boolean for whether or not the packages in package line should be considered as temperature dependent | [optional] 
**TempMin** | Pointer to **float32** | minimum temperature for packages in package line | [optional] 
**TempMax** | Pointer to **float32** | maximum temperature for packages in package line | [optional] 
**HsCode** | Pointer to **[]string** | Harmonized System code for packages | [optional] 
**Description** | Pointer to **string** | description of this package_line | [optional] 
**Links** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**ActivityLinks** | Pointer to [**[]LinkModel**](LinkModel.md) |  | [optional] 
**ActivityId** | Pointer to **[]int32** | Activity id | [optional] 

## Methods

### NewPackageLineFiltersModel

`func NewPackageLineFiltersModel() *PackageLineFiltersModel`

NewPackageLineFiltersModel instantiates a new PackageLineFiltersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageLineFiltersModelWithDefaults

`func NewPackageLineFiltersModelWithDefaults() *PackageLineFiltersModel`

NewPackageLineFiltersModelWithDefaults instantiates a new PackageLineFiltersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdatedAtSince

`func (o *PackageLineFiltersModel) GetUpdatedAtSince() time.Time`

GetUpdatedAtSince returns the UpdatedAtSince field if non-nil, zero value otherwise.

### GetUpdatedAtSinceOk

`func (o *PackageLineFiltersModel) GetUpdatedAtSinceOk() (*time.Time, bool)`

GetUpdatedAtSinceOk returns a tuple with the UpdatedAtSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtSince

`func (o *PackageLineFiltersModel) SetUpdatedAtSince(v time.Time)`

SetUpdatedAtSince sets UpdatedAtSince field to given value.

### HasUpdatedAtSince

`func (o *PackageLineFiltersModel) HasUpdatedAtSince() bool`

HasUpdatedAtSince returns a boolean if a field has been set.

### GetUpdatedAtTill

`func (o *PackageLineFiltersModel) GetUpdatedAtTill() time.Time`

GetUpdatedAtTill returns the UpdatedAtTill field if non-nil, zero value otherwise.

### GetUpdatedAtTillOk

`func (o *PackageLineFiltersModel) GetUpdatedAtTillOk() (*time.Time, bool)`

GetUpdatedAtTillOk returns a tuple with the UpdatedAtTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtTill

`func (o *PackageLineFiltersModel) SetUpdatedAtTill(v time.Time)`

SetUpdatedAtTill sets UpdatedAtTill field to given value.

### HasUpdatedAtTill

`func (o *PackageLineFiltersModel) HasUpdatedAtTill() bool`

HasUpdatedAtTill returns a boolean if a field has been set.

### GetId

`func (o *PackageLineFiltersModel) GetId() []int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PackageLineFiltersModel) GetIdOk() (*[]int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PackageLineFiltersModel) SetId(v []int32)`

SetId sets Id field to given value.

### HasId

`func (o *PackageLineFiltersModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNr

`func (o *PackageLineFiltersModel) GetNr() []string`

GetNr returns the Nr field if non-nil, zero value otherwise.

### GetNrOk

`func (o *PackageLineFiltersModel) GetNrOk() (*[]string, bool)`

GetNrOk returns a tuple with the Nr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNr

`func (o *PackageLineFiltersModel) SetNr(v []string)`

SetNr sets Nr field to given value.

### HasNr

`func (o *PackageLineFiltersModel) HasNr() bool`

HasNr returns a boolean if a field has been set.

### GetStatusId

`func (o *PackageLineFiltersModel) GetStatusId() []int32`

GetStatusId returns the StatusId field if non-nil, zero value otherwise.

### GetStatusIdOk

`func (o *PackageLineFiltersModel) GetStatusIdOk() (*[]int32, bool)`

GetStatusIdOk returns a tuple with the StatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusId

`func (o *PackageLineFiltersModel) SetStatusId(v []int32)`

SetStatusId sets StatusId field to given value.

### HasStatusId

`func (o *PackageLineFiltersModel) HasStatusId() bool`

HasStatusId returns a boolean if a field has been set.

### GetActive

`func (o *PackageLineFiltersModel) GetActive() []int32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *PackageLineFiltersModel) GetActiveOk() (*[]int32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *PackageLineFiltersModel) SetActive(v []int32)`

SetActive sets Active field to given value.

### HasActive

`func (o *PackageLineFiltersModel) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetStatusName

`func (o *PackageLineFiltersModel) GetStatusName() []string`

GetStatusName returns the StatusName field if non-nil, zero value otherwise.

### GetStatusNameOk

`func (o *PackageLineFiltersModel) GetStatusNameOk() (*[]string, bool)`

GetStatusNameOk returns a tuple with the StatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusName

`func (o *PackageLineFiltersModel) SetStatusName(v []string)`

SetStatusName sets StatusName field to given value.

### HasStatusName

`func (o *PackageLineFiltersModel) HasStatusName() bool`

HasStatusName returns a boolean if a field has been set.

### GetNrOfPackages

`func (o *PackageLineFiltersModel) GetNrOfPackages() float32`

GetNrOfPackages returns the NrOfPackages field if non-nil, zero value otherwise.

### GetNrOfPackagesOk

`func (o *PackageLineFiltersModel) GetNrOfPackagesOk() (*float32, bool)`

GetNrOfPackagesOk returns a tuple with the NrOfPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrOfPackages

`func (o *PackageLineFiltersModel) SetNrOfPackages(v float32)`

SetNrOfPackages sets NrOfPackages field to given value.

### HasNrOfPackages

`func (o *PackageLineFiltersModel) HasNrOfPackages() bool`

HasNrOfPackages returns a boolean if a field has been set.

### GetPackageTypeName

`func (o *PackageLineFiltersModel) GetPackageTypeName() []string`

GetPackageTypeName returns the PackageTypeName field if non-nil, zero value otherwise.

### GetPackageTypeNameOk

`func (o *PackageLineFiltersModel) GetPackageTypeNameOk() (*[]string, bool)`

GetPackageTypeNameOk returns a tuple with the PackageTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageTypeName

`func (o *PackageLineFiltersModel) SetPackageTypeName(v []string)`

SetPackageTypeName sets PackageTypeName field to given value.

### HasPackageTypeName

`func (o *PackageLineFiltersModel) HasPackageTypeName() bool`

HasPackageTypeName returns a boolean if a field has been set.

### GetPackageTypeId

`func (o *PackageLineFiltersModel) GetPackageTypeId() []int32`

GetPackageTypeId returns the PackageTypeId field if non-nil, zero value otherwise.

### GetPackageTypeIdOk

`func (o *PackageLineFiltersModel) GetPackageTypeIdOk() (*[]int32, bool)`

GetPackageTypeIdOk returns a tuple with the PackageTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageTypeId

`func (o *PackageLineFiltersModel) SetPackageTypeId(v []int32)`

SetPackageTypeId sets PackageTypeId field to given value.

### HasPackageTypeId

`func (o *PackageLineFiltersModel) HasPackageTypeId() bool`

HasPackageTypeId returns a boolean if a field has been set.

### GetAppliedCapacities

`func (o *PackageLineFiltersModel) GetAppliedCapacities() map[string]interface{}`

GetAppliedCapacities returns the AppliedCapacities field if non-nil, zero value otherwise.

### GetAppliedCapacitiesOk

`func (o *PackageLineFiltersModel) GetAppliedCapacitiesOk() (*map[string]interface{}, bool)`

GetAppliedCapacitiesOk returns a tuple with the AppliedCapacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedCapacities

`func (o *PackageLineFiltersModel) SetAppliedCapacities(v map[string]interface{})`

SetAppliedCapacities sets AppliedCapacities field to given value.

### HasAppliedCapacities

`func (o *PackageLineFiltersModel) HasAppliedCapacities() bool`

HasAppliedCapacities returns a boolean if a field has been set.

### GetCapacities

`func (o *PackageLineFiltersModel) GetCapacities() []CapacityModel`

GetCapacities returns the Capacities field if non-nil, zero value otherwise.

### GetCapacitiesOk

`func (o *PackageLineFiltersModel) GetCapacitiesOk() (*[]CapacityModel, bool)`

GetCapacitiesOk returns a tuple with the Capacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacities

`func (o *PackageLineFiltersModel) SetCapacities(v []CapacityModel)`

SetCapacities sets Capacities field to given value.

### HasCapacities

`func (o *PackageLineFiltersModel) HasCapacities() bool`

HasCapacities returns a boolean if a field has been set.

### GetBarcode

`func (o *PackageLineFiltersModel) GetBarcode() []string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *PackageLineFiltersModel) GetBarcodeOk() (*[]string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *PackageLineFiltersModel) SetBarcode(v []string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *PackageLineFiltersModel) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetADR

`func (o *PackageLineFiltersModel) GetADR() bool`

GetADR returns the ADR field if non-nil, zero value otherwise.

### GetADROk

`func (o *PackageLineFiltersModel) GetADROk() (*bool, bool)`

GetADROk returns a tuple with the ADR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADR

`func (o *PackageLineFiltersModel) SetADR(v bool)`

SetADR sets ADR field to given value.

### HasADR

`func (o *PackageLineFiltersModel) HasADR() bool`

HasADR returns a boolean if a field has been set.

### GetADRClass

`func (o *PackageLineFiltersModel) GetADRClass() []int32`

GetADRClass returns the ADRClass field if non-nil, zero value otherwise.

### GetADRClassOk

`func (o *PackageLineFiltersModel) GetADRClassOk() (*[]int32, bool)`

GetADRClassOk returns a tuple with the ADRClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADRClass

`func (o *PackageLineFiltersModel) SetADRClass(v []int32)`

SetADRClass sets ADRClass field to given value.

### HasADRClass

`func (o *PackageLineFiltersModel) HasADRClass() bool`

HasADRClass returns a boolean if a field has been set.

### GetADRUNNr

`func (o *PackageLineFiltersModel) GetADRUNNr() []int32`

GetADRUNNr returns the ADRUNNr field if non-nil, zero value otherwise.

### GetADRUNNrOk

`func (o *PackageLineFiltersModel) GetADRUNNrOk() (*[]int32, bool)`

GetADRUNNrOk returns a tuple with the ADRUNNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADRUNNr

`func (o *PackageLineFiltersModel) SetADRUNNr(v []int32)`

SetADRUNNr sets ADRUNNr field to given value.

### HasADRUNNr

`func (o *PackageLineFiltersModel) HasADRUNNr() bool`

HasADRUNNr returns a boolean if a field has been set.

### GetTemp

`func (o *PackageLineFiltersModel) GetTemp() bool`

GetTemp returns the Temp field if non-nil, zero value otherwise.

### GetTempOk

`func (o *PackageLineFiltersModel) GetTempOk() (*bool, bool)`

GetTempOk returns a tuple with the Temp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemp

`func (o *PackageLineFiltersModel) SetTemp(v bool)`

SetTemp sets Temp field to given value.

### HasTemp

`func (o *PackageLineFiltersModel) HasTemp() bool`

HasTemp returns a boolean if a field has been set.

### GetTempMin

`func (o *PackageLineFiltersModel) GetTempMin() float32`

GetTempMin returns the TempMin field if non-nil, zero value otherwise.

### GetTempMinOk

`func (o *PackageLineFiltersModel) GetTempMinOk() (*float32, bool)`

GetTempMinOk returns a tuple with the TempMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempMin

`func (o *PackageLineFiltersModel) SetTempMin(v float32)`

SetTempMin sets TempMin field to given value.

### HasTempMin

`func (o *PackageLineFiltersModel) HasTempMin() bool`

HasTempMin returns a boolean if a field has been set.

### GetTempMax

`func (o *PackageLineFiltersModel) GetTempMax() float32`

GetTempMax returns the TempMax field if non-nil, zero value otherwise.

### GetTempMaxOk

`func (o *PackageLineFiltersModel) GetTempMaxOk() (*float32, bool)`

GetTempMaxOk returns a tuple with the TempMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempMax

`func (o *PackageLineFiltersModel) SetTempMax(v float32)`

SetTempMax sets TempMax field to given value.

### HasTempMax

`func (o *PackageLineFiltersModel) HasTempMax() bool`

HasTempMax returns a boolean if a field has been set.

### GetHsCode

`func (o *PackageLineFiltersModel) GetHsCode() []string`

GetHsCode returns the HsCode field if non-nil, zero value otherwise.

### GetHsCodeOk

`func (o *PackageLineFiltersModel) GetHsCodeOk() (*[]string, bool)`

GetHsCodeOk returns a tuple with the HsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsCode

`func (o *PackageLineFiltersModel) SetHsCode(v []string)`

SetHsCode sets HsCode field to given value.

### HasHsCode

`func (o *PackageLineFiltersModel) HasHsCode() bool`

HasHsCode returns a boolean if a field has been set.

### GetDescription

`func (o *PackageLineFiltersModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PackageLineFiltersModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PackageLineFiltersModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PackageLineFiltersModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLinks

`func (o *PackageLineFiltersModel) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PackageLineFiltersModel) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PackageLineFiltersModel) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PackageLineFiltersModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetActivityLinks

`func (o *PackageLineFiltersModel) GetActivityLinks() []LinkModel`

GetActivityLinks returns the ActivityLinks field if non-nil, zero value otherwise.

### GetActivityLinksOk

`func (o *PackageLineFiltersModel) GetActivityLinksOk() (*[]LinkModel, bool)`

GetActivityLinksOk returns a tuple with the ActivityLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityLinks

`func (o *PackageLineFiltersModel) SetActivityLinks(v []LinkModel)`

SetActivityLinks sets ActivityLinks field to given value.

### HasActivityLinks

`func (o *PackageLineFiltersModel) HasActivityLinks() bool`

HasActivityLinks returns a boolean if a field has been set.

### GetActivityId

`func (o *PackageLineFiltersModel) GetActivityId() []int32`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *PackageLineFiltersModel) GetActivityIdOk() (*[]int32, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *PackageLineFiltersModel) SetActivityId(v []int32)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *PackageLineFiltersModel) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


