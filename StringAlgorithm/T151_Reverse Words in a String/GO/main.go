package main

import (
	"fmt"
)

func main() {
	s := " the  sky is blue "
	fmt.Println("Output:", reverseWords(s))

}

func reverseWords(s string) string {
	//1.移除多余空格
	tmp := []byte(s)
	//1.1删除前面和后面的空格
	i, j := 0, 0

	for i = 0; i < len(tmp); i++ {
		if tmp[i] != ' ' {
			break
		}
	}

	for j = len(tmp) - 1; j >= 0; j-- {
		if tmp[j] != ' ' {
			break
		}
	}
	tmp = tmp[i : j+1]
	//1.2 删除中间多余的空格(快慢指针法)
	fast, slow := 0, 0
	for ; fast < len(tmp); fast++ {
		if fast-1 > 0 && tmp[fast-1] == tmp[fast] && tmp[fast] == ' ' {
			continue
		}
		tmp[slow] = tmp[fast]
		slow++
	}

	tmp = tmp[:slow]

	//2.将整个字符串反转
	reverse(&tmp, 0, len(tmp)-1)

	//3. 将每个单词反转
	i = 0
	for i < len(tmp) {
		j = i
		for ; j < len(tmp) && tmp[j] != ' '; j++ {

		}
		reverse(&tmp, i, j-1)
		i = j
		i++
	}
	return string(tmp)
}

func reverse(b *[]byte, left, right int) {
	for left < right {
		(*b)[left], (*b)[right] = (*b)[right], (*b)[left]
		left++
		right--
	}
}
