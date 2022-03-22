package main

import "fmt"

func main() {
	s := "abc"
	t := "ahbgdc"
	fmt.Println("case1:", isSubsequence(s, t))
}

func isSubsequence(s string, t string) bool {
	sLen := len(s)
	tLen := len(t)
	i, j := 0, 0
	for i < sLen && j < tLen {
		if s[i] == t[j] {
			i++
		}
		j++

	}

	return i == sLen

}
