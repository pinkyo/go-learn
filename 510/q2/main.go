package main

func main() {
}

const mod = 1000000007

func minimumCost(nums []int, k int) int {
	cnt := uint64(0)
	sum := uint64(0)
	for _, num := range nums {
		sum += uint64(num)
		cnt += sum / uint64(k)
		sum %= uint64(k)
	}
	if sum > 0 {
		cnt++
	}
	cnt--
	cnt %= mod
	result := cnt * (cnt + 1) / 2 % mod
	return int(result)
}
