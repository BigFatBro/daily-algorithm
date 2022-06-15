package main

import "fmt"

func main() {
	grid := [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}
	fmt.Println("case1:", islandPerimeter(grid))
}

func islandPerimeter(grid [][]int) int {
	rowNum := len(grid)
	colNum := len(grid[0])

	res := 0

	for i := 0; i < rowNum; i++ {
		for j := 0; j < colNum; j++ {
			if grid[i][j] == 1 {
				// 检查四个方向
				if i == 0 || grid[i-1][j] == 0 {
					// 上方
					res++
				}
				if i == rowNum-1 || grid[i+1][j] == 0 {
					// 下方
					res++
				}
				if j == 0 || grid[i][j-1] == 0 {
					//左方
					res++
				}
				if j == colNum-1 || grid[i][j+1] == 0 {
					// 右方
					res++
				}
			}
		}
	}
	return res

}
