package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 6, 7, 7}
	fmt.Println("output:", findSubsequences(nums))
}

func findSubsequences(nums []int) [][]int {
	result := [][]int{}
	path := []int{}

	var backtracking func(int)
	backtracking = func(startIndex int) {
		if len(path) > 1 {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
		}

		//记录本层已经用过的元素
		set := make(map[int]struct{})
		for i := startIndex; i < len(nums); i++ {
			_, exists := set[nums[i]]
			if (len(path) > 0 && nums[i] < path[len(path)-1]) || exists {
				continue
			}
			set[nums[i]] = struct{}{}
			path = append(path, nums[i])
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtracking(0)
	return result

}
