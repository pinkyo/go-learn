package main

func main() {

}

func numIdenticalPairs(nums []int) int {
	cntMap := make(map[int]int)
	var result int
	for _, num := range nums {
		result += cntMap[num]
		cntMap[num]++
	}
	return result
}
