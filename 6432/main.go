package main

import "fmt"

func main() {
	// fmt.Println(countCompleteComponents(6, [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}}))
	fmt.Println(countCompleteComponents(5, [][]int{{2, 1}, {3, 2}, {4, 0}, {4, 2}}))
}

func countCompleteComponents(n int, edges [][]int) int {
	pre := make([]int, n)
	cntMap := make(map[int]int)
	nodeCntMap := make(map[int]int)

	for i := 0; i < n; i++ {
		pre[i] = i
	}

	for i := 0; i < len(edges); i++ {
		edge := edges[i]
		pre1 := edge[0]
		for pre[pre1] != pre1 {
			pre1 = pre[pre1]
		}
		pre2 := edge[1]
		for pre[pre2] != pre2 {
			pre2 = pre[pre2]
		}
		preVal := min(pre1, pre2)
		pre[pre1] = preVal
		pre[pre2] = preVal
		cntMap[edge[0]]++
		cntMap[edge[1]]++
	}

	for i := 0; i < n; i++ {
		val := i
		for pre[val] != val {
			val = pre[val]
		}
		pre[i] = val
		if i != val {
			cntMap[val] += cntMap[i]
		}
		nodeCntMap[val]++
	}

	var result int
	for i := 0; i < n; i++ {
		if pre[i] == i && cntMap[i] == nodeCntMap[i]*(nodeCntMap[i]-1) {
			result++
		}
	}

	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
