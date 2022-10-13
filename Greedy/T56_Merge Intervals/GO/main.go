package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}

	fmt.Println("Output:", merge(intervals))
}

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	//标记最后一个区间有没有合并
	flag := false
	result := [][]int{}

	length := len(intervals)
	for i := 1; i < length; i++ {
		start := intervals[i-1][0]
		end := intervals[i-1][1]
		for i < length && intervals[i][0] <= end {
			if intervals[i][1] > end {
				end = intervals[i][1]
			}
			if i == length-1 {
				flag = true
			}
			i++
		}
		result = append(result, []int{start, end})
	}
	if flag == false {
		result = append(result, intervals[length-1])
	}

	return result
}
