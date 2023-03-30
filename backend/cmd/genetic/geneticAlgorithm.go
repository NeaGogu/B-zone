package genetic

import (
	"bzone/backend/internal/models"
	"errors"
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
// NOffspring is the number of offspring generated in each generation and should be >0.
// NParents is the number of parents selected each generation and should be >0.
// NGenerations is the number of generations to run the genetic algorithm for and should be >0.
// TournamentSize is the number of solutions that take part in the parent selection tournaments and should be >0.
// A greater TournamentSize results in more selection pressure.
// MaxMutations is the maximum number of mutations that could potentially occur in a single solution and should be >0.
// MutationRate is the probability that a mutation will occur during reproduction and should be in the range [0,1].
// The actual number of mutations that are applied to a solution follows a binomial distribution with n=MaxMutations
// and p=MutationRate.
// CrossoverRate is the probability that two parents will be selected for crossover and should be in the range [0,1];
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
// Returns an error if some fields of GenAlgHyperParams are not within their bounds.
func RunGeneticAlgorithm(activities []models.ActivityModelBumbal, nZones, nGenerations int,
	timeout time.Duration) ([]models.ZoneModel, error) {
	inst := GenerateMDMTSPInstance(activities, nZones)
	params := GenAlgHyperParams{
		NOffspring: 100, NParents: 100, NGenerations: nGenerations, TournamentSize: 5,
		MaxMutations: 100, MutationRate: 0.05, CrossoverRate: 0.5,
	}
	sol, err := GeneticAlgorithm(inst, params, timeout)
	if err != nil {
		return nil, err
	}
	return Solution2ZoneModels(sol), nil
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
		var act, depot Pos

		act, err = parseActivity(activity)
		if err == nil {
			inst.Activities[i] = act
		}

		depot, err = parseDepot(activity)
		if err == nil {
			inst.Depots[i] = depot
		}
	}

	// remove duplicate depots
	inst.Depots = fp.Unique(inst.Depots)
	return inst
}

// parseActivity extracts the Pos of the activity from the models.ActivityModelBumbal.
// If either the latitude, longitude, or zip code can't be converted, an error is returned.
func parseActivity(activity models.ActivityModelBumbal) (Pos, error) {
	var err error
	var actLat, actLon float64
	var actZip int
	// convert activity.AddressApplied string values to float64 and int
	actLat, err = strconv.ParseFloat(*activity.AddressApplied.Latitude, 64)
	if err != nil {
		return Pos{}, err
	}
	actLon, err = strconv.ParseFloat(*activity.AddressApplied.Longitude, 64)
	if err != nil {
		return Pos{}, err
	}
	zipStr := *activity.AddressApplied.Zipcode
	if len(zipStr) < 4 {
		return Pos{}, errors.New("zip codes should be at least 4 characters long")
	}
	actZip, err = strconv.Atoi(zipStr[:4])
	if err != nil {
		return Pos{}, err
	}
	return Pos{actLat, actLon, actZip}, nil
}

// parseDepot extracts the Pos of the depot from the models.ActivityModelBumbal.
// If either the latitude, longitude, or zip code can't be converted, an error is returned.
func parseDepot(activity models.ActivityModelBumbal) (Pos, error) {
	var err error
	var depotLat, depotLon float64
	var depotZip int
	// convert activity.DepotAddress string values to float64 and int
	depotLat, err = strconv.ParseFloat(*activity.DepotAddress.Latitude, 64)
	if err != nil {
		return Pos{}, err
	}
	depotLon, err = strconv.ParseFloat(*activity.DepotAddress.Longitude, 64)
	if err != nil {
		return Pos{}, err
	}
	zipStr := *activity.DepotAddress.Zipcode
	if len(zipStr) < 4 {
		return Pos{}, errors.New("zip codes should be at least 4 characters long")
	}
	depotZip, err = strconv.Atoi(zipStr[:4])
	if err != nil {
		return Pos{}, err
	}
	return Pos{depotLat, depotLon, depotZip}, nil
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
// Returns an error if some fields of GenAlgHyperParams are not within their bounds or if an unexpected error occurs.
func GeneticAlgorithm(inst MDMTSPInstance, params GenAlgHyperParams, timeout time.Duration) (Solution, error) {
	err := checkGenAlgHyperParamsBounds(params)
	if err != nil {
		return Solution{}, err
	}

	t0 := time.Now()
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
			return Solution{}, err
		}

		parents, err := selectParents(population, params.NParents, params.TournamentSize)
		if err != nil {
			return Solution{}, err
		}
		population, err = makeOffspring(inst, parents, params.NOffspring, params.MaxMutations, params.MutationRate,
			params.CrossoverRate)
		if err != nil {
			return Solution{}, err
		}
		// elitism
		population[0] = bestSol
	}
	population.calcCosts()
	// ignore error because len(population) always > 0
	bestSol, err = population.getBest()
	if err != nil {
		return Solution{}, err
	}
	return bestSol, nil
}

