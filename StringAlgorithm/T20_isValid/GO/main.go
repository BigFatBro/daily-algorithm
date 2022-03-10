package main

import "fmt"

func main() {
	case1 := "()[]{}"
	case2 := "(]"
	case3 := "([)]"
	case4 := "{[]}"

	fmt.Println("case1:", isValid(case1))
	fmt.Println("case2:", isValid(case2))
	fmt.Println("case3:", isValid(case3))
	fmt.Println("case4:", isValid(case4))

}

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	//在golang中用切片实现栈
	stack := []byte{}

	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			//出栈
			stack = stack[:len(stack)-1]

		} else {
			//入栈
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0

}
