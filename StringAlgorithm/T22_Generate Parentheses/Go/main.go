package main

import "fmt"

func main() {
	fmt.Println("Output:", generateParenthesis(3))

}

func generateParenthesis(n int) []string {
	result := []string{}
	path := ""

	var backtrack func(string, int, int)
	backtrack = func(path string, left int, right int) {
		if left > n || right > left {
			return
		}
		if len(path) == 2*n {
			result = append(result, path)
			return
		}
		// 添加一个左括号
		backtrack(path+"(", left+1, right)
		//回溯直接省略了

		//添加一个右括号
		backtrack(path+")", left, right+1)
		//回溯省略了
	}

	backtrack(path, 0, 0)
	return result

}
