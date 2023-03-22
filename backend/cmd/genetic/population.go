package genetic

import (
	"errors"
	fp "github.com/rjNemo/underscore"
	"math/rand"
	"sync"
)

type Population []Solution

// randomPopulation initializes and returns a Population with nIndividuals random Solution
func randomPopulation(inst MDVRPInstance, nIndividuals int) Population {
	return fp.Map(make([]Solution, nIndividuals),
		func(t Solution) Solution { return randomSolution(inst) })
}

// calcCosts calculates the cost for all solutions in population
func (population *Population) calcCosts() {
	var wg sync.WaitGroup
	wg.Add(len(*population))
	for i := 0; i < len(*population); i++ {
		go func(sol *Solution) {
			defer wg.Done()
			sol.calcCost()
		}(&(*population)[i])
	}
	wg.Wait()
}

// getBest returns the best Solution in the Population.
// if len(population)==0, returns err!=nil
func (population *Population) getBest() (Solution, error) {
	if len(*population) == 0 {
		return Solution{}, errors.New("population must not be empty")
	}
	sol := fp.Reduce(*population,
		func(sol Solution, best Solution) Solution {
			if sol.Cost < best.Cost {
				return sol
			} else {
				return best
			}
		}, (*population)[0])
	return sol.copy(), nil
}

// tournamentSelection performs tournament selection on the Population.
// tournamentSize is the tournament size. Greater tournament size results in more selection pressure.
// if len(population)==0 || tournamentSize<=0, returns err!=nil
func (population *Population) tournamentSelection(tournamentSize int) (Solution, error) {
	if len(*population) == 0 {
		return Solution{}, errors.New("population must not be empty")
	} else if tournamentSize <= 0 {
		return Solution{}, errors.New("tournamentSize must be positive")
	}
	n := len(*population)
	var best *Solution
	for i := 0; i < tournamentSize; i++ {
		opponent := (*population)[rand.Intn(n)]
		if best == nil || opponent.Cost < best.Cost {
			best = &opponent
		}
	}
	return (*best).copy(), nil
}
