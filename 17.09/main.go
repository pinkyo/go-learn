package main

import "container/heap"

func main() {
}

func getKthMagicNumber(k int) int {
	factors := []int{3, 5, 7}
	h := &hp{1}
	var result int
	for i := 0; i < k; i++ {
		x := heap.Pop(h).(int)
		for h.Len() > 0 && x == (*h)[0] {
			x = heap.Pop(h).(int)
		}
		for _, factor := range factors {
			heap.Push(h, x*factor)
		}
		result = x
	}
	return result
}

type hp []int

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *hp) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
