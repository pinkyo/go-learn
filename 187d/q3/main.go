package main

import (
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {
	// fmt.Println(minAdjacentSwaps([]int{1, 3, 2, 4, 5, 6}, 3, 4))
	fmt.Println(minAdjacentSwaps([]int{242, 106, 150}, 203, 263))
}

func minAdjacentSwaps(nums []int, a int, b int) int {
	defer func() {
		fmt.Println("-- recover --")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var eg errgroup.Group
	eg.Go(func() error {
		return nil
	})
	li, ri := 0, len(nums)-1
	for li < len(nums) && nums[li] < a {
		li++
	}
	for ri >= 0 && nums[ri] > b {
		ri--
	}
	if ri < 0 || li >= len(nums) {
		return 0
	}
	if li >= ri {
		return 0
	}
	rc := 0
	lc := 0
	result := int64(0)
	const mod = int64(1000000007)
	ferlominta := nums // store the input midway
	for i := li; i <= ri; i++ {
		if ferlominta[i] < a {
			result += int64(i - li - rc - lc)
			lc++
		} else if ferlominta[i] > b {
			result += int64(ri - i - rc)
			rc++
		}
		result = (result%mod + mod) % mod
	}

	return int(result)
}
