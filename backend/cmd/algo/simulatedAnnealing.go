package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
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

type solutionModel struct {
	id       int
	solution []int
	cost     int
	costDiff int
}

func MinIntSlice(v []int) (m int) {
	for i := 0; i < len(v); i++ {
		if i == 0 || v[i] < m {
			m = v[i]
		}
	}
	return
}

func MaxIntSlice(v []int) (m int) {
	for i := 0; i < len(v); i++ {
		if i == 0 || v[i] > m {
			m = v[i]
		}
	}
	return
}

func euclideanDistance(a, b, c, d string) (m int) {
	//Ugly but can't find a way to do this better
	x1, _ := strconv.Atoi(a)
	x2, _ := strconv.Atoi(b)
	y1, _ := strconv.Atoi(c)
	y2, _ := strconv.Atoi(d)

	m = int(math.Sqrt(math.Pow(float64(x1-x2), 2) + math.Pow(float64(y1-y2), 2)))

	return
}

func calculateRouteCost(currentRoute []ActivityModel) (routeCost int) {
	for i := range currentRoute {
		if i != len(currentRoute)-1 {
			routeCost += euclideanDistance(currentRoute[i].address.latitude, currentRoute[i+1].address.latitude, currentRoute[i].address.longitude, currentRoute[i+1].address.longitude)
		} else if i == len(currentRoute)-1 && len(currentRoute) > 1 {
			routeCost += euclideanDistance(currentRoute[i].address.latitude, currentRoute[0].address.latitude, currentRoute[i].address.longitude, currentRoute[0].address.longitude)
		}
	}

	return
}

func calculateSolutionCost(activityArray []ActivityModel, solution solutionModel, n int, p int) (totalCost int, costDiff int) {
	currentRoute := []ActivityModel{}
	routeCostArray := []int{}

	for i := 0; i < n+p-1; i++ {
		if solution.solution[i] == 0 {
			routeCost := calculateRouteCost(currentRoute)
			routeCostArray = append(routeCostArray, routeCost)
			totalCost += routeCost
			currentRoute = nil
		} else {
			for k := range activityArray {
				if activityArray[k].id == solution.solution[i] {
					currentRoute = append(currentRoute, activityArray[k])
				}
			}
			if i == len(solution.solution)-1 {
				routeCost := calculateRouteCost(currentRoute)
				totalCost += routeCost
				currentRoute = nil
				if routeCost > 0 {
					routeCostArray = append(routeCostArray, routeCost)
				}
			}
		}
	}

	costDiff = MaxIntSlice(routeCostArray) - MinIntSlice(routeCostArray)
	return
}

func randomSolutions(activityArray []ActivityModel, p int) (initialSolution []int) {
	//Create set of exampleActivities for random permutation
	//Set of exampleActivities
	for i := range activityArray {
		initialSolution = append(initialSolution, activityArray[i].id)
	}
	//Zeros to indicate routes
	for i := 0; i < p-1; i++ {
		initialSolution = append(initialSolution, 0)
	}

	//Create random permutation
	for i := range initialSolution {
		j := rand.Intn(i + 1)
		initialSolution[i], initialSolution[j] = initialSolution[j], initialSolution[i]
	}

	//Last value is always a zero
	//initialSolution = append(initialSolution, 0)

	return
}

func swap(solution solutionModel) (newSolution solutionModel) {
	a := rand.Intn(len(solution.solution))
	b := rand.Intn(len(solution.solution))

	solution.solution[a], solution.solution[b] = solution.solution[b], solution.solution[a]

	return solution
}

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

