package main

import "fmt"

func main() {
	n := 4
	k := 2
	fmt.Println("output:", combine(n, k))

}

func combine(n int, k int) [][]int {
	result := [][]int{}
	path := []int{}

	var backtracking func(int, int, int)
	backtracking = func(n, k, startIndex int) {
		if len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			result = append(result, temp)
			return
		}
		for i := startIndex; i <= n; i++ {
			path = append(path, i)
			backtracking(n, k, i+1)
			path = path[:len(path)-1]
		}
	}
	backtracking(n, k, 1)
	return result

}

//剪枝
func combine_cut(n int, k int) [][]int {
	result := [][]int{}
	path := []int{}
	var backtracking func(int, int, int)
	backtracking = func(n, k, startIndex int) {
		if len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			result = append(result, temp)
			return
		}
		for i := startIndex; i <= n-len(path)+1; i++ {
			path = append(path, i)
			backtracking(n, k, i+1)
			path = path[:len(path)-1]
		}
	}
	backtracking(n, k, 1)
	return result
}
