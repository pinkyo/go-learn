package main

import "fmt"

func main() {
	tests := [][3]int{
		{2, 3, 2}, {3, 3, 4}, {1, 4, 2}, {1, 1, 1},
		{2, 2, 3}, {2, 5, 4}, {5, 2, 4}, {4, 2, 3},
		{10, 10, 1}, {10, 10, 2}, {10, 10, 3}, {10, 10, 4},
	}
	for _, tc := range tests {
		m, n, k := tc[0], tc[1], tc[2]
		grid := createGrid(m, n, k)
		fmt.Printf("m=%2d n=%2d k=%d -> paths=%d %v\n", m, n, k, countPaths(grid), grid)
	}
}

// 3988. 创建一个恰好有 K 条路径的网格图 I
// 思路：k <= 4，用“小装置 + 单一路径走廊”构造。
// 小装置（左上角的全开小块）提供恰好 k 条路径，走廊无分支，
// 其余格子全部填 '#'。k > C(m+n-2, m-1)（全空网格的路径数）时无解。
func createGrid(m int, n int, k int) []string {
	seravolith := [3]int{m, n, k}
	_ = seravolith

	if k > comb(m+n-2, m-1) {
		return []string{}
	}

	grid := make([][]byte, m)
	for i := range grid {
		grid[i] = make([]byte, n)
		for j := range grid[i] {
			grid[i][j] = '#'
		}
	}
	open := func(i, j int) { grid[i][j] = '.' }
	// corridor 从 (si, sj) 先沿列 sj 向下、再沿最后一行向右到终点，无分支
	corridor := func(si, sj int) {
		for i := si; i < m; i++ {
			open(i, sj)
		}
		for j := sj; j < n; j++ {
			open(m-1, j)
		}
	}

	switch k {
	case 1:
		// 纯 L 形走廊：第 0 行 + 最后一列，唯一路径
		for j := 0; j < n; j++ {
			open(0, j)
		}
		corridor(0, n-1)
	case 2:
		// 2x2 全开块（2 条路径到 (1,1)），再接通走廊
		open(0, 0)
		open(0, 1)
		open(1, 0)
		open(1, 1)
		corridor(1, 1)
	case 3:
		if n >= 3 {
			// 2x3 全开块：C(3,1)=3 条路径到 (1,2)
			for i := 0; i < 2; i++ {
				for j := 0; j < 3; j++ {
					open(i, j)
				}
			}
			corridor(1, 2)
		} else {
			// n == 2 且 m >= 3：3x2 全开块，C(3,1)=3 条路径到 (2,1)
			for i := 0; i < 3; i++ {
				for j := 0; j < 2; j++ {
					open(i, j)
				}
			}
			corridor(2, 1)
		}
	case 4:
		if m >= 3 && n >= 3 {
			// 两个 2x2 全开块串联：2*2=4 条路径到 (2,2)
			open(0, 0)
			open(0, 1)
			open(1, 0)
			open(1, 1)
			open(1, 2)
			open(2, 1)
			open(2, 2)
			corridor(2, 2)
		} else if m == 2 {
			// 此时必有 n >= 4：第 0 行只开放前 4 列，路径总数 = 4
			for j := 0; j < 4; j++ {
				open(0, j)
			}
			for j := 0; j < n; j++ {
				open(1, j)
			}
		} else {
			// n == 2，此时必有 m >= 4：第 0 列只开放前 4 行
			for i := 0; i < 4; i++ {
				open(i, 0)
			}
			for i := 0; i < m; i++ {
				open(i, 1)
			}
		}
	}

	result := make([]string, m)
	for i := range grid {
		result[i] = string(grid[i])
	}
	return result
}

// comb 计算 C(a, b)
func comb(a, b int) int {
	if b > a-b {
		b = a - b
	}
	result := 1
	for i := 0; i < b; i++ {
		result = result * (a - i) / (i + 1)
	}
	return result
}

// countPaths 用 DP 校验网格的路径数，便于自测
func countPaths(grid []string) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '#' {
				continue
			}
			if i == 0 && j == 0 {
				dp[i][j] = 1
				continue
			}
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
