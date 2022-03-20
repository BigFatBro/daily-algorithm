package main

import "fmt"

func main() {
	s := "leetcode"
	fmt.Println("case1:", firstUniqChar(s))

}

func firstUniqChar(s string) int {
	sLen := len(s)
	cnt := [26]int{}

	for i := 0; i < sLen; i++ {
		cnt[s[i]-'a']++
	}

	for i := 0; i < sLen; i++ {
		if cnt[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1

}