// checkGenAlgHyperParamsBounds checks if all the fields of params are within their bounds.
// Returns an error if a field is not within bounds; otherwise returns nil.
func checkGenAlgHyperParamsBounds(params GenAlgHyperParams) error {
	if params.NOffspring <= 0 {
		return errors.New("params.NOffspring is not greater than 0")
	} else if params.NParents <= 0 {
		return errors.New("params.NParents is not greater than 0")
	} else if params.TournamentSize <= 0 {
		return errors.New("params.TournamentSize is not greater than 0")
	} else if params.MaxMutations <= 0 {
		return errors.New("params.MaxMutations is not greater than 0")
	} else if params.NGenerations <= 0 {
		return errors.New("params.NGenerations is not greater than 0")
	} else if params.MutationRate < 0 || params.MutationRate > 1 {
		return errors.New("params.MutationRate is not in the range [0,1]")
	} else if params.CrossoverRate < 0 || params.CrossoverRate > 1 {
		return errors.New("params.CrossoverRate is not in the range [0,1]")
	}
	return nil
}

// selectParents selects nParents parents from population using tournament selection with tournamentSize.
// It selects all parents concurrently and waits for all parents to be selected before returning them.
// If len(population)==0 or tournamentSize<=0, an error is returned;
// otherwise, err==nil and the selected parents Population is returned.
func selectParents(population Population, nParents, tournamentSize int) (parents Population, err error) {
	if len(population) == 0 {
		return population, errors.New("population must not be empty")
	} else if tournamentSize <= 0 {
		return population, errors.New("tournamentSize must be positive")
	}

	var wg sync.WaitGroup
	wg.Add(nParents)
	// Create a population of parents
	parents = make(Population, nParents)
	// Concurrently select parents using tournament selection
	for i := 0; i < nParents; i++ {
		go func(j int) {
			defer wg.Done()
			// Error is ignored because len(population)>0 and tournamentSize>0
			parents[j], _ = population.tournamentSelection(tournamentSize)
		}(i)
	}
	// Wait for all parents to be selected before returning them
	wg.Wait()
	return parents, nil
}

// makeOffspring creates nOffspring from the parents. For each offspring, with a probability
// of crossoverRate, two parents are randomly selected and combined using crossover. Otherwise,
// the offspring is a copy of a single parent. The resulting offspring is then mutated with at
// most maxMutations, where each mutation has a mutationRate chance of happening.
// If an error occurs, an error is returned and the offspring Population is nil.
func makeOffspring(inst MDMTSPInstance, parents Population, nOffspring, maxMutations int,
	mutationRate, crossoverRate float64) (Population, error) {
	var wg sync.WaitGroup
	wg.Add(nOffspring)

	// Create a population of offspring
	offspring := make(Population, nOffspring)
	nParents := len(parents)
	var errorOccurred error

	// Concurrently create each offspring
	for i := 0; i < nOffspring; i++ {
		go func(j int) {
			defer wg.Done()
			var err error

			// Select a single parent
			parent1 := parents[rand.Intn(nParents)].copy()

			// With a probability of crossoverRate, select a second parent and combine them
			if rand.Float64() < crossoverRate {
				parent2 := parents[rand.Intn(nParents)].copy()
				offspring[j], err = crossover(parent1, parent2)
				if err != nil {
					errorOccurred = err
					return
				}
			} else {
				offspring[j] = parent1
			}

			// Mutate the offspring with at most maxMutations and mutationRate chance of happening
			offspring[j].mutate(inst, maxMutations, mutationRate)
		}(i)
	}

	// Wait for all offspring to be created before returning the population
	wg.Wait()

	if errorOccurred != nil {
		return nil, errorOccurred
	}

	return offspring, nil
}

// crossover combines parent1 and parent2 to create a child Solution.
// The function first selects nRoutes-1 routes randomly from either parent.
// A route is selected by randomly choosing one parent, and then adding the route to the
// child. The last route is created greedily using all remaining positions.
// The depot of the longest route that remained is chosen as the depot for the greedy route.
func crossover(parent1 Solution, parent2 Solution) (Solution, error) {
	nRoutes := len(parent1.Routes)
	child := Solution{Routes: make([]Route, nRoutes)}
	// pick nRoutes-1 routes repeatedly from a random parent and add it to the child
	for i := 0; i < nRoutes-1; i++ {
		var err error
		if rand.Float64() < 0.5 {
			err = passOnRoute(&parent1, &parent2, &child, i)
		} else {
			err = passOnRoute(&parent2, &parent1, &child, i)
		}
		if err != nil {
			return Solution{}, err
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

	return child.copy(), nil
}

// passOnRoute removes a random route from parent1 and adds that route to the child at the specified index.
// It then removes all points in that route from parent2 to prevent duplicate routes in the child.
// Returns an error is len(parent1.routes)==0
func passOnRoute(parent1, parent2, child *Solution, index int) error {
	var route Route
	var err error
	parent1.Routes, route, err = remove(parent1.Routes, rand.Intn(len(parent1.Routes)))
	if err != nil {
		return err
	}
	child.Routes[index] = route
	parent2.removeActivities(route.Activities)
	return nil
}
