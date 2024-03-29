package main

import "fmt"

func main() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println("Output:", minPathSum(grid))

}

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	//二维数组存储
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	//初始状态
	dp[0][0] = grid[0][0]

	// 处理最上边的情况：上一个点只能来自左方
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	//处理最左边的情况:上一个点只能来自上方
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	//状态转移方程：dp[i][j] = min(dp[i-1][j] + grid[i][j], dp[i][j-1] +  grid[i][j])
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j]+grid[i][j], dp[i][j-1]+grid[i][j])
		}
	}

	return dp[m-1][n-1]

}

func min(x, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}
