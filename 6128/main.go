package main

func main() {

}

func bestHand(ranks []int, suits []byte) string {
	i := 1
	for ; i < len(suits); i++ {
		if suits[i] != suits[i-1] {
			break
		}
	}
	if i == len(suits) {
		return "Flush"
	}

	cntMap := make(map[int]int)
	for _, rank := range ranks {
		cntMap[rank]++
	}
	var max int
	for _, v := range cntMap {
		if v > max {
			max = v
		}
	}
	switch max {
	case 3, 4, 5:
		return "Three of a Kind"
	case 2:
		return "Pair"
	case 1:
		return "High Card"
	}
	return ""
}
