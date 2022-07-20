package main

import "fmt"

func main() {

	fmt.Println("n = 19 output =", isHappy(19))
	fmt.Println("n=2 outpu t=", isHappy(2))

}

//获取各个位上数字的平方和
func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum

}

func isHappy(n int) bool {
	mp := map[int]bool{}
	for {
		sum := getSum(n)
		if sum == 1 {
			return true
		}
		if _, ok := mp[sum]; ok {
			return false
		} else {
			mp[sum] = true
		}
		n = sum
	}
}
