package main

import "fmt"

func main() {
	nums1 := []int{1, 5, 11, 5}
	fmt.Println("Outrput1:", canPartition(nums1))

}

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%2 == 1 {
		return false
	}

	target := sum / 2
	dp := make([]int, 10001)
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}

	if dp[target] == target {
		return true
	}
	return false

}

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}
