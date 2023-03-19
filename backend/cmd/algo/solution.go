package main

import (
	"errors"
	"fmt"
	fp "github.com/rjNemo/underscore"
	"math"
	"math/rand"
	"sync"
)

type VRPInstance struct {
	activities []Pos
	depots     []Pos
	nRoutes    int
}

// Pos is the position of an activity
type Pos struct {
	latitude  float64
	longitude float64
	zipcode   int
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
		sol.routes[i].depot = inst.depots[rand.Intn(len(inst.depots))]
	}
	for _, activity := range inst.activities {
		i := rand.Intn(inst.nRoutes)
		sol.routes[i].activities = append(sol.routes[i].activities, activity)
	}
	for i := 0; i < len(sol.routes); i++ {
		rand.Shuffle(len(sol.routes[i].activities), func(a, b int) {
			sol.routes[i].activities[a], sol.routes[i].activities[b] =
				sol.routes[i].activities[b], sol.routes[i].activities[a]
		})
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
		if len(r.activities) > 0 {
			sol.cost += dist(r.activities[0], r.depot)
			sol.cost += dist(r.activities[len(r.activities)-1], r.depot)
		}
	}
}

func (sol *Solution) mutate(inst VRPInstance, maxMutations int, mutationRate float64) {
	for mut := 0; mut < maxMutations; mut++ {
		if rand.Float64() > mutationRate {
			continue
		}
		switch rand.Intn(6) {
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
		}
	}
}

func (sol *Solution) randSolSwap() {
	r0 := rand.Intn(len(sol.routes))
	for len(sol.routes[r0].activities) == 0 {
		r0 = rand.Intn(len(sol.routes))
	}
	i := rand.Intn(len(sol.routes[r0].activities))
	r1 := rand.Intn(len(sol.routes))
	for len(sol.routes[r1].activities) == 0 {
		r1 = rand.Intn(len(sol.routes))
	}
	j := rand.Intn(len(sol.routes[r1].activities))
	err := sol.applySwap(r0, i, r1, j)
	if err != nil {
		panic(err)
	}
}

func (sol *Solution) randMigrate() {
	r0 := rand.Intn(len(sol.routes))
	for len(sol.routes[r0].activities) == 0 {
		r0 = rand.Intn(len(sol.routes))
	}
	i := rand.Intn(len(sol.routes[r0].activities))
	r1 := rand.Intn(len(sol.routes))
	for r0 == r1 {
		r1 = rand.Intn(len(sol.routes))
	}
	j := rand.Intn(len(sol.routes[r1].activities) + 1)
	err := sol.applyMigrate(r0, i, r1, j)
	if err != nil {
		panic(err)
	}
}

func (sol *Solution) randRouteSwap() {
	r := rand.Intn(len(sol.routes))
	for len(sol.routes[r].activities) == 0 {
		r = rand.Intn(len(sol.routes))
	}
	n := len(sol.routes[r].activities)
	i := rand.Intn(n)
	j := rand.Intn(n)
	err := sol.routes[r].applySwap(i, j)
	if err != nil {
		panic(err)
	}
}

func (sol *Solution) rand2Opt() {
	r := rand.Intn(len(sol.routes))
	for len(sol.routes[r].activities) == 0 {
		r = rand.Intn(len(sol.routes))
	}
	n := len(sol.routes[r].activities)
	i := rand.Intn(n)
	j := rand.Intn(n)
	err := sol.routes[r].apply2Opt(i, j)
	if err != nil {
		panic(err)
	}
}

func (sol *Solution) randRouteGreedy() {
	r := rand.Intn(len(sol.routes))
	for len(sol.routes[r].activities) == 0 {
		r = rand.Intn(len(sol.routes))
	}
	n := len(sol.routes[r].activities)
	i := rand.Intn(n)
	err := sol.routes[r].applyGreedy(i)
	if err != nil {
		panic(err)
	}
}

func (sol *Solution) randChangeDepot(inst VRPInstance) {
	r := rand.Intn(len(sol.routes))
	d := rand.Intn(len(inst.depots))
	err := sol.routes[r].applyChangeDepot(inst, d)
	if err != nil {
		panic(err)
	}
}

