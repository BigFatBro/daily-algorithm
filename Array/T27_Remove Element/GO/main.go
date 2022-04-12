package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	fmt.Println("case1:", removeElement(nums, val), nums)
}

//双指针法
func removeElement(nums []int, val int) int {
	numsLen := len(nums)
	slow := 0
	for fast := 0; fast < numsLen; fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}

	}
	return slow
}
