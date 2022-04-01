package main

import "fmt"

func main() {
	s := "Let's take LeetCode contest"
	fmt.Println("case1:", reverseWords(s))
}

func reverseWords(s string) string {
	t := []byte(s)

	i := 0
	j := 0
	for i < len(s) {
		if s[i] == ' ' {
			i++
		} else {
			j = i
			for j < len(s) {
				if s[j] != ' ' {
					j++

				} else {
					sub := t[i:j]
					for m, n := 0, len(sub); m < n/2; m++ {
						sub[m], sub[n-1-m] = sub[n-1-m], sub[m]
					}
					break
				}
			}
			if j == len(s) {
				sub := t[i:]
				for m, n := 0, len(sub); m < n/2; m++ {
					sub[m], sub[n-1-m] = sub[n-1-m], sub[m]
				}
				break
			}
			i = j
		}
	}

	return string(t)

}
