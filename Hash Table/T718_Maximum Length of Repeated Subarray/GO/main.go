package main

import "fmt"

func main() {

	nums1 := []int{
		1, 2, 3, 2, 1,
	}
	nums2 := []int{
		3, 2, 1, 4, 7,
	}

	fmt.Println("Output:", findLength(nums1, nums2))

}

func findLength(nums1 []int, nums2 []int) int {
	if len(nums1) < len(nums2) {
		return findMax(nums1, nums2)
	} else {
		return findMax(nums2, nums1)
	}
}

//默认nums1的长度小于nums2
func findMax(nums1, nums2 []int) int {
	maxLength := 0
	m, n := len(nums1), len(nums2)

	// nums1是较短的数组，nums1不动，滑动nums2
	// 第一阶段：nums2与nums1重合部分不断增加
	for i := 1; i < m; i++ {
		maxLength = max(maxLength, maxLen(nums1, 0, nums2, n-i, i))
	}

	// 第二阶段：nums2与nums1重合部分长度不变，nums1位于nums2的“腹中”
	for i := n - m; i >= 0; i-- {
		maxLength = max(maxLength, maxLen(nums1, 0, nums2, i, m))
	}

	//第三阶段： nums2与nums1重合部分不断减小
	for i := m - 1; i > 0; i-- {
		maxLength = max(maxLength, maxLen(nums1, m-i, nums2, 0, i))
	}

	return maxLength

}

//计算nums1中以下标i开始长度为length的子数组 和 nums2中以下标j长度为length的子数组 的最长公共部分的长度
func maxLen(nums1 []int, i int, nums2 []int, j int, length int) int {
	count := 0
	res := 0
	for k := 0; k < length; k++ {
		if nums1[i+k] == nums2[j+k] {
			count++
		} else if count > 0 {
			//公共部分不在连续
			res = max(count, res)
			count = 0
		}
	}
	if count > 0 {
		res = max(count, res)
	}
	return res

}

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}
