package main

import (
	"fmt"
	"math"
	"strconv"
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

type Cost struct {
	zoneID int
	cost   int
}

func euclideanDistance(a, b, c, d string) (m int) {
	//Ugly but can't find a way to do this better
	x1, _ := strconv.Atoi(a)
	x2, _ := strconv.Atoi(b)
	y1, _ := strconv.Atoi(c)
	y2, _ := strconv.Atoi(d)

	m = int(math.Sqrt(math.Pow(float64(x1-x2), 2) + math.Pow(float64(y1-y2), 2)))
	return
}

// Does not correctly calculate the cost yet, only sums up distances but not total route
func calculateCost(activityArray []ActivityModel, n int) (costArray []Cost, newCost int) {
	s := 0
	for s <= n {
		totalZoneCost := 0
		for i := range activityArray {
			for j := range activityArray[i].zone {
				if activityArray[i].zone[j].id == s {
					totalZoneCost += euclideanDistance(activityArray[i].address.latitude, activityArray[i].depot_address.latitude, activityArray[i].address.longitude, activityArray[i].depot_address.longitude)
				}
			}
		}
		costArray = append(costArray, Cost{s, totalZoneCost})
		s++
	}

	newCost = 0
	for s := range costArray {
		newCost += costArray[s].cost
	}

	return
}

func MinIntSlice(v []Cost) (m Cost) {
	for i, e := range v {
		if i == 0 || e.cost < m.cost {
			m = e
		}
	}
	return
}

func MaxIntSlice(v []Cost) (m Cost) {
	for i, e := range v {
		if i == 0 || e.cost > m.cost {
			m = e
		}
	}
	return
}

func calculateMaxDistance(activityArray []ActivityModel, maxZone Cost) (m ActivityModel) {
	for i, e := range activityArray {
		for j := range e.zone {
			if e.zone[j].id == maxZone.zoneID && (i == 0 || euclideanDistance(e.address.latitude, e.depot_address.latitude, e.address.longitude, e.depot_address.longitude) >
				euclideanDistance(m.address.latitude, m.depot_address.latitude, m.address.longitude, m.depot_address.longitude)) {
				m = e
			}
		}
	}
	return
}

func main() {
	var prevCost = math.MaxInt
	var newCost = math.MaxInt
	fmt.Println("prevCost, newCost =", prevCost, newCost)

	var costArray = []Cost{
		{0, math.MinInt}, //Write function to calculate this cost
		{1, math.MinInt}, //Write function to calculate this cost
	}
	fmt.Println("costArray =", costArray)

	var activityArray = []ActivityModel{
		{
			id: 0,
			address: AddressModel{
				id:        0,
				latitude:  "3",
				longitude: "5",
			},
			depot_address: AddressModel{
				id:        0,
				latitude:  "1",
				longitude: "3",
			},
			zone: []ZoneModel{{
				id:          0,
				zone_ranges: nil,
			}},
		},
		{
			id: 1,
			address: AddressModel{
				id:        1,
				latitude:  "6",
				longitude: "2",
			},
			depot_address: AddressModel{
				id:        0,
				latitude:  "1",
				longitude: "3",
			},
			zone: []ZoneModel{{
				id:          1,
				zone_ranges: nil,
			}},
		},
	}
	fmt.Println("activityArray =", activityArray)

	costArray, newCost = calculateCost(activityArray, 1)
	fmt.Println("costArray", costArray)

	fmt.Println("\n(*) Start looping")
	for newCost < prevCost {
		fmt.Println("\n(*) Next loop")
		//Move the activity that is the furthest from the depot to the zone with the lowest cost
		//1. Pick what zone has the lowest cost
		minZone := MinIntSlice(costArray)
		fmt.Println("minZone =", minZone)

		//2. Pick what zone has the highest cost
		maxZone := MaxIntSlice(costArray)
		fmt.Println("maxZone =", maxZone)

		//3. Pick from the that zone that has the highest cost the service with the largest distance to depot
		maxDistanceService := calculateMaxDistance(activityArray, maxZone)
		fmt.Println("maxDistanceService =", maxDistanceService)

		//4. Move the selected service to the zone with the lowest cost
		activityArray[maxDistanceService.id].zone[0].id = minZone.zoneID
		fmt.Println("serviceArray =", activityArray)

		//Set previous cost to the old current cost
		prevCost = newCost

		//Recalculate costs for zones and total cost
		costArray, newCost = calculateCost(activityArray, 1)
		fmt.Println("costArray, newCost =", costArray, newCost)

		//Print the new costs
		fmt.Println("newCost is", newCost)
		fmt.Println("prevCost is", prevCost)
	}
}
