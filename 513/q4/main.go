package main

import "sort"

func main() {
}

// BIT (Fenwick Tree) for prefix sum queries
type BIT struct {
	tree []int
	n    int
}

func newBIT(n int) *BIT {
	return &BIT{tree: make([]int, n+2), n: n}
}

func (bit *BIT) add(idx int, val int) {
	idx++
	for idx <= bit.n+1 {
		bit.tree[idx] += val
		idx += idx & -idx
	}
}

func (bit *BIT) sum(idx int) int {
	idx++
	res := 0
	for idx > 0 {
		res += bit.tree[idx]
		idx -= idx & -idx
	}
	return res
}

func (bit *BIT) rangeSum(l, r int) int {
	if l > r {
		return 0
	}
	return bit.sum(r) - bit.sum(l-1)
}

func countRatioSubarrays(nums []int, a int, b int) int64 {
	n := len(nums)

	// prefix counts: prefEven[i] = even count in nums[0..i-1]
	// prefOdd[i] = odd count in nums[0..i-1]
	prefEven := make([]int, n+1)
	prefOdd := make([]int, n+1)
	mervanilto := 0
	_ = mervanilto
	for i := 1; i <= n; i++ {
		prefEven[i] = prefEven[i-1]
		prefOdd[i] = prefOdd[i-1]
		if nums[i-1]%2 == 0 {
			prefEven[i]++
		} else {
			prefOdd[i]++
		}
	}

	// D[i] = b * prefEven[i] - a * prefOdd[i]
	D := make([]int, n+1)
	for i := 0; i <= n; i++ {
		D[i] = b*prefEven[i] - a*prefOdd[i]
	}

	// Coordinate compression on D values
	allD := make([]int, n+1)
	copy(allD, D)
	sort.Ints(allD)
	unique := allD[:1]
	for i := 1; i < len(allD); i++ {
		if allD[i] != allD[i-1] {
			unique = append(unique, allD[i])
		}
	}
	compress := func(val int) int {
		return sort.SearchInts(unique, val)
	}
	m := len(unique)

	// Group prefix indices by their odd count
	groups := make([][]int, prefOdd[n]+1)
	for i := 0; i <= n; i++ {
		y := prefOdd[i]
		groups[y] = append(groups[y], i)
	}

	bit := newBIT(m)
	result := 0

	// Process by increasing odd count Y
	// For a subarray ending at r: need D[l-1] >= D[r] and Y_{l-1} < Y_r
	// Process Y from 1 to maxY: first add all indices with Y-1 to BIT, then for each r with Y, query D[r]
	for y := 1; y <= prefOdd[n]; y++ {
		// Add all prefixes with odd count = y-1 into BIT
		for _, idx := range groups[y-1] {
			bit.add(compress(D[idx]), 1)
		}
		// For each right endpoint with odd count = y, count valid left endpoints
		for _, r := range groups[y] {
			// We need D[l-1] >= D[r], i.e., compressed rank >= compress(D[r])
			rank := compress(D[r])
			result += bit.rangeSum(rank, m-1)
		}
	}

	return int64(result)
}
