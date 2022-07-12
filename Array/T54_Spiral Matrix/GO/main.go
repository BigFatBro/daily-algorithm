package main

import "fmt"

func main() {

	case1 := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(case1)
	fmt.Println("case1:", spiralOrder(case1))

}

func spiralOrder(matrix [][]int) []int {
	ans := []int{}

	left, right := 0, len(matrix[0])-1
	top, bottom := 0, len(matrix)-1

	total := len(matrix[0]) * len(matrix)

	for len(ans) <= total {
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[top][i])
		}
		top++

		if len(ans) >= total {
			break
		}

		for i := top; i <= bottom; i++ {
			ans = append(ans, matrix[i][right])
		}
		right--

		if len(ans) >= total {
			break
		}

		for i := right; i >= left; i-- {
			ans = append(ans, matrix[bottom][i])
		}
		bottom--

		if len(ans) >= total {
			break
		}

		for i := bottom; i >= top; i-- {
			ans = append(ans, matrix[i][left])
		}

		left++
	}

	return ans

}
