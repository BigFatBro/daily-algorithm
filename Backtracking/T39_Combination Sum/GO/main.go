package main

import "fmt"

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println("Output:", combinationSum(candidates, target))

}

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	path := []int{}

	var backtracking func(int, int)
	backtracking = func(sum, startIndex int) {
		if sum > target {
			return
		}
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
		}

		for i := startIndex; i < len(candidates); i++ {
			sum += candidates[i]
			path = append(path, candidates[i])
			backtracking(sum, i)
			sum -= candidates[i]
			path = path[:len(path)-1]
		}
	}

	backtracking(0, 0)
	return result

}
