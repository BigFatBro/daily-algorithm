package main

import "fmt"

func main() {

	s := "LOvElY"

	fmt.Println("case1:", toLowerCase(s))

}

func toLowerCase(s string) string {
	b := []byte(s)

	for i := 0; i < len(b); i++ {
		if b[i] >= 65 && b[i] <= 90 {
			b[i] = b[i] + 32
		}
	}

	return string(b)
}
