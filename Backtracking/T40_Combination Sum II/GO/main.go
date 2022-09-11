package main

import (
	"fmt"
	"sort"
)

func main() {

	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println("Output:", combinationSum2(candidates, target))

}

func combinationSum2(candidates []int, target int) [][]int {
	result := [][]int{}
	path := []int{}

	var backtracking func(int, int, []bool)
	backtracking = func(sum, startIndex int, used []bool) {
		if sum > target {
			return
		}
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
		}

		for i := startIndex; i < len(candidates) && sum+candidates[i] <= target; i++ {
			if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == false {
				continue
			}

			sum += candidates[i]
			path = append(path, candidates[i])
			used[i] = true
			backtracking(sum, i+1, used)
			used[i] = false
			sum -= candidates[i]
			path = path[:len(path)-1]
		}
	}

	used := make([]bool, len(candidates))
	sort.Ints(candidates)
	backtracking(0, 0, used)
	return result
}
