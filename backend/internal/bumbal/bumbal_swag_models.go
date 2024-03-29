package bumbal

// this file contains models generated by swagger from bumbal's swagger.json
// but they are altered to only include what we need/ fix some missmatching types

// ZoneListResponse struct for ZoneListResponse
type BumbalZoneListResponse struct {
	Items []BumbalZoneModel `json:"items,omitempty"`
	// Count of total items without filters in place
	CountUnfiltered *int32 `json:"count_unfiltered,omitempty"`
}

// ZoneModel struct for ZoneModel
type BumbalZoneModel struct {
	// Unique ID
	Id string `json:"id,omitempty"`
	// Zone Name
	Name string `json:"name,omitempty"`
	// Tag names
	ZoneRanges *[]BumbalZoneRangeModel `json:"zone_ranges"`
}


// ZoneRangeModel struct for ZoneRangeModel
type BumbalZoneRangeModel struct {
	// Unique Zone type ID
	Id string `json:"id,omitempty"`
	// Zipcode range start
	ZipcodeFrom string `json:"zipcode_from"`
	// Zipcode range end
	ZipcodeTo string `json:"zipcode_to"`
	// 
	IsoCountry string `json:"iso_country,omitempty"`
}
