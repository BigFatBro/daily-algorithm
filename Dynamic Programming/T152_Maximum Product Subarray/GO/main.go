package main

import (
	"fmt"
	"math"
)

func main() {
	case1 := []int{2, 3, -2, 4}
	//case2 := []int{-2, 0, -1}
	fmt.Println("Output1:", maxProduct(case1))
	//fmt.Println("Output2:", maxProduct(case2))

}

func maxProduct(nums []int) int {
	// 定义dp数组
	maxF := make([]int, len(nums))
	minF := make([]int, len(nums))
	maxF[0] = nums[0]
	minF[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		maxF[i] = max(maxF[i-1]*nums[i], minF[i-1]*nums[i], nums[i])
		minF[i] = min(maxF[i-1]*nums[i], minF[i-1]*nums[i], nums[i])
	}
	ans := math.MinInt
	for i := 0; i < len(maxF); i++ {
		ans = max(ans, maxF[i])
	}
	return ans

}

func max(nums ...int) int {
	maxValue := math.MinInt
	for _, v := range nums {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}

func min(nums ...int) int {
	minValue := math.MaxInt
	for _, v := range nums {
		if v < minValue {
			minValue = v
		}
	}
	return minValue
}
