package main

import "fmt"

func main() {
	fmt.Println(distance([]int{1, 3, 1, 1, 2}))
}

func distance(nums []int) []int64 {
	posMap := make(map[int][]int)
	for i, num := range nums {
		posMap[num] = append(posMap[num], i)
	}
	sumMap := make(map[int][]int)
	for val, postList := range posMap {
		sum := make([]int, len(postList))
		sum[0] = postList[0]
		for i := 1; i < len(postList); i++ {
			sum[i] = sum[i-1] + postList[i]
		}
		sumMap[val] = sum
	}
	result := make([]int64, len(nums))
	for i, num := range nums {
		pos := binarySearch(posMap[num], i)
		sumArr := sumMap[num]
		temp := sumArr[len(sumArr)-1] - sumArr[pos]
		if pos > 0 {
			temp -= sumArr[pos-1]
		}
		temp -= (len(posMap[num]) - pos - 1) * i
		temp += pos * i
		result[i] = int64(temp)
	}

	return result
}

func binarySearch(arr []int, target int) int {
	l, h := 0, len(arr)-1
	for l <= h {
		mid := (l + h) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
