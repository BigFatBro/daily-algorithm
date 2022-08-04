package main

import "fmt"

func main() {
	s := "We are happy."
	fmt.Println("Output:", replaceSpace(s))
}

func replaceSpace(s string) string {
	sLen := len(s)
	ans := make([]byte, 0, sLen*3)
	for i := 0; i < sLen; i++ {
		if s[i] == ' ' {
			ans = append(ans, '%', '2', '0')
		} else {
			ans = append(ans, s[i])
		}
	}
	return string(ans)
}
