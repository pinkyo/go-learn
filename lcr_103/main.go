package main

func main() {
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	dp[0] = 0
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			if dp[i-coin] == -1 {
				continue
			}
			val := dp[i-coin] + 1
			if dp[i] == -1 || dp[i] > val {
				dp[i] = val
			}
		}
	}
	return dp[amount]
}
