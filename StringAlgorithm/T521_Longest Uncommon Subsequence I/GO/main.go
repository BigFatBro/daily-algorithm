package main

import "fmt"

func main() {
	a1 := "abc"
	b1 := "bcd"
	fmt.Println("case1:", findLUSlength(a1, b1))
	a2 := "abc"
	b2 := "abc"
	fmt.Println("case2:", findLUSlength(a2, b2))
}

func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}

	if len(a) > len(b) {
		return len(a)
	} else {
		return len(b)
	}

}
