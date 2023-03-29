package genetic

import (
	"bzone/backend/internal/models"
	fp "github.com/rjNemo/underscore"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// MDMTSPInstance is an instance of the Multi-Depot Multiple Traveling Salesman Problem (MDMTSP).
// The problem consists of finding a set of tours for NRoutes salesmen to visit all of the locations in Activities,
// starting and ending at one of the depots in Depots.
// Activities is the set of locations that need to be visited.
// Depots is the set of depots that can be the starting point of the tours.
// NRoutes is the number of salesmen there are and, thus, the number of routes that they need to travel.
type MDMTSPInstance struct {
	Activities []Pos
	Depots     []Pos
	NRoutes    int
}

// GenAlgHyperParams contains the hyperparameters for the genetic algorithm.
// NOffspring is the number of offspring generated in each generation.
// NParents is the number of parents selected each generation.
// NGenerations is the number of generations to run the genetic algorithm for.
// TournamentSize is the number of solutions that take part in the parent selection tournaments.
// A greater TournamentSize results in more selection pressure.
// MaxMutations is the maximum number of mutations that could potentially occur in a single solution.
// MutationRate is the probability that a mutation will occur during reproduction.
// The actual number of mutations that are applied to a solution follows a binomial distribution with n=MaxMutations
// and p=MutationRate.
// CrossoverRate is the probability that two parents will be selected for crossover;
// otherwise, only a single parent is selected.
type GenAlgHyperParams struct {
	NOffspring     int
	NParents       int
	NGenerations   int
	TournamentSize int
	MaxMutations   int
	MutationRate   float64
	CrossoverRate  float64
}

// RunGeneticAlgorithm converts the list of activities ([]models.ActivityModelBumbal) into an MDMTSPInstance,
// then solved the MDMTSPInstance by running the genetic algorithm with default parameters and a specified timeout.
// Finally, it converts the output into a list of zones ([]models.ZoneModel).
// If the timeout elapses, the algorithm is aborted and the best solution found so far is returned.
func RunGeneticAlgorithm(activities []models.ActivityModelBumbal, nZones, nGenerations int,
	timeout time.Duration) []models.ZoneModel {
	inst := GenerateMDMTSPInstance(activities, nZones)
	params := GenAlgHyperParams{
		NOffspring: 100, NParents: 100, NGenerations: nGenerations, TournamentSize: 5,
		MaxMutations: 100, MutationRate: 0.05, CrossoverRate: 0.5,
	}
	sol := GeneticAlgorithm(inst, params, timeout)
	return Solution2ZoneModels(sol)
}

// GenerateMDMTSPInstance converts a slice of models.ActivityModelBumbal to a MDMTSPInstance
func GenerateMDMTSPInstance(activities []models.ActivityModelBumbal, nRoutes int) MDMTSPInstance {
	inst := MDMTSPInstance{NRoutes: nRoutes}
	inst.Activities = make([]Pos, len(activities))
	inst.Depots = make([]Pos, len(activities))
	if len(activities) == 0 {
		return inst
	}

	for i, activity := range activities {
		var err error
		var actLat, actLon, depotLat, depotLon float64
		var actZip, depotZip int

		// convert activity.AddressApplied string values to float64 and int
		actLat, err = strconv.ParseFloat(*activity.AddressApplied.Latitude, 64)
		if err != nil {
			panic(err)
		}
		actLon, err = strconv.ParseFloat(*activity.AddressApplied.Longitude, 64)
		if err != nil {
			panic(err)
		}
		zipStr := *activity.AddressApplied.Zipcode
		actZip, err = strconv.Atoi(zipStr[:4])
		if err != nil {
			panic(err)
		}
		// add an activity with the just converted values
		inst.Activities[i] = Pos{
			Latitude:  actLat,
			Longitude: actLon,
			Zipcode:   actZip,
		}

		// convert activity.DepotAddress string values to float64 and int
		depotLat, err = strconv.ParseFloat(*activity.DepotAddress.Latitude, 64)
		if err != nil {
			panic(err)
		}
		depotLon, err = strconv.ParseFloat(*activity.DepotAddress.Longitude, 64)
		if err != nil {
			panic(err)
		}
		zipStr = *activity.DepotAddress.Zipcode
		depotZip, err = strconv.Atoi(zipStr[:4])
		if err != nil {
			panic(err)
		}
		// add a depot with the just converted values
		inst.Depots[i] = Pos{
			Latitude:  depotLat,
			Longitude: depotLon,
			Zipcode:   depotZip,
		}
	}

	// remove duplicate depots
	inst.Depots = fp.Unique(inst.Depots)
	return inst
}

// Solution2ZoneModels converts a Solution to a slice of models.ZoneModel
func Solution2ZoneModels(solution Solution) []models.ZoneModel {
	zones := make([]models.ZoneModel, len(solution.Routes))
	for i, route := range solution.Routes {
		// get a slice of all unique zipcodes in this route
		zips := fp.Unique(fp.Map(route.Activities, func(pos Pos) int { return pos.Zipcode }))
		zones[i] = models.ZoneModel{
			// converts every zipcode to a ZoneRangeModel of only a single zipcode
			ZoneRanges: fp.Map(zips, func(zip int) models.ZoneRangeModel {
				return models.ZoneRangeModel{ZipcodeFrom: zip, ZipcodeTo: zip}
			}),
		}
	}
	return zones
}

// GeneticAlgorithm runs a genetic algorithm with the specified hyperparameters and returns the best solution found.
// The algorithm iteratively evolves a population of solutions by selecting the best solutions from the current
// generation and creating new offspring through crossover and mutations. Elitism is also applied, to ensure that the
// best solution from the previous generation is always included in the next generation. Once the specified timeout
// duration has elapsed, the algorithm is aborted and the best solution found so far is returned.
func GeneticAlgorithm(inst MDMTSPInstance, params GenAlgHyperParams, timeout time.Duration) Solution {
	t0 := time.Now()
	var err error
	bestSol := randomSolution(inst)
	bestSol.calcCost()
	population := randomPopulation(inst, params.NOffspring)
	for generation := 0; generation < params.NGenerations; generation++ {
		elapsedTime := time.Since(t0)
		if timeout < elapsedTime {
			break
		}

		population.calcCosts()
		// save the best solution of the previous generation
		bestSol, err = population.getBest()
		if err != nil {
			panic(err)
		}
		parents := selectParents(population, params.NParents, params.TournamentSize)
		population = makeOffspring(inst, parents, params.NOffspring,
			params.MaxMutations, params.MutationRate, params.CrossoverRate)
		// elitism
		population[0] = bestSol
	}
	population.calcCosts()
	bestSol, err = population.getBest()
	if err != nil {
		panic(err)
	}
	return bestSol
}

// selectParents selects nParents parents from population using tournament selection with tournamentSize.
// It selects all parents concurrently and waits for all parents to be selected before returning them.
func selectParents(population Population, nParents, tournamentSize int) Population {
	var wg sync.WaitGroup
	wg.Add(nParents)
	// Create a population of parents
	parents := make(Population, nParents)
	// Concurrently select parents using tournament selection
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
	// Wait for all parents to be selected before returning them
	wg.Wait()
	return parents
}

// makeOffspring creates nOffspring from the parents. For each offspring, with a probability
// of crossoverRate, two parents are randomly selected and combined using crossover. Otherwise,
// the offspring is a copy of a single parent. The resulting offspring is then mutated with at
// most maxMutations, where each mutation has a mutationRate chance of happening.
func makeOffspring(inst MDMTSPInstance, parents Population, nOffspring, maxMutations int,
	mutationRate, crossoverRate float64) Population {
	var wg sync.WaitGroup
	wg.Add(nOffspring)

	// Create a population of offspring
	offspring := make(Population, nOffspring)
	nParents := len(parents)

	// Concurrently create each offspring
	for i := 0; i < nOffspring; i++ {
		go func(j int) {
			defer wg.Done()

			// Select a single parent
			parent1 := parents[rand.Intn(nParents)].copy()

			// With a probability of crossoverRate, select a second parent and combine them
			if rand.Float64() < crossoverRate {
				parent2 := parents[rand.Intn(nParents)].copy()
				offspring[j] = crossover(parent1, parent2)
			} else {
				offspring[j] = parent1
			}

			// Mutate the offspring with at most maxMutations and mutationRate chance of happening
			offspring[j].mutate(inst, maxMutations, mutationRate)
		}(i)
	}

	// Wait for all offspring to be created before returning the population
	wg.Wait()
	return offspring
}

// crossover combines parent1 and parent2 to create a child Solution.
// The function first selects nRoutes-1 routes randomly from either parent.
// A route is selected by randomly choosing one parent, and then adding the route to the
// child. The last route is created greedily using all remaining positions.
// The depot of the longest route that remained is chosen as the depot for the greedy route.
func crossover(parent1 Solution, parent2 Solution) Solution {
	nRoutes := len(parent1.Routes)
	child := Solution{Routes: make([]Route, nRoutes)}
	// pick nRoutes-1 routes repeatedly from a random parent and add it to the child
	for i := 0; i < nRoutes-1; i++ {
		if rand.Float64() < 0.5 {
			passOnRoute(&parent1, &parent2, &child, i)
		} else {
			passOnRoute(&parent2, &parent1, &child, i)
		}
	}

	// combine all remaining positions into a single slice
	remainingPositions := make([]Pos, 0)
	// the depot of the longest route that still remained will be the depot of the greedy route
	longest := 0
	var longI int
	for i, route := range parent1.Routes {
		if l := len(route.Activities); l > longest {
			longest = l
			longI = i
		}
		for _, activity := range route.Activities {
			remainingPositions = append(remainingPositions, activity)
		}
	}
	depot := parent1.Routes[longI].Depot
	// combine all remaining positions greedily into a single route
	route := greedyRoute(depot, remainingPositions)
	child.Routes[nRoutes-1] = route

	return child.copy()
}

// passOnRoute removes a random route from parent1 and adds that route to the child at the specified index.
// It then removes all points in that route from parent2 to prevent duplicate routes in the child.
func passOnRoute(parent1, parent2, child *Solution, index int) {
	var route Route
	var err error
	parent1.Routes, route, err = remove(parent1.Routes, rand.Intn(len(parent1.Routes)))
	if err != nil {
		panic(err)
	}
	child.Routes[index] = route
	parent2.removeActivities(route.Activities)
}
