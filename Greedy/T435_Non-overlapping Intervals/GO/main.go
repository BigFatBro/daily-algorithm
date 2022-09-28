package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{
		{1, 2},
		{2, 3},
		{3, 4},
		{1, 3},
	}
	fmt.Println("Output:", eraseOverlapIntervals(intervals))
}

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	// 按右边界排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	//非交叉区间的个数
	count := 1
	//区间分割点
	end := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if end <= intervals[i][0] {
			end = intervals[i][1]
			count++
		}
	}
	return len(intervals) - count

}
