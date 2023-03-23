package main

import (
	structures "bzone/backend/cmd/algo/exampleActivities"
	"encoding/json"
	"fmt"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geo"
	"log"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

//https://reader.elsevier.com/reader/sd/pii/S0360835217305053?token=4AA2F5BA18703003EE63F26E53D954A7B1D31F1E63B9A33C86418C6FD25267C6BD0D9C5A282A3DD266542897A56EDC21&originRegion=eu-west-1&originCreation=20230307132719

//Decision variables
//yik = 1 if customer i is visited in path k, 0 otherwise
//xijk = 1 if, in path k, a visit to customer i is followed by a visit to customer j, 0 otherwise
//sik = starting time of the service to customer i in path k
//Mk = 1 if path k is used, 0 otherwise

//Parameters
//n Number of customers
//gamma Constant cost of a path
//p Maximum number of paths (= number of zones)
//vi If 1, node i must be visited (1 for all nodes in our model)
//si Score of customer i
//tmax Time budget of a path
//tij Time distance from customer i to j
//ti Service time of customer i
//oi Opening time of time window for customer i
//ci Closing time of time window for customer i
//L A large constant

// Solution.solution creates a list of activities, where the order indicates a route that is
// terminated by an activity with activity.id 0 - see paper for specifications of solution model
// Also stores the total cost of the solution and the variance between the highest and lowest cost
type solutionModel struct {
	id       int
	solution []structures.ActivityModel
	cost     float64
	costDiff float64
}

// Finds the lowest float in an array of floats
func MinIntSlice(v []float64) (m float64) {
	for i := 0; i < len(v); i++ {
		if i == 0 || v[i] < m {
			m = v[i]
		}
	}
	return
}

// Finds the largest float in an array of floats
func MaxIntSlice(v []float64) (m float64) {
	for i := 0; i < len(v); i++ {
		if i == 0 || v[i] > m {
			m = v[i]
		}
	}
	return
}

// Function that checks whether a distanceInstance is already in the array of previously calculated distanceInstances
func contains(s []distanceInstance, str distanceInstance) (m bool, e float64) {
	for _, e := range s {
		if e.from_lat == str.from_lat && e.to_lat == str.to_lat && e.from_long == str.from_long && e.to_long == str.to_long {
			return true, e.distance
		}
	}

	return false, -1
}

// Structure used in next function
type distanceInstance struct {
	from_lat  string
	from_long string
	to_lat    string
	to_long   string
	distance  float64
}

// Array containing already calculated distances between different points
var distanceMatrix []distanceInstance

// Calculates the distance between two points given by string longitudes and latitudes
func geoDistance(a, b, c, d string) float64 {
	// Ugly but can't find a better way to convert the string data storage type to something usable for geo.Distance
	x1, _ := strconv.ParseFloat(strings.Replace(a, ",", ".", 1), 64)
	x2, _ := strconv.ParseFloat(strings.Replace(b, ",", ".", 1), 64)
	y1, _ := strconv.ParseFloat(strings.Replace(c, ",", ".", 1), 64)
	y2, _ := strconv.ParseFloat(strings.Replace(d, ",", ".", 1), 64)

	return geo.Distance(orb.Point{x1, y1}, orb.Point{x2, y2})
}

// Calculates for a given set of activities a set of zone ranges that contain all values within the zone
func calculateZoneRanges(currentZone []structures.ActivityModel, zoneId int) (completeZone structures.ZoneModel) {
	var zipcodes []string

	// List all zipcodes of the set of activities
	for i := range currentZone {
		zipcodes = append(zipcodes, currentZone[i].Address.Zipcode)
	}

	// Sort the zipcodes in ascending order
	sort.Strings(zipcodes)

	var k int
	var minzip int
	var e error

	// The lowest zipcode is the lowest value after sorting
	if len(zipcodes) > 0 {
		minzip, e = strconv.Atoi(zipcodes[0])
		if e != nil {
			log.Fatal(e)
		}
	}

	for i := 0; i < len(zipcodes)-1; i++ {
		// Store the current zipcode
		currentZipcode, e := strconv.Atoi(zipcodes[i])
		if e != nil {
			log.Fatal(e)
		}

		// Store the next zipcode
		nextZipcode, e := strconv.Atoi(zipcodes[i+1])
		if e != nil {
			log.Fatal(e)
		}

		// If the next zipcode is not directly after the current zipcode, a new zone range has to be created
		if (!(currentZipcode+1 == nextZipcode) && !(currentZipcode == nextZipcode)) || i == len(zipcodes)-2 {
			completeZone.Id = zoneId

			completeZone.Zone_ranges = append(completeZone.Zone_ranges, structures.ZoneRangeModel{
				Id:           k,
				Zipcode_from: minzip,
				Zipcode_to:   currentZipcode,
			})

			minzip = nextZipcode
			k++
		}
	}

	return
}

// Calculates the route cost of a set of activities
func calculateRouteCost(currentRoute []structures.ActivityModel) (routeCost float64) {
	start := time.Now()
	for i := range currentRoute {
		// Continue until len(currentRoute)-1 to prevent indexOutOfBounds
		if i != len(currentRoute)-1 {
			// Check if the value that has to be calculated has already been calculated
			alreadyCalculated, calculatedDistance := contains(distanceMatrix, distanceInstance{currentRoute[i].Latitude, currentRoute[i].Latitude, currentRoute[i+1].Longitude, currentRoute[i+1].Longitude, -1})
			if alreadyCalculated {
				routeCost = calculatedDistance
			} else {
				// If the value has not been calculated yet, calculate it and add it to distanceMatrix
				routeCost += geoDistance(currentRoute[i].Latitude, currentRoute[i+1].Latitude, currentRoute[i].Longitude, currentRoute[i+1].Longitude)
				distanceMatrix = append(distanceMatrix, distanceInstance{currentRoute[i].Latitude, currentRoute[i].Latitude, currentRoute[i+1].Longitude, currentRoute[i+1].Longitude, routeCost})
			}
		} else if i == len(currentRoute)-1 && len(currentRoute) > 1 {
			// Calculate for the last distance the distance back to the first activity
			alreadyCalculated, calculatedDistance := contains(distanceMatrix, distanceInstance{currentRoute[i].Latitude, currentRoute[i].Latitude, currentRoute[0].Longitude, currentRoute[0].Longitude, -1})
			if alreadyCalculated {
				routeCost = calculatedDistance
			} else {
				routeCost += geoDistance(currentRoute[i].Latitude, currentRoute[0].Latitude, currentRoute[i].Longitude, currentRoute[0].Longitude)
				distanceMatrix = append(distanceMatrix, distanceInstance{currentRoute[i].Latitude, currentRoute[i].Latitude, currentRoute[0].Longitude, currentRoute[0].Longitude, routeCost})
			}
		}
	}

	elapsed := time.Since(start)
	log.Printf("Calculating total route cost took %s and distanceMatrix contains %d elements", elapsed, len(distanceMatrix))

	return
}

// Calculate the cost of a complete solution
func calculateSolutionCost(solution solutionModel, n int, p int, change string, j int, k int) (totalCost float64, costDiff float64) {
	var currentRoute []structures.ActivityModel
	var routeCostArray []float64

	if change == "reversion" || (change == "swap" && (solution.solution[j].Id == 0 || solution.solution[k].Id == 0)) {
		// For every value in the solution
		for i := 0; i < n+p-1; i++ {
			// Check if the route is terminated by checking if the solution[i].id is zero, continue otherwise
			if solution.solution[i].Id == 0 || i == len(solution.solution)-1 {
				routeCost := calculateRouteCost(currentRoute)
				routeCostArray = append(routeCostArray, routeCost)
				totalCost += routeCost
				currentRoute = nil
			} else {
				currentRoute = append(currentRoute, solution.solution[i])
			}
		}
	} else if change == "swap" {
		swapValues := []int{k - 1, k + 1, j - 1, j + 1}
		swapValuesBase := []int{k, k, j, j}

		for i, e := range swapValues {
			_, calculatedDistance := contains(distanceMatrix, distanceInstance{solution.solution[e].Latitude, solution.solution[e].Latitude, solution.solution[swapValuesBase[i]].Longitude, solution.solution[swapValuesBase[i]].Longitude, -1})
			totalCost -= calculatedDistance

			i++
		}

		for i, e := range swapValues {
			alreadyCalculated, calculatedDistance := contains(distanceMatrix, distanceInstance{solution.solution[e].Latitude, solution.solution[e].Latitude, solution.solution[swapValuesBase[i]].Longitude, solution.solution[swapValuesBase[i]].Longitude, -1})
			if alreadyCalculated {
				totalCost += calculatedDistance
			} else {
				routeCost := geoDistance(solution.solution[e].Latitude, solution.solution[swapValuesBase[i]].Latitude, solution.solution[e].Longitude, solution.solution[swapValuesBase[i]].Longitude)
				totalCost += routeCost
				distanceMatrix = append(distanceMatrix, distanceInstance{solution.solution[i].Latitude, solution.solution[0].Latitude, solution.solution[i].Longitude, solution.solution[0].Longitude, routeCost})
			}

			i++
		}

	} else {
		fmt.Printf("Error: the function %s does not exist", change)
	}

	// Calculate the variance within the cost array
	costDiff = MaxIntSlice(routeCostArray) - MinIntSlice(routeCostArray)
	return
}

// Initializes a solution with a random order of activities
func randomSolutions(activityArray []structures.ActivityModel, p int) (initialSolution []structures.ActivityModel) {
	//Create set of exampleActivities for random permutation
	//Set of exampleActivities
	for i := range activityArray {
		initialSolution = append(initialSolution, activityArray[i])
	}
	//Zeros to indicate routes
	for i := 0; i < p-1; i++ {
		initialSolution = append(initialSolution, structures.ActivityModel{Id: 0})
	}

	//Create random permutation
	for i := range initialSolution {
		j := rand.Intn(i + 1)
		initialSolution[i], initialSolution[j] = initialSolution[j], initialSolution[i]
	}

	return
}

// Swaps two values in the solution
func swap(solution solutionModel) (newSolution solutionModel, a int, b int) {
	a = rand.Intn(len(solution.solution))
	b = rand.Intn(len(solution.solution))

	solution.solution[a], solution.solution[b] = solution.solution[b], solution.solution[a]

	return solution, a, b
}

// Reverses the whole solution
func reversion(solution solutionModel) (newSolution solutionModel) {
	start := 0
	end := len(solution.solution) - 1

	for start < end {
		solution.solution[start], solution.solution[end] = solution.solution[end], solution.solution[start]
		start++
		end--
	}

	return solution
}

// Structure that is used to run the concurrency, stores a global best solution
type sigmaBestMu struct {
	mutex     sync.Mutex
	sigmaBest solutionModel
}

// Function on sigmaBestMu that does the brunt of the calculating work, is called concurrently
func (mu *sigmaBestMu) step3456(solutionArray []solutionModel, activityArray []structures.ActivityModel, foundBestSol bool, T float64, Iiter int, N int, nNonImproving int, alpha float64, n int, p int, R int, k int) {
	//Step 3. N = N + 1
	N++
	//No best solution has been found this run so far
	foundBestSol = false

	fmt.Printf("(%d/%d) Starting new iteration/Current sigmaBest.cost = %.2f with costRange %.2f\n", N, Iiter, mu.sigmaBest.cost, mu.sigmaBest.costDiff)

	//Step 4. For k=1 to p
	sigmaNew := solutionModel{cost: math.MaxInt}
	//Step 4.1 Generating a solution sigmaNew based on sigmaK
	//Generate s~U(0,1)
	s := rand.Float64()

	var i = 0
	var j = 0

	//If (s <= 2/3) do swap, else do reversion
	if s <= 1.0/3.0 {
		sigmaNew, i, j = swap(solutionArray[k])
		fmt.Printf("		(%d/%d) Swapped values in solutionArray %d\n", N, Iiter, k)
	} else if s > 1.0/3.0 && s <= 2.0/3.0 {
		sigmaNew, i, j = swap(solutionArray[k])
		fmt.Printf("		(%d/%d) Swapped values in solutionArray %d\n", N, Iiter, k)
	} else {
		sigmaNew = reversion(solutionArray[k])
		fmt.Printf("		(%d/%d) Reversed values in solutionArray %d\n", N, Iiter, k)
	}

	//Step 4.2 If deltaK >= 0 {let solutionArray[k] be sigmaNew} else {random chance that it replaces solutionArray[k] regardless}
	if i > 0 && j > 0 {
		sigmaNew.cost, sigmaNew.costDiff = calculateSolutionCost(sigmaNew, n, p, "swap", i, j)
	} else {
		sigmaNew.cost, sigmaNew.costDiff = calculateSolutionCost(sigmaNew, n, p, "reversion", 0, 0)

	}

	deltaK := sigmaNew.cost - solutionArray[k].cost

	if deltaK >= 0 {
		solutionArray[k] = sigmaNew
	} else {
		//Generate r~U(0,1)
		r := rand.Float64()
		if r < math.Pow(math.E, deltaK/T) {
			solutionArray[k] = sigmaNew
		}
	}

	//Here we evaluate whether the global sigmaBest has to be updated such that a better solution has been found
	mu.mutex.Lock()
	//Step 4.3 If solutionArray[k].cost > sigmaBest.cost we have found a new best solution
	if solutionArray[k].cost > mu.sigmaBest.cost && solutionArray[k].costDiff <= mu.sigmaBest.costDiff {
		mu.sigmaBest = solutionArray[k]
		R = 0
		foundBestSol = true
	}
	mu.mutex.Unlock()

	fmt.Printf("		(%d/%d) Evaluated new solution\n", N, Iiter)

	//Step 5
	if N == Iiter {
		T = T * alpha
		N = 0
		if foundBestSol == false {
			R = R + 1
			foundBestSol = true
		}
	}

	fmt.Printf("		(%d/%d) Ending iteration, improved best solution last %d iterations ago\n", N, Iiter, R)
	//Step 6
	if R == nNonImproving {
		fmt.Printf("(*) Goroutine %d has found an optimal solution with cost %f and costDiff %f\n", mu.sigmaBest.id, mu.sigmaBest.cost, mu.sigmaBest.costDiff)

		var zonesArray []structures.ZoneModel
		var currentZone []structures.ActivityModel

		var j int

		//For every location in a route in the final solution, add the zipcode associated to that location to a zone
		for i := range mu.sigmaBest.solution {
			if mu.sigmaBest.solution[i].Id == 0 || i == len(mu.sigmaBest.solution)-1 {
				zonesArray = append(zonesArray, calculateZoneRanges(currentZone, j))
				j++
			} else {
				currentZone = append(currentZone, mu.sigmaBest.solution[i])
			}
		}

		fmt.Println("(*) Successfully exported optimal solution to zonemodel")
		os.Exit(3)

	} else {
		go mu.step3456(solutionArray, activityArray, foundBestSol, T, Iiter, N, nNonImproving, alpha, n, p, R, k)
	}
}

func main() {
	fmt.Println("(*) Starting with initialization now")

	// read our opened jsonFile as a byte array.
	data, err := os.ReadFile("./exampleActivities/csvjson.json")
	if err != nil {
		fmt.Print(err)
	}

	var activityArray []structures.ActivityModel

	// Unpack the jsonFile to an activityArray model
	err = json.Unmarshal(data, &activityArray)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("	(**) File has been loaded...")

	nrOfActivities := len(activityArray)
	n := nrOfActivities
	p := 1

	var solutionArray []solutionModel

	for i := 0; i < p; i++ {
		solutionArray = append(solutionArray, solutionModel{id: i})
	}

	//Step 1a. Generating the initial solutions
	for i := range solutionArray {
		solutionArray[i].solution = randomSolutions(activityArray, p)
	}

	fmt.Println("	(**) Random solutions have been created...")

	//Step 1b. Calculate costs for each of the initial solutions
	for i := 0; i < p; i++ {
		solutionArray[i].cost, solutionArray[i].costDiff = calculateSolutionCost(solutionArray[i], n, p, "reversion", 0, 0)
	}

	fmt.Println("	(**) Costs for solutions have been calculated...")

	//Step 2. Let T = t0 ; N=0; sigmabest = the best solution/sigmak among the p solutions;
	//fbest = sigmabest.cost; R=0; foundBestSol=false;
	T := 3.0
	N := 0

	mu := sigmaBestMu{sigmaBest: solutionModel{cost: math.MaxInt}}

	for i := range solutionArray {
		if solutionArray[i].cost < mu.sigmaBest.cost {
			mu.sigmaBest = solutionArray[i]
		}
	}

	R := 0
	foundBestSol := false

	fmt.Println("	(**) Variables have been set...")

	fmt.Printf("(*) Initialization complete, running %d concurrent iterations now\n\n", p)

	//Step 3, 4, 5, 6. More details in the function above
	for k := 0; k < p; k++ {
		go mu.step3456(solutionArray, activityArray, foundBestSol, T, 1000, N, 15, 0.9, n, p, R, k)

		var input string
		fmt.Scanln(&input)
	}
}
