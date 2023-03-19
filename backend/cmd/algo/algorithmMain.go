package main

import (
	"bzone/backend/cmd/algo/genetic"
	"fmt"
)

func main() {
	inst := genetic.MDVRPInstance{
		Activities: []genetic.Pos{
			{3, 5, 42},
			{6, 5, 42},
			{3, 4, 42},
			{10, 10, 69},
			{1, 2, 42},
			{33, 7, 69},
		},
		Depots:  []genetic.Pos{{5, 5, 42}, {9, 9, 50}},
		NRoutes: 2,
	}
	// inst := MDVRPInstance{
	// 	activities: make([]Pos, 1000),
	// 	depot:      Pos{0, 0},
	// 	nRoutes:    4,
	// }
	// for i := 0; i < 1000; i++ {
	// 	inst.activities[i] = Pos{rand.Float64()*200 - 100, rand.Float64()*200 - 100}
	// }

	sol := genetic.GeneticAlgorithm(inst, 100, 100, 100, 5, 5, 0.5, 0.5)
	fmt.Println("Cost:", sol.Cost)
	for i, route := range sol.Routes {
		fmt.Println("Route", i, "length:", len(route.Activities))
	}
	fmt.Println(sol)
	fmt.Println(genetic.Solution2ZoneMap(sol))
}
