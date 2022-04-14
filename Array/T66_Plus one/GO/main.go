package main

import "fmt"

func main() {
	digits := []int{1, 2, 3, 9}
	fmt.Println("case1:", plusOne(digits))

}

func plusOne(digits []int) []int {
	digitsLen := len(digits)
	for i := digitsLen - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			for j := i + 1; j < digitsLen; j++ {
				digits[j] = 0
			}
			return digits
		}
	}

	digits = make([]int, digitsLen+1)
	digits[0] = 1
	return digits
}
