package main

func main() {
}

func hardestWorker(n int, logs [][]int) int {
	max := 0
	var result int
	pre := 0
	for i := 0; i < len(logs); i++ {
		t := logs[i][1] - pre
		if t > max {
			max = t
			result = logs[i][0]
		} else if t == max && logs[i][0] < result {
			result = logs[i][0]
		}
		pre = logs[i][1]
	}
	return result
}
