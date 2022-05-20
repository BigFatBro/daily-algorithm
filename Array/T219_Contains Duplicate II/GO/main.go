package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1}
	k := 3
	fmt.Println("Output1:", containsNearbyDuplicateV1(nums, k))

	nums2 := []int{1, 0, 1, 1}
	k2 := 1
	fmt.Println("Output2:", containsNearbyDuplicateV2(nums2, k2))

}

// 方法1：暴力搜索（遍历每个滑动窗口）
func containsNearbyDuplicateV1(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums) && j-i <= k; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

// 方法2：哈希表
func containsNearbyDuplicateV2(nums []int, k int) bool {
	pos := make(map[int]int)
	for i, num := range nums {
		if p, ok := pos[num]; ok {
			if i-p <= k {
				return true
			}
		}
		pos[num] = i
	}
	return false
}
