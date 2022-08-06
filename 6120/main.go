package main

func main() {

}
func numberOfPairs(nums []int) []int {
	cntMap := make(map[int]int)
	for _, v := range nums {
		cntMap[v]++
	}
	res := make([]int, 2)
	for _, v := range cntMap {
		res[0] += v / 2
		res[1] += v % 2
	}
	return res
}
