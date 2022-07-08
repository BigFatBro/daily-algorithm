package main

import "fmt"

func main() {
	s := "ab#c"
	t := "ad#c"
	fmt.Println("case1:", backspaceCompare(s, t))

}

// 最简单的方法是用栈
// 此处的方法是双指针逆序遍历
func backspaceCompare(s string, t string) bool {
	//双指针法
	skipS, skipT := 0, 0
	i, j := len(s)-1, len(t)-1
	for i >= 0 || j >= 0 {
		// 循环直至找到一个确定不会被删除的字符
		for i >= 0 {
			if s[i] == '#' {
				skipS++
				i--
			} else if skipS > 0 {
				//跳过非#字符
				i--
				skipS--
			} else {
				break
			}
		}

		for j >= 0 {
			if t[j] == '#' {
				skipT++
				j--
			} else if skipT > 0 {
				j--
				skipT--
			} else {
				break
			}
		}

		// 比较两个字符串已经确定的字符是否相同
		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			// 一个字符串又找到了一个确定的字符，但是另一个字符串已经遍历完了
			return false

		}
		i--
		j--
	}
	return true
}
