package main

func main() {
}

// This function takes in two slices of integers, energyDrinkA and energyDrinkB, and returns an int64 value representing the maximum energy boost that can be achieved.
func maxEnergyBoost(energyDrinkA []int, energyDrinkB []int) int64 {
	// Create a 2D slice of int64 values to store the maximum energy boost for each drink
	dp := make([][]int64, len(energyDrinkA)+1)
	// Initialize the first row of the slice to 0
	dp[0] = []int64{0, 0}
	// Iterate through the slices of energy drinks
	for i := 1; i < len(dp); i++ {
		// Initialize the current row of the slice
		dp[i] = make([]int64, 2)
		// Update the maximum energy boost for the current drink
		dp[i][0] = max(dp[i][0], dp[i-1][0]+int64(energyDrinkA[i-1]))
		dp[i][1] = max(dp[i][1], dp[i-1][1]+int64(energyDrinkB[i-1]))
		// If the current index is greater than 1, update the maximum energy boost for the previous drink
		if i > 1 {
			dp[i][0] = max(dp[i][0], dp[i-2][1]+int64(energyDrinkA[i-1]))
			dp[i][1] = max(dp[i][1], dp[i-2][0]+int64(energyDrinkB[i-1]))
		}
	}

	// Return the maximum energy boost from the last row of the slice
	return max(dp[len(dp)-1][0], dp[len(dp)-1][1])
}

// This function takes in a variable number of int64 arguments and returns the maximum value
func max(arr ...int64) int64 {
	// Initialize the maximum value to the first element in the array
	max := arr[0]
	// Loop through the array
	for _, v := range arr {
		// If the current value is greater than the maximum value, update the maximum value
		if v > max {
			max = v
		}
	}
	// Return the maximum value
	return max
}
