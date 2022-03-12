package main

import "fmt"

func main() {
	//testC()
	case1 := 1
	case2 := 27
	case3 := 701
	fmt.Println("case1:", convertToTitle(case1))
	fmt.Println("case2:", convertToTitle(case2))
	fmt.Println("case3:", convertToTitle(case3))

}

func convertToTitle(columnNumber int) string {
	ans := ""
	for columnNumber > 0 {
		columnNumber--
		ans = string((columnNumber%26)+'A') + ans
		columnNumber = columnNumber / 26
	}
	return ans

}

func testC() {
	b := 2
	fmt.Println(string(b + 'A'))
}
