package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"Hello", "Alaska", "Dad", "Peace"}
	fmt.Println("case1:", findWords(words))
}

func findWords(words []string) []string {
	ans := []string{}
	t := "12210111011122000010020202"
	for _, v := range words {
		s := strings.ToLower(v)
		sLen := len(s)
		rowindex := t[s[0]-'a']
		i := 0
		for ; i < sLen; i++ {
			if t[s[i]-'a'] != rowindex {
				break
			}
		}

		if i == sLen {
			ans = append(ans, v)
		}
	}

	return ans
}
