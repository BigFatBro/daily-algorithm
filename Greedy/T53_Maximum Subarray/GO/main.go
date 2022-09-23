package main

import "math"

func main() {

}

// 暴力搜索法:超出时间限制
func maxSubArrayV1(nums []int) int {
	res := math.MinInt
	for i := 0; i < len(nums); i++ {
		count := 0
		for j := i; j < len(nums); j++ {
			count += nums[j]
			if count > res {
				res = count
			}
		}
	}
	return res

}

//贪心算法
func maxSubArrayV2(nums []int) int {
	res := math.MinInt
	count := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if count > res {
			res = count
		}
		if count <= 0 {
			count = 0
		}
	}

	return res
}
