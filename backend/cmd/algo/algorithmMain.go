package main

import (
	"bzone/backend/cmd/algo/genetic"
	model "bzone/backend/internal/models"
	"fmt"
	"math/rand"
)

func main() {
	// inst := genetic.MDVRPInstance{
	// 	Activities: []genetic.Pos{
	// 		{3, 5, 42},
	// 		{6, 5, 42},
	// 		{3, 4, 42},
	// 		{10, 10, 69},
	// 		{1, 2, 42},
	// 		{33, 7, 69},
	// 	},
	// 	Depots:  []genetic.Pos{{5, 5, 42}, {9, 9, 50}},
	// 	NRoutes: 2,
	// }

	nActivities := 100
	activities := make([]model.ActivityModelBumbal, nActivities)
	for i := 0; i < nActivities; i++ {
		lat := rand.Intn(100)
		lon := rand.Intn(100)
		dLat := rand.Intn(2)*50 + 25
		dLon := rand.Intn(2)*50 + 25
		activities[i] = *makeActivity(
			fmt.Sprintf("%d", lat),
			fmt.Sprintf("%d", lon),
			fmt.Sprintf("%d", int(lat)/25*10+int(lon)/25),
			fmt.Sprintf("%d", dLat),
			fmt.Sprintf("%d", dLon),
			fmt.Sprintf("%d", int(dLat)/25*10+int(dLon)/25),
		)
	}
	inst := genetic.GenerateMDVRPInstance(activities, 12)

	sol := genetic.GeneticAlgorithm(inst, 100, 100, 1000, 5, 100, 0.05, 0.5)
	fmt.Println("Cost:", sol.Cost)
	for i, route := range sol.Routes {
		fmt.Println("Route", i, "length:", len(route.Activities))
	}
	fmt.Println(sol)
	// fmt.Println(genetic.Solution2ZoneMap(sol))
	fmt.Println(genetic.Solution2ZoneModels(sol))
}

func makeActivity(actLat string, actLon string, actZip string, depotLat string, depotLon string, depotZip string) *model.ActivityModelBumbal {
	activity := model.NewActivityModel()
	address := model.NewAddressAppliedModel()
	address.SetLatitude(actLat)
	address.SetLongitude(actLon)
	address.SetZipcode(actZip)
	activity.SetAddressApplied(*address)

	depotAddress := model.NewAddressModel()
	depotAddress.SetLatitude(depotLat)
	depotAddress.SetLongitude(depotLon)
	depotAddress.SetZipcode(depotZip)
	activity.SetDepotAddress(*depotAddress)

	return activity
}
