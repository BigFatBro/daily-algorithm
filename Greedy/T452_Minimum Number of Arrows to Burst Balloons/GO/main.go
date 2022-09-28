package main

import "sort"

func main() {
}

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	result := 1

	for i := 1; i < len(points); i++ {
		if points[i][0] > points[i-1][1] {
			result++
		} else {
			//和上一个有重合,取二者中较小的右边界1
			if points[i-1][1] < points[i][1] {
				points[i][1] = points[i-1][1]
			}

		}
	}
	return result
}
