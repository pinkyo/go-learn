package main

func main() {

}

func rowAndMaximumOnes(mat [][]int) []int {
	result := make([]int, 2)
	for i, row := range mat {
		var cnt int
		for _, col := range row {
			if col == 1 {
				cnt += 1
			}
		}
		if cnt > result[1] {
			result[1] = cnt
			result[0] = i
		}
	}
	return result
}
