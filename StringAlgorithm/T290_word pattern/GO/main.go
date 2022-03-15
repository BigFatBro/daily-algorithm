package main

import (
	"fmt"
	"strings"
)

func main() {
	pattern := "abba"
	s := "dog cat cat dog"
	fmt.Println("case1:", wordPattern(pattern, s))

}

func wordPattern(pattern string, s string) bool {
	sArr := strings.Split(s, " ")
	if len(pattern) != len(sArr) {
		return false
	}
	patternMap := make(map[byte]string)
	sMap := make(map[string]byte)

	patternLen := len(pattern)

	for i := 0; i < patternLen; i++ {
		patternMapValue, pOK := patternMap[pattern[i]]
		sMapValue, sOK := sMap[sArr[i]]
		if (pOK && patternMapValue != sArr[i]) || (sOK && sMapValue != pattern[i]) {
			return false
		} else {
			patternMap[pattern[i]] = sArr[i]
			sMap[sArr[i]] = pattern[i]
		}
	}
	return true

}
