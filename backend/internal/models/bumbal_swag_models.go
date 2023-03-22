package models

// Models for retrieving activities from Bumbal
type ActivityListResponseBumbal struct {
	Items           *[]ActivityModelBumbal `json:"items,omitempty"`
	CountFiltered   int32                  `json:"count_filtered,omitempty"`
	CountUnfiltered int32                  `json:"count_unfiltered,omitempty"`
	CountLimited    int32                  `json:"count_limited,omitempty"`
}

type ActivityModelBumbal struct {
	Id             *string                    `json:"id,omitempty"`
	AddressApplied *AddressAppliedModelBumbal `json:"address_applied,omitempty"`
	DepotAddress   *AddressModelBumbal        `json:"depot_address,omitempty"`
}

type AddressAppliedModelBumbal struct {
	Zipcode   *string `json:"zipcode,omitempty"`
	Latitude  *string `json:"latitude,omitempty"`
	Longitude *string `json:"longitude,omitempty"`
}

type AddressModelBumbal struct {
	Zipcode   *string `json:"zipcode,omitempty"`
	Latitude  *string `json:"latitude,omitempty"`
	Longitude *string `json:"longitude,omitempty"`
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ActivityModelBumbal) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// NewActivityModel instantiates a new ActivityModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivityModel() *ActivityModelBumbal {
	this := ActivityModelBumbal{}
	return &this
}

// NewAddressAppliedModel instantiates a new AddressAppliedModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressAppliedModel() *AddressAppliedModelBumbal {
	this := AddressAppliedModelBumbal{}
	return &this
}

// NewAddressModel instantiates a new AddressModelBumbal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressModel() *AddressModelBumbal {
	this := AddressModelBumbal{}
	return &this
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *AddressAppliedModelBumbal) GetLatitude() string {
	if o == nil || o.Latitude == nil {
		var ret string
		return ret
	}
	return *o.Latitude
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *AddressAppliedModelBumbal) GetLongitude() string {
	if o == nil || o.Longitude == nil {
		var ret string
		return ret
	}
	return *o.Longitude
}

// GetZipcode returns the Zipcode field value if set, zero value otherwise.
func (o *AddressAppliedModelBumbal) GetZipcode() string {
	if o == nil || o.Zipcode == nil {
		var ret string
		return ret
	}
	return *o.Zipcode
}

// SetAddressApplied gets a reference to the given AddressAppliedModel and assigns it to the AddressApplied field.
func (o *ActivityModelBumbal) SetAddressApplied(v AddressAppliedModelBumbal) {
	o.AddressApplied = &v
}

// SetDepotAddress gets a reference to the given AddressModel and assigns it to the DepotAddress field.
func (o *ActivityModelBumbal) SetDepotAddress(v AddressModelBumbal) {
	o.DepotAddress = &v
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ActivityModelBumbal) SetId(v string) {
	o.Id = &v
}

// SetLatitude gets a reference to the given string and assigns it to the Latitude field.
func (o *AddressAppliedModelBumbal) SetLatitude(v string) {
	o.Latitude = &v
}

// SetLongitude gets a reference to the given string and assigns it to the Longitude field.
func (o *AddressAppliedModelBumbal) SetLongitude(v string) {
	o.Longitude = &v
}

// SetZipcode gets a reference to the given string and assigns it to the Zipcode field.
func (o *AddressAppliedModelBumbal) SetZipcode(v string) {
	o.Zipcode = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *AddressModelBumbal) GetLatitude() string {
	if o == nil || o.Latitude == nil {
		var ret string
		return ret
	}
	return *o.Latitude
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *AddressModelBumbal) GetLongitude() string {
	if o == nil || o.Longitude == nil {
		var ret string
		return ret
	}
	return *o.Longitude
}

// GetZipcode returns the Zipcode field value if set, zero value otherwise.
func (o *AddressModelBumbal) GetZipcode() string {
	if o == nil || o.Zipcode == nil {
		var ret string
		return ret
	}
	return *o.Zipcode
}

// SetLatitude gets a reference to the given string and assigns it to the Latitude field.
func (o *AddressModelBumbal) SetLatitude(v string) {
	o.Latitude = &v
}

// SetLongitude gets a reference to the given string and assigns it to the Longitude field.
func (o *AddressModelBumbal) SetLongitude(v string) {
	o.Longitude = &v
}

// SetZipcode gets a reference to the given string and assigns it to the Zipcode field.
func (o *AddressModelBumbal) SetZipcode(v string) {
	o.Zipcode = &v
}
