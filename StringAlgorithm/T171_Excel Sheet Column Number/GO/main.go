package main

import "fmt"

func main() {
	case1 := "A"
	case2 := "AB"
	case3 := "ZY"
	fmt.Println("case1:", titleToNumber(case1))
	fmt.Println("case2:", titleToNumber(case2))
	fmt.Println("case3:", titleToNumber(case3))

}

func titleToNumber(columnTitle string) int {
	ans := 0
	//sLen := len(columnTitle)
	for _, v := range columnTitle {
		ans = ans*26 + int(v-'A') + 1
	}

	return ans

}
