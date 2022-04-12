package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println("case1:", twoSum(nums, target))

}

func twoSum(nums []int, target int) []int {

	numMap := map[int]int{}
	for i, v := range nums {
		value, ok := numMap[target-v]
		if ok {
			return []int{i, value}

		}
		numMap[v] = i
	}
	return []int{}

}
