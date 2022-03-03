package main

import "fmt"

func main() {
	var case1 = []string{"flower", "flow", "flight"}
	fmt.Println("case1:", longestCommonPrefix(case1))
	var case2 = []string{"dog", "racecar", "car"}
	fmt.Println("case2:", longestCommonPrefix(case2))

}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 || strs == nil {
		return ""
	}
	prefix := strs[0]
	strsLen := len(strs)
	for i := 1; i < strsLen; i++ {
		prefix = longestCommonPrefixOn2Str(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix

}

func longestCommonPrefixOn2Str(str1, str2 string) string {

	var sMinLength int
	if len(str1) < len(str2) {
		sMinLength = len(str1)
	} else {
		sMinLength = len(str2)
	}

	endIndex := 0

	for ; endIndex < sMinLength; endIndex++ {
		if str1[endIndex] != str2[endIndex] {
			break
		}
	}
	return str1[0:endIndex]

}
