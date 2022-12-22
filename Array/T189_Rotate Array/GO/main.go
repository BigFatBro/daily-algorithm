package main

import "fmt"

func main() {
	nums := []int{
		1, 2, 3, 4, 5, 6, 7,
	}
	rotate(nums, 3)
	fmt.Println("Output:", nums)
}

func rotate(nums []int, k int) {
	if len(nums) <= 1 {
		return
	}
	// 先把整个数组倒置:7,6,5,4,3,2,1
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	//倒置前k个元素:5,6,7,4,3,2,1
	k = k % len(nums)
	for i, j := 0, k-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	//倒置后面的元素:5,6,7,1,2,3,4
	for i, j := k, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
