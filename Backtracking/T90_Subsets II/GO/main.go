package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 2}
	fmt.Println("Output:", subsetsWithDup(nums))
}

func subsetsWithDup(nums []int) [][]int {
	result := [][]int{}
	path := []int{}
	used := make([]bool, len(nums))

	var backtracking func(int)
	backtracking = func(startIndex int) {
		//获取所有子集既收集树的所有节点
		temp := make([]int, len(path))
		copy(temp, path)
		result = append(result, temp)

		for i := startIndex; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && used[i-1] == false {

				continue
			}

			path = append(path, nums[i])
			used[i] = true
			backtracking(i + 1)
			used[i] = false
			path = path[:len(path)-1]
		}
	}

	//先排序
	sort.Ints(nums)
	backtracking(0)
	return result
}
