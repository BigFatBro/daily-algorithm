package main

import "fmt"

func main() {
	s := "abcdefg"
	k := 2
	fmt.Println("Output:", reverseLeftWords(s, k))
}

func reverseLeftWords(s string, n int) string {
	b := []byte(s)
	//1.先反转前n个字符
	reverse(&b, 0, n-1)
	//2.反转n到末尾的字符
	reverse(&b, n, len(s)-1)
	//3.反转整个字符串
	reverse(&b, 0, len(s)-1)
	return string(b)

}

func reverse(b *[]byte, left, right int) {
	for left < right {
		(*b)[left], (*b)[right] = (*b)[right], (*b)[left]
		left++
		right--
	}
}
