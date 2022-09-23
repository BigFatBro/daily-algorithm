package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Output:", solveNQueens(4))

}

func solveNQueens(n int) [][]string {
	result := [][]string{}
	chessboard := make([][]string, n)
	for i := 0; i < n; i++ {
		chessboard[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			chessboard[i][j] = "."
		}
	}

	var backtracking func(int)
	backtracking = func(row int) {
		if row == n {
			temp := make([]string, n)
			for i, rowStr := range chessboard {
				temp[i] = strings.Join(rowStr, "")
			}
			result = append(result, temp)
			return
		}

		for col := 0; col < n; col++ {
			if isValid(row, col, n, chessboard) {
				chessboard[row][col] = "Q"
				backtracking(row + 1)
				chessboard[row][col] = "."
			}
		}
	}

	backtracking(0)
	return result

}

func isValid(row, col, n int, chessboard [][]string) bool {

	//检查列
	for i := 0; i < row; i++ {
		if chessboard[i][col] == "Q" {
			return false
		}
	}

	// 检查45度角
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}

	//检查135度角
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}

	return true

}