// applySwap swaps activity i in route r0 with the activity j in route r1.
// Returns an error if r0,r1<0 or r0,r1>=len(sol.routes)
// or i<0 or i>=len(sol.routes[r0].activities)
// or j<0 or j>=len(sol.routes[r1].activities)
func (sol *Solution) applySwap(r0, i, r1, j int) error {
	if r0 < 0 || r0 >= len(sol.routes) || r1 < 0 || r1 >= len(sol.routes) {
		return errors.New(fmt.Sprintf("index out of bounds: either r0=%d, r1=%d not in range [0,len(sol.routes)) = [0,%d)", r0, r1, len(sol.routes)))
	}
	if i < 0 || i >= len(sol.routes[r0].activities) {
		return errors.New(fmt.Sprintf("index out of bounds: i=%d not in range [0,len(sol.routes[r0].activities)) = [0,%d)", i, len(sol.routes[r0].activities)))
	}
	if j < 0 || j >= len(sol.routes[r1].activities) {
		return errors.New(fmt.Sprintf("index out of bounds: j=%d not in range [0,len(sol.routes[r1].activities)) = [0,%d)", j, len(sol.routes[r1].activities)))
	}

	sol.routes[r0].activities[i], sol.routes[r1].activities[j] = sol.routes[r1].activities[j], sol.routes[r0].activities[i]
	return nil
}

// applyMigrate removes activity i from route r0 and inserts it at index j in route r1.
// Returns an error if r0,r1<0 or r0,r1>=len(sol.routes)
// or i<0 or i>=len(sol.routes[r0].activities)
// or j<0 or j>len(sol.routes[r1].activities)
// or r0==r1
func (sol *Solution) applyMigrate(r0, i, r1, j int) error {
	if r0 < 0 || r0 >= len(sol.routes) || r1 < 0 || r1 >= len(sol.routes) {
		return errors.New(fmt.Sprintf("index out of bounds: either r0=%d, r1=%d not in range [0,len(sol.routes)) = [0,%d)", r0, r1, len(sol.routes)))
	}
	if r0 == r1 {
		return errors.New(fmt.Sprintf("r0 should not equal r1: r0=%d, r1=%d", r0, r1))
	}

	var act Pos
	var err error
	sol.routes[r0].activities, act, err = remove(sol.routes[r0].activities, i)
	if err != nil {
		return err
	}
	sol.routes[r1].activities, err = insert(sol.routes[r1].activities, act, j)
	if err != nil {
		return err
	}
	return nil
}

// applyGreedy removes activity i from route r and re-inserts it in any route at the position that makes the cost the lowest
func (sol *Solution) applyGreedy(r, i int) error {
	// TODO: applyGreedy
	return nil
}

func (sol *Solution) removePoints(points []Pos) {
	for _, point := range points {
		for r, route := range sol.routes {
			if i := indexOf(route.activities, point); i >= 0 {
				sol.routes[r].activities = append(route.activities[:i], route.activities[i+1:]...)
			}
		}
	}
}

func (sol *Solution) copy() Solution {
	return Solution{
		routes: fp.Map(sol.routes, func(route Route) Route {
			return route.copy()
		}),
		cost: sol.cost,
	}
}

func (sol *Solution) activityCount() int {
	return fp.Reduce(sol.routes,
		func(route Route, count int) int {
			return count + len(route.activities)
		}, 0)
}

// randomPopulation initializes and returns a Population with nIndividuals random Solution
func randomPopulation(inst VRPInstance, nIndividuals int) Population {
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
			if sol.cost < best.cost {
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
		if best == nil || opponent.cost < best.cost {
			best = &opponent
		}
	}
	return (*best).copy(), nil
}

func greedyRoute(depot Pos, points []Pos) Route {
	route := Route{depot, make([]Pos, len(points))}
	if len(points) == 0 {
		return route
	}

	minDist := math.MaxFloat64
	var minI int
	for i, point := range points {
		if d := dist(point, depot); d < minDist {
			minDist = d
			minI = i
		}
	}
	points, route.activities[0], _ = remove(points, minI)

	j := 1
	for len(points) > 0 {
		minDist = math.MaxFloat64
		for i, point := range points {
			if d := dist(point, route.activities[j-1]); d < minDist {
				minDist = d
				minI = i
			}
		}
		points, route.activities[j], _ = remove(points, minI)
		j++
	}

	return route
}

// applySwap swaps activity i and j in the route.
// Returns an error if i,j<0 or i,j>=len(route.activities)
func (route *Route) applySwap(i, j int) error {
	if i < 0 || i >= len(route.activities) || j < 0 || j >= len(route.activities) {
		return errors.New(fmt.Sprintf("index out of bounds: either i=%d, j=%d not in range [0,len(route.activities)) = [0,%d)", i, j, len(route.activities)))
	}
	if len(route.activities) > 1 && i != j {
		route.activities[i], route.activities[j] = route.activities[j], route.activities[i]
	}
	return nil
}

