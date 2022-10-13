package main

import "fmt"

func main() {
	cost1 := []int{
		1, 100, 1, 1, 1, 100, 1, 1, 100, 1,
	}
	cost2 := []int{
		10, 15, 20,
	}
	fmt.Println("Output1:", minCostClimbingStairs(cost1))
	fmt.Println("Output2:", minCostClimbingStairs(cost2))

}

func minCostClimbingStairs(cost []int) int {
	if len(cost) < 2 {
		return cost[len(cost)-1]
	}
	dp := make([]int, len(cost)+1)
	//初始化
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		if dp[i-1] < dp[i-2] {
			dp[i] = dp[i-1] + cost[i]
		} else {
			dp[i] = dp[i-2] + cost[i]
		}
	}
	if dp[len(cost)-1] < dp[len(cost)-2] {
		return dp[len(cost)-1]
	} else {
		return dp[len(cost)-2]
	}

}
