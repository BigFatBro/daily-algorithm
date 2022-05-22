package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{3, 0, 1}
	nums2 := []int{0, 1}
	fmt.Println("Output_V1_1:", missingNumberV1(nums1))
	fmt.Println("Output_V1_2:", missingNumberV1(nums2))
	fmt.Println("Output_V2_1:", missingNumberV2(nums1))
	fmt.Println("Output_v2_2:", missingNumberV2(nums2))
	fmt.Println("Output_V3_1:", missingNumberV3(nums1))
	fmt.Println("Output_v3_2:", missingNumberV3(nums2))
}

// 方法1：哈希表，两遍遍历
func missingNumberV1(nums []int) int {
	m := map[int]bool{}
	for _, v := range nums {
		m[v] = true
	}

	var i int
	for i = 0; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}
	return i
}

// 方法2：排序
func missingNumberV2(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	return len(nums)
}

// 方法3：等差数列求和：Sn=n*a1+n(n-1)d/2或Sn=n(a1+an)/2
func missingNumberV3(nums []int) int {
	numsLen := len(nums)
	stdSum := (numsLen + 1) * (0 + numsLen) / 2
	realSum := 0
	for _, v := range nums {
		realSum += v
	}
	return stdSum - realSum

}
