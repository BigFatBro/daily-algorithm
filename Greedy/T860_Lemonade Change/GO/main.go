package main

import "fmt"

func main() {
	bills := []int{5, 5, 5, 10, 20}
	fmt.Println("Output:", lemonadeChange(bills))
}

func lemonadeChange(bills []int) bool {
	five, ten, twenty := 0, 0, 0
	for _, bill := range bills {
		if bill == 5 {
			five++
		} else if bill == 10 {
			if five == 0 {
				return false
			} else {
				five--
				ten++
			}
		} else {
			if five > 0 && ten > 0 {
				five--
				ten--
				twenty++
			} else if five >= 3 {
				five -= 3
				twenty++
			} else {
				return false
			}
		}
	}
	return true
}
