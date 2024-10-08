package main

import (
	"sort"
	"strings"
)

func main() {
}

func digitCount(num string) bool {
	cntMap := make(map[int]int)
	for i := 0; i < len(num); i++ {
		cntMap[int(num[i]-'0')]++
	}
	for i := 0; i < len(num); i++ {
		if cntMap[i] != int(num[i]-'0') {
			return false
		}
	}
	return true
}

func largestWordCount(messages []string, senders []string) string {
	cntMap := make(map[string]int)
	for i := 0; i < len(messages); i++ {
		cntMap[senders[i]] += len(strings.Fields(messages[i]))
	}
	max := 0
	var result string
	for k, v := range cntMap {
		if v > max {
			result = k
			max = v
		} else if v == max && result < k {
			result = k
		}
	}
	return result
}

func maximumImportance(n int, roads [][]int) int64 {
	cntArr := make([]int, n)
	for _, v := range roads {
		cntArr[v[0]]++
		cntArr[v[1]]++
	}
	sort.Ints(cntArr)
	var result int64
	for i, v := range cntArr {
		result += int64(v * (i + 1))
	}
	return result
}

type BookMyShow struct {
	n      int
	m      int
	from   int
	remain int
	state  []int
}

func Constructor(n int, m int) BookMyShow {
	// fmt.Println("------")
	return BookMyShow{n: n, m: m, from: 0, remain: n * m, state: make([]int, n)}
}

func (this *BookMyShow) Gather(k int, maxRow int) []int {
	// fmt.Println(this.state)
	cur := this.from
	for cur <= maxRow {
		if this.state[cur]+k <= this.m {
			pre := this.state[cur]
			this.state[cur] += k
			this.remain -= k
			return []int{cur, pre}
		}
		cur++
	}
	return []int{}
}

func (this *BookMyShow) Scatter(k int, maxRow int) bool {
	// fmt.Println(this.state)
	// Check if there are enough seats available to satisfy the request
	if this.remain < k {
		return false
	}
	// Initialize the remaining seats to be allocated
	remain := k
	// Initialize the current row to be allocated
	cur := this.from
	// Loop through the rows until the maximum row or until all seats have been allocated
	for cur <= maxRow {
		// Check if there are any seats available in the current row
		if this.state[cur] < this.m {
			// Subtract the number of available seats from the remaining seats to be allocated
			remain -= this.m - this.state[cur]
		}
		// Check if all seats have been allocated
		if remain <= 0 {
			break
		}
		// Move to the next row
		cur++
	}
	// Check if there are still seats remaining to be allocated
	if remain > 0 {
		return false
	}

	// Initialize the remaining seats to be allocated
	remain = k
	// Initialize the current row to be allocated
	cur = this.from
	// Loop through the rows until the maximum row or until all seats have been allocated
	for cur <= maxRow {
		// Check if there are any seats available in the current row
		if this.state[cur] < this.m {
			// Store the number of available seats in the current row
			pre := this.state[cur]
			// Allocate the minimum of the available seats and the remaining seats to be allocated
			this.state[cur] += min(this.m-this.state[cur], remain)
			// Subtract the number of allocated seats from the remaining seats to be allocated
			remain -= this.m - pre
		}
		// Check if all seats have been allocated
		if remain <= 0 {
			// Update the starting row for the next allocation
			this.from = cur
			break
		}
		// Move to the next row
		cur++
	}

	// Subtract the number of allocated seats from the total number of remaining seats
	this.remain -= k
	// Return true to indicate that the allocation was successful
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
