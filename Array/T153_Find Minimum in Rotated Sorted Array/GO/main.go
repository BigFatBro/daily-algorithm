package main

import "fmt"

func main() {
	nums := []int{11, 13, 15, 17}
	fmt.Println("Output:", findMin(nums))
}

func findMin(nums []int) int {
	low := 0
	high := len(nums) - 1
	// 二分查找
	for low < high {
		pivot := low + (high-low)/2
		if nums[pivot] < nums[high] {
			high = pivot
		} else {
			low = pivot + 1
		}
	}

	return nums[low]

}
