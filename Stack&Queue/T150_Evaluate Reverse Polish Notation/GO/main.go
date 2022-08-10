package main

import (
	"fmt"
	"strconv"
)

func main() {

	tokens := []string{"2", "1", "+", "3", "*"}
	fmt.Println("Output:", evalRPN(tokens))

}

func evalRPN(tokens []string) int {
	tLen := len(tokens)
	stack := make([]int, 0)
	for i := 0; i < tLen; i++ {
		if tokens[i] == "+" {
			//取出栈顶两个元素，计算结果，并将结果压入栈中
			num1, num2 := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, num1+num2)

		} else if tokens[i] == "-" {
			num1, num2 := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, num2-num1)

		} else if tokens[i] == "*" {
			num1, num2 := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, num1*num2)

		} else if tokens[i] == "/" {
			num1, num2 := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, num2/num1)

		} else {
			num, _ := strconv.Atoi(tokens[i])
			stack = append(stack, num)
		}
	}
	return stack[len(stack)-1]

}
