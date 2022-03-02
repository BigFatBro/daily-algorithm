package main

import "fmt"

func main() {
	// 暴力法寻找最长回文子串
	case1 := "babad"
	fmt.Println("case1:", longestPalindrome(case1))
	case2 := "cbbd"
	fmt.Println("case2:", longestPalindrome(case2))
}

// isPalindrome 判断字符串是否是回文
func isPalindromic(s string) bool {
	sLen := len(s)
	for i := 0; i < sLen; i++ {
		if s[i] != s[sLen-1-i] {
			return false
		}

	}
	return true
}

//longestPalindrome 寻找最长回文子串
func longestPalindrome(s string) string {
	var ans string
	max := 0
	sLen := len(s)
	for i := 0; i < sLen; i++ {
		for j := i + 1; j <= sLen; j++ {
			subStr := s[i:j]
			if isPalindromic(subStr) && len(subStr) > max {
				ans = subStr
				max = len(subStr)
			}
		}
	}
	return ans
}
