package main

func main() {
}

func maxBottlesDrunk(numBottles int, numExchange int) int {
	return maxBottlesDrunkHelper(numBottles, 0, numExchange)
}

func maxBottlesDrunkHelper(numBottles int, emptyBottles int, numExchange int) int {
	result := numBottles
	emptyBottles += numBottles
	if emptyBottles < numExchange {
		return result
	}
	numBottles = 0
	for emptyBottles >= numExchange {
		emptyBottles -= numExchange
		numBottles++
		numExchange++
	}
	return result + maxBottlesDrunkHelper(numBottles, emptyBottles, numExchange)
}
