package main

import "fmt"

func main() {

	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	fmt.Println("output:", search(nums, target))

}

func search(nums []int, target int) int {

	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	// 特殊的二分查找
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if target == nums[mid] {
			return mid
		}

		// 判断两侧哪一侧是有序的
		if nums[0] <= nums[mid] {
			//左侧有序
			if target >= nums[l] && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}

		} else {
			//右侧有序
			if target > nums[mid] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	// 未找到目标值
	return -1
}
