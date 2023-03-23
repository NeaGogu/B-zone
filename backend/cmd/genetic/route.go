package genetic

import (
	"errors"
	"fmt"
	fp "github.com/rjNemo/underscore"
	"math"
)

// Route represents a tour stating at depot and going through all activities
// (depot -> activities[0] -> activities[1] -> ... -> activities[n] -> depot)
type Route struct {
	Depot      Pos
	Activities []Pos
}

// greedyRoute greedily builds a Route from the points and depot.
// Starting from the depot repeatedly the next Pos in the Route is the closest point to the last added Pos.
func greedyRoute(depot Pos, points []Pos) Route {
	route := Route{depot, make([]Pos, len(points))}
	if len(points) == 0 {
		return route
	}

	// Find the closest point to the depot
	minDist := math.MaxFloat64
	var minI int
	for i, point := range points {
		if d := dist(point, depot); d < minDist {
			minDist = d
			minI = i
		}
	}
	points, route.Activities[0], _ = remove(points, minI)

	// repeatedly add the point that is closest to the last added point
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
// Returns an error if i,j<0 or i,j>=len(route.activities); otherwise returns nil
func (route *Route) applySwap(i, j int) error {
	if i < 0 || i >= len(route.Activities) || j < 0 || j >= len(route.Activities) {
		return errors.New(
			fmt.Sprintf("index out of bounds: either i=%d, j=%d not in range [0,len(route.activities)) = [0,%d)",
				i, j, len(route.Activities)))
	}
	if len(route.Activities) > 1 && i != j {
		route.Activities[i], route.Activities[j] = route.Activities[j], route.Activities[i]
	}
	return nil
}

// apply2Opt reverses the slice route.activities[i:j+1] in place.
// Returns an error if i,j<0 or i,j>=len(route.activities); otherwise returns nil
func (route *Route) apply2Opt(i, j int) error {
	if i < 0 || i >= len(route.Activities) || j < 0 || j >= len(route.Activities) {
		return errors.New(
			fmt.Sprintf("index out of bounds: either i=%d, j=%d not in range [0,len(route.activities)) = [0,%d)",
				i, j, len(route.Activities)))
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

// applyGreedy removes activity i from the route and re-inserts it back in route
// in the position that makes route the shortest.
// Returns an error if i<0 or i>=len(route.activities); otherwise returns nil
func (route *Route) applyGreedy(i int) error {
	if i < 0 || i >= len(route.Activities) {
		return errors.New(
			fmt.Sprintf("index out of bounds: i=%d not in range [0,len(route.activities)) = [0,%d)",
				i, len(route.Activities)))
	}
	if len(route.Activities) <= 1 {
		return nil
	}

	var act Pos
	var err error
	// get and remove the activity at index i
	route.Activities, act, err = remove(route.Activities, i)
	if err != nil {
		return err
	}

	// find the best place to insert the activity
	j := 0
	minCost := dist(route.Depot, act) + dist(act, route.Activities[0]) - dist(route.Depot, route.Activities[0])
	for k, pos := range route.Activities[1:] {
		cost := dist(route.Activities[k], act) + dist(act, pos) - dist(route.Activities[k], pos)
		if cost < minCost {
			minCost = cost
			j = k + 1
		}
	}
	if dist(route.Depot, act)+dist(act, route.Activities[len(route.Activities)-1])-
		dist(route.Depot, route.Activities[len(route.Activities)-1]) < minCost {
		j = len(route.Activities)
	}
	// insert the activity in the best place
	route.Activities, err = insert(route.Activities, act, j)
	if err != nil {
		return err
	}
	return nil
}

// applyInsertGreedy inserts pos greedily in route.
func (route *Route) applyInsertGreedy(pos Pos) {
	route.Activities = append(route.Activities, pos)
	// error can be ignored because it will never happen.
	// given: len(route.Activities) >= 1
	// 0 <= len(route.Activities)-1 < len(route.Activities) always holds
	_ = route.applyGreedy(len(route.Activities) - 1)
}

// applyChangeDepot changes the Depot of route
// Returns an error if d<0 or d>=len(inst.depots); otherwise returns nil
func (route *Route) applyChangeDepot(inst MDMTSPInstance, d int) error {
	if d < 0 || d >= len(inst.Depots) {
		return errors.New(fmt.Sprintf("index out of bounds: d=%d not in range [0,len(inst.depots)) = [0,%d)",
			d, len(inst.Depots)))
	}
	route.Depot = inst.Depots[d]
	return nil
}

// copy deep copies route and returns the copied Route.
func (route *Route) copy() Route {
	return Route{
		Depot:      route.Depot.copy(),
		Activities: fp.Map(route.Activities, func(pos Pos) Pos { return pos.copy() }),
	}
}
