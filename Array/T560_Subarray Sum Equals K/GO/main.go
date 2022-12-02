package main

import "fmt"

func main() {

	nums := []int{
		1, 1, 1,
	}
	fmt.Println("Output:", subarraySum(nums, 2))

}

// 前缀和+哈希表
func subarraySum(nums []int, k int) int {
	n := len(nums)
	result := 0
	// 哈希表查询前缀和为key的元素个数
	m := map[int]int{}
	m[0] = 1
	//前缀和
	preSum := 0
	for i := 0; i < n; i++ {
		preSum += nums[i]
		if _, ok := m[preSum-k]; ok {
			result += m[preSum-k]
		}
		m[preSum] += 1
	}
	return result

}
