package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("Output:", monotoneIncreasingDigits(332))

}

func monotoneIncreasingDigits(n int) int {
	strNum := strconv.Itoa(n)
	strNumByte := []byte(strNum)
	length := len(strNumByte)
	if length <= 1 {
		return n
	}

	for i := length - 1; i > 0; i-- {
		if strNumByte[i-1] > strNumByte[i] {
			strNumByte[i-1] -= 1
			for j := i; j < length; j++ {
				strNumByte[j] = '9'
			}
		}

	}

	res, _ := strconv.Atoi(string(strNumByte))
	return res

}
