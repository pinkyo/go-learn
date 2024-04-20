package main

func main() {
}

func alternatingSubarray(nums []int) int {
	result := -1
	current := -1
	for i := 1; i < len(nums); i++ {
		if current == -1 && nums[i]-nums[i-1] == 1 {
			current = 2
			continue
		}
		if current > 0 && nums[i] == nums[i-2] {
			current++
			continue
		}
		if current > result {
			result = current
		}
		current = -1
		if current == -1 && nums[i]-nums[i-1] == 1 {
			current = 2
		}
	}
	if current > result {
		result = current
	}
	return result
}

// func abs(a int) int {
// 	if a < 0 {
// 		return -a
// 	}
// 	return a
// }
