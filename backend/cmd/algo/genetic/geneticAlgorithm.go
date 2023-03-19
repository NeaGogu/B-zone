package genetic

import (
	openapi "bzone/backend/internal/swag_gen"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

// GenerateMDVRPInstance converts a slice of openapi.ActivityModel to a MDVRPInstance
func GenerateMDVRPInstance(activities []openapi.ActivityModel, nRoutes int) MDVRPInstance {
	inst := MDVRPInstance{NRoutes: nRoutes}
	inst.Activities = make([]Pos, len(activities))
	depots := make(map[Pos]bool)

	for i, activity := range activities {
		var err error
		var actLat, actLon, depotLat, depotLon float64
		var actZip, depotZip int

		actLat, err = strconv.ParseFloat(*activity.AddressApplied.Latitude, 64)
		if err != nil {
			panic(err)
		}
		actLon, err = strconv.ParseFloat(*activity.AddressApplied.Longitude, 64)
		if err != nil {
			panic(err)
		}
		actZip, err = strconv.Atoi(*activity.AddressApplied.Zipcode)
		if err != nil {
			panic(err)
		}
		inst.Activities[i] = Pos{
			Latitude:  actLat,
			Longitude: actLon,
			Zipcode:   actZip,
		}

		depotLat, err = strconv.ParseFloat(*activity.DepotAddress.Latitude, 64)
		if err != nil {
			panic(err)
		}
		depotLon, err = strconv.ParseFloat(*activity.DepotAddress.Longitude, 64)
		if err != nil {
			panic(err)
		}
		depotZip, err = strconv.Atoi(*activity.DepotAddress.Zipcode)
		if err != nil {
			panic(err)
		}
		depotPos := Pos{
			Latitude:  depotLat,
			Longitude: depotLon,
			Zipcode:   depotZip,
		}
		depots[depotPos] = true
	}

	inst.Depots = make([]Pos, len(depots))
	i := 0
	for depot := range depots {
		inst.Depots[i] = depot
		i++
	}

	return inst
}

// GeneticAlgorithm runs a genetic algorithm with the specified hyperparameters and returns the best solution found
func GeneticAlgorithm(inst MDVRPInstance, nOffspring, nParents, nGenerations, tournamentSize, maxMutations int,
	mutationRate, crossoverRate float64) Solution {
	var err error
	bestSol := randomSolution(inst)
	bestSol.calcCost()
	population := randomPopulation(inst, nOffspring)
	for gen := 0; gen < nGenerations; gen++ {
		population.calcCosts()
		bestSol, err = population.getBest()
		if err != nil {
			panic(err)
		}
		fmt.Println("Generation:", gen, "best cost:", bestSol.Cost)
		parents := selectParents(population, nParents, tournamentSize)
		population = makeOffspring(inst, parents, nOffspring, maxMutations, mutationRate, crossoverRate)
		population[0] = bestSol
	}
	population.calcCosts()
	bestSol, err = population.getBest()
	if err != nil {
		panic(err)
	}
	return bestSol
}

// selectParents selects nParents with replacement from population using tournament selection with tournament size tournamentSize
func selectParents(population Population, nParents, tournamentSize int) Population {
	var wg sync.WaitGroup
	wg.Add(nParents)
	parents := make(Population, nParents)
	for i := 0; i < nParents; i++ {
		go func(j int) {
			defer wg.Done()
			var err error
			parents[j], err = population.tournamentSelection(tournamentSize)
			if err != nil {
				panic(err)
			}
		}(i)
	}
	wg.Wait()
	return parents
}

func makeOffspring(inst MDVRPInstance, parents Population, nOffspring, maxMutations int, mutationRate, crossoverRate float64) Population {
	var wg sync.WaitGroup
	wg.Add(nOffspring)
	n := len(parents)
	offspring := make(Population, nOffspring)
	for i := 0; i < nOffspring; i++ {
		go func(j int) {
			defer wg.Done()
			parent1 := parents[rand.Intn(n)].copy()
			if rand.Float64() < crossoverRate {
				parent2 := parents[rand.Intn(n)].copy()
				offspring[j] = crossover(parent1, parent2)
			} else {
				offspring[j] = parent1
			}
			offspring[j].mutate(inst, maxMutations, mutationRate)
		}(i)
	}
	wg.Wait()
	return offspring
}

func crossover(parent1 Solution, parent2 Solution) Solution {
	nRoutes := len(parent1.Routes)
	child := Solution{Routes: make([]Route, nRoutes)}
	// put nRoutes-1 complete routes from the parents in the child
	for i := 0; i < nRoutes-1; i++ {
		selectParent1 := rand.Float64()
		if selectParent1 < 0.5 {
			passOnRoute(&parent1, &parent2, &child, i)
		} else {
			passOnRoute(&parent2, &parent1, &child, i)
		}
	}

	// combine all remaining points greedily
	remaining := make([]Pos, 0)
	longest := 0
	var longI int
	for i, route := range parent1.Routes {
		if l := len(route.Activities); l > longest {
			longest = l
			longI = i
		}
		for _, activity := range route.Activities {
			remaining = append(remaining, activity)
		}
	}
	depot := parent1.Routes[longI].Depot
	route := greedyRoute(depot, remaining)
	child.Routes[nRoutes-1] = route

	return child.copy()
}

// passOnRoute removes a random route from parent1, adds that route to the child at index, and removes all points in that route from parent2
func passOnRoute(parent1, parent2, child *Solution, index int) {
	var route Route
	var err error
	parent1.Routes, route, err = remove(parent1.Routes, rand.Intn(len(parent1.Routes)))
	if err != nil {
		fmt.Println(err)
		return
	}
	child.Routes[index] = route
	parent2.removePoints(route.Activities)
}
