package main

import "fmt"

func main() {
	case1 := "aba"
	case2 := "abca"
	case3 := "abc"
	case4 := "bddb"
	fmt.Println("case1:", validPalindrome(case1))
	fmt.Println("case2:", validPalindrome(case2))
	fmt.Println("case3:", validPalindrome(case3))
	fmt.Println("case4:", validPalindrome(case4))

}

func validPalindrome(s string) bool {
	sLen := len(s)
	if sLen <= 2 {
		return true
	}

	i, j := 0, sLen-1
	for {
		if i >= j {
			return true
		}

		if s[i] == s[j] {
			i++
			j--
		} else {
			//判断删除后一个字母后是否是回文串
			return isPalindrome(s[i+1:j+1]) || isPalindrome(s[i:j+1-1])
		}
	}
}

func isPalindrome(s string) bool {
	sLen := len(s)
	i, j := 0, sLen-1
	for {
		if i >= j {
			return true
		} else {
			if s[i] == s[j] {
				i++
				j--
			} else {
				return false
			}
		}
	}
}
