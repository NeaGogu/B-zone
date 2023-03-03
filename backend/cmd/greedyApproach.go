package main

import (
	"fmt"
	"math"
)

type depot struct {
	long int
	lat  int
}

type Cost struct {
	zoneID int
	cost   int
}

type service struct {
	serviceID int
	long      int
	lat       int
	zoneID    int
}

func euclideanDistance(x1, x2, y1, y2 int) (m int) {
	m = int(math.Sqrt(math.Pow(float64(x1-x2), 2) + math.Pow(float64(y1-y2), 2)))
	return
}

func calculateCost(serviceArray []service, depotArray []depot, n int) (costArray []Cost, newCost int) {
	s := 0
	for s <= n {
		totalZoneCost := 0
		for i := range serviceArray {
			if serviceArray[i].zoneID == s {
				totalZoneCost += euclideanDistance(serviceArray[i].lat, depotArray[0].lat, serviceArray[i].long, depotArray[0].long)
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

func calculateMaxDistance(serviceArray []service, depotArray []depot, maxZone Cost) (m service) {
	for i, e := range serviceArray {
		if e.zoneID == maxZone.zoneID && (i == 0 || euclideanDistance(e.lat, depotArray[0].lat, e.long, depotArray[0].long) >
			euclideanDistance(m.lat, depotArray[0].lat, m.long, depotArray[0].long)) {
			m = e
		}
	}
	return
}

func main() {
	var prevCost = math.MaxInt
	var newCost = math.MaxInt
	fmt.Println("prevCost, newCost =", prevCost, newCost)

	var depotArray = []depot{
		{1, 3},
	}
	fmt.Println("depotArray =", depotArray)

	var costArray = []Cost{
		{0, math.MinInt}, //Write function to calculate this cost
		{1, math.MinInt}, //Write function to calculate this cost
	}
	fmt.Println("costArray =", costArray)

	var serviceArray = []service{
		{0, 5, 2, 0},
		{1, 2, 4, 1},
		{2, 3, 5, 1},
		{3, 3, 6, 0},
	}
	fmt.Println("serviceArray =", serviceArray)

	costArray, newCost = calculateCost(serviceArray, depotArray, 1)
	fmt.Println("costArray", costArray)

	for newCost < prevCost {
		//Move the activity that is the furthest from the depot to the zone with the lowest cost
		//1. Pick what zone has the lowest cost
		minZone := MinIntSlice(costArray)
		fmt.Println("minZone =", minZone)

		//2. Pick what zone has the highest cost
		maxZone := MaxIntSlice(costArray)
		fmt.Println("maxZone =", maxZone)

		//3. Pick from the that zone that has the highest cost the service with the largest distance to depot
		maxDistanceService := calculateMaxDistance(serviceArray, depotArray, maxZone)
		fmt.Println("maxDistanceService =", maxDistanceService)

		//4. Move the selected service to the zone with the lowest cost
		serviceArray[maxDistanceService.serviceID].zoneID = minZone.zoneID
		fmt.Println("serviceArray =", serviceArray)

		//Set previous cost to the old current cost
		prevCost = newCost

		//Recalculate costs for zones and total cost
		costArray, newCost = calculateCost(serviceArray, depotArray, 1)
		fmt.Println("costArray, newCost =", costArray, newCost)

		//Print the new costs
		fmt.Println("newCost is", newCost)
		fmt.Println("prevCost is", prevCost)
	}
}
