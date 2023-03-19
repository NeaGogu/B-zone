package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// geneticAlgorithm runs a genetic algorithm with the specified hyperparameters and returns the best solution found
func geneticAlgorithm(inst VRPInstance, nOffspring, nParents, nGenerations, tournamentSize, maxMutations int,
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
		fmt.Println("Generation:", gen, "best cost:", bestSol.cost)
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

func makeOffspring(inst VRPInstance, parents Population, nOffspring, maxMutations int, mutationRate, crossoverRate float64) Population {
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
	nRoutes := len(parent1.routes)
	child := Solution{routes: make([]Route, nRoutes)}
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
	for i, route := range parent1.routes {
		if l := len(route.activities); l > longest {
			longest = l
			longI = i
		}
		for _, activity := range route.activities {
			remaining = append(remaining, activity)
		}
	}
	depot := parent1.routes[longI].depot
	route := greedyRoute(depot, remaining)
	child.routes[nRoutes-1] = route

	return child.copy()
}

// passOnRoute removes a random route from parent1, adds that route to the child at index, and removes all points in that route from parent2
func passOnRoute(parent1, parent2, child *Solution, index int) {
	var route Route
	var err error
	parent1.routes, route, err = remove(parent1.routes, rand.Intn(len(parent1.routes)))
	if err != nil {
		fmt.Println(err)
		return
	}
	child.routes[index] = route
	parent2.removePoints(route.activities)
}
