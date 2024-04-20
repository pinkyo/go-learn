package main

import "fmt"

func main() {
	fmt.Println(storeWater([]int{9, 0, 1}, []int{0, 2, 2}))
}

func storeWater(bucket []int, vat []int) int {
	var max int
	for i := 0; i < len(vat); i++ {
		if bucket[i] == 0 {
			if vat[i] > max {
				max = vat[i]
			}
		} else if vat[i]/bucket[i] > max {
			max = (vat[i] + bucket[i] - 1) / bucket[i]
		}
	}
	if max == 0 {
		return 0
	}

	ans := max + len(bucket)
	for i := 1; i <= max; i++ {
		var sum int
		for j := 0; j < len(bucket); j++ {
			temp := (vat[j]+i-1)/i - bucket[j]
			if temp > 0 {
				sum += temp
			}
		}
		sum += i
		if sum < ans {
			ans = sum
		}
	}
	return ans
}