func step3456(solutionArray []solutionModel, activityArray []ActivityModel, sigmaBest solutionModel, foundBestSol bool, T float64, Iiter int, N int, nNonImproving int, alpha float64, n int, p int, R int) {
	//Step 3. N = N + 1
	N++
	//No best solution has been found this run
	foundBestSol = false

	//Step 4. For k=1 to p
	sigmaNew := solutionModel{cost: math.MaxInt}
	//Arrays start counting at 0 instead of 1
	for k := 0; k < p-1; k++ {
		//Step 4.1 Generating a solution sigmaNew based on sigmaK
		//Generate s~U(0,1)
		s := rand.Float64()

		//If (s <= 2/3) do swap, else do reversion
		if s <= float64(1)/float64(3) {
			sigmaNew = swap(solutionArray[k])
		} else if s > float64(1)/float64(3) && s <= float64(2)/float64(3) {
			sigmaNew = swap(solutionArray[k])
		} else {
			sigmaNew = reversion(solutionArray[k])
		}

		//Step 4.2 If deltaK >= 0 {let solutionArray[k] be sigmaNew} else {random chance that it replaces solutionArray[k] regardless}
		sigmaNew.cost, sigmaNew.costDiff = calculateSolutionCost(activityArray, sigmaNew, n, p)
		deltaK := sigmaNew.cost - solutionArray[k].cost

		if deltaK >= 0 {
			solutionArray[k] = sigmaNew
		} else {
			//Generate r~U(0,1)
			r := rand.Float64()
			if r < math.Pow(math.E, float64(deltaK)/T) {
				solutionArray[k] = sigmaNew
			}
		}

		//Step 4.3 If solutionArray[k].cost > sigmaBest.cost we have found a new best solution
		if solutionArray[k].cost > sigmaBest.cost && solutionArray[k].costDiff <= sigmaBest.costDiff {
			sigmaBest = solutionArray[k]
			R = 0
			foundBestSol = true
		}
	}

	//Step 5
	if N == Iiter {
		T = T * alpha
		N = 0
		if foundBestSol == false {
			R = R + 1
			foundBestSol = true
		}
	} /*else {
		step3456(solutionArray, activityArray, sigmaBest, foundBestSol, T, Iiter, N, nNonImproving, alpha, n, p, R)
	}*/

	//Step 6
	if R == nNonImproving {
		fmt.Printf("The optimal solution is %d with cost %d and costDiff %d\n", sigmaBest.solution, sigmaBest.cost, sigmaBest.costDiff)
		os.Exit(3)
	} else {
		step3456(solutionArray, activityArray, sigmaBest, foundBestSol, T, Iiter, N, nNonImproving, alpha, n, p, R)
	}
}

func main() {
	n := 6
	p := 3

	solutionArray := []solutionModel{
		{id: 1}, {id: 2}, {id: 3},
	}

	activityArray := []ActivityModel{
		{id: 1, address: AddressModel{longitude: "3", latitude: "5"}},
		{id: 2, address: AddressModel{longitude: "6", latitude: "5"}},
		{id: 3, address: AddressModel{longitude: "3", latitude: "4"}},
		{id: 4, address: AddressModel{longitude: "10", latitude: "10"}},
		{id: 5, address: AddressModel{longitude: "1", latitude: "2"}},
		{id: 6, address: AddressModel{longitude: "33", latitude: "7"}},
	}

	//Step 1a. Generating the initial solutions
	for i := range solutionArray {
		solutionArray[i].solution = randomSolutions(activityArray, p)
	}

	//Step 1b. Calculate costs for each of the initial solutions
	for i := 0; i < p; i++ {
		solutionArray[i].cost, solutionArray[i].costDiff = calculateSolutionCost(activityArray, solutionArray[i], n, p)
	}

	//Step 2. Let T = t0 ; N=0; sigmabest = the best solution/sigmak among the p solutions;
	//fbest = sigmabest.cost; R=0; foundBestSol=false;
	T := float64(3)
	N := 0

	sigmaBest := solutionModel{cost: math.MaxInt}
	for i := range solutionArray {
		if solutionArray[i].cost < sigmaBest.cost {
			sigmaBest = solutionArray[i]
		}
	}

	R := 0
	foundBestSol := false

	//Step 3, 4, 5, 6. More details in the function above
	step3456(solutionArray, activityArray, sigmaBest, foundBestSol, T, 5000, N, 15, 0.9, n, p, R)
}
