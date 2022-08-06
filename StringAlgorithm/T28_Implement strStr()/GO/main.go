package main

import "fmt"

func main() {
	haystack := "hello"
	needle := "ll"
	fmt.Println("Output:", strStr(haystack, needle))

}

//KMP算法
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	next := make([]int, len(needle))
	getNext(&next, needle)

	j := 0
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}

		if j == len(needle) {
			return (i - len(needle) + 1)
		}
	}
	return -1

}

func getNext(next *[]int, s string) {
	j := 0
	(*next)[0] = 0
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
