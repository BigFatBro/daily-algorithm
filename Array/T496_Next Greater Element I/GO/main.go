package main

import "fmt"

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	fmt.Println("case1:", nextGreaterElement(nums1, nums2))

}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	table := map[int]int{}
	stack := []int{}
	nums2Len := len(nums2)
	for i := nums2Len - 1; i >= 0; i-- {
		// 先出栈
		for len(stack) > 0 && stack[len(stack)-1] <= nums2[i] {
			stack = stack[:len(stack)-1]
		}
		//记录到哈希表
		if len(stack) > 0 {
			table[nums2[i]] = stack[len(stack)-1]
		} else {
			table[nums2[i]] = -1
		}
		stack = append(stack, nums2[i])
	}

	res := []int{}
	for _, v := range nums1 {
		res = append(res, table[v])
	}
	return res
}
