package main

import "fmt"

func main() {
	s := "abcd"
	t := "abcde"
	fmt.Printf("case1:%c\n", findTheDifference(s, t))
}

func findTheDifference(s string, t string) byte {
	sLen := len(s)
	cnt := [26]int{}

	for i := 0; i < sLen; i++ {
		cnt[s[i]-'a']++
	}
	for i := 0; ; i++ {
		ch := t[i]
		cnt[ch-'a']--
		if cnt[ch-'a'] < 0 {
			return ch
		}
	}
}
