package main

import "fmt"

func main() {
	fmt.Println(maxScore([]int{1, 2, 3, 4, 5, 6, 1}, 3))
}

func maxScore(cardPoints []int, k int) int {
	sumArr := make([]int, len(cardPoints))
	sumArr[0] = cardPoints[0]
	for i := 1; i < len(cardPoints); i++ {
		sumArr[i] = sumArr[i-1] + cardPoints[i]
	}

	result := sumArr[k-1]
	var sum int
	for i := len(cardPoints) - 1; i >= len(cardPoints)+1-k; i-- {
		sum += cardPoints[i]
		tmp := sum + sumArr[k+i-len(cardPoints)-1]
		if tmp > result {
			result = tmp
		}
	}
	sum += cardPoints[len(cardPoints)-k]
	if sum > result {
		result = sum
	}
	return result
}
