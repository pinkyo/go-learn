package main

import "math"

func main() {

}

func average(salary []int) float64 {
	max := math.MinInt
	min := math.MaxInt
	var sum int
	for _, s := range salary {
		sum += s
		if s > max {
			max = s
		}
		if s < min {
			min = s
		}
	}
	sum -= max
	sum -= min
	return float64(sum) / float64(len(salary)-2)
}
