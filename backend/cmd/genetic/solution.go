package genetic

import (
	"fmt"
	fp "github.com/rjNemo/underscore"
	"math/rand"
)

// Solution consists of a Route and a cost of that Route.
// The cost is saved, so it does not need to be recalculated every time it is needed.
type Solution struct {
	Routes []Route
	Cost   float64
}

// randomSolution initializes and returns a random Solution for the given MDMTSPInstance.
// It first chooses a random depot for each route in the solution, then randomly assigns activities
// from the instance to each route.
func randomSolution(inst MDMTSPInstance) Solution {
	var sol Solution
	sol.Routes = make([]Route, inst.NRoutes)
	// Choose a random depot for every route
	for i := 0; i < inst.NRoutes; i++ {
		sol.Routes[i].Depot = inst.Depots[rand.Intn(len(inst.Depots))]
	}
	// Append the activities to a random route
	for _, activity := range inst.Activities {
		i := rand.Intn(inst.NRoutes)
		sol.Routes[i].Activities = append(sol.Routes[i].Activities, activity)
	}
	// Shuffle the activities of every route to randomize their order
	for i := 0; i < len(sol.Routes); i++ {
		rand.Shuffle(len(sol.Routes[i].Activities), func(a, b int) {
			sol.Routes[i].Activities[a], sol.Routes[i].Activities[b] =
				sol.Routes[i].Activities[b], sol.Routes[i].Activities[a]
		})
	}
	return sol
}

// calcCost calculates the cost of the solution based on several factors. The main factor is the total distance
// of all routes. The function also takes into account the variance of the number of activities in each route,
// which is a measure of how evenly the activities are distributed among the routes.
// Finally, the function penalizes the use of multiple depots by adding a term inversely proportional to the number of
// depots used. A lower cost means the solution is better.
func (sol *Solution) calcCost() {
	sol.Cost = 0
	// Calculate the total distance of all routes
	for _, route := range sol.Routes {
		sol.Cost += route.length()
	}

	// Add penalty for not utilizing as many depots as possible
	if nDepots := sol.nDepots(); nDepots > 0 {
		sol.Cost += sol.Cost / (10.0 * float64(nDepots))
	}

	// Add penalty if the activities are not distributes evenly among routes
	sol.Cost += sol.routeLengthVariance()
}

// nDepots counts the number of unique depots used in the Solution.
func (sol *Solution) nDepots() int {
	return len(fp.Unique(fp.Map(sol.Routes, func(route Route) Pos { return route.Depot })))
}

// zipsPerRouteSum counts the number of unique zipcodes visited by each route separately and
// returns the sum of those counts.
func (sol *Solution) zipsPerRouteSum() int {
	// Count the number of unique zipcodes visited by each route
	zipsPerRoute := fp.Map(sol.Routes, func(route Route) int {
		return len(fp.Unique(fp.Map(route.Activities, func(pos Pos) int { return pos.Zipcode })))
	})
	return fp.Sum(zipsPerRoute)
}

// routeLengthVariance calculates the variance of the number of activities in each route.
func (sol *Solution) routeLengthVariance() float64 {
	if len(sol.Routes) <= 1 {
		return 0
	}
	sum := fp.Reduce(sol.Routes, func(route Route, s int) int {
		return s + len(route.Activities)
	}, 0)
	mean := float64(sum) / float64(len(sol.Routes))
	sos := fp.Reduce(sol.Routes, func(route Route, acc float64) float64 {
		diff := float64(len(route.Activities)) - mean
		return acc + diff*diff
	}, 0)
	variance := sos / float64(len(sol.Routes)-1)
	return variance
}

// mutate applies a random number of mutations to the current solution with a given mutation rate.
// The function randomly selects one of the five mutation
// operators to apply: swap two points in a route, 2-opt within a route, greedily move a point within a route,
// change the depot of a route, or move all activities with the same zipcode from one route to another.
// The number of mutations applied is at most maxMutations and mutationRate is the probability that a mutation will
// occur. The actual number of mutations that are applied follows a binomial distribution with n=maxMutations and
// p=mutationRate.
func (sol *Solution) mutate(inst MDMTSPInstance, maxMutations int, mutationRate float64) {
	for mut := 0; mut < maxMutations; mut++ {
		if rand.Float64() > mutationRate {
			continue
		}
		switch rand.Intn(5) {
		case 0:
			sol.randRouteSwap()
		case 1:
			sol.rand2Opt()
		case 2:
			sol.randRouteGreedy()
		case 3:
			sol.randChangeDepot(inst)
		case 4:
			sol.randMigrateZip()
		}
	}
}

