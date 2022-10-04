package main

import "sort"

func main() {
}

func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)
	var idx int
	var result int
	for i := 0; i < len(players); i++ {
		for idx < len(trainers) && players[i] > trainers[idx] {
			idx++
		}
		if idx == len(trainers) {
			break
		}
		result++
		idx++
	}
	return result
}
