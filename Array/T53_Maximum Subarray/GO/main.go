package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println("case1:", maxSubArray(nums))

}

// 动态规划算法:如果当前元素+之前元素组成的子数列的sum > 当前元素，则更新当前元素处的sum， 这个sum是当前元素+ 之前元素 组成的子数列的sum
// 否则：当前元素处的sum 就是当前元素的值
// 和贪心算法算法的思路一样
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1]+nums[i] > nums[i] {
			nums[i] = nums[i-1] + nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
