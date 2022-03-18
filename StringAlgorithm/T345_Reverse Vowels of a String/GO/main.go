package main

import "fmt"

func main() {
	s := "hello"
	fmt.Println("case1:", reverseVowels(s))
}

func reverseVowels(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	arr := []byte(s)

	left := 0
	right := n - 1
	for {
		if left >= right {
			break
		}
		if isVowel(s[left]) && isVowel(s[right]) {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
		if !isVowel(arr[left]) {
			left++
		}
		if !isVowel(arr[right]) {
			right--
		}

	}
	return string(arr[:])
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'A' || c == 'e' || c == 'E' || c == 'i' || c == 'I' || c == 'o' || c == 'O' || c == 'u' || c == 'U'

}
