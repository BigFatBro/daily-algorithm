package main

import "fmt"

func main() {
	var case1 string = "III"
	var case2 string = "IV"
	var case3 string = "IX"
	var case4 string = "LVIII"
	var case5 string = "MCMXCIV"
	fmt.Println("case1:", romanToInt(case1))
	fmt.Println("case2:", romanToInt(case2))
	fmt.Println("case3:", romanToInt(case3))
	fmt.Println("case4:", romanToInt(case4))
	fmt.Println("case5:", romanToInt(case5))

}

func romanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}

	lookUp := map[string]int{
		"CD": 400,
		"CM": 900,
		"XL": 40,
		"XC": 90,
		"IV": 4,
		"IX": 9,
		"M":  1000,
		"D":  500,
		"C":  100,
		"L":  50,
		"X":  10,
		"V":  5,
		"I":  1,
	}

	i := 0
	ans := 0
	sLen := len(s)
	//遍历到倒数第二个字母
	for i < sLen-1 {
		value, ok := lookUp[s[i:i+2]]
		if ok {
			ans = ans + value
			i = i + 2
		} else {
			value, _ := lookUp[s[i:i+1]]
			ans = ans + value
			i = i + 1
		}

	}
	if i == sLen-1 {
		value, _ := lookUp[s[i:i+1]]
		ans = ans + value
	}

	return ans

}
