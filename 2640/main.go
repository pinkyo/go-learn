package main

func main() {

}

func findPrefixScore(nums []int) []int64 {
	var result []int64
	var sum int
	var max int
	for _, num := range nums {
		if num > max {
			max = num
		}
		sum += (max + num)
		result = append(result, int64(sum))
	}

	return result
}
