package main

import "fmt"

func main() {

	nums := []int{3, 2, 4}
	target := 6
	fmt.Println("Case1:", twoSum(nums, target))

}

func twoSum(nums []int, target int) []int {
	mp := map[int]int{}

	for i, n := range nums {
		if v, ok := mp[target-n]; ok {
			return []int{v, i}
		}

		mp[n] = i
	}
	return []int{}

}
