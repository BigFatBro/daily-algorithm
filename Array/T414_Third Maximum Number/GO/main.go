package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	case1 := []int{3, 2, 1}
	case2 := []int{2, 2, 3, 1}
	case3 := []int{1, 2}
	fmt.Println("case1:", thirdMax(case1))
	fmt.Println("case2:", thirdMax(case2))
	fmt.Println("case3:", thirdMax(case3))
	fmt.Println("case2V2:", thirdMaxV2(case2))

}

//直观的做法：想去重，再排序，最后取第三大的数，没有则取最大的数
// 做完之后发现：不用去重也可以，直接排序，然后遍历，找第个不同的元素
func thirdMax(nums []int) int {
	set := make([]int, 0)
	m := make(map[int]bool)
	for _, v := range nums {
		if _, ok := m[v]; !ok {
			set = append(set, v)
			m[v] = true

		}
	}

	sort.Ints(set)
	if len(set) >= 3 {
		return set[len(set)-3]
	} else {
		return set[len(set)-1]
	}

}

//一次遍历：用三个变量分别表示最大的前三个值，不断更新这三个值
func thirdMaxV2(nums []int) int {
	a, b, c := math.MinInt64, math.MinInt64, math.MinInt64
	for _, v := range nums {
		if v > a {
			a, b, c = v, a, b
		} else if v < a && v > b {
			b, c = v, b
		} else if v < b && v > c {
			c = v
		}
	}
	if c == math.MinInt64 {
		return a
	}
	return c

}
