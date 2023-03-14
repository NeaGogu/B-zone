package main

import "fmt"

func main() {
	nRoutes := 2
	depot := Pos{5, 5}
	activities := []Pos{
		{3, 5},
		{6, 5},
		{3, 4},
		{10, 10},
		{1, 2},
		{33, 7},
	}
	inst := VRPInstance{activities, depot, nRoutes}

	sol := geneticAlgorithm(inst, 5, 5, 5, 3, 1, 0.5)
	fmt.Println(sol)
}
