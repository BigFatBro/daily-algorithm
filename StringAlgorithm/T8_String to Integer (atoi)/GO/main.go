package main

import (
	"fmt"
	"math"
)

func main() {
	s1 := "   -42"
	s2 := "-91283472332"
	fmt.Println("Output:", myAtoi(s1))
	fmt.Println("Output:", myAtoi(s2))

}

func myAtoi(s string) int {
	max := int(math.Pow(2, 31) - 1)
	min := int(math.Pow(-2, 31))
	result := 0
	sign := 1
	index := 0
	length := len(s)

	//1.去掉头部空格
	for index < length && s[index] == ' ' {
		index++
	}

	//2.读取符号
	if index < length && s[index] == '-' {
		sign = sign * -1
		index++
	} else if index < length && s[index] == '+' {
		index++
	}

	//3.读取数字
	for index < length {
		curChar := s[index]
		//先判断是不是数字
		if curChar < '0' || curChar > '9' {
			break
		}
		digit := int(curChar - '0')
		//判断是否溢出
		if result > max/10 || (result == max/10 && digit > (max%10)) {
			return max
		}

		if result < min/10 || (result == min/10 && digit > -(min%10)) {
			return min
		}

		// 不溢出时将数字加入结果

		result = result*10 + sign*(digit)
		index++
	}

	return result

}
