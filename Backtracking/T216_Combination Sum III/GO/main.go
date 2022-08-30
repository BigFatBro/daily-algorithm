package main

import "fmt"

func main() {
	k := 3
	n := 7
	fmt.Println("Output:", combinationSum3(k, n))
	fmt.Println("Output:", combinationSum3(3, 9))

}

func combinationSum3(k int, n int) [][]int {
	result := [][]int{}
	path := []int{}
	var backtracking func(int, int, int, int)
	backtracking = func(targetSum, k, sum, startIndex int) {
		if len(path) == k && sum == targetSum {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
		}
		for i := startIndex; i <= 9; i++ {
			path = append(path, i)
			sum += i
			backtracking(targetSum, k, sum, i+1)
			sum -= i
			path = path[:len(path)-1]
		}
	}

	backtracking(n, k, 0, 1)
	return result
}

//剪枝操作

func combinationSum3_Cut(k int, n int) [][]int {
	result := [][]int{}
	path := []int{}
	var backtracking func(int, int, int, int)
	backtracking = func(targetSum, k, sum, startIndex int) {
		if len(path) == k && sum == targetSum {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}
		for i := startIndex; i <= 9-k+len(path)+1; i++ {
			path = append(path, i)
			sum += i
			backtracking(targetSum, k, sum, i+1)
			sum -= i
			path = path[:len(path)-1]
		}
	}

	backtracking(n, k, 0, 1)
	return result
}
