package main

import (
	"fmt"
	"math"
)

func main() {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}

	fmt.Println("case1:", minSubArrayLen(target, nums))

}

func minSubArrayLen(target int, nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	start, end, sum, length := 0, 0, 0, 0
	minLength := math.MaxInt32

	for end = 0; end < len(nums); end++ {
		sum += nums[end]
		for sum >= target {
			length = end - start + 1
			if length < minLength {
				minLength = length
			}
			//起始指针前移,直到窗口内的sum小于target
			sum -= nums[start]
			start++
		}

	}

	if minLength == math.MaxInt32 {
		return 0
	}

	return minLength

}
