package main

import (
	"fmt"
	"strings"
)

func main() {

	words := []string{"Hello", "Alaska", "Dad", "Peace"}
	ans := findWords(words)
	fmt.Println("case1:", ans)

}

func findWords(words []string) []string {
	ans := []string{}
	rowIdx := "12210111011122000010020202"
	for _, word := range words {
		isVaild := true
		lowerWord := strings.ToLower(word)
		idx := rowIdx[lowerWord[0]-'a']
		for i := 1; i < len(lowerWord); i++ {
			if rowIdx[lowerWord[i]-'a'] != idx {
				isVaild = false
				break
			}
		}
		if isVaild {
			ans = append(ans, word)
		}

	}
	return ans
}
