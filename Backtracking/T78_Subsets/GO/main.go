package main

import "fmt"

func main() {

	nums := []int{1, 2, 3}
	fmt.Println("output:", subsets(nums))

}

func subsets(nums []int) [][]int {
	result := [][]int{}
	path := []int{}
	var backtracking func(int)
	backtracking = func(startIndex int) {
		temp := make([]int, len(path))
		copy(temp, path)
		result = append(result, temp) //收集树中的每一个节点
		if startIndex >= len(nums) {
			return
		}
		for i := startIndex; i < len(nums); i++ {
			path = append(path, nums[i])
			backtracking(i + 1)
			path = path[:len(path)-1]
		}

	}
	backtracking(0)
	return result

}
