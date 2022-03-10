package main

import "fmt"

func main() {
	case1 := "Hello World"
	case2 := "   fly me   to   the moon  "
	case3 := "luffy is still joyboy"
	case4 := "a"
	fmt.Println("case1:", lengthOfLastWord(case1))
	fmt.Println("case2:", lengthOfLastWord(case2))
	fmt.Println("case3:", lengthOfLastWord(case3))
	fmt.Println("case4:", lengthOfLastWord(case4))

}

func lengthOfLastWord(s string) int {
	n := len(s)
	lasWordEnd := n - 1

	for ; ; lasWordEnd-- {
		if s[lasWordEnd] != ' ' {
			break
		}
	}

	lasWordBegin := lasWordEnd
	for ; ; lasWordBegin-- {
		if lasWordBegin < 0 || s[lasWordBegin] == ' ' {
			break
		}
	}

	return lasWordEnd - lasWordBegin
}
