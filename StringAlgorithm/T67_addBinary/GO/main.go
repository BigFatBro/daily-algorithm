package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "11"
	b := "1"
	fmt.Println(addBinary(a, b))
}

func addBinary(a string, b string) string {

	ans := ""
	//进位
	ca := 0
	i := len(a) - 1
	j := len(b) - 1
	for i >= 0 || j >= 0 {
		sum := ca
		if i >= 0 {
			sum += int(a[i] - '0')
		} else {
			sum += 0
		}

		if j >= 0 {
			sum += int(b[j] - '0')
		} else {
			sum += 0
		}

		ans = strconv.Itoa((sum % 2)) + ans
		ca = sum / 2
		i--
		j--

	}

	if ca == 1 {
		ans = strconv.Itoa(ca) + ans
	}
	return ans

}
