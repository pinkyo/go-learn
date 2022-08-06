package main

func main() {
}

func taskSchedulerII(tasks []int, space int) int64 {
	var result int64
	preMap := make(map[int]int64)
	for _, task := range tasks {
		if pre, ok := preMap[task]; ok {
			next := pre + int64(space)
			if next > result {
				result = next
			}
		}
		result++
		preMap[task] = result
	}
	return result
}
