package main

func main() {
}

func minimumIndex(nums []int) int {
	cntMap := make(map[int]int)
	for _, num := range nums {
		cntMap[num]++
	}

	cnt2Map := make(map[int]int)
	maxVal := -1
	for idx, num := range nums {
		cnt2Map[num]++
		if cnt2Map[num] > cnt2Map[maxVal] {
			maxVal = num
		}
		if cnt2Map[maxVal]*2 > idx+1 && (cntMap[maxVal]-cnt2Map[maxVal])*2 > len(nums)-idx-1 {
			return idx
		}
	}

	return -1
}
