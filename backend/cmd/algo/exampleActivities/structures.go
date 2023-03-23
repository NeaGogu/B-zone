package exampleActivities

type ZoneRangeModel struct {
	Id           int
	Zipcode_from int
	Zipcode_to   int
}

type ZoneModel struct {
	Id          int
	Zone_ranges []ZoneRangeModel
}

type AddressModel struct {
	Id        int
	Latitude  string
	Longitude string
	Zipcode   string
}

type ActivityModel struct {
	Id      int
	Address AddressModel
	//Depot_address    AddressModel
	Longitude        string
	Latitude         string
	Depot_longitude  string
	Depot_latitude   string
	Depot_address_id int
	Zone             []ZoneModel
}
