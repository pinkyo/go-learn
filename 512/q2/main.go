package main

func main() {
}

func aggregateTimeSeries(series1 [][]int, series2 [][]int) [][]int {
	result := make([][]int, 0, len(series1)+len(series2))
	si, sj := 0, 0
	for si < len(series1) && sj < len(series2) {
		if series1[si][0] < series2[sj][0] {
			result = append(result, []int{series1[si][0], series1[si][1] + series2[sj][1]})
			si++
			continue
		} else if series1[si][0] > series2[sj][0] {
			result = append(result, []int{series2[sj][0], series1[si][1] + series2[sj][1]})
			sj++
			continue
		}
		result = append(result, []int{series1[si][0], series1[si][1] + series2[sj][1]})
		si++
		sj++
	}
	for si < len(series1) {
		result = append(result, []int{series1[si][0], series1[si][1]})
		si++
	}
	for sj < len(series2) {
		result = append(result, []int{series2[sj][0], series2[sj][1]})
		sj++
	}
	return result
}
