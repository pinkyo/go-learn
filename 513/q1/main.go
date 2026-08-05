package main

func main() {
}

func maxPairStrength(nums []int) int64 {
	result := int64(0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			g := gcd(int64(nums[i]), int64(nums[j]))
			result = max(result, int64(nums[i])*int64(nums[j])/g/g)
		}
	}
	return result
}

func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
