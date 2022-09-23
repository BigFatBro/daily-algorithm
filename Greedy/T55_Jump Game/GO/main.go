package main

import "fmt"

func main() {
	nums := []int{3, 2, 1, 0, 4}
	fmt.Println("Output:", canJump(nums))
	fmt.Println("Output:", canJumpV2(nums))

}

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	maxIndex := 0
	i := 0
	for i <= maxIndex {
		if i+nums[i] > maxIndex {
			maxIndex = i + nums[i]
		}

		if maxIndex >= len(nums)-1 {
			return true
		}
		i++
	}

	return false

}

//动态规划解法
func canJumpV2(nums []int) bool {
	//定义数组，存储状态
	f := make([]bool, len(nums))
	// 定初始条件
	f[0] = true
	for i := 0; i < len(nums); i++ {
		f[i] = false
		for j := 0; j < i; j++ {
			if f[j] && (nums[j]+j > i) {
				f[i] = true
				break
			}

		}

	}
	return f[len(nums)-1]

}
