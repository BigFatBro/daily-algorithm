package main

import "fmt"

func main() {
	s := "Hello, my name is John"
	fmt.Println("case1:", countSegments(s))
}

func countSegments(s string) int {
	sLen := len(s)
	count := 0
	for i := 0; i < sLen; i++ {
		if (s[i] != ' ') && (i == 0 || s[i-1] == ' ') {
			count++
		}
	}
	return count
}
