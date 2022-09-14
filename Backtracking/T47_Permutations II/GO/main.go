package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 2}
	fmt.Println("Output:", permuteUnique(nums))
}

func permuteUnique(nums []int) [][]int {
	result := [][]int{}
	path := []int{}
	used := make([]bool, len(nums))
	var backtracking func()
	backtracking = func() {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
		}

		for i := 0; i < len(nums); i++ {
			if (i > 0 && nums[i] == nums[i-1] && used[i-1] == false) || used[i] == true {
				continue
			}

			used[i] = true
			path = append(path, nums[i])
			backtracking()
			path = path[:len(path)-1]
			used[i] = false

		}
	}

	sort.Ints(nums)
	backtracking()
	return result

}
