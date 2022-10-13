package main

import "fmt"

func main() {
	prices := []int{1, 3, 2, 8, 4, 9}
	fee := 2
	fmt.Println("Output:", maxProfit(prices, fee))

}

func maxProfit(prices []int, fee int) int {
	result := 0
	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		}

		if prices[i] > minPrice && prices[i] <= minPrice+fee {
			continue
		}

		if prices[i] > minPrice+fee {
			result += prices[i] - minPrice - fee
			minPrice = prices[i] - fee
		}
	}

	return result
}
