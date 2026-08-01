package main

import "sort"

func main() {
}

func findXSum(nums []int, k int, x int) []int {
	result := make([]int, 0)
	for i := 0; i < len(nums)-k+1; i++ {
		result = append(result, calculate(nums[i:i+k], x))
	}

	return result
}

func calculate(nums []int, x int) int {
	cntMap := make(map[int]int)
	for _, num := range nums {
		cntMap[num]++
	}
	arr := make([][2]int, 0)
	for k, v := range cntMap {
		arr = append(arr, [2]int{k, v})
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][1] == arr[j][1] {
			return arr[i][0] > arr[j][0]
		}
		return arr[i][1] > arr[j][1]
	})

	result := 0
	for i := 0; i < x && i < len(arr); i++ {
		result += arr[i][0] * arr[i][1]
	}
	return result
}
