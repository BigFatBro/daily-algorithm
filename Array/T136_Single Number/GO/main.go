package main

import "fmt"

func main() {
	nums := []int{2, 2, 1}
	fmt.Println("Output:", singleNumber(nums))
}

func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}
