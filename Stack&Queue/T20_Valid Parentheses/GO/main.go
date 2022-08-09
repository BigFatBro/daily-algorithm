package main

import "fmt"

func main() {
	s := "()[]{}"
	fmt.Println("Output:", isValid(s))
}

func isValid(s string) bool {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, ')')
		} else if s[i] == '[' {
			stack = append(stack, ']')
		} else if s[i] == '{' {
			stack = append(stack, '}')
		} else if len(stack) == 0 || stack[len(stack)-1] != s[i] {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
