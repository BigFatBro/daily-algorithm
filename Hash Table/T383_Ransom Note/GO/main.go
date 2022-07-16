package main

import "fmt"

func main() {
	ransomNote := "aa"
	magazine := "ab"
	fmt.Println("case1:", canConstruct(ransomNote, magazine))
}

func canConstruct(ransomNote string, magazine string) bool {
	record := [26]int{}
	for _, v := range magazine {
		record[v-'a']++
	}

	for _, v := range ransomNote {
		record[v-'a']--
	}

	for _, v := range record {
		if v < 0 {
			return false
		}
	}
	return true

}
