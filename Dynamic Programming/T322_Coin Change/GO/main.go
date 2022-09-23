package main

import (
	"fmt"
	"math"
)

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println("output:", coinChange(coins, amount))

}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	//初始化
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
	}

	for i := 0; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] && dp[i-coins[j]] != math.MaxInt {
				if dp[i-coins[j]]+1 < dp[i] {
					dp[i] = dp[i-coins[j]] + 1
				}
			}
		}
	}

	if dp[amount] == math.MaxInt {
		dp[amount] = -1
	}
	return dp[amount]

}
