package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(minNumberOfSeconds(100000, []int{1000000}))
	fmt.Println(canReach(100000, []int{1000000}, 5000050000000000))
}

func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	l, h := 0, workerTimes[0]*mountainHeight*(mountainHeight+1)/2
	for l <= h {
		mid := (l + h) >> 1
		canReach := canReach(mountainHeight, workerTimes, mid)
		if canReach {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}

	return int64(l)
}

func canReach(mountainHeight int, workerTimes []int, seconds int) bool {
	cnt := 0
	for _, wt := range workerTimes {
		// 使用浮点数进行计算
		c := -float64(seconds) * 2.0
		a := float64(wt)
		b := float64(wt)

		// 计算判别式的值
		discriminant := b*b - 4*a*c
		if discriminant < 0 {
			continue // 无实数解，跳过
		}

		// 计算根
		val := (math.Sqrt(discriminant) - b) / (2 * a)

		// 累加结果
		cnt += int(val)
		if cnt >= mountainHeight {
			return true
		}
	}
	return false
}
