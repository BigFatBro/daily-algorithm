package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println("Output:", permute(nums))

}

func permute(nums []int) [][]int {
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
			if used[i] == true {
				continue
			}

			used[i] = true
			path = append(path, nums[i])
			backtracking()
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtracking()
	return result

}
