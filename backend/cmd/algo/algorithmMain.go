package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// inst := VRPInstance{
	// 	activities: []Pos{
	// 		{3, 5},
	// 		{6, 5},
	// 		{3, 4},
	// 		{10, 10},
	// 		{1, 2},
	// 		{33, 7},
	// 	},
	// 	depot:   Pos{5, 5},
	// 	nRoutes: 2,
	// }
	inst := VRPInstance{
		activities: make([]Pos, 1000),
		depot:      Pos{0, 0},
		nRoutes:    4,
	}
	for i := 0; i < 1000; i++ {
		inst.activities[i] = Pos{rand.Float64()*200 - 100, rand.Float64()*200 - 100}
	}

	sol := geneticAlgorithm(inst, 100, 100, 100, 5, 5, 0.5, 0.5)
	fmt.Println("Cost:", sol.cost)
	for i, route := range sol.routes {
		fmt.Println("Route", i, "length:", len(route.activities))
	}
	// fmt.Println(sol)
}
