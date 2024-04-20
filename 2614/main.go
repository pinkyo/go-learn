package main

func main() {

}

func diagonalPrime(nums [][]int) int {
	isPrime := func(nums int) bool {
		if nums < 2 {
			return false
		}
		for i := 2; i < nums; i++ {
			if nums%i == 0 {
				return false
			}
		}
		return true
	}

	var result int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			if i == j || i == len(nums)-1-j {
				if nums[i][j] > result && isPrime(nums[i][j]) {
					result = nums[i][j]
				}
			}
		}
	}

	return result
}
