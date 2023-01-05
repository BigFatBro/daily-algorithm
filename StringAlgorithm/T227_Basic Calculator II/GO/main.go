package main

import "fmt"

func main() {
	s1 := "3+2*2"
	s2 := " 3/2 "
	s3 := " 3+5 / 2 "
	fmt.Println("Output:", calculate(s1))
	fmt.Println("Output:", calculate(s2))
	fmt.Println("Output:", calculate(s3))

}

func calculate(s string) int {
	digitStack := []int{}
	symbolStack := []byte{}
	sLen := len(s)
	i := 0
	for i < sLen {
		//把数字和符号拆开，加入到栈中
		if s[i] == '+' || s[i] == '-' {
			symbolStack = append(symbolStack, s[i])
			i++

		} else if s[i] == '*' || s[i] == '/' {
			//寻找下一个操作数，与栈顶元素进行乘除，然后将结果入栈
			num := 0
			end := i + 1
			for ; end < sLen && s[end] != '+' && s[end] != '-' && s[end] != '*' && s[end] != '/'; end++ {
				if s[end] == ' ' {
					continue
				} else {
					num = num*10 + int(s[end]-'0')
				}
			}
			//出栈
			top := digitStack[len(digitStack)-1]
			digitStack = digitStack[:len(digitStack)-1]
			//与栈顶元素进行乘除,并将结果入栈
			if s[i] == '*' {
				digitStack = append(digitStack, top*num)
			} else {
				digitStack = append(digitStack, top/num)
			}
			i = end

		} else if s[i] == ' ' {
			i++
		} else {
			//s[i]是数字时，找到完整数字然后入栈
			num := 0
			end := i
			for ; end < sLen && s[end] != '+' && s[end] != '-' && s[end] != '*' && s[end] != '/'; end++ {
				if s[end] == ' ' {
					continue
				} else {
					num = num*10 + int(s[end]-'0')
				}
			}
			digitStack = append(digitStack, num)
			i = end
		}
	}

	// 符号栈中只有加减号了，处理加减法
	result := digitStack[0]
	for i := 0; i < len(symbolStack); i++ {
		if symbolStack[i] == '+' {
			result = result + digitStack[i+1]
		} else {
			result = result - digitStack[i+1]
		}

	}

	return result

}
