package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(kthSmallestPrimeFraction([]int{1, 3, 11, 13, 37, 53, 59}, 6))
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	priorityQueue := make(priorityQueue, 0, len(arr))
	heap.Init(&priorityQueue)

	var cnt int
	var result pair
	// OUTER:
	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; j > i; j-- {
			p := pair{arr[i], arr[j]}
			// for len(priorityQueue) > 0 && priorityQueue[0][0]*p[1] < p[0]*priorityQueue[0][1] {
			// 	val := heap.Pop(&priorityQueue)
			// 	cnt++
			// 	if cnt == k {
			// 		result, _ = val.(pair)
			// 		break OUTER
			// 	}
			// }
			heap.Push(&priorityQueue, p)
		}
	}
	for len(priorityQueue) > 0 && cnt < k {
		val := heap.Pop(&priorityQueue)
		result, _ = val.(pair)
		cnt++
	}
	return result[:]
}

type pair [2]int

// priorityQueue implements heap.Interface and holds pairs.
type priorityQueue []pair

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i][0]*pq[j][1] < pq[j][0]*pq[i][1]
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(pair))
}

func (pq *priorityQueue) Pop() interface{} {
	n := len(*pq)
	x := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return x
}
