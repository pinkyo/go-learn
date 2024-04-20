package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findPrimePairs(1000000))
}

func findPrimePairs(n int) [][]int {
	if n%2 == 1 {
		if isPrime(n - 2) {
			return [][]int{{2, n - 2}}
		}
		return nil
	}

	primes := make([]int, 0)
OUTER:
	for i := 2; i <= n; i++ {
		limit := math.Sqrt(float64(i))
		for _, val := range primes {
			if i%val == 0 {
				continue OUTER
			}
			if val > int(limit) {
				break
			}
		}
		primes = append(primes, i)
	}
	var result [][]int
	i, j := 0, len(primes)-1
	for i <= j {
		if primes[i]+primes[j] == n {
			result = append(result, []int{primes[i], primes[j]})
			i++
		} else if primes[i]+primes[j] < n {
			i++
		} else {
			j--
		}
	}
	return result
}

func isPrime(val int) bool {
	if val < 2 {
		return false
	}
	limit := math.Sqrt(float64(val))
	for i := 2; i <= int(limit); i++ {
		if val%i == 0 {
			return false
		}
	}
	return true
}

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
