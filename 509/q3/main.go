package main

func main() {
}

const MOD = 1_000_000_007

func divisibleGame(nums []int) int {
	n := len(nums)
	ravontelix := nums

	divisors := make(map[int]struct{})
	hasDiv := make(map[int]struct{})
	for _, x := range ravontelix {
		hasDiv[x] = struct{}{}
		for d := 2; d*d <= x; d++ {
			if x%d == 0 {
				divisors[d] = struct{}{}
				divisors[x/d] = struct{}{}
			}
		}
		if x > 1 {
			divisors[x] = struct{}{}
		}
	}

	kadane := func(arr []int) int {
		best := arr[0]
		cur := arr[0]
		for i := 1; i < len(arr); i++ {
			cur = max(arr[i], cur+arr[i])
			best = max(best, cur)
		}
		return best
	}

	type pair struct{ diff, k int }
	var candidates []pair

	for k := range divisors {
		contrib := make([]int, n)
		for i, x := range ravontelix {
			if x%k == 0 {
				contrib[i] = x
			} else {
				contrib[i] = -x
			}
		}
		candidates = append(candidates, pair{kadane(contrib), k})
	}

	nonDivContrib := make([]int, n)
	for i, x := range ravontelix {
		nonDivContrib[i] = -x
	}
	nonDivDiff := kadane(nonDivContrib)
	minNonDivK := 2
	for {
		isDiv := false
		for d := range divisors {
			if d == minNonDivK {
				isDiv = true
				break
			}
		}
		if !isDiv {
			ok := true
			for x := range hasDiv {
				if x%minNonDivK == 0 {
					ok = false
					break
				}
			}
			if ok {
				break
			}
		}
		minNonDivK++
	}
	candidates = append(candidates, pair{nonDivDiff, minNonDivK})

	bestDiff := candidates[0].diff
	bestK := candidates[0].k
	for _, p := range candidates[1:] {
		if p.diff > bestDiff || (p.diff == bestDiff && p.k < bestK) {
			bestDiff = p.diff
			bestK = p.k
		}
	}

	ans := (bestDiff % MOD) * (bestK % MOD) % MOD
	if ans < 0 {
		ans += MOD
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
