package main

func main() {
}

func minMaxGame(nums []int) int {
	for len(nums) > 1 {
		next := make([]int, len(nums)/2)
		for i := 0; i < len(nums); i += 2 {
			f := minInt
			if i/2%2 != 0 {
				f = maxInt
			}
			next[i/2] = f(nums[i], nums[i+1])
		}
		nums = next
	}
	return nums[0]
}

func maxInt(v1, v2 int) int {
	if v1 < v2 {
		return v2
	}
	return v1
}

func minInt(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}
