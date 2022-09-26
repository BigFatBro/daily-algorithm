package main

import "fmt"

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println("Output:", canCompleteCircuit(gas, cost))
}

// 暴力解法:超出时间限制
func canCompleteCircuit(gas []int, cost []int) int {
	//遍历起点
	for i := 0; i < len(gas); i++ {
		restOil := gas[i] - cost[i]
		index := (i + 1) % len(cost)
		for restOil > 0 && index != i {
			restOil += gas[index] - cost[index]
			index = (index + 1) % len(cost)
		}

		//如果回到了出发点并且油量有剩余则找到了起点
		if restOil >= 0 && index == i {
			return i
		}
	}
	return -1
}

//贪心法
func canCompleteCircuitV2(gas []int, cost []int) int {
	curSum := 0
	totalSum := 0
	start := 0
	for i := 0; i < len(gas); i++ {
		curSum += gas[i] - cost[i]
		totalSum += (gas[i] - cost[i])
		if curSum < 0 {
			start = i + 1
			curSum = 0
		}
	}
	//totalSum大于0说明汽油总量大于消耗量，一定可以跑完全程
	if totalSum >= 0 {
		return start
	} else {
		return -1
	}
}
