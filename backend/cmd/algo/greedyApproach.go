package main

import (
	structures "bzone/backend/cmd/algo/exampleActivities"
	"encoding/json"
	"fmt"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geo"
	"math"
	"os"
	"strconv"
	"strings"
)

type Cost struct {
	zoneID int
	cost   float64
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

func geoDistance(a, b, c, d string) float64 {
	x1, _ := strconv.ParseFloat(strings.Replace(a, ",", ".", 1), 64)
	x2, _ := strconv.ParseFloat(strings.Replace(b, ",", ".", 1), 64)
	y1, _ := strconv.ParseFloat(strings.Replace(c, ",", ".", 1), 64)
	y2, _ := strconv.ParseFloat(strings.Replace(d, ",", ".", 1), 64)

	return geo.Distance(orb.Point{x1, y1}, orb.Point{x2, y2})
}

func contains(s []structures.ZoneModel, str structures.ZoneModel) bool {
	for _, e := range s {
		if e.Id == str.Id {
			return true
		}
	}

	return false
}

// Does not correctly calculate the cost yet, only sums up distances but not total route
func calculateCost(activityArray []structures.ActivityModel, n int) (costArray []Cost, newCost float64) {
	s := 0
	for s <= n {
		totalZoneCost := 0.0
		for i := range activityArray {
			for j := range activityArray[i].Zone {
				if activityArray[i].Zone[j].Id == s {
					totalZoneCost += geoDistance(activityArray[i].Latitude, activityArray[i].Depot_latitude, activityArray[i].Depot_longitude, activityArray[i].Depot_longitude)
				}
			}
		}
		costArray = append(costArray, Cost{s, totalZoneCost})
		s++
	}

	newCost = 0.0
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

func calculateMaxDistance(activityArray []structures.ActivityModel, maxZone Cost) (m structures.ActivityModel) {
	for i, e := range activityArray {
		for j := range e.Zone {
			if e.Zone[j].Id == maxZone.zoneID && (i == 0 || geoDistance(e.Latitude, e.Depot_latitude, e.Longitude, e.Depot_longitude) >
				geoDistance(m.Latitude, m.Depot_latitude, m.Longitude, m.Depot_longitude)) {
				m = e
			}
		}
	}
	return
}

func main() {
	var prevCost = math.MaxFloat64
	var newCost = math.MaxFloat64
	fmt.Println("prevCost, newCost =", prevCost, newCost)

	var costArray = []Cost{
		{0, math.MaxFloat64}, //Write function to calculate this cost
		{1, math.MaxFloat64}, //Write function to calculate this cost
	}
	fmt.Println("costArray =", costArray)

	/*
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
		}*/

	fmt.Println("(*) Starting with initialization now")
	// read our opened jsonFile as a byte array.
	data, err := os.ReadFile("./exampleActivities/csvjson.json")
	if err != nil {
		fmt.Print(err)
	}

	var activityArray []structures.ActivityModel

	err = json.Unmarshal(data, &activityArray)
	if err != nil {
		fmt.Println("error:", err)
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
		for i := 0; i < len(activityArray[maxDistanceService.Id].Zone); i++ {
			activityArray[maxDistanceService.Id].Zone[i].Id = minZone.zoneID
		}
		//fmt.Println("serviceArray =", activityArray)

		//Set previous cost to the old current cost
		prevCost = newCost

		//Recalculate costs for zones and total cost
		costArray, newCost = calculateCost(activityArray, 1)
		fmt.Println("costArray, newCost =", costArray, newCost)

		//Print the new costs
		fmt.Println("newCost is", newCost)
		fmt.Println("prevCost is", prevCost)
	}

	var allZones []structures.ZoneModel

	for i := range activityArray {
		for j := range activityArray[i].Zone {
			if !contains(allZones, activityArray[i].Zone[j]) {
				allZones = append(allZones, activityArray[i].Zone[j])
			}
		}
	}

	fmt.Println("The final list of all zones is stored in: allZones and is equal to", allZones)
}
