package main

func main() {
}

func minCost(colors string, neededTime []int) int {
	result := 0
	var preColor byte
	var preIdx int
	for i := 0; i < len(colors); i++ {
		c := colors[i]
		if preColor != c {
			preColor = c
			preIdx = i
			continue
		}
		if neededTime[i] < neededTime[preIdx] {
			result += neededTime[i]
		} else {
			result += neededTime[preIdx]
			preIdx = i
		}
	}

	return result
}
