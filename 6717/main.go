package main

func main() {
}

func findSubarrays(nums []int) bool {
	exist := make(map[int]struct{})
	for i := 0; i < len(nums)-1; i++ {
		val := nums[i] + nums[i+1]
		if _, ok := exist[val]; ok {
			return true
		}
		exist[val] = struct{}{}
	}
	return false
}
