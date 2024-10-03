package main

import "fmt"

// This function is the main function of the program
func main() {
	// This line prints the result of the shortestDistanceAfterQueries function with the parameters 5 and [][]int{{2, 4}, {0, 2}, {0, 4}}
	fmt.Println(shortestDistanceAfterQueries(5, [][]int{{2, 4}, {0, 2}, {0, 4}}))
}

// This function takes in an integer n and a 2D slice of integers queries and returns a slice of integers
func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	// Create a slice of integers dis with length n and initialize each element to its index
	dis := make([]int, n)
	for i := range dis {
		dis[i] = i
	}
	// Create a 2D slice of integers adj with length n and initialize each element to an empty slice
	adj := make([][]int, n)
	for i := 0; i < n-1; i++ {
		adj[i] = make([]int, 0)
		// Append the next index to the current index
		adj[i] = append(adj[i], i+1)
	}

	// Create a slice of integers result with length equal to the length of queries
	result := make([]int, len(queries))
	// Iterate through the queries
	for idx, q := range queries {
		// Append the second element of the current query to the first element of the current query in adj
		adj[q[0]] = append(adj[q[0]], q[1])
		// If the distance to the first element of the current query plus 1 is greater than or equal to the distance to the second element of the current query, continue to the next query
		if dis[q[0]]+1 >= dis[q[1]] {
			continue
		}
		// Set the distance to the second element of the current query to the distance to the first element of the current query plus 1
		dis[q[1]] = dis[q[0]] + 1
		// Iterate through the indices from the second element of the current query to the end of the slice
		for i := q[1]; i < n; i++ {
			// Iterate through the adjacent indices of the current index
			for _, p := range adj[i] {
				// Set the distance to the current adjacent index to the minimum of the current distance and the distance to the current index plus 1
				dis[p] = min(dis[p], dis[i]+1)
			}
		}
		// Set the current index of result to the distance to the last index
		result[idx] = dis[n-1]
	}
	// Return the result
	return result
}

// This function takes two integers as parameters and returns the smaller of the two
func min(a, b int) int {
	// If a is less than b, return a
	if a < b {
		return a
	}
	// Otherwise, return b
	return b
}
