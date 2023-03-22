package genetic

import (
	"errors"
	"fmt"
	fp "github.com/rjNemo/underscore"
	"math"
	"math/rand"
)

type Solution struct {
	Routes []Route
	Cost   float64
}

// randomSolution initializes and returns a random Solution
func randomSolution(inst MDVRPInstance) Solution {
	var sol Solution
	sol.Routes = make([]Route, inst.NRoutes)
	for i := 0; i < inst.NRoutes; i++ {
		sol.Routes[i].Depot = inst.Depots[rand.Intn(len(inst.Depots))]
	}
	for _, activity := range inst.Activities {
		i := rand.Intn(inst.NRoutes)
		sol.Routes[i].Activities = append(sol.Routes[i].Activities, activity)
	}
	for i := 0; i < len(sol.Routes); i++ {
		rand.Shuffle(len(sol.Routes[i].Activities), func(a, b int) {
			sol.Routes[i].Activities[a], sol.Routes[i].Activities[b] =
				sol.Routes[i].Activities[b], sol.Routes[i].Activities[a]
		})
	}
	return sol
}

// calcCost calculates the Cost of the Solution. The Cost is the total distance of all Routes
func (sol *Solution) calcCost() {
	sol.Cost = 0
	for _, r := range sol.Routes {
		if len(r.Activities) == 0 {
			continue
		}
		for i, pos := range r.Activities[1:] {
			sol.Cost += dist(pos, r.Activities[i])
		}
		if len(r.Activities) > 0 {
			sol.Cost += dist(r.Activities[0], r.Depot)
			sol.Cost += dist(r.Activities[len(r.Activities)-1], r.Depot)
		}
	}

	sol.Cost += sol.routeLengthVariance() + 1.0

	sol.Cost += float64(sol.zipsPerRouteSum()) * 10.0

	sol.Cost += 100.0 / float64(sol.nDepots())
}

func (sol *Solution) nDepots() int {
	return len(fp.Unique(fp.Map(sol.Routes, func(route Route) Pos {
		return route.Depot
	})))
}

func (sol *Solution) zipsPerRouteSum() int {
	zipsPerRoute := fp.Map(sol.Routes, func(route Route) int {
		return len(fp.Unique(fp.Map(route.Activities, func(pos Pos) int { return pos.Zipcode })))
	})
	return fp.Sum(zipsPerRoute)
}

func (sol *Solution) routeLengthVariance() float64 {
	sum := fp.Reduce(sol.Routes, func(route Route, s int) int {
		return s + len(route.Activities)
	}, 0)
	mean := float64(sum) / float64(len(sol.Routes))
	sos := fp.Reduce(sol.Routes, func(route Route, acc float64) float64 {
		return acc + math.Pow(float64(len(route.Activities))-mean, 2)
	}, 0)
	variance := sos / float64(len(sol.Routes)-1)
	return variance
}

func (sol *Solution) mutate(inst MDVRPInstance, maxMutations int, mutationRate float64) {
	for mut := 0; mut < maxMutations; mut++ {
		if rand.Float64() > mutationRate {
			continue
		}
		switch rand.Intn(7) {
		case 0:
			sol.randSolSwap()
		case 1:
			sol.randMigrate()
		case 2:
			sol.randRouteSwap()
		case 3:
			sol.rand2Opt()
		case 4:
			sol.randRouteGreedy()
		case 5:
			sol.randChangeDepot(inst)
		case 6:
			sol.randMigrateZip()
		}
	}
}

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

func (sol *Solution) randChangeDepot(inst MDVRPInstance) {
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
// or j<0 or j>=len(sol.routes[r1].activities)
func (sol *Solution) applySwap(r0, i, r1, j int) error {
	if r0 < 0 || r0 >= len(sol.Routes) || r1 < 0 || r1 >= len(sol.Routes) {
		return errors.New(fmt.Sprintf("index out of bounds: either r0=%d, r1=%d not in range [0,len(sol.routes)) = [0,%d)", r0, r1, len(sol.Routes)))
	}
	if i < 0 || i >= len(sol.Routes[r0].Activities) {
		return errors.New(fmt.Sprintf("index out of bounds: i=%d not in range [0,len(sol.routes[r0].activities)) = [0,%d)", i, len(sol.Routes[r0].Activities)))
	}
	if j < 0 || j >= len(sol.Routes[r1].Activities) {
		return errors.New(fmt.Sprintf("index out of bounds: j=%d not in range [0,len(sol.routes[r1].activities)) = [0,%d)", j, len(sol.Routes[r1].Activities)))
	}

	sol.Routes[r0].Activities[i], sol.Routes[r1].Activities[j] = sol.Routes[r1].Activities[j], sol.Routes[r0].Activities[i]
	return nil
}

// applyMigrate removes activity i from route r0 and inserts it at index j in route r1.
// Returns an error if r0,r1<0 or r0,r1>=len(sol.routes)
// or i<0 or i>=len(sol.routes[r0].activities)
// or j<0 or j>len(sol.routes[r1].activities)
// or r0==r1
func (sol *Solution) applyMigrate(r0, i, r1, j int) error {
	if r0 < 0 || r0 >= len(sol.Routes) || r1 < 0 || r1 >= len(sol.Routes) {
		return errors.New(fmt.Sprintf("index out of bounds: either r0=%d, r1=%d not in range [0,len(sol.routes)) = [0,%d)", r0, r1, len(sol.Routes)))
	}
	if r0 == r1 {
		return errors.New(fmt.Sprintf("r0 should not equal r1: r0=%d, r1=%d", r0, r1))
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
// or r0==r1
func (sol *Solution) applyMigrateZip(r0, r1, zip int) error {
	if r0 < 0 || r0 >= len(sol.Routes) || r1 < 0 || r1 >= len(sol.Routes) {
		return errors.New(fmt.Sprintf("index out of bounds: either r0=%d, r1=%d not in range [0,len(sol.routes)) = [0,%d)", r0, r1, len(sol.Routes)))
	}
	if r0 == r1 {
		return errors.New(fmt.Sprintf("r0 should not equal r1: r0=%d, r1=%d", r0, r1))
	}

	for _, activity := range sol.Routes[r0].Activities {
		if activity.Zipcode == zip {
			sol.Routes[r1].applyInsertGreedy(activity)
		}
	}
	sol.Routes[r0].Activities = fp.Filter(sol.Routes[r0].Activities, func(pos Pos) bool {
		return pos.Zipcode != zip
	})
	return nil
}

func (sol *Solution) removePoints(points []Pos) {
	for _, point := range points {
		for r, route := range sol.Routes {
			if i := indexOf(route.Activities, point); i >= 0 {
				sol.Routes[r].Activities = append(route.Activities[:i], route.Activities[i+1:]...)
			}
		}
	}
}

func (sol *Solution) copy() Solution {
	return Solution{
		Routes: fp.Map(sol.Routes, func(route Route) Route {
			return route.copy()
		}),
		Cost: sol.Cost,
	}
}

func (sol *Solution) activityCount() int {
	return fp.Reduce(sol.Routes,
		func(route Route, count int) int {
			return count + len(route.Activities)
		}, 0)
}