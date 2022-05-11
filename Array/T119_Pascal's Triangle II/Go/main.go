package main

import "fmt"

func main() {
	rowIndex := 3
	fmt.Println("output_1:", getRow(rowIndex))
	fmt.Println("output_2:", getRowV2(rowIndex))
}

func getRow(rowIndex int) []int {
	ans := make([]int, rowIndex+1, rowIndex+1)
	copyAns := make([]int, rowIndex+1, rowIndex+1)
	ans[0] = 1
	for i := 1; i <= rowIndex; i++ {
		copy(copyAns, ans)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				ans[j] = 1
			} else {
				ans[j] = copyAns[j-1] + copyAns[j]
			}
		}
	}
	return ans
}

// 空间上的优化,只用一个数组：倒着计算
func getRowV2(rowIndex int) []int {
	ans := make([]int, rowIndex+1, rowIndex+1)
	ans[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i; j >= 0; j-- {
			if j == 0 || j == i {
				ans[j] = 1
			} else {
				ans[j] = ans[j-1] + ans[j]
			}
		}
	}
	return ans
}
