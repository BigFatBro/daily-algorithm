package main

import "fmt"

func main() {

	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	fmt.Println("case1:", search(nums, target))

}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	midlde := left + (right-left)/2
	for {
		if left > right {
			break
		}

		midlde = left + (right-left)/2
		if nums[midlde] > target {
			right = midlde - 1
		} else if nums[midlde] < target {
			left = midlde + 1
		} else {
			return midlde
		}
	}

	return -1

}
