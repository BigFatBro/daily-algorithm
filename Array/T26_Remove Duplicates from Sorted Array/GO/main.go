package main

import "fmt"

func main() {
	nums := []int{1, 1, 2}
	fmt.Println("case1:", removeDuplicates(nums), nums)
}

func removeDuplicates(nums []int) int {
	numsMap := map[int]bool{}

	p := 0
	for _, v := range nums {
		_, ok := numsMap[v]
		if !ok {
			nums[p] = v
			p++
			numsMap[v] = true
		}
	}
	return p

}

// 双指针法
func removeDuplicates_v2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
