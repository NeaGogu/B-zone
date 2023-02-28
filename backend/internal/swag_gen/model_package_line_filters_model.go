/*
 * Bumbal Client Api
 *
 * Bumbal API documentation
 *
 * API version: 2.0
 * Contact: info@bumbal.eu
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// PackageLineFiltersModel struct for PackageLineFiltersModel
type PackageLineFiltersModel struct {
	// Show updated since
	UpdatedAtSince *time.Time `json:"updated_at_since,omitempty"`
	// Show updated till
	UpdatedAtTill *time.Time `json:"updated_at_till,omitempty"`
	// Bumbal package line id's
	Id *[]int32 `json:"id,omitempty"`
	// PackageLine numbers
	Nr *[]string `json:"nr,omitempty"`
	// StatusIds of PackageLine, 31: package_line_cancelled, 23: package_line_incomplete, 24: package_line_new, 42: package_line_awaiting, 25: package_line_accepted, 10: package_line_planned, 11: package_line_in_progress, 12: package_line_executed
	StatusId *[]int32 `json:"status_id,omitempty"`
	// Active status of PackageLine, 0 values represent deleted PackageLines
	Active *[]int32 `json:"active,omitempty"`
	// PackageLine Status
	StatusName *[]string `json:"status_name,omitempty"`
	// Number of packages in package line
	NrOfPackages *float32 `json:"nr_of_packages,omitempty"`
	// Type of the Packages
	PackageTypeName *[]string `json:"package_type_name,omitempty"`
	// ID of the package type for the packages
	PackageTypeId *[]int32 `json:"package_type_id,omitempty"`
	// 
	AppliedCapacities *map[string]interface{} `json:"applied_capacities,omitempty"`
	// 
	Capacities *[]CapacityModel `json:"capacities,omitempty"`
	// barcode for packages
	Barcode *[]string `json:"barcode,omitempty"`
	// boolean for whether or not the packages in package line should be considered as ADR
	ADR *bool `json:"ADR,omitempty"`
	// ADR class of packages in package line
	ADRClass *[]int32 `json:"ADR_class,omitempty"`
	// ADR UN Nr of packages in package line
	ADRUNNr *[]int32 `json:"ADR_UN_nr,omitempty"`
	// boolean for whether or not the packages in package line should be considered as temperature dependent
	Temp *bool `json:"temp,omitempty"`
	// minimum temperature for packages in package line
	TempMin *float32 `json:"temp_min,omitempty"`
	// maximum temperature for packages in package line
	TempMax *float32 `json:"temp_max,omitempty"`
	// Harmonized System code for packages
	HsCode *[]string `json:"hs_code,omitempty"`
	// description of this package_line
	Description *string `json:"description,omitempty"`
	// 
	Links *[]LinkModel `json:"links,omitempty"`
	// 
	ActivityLinks *[]LinkModel `json:"activity_links,omitempty"`
	// Activity id
	ActivityId *[]int32 `json:"activity_id,omitempty"`
}

// NewPackageLineFiltersModel instantiates a new PackageLineFiltersModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageLineFiltersModel() *PackageLineFiltersModel {
	this := PackageLineFiltersModel{}
	return &this
}

// NewPackageLineFiltersModelWithDefaults instantiates a new PackageLineFiltersModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageLineFiltersModelWithDefaults() *PackageLineFiltersModel {
	this := PackageLineFiltersModel{}
	return &this
}

// GetUpdatedAtSince returns the UpdatedAtSince field value if set, zero value otherwise.
func (o *PackageLineFiltersModel) GetUpdatedAtSince() time.Time {
	if o == nil || o.UpdatedAtSince == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAtSince
}

// GetUpdatedAtSinceOk returns a tuple with the UpdatedAtSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLineFiltersModel) GetUpdatedAtSinceOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAtSince == nil {
		return nil, false
	}
	return o.UpdatedAtSince, true
}

// HasUpdatedAtSince returns a boolean if a field has been set.
func (o *PackageLineFiltersModel) HasUpdatedAtSince() bool {
	if o != nil && o.UpdatedAtSince != nil {
		return true
	}

	return false
}

// SetUpdatedAtSince gets a reference to the given time.Time and assigns it to the UpdatedAtSince field.
func (o *PackageLineFiltersModel) SetUpdatedAtSince(v time.Time) {
	o.UpdatedAtSince = &v
}

// GetUpdatedAtTill returns the UpdatedAtTill field value if set, zero value otherwise.
func (o *PackageLineFiltersModel) GetUpdatedAtTill() time.Time {
	if o == nil || o.UpdatedAtTill == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAtTill
}

// GetUpdatedAtTillOk returns a tuple with the UpdatedAtTill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLineFiltersModel) GetUpdatedAtTillOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAtTill == nil {
		return nil, false
	}
	return o.UpdatedAtTill, true
}

// HasUpdatedAtTill returns a boolean if a field has been set.
func (o *PackageLineFiltersModel) HasUpdatedAtTill() bool {
	if o != nil && o.UpdatedAtTill != nil {
		return true
	}

	return false
}

// SetUpdatedAtTill gets a reference to the given time.Time and assigns it to the UpdatedAtTill field.
func (o *PackageLineFiltersModel) SetUpdatedAtTill(v time.Time) {
	o.UpdatedAtTill = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PackageLineFiltersModel) GetId() []int32 {
	if o == nil || o.Id == nil {
		var ret []int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLineFiltersModel) GetIdOk() (*[]int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PackageLineFiltersModel) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given []int32 and assigns it to the Id field.
func (o *PackageLineFiltersModel) SetId(v []int32) {
	o.Id = &v
}

// GetNr returns the Nr field value if set, zero value otherwise.
func (o *PackageLineFiltersModel) GetNr() []string {
	if o == nil || o.Nr == nil {
		var ret []string
		return ret
	}
	return *o.Nr
}

// GetNrOk returns a tuple with the Nr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLineFiltersModel) GetNrOk() (*[]string, bool) {
	if o == nil || o.Nr == nil {
		return nil, false
	}
	return o.Nr, true
}

// HasNr returns a boolean if a field has been set.
func (o *PackageLineFiltersModel) HasNr() bool {
	if o != nil && o.Nr != nil {
		return true
	}

	return false
}

// SetNr gets a reference to the given []string and assigns it to the Nr field.
func (o *PackageLineFiltersModel) SetNr(v []string) {
	o.Nr = &v
}

// GetStatusId returns the StatusId field value if set, zero value otherwise.
func (o *PackageLineFiltersModel) GetStatusId() []int32 {
	if o == nil || o.StatusId == nil {
		var ret []int32
		return ret
	}
	return *o.StatusId
}

// GetStatusIdOk returns a tuple with the StatusId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLineFiltersModel) GetStatusIdOk() (*[]int32, bool) {
	if o == nil || o.StatusId == nil {
		return nil, false
	}
	return o.StatusId, true
}

// HasStatusId returns a boolean if a field has been set.
func (o *PackageLineFiltersModel) HasStatusId() bool {
	if o != nil && o.StatusId != nil {
		return true
	}

	return false
}

// SetStatusId gets a reference to the given []int32 and assigns it to the StatusId field.
func (o *PackageLineFiltersModel) SetStatusId(v []int32) {
	o.StatusId = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *PackageLineFiltersModel) GetActive() []int32 {
	if o == nil || o.Active == nil {
		var ret []int32
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLineFiltersModel) GetActiveOk() (*[]int32, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *PackageLineFiltersModel) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given []int32 and assigns it to the Active field.
func (o *PackageLineFiltersModel) SetActive(v []int32) {
	o.Active = &v
}

// GetStatusName returns the StatusName field value if set, zero value otherwise.
func (o *PackageLineFiltersModel) GetStatusName() []string {
	if o == nil || o.StatusName == nil {
		var ret []string
		return ret
	}
	return *o.StatusName
}

// GetStatusNameOk returns a tuple with the StatusName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLineFiltersModel) GetStatusNameOk() (*[]string, bool) {
	if o == nil || o.StatusName == nil {
		return nil, false
	}
	return o.StatusName, true
}

// HasStatusName returns a boolean if a field has been set.
func (o *PackageLineFiltersModel) HasStatusName() bool {
	if o != nil && o.StatusName != nil {
		return true
	}

	return false
}

// SetStatusName gets a reference to the given []string and assigns it to the StatusName field.
func (o *PackageLineFiltersModel) SetStatusName(v []string) {
	o.StatusName = &v
}

// GetNrOfPackages returns the NrOfPackages field value if set, zero value otherwise.
func (o *PackageLineFiltersModel) GetNrOfPackages() float32 {
	if o == nil || o.NrOfPackages == nil {
		var ret float32
		return ret
	}
	return *o.NrOfPackages
}

// GetNrOfPackagesOk returns a tuple with the NrOfPackages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLineFiltersModel) GetNrOfPackagesOk() (*float32, bool) {
	if o == nil || o.NrOfPackages == nil {
		return nil, false
	}
	return o.NrOfPackages, true
}

// HasNrOfPackages returns a boolean if a field has been set.
func (o *PackageLineFiltersModel) HasNrOfPackages() bool {
	if o != nil && o.NrOfPackages != nil {
		return true
	}

	return false
}

// SetNrOfPackages gets a reference to the given float32 and assigns it to the NrOfPackages field.
func (o *PackageLineFiltersModel) SetNrOfPackages(v float32) {
	o.NrOfPackages = &v
}

// GetPackageTypeName returns the PackageTypeName field value if set, zero value otherwise.
func (o *PackageLineFiltersModel) GetPackageTypeName() []string {
	if o == nil || o.PackageTypeName == nil {
		var ret []string
		return ret
	}
	return *o.PackageTypeName
}

// GetPackageTypeNameOk returns a tuple with the PackageTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLineFiltersModel) GetPackageTypeNameOk() (*[]string, bool) {
	if o == nil || o.PackageTypeName == nil {
		return nil, false
	}
	return o.PackageTypeName, true
}

// HasPackageTypeName returns a boolean if a field has been set.
func (o *PackageLineFiltersModel) HasPackageTypeName() bool {
	if o != nil && o.PackageTypeName != nil {
		return true
	}

	return false
}

// SetPackageTypeName gets a reference to the given []string and assigns it to the PackageTypeName field.
func (o *PackageLineFiltersModel) SetPackageTypeName(v []string) {
	o.PackageTypeName = &v
}

// GetPackageTypeId returns the PackageTypeId field value if set, zero value otherwise.
func (o *PackageLineFiltersModel) GetPackageTypeId() []int32 {
	if o == nil || o.PackageTypeId == nil {
		var ret []int32
		return ret
	}
	return *o.PackageTypeId
}

// GetPackageTypeIdOk returns a tuple with the PackageTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLineFiltersModel) GetPackageTypeIdOk() (*[]int32, bool) {
	if o == nil || o.PackageTypeId == nil {
		return nil, false
	}
	return o.PackageTypeId, true
}

// HasPackageTypeId returns a boolean if a field has been set.
func (o *PackageLineFiltersModel) HasPackageTypeId() bool {
	if o != nil && o.PackageTypeId != nil {
		return true
	}

	return false
}

// SetPackageTypeId gets a reference to the given []int32 and assigns it to the PackageTypeId field.
func (o *PackageLineFiltersModel) SetPackageTypeId(v []int32) {
	o.PackageTypeId = &v
}

// GetAppliedCapacities returns the AppliedCapacities field value if set, zero value otherwise.
func (o *PackageLineFiltersModel) GetAppliedCapacities() map[string]interface{} {
	if o == nil || o.AppliedCapacities == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.AppliedCapacities
}

// GetAppliedCapacitiesOk returns a tuple with the AppliedCapacities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLineFiltersModel) GetAppliedCapacitiesOk() (*map[string]interface{}, bool) {
	if o == nil || o.AppliedCapacities == nil {
		return nil, false
	}
	return o.AppliedCapacities, true
}

// HasAppliedCapacities returns a boolean if a field has been set.
func (o *PackageLineFiltersModel) HasAppliedCapacities() bool {
	if o != nil && o.AppliedCapacities != nil {
		return true
	}

	return false
}

// SetAppliedCapacities gets a reference to the given map[string]interface{} and assigns it to the AppliedCapacities field.
func (o *PackageLineFiltersModel) SetAppliedCapacities(v map[string]interface{}) {
	o.AppliedCapacities = &v
}

// GetCapacities returns the Capacities field value if set, zero value otherwise.
func (o *PackageLineFiltersModel) GetCapacities() []CapacityModel {
	if o == nil || o.Capacities == nil {
		var ret []CapacityModel
		return ret
	}
	return *o.Capacities
}

// GetCapacitiesOk returns a tuple with the Capacities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLineFiltersModel) GetCapacitiesOk() (*[]CapacityModel, bool) {
	if o == nil || o.Capacities == nil {
		return nil, false
	}
	return o.Capacities, true
}

// HasCapacities returns a boolean if a field has been set.
func (o *PackageLineFiltersModel) HasCapacities() bool {
	if o != nil && o.Capacities != nil {
		return true
	}

	return false
}

// SetCapacities gets a reference to the given []CapacityModel and assigns it to the Capacities field.
func (o *PackageLineFiltersModel) SetCapacities(v []CapacityModel) {
	o.Capacities = &v
}

// GetBarcode returns the Barcode field value if set, zero value otherwise.
func (o *PackageLineFiltersModel) GetBarcode() []string {
	if o == nil || o.Barcode == nil {
		var ret []string
		return ret
	}
	return *o.Barcode
}

// GetBarcodeOk returns a tuple with the Barcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLineFiltersModel) GetBarcodeOk() (*[]string, bool) {
	if o == nil || o.Barcode == nil {
		return nil, false
	}
	return o.Barcode, true
}

// HasBarcode returns a boolean if a field has been set.
func (o *PackageLineFiltersModel) HasBarcode() bool {
	if o != nil && o.Barcode != nil {
		return true
	}

	return false
}

// SetBarcode gets a reference to the given []string and assigns it to the Barcode field.
func (o *PackageLineFiltersModel) SetBarcode(v []string) {
	o.Barcode = &v
}

// GetADR returns the ADR field value if set, zero value otherwise.
func (o *PackageLineFiltersModel) GetADR() bool {
	if o == nil || o.ADR == nil {
		var ret bool
		return ret
	}
	return *o.ADR
}

// GetADROk returns a tuple with the ADR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLineFiltersModel) GetADROk() (*bool, bool) {
	if o == nil || o.ADR == nil {
		return nil, false
	}
	return o.ADR, true
}

// HasADR returns a boolean if a field has been set.
func (o *PackageLineFiltersModel) HasADR() bool {
	if o != nil && o.ADR != nil {
		return true
	}

	return false
}

// SetADR gets a reference to the given bool and assigns it to the ADR field.
func (o *PackageLineFiltersModel) SetADR(v bool) {
	o.ADR = &v
}

// GetADRClass returns the ADRClass field value if set, zero value otherwise.
func (o *PackageLineFiltersModel) GetADRClass() []int32 {
	if o == nil || o.ADRClass == nil {
		var ret []int32
		return ret
	}
	return *o.ADRClass
}

// GetADRClassOk returns a tuple with the ADRClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLineFiltersModel) GetADRClassOk() (*[]int32, bool) {
	if o == nil || o.ADRClass == nil {
		return nil, false
	}
	return o.ADRClass, true
}

// HasADRClass returns a boolean if a field has been set.
func (o *PackageLineFiltersModel) HasADRClass() bool {
	if o != nil && o.ADRClass != nil {
		return true
	}

	return false
}

// SetADRClass gets a reference to the given []int32 and assigns it to the ADRClass field.
func (o *PackageLineFiltersModel) SetADRClass(v []int32) {
	o.ADRClass = &v
}

// GetADRUNNr returns the ADRUNNr field value if set, zero value otherwise.
func (o *PackageLineFiltersModel) GetADRUNNr() []int32 {
	if o == nil || o.ADRUNNr == nil {
		var ret []int32
		return ret
	}
	return *o.ADRUNNr
}

// GetADRUNNrOk returns a tuple with the ADRUNNr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLineFiltersModel) GetADRUNNrOk() (*[]int32, bool) {
	if o == nil || o.ADRUNNr == nil {
		return nil, false
	}
	return o.ADRUNNr, true
}

// HasADRUNNr returns a boolean if a field has been set.
func (o *PackageLineFiltersModel) HasADRUNNr() bool {
	if o != nil && o.ADRUNNr != nil {
		return true
	}

	return false
}

// SetADRUNNr gets a reference to the given []int32 and assigns it to the ADRUNNr field.
func (o *PackageLineFiltersModel) SetADRUNNr(v []int32) {
	o.ADRUNNr = &v
}

// GetTemp returns the Temp field value if set, zero value otherwise.
func (o *PackageLineFiltersModel) GetTemp() bool {
	if o == nil || o.Temp == nil {
		var ret bool
		return ret
	}
	return *o.Temp
}

// GetTempOk returns a tuple with the Temp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLineFiltersModel) GetTempOk() (*bool, bool) {
	if o == nil || o.Temp == nil {
		return nil, false
	}
	return o.Temp, true
}

// HasTemp returns a boolean if a field has been set.
func (o *PackageLineFiltersModel) HasTemp() bool {
	if o != nil && o.Temp != nil {
		return true
	}

	return false
}

// SetTemp gets a reference to the given bool and assigns it to the Temp field.
func (o *PackageLineFiltersModel) SetTemp(v bool) {
	o.Temp = &v
}

// GetTempMin returns the TempMin field value if set, zero value otherwise.
func (o *PackageLineFiltersModel) GetTempMin() float32 {
	if o == nil || o.TempMin == nil {
		var ret float32
		return ret
	}
	return *o.TempMin
}

// GetTempMinOk returns a tuple with the TempMin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLineFiltersModel) GetTempMinOk() (*float32, bool) {
	if o == nil || o.TempMin == nil {
		return nil, false
	}
	return o.TempMin, true
}

// HasTempMin returns a boolean if a field has been set.
func (o *PackageLineFiltersModel) HasTempMin() bool {
	if o != nil && o.TempMin != nil {
		return true
	}

	return false
}

// SetTempMin gets a reference to the given float32 and assigns it to the TempMin field.
func (o *PackageLineFiltersModel) SetTempMin(v float32) {
	o.TempMin = &v
}

// GetTempMax returns the TempMax field value if set, zero value otherwise.
func (o *PackageLineFiltersModel) GetTempMax() float32 {
	if o == nil || o.TempMax == nil {
		var ret float32
		return ret
	}
	return *o.TempMax
}

// GetTempMaxOk returns a tuple with the TempMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLineFiltersModel) GetTempMaxOk() (*float32, bool) {
	if o == nil || o.TempMax == nil {
		return nil, false
	}
	return o.TempMax, true
}

// HasTempMax returns a boolean if a field has been set.
func (o *PackageLineFiltersModel) HasTempMax() bool {
	if o != nil && o.TempMax != nil {
		return true
	}

	return false
}

// SetTempMax gets a reference to the given float32 and assigns it to the TempMax field.
func (o *PackageLineFiltersModel) SetTempMax(v float32) {
	o.TempMax = &v
}

// GetHsCode returns the HsCode field value if set, zero value otherwise.
func (o *PackageLineFiltersModel) GetHsCode() []string {
	if o == nil || o.HsCode == nil {
		var ret []string
		return ret
	}
	return *o.HsCode
}

// GetHsCodeOk returns a tuple with the HsCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLineFiltersModel) GetHsCodeOk() (*[]string, bool) {
	if o == nil || o.HsCode == nil {
		return nil, false
	}
	return o.HsCode, true
}

// HasHsCode returns a boolean if a field has been set.
func (o *PackageLineFiltersModel) HasHsCode() bool {
	if o != nil && o.HsCode != nil {
		return true
	}

	return false
}

// SetHsCode gets a reference to the given []string and assigns it to the HsCode field.
func (o *PackageLineFiltersModel) SetHsCode(v []string) {
	o.HsCode = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PackageLineFiltersModel) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLineFiltersModel) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PackageLineFiltersModel) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PackageLineFiltersModel) SetDescription(v string) {
	o.Description = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PackageLineFiltersModel) GetLinks() []LinkModel {
	if o == nil || o.Links == nil {
		var ret []LinkModel
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLineFiltersModel) GetLinksOk() (*[]LinkModel, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PackageLineFiltersModel) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []LinkModel and assigns it to the Links field.
func (o *PackageLineFiltersModel) SetLinks(v []LinkModel) {
	o.Links = &v
}

// GetActivityLinks returns the ActivityLinks field value if set, zero value otherwise.
func (o *PackageLineFiltersModel) GetActivityLinks() []LinkModel {
	if o == nil || o.ActivityLinks == nil {
		var ret []LinkModel
		return ret
	}
	return *o.ActivityLinks
}

// GetActivityLinksOk returns a tuple with the ActivityLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLineFiltersModel) GetActivityLinksOk() (*[]LinkModel, bool) {
	if o == nil || o.ActivityLinks == nil {
		return nil, false
	}
	return o.ActivityLinks, true
}

// HasActivityLinks returns a boolean if a field has been set.
func (o *PackageLineFiltersModel) HasActivityLinks() bool {
	if o != nil && o.ActivityLinks != nil {
		return true
	}

	return false
}

// SetActivityLinks gets a reference to the given []LinkModel and assigns it to the ActivityLinks field.
func (o *PackageLineFiltersModel) SetActivityLinks(v []LinkModel) {
	o.ActivityLinks = &v
}

// GetActivityId returns the ActivityId field value if set, zero value otherwise.
func (o *PackageLineFiltersModel) GetActivityId() []int32 {
	if o == nil || o.ActivityId == nil {
		var ret []int32
		return ret
	}
	return *o.ActivityId
}

// GetActivityIdOk returns a tuple with the ActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageLineFiltersModel) GetActivityIdOk() (*[]int32, bool) {
	if o == nil || o.ActivityId == nil {
		return nil, false
	}
	return o.ActivityId, true
}

// HasActivityId returns a boolean if a field has been set.
func (o *PackageLineFiltersModel) HasActivityId() bool {
	if o != nil && o.ActivityId != nil {
		return true
	}

	return false
}

// SetActivityId gets a reference to the given []int32 and assigns it to the ActivityId field.
func (o *PackageLineFiltersModel) SetActivityId(v []int32) {
	o.ActivityId = &v
}

func (o PackageLineFiltersModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UpdatedAtSince != nil {
		toSerialize["updated_at_since"] = o.UpdatedAtSince
	}
	if o.UpdatedAtTill != nil {
		toSerialize["updated_at_till"] = o.UpdatedAtTill
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Nr != nil {
		toSerialize["nr"] = o.Nr
	}
	if o.StatusId != nil {
		toSerialize["status_id"] = o.StatusId
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.StatusName != nil {
		toSerialize["status_name"] = o.StatusName
	}
	if o.NrOfPackages != nil {
		toSerialize["nr_of_packages"] = o.NrOfPackages
	}
	if o.PackageTypeName != nil {
		toSerialize["package_type_name"] = o.PackageTypeName
	}
	if o.PackageTypeId != nil {
		toSerialize["package_type_id"] = o.PackageTypeId
	}
	if o.AppliedCapacities != nil {
		toSerialize["applied_capacities"] = o.AppliedCapacities
	}
	if o.Capacities != nil {
		toSerialize["capacities"] = o.Capacities
	}
	if o.Barcode != nil {
		toSerialize["barcode"] = o.Barcode
	}
	if o.ADR != nil {
		toSerialize["ADR"] = o.ADR
	}
	if o.ADRClass != nil {
		toSerialize["ADR_class"] = o.ADRClass
	}
	if o.ADRUNNr != nil {
		toSerialize["ADR_UN_nr"] = o.ADRUNNr
	}
	if o.Temp != nil {
		toSerialize["temp"] = o.Temp
	}
	if o.TempMin != nil {
		toSerialize["temp_min"] = o.TempMin
	}
	if o.TempMax != nil {
		toSerialize["temp_max"] = o.TempMax
	}
	if o.HsCode != nil {
		toSerialize["hs_code"] = o.HsCode
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.ActivityLinks != nil {
		toSerialize["activity_links"] = o.ActivityLinks
	}
	if o.ActivityId != nil {
		toSerialize["activity_id"] = o.ActivityId
	}
	return json.Marshal(toSerialize)
}

type NullablePackageLineFiltersModel struct {
	value *PackageLineFiltersModel
	isSet bool
}

func (v NullablePackageLineFiltersModel) Get() *PackageLineFiltersModel {
	return v.value
}

func (v *NullablePackageLineFiltersModel) Set(val *PackageLineFiltersModel) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageLineFiltersModel) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageLineFiltersModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageLineFiltersModel(val *PackageLineFiltersModel) *NullablePackageLineFiltersModel {
	return &NullablePackageLineFiltersModel{value: val, isSet: true}
}

func (v NullablePackageLineFiltersModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageLineFiltersModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


