package main

import "fmt"

func main() {
	s := "abccccdd"
	fmt.Println("s longestPalindrome:", longestPalindrome(s))
}

func longestPalindrome(s string) int {
	//计数
	sLen := len(s)
	cnt := [128]int{}

	for i := 0; i < sLen; i++ {
		cnt[s[i]]++
	}

	ans := 0
	for i := 0; i < 128; i++ {
		ans = ans + cnt[i]/2*2
		if (cnt[i]%2 == 1) && (ans%2 == 0) {
			ans++
		}
	}
	return ans
}
