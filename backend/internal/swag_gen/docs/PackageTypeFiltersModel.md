# PackageTypeFiltersModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdatedAtSince** | Pointer to **time.Time** | Show updated since | [optional] 
**UpdatedAtTill** | Pointer to **time.Time** | Show updated till | [optional] 
**Id** | Pointer to **[]int32** | Bumbal package line id&#39;s | [optional] 
**Nr** | Pointer to **[]string** | PackageType numbers | [optional] 
**StatusId** | Pointer to **[]int32** | StatusIds of PackageType, 31: package_line_cancelled, 23: package_line_incomplete, 24: package_line_new, 42: package_line_awaiting, 25: package_line_accepted, 10: package_line_planned, 11: package_line_in_progress, 12: package_line_executed | [optional] 
**StatusName** | Pointer to **[]string** | PackageType Status | [optional] 
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

### NewPackageTypeFiltersModel

`func NewPackageTypeFiltersModel() *PackageTypeFiltersModel`

NewPackageTypeFiltersModel instantiates a new PackageTypeFiltersModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageTypeFiltersModelWithDefaults

`func NewPackageTypeFiltersModelWithDefaults() *PackageTypeFiltersModel`

NewPackageTypeFiltersModelWithDefaults instantiates a new PackageTypeFiltersModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdatedAtSince

`func (o *PackageTypeFiltersModel) GetUpdatedAtSince() time.Time`

GetUpdatedAtSince returns the UpdatedAtSince field if non-nil, zero value otherwise.

### GetUpdatedAtSinceOk

`func (o *PackageTypeFiltersModel) GetUpdatedAtSinceOk() (*time.Time, bool)`

GetUpdatedAtSinceOk returns a tuple with the UpdatedAtSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtSince

`func (o *PackageTypeFiltersModel) SetUpdatedAtSince(v time.Time)`

SetUpdatedAtSince sets UpdatedAtSince field to given value.

### HasUpdatedAtSince

`func (o *PackageTypeFiltersModel) HasUpdatedAtSince() bool`

HasUpdatedAtSince returns a boolean if a field has been set.

### GetUpdatedAtTill

`func (o *PackageTypeFiltersModel) GetUpdatedAtTill() time.Time`

GetUpdatedAtTill returns the UpdatedAtTill field if non-nil, zero value otherwise.

### GetUpdatedAtTillOk

`func (o *PackageTypeFiltersModel) GetUpdatedAtTillOk() (*time.Time, bool)`

GetUpdatedAtTillOk returns a tuple with the UpdatedAtTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAtTill

`func (o *PackageTypeFiltersModel) SetUpdatedAtTill(v time.Time)`

SetUpdatedAtTill sets UpdatedAtTill field to given value.

### HasUpdatedAtTill

`func (o *PackageTypeFiltersModel) HasUpdatedAtTill() bool`

HasUpdatedAtTill returns a boolean if a field has been set.

### GetId

`func (o *PackageTypeFiltersModel) GetId() []int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PackageTypeFiltersModel) GetIdOk() (*[]int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PackageTypeFiltersModel) SetId(v []int32)`

SetId sets Id field to given value.

### HasId

