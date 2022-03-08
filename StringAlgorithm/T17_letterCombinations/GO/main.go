package main

import "fmt"

func main() {
	case1 := "23"
	case2 := ""
	case3 := "2"
	fmt.Println("case1:", letterCombinations(case1))
	fmt.Println("case2:", letterCombinations(case2))
	fmt.Println("case3:", letterCombinations(case3))

}

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	var combinations []string
	digitsLen := len(digits)
	if digitsLen == 0 {
		return combinations
	}

	backtrack(&combinations, digits, 0, "")
	return combinations

}

func backtrack(combinations *[]string, digits string, index int, combination string) {
	if index == len(digits) {
		*combinations = append(*combinations, combination)
	} else {
		//subCombination := digits[index : index+1]
		digit := string(digits[index])
		letters := phoneMap[digit]
		for i := 0; i < len(letters); i++ {
			//combination = combination + string(letters[i])
			backtrack(combinations, digits, index+1, combination+string(letters[i]))
			// 回溯时删除刚刚加入的字符
			//golang中字符串不可修改，所以combination + string(letters[i])并不会修改combination
			//combination = combination[:len(combination)-1]
		}

	}
}
