package main

func main() {
}

// findWinningPlayer Function to find the index of the winning player
func findWinningPlayer(skills []int, k int) int {
	// maxIdx is the index of the current highest skilled player
	maxIdx := 0
	// Traverse all players from beginning to end
	for i := 1; i < len(skills); i++ {
		// If the current player's skill is higher than the previously found highest skilled player
		if skills[i] > skills[maxIdx] {
			// Update maxIdx to the index of the current player
			maxIdx = i
		}
		// Calculate the number of consecutive games won by the current player
		cnt := i - maxIdx
		// If the current highest skilled player is not the first player
		if maxIdx > 0 {
			// Add one to the number of consecutive games won to include the first player
			cnt += 1
		}
		// If the current player has won consecutive games equal to or exceeding k
		if cnt >= k {
			// Exit the loop, the current player is the winner
			break
		}
	}
	// Return the index of the winning player
	return maxIdx
}
