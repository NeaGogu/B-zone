package main

import (
	fp "github.com/rjNemo/underscore"
	"math"
	"math/rand"
)

type VRPInstance struct {
	activities []Pos
	depot      Pos
	nRoutes    int
}

// Pos is the position of an activity
type Pos struct {
	latitude  float64
	longitude float64
}

// Route represents a tour stating at depot and going through all activities
// (depot -> activities[0] -> activities[1] -> ... -> activities[n] -> depot)
type Route struct {
	depot      Pos
	activities []Pos
}

type Solution struct {
	routes []Route
	cost   float64
}

type Population []Solution

// randomSolution initializes and returns a random Solution
func randomSolution(inst VRPInstance) Solution {
	var sol Solution
	sol.routes = make([]Route, inst.nRoutes)
	for i := 0; i < inst.nRoutes; i++ {
		sol.routes[i].depot = inst.depot
	}
	for _, activity := range inst.activities {
		i := rand.Intn(inst.nRoutes)
		sol.routes[i].activities = append(sol.routes[i].activities, activity)
	}
	return sol
}

// calcCost calculates the cost of the Solution. The cost is the total distance of all routes
func (sol *Solution) calcCost() {
	sol.cost = 0
	for _, r := range sol.routes {
		if len(r.activities) == 0 {
			continue
		}
		for i, pos := range r.activities[1:] {
			sol.cost += dist(pos, r.activities[i])
		}
		sol.cost += dist(r.activities[0], r.depot)
		sol.cost += dist(r.activities[len(r.activities)-1], r.depot)
	}
}

func (sol *Solution) mutate(nMutations int) {
	for i := 0; i < nMutations; i++ {
		r0 := rand.Intn(len(sol.routes))
		r1 := rand.Intn(len(sol.routes))
		for r0 == r1 {
			r1 = rand.Intn(len(sol.routes))
		}
		i := rand.Intn(len(sol.routes[r0].activities))

		switch rand.Intn(4) {
		case 0:
			j := rand.Intn(len(sol.routes[r1].activities))
			sol.applySwap(r0, i, r1, j)
		case 1:
			j := rand.Intn(len(sol.routes[r1].activities) + 1)
			sol.applyMigrate(r0, i, r1, j)
		case 2:
			j := rand.Intn(len(sol.routes[r0].activities))
			for i == j {
				j = rand.Intn(len(sol.routes[r0].activities))
			}
			sol.routes[r0].applySwap(i, j)
		case 3:
			j := rand.Intn(len(sol.routes[r0].activities))
			for i == j {
				j = rand.Intn(len(sol.routes[r0].activities))
			}
			sol.routes[r0].apply2Opt(i, j)
		}
	}
}

// applySwap swaps activity i in route r0 with the activity j in route r1
func (sol *Solution) applySwap(r0, i, r1, j int) {
	sol.routes[r0].activities[i], sol.routes[r1].activities[j] = sol.routes[r1].activities[j], sol.routes[r0].activities[i]
}

// applyMigrate removes activity i from route r0 and inserts it at index j in route r1
func (sol *Solution) applyMigrate(r0, i, r1, j int) {
	var act Pos
	sol.routes[r0].activities, act = remove(sol.routes[r0].activities, i)
	sol.routes[r1].activities = insert(sol.routes[r1].activities, act, j)
}

// applyGreedy removes activity i from route r and re-inserts it in any route at the position that makes the cost the lowest
func (sol *Solution) applyGreedy(r, i int) {
	// TODO: applyGreedy
}

func (sol *Solution) removePoints(points []Pos, removeRoutes bool) {
	for _, point := range points {
		for r, route := range sol.routes {
			if i := indexOf(route.activities, point); i >= 0 {
				sol.routes[r].activities = append(route.activities[:i], route.activities[i+1:]...)
			}
		}
	}
	if removeRoutes {
		for i := 0; i < len(sol.routes); i++ {
			if len(sol.routes[i].activities) == 0 {
				sol.routes = append(sol.routes[:i], sol.routes[i+1:]...)
				i--
			}
		}
	}
}

// randomPopulation initializes and returns a Population with nIndividuals random Solution
func randomPopulation(inst VRPInstance, nIndividuals int) Population {
	return fp.Map(make([]Solution, nIndividuals),
		func(t Solution) Solution { return randomSolution(inst) })
}

// calcCosts calculates the cost for all solutions in population
func (population *Population) calcCosts() {
	for i := 0; i < len(*population); i++ {
		(*population)[i].calcCost()
	}
}

// getBest returns the best Solution in the Population
func (population *Population) getBest() Solution {
	return fp.Reduce(*population,
		func(sol Solution, best Solution) Solution {
			if sol.cost < best.cost {
				return sol
			} else {
				return best
			}
		}, (*population)[0])
}

// tournamentSelection performs tournament selection on the Population.
// tournamentSize is the tournament size. Greater tournament size results in more selection pressure
func (population *Population) tournamentSelection(tournamentSize int) Solution {
	n := len(*population)
	var best *Solution
	for i := 0; i < tournamentSize; i++ {
		opponent := (*population)[rand.Intn(n)]
		if best == nil || opponent.cost < best.cost {
			best = &opponent
		}
	}
	return *best
}

// applySwap swaps activity i and j in the route
func (route *Route) applySwap(i, j int) {
	route.activities[i], route.activities[j] = route.activities[j], route.activities[i]
}

// apply2Opt reverses the slice route.activities[i:j+1] in place. i,j < len(route.activities)
func (route *Route) apply2Opt(i, j int) {
	if i > j {
		i, j = j, i
	}
	for k := 0; k < (j-i)/2+1; k++ {
		route.applySwap(i+k, j-k)
	}
}

// applyGreedy removes activity i from the route and re-inserts it back in route in the position that makes route the shortest
func (route *Route) applyGreedy(i int) {
	var act Pos
	route.activities, act = remove(route.activities, i)
	var j int
	minCost := math.MaxFloat64
	// FIXME: needs to include depot
	for k := 0; k < len(route.activities)-1; k++ {
		cost := dist(act, route.activities[k]) + dist(act, route.activities[k+1])
		if cost < minCost {
			minCost = cost
			j = k
		}
	}
	route.activities = insert(route.activities, act, j)
}

// insert inserts element a at index i in slice l and returns the resulting slice
func insert[A interface{}](l []A, a A, i int) []A {
	if i < 0 || i >= len(l) {
		panic("index out of bounds. ensure that 0 <= i <= len(l)")
	}
	tmp := append(l[:i], a)
	if i < len(l) {
		return append(tmp, l[i:]...)
	} else {
		return tmp
	}
}

// remove removes the element at index i from slice l and returns the resulting slice together with the element that was removed
func remove[A interface{}](l []A, i int) (s []A, elem A) {
	if i < 0 || i >= len(l) {
		panic("index out of bounds. ensure that 0 <= i < len(l)")
	}
	return append(l[:i], l[i+1:]...), l[i]
}

func indexOf[A comparable](l []A, a A) int {
	for i, v := range l {
		if v == a {
			return i
		}
	}
	return -1
}
