package main

import "fmt"

func main() {
	case1 := "00110011"
	fmt.Println("case1:", countBinarySubstrings(case1))

}

func countBinarySubstrings(s string) int {
	countArray := []int{}

	lastChar := s[0]
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == lastChar {
			count++
		} else {
			countArray = append(countArray, count)
			lastChar = s[i]
			count = 1
		}
	}

	countArray = append(countArray, count)

	ans := 0
	for i := 1; i < len(countArray); i++ {
		if countArray[i] < countArray[i-1] {
			ans = ans + countArray[i]
		} else {
			ans = ans + countArray[i-1]
		}
	}

	return ans

}
