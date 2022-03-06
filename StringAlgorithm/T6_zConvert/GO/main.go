package main

import "fmt"

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))

}

func convert(s string, numRows int) string {
	n := len(s)
	r := numRows
	if r == 1 || r >= n {
		return s
	}

	//一个周期需要的字符数
	t := r + r - 2

	//需要的矩阵列数
	c := (n + t - 1) / t * (r - 1)

	//创建一个r行c列
	mat := make([][]byte, r)
	for i := range mat {
		mat[i] = make([]byte, c)
	}

	j, k := 0, 0
	for i := 0; i < n; i++ {
		mat[j][k] = byte(s[i])
		if i%t < r-1 {
			//向下移动
			j++
		} else {
			//向右上移动
			j--
			k++
		}
	}

	ans := make([]byte, 0, n)
	for _, row := range mat {
		for _, ch := range row {
			if ch > 0 {
				ans = append(ans, ch)
			}

		}

	}
	return string(ans)

}
