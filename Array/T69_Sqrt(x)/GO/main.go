package main

import "fmt"

func main() {

	x := 8
	fmt.Println("case1:", mySqrt(x))

}

func mySqrt(x int) int {
	//二分查找法
	left := 0
	right := x
	ans := -1
	mid := left + (right-left)/2
	for left <= right {
		mid = left + (right-left)/2
		if mid*mid <= x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans

}
