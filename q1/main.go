package main

func main() {
}

type neighborSum struct {
	grid     [][]int
	position map[int][2]int
}

func Constructor(grid [][]int) neighborSum {
	position := make(map[int][2]int)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			position[grid[i][j]] = [2]int{i, j}
		}
	}
	return neighborSum{
		grid:     grid,
		position: position,
	}
}

func (this *neighborSum) AdjacentSum(value int) int {
	pos := this.position[value]
	result := 0
	if pos[0] > 0 {
		result += this.grid[pos[0]-1][pos[1]]
	}
	if pos[0] < len(this.grid)-1 {
		result += this.grid[pos[0]+1][pos[1]]
	}
	if pos[1] > 0 {
		result += this.grid[pos[0]][pos[1]-1]
	}
	if pos[1] < len(this.grid[0])-1 {
		result += this.grid[pos[0]][pos[1]+1]
	}
	return result
}

func (this *neighborSum) DiagonalSum(value int) int {
	pos := this.position[value]
	result := 0
	if pos[0] > 0 && pos[1] > 0 {
		result += this.grid[pos[0]-1][pos[1]-1]
	}
	if pos[0] > 0 && pos[1] < len(this.grid[0])-1 {
		result += this.grid[pos[0]-1][pos[1]+1]
	}
	if pos[0] < len(this.grid)-1 && pos[1] > 0 {
		result += this.grid[pos[0]+1][pos[1]-1]
	}
	if pos[0] < len(this.grid)-1 && pos[1] < len(this.grid[0])-1 {
		result += this.grid[pos[0]+1][pos[1]+1]
	}
	return result
}

/**
 * Your neighborSum object will be instantiated and called as such:
 * obj := Constructor(grid);
 * param_1 := obj.AdjacentSum(value);
 * param_2 := obj.DiagonalSum(value);
 */
