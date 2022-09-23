package main

import "fmt"

func main() {
	nums := []int{1, 7, 4, 9, 2, 5}
	fmt.Println("Output:", wiggleMaxLength(nums))

}

func wiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	res := 1
	curDiff := 0
	preDiff := 0
	for i := 0; i < len(nums)-1; i++ {
		curDiff = nums[i+1] - nums[i]
		if (curDiff > 0 && preDiff <= 0) || (preDiff >= 0 && curDiff < 0) {
			res++
			preDiff = curDiff
		}
	}

	return res
}
