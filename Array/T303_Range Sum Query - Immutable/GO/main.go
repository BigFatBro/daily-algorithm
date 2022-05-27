package main

import "fmt"

type NumArray struct {
	sums []int
}

// 思路：利用前缀
func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums)+1)
	for i, v := range nums {
		sums[i+1] = sums[i] + v

	}
	return NumArray{sums: sums}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sums[right+1] - this.sums[left]
}

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	numArray := Constructor(nums)
	fmt.Println("Output:", numArray.SumRange(0, 2))
}
