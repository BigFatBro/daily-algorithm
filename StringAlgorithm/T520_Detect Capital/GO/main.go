package main

import "fmt"

func main() {
	case1 := "USA"
	case2 := "FlaG"
	fmt.Println("case1:", detectCapitalUse(case1))
	fmt.Println("case2:", detectCapitalUse(case2))

}

func detectCapitalUse(word string) bool {
	if word[0] < 91 {
		//首字母为大写
		//都大写或者都小写则合格
		if len(word) < 2 {
			return true
		} else {
			falg := 0
			if word[1] < 91 {
				falg = 1
			}
			for i := 1; i < len(word); i++ {
				tempFlag := 0
				if word[i] < 91 {
					tempFlag = 1
				}
				if tempFlag != falg {
					return false
				}
			}
		}

	} else {
		//都小写则合格
		for i := 1; i < len(word); i++ {
			if word[i] < 91 {
				return false
			}
		}

	}

	return true

}
