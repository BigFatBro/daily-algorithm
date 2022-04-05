package main

import "fmt"

func main() {
	case1 := "UDD"
	fmt.Println("case1:", judgeCircle(case1))
}

func judgeCircle(moves string) bool {
	x := 0
	y := 0
	for i := 0; i < len(moves); i++ {
		if moves[i] == 'U' {
			y++
		} else if moves[i] == 'D' {
			y--
		} else if moves[i] == 'L' {
			x--
		} else {
			x++
		}
	}

	if x == 0 && y == 0 {
		return true
	} else {
		return false
	}

}
