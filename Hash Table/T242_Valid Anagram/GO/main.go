package main

import "fmt"

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram(s, t))

}

func isAnagram(s string, t string) bool {
	record := [26]int{}
	for _, v := range s {
		record[v-'a']++
	}
	for _, v := range t {
		record[v-'a']--
	}
	return record == [26]int{}

}
