package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	nc := Constructor()
	fmt.Println(nc.Find(10)) // 没有数字 10 ，所以返回 -1 。
	nc.Change(2, 10)         // 容器中下标为 2 处填入数字 10 。
	nc.Change(1, 10)         // 容器中下标为 1 处填入数字 10 。
	nc.Change(3, 10)         // 容器中下标为 3 处填入数字 10 。
	nc.Change(5, 10)         // 容器中下标为 5 处填入数字 10 。
	fmt.Println(nc.Find(10)) // 数字 10 所在的下标为 1 ，2 ，3 和 5 。因为最小下标为 1 ，所以返回 1 。
	nc.Change(1, 20)         // 容器中下标为 1 处填入数字 20 。注意，下标 1 处之前为 10 ，现在被替换为 20 。
	fmt.Println(nc.Find(10)) // 数字 10 所在下标为 2 ，3 和 5 。最小下标为 2 ，所以返回 2 。
}

type hi struct {
	sort.IntSlice
}

func (h *hi) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hi) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

type NumberContainers struct {
	idxMap     map[int]*hi
	valueMap   map[int]int
	removeMaps map[int]map[int]int
}

func Constructor() NumberContainers {
	return NumberContainers{
		idxMap:     make(map[int]*hi),
		valueMap:   make(map[int]int),
		removeMaps: make(map[int]map[int]int),
	}
}

func (this *NumberContainers) Change(index int, number int) {
	if value, ok := this.valueMap[index]; ok {
		if value == number {
			return
		}
		if _, ok := this.removeMaps[value]; !ok {
			this.removeMaps[value] = make(map[int]int)
		}
		removeMap := this.removeMaps[value]
		removeMap[index]++
	}
	this.valueMap[index] = number
	if _, ok := this.idxMap[number]; !ok {
		this.idxMap[number] = &hi{
			IntSlice: make([]int, 0),
		}
	}
	heap.Push(this.idxMap[number], index)
}

func (this *NumberContainers) Find(number int) int {
	if _, ok := this.idxMap[number]; !ok {
		return -1
	}
	removeMap := this.removeMaps[number]

	for this.idxMap[number].Len() > 0 {
		if _, ok := removeMap[this.idxMap[number].IntSlice[0]]; ok && removeMap[this.idxMap[number].IntSlice[0]] > 0 {
			removeMap[this.idxMap[number].IntSlice[0]]--
			heap.Pop(this.idxMap[number])
			continue
		}
		break
	}
	if this.idxMap[number].Len() == 0 {
		return -1
	}

	return this.idxMap[number].IntSlice[0]
}
