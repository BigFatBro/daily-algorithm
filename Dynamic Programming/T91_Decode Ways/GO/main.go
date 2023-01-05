package main

import (
	"fmt"
	"strconv"
)

func main() {
	s1 := "226"
	fmt.Println("Output:", numDecodings(s1))
	s2 := "06"
	fmt.Println("Output:", numDecodings(s2))

}

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	legalStr := []string{}
	for i := 1; i < 27; i++ {
		legalStr = append(legalStr, strconv.Itoa(i))
	}

	var isLegal func(string) bool
	isLegal = func(s string) bool {
		for _, v := range legalStr {
			if s == v {
				return true
			}
		}
		return false
	}

	dp := make([]int, len(s))

	//初始条件
	dp[0] = 1
	if s[1] == '0' {
		if isLegal(s[:2]) {
			dp[1] = 1
		} else {
			dp[1] = 0
		}
	} else {
		if isLegal(s[:2]) {
			dp[1] = 2
		} else {
			dp[1] = 1
		}
	}

	for i := 2; i < len(s); i++ {
		if s[i] == '0' {
			if isLegal(s[i-1 : i+1]) {
				dp[i] = dp[i-2]
			} else {
				return 0
			}
		} else {
			if isLegal(s[i-1 : i+1]) {
				dp[i] = dp[i-2] + dp[i-1]
			} else {
				dp[i] = dp[i-1]
			}
		}
	}
	return dp[len(dp)-1]

}
