package main

import (
	"fmt"
)

func main() {
	case1 := "abab"
	case2 := "aba"
	fmt.Println("case1:", repeatedSubstringPattern(case1))
	fmt.Println("case2:", repeatedSubstringPattern(case2))

}

//暴力解法：遍历子串结束位置，验证后续字符是否是子串的重复
func repeatedSubstringPattern(s string) bool {

	for i := 1; i*2 < len(s); i++ {
		if len(s)%i == 0 {
			match := true
			for j := i; j < len(s); j++ {
				if s[j] != s[j-i] {
					match = false
					break
				}
			}

			if match {
				return true
			}
		}
	}
	return false
}

// KMP解法
func repeatedSubstringPatternV2(s string) bool {
	sLen := len(s)
	if sLen == 0 {
		return false
	}

	next := make([]int, sLen)
	getNext(&next, s)
	if next[sLen-1] != 0 && sLen%(sLen-(next[sLen-1])) == 0 {
		return true
	}
	return false

}

func getNext(next *[]int, s string) {
	(*next)[0] = 0
	j := 0
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = (*next)[j-1]
		}
		if s[i] == s[j] {
			j++
		}

		(*next)[i] = j
	}
}
