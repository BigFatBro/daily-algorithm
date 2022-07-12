package main

import "fmt"

func main() {

	case1 := 3
	fmt.Println("case1", generateMatrix(case1))

}

func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}

	num := 1
	end := n * n
	//定义边界
	left, right := 0, n-1
	top, bottom := 0, n-1
	for num <= end {
		//遍历上边
		for i := left; i <= right; i++ {
			ans[top][i] = num
			num++
		}
		//遍历完第一行
		top++

		//遍历右边
		for i := top; i <= bottom; i++ {
			ans[i][right] = num
			num++
		}
		right--

		//遍历底边
		for i := right; i >= left; i-- {
			ans[bottom][i] = num
			num++
		}
		bottom--
		//遍历左边
		for i := bottom; i >= top; i-- {
			ans[i][left] = num
			num++
		}
		left++
	}

	return ans

}
