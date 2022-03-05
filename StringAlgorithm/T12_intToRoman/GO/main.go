package main

import "fmt"

var valueSymbols = []struct {
	value  int
	symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func main() {

	case1 := 3
	case2 := 4
	case3 := 9
	case4 := 58
	case5 := 1994
	fmt.Println("case1:", intToRoman(case1))
	fmt.Println("case2:", intToRoman(case2))
	fmt.Println("case3:", intToRoman(case3))
	fmt.Println("case4:", intToRoman(case4))
	fmt.Println("case5:", intToRoman(case5))

}

func intToRoman(num int) string {

	ansRoman := []byte{}
	for _, v := range valueSymbols {
		for num >= v.value {
			num = num - v.value
			ansRoman = append(ansRoman, v.symbol...)
		}
		if num == 0 {
			break
		}
	}
	return string(ansRoman)

}
