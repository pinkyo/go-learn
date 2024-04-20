package main

func main() {

}

func arithmeticTriplets(nums []int, diff int) int {
	idxMap := make(map[int][]int)
	for i, num := range nums {
		idxMap[num] = append(idxMap[num], i)
	}
	var result int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j]-nums[i] != diff {
				continue
			}
			for _, k := range idxMap[nums[j]+diff] {
				if k > j {
					result++
				}
			}
		}
	}

	return result
}