// apply2Opt reverses the slice route.activities[i:j+1] in place.
// Returns an error if i,j<0 or i,j>=len(route.activities)
func (route *Route) apply2Opt(i, j int) error {
	if i < 0 || i >= len(route.activities) || j < 0 || j >= len(route.activities) {
		return errors.New(fmt.Sprintf("index out of bounds: either i=%d, j=%d not in range [0,len(route.activities)) = [0,%d)", i, j, len(route.activities)))
	}
	if len(route.activities) > 1 && i != j {
		if i > j {
			i, j = j, i
		}
		for k := 0; k < (j-i)/2+1; k++ {
			err := route.applySwap(i+k, j-k)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// applyGreedy removes activity i from the route and re-inserts it back in route in the position that makes route the shortest.
// Returns an error if i<0 or i>=len(route.activities)
func (route *Route) applyGreedy(i int) error {
	if i < 0 || i >= len(route.activities) {
		return errors.New(fmt.Sprintf("index out of bounds: i=%d not in range [0,len(route.activities)) = [0,%d)", i, len(route.activities)))
	}
	if len(route.activities) <= 1 {
		return nil
	}

	var act Pos
	var err error
	route.activities, act, err = remove(route.activities, i)
	if err != nil {
		return err
	}

	j := 0
	minCost := dist(route.depot, act) + dist(act, route.activities[0])
	for k, pos := range route.activities[1:] {
		cost := dist(route.activities[k], act) + dist(act, pos)
		if cost < minCost {
			minCost = cost
			j = k + 1
		}
	}
	if dist(route.depot, act)+dist(act, route.activities[len(route.activities)-1]) < minCost {
		j = len(route.activities)
	}
	route.activities, err = insert(route.activities, act, j)
	if err != nil {
		return err
	}
	return nil
}

// applyChangeDepot changes the depot of route
// Returns an error if d<0 or d>=len(inst.depots)
func (route *Route) applyChangeDepot(inst VRPInstance, d int) error {
	if d < 0 || d >= len(inst.depots) {
		return errors.New(fmt.Sprintf("index out of bounds: d=%d not in range [0,len(inst.depots)) = [0,%d)", d, len(inst.depots)))
	}
	route.depot = inst.depots[d]
	return nil
}

func (route *Route) copy() Route {
	return Route{
		depot: route.depot.copy(),
		activities: fp.Map(route.activities, func(pos Pos) Pos {
			return pos.copy()
		}),
	}
}

// insert inserts element a at index i in slice l and returns the resulting slice.
// If 0 <= i <= len(l) then err==nil, otherwise err is an error
func insert[A interface{}](l []A, a A, i int) (slice []A, err error) {
	if i < 0 || i > len(l) {
		return l, errors.New(fmt.Sprintf("index out of bounds: i=%d not in range [0,len(l)] = [0,%d]", i, len(l)))
	}
	l = append(l, a)
	copy(l[i+1:], l[i:])
	l[i] = a
	return l, nil
}

// remove removes the element at index i from slice l and returns the resulting slice together with the element that was removed.
// If 0 <= i < len(l) then err==nil, otherwise err is an error
func remove[A interface{}](l []A, i int) (s []A, elem A, err error) {
	if i < 0 || i >= len(l) {
		var a A
		return l, a, errors.New(fmt.Sprintf("index out of bounds: i=%d not in range [0,len(l)) = [0,%d)", i, len(l)))
	}
	elem = l[i]
	return append(l[:i], l[i+1:]...), elem, nil
}

// indexOf returns the index of a in l. If a is not in l, -1 is returned
func indexOf[A comparable](l []A, a A) int {
	for i, v := range l {
		if v == a {
			return i
		}
	}
	return -1
}

// dist calculates the Euclidean distance between Pos p0 and p1
func dist(p0, p1 Pos) float64 {
	return math.Sqrt(math.Pow(p1.longitude-p0.longitude, 2) + math.Pow(p1.latitude-p0.latitude, 2))
}

func (p *Pos) copy() Pos {
	return Pos{
		latitude:  p.latitude,
		longitude: p.longitude,
		zipcode:   p.zipcode,
	}
}
