package main

import (
	"fmt"
)

func main() {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	bagWeight := 4
	fmt.Println("Output:", GetBestValue(weight, value, bagWeight))

}

func GetBestValue(weight, value []int, bagWeight int) int {
	//二维数组初始化
	dp := make([][]int, len(weight))
	for i := 0; i < len(weight); i++ {
		dp[i] = make([]int, bagWeight+1)
	}

	// 初始化
	for i := weight[0]; i <= bagWeight; i++ {
		dp[0][i] = value[0]
	}

	for i := 1; i < len(weight); i++ {
		for j := 0; j <= bagWeight; j++ {
			if j < weight[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}
	return dp[len(weight)-1][bagWeight]

}

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}