`func (o *PackageTypeFiltersModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNr

`func (o *PackageTypeFiltersModel) GetNr() []string`

GetNr returns the Nr field if non-nil, zero value otherwise.

### GetNrOk

`func (o *PackageTypeFiltersModel) GetNrOk() (*[]string, bool)`

GetNrOk returns a tuple with the Nr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNr

`func (o *PackageTypeFiltersModel) SetNr(v []string)`

SetNr sets Nr field to given value.

### HasNr

`func (o *PackageTypeFiltersModel) HasNr() bool`

HasNr returns a boolean if a field has been set.

### GetStatusId

`func (o *PackageTypeFiltersModel) GetStatusId() []int32`

GetStatusId returns the StatusId field if non-nil, zero value otherwise.

### GetStatusIdOk

`func (o *PackageTypeFiltersModel) GetStatusIdOk() (*[]int32, bool)`

GetStatusIdOk returns a tuple with the StatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusId

`func (o *PackageTypeFiltersModel) SetStatusId(v []int32)`

SetStatusId sets StatusId field to given value.

### HasStatusId

`func (o *PackageTypeFiltersModel) HasStatusId() bool`

HasStatusId returns a boolean if a field has been set.

### GetStatusName

`func (o *PackageTypeFiltersModel) GetStatusName() []string`

GetStatusName returns the StatusName field if non-nil, zero value otherwise.

### GetStatusNameOk

`func (o *PackageTypeFiltersModel) GetStatusNameOk() (*[]string, bool)`

GetStatusNameOk returns a tuple with the StatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusName

`func (o *PackageTypeFiltersModel) SetStatusName(v []string)`

SetStatusName sets StatusName field to given value.

### HasStatusName

`func (o *PackageTypeFiltersModel) HasStatusName() bool`

HasStatusName returns a boolean if a field has been set.

### GetNrOfPackages

`func (o *PackageTypeFiltersModel) GetNrOfPackages() float32`

GetNrOfPackages returns the NrOfPackages field if non-nil, zero value otherwise.

### GetNrOfPackagesOk

`func (o *PackageTypeFiltersModel) GetNrOfPackagesOk() (*float32, bool)`

GetNrOfPackagesOk returns a tuple with the NrOfPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrOfPackages

`func (o *PackageTypeFiltersModel) SetNrOfPackages(v float32)`

SetNrOfPackages sets NrOfPackages field to given value.

### HasNrOfPackages

`func (o *PackageTypeFiltersModel) HasNrOfPackages() bool`

HasNrOfPackages returns a boolean if a field has been set.

### GetPackageTypeName

`func (o *PackageTypeFiltersModel) GetPackageTypeName() []string`

GetPackageTypeName returns the PackageTypeName field if non-nil, zero value otherwise.

### GetPackageTypeNameOk

`func (o *PackageTypeFiltersModel) GetPackageTypeNameOk() (*[]string, bool)`

GetPackageTypeNameOk returns a tuple with the PackageTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageTypeName

`func (o *PackageTypeFiltersModel) SetPackageTypeName(v []string)`

SetPackageTypeName sets PackageTypeName field to given value.

### HasPackageTypeName

`func (o *PackageTypeFiltersModel) HasPackageTypeName() bool`

HasPackageTypeName returns a boolean if a field has been set.

### GetPackageTypeId

`func (o *PackageTypeFiltersModel) GetPackageTypeId() []int32`

GetPackageTypeId returns the PackageTypeId field if non-nil, zero value otherwise.

### GetPackageTypeIdOk

`func (o *PackageTypeFiltersModel) GetPackageTypeIdOk() (*[]int32, bool)`

GetPackageTypeIdOk returns a tuple with the PackageTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageTypeId

`func (o *PackageTypeFiltersModel) SetPackageTypeId(v []int32)`

SetPackageTypeId sets PackageTypeId field to given value.

### HasPackageTypeId

`func (o *PackageTypeFiltersModel) HasPackageTypeId() bool`

HasPackageTypeId returns a boolean if a field has been set.

### GetAppliedCapacities

`func (o *PackageTypeFiltersModel) GetAppliedCapacities() map[string]interface{}`

GetAppliedCapacities returns the AppliedCapacities field if non-nil, zero value otherwise.

### GetAppliedCapacitiesOk

`func (o *PackageTypeFiltersModel) GetAppliedCapacitiesOk() (*map[string]interface{}, bool)`

GetAppliedCapacitiesOk returns a tuple with the AppliedCapacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedCapacities

`func (o *PackageTypeFiltersModel) SetAppliedCapacities(v map[string]interface{})`

SetAppliedCapacities sets AppliedCapacities field to given value.

### HasAppliedCapacities

`func (o *PackageTypeFiltersModel) HasAppliedCapacities() bool`

HasAppliedCapacities returns a boolean if a field has been set.

### GetCapacities

`func (o *PackageTypeFiltersModel) GetCapacities() []CapacityModel`

GetCapacities returns the Capacities field if non-nil, zero value otherwise.

### GetCapacitiesOk

`func (o *PackageTypeFiltersModel) GetCapacitiesOk() (*[]CapacityModel, bool)`

GetCapacitiesOk returns a tuple with the Capacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacities

`func (o *PackageTypeFiltersModel) SetCapacities(v []CapacityModel)`

SetCapacities sets Capacities field to given value.

### HasCapacities

`func (o *PackageTypeFiltersModel) HasCapacities() bool`

HasCapacities returns a boolean if a field has been set.

### GetBarcode

`func (o *PackageTypeFiltersModel) GetBarcode() []string`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *PackageTypeFiltersModel) GetBarcodeOk() (*[]string, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *PackageTypeFiltersModel) SetBarcode(v []string)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *PackageTypeFiltersModel) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetADR

`func (o *PackageTypeFiltersModel) GetADR() bool`

GetADR returns the ADR field if non-nil, zero value otherwise.

### GetADROk

`func (o *PackageTypeFiltersModel) GetADROk() (*bool, bool)`

GetADROk returns a tuple with the ADR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADR

`func (o *PackageTypeFiltersModel) SetADR(v bool)`

SetADR sets ADR field to given value.

### HasADR

`func (o *PackageTypeFiltersModel) HasADR() bool`

HasADR returns a boolean if a field has been set.

### GetADRClass

`func (o *PackageTypeFiltersModel) GetADRClass() []int32`

GetADRClass returns the ADRClass field if non-nil, zero value otherwise.

### GetADRClassOk

`func (o *PackageTypeFiltersModel) GetADRClassOk() (*[]int32, bool)`

GetADRClassOk returns a tuple with the ADRClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADRClass

`func (o *PackageTypeFiltersModel) SetADRClass(v []int32)`

SetADRClass sets ADRClass field to given value.

### HasADRClass

`func (o *PackageTypeFiltersModel) HasADRClass() bool`

HasADRClass returns a boolean if a field has been set.

### GetADRUNNr

`func (o *PackageTypeFiltersModel) GetADRUNNr() []int32`

GetADRUNNr returns the ADRUNNr field if non-nil, zero value otherwise.

### GetADRUNNrOk

`func (o *PackageTypeFiltersModel) GetADRUNNrOk() (*[]int32, bool)`

GetADRUNNrOk returns a tuple with the ADRUNNr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADRUNNr

`func (o *PackageTypeFiltersModel) SetADRUNNr(v []int32)`

SetADRUNNr sets ADRUNNr field to given value.

### HasADRUNNr

`func (o *PackageTypeFiltersModel) HasADRUNNr() bool`

HasADRUNNr returns a boolean if a field has been set.

### GetTemp

`func (o *PackageTypeFiltersModel) GetTemp() bool`

GetTemp returns the Temp field if non-nil, zero value otherwise.

### GetTempOk

`func (o *PackageTypeFiltersModel) GetTempOk() (*bool, bool)`

GetTempOk returns a tuple with the Temp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemp

`func (o *PackageTypeFiltersModel) SetTemp(v bool)`

SetTemp sets Temp field to given value.

### HasTemp

`func (o *PackageTypeFiltersModel) HasTemp() bool`

HasTemp returns a boolean if a field has been set.

### GetTempMin

`func (o *PackageTypeFiltersModel) GetTempMin() float32`

GetTempMin returns the TempMin field if non-nil, zero value otherwise.

### GetTempMinOk

`func (o *PackageTypeFiltersModel) GetTempMinOk() (*float32, bool)`

GetTempMinOk returns a tuple with the TempMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempMin

`func (o *PackageTypeFiltersModel) SetTempMin(v float32)`

SetTempMin sets TempMin field to given value.

### HasTempMin

`func (o *PackageTypeFiltersModel) HasTempMin() bool`

HasTempMin returns a boolean if a field has been set.

### GetTempMax

`func (o *PackageTypeFiltersModel) GetTempMax() float32`

GetTempMax returns the TempMax field if non-nil, zero value otherwise.

### GetTempMaxOk

`func (o *PackageTypeFiltersModel) GetTempMaxOk() (*float32, bool)`

GetTempMaxOk returns a tuple with the TempMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempMax

`func (o *PackageTypeFiltersModel) SetTempMax(v float32)`

SetTempMax sets TempMax field to given value.

### HasTempMax

`func (o *PackageTypeFiltersModel) HasTempMax() bool`

HasTempMax returns a boolean if a field has been set.

### GetHsCode

`func (o *PackageTypeFiltersModel) GetHsCode() []string`

GetHsCode returns the HsCode field if non-nil, zero value otherwise.

### GetHsCodeOk

`func (o *PackageTypeFiltersModel) GetHsCodeOk() (*[]string, bool)`

GetHsCodeOk returns a tuple with the HsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsCode

`func (o *PackageTypeFiltersModel) SetHsCode(v []string)`

SetHsCode sets HsCode field to given value.

### HasHsCode

`func (o *PackageTypeFiltersModel) HasHsCode() bool`

HasHsCode returns a boolean if a field has been set.

### GetDescription

`func (o *PackageTypeFiltersModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PackageTypeFiltersModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PackageTypeFiltersModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PackageTypeFiltersModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLinks

`func (o *PackageTypeFiltersModel) GetLinks() []LinkModel`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PackageTypeFiltersModel) GetLinksOk() (*[]LinkModel, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PackageTypeFiltersModel) SetLinks(v []LinkModel)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PackageTypeFiltersModel) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetActivityLinks

`func (o *PackageTypeFiltersModel) GetActivityLinks() []LinkModel`

GetActivityLinks returns the ActivityLinks field if non-nil, zero value otherwise.

### GetActivityLinksOk

`func (o *PackageTypeFiltersModel) GetActivityLinksOk() (*[]LinkModel, bool)`

GetActivityLinksOk returns a tuple with the ActivityLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityLinks

`func (o *PackageTypeFiltersModel) SetActivityLinks(v []LinkModel)`

SetActivityLinks sets ActivityLinks field to given value.

### HasActivityLinks

`func (o *PackageTypeFiltersModel) HasActivityLinks() bool`

HasActivityLinks returns a boolean if a field has been set.

### GetActivityId

`func (o *PackageTypeFiltersModel) GetActivityId() []int32`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *PackageTypeFiltersModel) GetActivityIdOk() (*[]int32, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *PackageTypeFiltersModel) SetActivityId(v []int32)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *PackageTypeFiltersModel) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


