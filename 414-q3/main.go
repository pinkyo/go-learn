package main

func main() {
}

// findMaximumScore 计算给定整数切片的最大分数。
// 它通过比较从每个可能的起点开始到切片末尾所能获得的最大分数，最终返回最大的那个分数。
// 参数:
//
//	nums []int - 一个整数切片，代表需要计算最大分数的序列。
//
// 返回值:
//
//	int64 - 表示序列的最大分数。
func findMaximumScore(nums []int) int64 {
	// 初始化一个与输入切片长度相同的整数切片result，用于存储每个位置的最大分数。
	result := make([]int, len(nums))

	// 遍历输入切片中的每个元素，计算并更新result切片。
	for i := 0; i < len(nums); i++ {
		// 从当前位置i开始，向前遍历直到切片末尾，计算可能的最大分数。
		for j := i + 1; j < len(nums); j++ {
			// 计算从当前位置i到j所能获得的分数，并考虑是否更新result[j]。
			nval := result[i] + (j-i)*nums[i]
			if nval > result[j] {
				result[j] = nval
			} else if nums[j] >= nums[i] {
				// 如果遇到的下一个元素大于等于当前元素，停止内层循环，因为后续无法获得更高分数。
				break
			}
		}
	}
	// 返回result切片最后一个元素作为最大分数。
	return int64(result[len(result)-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
