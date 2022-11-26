package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
	fmt.Println("Output:", matrix)

}

func rotate(matrix [][]int) {
	//先水平翻转
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j], matrix[n-1-i][j] = matrix[n-1-i][j], matrix[i][j]
		}
	}
	//再沿对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

}
