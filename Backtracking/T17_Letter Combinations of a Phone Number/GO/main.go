package main

import "fmt"

func main() {
	digits := "23"
	fmt.Println("Output:", letterCombinations(digits))
}

func letterCombinations(digits string) []string {
	result := []string{}
	tempStr := ""

	digitsMap := [10]string{
		"",     // 0
		"",     // 1
		"abc",  // 2
		"def",  // 3
		"ghi",  // 4
		"jkl",  // 5
		"mno",  // 6
		"pqrs", // 7
		"tuv",  // 8
		"wxyz", // 9
	}

	var backtracking func(string, int)
	backtracking = func(digits string, index int) {
		if len(tempStr) == len(digits) {
			temp := make([]byte, len(tempStr))
			copy(temp, []byte(tempStr))
			result = append(result, string(temp))
			return
		}
		tmpK := digits[index] - '0' // 将index指向的数字转为int（确定下一个数字）
		letter := digitsMap[tmpK]   // 取数字对应的字符集
		for i := 0; i < len(letter); i++ {
			tempStr = tempStr + string(letter[i])
			backtracking(digits, index+1)
			tempStr = tempStr[:len(tempStr)-1]

		}
	}

	if len(digits) == 0 {
		return result
	}

	backtracking(digits, 0)
	return result
}
