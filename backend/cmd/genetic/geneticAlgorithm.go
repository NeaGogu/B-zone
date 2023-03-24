package genetic

import (
	"bzone/backend/internal/models"
	"fmt"
	fp "github.com/rjNemo/underscore"
	"math/rand"
	"strconv"
	"sync"
)

// MDMTSPInstance is an instance of the Multi-Depot Multiple Traveling Salesman Problem (MDMTSP).
// Activities is the set of locations that need to be visited.
// Depots is the set of depots that can be the starting point of the tours.
// NRoutes is the number of salesmen there are and, thus, the number of routes that they need to travel.
type MDMTSPInstance struct {
	Activities []Pos
	Depots     []Pos
	NRoutes    int
}

// GenAlgHyperParams are the hyperparameters for the genetic algorithm.
// nOffspring is the number of offspring made each generation.
// nParents is the number of parents to be selected each generation.
// nGenerations is the number of generations to run the genetic algorithm for.
// tournamentSize is the number of solutions that take part in the parent selection tournaments.
// Greater tournamentSize = more selection pressure.
// maxMutations is the maximum number of mutations that could possibly happen and
// mutationRate is the chance that a mutation happens. As a result of these two parameters, the actual number of
// mutations that are applied to a solution follows a binomial distribution with n=maxMutations and p=mutationRate.
// crossoverRate the chance that 2 parents are selected and selected with crossover;
// otherwise only a single parent is selected
type GenAlgHyperParams struct {
	NOffspring     int
	NParents       int
	NGenerations   int
	TournamentSize int
	MaxMutations   int
	MutationRate   float64
	CrossoverRate  float64
}

// RunGeneticAlgorithm converts the activities ([]models.ActivityModelBumbal) into the input for GeneticAlgorithm,
// runs GeneticAlgorithm with some default parameters, and, finally, converts the output into a []models.ZoneModel.
func RunGeneticAlgorithm(activities []models.ActivityModelBumbal, nZones, nGenerations int) []models.ZoneModel {
	inst := GenerateMDMTSPInstance(activities, nZones)
	params := GenAlgHyperParams{
		NOffspring: 100, NParents: 100, NGenerations: nGenerations, TournamentSize: 5,
		MaxMutations: 100, MutationRate: 0.05, CrossoverRate: 0.5,
	}
	sol := GeneticAlgorithm(inst, params)
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
func GeneticAlgorithm(inst MDMTSPInstance, params GenAlgHyperParams) Solution {
	var err error
	bestSol := randomSolution(inst)
	bestSol.calcCost()
	population := randomPopulation(inst, params.NOffspring)
	for gen := 0; gen < params.NGenerations; gen++ {
		fmt.Println("gen:", gen, "\tcost:", bestSol.Cost)
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

// selectParents selects nParents with replacement from population using tournament selection with tournamentSize
func selectParents(population Population, nParents, tournamentSize int) Population {
	var wg sync.WaitGroup
	wg.Add(nParents)
	parents := make(Population, nParents)
	for i := 0; i < nParents; i++ {
		// concurrently select parents using tournament selection
		go func(j int) {
			defer wg.Done()
			var err error
			parents[j], err = population.tournamentSelection(tournamentSize)
			if err != nil {
				panic(err)
			}
		}(i)
	}
	// block until all parents have been selected
	wg.Wait()
	return parents
}

// makeOffspring makes nOffspring from the parents. With a rate of
// crossoverRate, 2 parents are randomly selected and are combined using
// crossover, otherwise it is just the parent. The resulting offspring is
// mutated with at most maxMutations, where each mutation has a mutationRate
// chance of happening.
func makeOffspring(inst MDMTSPInstance, parents Population, nOffspring, maxMutations int,
	mutationRate, crossoverRate float64) Population {
	var wg sync.WaitGroup
	wg.Add(nOffspring)
	n := len(parents)
	offspring := make(Population, nOffspring)
	for i := 0; i < nOffspring; i++ {
		// concurrently make all offspring
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
	// block until all offspring have been made
	wg.Wait()
	return offspring
}

// crossover combines parent1 and parent2 and returns the combined Solution.
// crossover works by repeatedly picking a nRoutes-1 route from a parent and
// putting that whole route in the child Solution. The last route is greedily
// made with the remaining points.
func crossover(parent1 Solution, parent2 Solution) Solution {
	nRoutes := len(parent1.Routes)
	child := Solution{Routes: make([]Route, nRoutes)}
	// pick nRoutes-1 routes repeatedly from a random parents and add it to the child
	for i := 0; i < nRoutes-1; i++ {
		if rand.Float64() < 0.5 {
			passOnRoute(&parent1, &parent2, &child, i)
		} else {
			passOnRoute(&parent2, &parent1, &child, i)
		}
	}

	// combine all remaining position into a single slice
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

// passOnRoute removes a random route from parent1, adds that route to the child at index,
// and removes all points in that route from parent2
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
