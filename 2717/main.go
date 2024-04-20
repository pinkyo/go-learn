package main

func main() {

}

func semiOrderedPermutation(nums []int) int {
	var pos1 int
	var posn int
	for i, num := range nums {
		if num == 1 {
			pos1 = i
		}
		if num == len(nums) {
			posn = i
		}
	}
	if pos1 < posn {
		return pos1 + len(nums) - 1 - posn
	}
	return pos1 + len(nums) - 1 - posn - 1
}
