package main

import (
	"fmt"
	"math/rand"
)

type ZoneRangeModel struct {
	id           int
	zipcode_from int
	zipcode_to   int
}

type ZoneModel struct {
	id          int
	zone_ranges []ZoneRangeModel
}

type AddressModel struct {
	id        int
	latitude  string
	longitude string
}

type ActivityModel struct {
	id               int
	address          AddressModel
	depot_address    AddressModel
	depot_address_id int
	zone             []ZoneModel
}

func main() {
	nrOfZones := 2

	var ActivityArray = []ActivityModel{
		{id: 0}, {id: 1}, {id: 3},
	}

	var ZonesArray [][]ZoneModel
	for i := 0; i < nrOfZones; i++ {
		ZonesArray = append(ZonesArray, []ZoneModel{{
			id:          i,
			zone_ranges: nil,
		}})
	}

	for e := range ActivityArray {
		ActivityArray[e].zone = ZonesArray[rand.Intn(nrOfZones)]
		fmt.Println("Activity id", ActivityArray[e].id, "has zone", ActivityArray[e].zone)
	}
}
