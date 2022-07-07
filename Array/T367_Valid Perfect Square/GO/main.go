package main

import "fmt"

func main() {
	fmt.Println("case1: num = 16, output = ", isPerfectSquare(16))
	fmt.Println("case1: num = 15, output = ", isPerfectSquare(15))
}

func isPerfectSquare(num int) bool {
	// 二分查找法
	left, right := 0, num
	mid := -1
	for left <= right {
		mid = left + (right-left)/2
		if mid*mid == num {
			return true
		} else if mid*mid < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false

}
