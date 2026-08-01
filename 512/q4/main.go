package main

import "fmt"

func main() {
	fmt.Println(minCost(2, 2, [][]int{
		{0, 7},
		{3, 2},
	}))
}

func minCost(m int, n int, penalty [][]int) int64 {
	oddCost := make([][]int, m)
	evenCost := make([][]int, m)
	for i := 0; i < m; i++ {
		oddCost[i] = make([]int, n)
		evenCost[i] = make([]int, n)
	}
	evenCost[0][0] = 1
	queue := make([][]int, 0)
	queue = append(queue, []int{0, 0, 0})

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		x, y, steps := cur[0], cur[1], cur[2]
		source := evenCost
		target := oddCost
		if steps%2 != 0 {
			source = oddCost
			target = evenCost
		}
		stop := source[x][y] + pen(penalty, x, y)
		if stop < target[x][y] || target[x][y] == 0 {
			target[x][y] = stop
			queue = append(queue, []int{x, y, steps + 1})
		}

		cost1 := source[x][y] + x*(y+1)
		cost2 := source[x][y] + (x+1)*y
		cost3 := source[x][y] + (x+1)*(y+2)
		cost4 := source[x][y] + (x+2)*(y+1)
		px := pen(penalty, x, y)
		if (steps+1)%2 == 0 {
			cost3 += px
			cost4 += px
		} else {
			cost1 += px
			cost2 += px
		}
		if x-1 >= 0 && (cost1 < target[x-1][y] || target[x-1][y] == 0) {
			target[x-1][y] = cost1
			queue = append(queue, []int{x - 1, y, steps + 1})
		}
		if y-1 >= 0 && (cost2 < target[x][y-1] || target[x][y-1] == 0) {
			target[x][y-1] = cost2
			queue = append(queue, []int{x, y - 1, steps + 1})
		}
		if y+1 < n && (cost3 < target[x][y+1] || target[x][y+1] == 0) {
			target[x][y+1] = cost3
			queue = append(queue, []int{x, y + 1, steps + 1})
		}
		if x+1 < m && (cost4 < target[x+1][y] || target[x+1][y] == 0) {
			target[x+1][y] = cost4
			queue = append(queue, []int{x + 1, y, steps + 1})
		}

	}

	return int64(minNonZero(oddCost[m-1][n-1], evenCost[m-1][n-1]))
}

// pen 安全地读取 penalty[x][y]；当下标越界（如测试数据维度不匹配）时按 0 罚分处理
func pen(p [][]int, x, y int) int {
	if x >= 0 && x < len(p) && y >= 0 && y < len(p[x]) {
		return p[x][y]
	}
	return 0
}

func minNonZero(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	if a < b {
		return a
	}
	return b
}
