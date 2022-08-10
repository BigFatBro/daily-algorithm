package main

import "fmt"

func main() {

	s := "abbaca"
	fmt.Println("Output:", removeDuplicates(s))

}

func removeDuplicates(s string) string {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 || s[i] != stack[len(stack)-1] {
			stack = append(stack, s[i])
		} else {
			//弹出栈顶元素
			stack = stack[:len(stack)-1]
		}

	}

	return string(stack)

}
