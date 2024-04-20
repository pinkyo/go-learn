package main

func main() {
}

func lengthOfLongestSubsequence(nums []int, target int) int {
	dp := make([]int, target+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = -1
	}
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			if dp[j-nums[i]] != -1 && dp[j-nums[i]]+1 > dp[j] {
				dp[j] = dp[j-nums[i]] + 1
			}
		}
	}

	return dp[target]
}
