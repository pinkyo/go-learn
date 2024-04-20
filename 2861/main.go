package main

func main() {
}

func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {
	var result int
	for i := 0; i < k; i++ {
		val := calculateAlloysOfMachine(n, budget, composition[i], stock, cost)
		if val > result {
			result = val
		}
	}
	return result
}

func calculateAlloysOfMachine(n int, budget int, composition []int, stock []int, cost []int) int {
	l := 0
	h := (budget/cost[0] + stock[0]) / composition[0]
	for l <= h {
		mid := (l + h) / 2
		sum := 0
		for i, item := range composition {
			if item*mid > stock[i] {
				sum += (item*mid - stock[i]) * cost[i]
			}
		}
		if sum > budget {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	return h
}