// randSolSwap swaps 2 random activities from 2 random routes.
// During performance testing it was found that when this function was used in mutate()
// it negatively affected the solution quality with the current implementation.
func (sol *Solution) randSolSwap() {
	r0 := rand.Intn(len(sol.Routes))
	for len(sol.Routes[r0].Activities) == 0 {
		r0 = rand.Intn(len(sol.Routes))
	}
	i := rand.Intn(len(sol.Routes[r0].Activities))
	r1 := rand.Intn(len(sol.Routes))
	for len(sol.Routes[r1].Activities) == 0 {
		r1 = rand.Intn(len(sol.Routes))
	}
	j := rand.Intn(len(sol.Routes[r1].Activities))
	err := sol.applySwap(r0, i, r1, j)
	if err != nil {
		panic(err)
	}
}

// randMigrate moves a random activity from a random route to a different random route.
// During performance testing it was found that when this function was used in mutate()
// it negatively affected the solution quality with the current implementation.
func (sol *Solution) randMigrate() {
	r0 := rand.Intn(len(sol.Routes))
	for len(sol.Routes[r0].Activities) == 0 {
		r0 = rand.Intn(len(sol.Routes))
	}
	i := rand.Intn(len(sol.Routes[r0].Activities))
	r1 := rand.Intn(len(sol.Routes))
	for r0 == r1 {
		r1 = rand.Intn(len(sol.Routes))
	}
	j := rand.Intn(len(sol.Routes[r1].Activities) + 1)
	err := sol.applyMigrate(r0, i, r1, j)
	if err != nil {
		panic(err)
	}
}

// randMigrateZip moves all activities with a random zipcode from a random route to a different random route.
func (sol *Solution) randMigrateZip() {
	r0 := rand.Intn(len(sol.Routes))
	for len(sol.Routes[r0].Activities) == 0 {
		r0 = rand.Intn(len(sol.Routes))
	}
	zip := sol.Routes[r0].Activities[rand.Intn(len(sol.Routes[r0].Activities))].Zipcode
	r1 := rand.Intn(len(sol.Routes))
	for r0 == r1 {
		r1 = rand.Intn(len(sol.Routes))
	}
	err := sol.applyMigrateZip(r0, r1, zip)
	if err != nil {
		panic(err)
	}
}

// randRouteSwap swaps two random activities within a random route.
func (sol *Solution) randRouteSwap() {
	r := rand.Intn(len(sol.Routes))
	for len(sol.Routes[r].Activities) == 0 {
		r = rand.Intn(len(sol.Routes))
	}
	n := len(sol.Routes[r].Activities)
	i := rand.Intn(n)
	j := rand.Intn(n)
	err := sol.Routes[r].applySwap(i, j)
	if err != nil {
		panic(err)
	}
}

// rand2Opt performs 2-OPT on two random activities within a random route.
func (sol *Solution) rand2Opt() {
	r := rand.Intn(len(sol.Routes))
	for len(sol.Routes[r].Activities) == 0 {
		r = rand.Intn(len(sol.Routes))
	}
	n := len(sol.Routes[r].Activities)
	i := rand.Intn(n)
	j := rand.Intn(n)
	err := sol.Routes[r].apply2Opt(i, j)
	if err != nil {
		panic(err)
	}
}

// randRouteGreedy greedily moves a random activity to a better position within a random route.
func (sol *Solution) randRouteGreedy() {
	r := rand.Intn(len(sol.Routes))
	for len(sol.Routes[r].Activities) == 0 {
		r = rand.Intn(len(sol.Routes))
	}
	n := len(sol.Routes[r].Activities)
	i := rand.Intn(n)
	err := sol.Routes[r].applyGreedy(i)
	if err != nil {
		panic(err)
	}
}

// randChangeDepot changes the depot of a route to a random (possibly different) depot.
func (sol *Solution) randChangeDepot(inst MDMTSPInstance) {
	r := rand.Intn(len(sol.Routes))
	d := rand.Intn(len(inst.Depots))
	err := sol.Routes[r].applyChangeDepot(inst, d)
	if err != nil {
		panic(err)
	}
}

