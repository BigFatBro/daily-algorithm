package main

import (
	"fmt"
	"strconv"
)

func main() {
	num1 := "11"
	num2 := "123"
	fmt.Println("case1:", addStrings(num1, num2))

}

func addStrings(num1 string, num2 string) string {
	i := len(num1) - 1
	j := len(num2) - 1
	jinwei := 0
	ans := ""
	for i >= 0 || j >= 0 || jinwei != 0 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}

		res := x + y + jinwei
		ans = strconv.Itoa(res%10) + ans
		jinwei = res / 10
		i--
		j--
	}

	return ans
}
