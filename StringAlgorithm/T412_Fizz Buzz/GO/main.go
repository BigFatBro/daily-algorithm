package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 15
	ans := fizzBuzz(n)
	for _, v := range ans {
		fmt.Println(v)
	}

}

func fizzBuzz(n int) []string {
	var ans = []string{}
	for i := 1; i <= n; i++ {
		if (i%3 == 0) && (i%5 == 0) {
			ans = append(ans, "FizzBuzz")
		} else if i%3 == 0 {
			ans = append(ans, "Fizz")
		} else if i%5 == 0 {
			ans = append(ans, "Buzz")
		} else {
			ans = append(ans, strconv.Itoa(i))
		}
	}

	return ans

}
