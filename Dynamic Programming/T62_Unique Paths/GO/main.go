package main

import "fmt"

func main() {
	fmt.Println("Output:", uniquePaths(3, 7))
}

func uniquePaths(m int, n int) int {
	f := make([][]int, m)
	for i := 0; i < m; i++ {
		f[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				f[i][j] = 1
			} else {
				f[i][j] = f[i-1][j] + f[i][j-1]
			}
		}
	}
	return f[m-1][n-1]
}
