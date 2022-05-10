package main

import "fmt"

func main() {
	numRows := 5
	fmt.Println("Output:", generate(numRows))
}

func generate(numRows int) [][]int {
	ans := [][]int{{1}}

	// 外层循环，循环numRows次
	for i := 1; i < numRows; i++ {
		temp := []int{}
		// 内层循环，循环ans切片中最后一个元素的长度+1次
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				temp = append(temp, 1)
			} else {
				temp = append(temp, ans[len(ans)-1][j-1]+ans[len(ans)-1][j])
			}

		}
		ans = append(ans, temp)
	}
	return ans
}
