package main

import "fmt"

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	for _, v := range s {
		fmt.Printf("%c ", v)
	}

}

func reverseString(s []byte) {
	sLen := len(s)
	for i := 0; i < sLen/2; i++ {
		s[i], s[sLen-i-1] = s[sLen-i-1], s[i]
	}
}
