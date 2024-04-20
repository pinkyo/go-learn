package main

func main() {
}

func numTriplets(nums1 []int, nums2 []int) int {
	return numTripletsHelper(nums1, nums2) + numTripletsHelper(nums2, nums1)
}

func numTripletsHelper(nums1 []int, nums2 []int) int {
	cntMap := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		for j := i + 1; j < len(nums1); j++ {
			cntMap[nums1[i]*nums1[j]]++
		}
	}
	cntMap2 := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		cntMap2[nums2[i]*nums2[i]]++
	}
	var result int
	for k, v := range cntMap {
		if v2, ok := cntMap2[k]; ok {
			result += v * v2
		}
	}

	return result
}
