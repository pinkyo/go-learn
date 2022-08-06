package main

func main() {

}

func countPairs(n int, edges [][]int) int64 {
	preArr := make([]int, n)
	for i := 0; i < n; i++ {
		preArr[i] = i
	}
	for _, v := range edges {
		pre1 := v[0]
		for preArr[pre1] != pre1 {
			pre1 = preArr[pre1]
		}
		pre2 := v[1]
		for preArr[pre2] != pre2 {
			pre2 = preArr[pre2]
		}
		preArr[pre1] = min(pre1, pre2)
		preArr[pre2] = min(pre1, pre2)
	}
	preMap := make(map[int]int)
	for i := 0; i < n; i++ {
		val := preArr[i]
		for preArr[val] != val {
			val = preArr[val]
		}
		preMap[val]++
	}
	var result int
	for _, v := range preMap {
		result += v * (n - v)
	}
	return int64(result / 2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
