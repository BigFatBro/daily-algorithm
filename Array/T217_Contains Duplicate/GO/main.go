package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println("Output:", containsDuplicateV1(nums))

}

// 思路：先排序，重复的元素排序后会相邻
func containsDuplicateV1(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false

}

// 哈希表
func containsDuplicateV2(nums []int) bool {
	set := map[int]struct{}{}
	for _, v := range nums {
		if _, ok := set[v]; ok {
			return true
		}
		set[v] = struct{}{}
	}
	return false
}
