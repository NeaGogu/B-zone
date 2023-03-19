package genetic

import (
	"errors"
	"fmt"
	fp "github.com/rjNemo/underscore"
	"math"
	"math/rand"
	"sync"
)

type MDVRPInstance struct {
	Activities []Pos
	Depots     []Pos
	NRoutes    int
}

// Pos is the position of an activity
type Pos struct {
	Latitude  float64
	Longitude float64
	Zipcode   int
}

// Route represents a tour stating at depot and going through all activities
// (depot -> activities[0] -> activities[1] -> ... -> activities[n] -> depot)
type Route struct {
	Depot      Pos
	Activities []Pos
}

type Solution struct {
	Routes []Route
	Cost   float64
}

type Population []Solution

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
}

func (sol *Solution) mutate(inst MDVRPInstance, maxMutations int, mutationRate float64) {
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
	points, route.Activities[0], _ = remove(points, minI)

	j := 1
	for len(points) > 0 {
		minDist = math.MaxFloat64
		for i, point := range points {
			if d := dist(point, route.Activities[j-1]); d < minDist {
				minDist = d
				minI = i
			}
		}
		points, route.Activities[j], _ = remove(points, minI)
		j++
	}

	return route
}

// applySwap swaps activity i and j in the route.
// Returns an error if i,j<0 or i,j>=len(route.activities)
func (route *Route) applySwap(i, j int) error {
	if i < 0 || i >= len(route.Activities) || j < 0 || j >= len(route.Activities) {
		return errors.New(fmt.Sprintf("index out of bounds: either i=%d, j=%d not in range [0,len(route.activities)) = [0,%d)", i, j, len(route.Activities)))
	}
	if len(route.Activities) > 1 && i != j {
		route.Activities[i], route.Activities[j] = route.Activities[j], route.Activities[i]
	}
	return nil
}

// apply2Opt reverses the slice route.activities[i:j+1] in place.
// Returns an error if i,j<0 or i,j>=len(route.activities)
func (route *Route) apply2Opt(i, j int) error {
	if i < 0 || i >= len(route.Activities) || j < 0 || j >= len(route.Activities) {
		return errors.New(fmt.Sprintf("index out of bounds: either i=%d, j=%d not in range [0,len(route.activities)) = [0,%d)", i, j, len(route.Activities)))
	}
	if len(route.Activities) > 1 && i != j {
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
	if i < 0 || i >= len(route.Activities) {
		return errors.New(fmt.Sprintf("index out of bounds: i=%d not in range [0,len(route.activities)) = [0,%d)", i, len(route.Activities)))
	}
	if len(route.Activities) <= 1 {
		return nil
	}

	var act Pos
	var err error
	route.Activities, act, err = remove(route.Activities, i)
	if err != nil {
		return err
	}

	j := 0
	minCost := dist(route.Depot, act) + dist(act, route.Activities[0])
	for k, pos := range route.Activities[1:] {
		cost := dist(route.Activities[k], act) + dist(act, pos)
		if cost < minCost {
			minCost = cost
			j = k + 1
		}
	}
	if dist(route.Depot, act)+dist(act, route.Activities[len(route.Activities)-1]) < minCost {
		j = len(route.Activities)
	}
	route.Activities, err = insert(route.Activities, act, j)
	if err != nil {
		return err
	}
	return nil
}

// applyChangeDepot changes the Depot of route
// Returns an error if d<0 or d>=len(inst.depots)
func (route *Route) applyChangeDepot(inst MDVRPInstance, d int) error {
	if d < 0 || d >= len(inst.Depots) {
		return errors.New(fmt.Sprintf("index out of bounds: d=%d not in range [0,len(inst.depots)) = [0,%d)", d, len(inst.Depots)))
	}
	route.Depot = inst.Depots[d]
	return nil
}

func (route *Route) copy() Route {
	return Route{
		Depot: route.Depot.copy(),
		Activities: fp.Map(route.Activities, func(pos Pos) Pos {
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
	return math.Sqrt(math.Pow(p1.Longitude-p0.Longitude, 2) + math.Pow(p1.Latitude-p0.Latitude, 2))
}

func (p *Pos) copy() Pos {
	return Pos{
		Latitude:  p.Latitude,
		Longitude: p.Longitude,
		Zipcode:   p.Zipcode,
	}
}
