package main

import "fmt"

func main() {
	nums := []int{-7, -3, 2, 3, 11}
	fmt.Println("case1:", sortedSquares(nums))

}

func sortedSquares(nums []int) []int {
	//先找到最大的负数的下标
	lastN := -1
	for i := 0; i < len(nums) && nums[i] < 0; i++ {
		lastN = i
	}

	p := lastN + 1
	ans := []int{}
	for lastN >= 0 && p < len(nums) {
		if nums[p] < -nums[lastN] {
			ans = append(ans, nums[p]*nums[p])
			p++
		} else {
			ans = append(ans, nums[lastN]*nums[lastN])
			lastN--
		}
	}
	// 处理剩余部分
	for i := lastN; i >= 0; i-- {
		ans = append(ans, nums[i]*nums[i])
	}

	for i := p; i < len(nums); i++ {
		ans = append(ans, nums[i]*nums[i])
	}

	return ans

}
