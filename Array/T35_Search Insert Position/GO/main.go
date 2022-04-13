package main

import "fmt"

func main() {
	nums1 := []int{1, 3, 5, 6}
	target1 := 7

	nums2 := []int{1, 3, 5, 6}
	target2 := 2

	fmt.Println("case1:", searchInsert(nums1, target1))
	fmt.Println("case2", searchInsert(nums2, target2))

}

func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if target <= nums[i] {
			return i
		}
	}

	return len(nums)
}
