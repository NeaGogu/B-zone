package bumbal

// Models for retrieving activities from Bumbal
type ActivityListResponseBumbal struct {
	Items           *[]ActivityModelBumbal `json:"items,omitempty"`
	CountFiltered   int32                  `json:"count_filtered,omitempty"`
	CountUnfiltered int32                  `json:"count_unfiltered,omitempty"`
	CountLimited    int32                  `json:"count_limited,omitempty"`
}

type ActivityModelBumbal struct {
	Id             string                    `json:"id,omitempty"`
	AddressApplied AddressAppliedModelBumbal `json:"address_applied,omitempty"`
}

type AddressAppliedModelBumbal struct {
	Zipcode   string `json:"zipcode,omitempty"`
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
}
