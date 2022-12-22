package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println("Output:", nums)

}

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1

	//找A[i] < A[j]
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	//在[j,end)中找一个满足A[i] < A[k]
	if i >= 0 {
		for nums[i] >= nums[k] {
			k--
		}
		// 交换
		nums[i], nums[k] = nums[k], nums[i]
	}

	// 倒置A[j:end]
	for i, j := j, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
