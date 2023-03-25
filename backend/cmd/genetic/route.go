package genetic

import (
	"fmt"
	fp "github.com/rjNemo/underscore"
	"math"
)

// Route represents a tour stating at depot, going through all activities, and ending at the depot.
// (depot -> activities[0] -> activities[1] -> ... -> activities[n] -> depot)
type Route struct {
	Depot      Pos
	Activities []Pos
}

// greedyRoute builds a Route from the depot and a list of points using a greedy algorithm.
// Starting from the depot, it repeatedly selects the next point to visit as the closest point to the last added point.
// If there are no points in the input, the function returns a Route with only the depot.
func greedyRoute(depot Pos, points []Pos) Route {
	route := Route{depot, make([]Pos, len(points))}
	if len(points) == 0 {
		return route
	}
	// Find the closest point to the depot and remove it from the list
	minDist := math.MaxFloat64
	var minI int
	for i, point := range points {
		if d := dist(point, depot); d < minDist {
			minDist = d
			minI = i
		}
	}
	points, route.Activities[0], _ = remove(points, minI)

	// Repeatedly add the closest point to the last added point until all points are visited
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

// applySwap swaps the position of activity i and j in the route.
// Returns an error if i or j is out of range [0, len(route.Activities)), otherwise returns nil.
func (route *Route) applySwap(i, j int) error {
	if i < 0 || i >= len(route.Activities) || j < 0 || j >= len(route.Activities) {
		return fmt.Errorf("index out of bounds: either i=%d or j=%d not in range [0, len(route.Activities)) = [0, %d)",
			i, j, len(route.Activities))
	}
	if len(route.Activities) > 1 && i != j {
		route.Activities[i], route.Activities[j] = route.Activities[j], route.Activities[i]
	}
	return nil
}

// apply2Opt applies the 2-opt move to the slice route.Activities[i:j+1], which reverses the order of the activities
// between indices i and j, inclusively. The method does this in place, so it modifies the route object it is called on.
// Returns an error if i, j < 0 or i, j >= len(route.Activities), indicating that the indices are out of bounds.
// Otherwise, it returns nil to indicate that the method executed successfully.
func (route *Route) apply2Opt(i, j int) error {
	if i < 0 || i >= len(route.Activities) || j < 0 || j >= len(route.Activities) {
		return fmt.Errorf("index out of bounds: either i=%d or j=%d not in range [0,len(route.activities)) = [0,%d)",
			i, j, len(route.Activities))
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
// in the position that minimizes the total distance traveled by the vehicles.
// Returns an error if i<0 or i>=len(route.activities); otherwise returns nil
func (route *Route) applyGreedy(i int) error {
	if i < 0 || i >= len(route.Activities) {
		return fmt.Errorf("index out of bounds: i=%d is not within the valid range [0,len(route.activities)) = [0,%d)",
			i, len(route.Activities))
	}
	if len(route.Activities) <= 1 {
		return nil
	}
	var act Pos
	var err error
	// Get and remove the activity at index i
	route.Activities, act, err = remove(route.Activities, i)
	if err != nil {
		return err
	}

	// Find the best place to insert the activity
	j := 0
	// TODO: cache the last calculated distance as that distance is needed again in the next iteration
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
	// Insert the activity in the best place
	route.Activities, err = insert(route.Activities, act, j)
	if err != nil {
		return err
	}
	return nil
}

// applyInsertGreedy inserts pos greedily in route in the place that minimized the distance added by inserting this Pos.
func (route *Route) applyInsertGreedy(pos Pos) {
	route.Activities = append(route.Activities, pos)
	// error can be ignored because it will never happen.
	// given: len(route.Activities) >= 1
	// 0 <= len(route.Activities)-1 < len(route.Activities) always holds
	_ = route.applyGreedy(len(route.Activities) - 1)
}

// applyChangeDepot changes the Depot of route to the d-th depot in inst.depots
// Returns an error if d<0 or d>=len(inst.depots); otherwise returns nil
func (route *Route) applyChangeDepot(inst MDMTSPInstance, d int) error {
	if d < 0 || d >= len(inst.Depots) {
		return fmt.Errorf("index out of bounds: d=%d is not within the valid range [0,len(inst.depots)) = [0,%d)",
			d, len(inst.Depots))
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
