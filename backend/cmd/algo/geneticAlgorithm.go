package main

import (
	"math"
	"math/rand"
)

// geneticAlgorithm runs a genetic algorithm with the specified hyperparameters and returns the best solution found
func geneticAlgorithm(inst VRPInstance, nOffspring, nParents, nGenerations, tournamentSize, nMutations int,
	crossoverRate float64) Solution {
	bestSol := randomSolution(inst)
	bestSol.calcCost()
	population := randomPopulation(inst, nOffspring)
	for gen := 0; gen < nGenerations; gen++ {
		population.calcCosts()
		bestSol = population.getBest()
		parents := selectParents(population, nParents, tournamentSize)
		population = makeOffspring(parents, nOffspring, nMutations, crossoverRate)
		population[0] = bestSol
	}
	population.calcCosts()
	return population.getBest()
}

// selectParents selects nParents with replacement from population using tournament selection with tournament size tournamentSize
func selectParents(population Population, nParents, tournamentSize int) Population {
	parents := make(Population, nParents)
	for i := 0; i < nParents; i++ {
		parents[i] = population.tournamentSelection(tournamentSize)
	}
	return parents
}

func makeOffspring(parents Population, nOffspring, nMutations int, crossoverRate float64) Population {
	n := len(parents)
	offspring := make(Population, nOffspring)
	for i := 0; i < nOffspring; i++ {
		parent1 := parents[rand.Intn(n)]
		if rand.Float64() < crossoverRate {
			parent2 := parents[rand.Intn(n)]
			offspring[i] = crossover(parent1, parent2)
		} else {
			offspring[i] = parent1
		}
		offspring[i].mutate(nMutations)
	}
	return offspring
}

func crossover(parent1 Solution, parent2 Solution) Solution {
	child := Solution{routes: make([]Route, 0)}
	for len(parent1.routes)+len(parent2.routes) > 0 {
		selectParent1 := rand.Float64()
		var route Route
		if len(parent1.routes) > 0 && (len(parent2.routes) == 0 || selectParent1 < 0.5) {
			parent1.routes, route = remove(parent1.routes, rand.Intn(len(parent1.routes)))
			child.routes = append(child.routes, route)
			parent2.removePoints(route.activities, true)
		} else {
			parent2.routes, route = remove(parent2.routes, rand.Intn(len(parent2.routes)))
			child.routes = append(child.routes, route)
			parent1.removePoints(route.activities, true)
		}
	}
	return child
}

// dist calculates the Euclidean distance between Pos p0 and p1
func dist(p0, p1 Pos) float64 {
	return math.Sqrt(math.Pow(p1.longitude-p0.longitude, 2) + math.Pow(p1.latitude-p0.latitude, 2))
}
