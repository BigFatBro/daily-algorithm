package main

import "fmt"

func main() {
	case1 := "PPALLP"
	case2 := "PPALLL"
	fmt.Println("case1:", checkRecord(case1))
	fmt.Println("case2:", checkRecord(case2))
}

func checkRecord(s string) bool {
	countA := 0
	countL := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			countA++
			countL = 0
			if countA >= 2 {
				return false
			}
		} else if s[i] == 'L' {
			countL++
			if countL >= 3 {
				return false
			}
		} else {
			countL = 0
		}
	}

	return true

}