// applySwap swaps activity i in route r0 with the activity j in route r1.
// Returns an error if r0,r1<0 or r0,r1>=len(sol.routes)
// or i<0 or i>=len(sol.routes[r0].activities)
// or j<0 or j>=len(sol.routes[r1].activities); otherwise returns nil.
func (sol *Solution) applySwap(r0, i, r1, j int) error {
	if r0 < 0 || r0 >= len(sol.Routes) || r1 < 0 || r1 >= len(sol.Routes) {
		return fmt.Errorf("invalid route indices r0=%d, r1=%d; valid range is [0,%d)", r0, r1, len(sol.Routes))
	}
	if i < 0 || i >= len(sol.Routes[r0].Activities) {
		return fmt.Errorf("invalid activity index i=%d for route r0=%d; valid range is [0,%d)",
			i, r0, len(sol.Routes[r0].Activities))
	}
	if j < 0 || j >= len(sol.Routes[r1].Activities) {
		return fmt.Errorf("invalid activity index j=%d for route r1=%d; valid range is [0,%d)",
			j, r1, len(sol.Routes[r1].Activities))
	}

	sol.Routes[r0].Activities[i], sol.Routes[r1].Activities[j] = sol.Routes[r1].Activities[j], sol.Routes[r0].Activities[i]
	return nil
}

// applyMigrate removes activity i from route r0 and inserts it at index j in route r1.
// Returns an error if r0,r1<0 or r0,r1>=len(sol.routes)
// or i<0 or i>=len(sol.routes[r0].activities)
// or j<0 or j>len(sol.routes[r1].activities)
// or r0==r1; otherwise returns nil.
func (sol *Solution) applyMigrate(r0, i, r1, j int) error {
	if r0 < 0 || r0 >= len(sol.Routes) || r1 < 0 || r1 >= len(sol.Routes) {
		return fmt.Errorf("invalid route indices r0=%d, r1=%d; valid range is [0,%d)", r0, r1, len(sol.Routes))
	}
	if r0 == r1 {
		return fmt.Errorf("r0 should not equal r1: r0=%d, r1=%d", r0, r1)
	}

	var act Pos
	var err error
	sol.Routes[r0].Activities, act, err = remove(sol.Routes[r0].Activities, i)
	if err != nil {
		return err
	}
	sol.Routes[r1].Activities, err = insert(sol.Routes[r1].Activities, act, j)
	if err != nil {
		return err
	}
	return nil
}

// applyMigrateZip moves all activities with zipcode zip from route r0 to route r1 greedily
// Returns an error if r0,r1<0 or r0,r1>=len(sol.routes)
// or r0==r1; otherwise returns nil.
func (sol *Solution) applyMigrateZip(r0, r1, zip int) error {
	if r0 < 0 || r0 >= len(sol.Routes) || r1 < 0 || r1 >= len(sol.Routes) {
		return fmt.Errorf("invalid route indices r0=%d, r1=%d; valid range is [0,%d)", r0, r1, len(sol.Routes))
	}
	if r0 == r1 {
		return fmt.Errorf("r0 should not equal r1: r0=%d, r1=%d", r0, r1)
	}

	// Insert activities with zipcode zip in Route r1
	for _, activity := range sol.Routes[r0].Activities {
		if activity.Zipcode == zip {
			sol.Routes[r1].applyInsertGreedy(activity)
		}
	}
	// Remove activities with zipcode zip from Route r0
	sol.Routes[r0].Activities = fp.Filter(sol.Routes[r0].Activities, func(pos Pos) bool { return pos.Zipcode != zip })
	return nil
}

// removeActivities removes the first occurrence of a Pos from the Solution for every Pos in activities
func (sol *Solution) removeActivities(activities []Pos) {
	for _, activity := range activities {
		for r, route := range sol.Routes {
			if i := indexOf(route.Activities, activity); i >= 0 {
				sol.Routes[r].Activities, _, _ = remove(route.Activities, i)
				break
			}
		}
	}
}

// copy deep copies sol and returns the copied Solution.
func (sol *Solution) copy() Solution {
	return Solution{
		Routes: fp.Map(sol.Routes, func(route Route) Route { return route.copy() }),
		Cost:   sol.Cost,
	}
}
