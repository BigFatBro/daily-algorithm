package main

import (
	"fmt"
	"math"
)

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("output: ", maxProfit(prices))
}

// 方法1：暴力求解o(n^2)，不可取
// 方法2：一次遍历
func maxProfit(prices []int) int {
	minPrice := math.MaxInt
	maxEarn := 0
	for _, v := range prices {
		if v < minPrice {
			minPrice = v
		} else {
			if v-minPrice > maxEarn {
				maxEarn = v - minPrice
			}
		}
	}
	return maxEarn
}
